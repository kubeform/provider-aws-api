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

	chimev1alpha1 "kubeform.dev/provider-aws-api/apis/chime/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/chime/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// VoiceConnectorOriginationInformer provides access to a shared informer and lister for
// VoiceConnectorOriginations.
type VoiceConnectorOriginationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VoiceConnectorOriginationLister
}

type voiceConnectorOriginationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVoiceConnectorOriginationInformer constructs a new informer for VoiceConnectorOrigination type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVoiceConnectorOriginationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVoiceConnectorOriginationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVoiceConnectorOriginationInformer constructs a new informer for VoiceConnectorOrigination type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVoiceConnectorOriginationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ChimeV1alpha1().VoiceConnectorOriginations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ChimeV1alpha1().VoiceConnectorOriginations(namespace).Watch(context.TODO(), options)
			},
		},
		&chimev1alpha1.VoiceConnectorOrigination{},
		resyncPeriod,
		indexers,
	)
}

func (f *voiceConnectorOriginationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVoiceConnectorOriginationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *voiceConnectorOriginationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&chimev1alpha1.VoiceConnectorOrigination{}, f.defaultInformer)
}

func (f *voiceConnectorOriginationInformer) Lister() v1alpha1.VoiceConnectorOriginationLister {
	return v1alpha1.NewVoiceConnectorOriginationLister(f.Informer().GetIndexer())
}
