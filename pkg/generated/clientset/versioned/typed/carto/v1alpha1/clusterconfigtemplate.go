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

// ClusterConfigTemplatesGetter has a method to return a ClusterConfigTemplateInterface.
// A group's client should implement this interface.
type ClusterConfigTemplatesGetter interface {
	ClusterConfigTemplates(namespace string) ClusterConfigTemplateInterface
}

// ClusterConfigTemplateInterface has methods to work with ClusterConfigTemplate resources.
type ClusterConfigTemplateInterface interface {
	Create(ctx context.Context, clusterConfigTemplate *v1alpha1.ClusterConfigTemplate, opts v1.CreateOptions) (*v1alpha1.ClusterConfigTemplate, error)
	Update(ctx context.Context, clusterConfigTemplate *v1alpha1.ClusterConfigTemplate, opts v1.UpdateOptions) (*v1alpha1.ClusterConfigTemplate, error)
	UpdateStatus(ctx context.Context, clusterConfigTemplate *v1alpha1.ClusterConfigTemplate, opts v1.UpdateOptions) (*v1alpha1.ClusterConfigTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterConfigTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterConfigTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterConfigTemplate, err error)
	ClusterConfigTemplateExpansion
}

// clusterConfigTemplates implements ClusterConfigTemplateInterface
type clusterConfigTemplates struct {
	client rest.Interface
	ns     string
}

// newClusterConfigTemplates returns a ClusterConfigTemplates
func newClusterConfigTemplates(c *CartoV1alpha1Client, namespace string) *clusterConfigTemplates {
	return &clusterConfigTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterConfigTemplate, and returns the corresponding clusterConfigTemplate object, and an error if there is any.
func (c *clusterConfigTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterConfigTemplate, err error) {
	result = &v1alpha1.ClusterConfigTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterConfigTemplates that match those selectors.
func (c *clusterConfigTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterConfigTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterConfigTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterConfigTemplates.
func (c *clusterConfigTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterConfigTemplate and creates it.  Returns the server's representation of the clusterConfigTemplate, and an error, if there is any.
func (c *clusterConfigTemplates) Create(ctx context.Context, clusterConfigTemplate *v1alpha1.ClusterConfigTemplate, opts v1.CreateOptions) (result *v1alpha1.ClusterConfigTemplate, err error) {
	result = &v1alpha1.ClusterConfigTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterConfigTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterConfigTemplate and updates it. Returns the server's representation of the clusterConfigTemplate, and an error, if there is any.
func (c *clusterConfigTemplates) Update(ctx context.Context, clusterConfigTemplate *v1alpha1.ClusterConfigTemplate, opts v1.UpdateOptions) (result *v1alpha1.ClusterConfigTemplate, err error) {
	result = &v1alpha1.ClusterConfigTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		Name(clusterConfigTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterConfigTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterConfigTemplates) UpdateStatus(ctx context.Context, clusterConfigTemplate *v1alpha1.ClusterConfigTemplate, opts v1.UpdateOptions) (result *v1alpha1.ClusterConfigTemplate, err error) {
	result = &v1alpha1.ClusterConfigTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		Name(clusterConfigTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterConfigTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterConfigTemplate and deletes it. Returns an error if one occurs.
func (c *clusterConfigTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterConfigTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterConfigTemplate.
func (c *clusterConfigTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterConfigTemplate, err error) {
	result = &v1alpha1.ClusterConfigTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterconfigtemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
