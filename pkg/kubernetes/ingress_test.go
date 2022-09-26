package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	networkingV1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestIngress(t *testing.T) {
	testCases := []struct {
		name          string
		ingress       []runtime.Object
		expectSuccess bool
	}{
		{
			name: "existing_ingress_class_nginx_should_pass",
			ingress: []runtime.Object{
				&networkingV1.IngressClass{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "nginx",
					},
					Spec: networkingV1.IngressClassSpec{},
				},
			},
			expectSuccess: true,
		},
		{
			name: "existing_ingress_class_nginx_should_fail",
			ingress: []runtime.Object{
				&networkingV1.IngressClass{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "nginx2",
					},
					Spec: networkingV1.IngressClassSpec{},
				},
			},
			expectSuccess: false,
		},
	}

	drifter := model.Drifter{
		Helm: model.K8sHelm{},
		Kubernetes: model.Kubernetes{
			Ingress: model.K8sIngress{IngressClasses: []string{"nginx"}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.ingress...)
			va := CheckIngressClass(drifter,
				fakeClientSet, context.Background())
			if va && test.expectSuccess {
				t.Fatalf("unexpected error getting ingress:")
			}
		})
	}
}
