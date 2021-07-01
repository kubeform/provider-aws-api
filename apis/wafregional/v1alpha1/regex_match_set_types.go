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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type RegexMatchSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegexMatchSetSpec   `json:"spec,omitempty"`
	Status            RegexMatchSetStatus `json:"status,omitempty"`
}

type RegexMatchSetSpecRegexMatchTupleFieldToMatch struct {
	// +optional
	Data *string `json:"data,omitempty" tf:"data"`
	Type *string `json:"type" tf:"type"`
}

type RegexMatchSetSpecRegexMatchTuple struct {
	FieldToMatch       *RegexMatchSetSpecRegexMatchTupleFieldToMatch `json:"fieldToMatch" tf:"field_to_match"`
	RegexPatternSetID  *string                                       `json:"regexPatternSetID" tf:"regex_pattern_set_id"`
	TextTransformation *string                                       `json:"textTransformation" tf:"text_transformation"`
}

type RegexMatchSetSpec struct {
	KubeformOutput *RegexMatchSetSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource RegexMatchSetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RegexMatchSetSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Name *string `json:"name" tf:"name"`
	// +optional
	RegexMatchTuple []RegexMatchSetSpecRegexMatchTuple `json:"regexMatchTuple,omitempty" tf:"regex_match_tuple"`
}

type RegexMatchSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RegexMatchSetList is a list of RegexMatchSets
type RegexMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RegexMatchSet CRD objects
	Items []RegexMatchSet `json:"items,omitempty"`
}
