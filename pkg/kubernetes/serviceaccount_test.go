package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestServiceAccounts(t *testing.T) {
	testCases := []struct {
		name          string
		namespace     []runtime.Object
		expectSuccess bool
	}{
		{
			name: "existing_service_account_should_pass",
			namespace: []runtime.Object{
				&corev1.ServiceAccount{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "anetd-secret",
						Namespace: "kube-system",
					},
				},
			},
			expectSuccess: true,
		},

		{
			name: "non_existing_service_account_should_fail",
			namespace: []runtime.Object{
				&corev1.ServiceAccount{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "anetd2",
						Namespace: "kube-system",
					},
				},
			},
			expectSuccess: false,
		},
	}

	drifter := model.Drifter{
		Kubernetes: model.Kubernetes{
			ServiceAccounts: []model.ServiceAccounts{{
				NameSpace: "kube-system",
				Names:     []string{"anetd-secret"},
			}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.namespace...)
			va := CheckServiceAccounts(drifter,
				fakeClientSet, context.Background())
			if va && test.expectSuccess {
				t.Fatalf("unexpected error getting namespace:")
			}
		})
	}
}
