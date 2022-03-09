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
	// AggregateAuthorizations returns a AggregateAuthorizationInformer.
	AggregateAuthorizations() AggregateAuthorizationInformer
	// ConfigRules returns a ConfigRuleInformer.
	ConfigRules() ConfigRuleInformer
	// ConfigurationAggregators returns a ConfigurationAggregatorInformer.
	ConfigurationAggregators() ConfigurationAggregatorInformer
	// ConfigurationRecorders returns a ConfigurationRecorderInformer.
	ConfigurationRecorders() ConfigurationRecorderInformer
	// ConfigurationRecorderStatuses returns a ConfigurationRecorderStatusInformer.
	ConfigurationRecorderStatuses() ConfigurationRecorderStatusInformer
	// ConformancePacks returns a ConformancePackInformer.
	ConformancePacks() ConformancePackInformer
	// DeliveryChannels returns a DeliveryChannelInformer.
	DeliveryChannels() DeliveryChannelInformer
	// OrganizationConformancePacks returns a OrganizationConformancePackInformer.
	OrganizationConformancePacks() OrganizationConformancePackInformer
	// OrganizationCustomRules returns a OrganizationCustomRuleInformer.
	OrganizationCustomRules() OrganizationCustomRuleInformer
	// OrganizationManagedRules returns a OrganizationManagedRuleInformer.
	OrganizationManagedRules() OrganizationManagedRuleInformer
	// RemediationConfigurations returns a RemediationConfigurationInformer.
	RemediationConfigurations() RemediationConfigurationInformer
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

// AggregateAuthorizations returns a AggregateAuthorizationInformer.
func (v *version) AggregateAuthorizations() AggregateAuthorizationInformer {
	return &aggregateAuthorizationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigRules returns a ConfigRuleInformer.
func (v *version) ConfigRules() ConfigRuleInformer {
	return &configRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigurationAggregators returns a ConfigurationAggregatorInformer.
func (v *version) ConfigurationAggregators() ConfigurationAggregatorInformer {
	return &configurationAggregatorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigurationRecorders returns a ConfigurationRecorderInformer.
func (v *version) ConfigurationRecorders() ConfigurationRecorderInformer {
	return &configurationRecorderInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigurationRecorderStatuses returns a ConfigurationRecorderStatusInformer.
func (v *version) ConfigurationRecorderStatuses() ConfigurationRecorderStatusInformer {
	return &configurationRecorderStatusInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConformancePacks returns a ConformancePackInformer.
func (v *version) ConformancePacks() ConformancePackInformer {
	return &conformancePackInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DeliveryChannels returns a DeliveryChannelInformer.
func (v *version) DeliveryChannels() DeliveryChannelInformer {
	return &deliveryChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OrganizationConformancePacks returns a OrganizationConformancePackInformer.
func (v *version) OrganizationConformancePacks() OrganizationConformancePackInformer {
	return &organizationConformancePackInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OrganizationCustomRules returns a OrganizationCustomRuleInformer.
func (v *version) OrganizationCustomRules() OrganizationCustomRuleInformer {
	return &organizationCustomRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OrganizationManagedRules returns a OrganizationManagedRuleInformer.
func (v *version) OrganizationManagedRules() OrganizationManagedRuleInformer {
	return &organizationManagedRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RemediationConfigurations returns a RemediationConfigurationInformer.
func (v *version) RemediationConfigurations() RemediationConfigurationInformer {
	return &remediationConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
