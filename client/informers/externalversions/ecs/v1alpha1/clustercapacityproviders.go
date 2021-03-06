/*
Copyright AppsCode Inc. and Contributors

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

	ecsv1alpha1 "kubeform.dev/provider-aws-api/apis/ecs/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/ecs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterCapacityProvidersInformer provides access to a shared informer and lister for
// ClusterCapacityProviderses.
type ClusterCapacityProvidersInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterCapacityProvidersLister
}

type clusterCapacityProvidersInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterCapacityProvidersInformer constructs a new informer for ClusterCapacityProviders type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterCapacityProvidersInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterCapacityProvidersInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterCapacityProvidersInformer constructs a new informer for ClusterCapacityProviders type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterCapacityProvidersInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EcsV1alpha1().ClusterCapacityProviderses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EcsV1alpha1().ClusterCapacityProviderses(namespace).Watch(context.TODO(), options)
			},
		},
		&ecsv1alpha1.ClusterCapacityProviders{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterCapacityProvidersInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterCapacityProvidersInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterCapacityProvidersInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ecsv1alpha1.ClusterCapacityProviders{}, f.defaultInformer)
}

func (f *clusterCapacityProvidersInformer) Lister() v1alpha1.ClusterCapacityProvidersLister {
	return v1alpha1.NewClusterCapacityProvidersLister(f.Informer().GetIndexer())
}
