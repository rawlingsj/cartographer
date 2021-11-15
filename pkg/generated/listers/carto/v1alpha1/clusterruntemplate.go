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

// ClusterRunTemplateLister helps list ClusterRunTemplates.
// All objects returned here must be treated as read-only.
type ClusterRunTemplateLister interface {
	// List lists all ClusterRunTemplates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterRunTemplate, err error)
	// ClusterRunTemplates returns an object that can list and get ClusterRunTemplates.
	ClusterRunTemplates(namespace string) ClusterRunTemplateNamespaceLister
	ClusterRunTemplateListerExpansion
}

// clusterRunTemplateLister implements the ClusterRunTemplateLister interface.
type clusterRunTemplateLister struct {
	indexer cache.Indexer
}

// NewClusterRunTemplateLister returns a new ClusterRunTemplateLister.
func NewClusterRunTemplateLister(indexer cache.Indexer) ClusterRunTemplateLister {
	return &clusterRunTemplateLister{indexer: indexer}
}

// List lists all ClusterRunTemplates in the indexer.
func (s *clusterRunTemplateLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterRunTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterRunTemplate))
	})
	return ret, err
}

// ClusterRunTemplates returns an object that can list and get ClusterRunTemplates.
func (s *clusterRunTemplateLister) ClusterRunTemplates(namespace string) ClusterRunTemplateNamespaceLister {
	return clusterRunTemplateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterRunTemplateNamespaceLister helps list and get ClusterRunTemplates.
// All objects returned here must be treated as read-only.
type ClusterRunTemplateNamespaceLister interface {
	// List lists all ClusterRunTemplates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterRunTemplate, err error)
	// Get retrieves the ClusterRunTemplate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterRunTemplate, error)
	ClusterRunTemplateNamespaceListerExpansion
}

// clusterRunTemplateNamespaceLister implements the ClusterRunTemplateNamespaceLister
// interface.
type clusterRunTemplateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterRunTemplates in the indexer for a given namespace.
func (s clusterRunTemplateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterRunTemplate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterRunTemplate))
	})
	return ret, err
}

// Get retrieves the ClusterRunTemplate from the indexer for a given namespace and name.
func (s clusterRunTemplateNamespaceLister) Get(name string) (*v1alpha1.ClusterRunTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterruntemplate"), name)
	}
	return obj.(*v1alpha1.ClusterRunTemplate), nil
}