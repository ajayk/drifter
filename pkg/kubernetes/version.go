// Package kubernetes
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

package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	"k8s.io/client-go/kubernetes"
	"log"
)

func CheckVersion(clusterConfig model.Drifter, client kubernetes.Interface, ctx context.Context) bool {
	hasDrifts := false
	if len(clusterConfig.Kubernetes.Version) > 0 {
		version, err := client.Discovery().ServerVersion()
		if err != nil {
			log.Fatal("Unable to list server version", err)
		}

		if version.String() != clusterConfig.Kubernetes.Version {
			hasDrifts = true
			log.Printf("Expected %s version of kubernetes found %s", clusterConfig.Kubernetes.Version, version.String())
		}

	}
	return hasDrifts
}
