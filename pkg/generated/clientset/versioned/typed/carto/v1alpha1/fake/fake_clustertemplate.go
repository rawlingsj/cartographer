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

// FakeClusterTemplates implements ClusterTemplateInterface
type FakeClusterTemplates struct {
	Fake *FakeCartoV1alpha1
	ns   string
}

var clustertemplatesResource = schema.GroupVersionResource{Group: "carto.run", Version: "v1alpha1", Resource: "clustertemplates"}

var clustertemplatesKind = schema.GroupVersionKind{Group: "carto.run", Version: "v1alpha1", Kind: "ClusterTemplate"}

// Get takes name of the clusterTemplate, and returns the corresponding clusterTemplate object, and an error if there is any.
func (c *FakeClusterTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustertemplatesResource, c.ns, name), &v1alpha1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTemplate), err
}

// List takes label and field selectors, and returns the list of ClusterTemplates that match those selectors.
func (c *FakeClusterTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustertemplatesResource, clustertemplatesKind, c.ns, opts), &v1alpha1.ClusterTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterTemplateList{ListMeta: obj.(*v1alpha1.ClusterTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTemplates.
func (c *FakeClusterTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustertemplatesResource, c.ns, opts))

}

// Create takes the representation of a clusterTemplate and creates it.  Returns the server's representation of the clusterTemplate, and an error, if there is any.
func (c *FakeClusterTemplates) Create(ctx context.Context, clusterTemplate *v1alpha1.ClusterTemplate, opts v1.CreateOptions) (result *v1alpha1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustertemplatesResource, c.ns, clusterTemplate), &v1alpha1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTemplate), err
}

// Update takes the representation of a clusterTemplate and updates it. Returns the server's representation of the clusterTemplate, and an error, if there is any.
func (c *FakeClusterTemplates) Update(ctx context.Context, clusterTemplate *v1alpha1.ClusterTemplate, opts v1.UpdateOptions) (result *v1alpha1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustertemplatesResource, c.ns, clusterTemplate), &v1alpha1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterTemplates) UpdateStatus(ctx context.Context, clusterTemplate *v1alpha1.ClusterTemplate, opts v1.UpdateOptions) (*v1alpha1.ClusterTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clustertemplatesResource, "status", c.ns, clusterTemplate), &v1alpha1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTemplate), err
}

// Delete takes name of the clusterTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClusterTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustertemplatesResource, c.ns, name), &v1alpha1.ClusterTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustertemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clusterTemplate.
func (c *FakeClusterTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustertemplatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTemplate), err
}