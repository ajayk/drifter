package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	mapset "github.com/deckarep/golang-set/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

func CheckStatefulSets(clusterConfig model.Drifter, client kubernetes.Interface, ctx context.Context) bool {
	hasDrifts := false
	var driftCount = 0
	if len(clusterConfig.Kubernetes.StatefulSets) > 0 {
		for _, sts := range clusterConfig.Kubernetes.StatefulSets {

			dsList, err := client.AppsV1().StatefulSets(sts.NameSpace).List(ctx, v1.ListOptions{})
			if err != nil {
				log.Fatal("Unable to get Stateful-sets ", err)
			}
			if len(dsList.Items) == 0 {
				driftCount++
				// When the namespace has no ds items, we are checking for a ds item then there is a drift
				hasDrifts = true

			} else {
				fetched := mapset.NewSet[string]()
				required := mapset.NewSet[string]()

				for _, d := range sts.Names {
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
						log.Printf("Missing Statefulset %s\n", d)
					}
				}

			}
		}
	}

	return hasDrifts

}
