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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/vmware-tanzu/cartographer/pkg/apis/carto/v1alpha1"
	scheme "github.com/vmware-tanzu/cartographer/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterSupplyChainsGetter has a method to return a ClusterSupplyChainInterface.
// A group's client should implement this interface.
type ClusterSupplyChainsGetter interface {
	ClusterSupplyChains() ClusterSupplyChainInterface
}

// ClusterSupplyChainInterface has methods to work with ClusterSupplyChain resources.
type ClusterSupplyChainInterface interface {
	Create(ctx context.Context, clusterSupplyChain *v1alpha1.ClusterSupplyChain, opts v1.CreateOptions) (*v1alpha1.ClusterSupplyChain, error)
	Update(ctx context.Context, clusterSupplyChain *v1alpha1.ClusterSupplyChain, opts v1.UpdateOptions) (*v1alpha1.ClusterSupplyChain, error)
	UpdateStatus(ctx context.Context, clusterSupplyChain *v1alpha1.ClusterSupplyChain, opts v1.UpdateOptions) (*v1alpha1.ClusterSupplyChain, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterSupplyChain, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterSupplyChainList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSupplyChain, err error)
	ClusterSupplyChainExpansion
}

// clusterSupplyChains implements ClusterSupplyChainInterface
type clusterSupplyChains struct {
	client rest.Interface
}

// newClusterSupplyChains returns a ClusterSupplyChains
func newClusterSupplyChains(c *CartoV1alpha1Client) *clusterSupplyChains {
	return &clusterSupplyChains{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterSupplyChain, and returns the corresponding clusterSupplyChain object, and an error if there is any.
func (c *clusterSupplyChains) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterSupplyChain, err error) {
	result = &v1alpha1.ClusterSupplyChain{}
	err = c.client.Get().
		Resource("clustersupplychains").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterSupplyChains that match those selectors.
func (c *clusterSupplyChains) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterSupplyChainList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterSupplyChainList{}
	err = c.client.Get().
		Resource("clustersupplychains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterSupplyChains.
func (c *clusterSupplyChains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustersupplychains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterSupplyChain and creates it.  Returns the server's representation of the clusterSupplyChain, and an error, if there is any.
func (c *clusterSupplyChains) Create(ctx context.Context, clusterSupplyChain *v1alpha1.ClusterSupplyChain, opts v1.CreateOptions) (result *v1alpha1.ClusterSupplyChain, err error) {
	result = &v1alpha1.ClusterSupplyChain{}
	err = c.client.Post().
		Resource("clustersupplychains").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSupplyChain).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterSupplyChain and updates it. Returns the server's representation of the clusterSupplyChain, and an error, if there is any.
func (c *clusterSupplyChains) Update(ctx context.Context, clusterSupplyChain *v1alpha1.ClusterSupplyChain, opts v1.UpdateOptions) (result *v1alpha1.ClusterSupplyChain, err error) {
	result = &v1alpha1.ClusterSupplyChain{}
	err = c.client.Put().
		Resource("clustersupplychains").
		Name(clusterSupplyChain.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSupplyChain).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterSupplyChains) UpdateStatus(ctx context.Context, clusterSupplyChain *v1alpha1.ClusterSupplyChain, opts v1.UpdateOptions) (result *v1alpha1.ClusterSupplyChain, err error) {
	result = &v1alpha1.ClusterSupplyChain{}
	err = c.client.Put().
		Resource("clustersupplychains").
		Name(clusterSupplyChain.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterSupplyChain).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterSupplyChain and deletes it. Returns an error if one occurs.
func (c *clusterSupplyChains) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustersupplychains").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterSupplyChains) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustersupplychains").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterSupplyChain.
func (c *clusterSupplyChains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterSupplyChain, err error) {
	result = &v1alpha1.ClusterSupplyChain{}
	err = c.client.Patch(pt).
		Resource("clustersupplychains").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
