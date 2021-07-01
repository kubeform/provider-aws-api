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
	// Apis returns a ApiInformer.
	Apis() ApiInformer
	// ApiMappings returns a ApiMappingInformer.
	ApiMappings() ApiMappingInformer
	// Authorizers returns a AuthorizerInformer.
	Authorizers() AuthorizerInformer
	// Deployments returns a DeploymentInformer.
	Deployments() DeploymentInformer
	// DomainNames returns a DomainNameInformer.
	DomainNames() DomainNameInformer
	// Integrations returns a IntegrationInformer.
	Integrations() IntegrationInformer
	// IntegrationResponses returns a IntegrationResponseInformer.
	IntegrationResponses() IntegrationResponseInformer
	// Models returns a ModelInformer.
	Models() ModelInformer
	// Routes returns a RouteInformer.
	Routes() RouteInformer
	// RouteResponses returns a RouteResponseInformer.
	RouteResponses() RouteResponseInformer
	// Stages returns a StageInformer.
	Stages() StageInformer
	// VpcLinks returns a VpcLinkInformer.
	VpcLinks() VpcLinkInformer
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

// Apis returns a ApiInformer.
func (v *version) Apis() ApiInformer {
	return &apiInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiMappings returns a ApiMappingInformer.
func (v *version) ApiMappings() ApiMappingInformer {
	return &apiMappingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Authorizers returns a AuthorizerInformer.
func (v *version) Authorizers() AuthorizerInformer {
	return &authorizerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Deployments returns a DeploymentInformer.
func (v *version) Deployments() DeploymentInformer {
	return &deploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainNames returns a DomainNameInformer.
func (v *version) DomainNames() DomainNameInformer {
	return &domainNameInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Integrations returns a IntegrationInformer.
func (v *version) Integrations() IntegrationInformer {
	return &integrationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IntegrationResponses returns a IntegrationResponseInformer.
func (v *version) IntegrationResponses() IntegrationResponseInformer {
	return &integrationResponseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Models returns a ModelInformer.
func (v *version) Models() ModelInformer {
	return &modelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Routes returns a RouteInformer.
func (v *version) Routes() RouteInformer {
	return &routeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RouteResponses returns a RouteResponseInformer.
func (v *version) RouteResponses() RouteResponseInformer {
	return &routeResponseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Stages returns a StageInformer.
func (v *version) Stages() StageInformer {
	return &stageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VpcLinks returns a VpcLinkInformer.
func (v *version) VpcLinks() VpcLinkInformer {
	return &vpcLinkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
