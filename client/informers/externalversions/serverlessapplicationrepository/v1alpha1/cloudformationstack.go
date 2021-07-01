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
	serverlessapplicationrepositoryv1alpha1 "kubeform.dev/provider-aws-api/apis/serverlessapplicationrepository/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/serverlessapplicationrepository/v1alpha1"
)

// CloudformationStackInformer provides access to a shared informer and lister for
// CloudformationStacks.
type CloudformationStackInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CloudformationStackLister
}

type cloudformationStackInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCloudformationStackInformer constructs a new informer for CloudformationStack type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCloudformationStackInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloudformationStackInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCloudformationStackInformer constructs a new informer for CloudformationStack type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCloudformationStackInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServerlessapplicationrepositoryV1alpha1().CloudformationStacks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServerlessapplicationrepositoryV1alpha1().CloudformationStacks(namespace).Watch(context.TODO(), options)
			},
		},
		&serverlessapplicationrepositoryv1alpha1.CloudformationStack{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloudformationStackInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloudformationStackInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloudformationStackInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&serverlessapplicationrepositoryv1alpha1.CloudformationStack{}, f.defaultInformer)
}

func (f *cloudformationStackInformer) Lister() v1alpha1.CloudformationStackLister {
	return v1alpha1.NewCloudformationStackLister(f.Informer().GetIndexer())
}
