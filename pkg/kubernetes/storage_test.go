package kubernetes

import (
	"context"
	"github.com/ajayk/drifter/pkg/model"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestStorage(t *testing.T) {
	testCases := []struct {
		name          string
		storageClass  []runtime.Object
		expectSuccess bool
	}{
		{
			name: "existing_ingress_class_nginx_should_pass",
			storageClass: []runtime.Object{
				&storagev1.StorageClass{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "ebs-sc",
					},
				},
			},
			expectSuccess: true,
		},

		{
			name: "existing_ingress_class_nginx_should_pass",
			storageClass: []runtime.Object{
				&storagev1.StorageClass{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name: "ebs-sc-1",
					},
				},
			},
			expectSuccess: false,
		},
	}

	drifter := model.Drifter{
		Kubernetes: model.Kubernetes{
			Storage: model.K8sStorage{StorageClasses: []string{"ebs-sc"}},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			fakeClientSet := fake.NewSimpleClientset(test.storageClass...)
			va := CheckStorageClasses(drifter,
				fakeClientSet, context.Background())
			if va && test.expectSuccess {
				t.Fatalf("unexpected error getting namespace:")
			}
		})
	}
}
