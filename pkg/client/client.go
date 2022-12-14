// Package client
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

package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/flowcontrol"
	"k8s.io/klog/v2"
	"math"
)

func GetKubernetesClient(configPath string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		klog.Errorf("Error reading the kubeconfig file: %v", err)
		return nil, err
	}
	config.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(math.MaxFloat32, math.MaxInt)

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Errorf("Error creating a k8s client: %v", err)
		return nil, err
	}

	return client, nil
}
