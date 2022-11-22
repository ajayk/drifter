package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	v1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestCLusterRole(t *testing.T) {
	testCases := []struct {
		name          string
		namespace     []runtime.Object
		expectSuccess bool
	}{
		{
			name: "existing_clusterrole_should_pass",
			namespace: []runtime.Object{
				&v1.ClusterRole{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "anetd-cm",
					},
				},
			},
			expectSuccess: true,
		},

		{
			name: "non_existing_cr_should_fail",
			namespace: []runtime.Object{
				&v1.ClusterRole{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "anetd2",
					},
				},
			},
			expectSuccess: false,
		},
	}

	drifter := model.Drifter{
		Kubernetes: model.Kubernetes{
			ClusterRoles: []model.ClusterRoles{{
				Names: []string{"anetd-cm"},
			}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.namespace...)
			va := CheckClusterRoles(drifter,
				fakeClientSet, context.Background())
			if va && test.expectSuccess {
				t.Fatalf("unexpected error getting namespace:")
			}
		})
	}
}
