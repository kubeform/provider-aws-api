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

// BucketServerSideEncryptionConfigurationInformer provides access to a shared informer and lister for
// BucketServerSideEncryptionConfigurations.
type BucketServerSideEncryptionConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BucketServerSideEncryptionConfigurationLister
}

type bucketServerSideEncryptionConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBucketServerSideEncryptionConfigurationInformer constructs a new informer for BucketServerSideEncryptionConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBucketServerSideEncryptionConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBucketServerSideEncryptionConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBucketServerSideEncryptionConfigurationInformer constructs a new informer for BucketServerSideEncryptionConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBucketServerSideEncryptionConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.S3V1alpha1().BucketServerSideEncryptionConfigurations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.S3V1alpha1().BucketServerSideEncryptionConfigurations(namespace).Watch(context.TODO(), options)
			},
		},
		&s3v1alpha1.BucketServerSideEncryptionConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *bucketServerSideEncryptionConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBucketServerSideEncryptionConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *bucketServerSideEncryptionConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&s3v1alpha1.BucketServerSideEncryptionConfiguration{}, f.defaultInformer)
}

func (f *bucketServerSideEncryptionConfigurationInformer) Lister() v1alpha1.BucketServerSideEncryptionConfigurationLister {
	return v1alpha1.NewBucketServerSideEncryptionConfigurationLister(f.Informer().GetIndexer())
}
