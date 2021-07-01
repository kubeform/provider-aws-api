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
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Aliases returns a AliasInformer.
	Aliases() AliasInformer
	// Ciphertexts returns a CiphertextInformer.
	Ciphertexts() CiphertextInformer
	// ExternalKeys returns a ExternalKeyInformer.
	ExternalKeys() ExternalKeyInformer
	// Grants returns a GrantInformer.
	Grants() GrantInformer
	// Keys returns a KeyInformer.
	Keys() KeyInformer
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

// Aliases returns a AliasInformer.
func (v *version) Aliases() AliasInformer {
	return &aliasInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Ciphertexts returns a CiphertextInformer.
func (v *version) Ciphertexts() CiphertextInformer {
	return &ciphertextInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ExternalKeys returns a ExternalKeyInformer.
func (v *version) ExternalKeys() ExternalKeyInformer {
	return &externalKeyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Grants returns a GrantInformer.
func (v *version) Grants() GrantInformer {
	return &grantInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Keys returns a KeyInformer.
func (v *version) Keys() KeyInformer {
	return &keyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
