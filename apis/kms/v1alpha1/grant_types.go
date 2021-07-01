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

type Grant struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GrantSpec   `json:"spec,omitempty"`
	Status            GrantStatus `json:"status,omitempty"`
}

type GrantSpecConstraints struct {
	// +optional
	EncryptionContextEquals *map[string]string `json:"encryptionContextEquals,omitempty" tf:"encryption_context_equals"`
	// +optional
	EncryptionContextSubset *map[string]string `json:"encryptionContextSubset,omitempty" tf:"encryption_context_subset"`
}

type GrantSpec struct {
	KubeformOutput *GrantSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource GrantSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GrantSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Constraints []GrantSpecConstraints `json:"constraints,omitempty" tf:"constraints"`
	// +optional
	GrantCreationTokens []string `json:"grantCreationTokens,omitempty" tf:"grant_creation_tokens"`
	// +optional
	GrantID *string `json:"grantID,omitempty" tf:"grant_id"`
	// +optional
	GrantToken       *string `json:"grantToken,omitempty" tf:"grant_token"`
	GranteePrincipal *string `json:"granteePrincipal" tf:"grantee_principal"`
	KeyID            *string `json:"keyID" tf:"key_id"`
	// +optional
	Name       *string  `json:"name,omitempty" tf:"name"`
	Operations []string `json:"operations" tf:"operations"`
	// +optional
	RetireOnDelete *bool `json:"retireOnDelete,omitempty" tf:"retire_on_delete"`
	// +optional
	RetiringPrincipal *string `json:"retiringPrincipal,omitempty" tf:"retiring_principal"`
}

type GrantStatus struct {
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

// GrantList is a list of Grants
type GrantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Grant CRD objects
	Items []Grant `json:"items,omitempty"`
}
