// Package helm
// Copyright © 2022 Ajay K <ajaykemparaj@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package helm

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	helmstoragev3 "helm.sh/helm/v3/pkg/storage"
	"k8s.io/client-go/kubernetes"

	"helm.sh/helm/v3/pkg/release"
	driverv3 "helm.sh/helm/v3/pkg/storage/driver"

	"log"
)

func CheckHelmComponents(clusterConfig model.Drifter, client kubernetes.Interface, ctx context.Context) bool {
	hasDrifts := false
	if len(clusterConfig.Helm.Components) > 0 {
		hs := driverv3.NewSecrets(client.CoreV1().Secrets(""))
		helmClient := helmstoragev3.Init(hs)
		releases, err := helmClient.ListDeployed()
		if err != nil {
			log.Fatal("Unable to list helm releases ", err)
		}
		installedHelmComponents := make(map[string]*release.Release)
		for _, release := range releases {
			installedHelmComponents[release.Name] = release
		}
		for _, s := range clusterConfig.Helm.Components {
			if release, ok := installedHelmComponents[s.Name]; ok {
				if release.Info.Status.String() == "deployed" {
					if s.Version != "" {
						if s.Version == release.Chart.Metadata.Version {

						} else {
							log.Printf("Mismatched helm chart %s expected %s found %s  ", s.Name, s.Version, release.Chart.Metadata.Version)
							hasDrifts = true
						}
					}

					if s.AppVersion != "" {
						if release.Chart.AppVersion() != s.AppVersion {
							log.Println("Need", s.AppVersion)
							log.Printf("App Version mismatch for %s , %s\n", s.Name, release.Chart.AppVersion())
							hasDrifts = true
						}
					}
				} else {
					log.Println("Missing Helm Deployment ", s.Name, release.Info.Status)
					hasDrifts = true
				}
				if s.Version != "" {
				}
			} else {
				hasDrifts = true
				log.Println("Missing Helm Deployment ", s.Name)
			}
		}
	}
	return hasDrifts
}
