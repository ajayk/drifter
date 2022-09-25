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
	"fmt"
	"github.com/ajayk/drifter/pkg/model"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func CheckNamespaces(clusterConfig model.Drifter, client *kubernetes.Clientset, ctx context.Context) {
	if len(clusterConfig.Kubernetes.Namespaces) > 0 {

		nsList, err := client.CoreV1().Namespaces().List(ctx, v1.ListOptions{})
		if err != nil {
			log.Fatal(err)
		}
		namespacesMap := make(map[string]corev1.Namespace)
		for _, ns := range nsList.Items {
			namespacesMap[ns.Name] = ns
		}

		for _, expectNs := range clusterConfig.Kubernetes.Namespaces {
			if _, ok := namespacesMap[expectNs]; ok {
				// Do Nothing
			} else {
				fmt.Println("Missing ", expectNs)
			}
		}

	}

}
