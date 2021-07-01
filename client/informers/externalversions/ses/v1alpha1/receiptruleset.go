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
	sesv1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/ses/v1alpha1"
)

// ReceiptRuleSetInformer provides access to a shared informer and lister for
// ReceiptRuleSets.
type ReceiptRuleSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReceiptRuleSetLister
}

type receiptRuleSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReceiptRuleSetInformer constructs a new informer for ReceiptRuleSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReceiptRuleSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReceiptRuleSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReceiptRuleSetInformer constructs a new informer for ReceiptRuleSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReceiptRuleSetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SesV1alpha1().ReceiptRuleSets(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SesV1alpha1().ReceiptRuleSets(namespace).Watch(context.TODO(), options)
			},
		},
		&sesv1alpha1.ReceiptRuleSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *receiptRuleSetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReceiptRuleSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *receiptRuleSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&sesv1alpha1.ReceiptRuleSet{}, f.defaultInformer)
}

func (f *receiptRuleSetInformer) Lister() v1alpha1.ReceiptRuleSetLister {
	return v1alpha1.NewReceiptRuleSetLister(f.Informer().GetIndexer())
}
