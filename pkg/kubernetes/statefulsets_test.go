package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestStatefulSets(t *testing.T) {
	testCases := []struct {
		name          string
		namespace     []runtime.Object
		expectSuccess bool
	}{
		{
			name: "existing_statefulset_should_pass",
			namespace: []runtime.Object{
				&appsv1.StatefulSet{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "anetd",
						Namespace: "kube-system",
					},
					Spec:   appsv1.StatefulSetSpec{},
					Status: appsv1.StatefulSetStatus{},
				},
			},
			expectSuccess: true,
		},

		{
			name: "existing_deployment_should_fail",
			namespace: []runtime.Object{
				&appsv1.StatefulSet{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "anetd2",
						Namespace: "kube-system",
					},
					Spec:   appsv1.StatefulSetSpec{},
					Status: appsv1.StatefulSetStatus{},
				},
			},
			expectSuccess: false,
		},
	}

	drifter := model.Drifter{
		Kubernetes: model.Kubernetes{
			StatefulSets: []model.StatefulSets{{
				NameSpace: "kube-system",
				Names:     []string{"anetd"},
			}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.namespace...)
			va := CheckStatefulSets(drifter,
				fakeClientSet, context.Background())
			if va && test.expectSuccess {
				t.Fatalf("unexpected error getting namespace:")
			}
		})
	}
}
