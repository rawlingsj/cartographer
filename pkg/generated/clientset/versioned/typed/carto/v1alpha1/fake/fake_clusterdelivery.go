/*
Copyright 2021 VMware

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/vmware-tanzu/cartographer/pkg/apis/carto/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterDeliveries implements ClusterDeliveryInterface
type FakeClusterDeliveries struct {
	Fake *FakeCartoV1alpha1
}

var clusterdeliveriesResource = schema.GroupVersionResource{Group: "carto.run", Version: "v1alpha1", Resource: "clusterdeliveries"}

var clusterdeliveriesKind = schema.GroupVersionKind{Group: "carto.run", Version: "v1alpha1", Kind: "ClusterDelivery"}

// Get takes name of the clusterDelivery, and returns the corresponding clusterDelivery object, and an error if there is any.
func (c *FakeClusterDeliveries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterDelivery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterdeliveriesResource, name), &v1alpha1.ClusterDelivery{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDelivery), err
}

// List takes label and field selectors, and returns the list of ClusterDeliveries that match those selectors.
func (c *FakeClusterDeliveries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterDeliveryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterdeliveriesResource, clusterdeliveriesKind, opts), &v1alpha1.ClusterDeliveryList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterDeliveryList{ListMeta: obj.(*v1alpha1.ClusterDeliveryList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterDeliveryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterDeliveries.
func (c *FakeClusterDeliveries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterdeliveriesResource, opts))
}

// Create takes the representation of a clusterDelivery and creates it.  Returns the server's representation of the clusterDelivery, and an error, if there is any.
func (c *FakeClusterDeliveries) Create(ctx context.Context, clusterDelivery *v1alpha1.ClusterDelivery, opts v1.CreateOptions) (result *v1alpha1.ClusterDelivery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterdeliveriesResource, clusterDelivery), &v1alpha1.ClusterDelivery{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDelivery), err
}

// Update takes the representation of a clusterDelivery and updates it. Returns the server's representation of the clusterDelivery, and an error, if there is any.
func (c *FakeClusterDeliveries) Update(ctx context.Context, clusterDelivery *v1alpha1.ClusterDelivery, opts v1.UpdateOptions) (result *v1alpha1.ClusterDelivery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterdeliveriesResource, clusterDelivery), &v1alpha1.ClusterDelivery{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDelivery), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterDeliveries) UpdateStatus(ctx context.Context, clusterDelivery *v1alpha1.ClusterDelivery, opts v1.UpdateOptions) (*v1alpha1.ClusterDelivery, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterdeliveriesResource, "status", clusterDelivery), &v1alpha1.ClusterDelivery{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDelivery), err
}

// Delete takes name of the clusterDelivery and deletes it. Returns an error if one occurs.
func (c *FakeClusterDeliveries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterdeliveriesResource, name), &v1alpha1.ClusterDelivery{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterDeliveries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterdeliveriesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterDeliveryList{})
	return err
}

// Patch applies the patch and returns the patched clusterDelivery.
func (c *FakeClusterDeliveries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterDelivery, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterdeliveriesResource, name, pt, data, subresources...), &v1alpha1.ClusterDelivery{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterDelivery), err
}
