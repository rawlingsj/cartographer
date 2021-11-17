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

// FakeClusterRunTemplates implements ClusterRunTemplateInterface
type FakeClusterRunTemplates struct {
	Fake *FakeCartoV1alpha1
}

var clusterruntemplatesResource = schema.GroupVersionResource{Group: "carto.run", Version: "v1alpha1", Resource: "clusterruntemplates"}

var clusterruntemplatesKind = schema.GroupVersionKind{Group: "carto.run", Version: "v1alpha1", Kind: "ClusterRunTemplate"}

// Get takes name of the clusterRunTemplate, and returns the corresponding clusterRunTemplate object, and an error if there is any.
func (c *FakeClusterRunTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterRunTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterruntemplatesResource, name), &v1alpha1.ClusterRunTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRunTemplate), err
}

// List takes label and field selectors, and returns the list of ClusterRunTemplates that match those selectors.
func (c *FakeClusterRunTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterRunTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterruntemplatesResource, clusterruntemplatesKind, opts), &v1alpha1.ClusterRunTemplateList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterRunTemplateList{ListMeta: obj.(*v1alpha1.ClusterRunTemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterRunTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRunTemplates.
func (c *FakeClusterRunTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterruntemplatesResource, opts))
}

// Create takes the representation of a clusterRunTemplate and creates it.  Returns the server's representation of the clusterRunTemplate, and an error, if there is any.
func (c *FakeClusterRunTemplates) Create(ctx context.Context, clusterRunTemplate *v1alpha1.ClusterRunTemplate, opts v1.CreateOptions) (result *v1alpha1.ClusterRunTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterruntemplatesResource, clusterRunTemplate), &v1alpha1.ClusterRunTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRunTemplate), err
}

// Update takes the representation of a clusterRunTemplate and updates it. Returns the server's representation of the clusterRunTemplate, and an error, if there is any.
func (c *FakeClusterRunTemplates) Update(ctx context.Context, clusterRunTemplate *v1alpha1.ClusterRunTemplate, opts v1.UpdateOptions) (result *v1alpha1.ClusterRunTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterruntemplatesResource, clusterRunTemplate), &v1alpha1.ClusterRunTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRunTemplate), err
}

// Delete takes name of the clusterRunTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClusterRunTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterruntemplatesResource, name), &v1alpha1.ClusterRunTemplate{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRunTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterruntemplatesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterRunTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clusterRunTemplate.
func (c *FakeClusterRunTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterRunTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterruntemplatesResource, name, pt, data, subresources...), &v1alpha1.ClusterRunTemplate{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRunTemplate), err
}
