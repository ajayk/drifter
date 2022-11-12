package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	mapset "github.com/deckarep/golang-set/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
	"log"
)

func CheckConfigMaps(clusterConfig model.Drifter, client kubernetes.Interface, ctx context.Context) bool {
	hasDrifts := false
	var driftCount = 0
	if len(clusterConfig.Kubernetes.ConfigMaps) > 0 {

		for _, ds := range clusterConfig.Kubernetes.ConfigMaps {

			dsList, err := client.CoreV1().ConfigMaps(ds.NameSpace).List(ctx, v1.ListOptions{})
			if err != nil {
				log.Fatal("Unable to get ConfigMaps ", err)
			}
			if len(dsList.Items) == 0 {
				driftCount++
				// When the namespace has no ds items, we are checking for a ds item then there is a drift
				hasDrifts = true

			} else {
				fetched := mapset.NewSet[string]()
				required := mapset.NewSet[string]()

				for _, d := range ds.Names {
					required.Add(d)
				}

				for _, d := range dsList.Items {
					fetched.Add(d.Name)
				}

				result := fetched.Intersect(required)
				if result.Equal(required) {
					//no op
				} else {
					driftCount++
					hasDrifts = true
					diffs := required.Difference(result)
					for _, d := range diffs.ToSlice() {
						log.Printf("Missing ConfigMap %s\n", d)
					}
				}
			}
		}
	}
	if driftCount > 0 {
		hasDrifts = true
	}
	return hasDrifts
}
