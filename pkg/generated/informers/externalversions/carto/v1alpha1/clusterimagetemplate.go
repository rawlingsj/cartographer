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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	cartov1alpha1 "github.com/vmware-tanzu/cartographer/pkg/apis/carto/v1alpha1"
	versioned "github.com/vmware-tanzu/cartographer/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/vmware-tanzu/cartographer/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vmware-tanzu/cartographer/pkg/generated/listers/carto/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterImageTemplateInformer provides access to a shared informer and lister for
// ClusterImageTemplates.
type ClusterImageTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterImageTemplateLister
}

type clusterImageTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterImageTemplateInformer constructs a new informer for ClusterImageTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterImageTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterImageTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterImageTemplateInformer constructs a new informer for ClusterImageTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterImageTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CartoV1alpha1().ClusterImageTemplates().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CartoV1alpha1().ClusterImageTemplates().Watch(context.TODO(), options)
			},
		},
		&cartov1alpha1.ClusterImageTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterImageTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterImageTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterImageTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cartov1alpha1.ClusterImageTemplate{}, f.defaultInformer)
}

func (f *clusterImageTemplateInformer) Lister() v1alpha1.ClusterImageTemplateLister {
	return v1alpha1.NewClusterImageTemplateLister(f.Informer().GetIndexer())
}
