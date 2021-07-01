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
	pinpointv1alpha1 "kubeform.dev/provider-aws-api/apis/pinpoint/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/pinpoint/v1alpha1"
)

// ApnsVoipChannelInformer provides access to a shared informer and lister for
// ApnsVoipChannels.
type ApnsVoipChannelInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ApnsVoipChannelLister
}

type apnsVoipChannelInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewApnsVoipChannelInformer constructs a new informer for ApnsVoipChannel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApnsVoipChannelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApnsVoipChannelInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredApnsVoipChannelInformer constructs a new informer for ApnsVoipChannel type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApnsVoipChannelInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PinpointV1alpha1().ApnsVoipChannels(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PinpointV1alpha1().ApnsVoipChannels(namespace).Watch(context.TODO(), options)
			},
		},
		&pinpointv1alpha1.ApnsVoipChannel{},
		resyncPeriod,
		indexers,
	)
}

func (f *apnsVoipChannelInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApnsVoipChannelInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *apnsVoipChannelInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pinpointv1alpha1.ApnsVoipChannel{}, f.defaultInformer)
}

func (f *apnsVoipChannelInformer) Lister() v1alpha1.ApnsVoipChannelLister {
	return v1alpha1.NewApnsVoipChannelLister(f.Informer().GetIndexer())
}
