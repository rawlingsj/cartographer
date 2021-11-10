// Copyright 2021 VMware
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vmware-tanzu/cartographer/pkg/apis/carto/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Delivery Validation", func() {
	Describe("#Create", func() {
		Context("Well formed delivery", func() {
			var delivery *v1alpha1.ClusterDelivery

			BeforeEach(func() {
				delivery = &v1alpha1.ClusterDelivery{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "delivery-resource",
						Namespace: "default",
					},
					Spec: v1alpha1.ClusterDeliverySpec{
						Resources: []v1alpha1.ClusterDeliveryResource{
							{
								Name: "source-provider",
								TemplateRef: v1alpha1.DeliveryClusterTemplateReference{
									Kind: "ClusterSourceTemplate",
									Name: "source-template",
								},
							},
							{
								Name: "other-source-provider",
								TemplateRef: v1alpha1.DeliveryClusterTemplateReference{
									Kind: "ClusterSourceTemplate",
									Name: "source-template",
								},
							},
						},
					},
				}
			})

			It("does not return an error", func() {
				Expect(delivery.ValidateCreate()).NotTo(HaveOccurred())
			})

		})

		Context("Duplicate resource names", func() {
			var delivery *v1alpha1.ClusterDelivery

			BeforeEach(func() {
				delivery = &v1alpha1.ClusterDelivery{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "delivery-resource",
						Namespace: "default",
					},
					Spec: v1alpha1.ClusterDeliverySpec{
						Resources: []v1alpha1.ClusterDeliveryResource{
							{
								Name: "source-provider",
								TemplateRef: v1alpha1.DeliveryClusterTemplateReference{
									Kind: "ClusterSourceTemplate",
									Name: "source-template",
								},
							},
							{
								Name: "source-provider",
								TemplateRef: v1alpha1.DeliveryClusterTemplateReference{
									Kind: "ClusterSourceTemplate",
									Name: "source-template",
								},
							},
						},
					},
				}
			})

			It("does not return an error", func() {
				err := delivery.ValidateCreate()
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring(`spec.resources[1].name "source-provider" cannot appear twice`)))
			})

		})
	})

	Describe("#Update", func() {
		var (
			previousDelivery *v1alpha1.ClusterDelivery
			newDelivery      *v1alpha1.ClusterDelivery
		)

		BeforeEach(func() {
			previousDelivery = &v1alpha1.ClusterDelivery{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "delivery-resource",
					Namespace: "default",
				},
				Spec: v1alpha1.ClusterDeliverySpec{
					Resources: []v1alpha1.ClusterDeliveryResource{
						{
							Name: "source-provider",
							TemplateRef: v1alpha1.DeliveryClusterTemplateReference{
								Kind: "ClusterSourceTemplate",
								Name: "source-template",
							},
						},
						{
							Name: "other-source-provider",
							TemplateRef: v1alpha1.DeliveryClusterTemplateReference{
								Kind: "ClusterSourceTemplate",
								Name: "source-template",
							},
						},
					},
				},
			}
		})

		Context("with a valid change", func() {
			BeforeEach(func() {
				newDelivery = previousDelivery.DeepCopy()
				newDelivery.Spec.Resources = newDelivery.Spec.Resources[:1]
			})

			It("does not return an error", func() {
				Expect(newDelivery.ValidateUpdate(previousDelivery)).NotTo(HaveOccurred())
			})
		})

		Context("Duplicate resource names", func() {
			BeforeEach(func() {
				newDelivery = previousDelivery.DeepCopy()
				newDelivery.Spec.Resources = append(newDelivery.Spec.Resources, v1alpha1.ClusterDeliveryResource{
					Name: "other-source-provider",
					TemplateRef: v1alpha1.DeliveryClusterTemplateReference{
						Kind: "ClusterSourceTemplate",
						Name: "source-template",
					},
				})
			})

			It("does not return an error", func() {
				err := newDelivery.ValidateUpdate(previousDelivery)
				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(ContainSubstring(`spec.resources[2].name "other-source-provider" cannot appear twice`)))
			})

		})

	})

	Describe("#Delete", func() {

	})
})
