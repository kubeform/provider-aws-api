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

type UserSpecHomeDirectoryMappings struct {
	Entry  *string `json:"entry" tf:"entry"`
	Target *string `json:"target" tf:"target"`
}

type UserSpecPosixProfile struct {
	Gid *int64 `json:"gid" tf:"gid"`
	// +optional
	SecondaryGids []int64 `json:"secondaryGids,omitempty" tf:"secondary_gids"`
	Uid           *int64  `json:"uid" tf:"uid"`
}

type UserSpec struct {
	State *UserSpecResource `json:"state,omitempty" tf:"-"`

	Resource UserSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type UserSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	HomeDirectory *string `json:"homeDirectory,omitempty" tf:"home_directory"`
	// +optional
	HomeDirectoryMappings []UserSpecHomeDirectoryMappings `json:"homeDirectoryMappings,omitempty" tf:"home_directory_mappings"`
	// +optional
	HomeDirectoryType *string `json:"homeDirectoryType,omitempty" tf:"home_directory_type"`
	// +optional
	Policy *string `json:"policy,omitempty" tf:"policy"`
	// +optional
	PosixProfile *UserSpecPosixProfile `json:"posixProfile,omitempty" tf:"posix_profile"`
	Role         *string               `json:"role" tf:"role"`
	ServerID     *string               `json:"serverID" tf:"server_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll  *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	UserName *string            `json:"userName" tf:"user_name"`
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
