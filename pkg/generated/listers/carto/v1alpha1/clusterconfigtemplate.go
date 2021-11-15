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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vmware-tanzu/cartographer/pkg/apis/carto/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterConfigTemplateLister helps list ClusterConfigTemplates.
// All objects returned here must be treated as read-only.
type ClusterConfigTemplateLister interface {
	// List lists all ClusterConfigTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterConfigTemplate, err error)
	// ClusterConfigTemplates returns an object that can list and get ClusterConfigTemplates.
	ClusterConfigTemplates(namespace string) ClusterConfigTemplateNamespaceLister
	ClusterConfigTemplateListerExpansion
}

// clusterConfigTemplateLister implements the ClusterConfigTemplateLister interface.
type clusterConfigTemplateLister struct {
	indexer cache.Indexer
}

// NewClusterConfigTemplateLister returns a new ClusterConfigTemplateLister.
func NewClusterConfigTemplateLister(indexer cache.Indexer) ClusterConfigTemplateLister {
	return &clusterConfigTemplateLister{indexer: indexer}
}

// List lists all ClusterConfigTemplates in the indexer.
func (s *clusterConfigTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterConfigTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterConfigTemplate))
	})
	return ret, err
}

// ClusterConfigTemplates returns an object that can list and get ClusterConfigTemplates.
func (s *clusterConfigTemplateLister) ClusterConfigTemplates(namespace string) ClusterConfigTemplateNamespaceLister {
	return clusterConfigTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterConfigTemplateNamespaceLister helps list and get ClusterConfigTemplates.
// All objects returned here must be treated as read-only.
type ClusterConfigTemplateNamespaceLister interface {
	// List lists all ClusterConfigTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterConfigTemplate, err error)
	// Get retrieves the ClusterConfigTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterConfigTemplate, error)
	ClusterConfigTemplateNamespaceListerExpansion
}

// clusterConfigTemplateNamespaceLister implements the ClusterConfigTemplateNamespaceLister
// interface.
type clusterConfigTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterConfigTemplates in the indexer for a given namespace.
func (s clusterConfigTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterConfigTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterConfigTemplate))
	})
	return ret, err
}

// Get retrieves the ClusterConfigTemplate from the indexer for a given namespace and name.
func (s clusterConfigTemplateNamespaceLister) Get(name string) (*v1alpha1.ClusterConfigTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterconfigtemplate"), name)
	}
	return obj.(*v1alpha1.ClusterConfigTemplate), nil
}