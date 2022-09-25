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
	"fmt"
	"github.com/ajayk/drifter/pkg/model"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/kube"
	"helm.sh/helm/v3/pkg/release"
	"log"
)

func CheckHelmComponents(clusterConfig model.Drifter, kubeconfig string) {
	if len(clusterConfig.Helm.Components) > 0 {
		actionConfig := new(action.Configuration)
		err := actionConfig.Init(kube.GetConfig(kubeconfig, "", ""), "", "", log.Printf)
		if err != nil {
			log.Fatal(err)
		}

		releases, err := action.NewList(actionConfig).Run()
		if err != nil {
			log.Fatal(err)
		}
		installedHelmComponents := make(map[string]*release.Release)
		for _, release := range releases {
			installedHelmComponents[release.Name] = release
		}

		for _, s := range clusterConfig.Helm.Components {
			if release, ok := installedHelmComponents[s.Name]; ok {
				if release.Info.Status.String() == "deployed" {
					if s.AppVersion != "" {
						if release.Chart.AppVersion() != s.AppVersion {
							fmt.Println("Need", s.AppVersion)
							fmt.Printf("App Version mismatch for %s , %s\n", s.Name, release.Chart.AppVersion())
						}
					}

				} else {
					fmt.Println("Missing Deployment ", s.Name, release.Info.Status)
				}
				if s.Version != "" {
				}
			} else {
				fmt.Println("Missing Deployment ", s.Name)
			}
		}
	}
}
