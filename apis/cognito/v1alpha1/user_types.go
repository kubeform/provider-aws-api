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

type User struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              UserSpec   `json:"spec,omitempty"`
	Status            UserStatus `json:"status,omitempty"`
}

type UserSpec struct {
	State *UserSpecResource `json:"state,omitempty" tf:"-"`

	Resource UserSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type UserSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Attributes *map[string]string `json:"attributes,omitempty" tf:"attributes"`
	// +optional
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty" tf:"client_metadata"`
	// +optional
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date"`
	// +optional
	DesiredDeliveryMediums []string `json:"desiredDeliveryMediums,omitempty" tf:"desired_delivery_mediums"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	ForceAliasCreation *bool `json:"forceAliasCreation,omitempty" tf:"force_alias_creation"`
	// +optional
	LastModifiedDate *string `json:"lastModifiedDate,omitempty" tf:"last_modified_date"`
	// +optional
	MessageAction *string `json:"messageAction,omitempty" tf:"message_action"`
	// +optional
	MfaSettingList []string `json:"mfaSettingList,omitempty" tf:"mfa_setting_list"`
	// +optional
	Password *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	PreferredMfaSetting *string `json:"preferredMfaSetting,omitempty" tf:"preferred_mfa_setting"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Sub *string `json:"sub,omitempty" tf:"sub"`
	// +optional
	TemporaryPassword *string `json:"-" sensitive:"true" tf:"temporary_password"`
	UserPoolID        *string `json:"userPoolID" tf:"user_pool_id"`
	Username          *string `json:"username" tf:"username"`
	// +optional
	ValidationData *map[string]string `json:"validationData,omitempty" tf:"validation_data"`
}

type UserStatus struct {
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

// UserList is a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of User CRD objects
	Items []User `json:"items,omitempty"`
}