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

type ResourceShareAccepter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceShareAccepterSpec   `json:"spec,omitempty"`
	Status            ResourceShareAccepterStatus `json:"status,omitempty"`
}

type ResourceShareAccepterSpec struct {
	State *ResourceShareAccepterSpecResource `json:"state,omitempty" tf:"-"`

	Resource ResourceShareAccepterSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ResourceShareAccepterSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	InvitationArn *string `json:"invitationArn,omitempty" tf:"invitation_arn"`
	// +optional
	ReceiverAccountID *string `json:"receiverAccountID,omitempty" tf:"receiver_account_id"`
	// +optional
	Resources []string `json:"resources,omitempty" tf:"resources"`
	// +optional
	SenderAccountID *string `json:"senderAccountID,omitempty" tf:"sender_account_id"`
	ShareArn        *string `json:"shareArn" tf:"share_arn"`
	// +optional
	ShareID *string `json:"shareID,omitempty" tf:"share_id"`
	// +optional
	ShareName *string `json:"shareName,omitempty" tf:"share_name"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
}

type ResourceShareAccepterStatus struct {
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

// ResourceShareAccepterList is a list of ResourceShareAccepters
type ResourceShareAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ResourceShareAccepter CRD objects
	Items []ResourceShareAccepter `json:"items,omitempty"`
}
