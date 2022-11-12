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

func TestConfigMaps(t *testing.T) {
	testCases := []struct {
		name          string
		namespace     []runtime.Object
		expectSuccess bool
	}{
		{
			name: "existing_cm_should_pass",
			namespace: []runtime.Object{
				&corev1.ConfigMap{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "anetd-cm",
						Namespace: "kube-system",
					},
				},
			},
			expectSuccess: true,
		},

		{
			name: "non_existing_cm_should_fail",
			namespace: []runtime.Object{
				&corev1.ConfigMap{
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
			ConfigMaps: []model.ConfigMaps{{
				NameSpace: "kube-system",
				Names:     []string{"anetd-cm"},
			}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.namespace...)
			va := CheckConfigMaps(drifter,
				fakeClientSet, context.Background())
			if va && test.expectSuccess {
				t.Fatalf("unexpected error getting namespace:")
			}
		})
	}
}
