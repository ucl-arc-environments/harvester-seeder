package controllers

import (
	"fmt"
	bmaasv1alpha1 "github.com/harvester/bmaas/pkg/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var _ = Describe("AddressPool controller tests", func() {
	var a *bmaasv1alpha1.AddressPool

	BeforeEach(func() {
		a = &bmaasv1alpha1.AddressPool{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "sample",
				Namespace: "default",
			},
			Spec: bmaasv1alpha1.AddressSpec{
				CIDR:    "192.168.1.1/29",
				Gateway: "192.168.1.7",
			},
		}

		Eventually(func() error {
			return k8sClient.Create(ctx, a)
		}, "30s", "5s").ShouldNot(HaveOccurred())
	})

	It("check addresspool is ready ", func() {
		Eventually(func() error {
			obj := &bmaasv1alpha1.AddressPool{}
			err := k8sClient.Get(ctx, types.NamespacedName{Namespace: a.Namespace, Name: a.Name}, obj)
			if err != nil {
				return err
			}

			if obj.Status.Status != bmaasv1alpha1.PoolReady {
				return fmt.Errorf("waiting for pool to be ready. current status is %s", obj.Status.Status)
			}
			return nil
		}, "30s", "5s").ShouldNot(HaveOccurred())
	})

	AfterEach(func() {
		Eventually(func() error {
			err := k8sClient.Delete(ctx, a)
			return err
		}).ShouldNot(HaveOccurred())
	})
})