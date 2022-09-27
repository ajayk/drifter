// Package cmd
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
package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/ajayk/drifter/pkg/client"
	"github.com/ajayk/drifter/pkg/helm"
	"github.com/ajayk/drifter/pkg/kubernetes"
	"github.com/ajayk/drifter/pkg/model"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var config string
var checkFile string

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "checks for any configuration drifts in a cluster",
	Long: `Running check figures out any configuration drifts in a cluster

Pass the expectation yaml file and the kube config file to the check command 
`,
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()
		if fflags.Changed("kubeconfig") == false { // check if the flag "git-org" is set
			fmt.Println("no kubeconfig file specified") // If not, we'll let the user know
			return                                      // and return
		}
		if fflags.Changed("check-file") == false { // check if the flag "harbor-file" is set
			fmt.Println("check file was not specified") // If not, we'll let the user know
			return                                      // and return
		}
		if checkFile == "" {
			errors.New("check yaml not specified")
		}
		f, err := os.Open(checkFile)
		if err != nil {
			errors.New("unable to open Cluster Yaml file")

		}
		defer f.Close()
		driftConfig := model.Drifter{}
		if err := yaml.NewDecoder(f).Decode(&driftConfig); err != nil {
			fmt.Println("Unable to parse yaml")
			fmt.Println(err)
			os.Exit(1)
		}

		kubernetesClientSet, err := client.GetKubernetesClient(config)
		if err != nil {
			fmt.Println("Unable to obtain kubernetes client")
		}

		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
		ctx, cancel := context.WithCancel(context.Background())
		var once sync.Once
		defer once.Do(cancel)
		go func() {
			for {
				select {
				case sig := <-sigs:
					klog.V(1).Infof("Received a stop signal: %v", sig)
					once.Do(cancel)
				case <-ctx.Done():
					klog.V(1).Info("Cancelled context and exiting program")
					return
				}
			}
		}()

		storageDrifts := kubernetes.CheckStorageClasses(driftConfig, kubernetesClientSet, ctx)
		namesPaceDrifts := kubernetes.CheckNamespaces(driftConfig, kubernetesClientSet, ctx)
		ingressDrifts := kubernetes.CheckIngressClass(driftConfig, kubernetesClientSet, ctx)
		helmDrifts := helm.CheckHelmComponents(driftConfig, kubernetesClientSet, ctx)
		if ingressDrifts || namesPaceDrifts || storageDrifts || helmDrifts {
			os.Exit(2)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.
	checkCmd.PersistentFlags().StringVarP(&config, "kubeconfig", "k", "", "full path to kubeconfig file")
	checkCmd.PersistentFlags().StringVarP(&checkFile, "check-file", "c", "", "path to cluster expectation yaml")

}
