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
	networkingV1 "k8s.io/api/networking/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func CheckIngressClass(clusterConfig model.Drifter, client *kubernetes.Clientset, ctx context.Context) bool {
	hasDrifts := false
	if len(clusterConfig.Kubernetes.Ingress.IngressClasses) > 0 {

		ingressList, err := client.NetworkingV1().IngressClasses().List(ctx, v1.ListOptions{})

		if err != nil {
			log.Fatal("Unable to get Ingress classes ", err)
		}
		installedIngress := make(map[string]networkingV1.IngressClass)
		for _, ic := range ingressList.Items {
			installedIngress[ic.Name] = ic
		}
		for _, expectSc := range clusterConfig.Kubernetes.Ingress.IngressClasses {
			if _, ok := installedIngress[expectSc]; ok {
				//do something here

			} else {
				log.Printf("Missing ingress class: %s\n", expectSc)
				hasDrifts = true
			}
		}
	}
	return hasDrifts
}
