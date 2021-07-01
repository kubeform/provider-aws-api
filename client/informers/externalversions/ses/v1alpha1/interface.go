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
	// ActiveReceiptRuleSets returns a ActiveReceiptRuleSetInformer.
	ActiveReceiptRuleSets() ActiveReceiptRuleSetInformer
	// ConfigurationSets returns a ConfigurationSetInformer.
	ConfigurationSets() ConfigurationSetInformer
	// DomainDkims returns a DomainDkimInformer.
	DomainDkims() DomainDkimInformer
	// DomainIdentities returns a DomainIdentityInformer.
	DomainIdentities() DomainIdentityInformer
	// DomainIdentityVerifications returns a DomainIdentityVerificationInformer.
	DomainIdentityVerifications() DomainIdentityVerificationInformer
	// DomainMailFroms returns a DomainMailFromInformer.
	DomainMailFroms() DomainMailFromInformer
	// EmailIdentities returns a EmailIdentityInformer.
	EmailIdentities() EmailIdentityInformer
	// EventDestinations returns a EventDestinationInformer.
	EventDestinations() EventDestinationInformer
	// IdentityNotificationTopics returns a IdentityNotificationTopicInformer.
	IdentityNotificationTopics() IdentityNotificationTopicInformer
	// IdentityPolicies returns a IdentityPolicyInformer.
	IdentityPolicies() IdentityPolicyInformer
	// ReceiptFilters returns a ReceiptFilterInformer.
	ReceiptFilters() ReceiptFilterInformer
	// ReceiptRules returns a ReceiptRuleInformer.
	ReceiptRules() ReceiptRuleInformer
	// ReceiptRuleSets returns a ReceiptRuleSetInformer.
	ReceiptRuleSets() ReceiptRuleSetInformer
	// Templates returns a TemplateInformer.
	Templates() TemplateInformer
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

// ActiveReceiptRuleSets returns a ActiveReceiptRuleSetInformer.
func (v *version) ActiveReceiptRuleSets() ActiveReceiptRuleSetInformer {
	return &activeReceiptRuleSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigurationSets returns a ConfigurationSetInformer.
func (v *version) ConfigurationSets() ConfigurationSetInformer {
	return &configurationSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainDkims returns a DomainDkimInformer.
func (v *version) DomainDkims() DomainDkimInformer {
	return &domainDkimInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainIdentities returns a DomainIdentityInformer.
func (v *version) DomainIdentities() DomainIdentityInformer {
	return &domainIdentityInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainIdentityVerifications returns a DomainIdentityVerificationInformer.
func (v *version) DomainIdentityVerifications() DomainIdentityVerificationInformer {
	return &domainIdentityVerificationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainMailFroms returns a DomainMailFromInformer.
func (v *version) DomainMailFroms() DomainMailFromInformer {
	return &domainMailFromInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EmailIdentities returns a EmailIdentityInformer.
func (v *version) EmailIdentities() EmailIdentityInformer {
	return &emailIdentityInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EventDestinations returns a EventDestinationInformer.
func (v *version) EventDestinations() EventDestinationInformer {
	return &eventDestinationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityNotificationTopics returns a IdentityNotificationTopicInformer.
func (v *version) IdentityNotificationTopics() IdentityNotificationTopicInformer {
	return &identityNotificationTopicInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityPolicies returns a IdentityPolicyInformer.
func (v *version) IdentityPolicies() IdentityPolicyInformer {
	return &identityPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ReceiptFilters returns a ReceiptFilterInformer.
func (v *version) ReceiptFilters() ReceiptFilterInformer {
	return &receiptFilterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ReceiptRules returns a ReceiptRuleInformer.
func (v *version) ReceiptRules() ReceiptRuleInformer {
	return &receiptRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ReceiptRuleSets returns a ReceiptRuleSetInformer.
func (v *version) ReceiptRuleSets() ReceiptRuleSetInformer {
	return &receiptRuleSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Templates returns a TemplateInformer.
func (v *version) Templates() TemplateInformer {
	return &templateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
