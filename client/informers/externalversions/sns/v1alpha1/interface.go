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
	// PlatformApplications returns a PlatformApplicationInformer.
	PlatformApplications() PlatformApplicationInformer
	// SmsPreferenceses returns a SmsPreferencesInformer.
	SmsPreferenceses() SmsPreferencesInformer
	// Topics returns a TopicInformer.
	Topics() TopicInformer
	// TopicPolicies returns a TopicPolicyInformer.
	TopicPolicies() TopicPolicyInformer
	// TopicSubscriptions returns a TopicSubscriptionInformer.
	TopicSubscriptions() TopicSubscriptionInformer
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

// PlatformApplications returns a PlatformApplicationInformer.
func (v *version) PlatformApplications() PlatformApplicationInformer {
	return &platformApplicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SmsPreferenceses returns a SmsPreferencesInformer.
func (v *version) SmsPreferenceses() SmsPreferencesInformer {
	return &smsPreferencesInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Topics returns a TopicInformer.
func (v *version) Topics() TopicInformer {
	return &topicInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TopicPolicies returns a TopicPolicyInformer.
func (v *version) TopicPolicies() TopicPolicyInformer {
	return &topicPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TopicSubscriptions returns a TopicSubscriptionInformer.
func (v *version) TopicSubscriptions() TopicSubscriptionInformer {
	return &topicSubscriptionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
