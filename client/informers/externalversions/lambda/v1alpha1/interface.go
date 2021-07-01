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
	// CodeSigningConfigs returns a CodeSigningConfigInformer.
	CodeSigningConfigs() CodeSigningConfigInformer
	// EventSourceMappings returns a EventSourceMappingInformer.
	EventSourceMappings() EventSourceMappingInformer
	// Functions returns a FunctionInformer.
	Functions() FunctionInformer
	// FunctionEventInvokeConfigs returns a FunctionEventInvokeConfigInformer.
	FunctionEventInvokeConfigs() FunctionEventInvokeConfigInformer
	// LayerVersions returns a LayerVersionInformer.
	LayerVersions() LayerVersionInformer
	// Permissions returns a PermissionInformer.
	Permissions() PermissionInformer
	// ProvisionedConcurrencyConfigs returns a ProvisionedConcurrencyConfigInformer.
	ProvisionedConcurrencyConfigs() ProvisionedConcurrencyConfigInformer
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

// CodeSigningConfigs returns a CodeSigningConfigInformer.
func (v *version) CodeSigningConfigs() CodeSigningConfigInformer {
	return &codeSigningConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EventSourceMappings returns a EventSourceMappingInformer.
func (v *version) EventSourceMappings() EventSourceMappingInformer {
	return &eventSourceMappingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Functions returns a FunctionInformer.
func (v *version) Functions() FunctionInformer {
	return &functionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FunctionEventInvokeConfigs returns a FunctionEventInvokeConfigInformer.
func (v *version) FunctionEventInvokeConfigs() FunctionEventInvokeConfigInformer {
	return &functionEventInvokeConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LayerVersions returns a LayerVersionInformer.
func (v *version) LayerVersions() LayerVersionInformer {
	return &layerVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Permissions returns a PermissionInformer.
func (v *version) Permissions() PermissionInformer {
	return &permissionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProvisionedConcurrencyConfigs returns a ProvisionedConcurrencyConfigInformer.
func (v *version) ProvisionedConcurrencyConfigs() ProvisionedConcurrencyConfigInformer {
	return &provisionedConcurrencyConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
