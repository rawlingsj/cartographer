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

// ClusterDeliveryLister helps list ClusterDeliveries.
// All objects returned here must be treated as read-only.
type ClusterDeliveryLister interface {
	// List lists all ClusterDeliveries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterDelivery, err error)
	// ClusterDeliveries returns an object that can list and get ClusterDeliveries.
	ClusterDeliveries(namespace string) ClusterDeliveryNamespaceLister
	ClusterDeliveryListerExpansion
}

// clusterDeliveryLister implements the ClusterDeliveryLister interface.
type clusterDeliveryLister struct {
	indexer cache.Indexer
}

// NewClusterDeliveryLister returns a new ClusterDeliveryLister.
func NewClusterDeliveryLister(indexer cache.Indexer) ClusterDeliveryLister {
	return &clusterDeliveryLister{indexer: indexer}
}

// List lists all ClusterDeliveries in the indexer.
func (s *clusterDeliveryLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterDelivery, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterDelivery))
	})
	return ret, err
}

// ClusterDeliveries returns an object that can list and get ClusterDeliveries.
func (s *clusterDeliveryLister) ClusterDeliveries(namespace string) ClusterDeliveryNamespaceLister {
	return clusterDeliveryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterDeliveryNamespaceLister helps list and get ClusterDeliveries.
// All objects returned here must be treated as read-only.
type ClusterDeliveryNamespaceLister interface {
	// List lists all ClusterDeliveries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterDelivery, err error)
	// Get retrieves the ClusterDelivery from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterDelivery, error)
	ClusterDeliveryNamespaceListerExpansion
}

// clusterDeliveryNamespaceLister implements the ClusterDeliveryNamespaceLister
// interface.
type clusterDeliveryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterDeliveries in the indexer for a given namespace.
func (s clusterDeliveryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterDelivery, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterDelivery))
	})
	return ret, err
}

// Get retrieves the ClusterDelivery from the indexer for a given namespace and name.
func (s clusterDeliveryNamespaceLister) Get(name string) (*v1alpha1.ClusterDelivery, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clusterdelivery"), name)
	}
	return obj.(*v1alpha1.ClusterDelivery), nil
}
