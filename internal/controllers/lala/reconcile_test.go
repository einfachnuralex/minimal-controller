package lala

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	lalav1alpha1 "github.com/einfachnuralex/minimal-controller/pkg/apis/lala/v1alpha1"
)

var _ = Describe("Lala Controller", func() {
	Context("When reconciling a resource", func() {
		const resourceName = "test-resource"

		ctx := context.Background()

		typeNamespacedName := types.NamespacedName{
			Name:      resourceName,
			Namespace: "default", // TODO(user):Modify as needed
		}
		lala := &lalav1alpha1.Lala{}

		BeforeEach(func() {
			By("creating the custom resource")
			err := k8sClient.Get(ctx, typeNamespacedName, lala)
			if err != nil && errors.IsNotFound(err) {
				resource := &lalav1alpha1.Lala{
					ObjectMeta: metav1.ObjectMeta{
						Name:      resourceName,
						Namespace: "default",
					},
					// TODO(user): Specify other spec details if needed.
				}
				Expect(k8sClient.Create(ctx, resource)).To(Succeed())
			}
		})

		It("Check for resource", func() {
			By("Checking the created resource")
			Eventually(func(g Gomega) {
				g.Expect(k8sClient.Get(ctx, typeNamespacedName, lala)).To(Succeed())
			}, timeout, interval).To(Succeed())

			By("By checking if the status is updated correctly")
			Eventually(func(g Gomega) {
				g.Expect(k8sClient.Get(ctx, typeNamespacedName, lala)).To(Succeed())
				g.Expect(lala.Status.InstanceID).To(Equal(instance_name))
			}, timeout, interval).Should(Succeed())
		})

		AfterEach(func() {
			resource := &lalav1alpha1.Lala{}
			err := k8sClient.Get(ctx, typeNamespacedName, resource)
			Expect(err).NotTo(HaveOccurred())

			By("Cleanup the specific resource")
			Expect(k8sClient.Delete(ctx, resource)).To(Succeed())
		})
	})
})
