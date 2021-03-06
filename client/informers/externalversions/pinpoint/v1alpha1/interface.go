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
	// AdmChannels returns a AdmChannelInformer.
	AdmChannels() AdmChannelInformer
	// ApnsChannels returns a ApnsChannelInformer.
	ApnsChannels() ApnsChannelInformer
	// ApnsSandboxChannels returns a ApnsSandboxChannelInformer.
	ApnsSandboxChannels() ApnsSandboxChannelInformer
	// ApnsVoipChannels returns a ApnsVoipChannelInformer.
	ApnsVoipChannels() ApnsVoipChannelInformer
	// ApnsVoipSandboxChannels returns a ApnsVoipSandboxChannelInformer.
	ApnsVoipSandboxChannels() ApnsVoipSandboxChannelInformer
	// Apps returns a AppInformer.
	Apps() AppInformer
	// BaiduChannels returns a BaiduChannelInformer.
	BaiduChannels() BaiduChannelInformer
	// EmailChannels returns a EmailChannelInformer.
	EmailChannels() EmailChannelInformer
	// EventStreams returns a EventStreamInformer.
	EventStreams() EventStreamInformer
	// GcmChannels returns a GcmChannelInformer.
	GcmChannels() GcmChannelInformer
	// SmsChannels returns a SmsChannelInformer.
	SmsChannels() SmsChannelInformer
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

// AdmChannels returns a AdmChannelInformer.
func (v *version) AdmChannels() AdmChannelInformer {
	return &admChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApnsChannels returns a ApnsChannelInformer.
func (v *version) ApnsChannels() ApnsChannelInformer {
	return &apnsChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApnsSandboxChannels returns a ApnsSandboxChannelInformer.
func (v *version) ApnsSandboxChannels() ApnsSandboxChannelInformer {
	return &apnsSandboxChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApnsVoipChannels returns a ApnsVoipChannelInformer.
func (v *version) ApnsVoipChannels() ApnsVoipChannelInformer {
	return &apnsVoipChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApnsVoipSandboxChannels returns a ApnsVoipSandboxChannelInformer.
func (v *version) ApnsVoipSandboxChannels() ApnsVoipSandboxChannelInformer {
	return &apnsVoipSandboxChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Apps returns a AppInformer.
func (v *version) Apps() AppInformer {
	return &appInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BaiduChannels returns a BaiduChannelInformer.
func (v *version) BaiduChannels() BaiduChannelInformer {
	return &baiduChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EmailChannels returns a EmailChannelInformer.
func (v *version) EmailChannels() EmailChannelInformer {
	return &emailChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EventStreams returns a EventStreamInformer.
func (v *version) EventStreams() EventStreamInformer {
	return &eventStreamInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GcmChannels returns a GcmChannelInformer.
func (v *version) GcmChannels() GcmChannelInformer {
	return &gcmChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SmsChannels returns a SmsChannelInformer.
func (v *version) SmsChannels() SmsChannelInformer {
	return &smsChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
