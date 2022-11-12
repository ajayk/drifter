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

func TestNamespace(t *testing.T) {
	testCases := []struct {
		name          string
		namespace     []runtime.Object
		expectSuccess bool
	}{
		{
			name: "existing_ingress_class_nginx_should_pass",
			namespace: []runtime.Object{
				&corev1.Namespace{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "kube-test-ns",
					},
				},
			},
			expectSuccess: true,
		},

		{
			name: "existing_ingress_class_nginx_should_pass",
			namespace: []runtime.Object{
				&corev1.Namespace{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "kube-test-ns-not-existing",
					},
				},
			},
			expectSuccess: false,
		},
	}

	drifter := model.Drifter{
		Kubernetes: model.Kubernetes{
			Namespaces: []model.Namespace{{
				Name: "kube-test-ns",
			},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.namespace...)
			va := CheckNamespaces(drifter,
				fakeClientSet, context.Background())
			if va && test.expectSuccess {
				t.Fatalf("unexpected error getting namespace:")
			}
		})
	}
}
