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

type Parameter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ParameterSpec   `json:"spec,omitempty"`
	Status            ParameterStatus `json:"status,omitempty"`
}

type ParameterSpec struct {
	KubeformOutput *ParameterSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ParameterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
}

type ParameterSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllowedPattern *string `json:"allowedPattern,omitempty" tf:"allowed_pattern"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	DataType *string `json:"dataType,omitempty" tf:"data_type"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	KeyID *string `json:"keyID,omitempty" tf:"key_id"`
	Name  *string `json:"name" tf:"name"`
	// +optional
	Overwrite *bool `json:"overwrite,omitempty" tf:"overwrite"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Tier  *string `json:"tier,omitempty" tf:"tier"`
	Type  *string `json:"type" tf:"type"`
	Value *string `json:"-" sensitive:"true" tf:"value"`
	// +optional
	Version *int64 `json:"version,omitempty" tf:"version"`
}

type ParameterStatus struct {
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

// ParameterList is a list of Parameters
type ParameterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Parameter CRD objects
	Items []Parameter `json:"items,omitempty"`
}
