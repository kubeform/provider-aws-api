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

	s3controlv1alpha1 "kubeform.dev/provider-aws-api/apis/s3control/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/s3control/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ObjectLambdaAccessPointPolicyInformer provides access to a shared informer and lister for
// ObjectLambdaAccessPointPolicies.
type ObjectLambdaAccessPointPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ObjectLambdaAccessPointPolicyLister
}

type objectLambdaAccessPointPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewObjectLambdaAccessPointPolicyInformer constructs a new informer for ObjectLambdaAccessPointPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewObjectLambdaAccessPointPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredObjectLambdaAccessPointPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredObjectLambdaAccessPointPolicyInformer constructs a new informer for ObjectLambdaAccessPointPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredObjectLambdaAccessPointPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.S3controlV1alpha1().ObjectLambdaAccessPointPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.S3controlV1alpha1().ObjectLambdaAccessPointPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&s3controlv1alpha1.ObjectLambdaAccessPointPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *objectLambdaAccessPointPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredObjectLambdaAccessPointPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *objectLambdaAccessPointPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&s3controlv1alpha1.ObjectLambdaAccessPointPolicy{}, f.defaultInformer)
}

func (f *objectLambdaAccessPointPolicyInformer) Lister() v1alpha1.ObjectLambdaAccessPointPolicyLister {
	return v1alpha1.NewObjectLambdaAccessPointPolicyLister(f.Informer().GetIndexer())
}
