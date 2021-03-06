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
	// DelegatedAdministrators returns a DelegatedAdministratorInformer.
	DelegatedAdministrators() DelegatedAdministratorInformer
	// Organizations returns a OrganizationInformer.
	Organizations() OrganizationInformer
	// OrganizationalUnits returns a OrganizationalUnitInformer.
	OrganizationalUnits() OrganizationalUnitInformer
	// Policies returns a PolicyInformer.
	Policies() PolicyInformer
	// PolicyAttachments returns a PolicyAttachmentInformer.
	PolicyAttachments() PolicyAttachmentInformer
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

// DelegatedAdministrators returns a DelegatedAdministratorInformer.
func (v *version) DelegatedAdministrators() DelegatedAdministratorInformer {
	return &delegatedAdministratorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Organizations returns a OrganizationInformer.
func (v *version) Organizations() OrganizationInformer {
	return &organizationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OrganizationalUnits returns a OrganizationalUnitInformer.
func (v *version) OrganizationalUnits() OrganizationalUnitInformer {
	return &organizationalUnitInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Policies returns a PolicyInformer.
func (v *version) Policies() PolicyInformer {
	return &policyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PolicyAttachments returns a PolicyAttachmentInformer.
func (v *version) PolicyAttachments() PolicyAttachmentInformer {
	return &policyAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
