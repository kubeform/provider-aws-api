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

	imagebuilderv1alpha1 "kubeform.dev/provider-aws-api/apis/imagebuilder/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/imagebuilder/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ContainerRecipeInformer provides access to a shared informer and lister for
// ContainerRecipes.
type ContainerRecipeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ContainerRecipeLister
}

type containerRecipeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewContainerRecipeInformer constructs a new informer for ContainerRecipe type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewContainerRecipeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredContainerRecipeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredContainerRecipeInformer constructs a new informer for ContainerRecipe type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredContainerRecipeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ImagebuilderV1alpha1().ContainerRecipes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ImagebuilderV1alpha1().ContainerRecipes(namespace).Watch(context.TODO(), options)
			},
		},
		&imagebuilderv1alpha1.ContainerRecipe{},
		resyncPeriod,
		indexers,
	)
}

func (f *containerRecipeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredContainerRecipeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *containerRecipeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&imagebuilderv1alpha1.ContainerRecipe{}, f.defaultInformer)
}

func (f *containerRecipeInformer) Lister() v1alpha1.ContainerRecipeLister {
	return v1alpha1.NewContainerRecipeLister(f.Informer().GetIndexer())
}
