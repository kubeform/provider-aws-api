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

	s3v1alpha1 "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/s3/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BucketCorsConfigurationInformer provides access to a shared informer and lister for
// BucketCorsConfigurations.
type BucketCorsConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BucketCorsConfigurationLister
}

type bucketCorsConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBucketCorsConfigurationInformer constructs a new informer for BucketCorsConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBucketCorsConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBucketCorsConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBucketCorsConfigurationInformer constructs a new informer for BucketCorsConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBucketCorsConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.S3V1alpha1().BucketCorsConfigurations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.S3V1alpha1().BucketCorsConfigurations(namespace).Watch(context.TODO(), options)
			},
		},
		&s3v1alpha1.BucketCorsConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *bucketCorsConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBucketCorsConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bucketCorsConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&s3v1alpha1.BucketCorsConfiguration{}, f.defaultInformer)
}

func (f *bucketCorsConfigurationInformer) Lister() v1alpha1.BucketCorsConfigurationLister {
	return v1alpha1.NewBucketCorsConfigurationLister(f.Informer().GetIndexer())
}
