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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	macie2v1alpha1 "kubeform.dev/provider-aws-api/apis/macie2/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/macie2/v1alpha1"
)

// FindingsFilterInformer provides access to a shared informer and lister for
// FindingsFilters.
type FindingsFilterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FindingsFilterLister
}

type findingsFilterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFindingsFilterInformer constructs a new informer for FindingsFilter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFindingsFilterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFindingsFilterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFindingsFilterInformer constructs a new informer for FindingsFilter type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFindingsFilterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Macie2V1alpha1().FindingsFilters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Macie2V1alpha1().FindingsFilters(namespace).Watch(context.TODO(), options)
			},
		},
		&macie2v1alpha1.FindingsFilter{},
		resyncPeriod,
		indexers,
	)
}

func (f *findingsFilterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFindingsFilterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *findingsFilterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&macie2v1alpha1.FindingsFilter{}, f.defaultInformer)
}

func (f *findingsFilterInformer) Lister() v1alpha1.FindingsFilterLister {
	return v1alpha1.NewFindingsFilterLister(f.Informer().GetIndexer())
}
