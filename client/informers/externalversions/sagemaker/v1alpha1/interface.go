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
	// Apps returns a AppInformer.
	Apps() AppInformer
	// AppImageConfigs returns a AppImageConfigInformer.
	AppImageConfigs() AppImageConfigInformer
	// CodeRepositories returns a CodeRepositoryInformer.
	CodeRepositories() CodeRepositoryInformer
	// Devices returns a DeviceInformer.
	Devices() DeviceInformer
	// DeviceFleets returns a DeviceFleetInformer.
	DeviceFleets() DeviceFleetInformer
	// Domains returns a DomainInformer.
	Domains() DomainInformer
	// Endpoints returns a EndpointInformer.
	Endpoints() EndpointInformer
	// EndpointConfigurations returns a EndpointConfigurationInformer.
	EndpointConfigurations() EndpointConfigurationInformer
	// FeatureGroups returns a FeatureGroupInformer.
	FeatureGroups() FeatureGroupInformer
	// FlowDefinitions returns a FlowDefinitionInformer.
	FlowDefinitions() FlowDefinitionInformer
	// HumanTaskUis returns a HumanTaskUiInformer.
	HumanTaskUis() HumanTaskUiInformer
	// Images returns a ImageInformer.
	Images() ImageInformer
	// ImageVersions returns a ImageVersionInformer.
	ImageVersions() ImageVersionInformer
	// Models returns a ModelInformer.
	Models() ModelInformer
	// ModelPackageGroups returns a ModelPackageGroupInformer.
	ModelPackageGroups() ModelPackageGroupInformer
	// ModelPackageGroupPolicies returns a ModelPackageGroupPolicyInformer.
	ModelPackageGroupPolicies() ModelPackageGroupPolicyInformer
	// NotebookInstances returns a NotebookInstanceInformer.
	NotebookInstances() NotebookInstanceInformer
	// NotebookInstanceLifecycleConfigurations returns a NotebookInstanceLifecycleConfigurationInformer.
	NotebookInstanceLifecycleConfigurations() NotebookInstanceLifecycleConfigurationInformer
	// Projects returns a ProjectInformer.
	Projects() ProjectInformer
	// StudioLifecycleConfigs returns a StudioLifecycleConfigInformer.
	StudioLifecycleConfigs() StudioLifecycleConfigInformer
	// UserProfiles returns a UserProfileInformer.
	UserProfiles() UserProfileInformer
	// Workforces returns a WorkforceInformer.
	Workforces() WorkforceInformer
	// Workteams returns a WorkteamInformer.
	Workteams() WorkteamInformer
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

// Apps returns a AppInformer.
func (v *version) Apps() AppInformer {
	return &appInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AppImageConfigs returns a AppImageConfigInformer.
func (v *version) AppImageConfigs() AppImageConfigInformer {
	return &appImageConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CodeRepositories returns a CodeRepositoryInformer.
func (v *version) CodeRepositories() CodeRepositoryInformer {
	return &codeRepositoryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Devices returns a DeviceInformer.
func (v *version) Devices() DeviceInformer {
	return &deviceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DeviceFleets returns a DeviceFleetInformer.
func (v *version) DeviceFleets() DeviceFleetInformer {
	return &deviceFleetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Domains returns a DomainInformer.
func (v *version) Domains() DomainInformer {
	return &domainInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Endpoints returns a EndpointInformer.
func (v *version) Endpoints() EndpointInformer {
	return &endpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EndpointConfigurations returns a EndpointConfigurationInformer.
func (v *version) EndpointConfigurations() EndpointConfigurationInformer {
	return &endpointConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FeatureGroups returns a FeatureGroupInformer.
func (v *version) FeatureGroups() FeatureGroupInformer {
	return &featureGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FlowDefinitions returns a FlowDefinitionInformer.
func (v *version) FlowDefinitions() FlowDefinitionInformer {
	return &flowDefinitionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HumanTaskUis returns a HumanTaskUiInformer.
func (v *version) HumanTaskUis() HumanTaskUiInformer {
	return &humanTaskUiInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Images returns a ImageInformer.
func (v *version) Images() ImageInformer {
	return &imageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ImageVersions returns a ImageVersionInformer.
func (v *version) ImageVersions() ImageVersionInformer {
	return &imageVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Models returns a ModelInformer.
func (v *version) Models() ModelInformer {
	return &modelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ModelPackageGroups returns a ModelPackageGroupInformer.
func (v *version) ModelPackageGroups() ModelPackageGroupInformer {
	return &modelPackageGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ModelPackageGroupPolicies returns a ModelPackageGroupPolicyInformer.
func (v *version) ModelPackageGroupPolicies() ModelPackageGroupPolicyInformer {
	return &modelPackageGroupPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NotebookInstances returns a NotebookInstanceInformer.
func (v *version) NotebookInstances() NotebookInstanceInformer {
	return &notebookInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NotebookInstanceLifecycleConfigurations returns a NotebookInstanceLifecycleConfigurationInformer.
func (v *version) NotebookInstanceLifecycleConfigurations() NotebookInstanceLifecycleConfigurationInformer {
	return &notebookInstanceLifecycleConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Projects returns a ProjectInformer.
func (v *version) Projects() ProjectInformer {
	return &projectInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StudioLifecycleConfigs returns a StudioLifecycleConfigInformer.
func (v *version) StudioLifecycleConfigs() StudioLifecycleConfigInformer {
	return &studioLifecycleConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UserProfiles returns a UserProfileInformer.
func (v *version) UserProfiles() UserProfileInformer {
	return &userProfileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Workforces returns a WorkforceInformer.
func (v *version) Workforces() WorkforceInformer {
	return &workforceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Workteams returns a WorkteamInformer.
func (v *version) Workteams() WorkteamInformer {
	return &workteamInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
