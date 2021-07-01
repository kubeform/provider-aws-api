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
	// AssessmentTargets returns a AssessmentTargetInformer.
	AssessmentTargets() AssessmentTargetInformer
	// AssessmentTemplates returns a AssessmentTemplateInformer.
	AssessmentTemplates() AssessmentTemplateInformer
	// ResourceGroups returns a ResourceGroupInformer.
	ResourceGroups() ResourceGroupInformer
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

// AssessmentTargets returns a AssessmentTargetInformer.
func (v *version) AssessmentTargets() AssessmentTargetInformer {
	return &assessmentTargetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AssessmentTemplates returns a AssessmentTemplateInformer.
func (v *version) AssessmentTemplates() AssessmentTemplateInformer {
	return &assessmentTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ResourceGroups returns a ResourceGroupInformer.
func (v *version) ResourceGroups() ResourceGroupInformer {
	return &resourceGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
