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
	// Accounts returns a AccountInformer.
	Accounts() AccountInformer
	// ApiKeys returns a ApiKeyInformer.
	ApiKeys() ApiKeyInformer
	// ApigatewayResources returns a ApigatewayResourceInformer.
	ApigatewayResources() ApigatewayResourceInformer
	// Authorizers returns a AuthorizerInformer.
	Authorizers() AuthorizerInformer
	// BasePathMappings returns a BasePathMappingInformer.
	BasePathMappings() BasePathMappingInformer
	// ClientCertificates returns a ClientCertificateInformer.
	ClientCertificates() ClientCertificateInformer
	// Deployments returns a DeploymentInformer.
	Deployments() DeploymentInformer
	// DocumentationParts returns a DocumentationPartInformer.
	DocumentationParts() DocumentationPartInformer
	// DocumentationVersions returns a DocumentationVersionInformer.
	DocumentationVersions() DocumentationVersionInformer
	// DomainNames returns a DomainNameInformer.
	DomainNames() DomainNameInformer
	// GatewayResponses returns a GatewayResponseInformer.
	GatewayResponses() GatewayResponseInformer
	// Integrations returns a IntegrationInformer.
	Integrations() IntegrationInformer
	// IntegrationResponses returns a IntegrationResponseInformer.
	IntegrationResponses() IntegrationResponseInformer
	// Methods returns a MethodInformer.
	Methods() MethodInformer
	// MethodResponses returns a MethodResponseInformer.
	MethodResponses() MethodResponseInformer
	// MethodSettingses returns a MethodSettingsInformer.
	MethodSettingses() MethodSettingsInformer
	// Models returns a ModelInformer.
	Models() ModelInformer
	// RequestValidators returns a RequestValidatorInformer.
	RequestValidators() RequestValidatorInformer
	// RestAPIs returns a RestAPIInformer.
	RestAPIs() RestAPIInformer
	// RestAPIPolicies returns a RestAPIPolicyInformer.
	RestAPIPolicies() RestAPIPolicyInformer
	// Stages returns a StageInformer.
	Stages() StageInformer
	// UsagePlans returns a UsagePlanInformer.
	UsagePlans() UsagePlanInformer
	// UsagePlanKeys returns a UsagePlanKeyInformer.
	UsagePlanKeys() UsagePlanKeyInformer
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

// Accounts returns a AccountInformer.
func (v *version) Accounts() AccountInformer {
	return &accountInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApiKeys returns a ApiKeyInformer.
func (v *version) ApiKeys() ApiKeyInformer {
	return &apiKeyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApigatewayResources returns a ApigatewayResourceInformer.
func (v *version) ApigatewayResources() ApigatewayResourceInformer {
	return &apigatewayResourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Authorizers returns a AuthorizerInformer.
func (v *version) Authorizers() AuthorizerInformer {
	return &authorizerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BasePathMappings returns a BasePathMappingInformer.
func (v *version) BasePathMappings() BasePathMappingInformer {
	return &basePathMappingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClientCertificates returns a ClientCertificateInformer.
func (v *version) ClientCertificates() ClientCertificateInformer {
	return &clientCertificateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Deployments returns a DeploymentInformer.
func (v *version) Deployments() DeploymentInformer {
	return &deploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DocumentationParts returns a DocumentationPartInformer.
func (v *version) DocumentationParts() DocumentationPartInformer {
	return &documentationPartInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DocumentationVersions returns a DocumentationVersionInformer.
func (v *version) DocumentationVersions() DocumentationVersionInformer {
	return &documentationVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainNames returns a DomainNameInformer.
func (v *version) DomainNames() DomainNameInformer {
	return &domainNameInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GatewayResponses returns a GatewayResponseInformer.
func (v *version) GatewayResponses() GatewayResponseInformer {
	return &gatewayResponseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Integrations returns a IntegrationInformer.
func (v *version) Integrations() IntegrationInformer {
	return &integrationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IntegrationResponses returns a IntegrationResponseInformer.
func (v *version) IntegrationResponses() IntegrationResponseInformer {
	return &integrationResponseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Methods returns a MethodInformer.
func (v *version) Methods() MethodInformer {
	return &methodInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MethodResponses returns a MethodResponseInformer.
func (v *version) MethodResponses() MethodResponseInformer {
	return &methodResponseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MethodSettingses returns a MethodSettingsInformer.
func (v *version) MethodSettingses() MethodSettingsInformer {
	return &methodSettingsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Models returns a ModelInformer.
func (v *version) Models() ModelInformer {
	return &modelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RequestValidators returns a RequestValidatorInformer.
func (v *version) RequestValidators() RequestValidatorInformer {
	return &requestValidatorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RestAPIs returns a RestAPIInformer.
func (v *version) RestAPIs() RestAPIInformer {
	return &restAPIInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RestAPIPolicies returns a RestAPIPolicyInformer.
func (v *version) RestAPIPolicies() RestAPIPolicyInformer {
	return &restAPIPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Stages returns a StageInformer.
func (v *version) Stages() StageInformer {
	return &stageInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UsagePlans returns a UsagePlanInformer.
func (v *version) UsagePlans() UsagePlanInformer {
	return &usagePlanInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UsagePlanKeys returns a UsagePlanKeyInformer.
func (v *version) UsagePlanKeys() UsagePlanKeyInformer {
	return &usagePlanKeyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VpcLinks returns a VpcLinkInformer.
func (v *version) VpcLinks() VpcLinkInformer {
	return &vpcLinkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
