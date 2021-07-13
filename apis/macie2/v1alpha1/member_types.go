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

type Member struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MemberSpec   `json:"spec,omitempty"`
	Status            MemberStatus `json:"status,omitempty"`
}

type MemberSpec struct {
	State *MemberSpecResource `json:"state,omitempty" tf:"-"`

	Resource MemberSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type MemberSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AccountID *string `json:"accountID" tf:"account_id"`
	// +optional
	AdministratorAccountID *string `json:"administratorAccountID,omitempty" tf:"administrator_account_id"`
	// +optional
	Arn   *string `json:"arn,omitempty" tf:"arn"`
	Email *string `json:"email" tf:"email"`
	// +optional
	InvitationDisableEmailNotification *string `json:"invitationDisableEmailNotification,omitempty" tf:"invitation_disable_email_notification"`
	// +optional
	InvitationMessage *string `json:"invitationMessage,omitempty" tf:"invitation_message"`
	// +optional
	Invite *bool `json:"invite,omitempty" tf:"invite"`
	// +optional
	InvitedAt *string `json:"invitedAt,omitempty" tf:"invited_at"`
	// +optional
	MasterAccountID *string `json:"masterAccountID,omitempty" tf:"master_account_id"`
	// +optional
	RelationshipStatus *string `json:"relationshipStatus,omitempty" tf:"relationship_status"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at"`
}

type MemberStatus struct {
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

// MemberList is a list of Members
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Member CRD objects
	Items []Member `json:"items,omitempty"`
}
