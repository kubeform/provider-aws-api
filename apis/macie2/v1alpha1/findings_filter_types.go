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

type FindingsFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FindingsFilterSpec   `json:"spec,omitempty"`
	Status            FindingsFilterStatus `json:"status,omitempty"`
}

type FindingsFilterSpecFindingCriteriaCriterion struct {
	// +optional
	Eq []string `json:"eq,omitempty" tf:"eq"`
	// +optional
	EqExactMatch []string `json:"eqExactMatch,omitempty" tf:"eq_exact_match"`
	Field        *string  `json:"field" tf:"field"`
	// +optional
	Gt *string `json:"gt,omitempty" tf:"gt"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lt *string `json:"lt,omitempty" tf:"lt"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
	// +optional
	Neq []string `json:"neq,omitempty" tf:"neq"`
}

type FindingsFilterSpecFindingCriteria struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	Criterion []FindingsFilterSpecFindingCriteriaCriterion `json:"criterion,omitempty" tf:"criterion"`
}

type FindingsFilterSpec struct {
	State *FindingsFilterSpecResource `json:"state,omitempty" tf:"-"`

	Resource FindingsFilterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FindingsFilterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Action *string `json:"action" tf:"action"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description     *string                            `json:"description,omitempty" tf:"description"`
	FindingCriteria *FindingsFilterSpecFindingCriteria `json:"findingCriteria" tf:"finding_criteria"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	Position *int64 `json:"position,omitempty" tf:"position"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type FindingsFilterStatus struct {
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

// FindingsFilterList is a list of FindingsFilters
type FindingsFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FindingsFilter CRD objects
	Items []FindingsFilter `json:"items,omitempty"`
}
