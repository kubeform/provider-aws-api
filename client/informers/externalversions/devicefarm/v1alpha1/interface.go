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
	// DevicePools returns a DevicePoolInformer.
	DevicePools() DevicePoolInformer
	// InstanceProfiles returns a InstanceProfileInformer.
	InstanceProfiles() InstanceProfileInformer
	// NetworkProfiles returns a NetworkProfileInformer.
	NetworkProfiles() NetworkProfileInformer
	// Projects returns a ProjectInformer.
	Projects() ProjectInformer
	// TestGridProjects returns a TestGridProjectInformer.
	TestGridProjects() TestGridProjectInformer
	// Uploads returns a UploadInformer.
	Uploads() UploadInformer
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

// DevicePools returns a DevicePoolInformer.
func (v *version) DevicePools() DevicePoolInformer {
	return &devicePoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstanceProfiles returns a InstanceProfileInformer.
func (v *version) InstanceProfiles() InstanceProfileInformer {
	return &instanceProfileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkProfiles returns a NetworkProfileInformer.
func (v *version) NetworkProfiles() NetworkProfileInformer {
	return &networkProfileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Projects returns a ProjectInformer.
func (v *version) Projects() ProjectInformer {
	return &projectInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TestGridProjects returns a TestGridProjectInformer.
func (v *version) TestGridProjects() TestGridProjectInformer {
	return &testGridProjectInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Uploads returns a UploadInformer.
func (v *version) Uploads() UploadInformer {
	return &uploadInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
