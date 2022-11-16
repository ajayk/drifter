// Package model
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
package model

type Drifter struct {
	Helm       K8sHelm    `yaml:"helm"`
	Kubernetes Kubernetes `yaml:"kubernetes"`
}

type Kubernetes struct {
	Namespaces      []Namespace       `yaml:"namespaces"`
	DaemonSets      []DaemonSets      `yaml:"daemonsets"`
	Deployments     []Deployments     `yaml:"deployments"`
	StatefulSets    []StatefulSets    `yaml:"statefulsets"`
	Storage         K8sStorage        `yaml:"storage"`
	Ingress         K8sIngress        `yaml:"ingress"`
	Secrets         []Secrets         `yaml:"secrets"`
	ServiceAccounts []ServiceAccounts `yaml:"serviceAccounts"`

	ConfigMaps []ConfigMaps `yaml:"configmaps"`
}
type ServiceAccounts struct {
	NameSpace string   `yaml:"namespace"`
	Names     []string `yaml:"names"`
}

type ConfigMaps struct {
	NameSpace string   `yaml:"namespace"`
	Names     []string `yaml:"names"`
}

type Secrets struct {
	NameSpace string   `yaml:"namespace"`
	Names     []string `yaml:"names"`
}
type Deployments struct {
	NameSpace string   `yaml:"namespace"`
	Names     []string `yaml:"names"`
}

type StatefulSets struct {
	NameSpace string   `yaml:"namespace"`
	Names     []string `yaml:"names"`
}

type DaemonSets struct {
	NameSpace string   `yaml:"namespace"`
	Names     []string `yaml:"names"`
}

type Namespace struct {
	Name string `yaml:"name"`
}

type K8sIngress struct {
	IngressClasses []string `yaml:"classes"`
}

type K8sHelm struct {
	Components []HelmComponent `yaml:"components"`
}
type HelmComponent struct {
	Name       string `yaml:"name"`
	Version    string `yaml:"version"`
	AppVersion string `yaml:"appVersion"`
}

type K8sStorage struct {
	StorageClasses []string `yaml:"classes"`
}
