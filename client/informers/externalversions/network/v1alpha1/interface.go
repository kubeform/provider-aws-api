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

	networkv1alpha1 "kubeform.dev/provider-aws-api/apis/network/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/network/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Acls returns a AclInformer.
	Acls() AclInformer
	// AclRules returns a AclRuleInformer.
	AclRules() AclRuleInformer
	// Interfaces returns a InterfaceInformer.
	Interfaces() InterfaceInformer
	// InterfaceAttachments returns a InterfaceAttachmentInformer.
	InterfaceAttachments() InterfaceAttachmentInformer
	// InterfaceSgAttachments returns a InterfaceSgAttachmentInformer.
	InterfaceSgAttachments() InterfaceSgAttachmentInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Acls returns a AclInformer.
func (v *version) Acls() AclInformer {
	return &aclInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AclRules returns a AclRuleInformer.
func (v *version) AclRules() AclRuleInformer {
	return &aclRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Interfaces returns a InterfaceInformer.
func (v *version) Interfaces() InterfaceInformer {
	return &interfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceAttachments returns a InterfaceAttachmentInformer.
func (v *version) InterfaceAttachments() InterfaceAttachmentInformer {
	return &interfaceAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceSgAttachments returns a InterfaceSgAttachmentInformer.
func (v *version) InterfaceSgAttachments() InterfaceSgAttachmentInformer {
	return &interfaceSgAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InterfaceInformer provides access to a shared informer and lister for
// Interfaces.
type InterfaceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.InterfaceLister
}

type interfaceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewInterfaceInformer constructs a new informer for Interface type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInterfaceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInterfaceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredInterfaceInformer constructs a new informer for Interface type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInterfaceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().Interfaces(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkV1alpha1().Interfaces(namespace).Watch(context.TODO(), options)
			},
		},
		&networkv1alpha1.Interface{},
		resyncPeriod,
		indexers,
	)
}

func (f *interfaceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInterfaceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *interfaceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkv1alpha1.Interface{}, f.defaultInformer)
}

func (f *interfaceInformer) Lister() v1alpha1.InterfaceLister {
	return v1alpha1.NewInterfaceLister(f.Informer().GetIndexer())
}
