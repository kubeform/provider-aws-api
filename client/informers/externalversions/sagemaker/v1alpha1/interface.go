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
	// Domains returns a DomainInformer.
	Domains() DomainInformer
	// Endpoints returns a EndpointInformer.
	Endpoints() EndpointInformer
	// EndpointConfigurations returns a EndpointConfigurationInformer.
	EndpointConfigurations() EndpointConfigurationInformer
	// FeatureGroups returns a FeatureGroupInformer.
	FeatureGroups() FeatureGroupInformer
	// Images returns a ImageInformer.
	Images() ImageInformer
	// ImageVersions returns a ImageVersionInformer.
	ImageVersions() ImageVersionInformer
	// Models returns a ModelInformer.
	Models() ModelInformer
	// ModelPackageGroups returns a ModelPackageGroupInformer.
	ModelPackageGroups() ModelPackageGroupInformer
	// NotebookInstances returns a NotebookInstanceInformer.
	NotebookInstances() NotebookInstanceInformer
	// NotebookInstanceLifecycleConfigurations returns a NotebookInstanceLifecycleConfigurationInformer.
	NotebookInstanceLifecycleConfigurations() NotebookInstanceLifecycleConfigurationInformer
	// UserProfiles returns a UserProfileInformer.
	UserProfiles() UserProfileInformer
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

// NotebookInstances returns a NotebookInstanceInformer.
func (v *version) NotebookInstances() NotebookInstanceInformer {
	return &notebookInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NotebookInstanceLifecycleConfigurations returns a NotebookInstanceLifecycleConfigurationInformer.
func (v *version) NotebookInstanceLifecycleConfigurations() NotebookInstanceLifecycleConfigurationInformer {
	return &notebookInstanceLifecycleConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UserProfiles returns a UserProfileInformer.
func (v *version) UserProfiles() UserProfileInformer {
	return &userProfileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
