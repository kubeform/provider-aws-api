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

	worklinkv1alpha1 "kubeform.dev/provider-aws-api/apis/worklink/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/worklink/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WebsiteCertificateAuthorityAssociationInformer provides access to a shared informer and lister for
// WebsiteCertificateAuthorityAssociations.
type WebsiteCertificateAuthorityAssociationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.WebsiteCertificateAuthorityAssociationLister
}

type websiteCertificateAuthorityAssociationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWebsiteCertificateAuthorityAssociationInformer constructs a new informer for WebsiteCertificateAuthorityAssociation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWebsiteCertificateAuthorityAssociationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWebsiteCertificateAuthorityAssociationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWebsiteCertificateAuthorityAssociationInformer constructs a new informer for WebsiteCertificateAuthorityAssociation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWebsiteCertificateAuthorityAssociationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorklinkV1alpha1().WebsiteCertificateAuthorityAssociations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WorklinkV1alpha1().WebsiteCertificateAuthorityAssociations(namespace).Watch(context.TODO(), options)
			},
		},
		&worklinkv1alpha1.WebsiteCertificateAuthorityAssociation{},
		resyncPeriod,
		indexers,
	)
}

func (f *websiteCertificateAuthorityAssociationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWebsiteCertificateAuthorityAssociationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *websiteCertificateAuthorityAssociationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&worklinkv1alpha1.WebsiteCertificateAuthorityAssociation{}, f.defaultInformer)
}

func (f *websiteCertificateAuthorityAssociationInformer) Lister() v1alpha1.WebsiteCertificateAuthorityAssociationLister {
	return v1alpha1.NewWebsiteCertificateAuthorityAssociationLister(f.Informer().GetIndexer())
}
