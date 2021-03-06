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

type HostedConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedConnectionSpec   `json:"spec,omitempty"`
	Status            HostedConnectionStatus `json:"status,omitempty"`
}

type HostedConnectionSpec struct {
	State *HostedConnectionSpecResource `json:"state,omitempty" tf:"-"`

	Resource HostedConnectionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type HostedConnectionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AwsDevice    *string `json:"awsDevice,omitempty" tf:"aws_device"`
	Bandwidth    *string `json:"bandwidth" tf:"bandwidth"`
	ConnectionID *string `json:"connectionID" tf:"connection_id"`
	// +optional
	HasLogicalRedundancy *string `json:"hasLogicalRedundancy,omitempty" tf:"has_logical_redundancy"`
	// +optional
	JumboFrameCapable *bool `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable"`
	// +optional
	LagID *string `json:"lagID,omitempty" tf:"lag_id"`
	// +optional
	LoaIssueTime *string `json:"loaIssueTime,omitempty" tf:"loa_issue_time"`
	// +optional
	Location       *string `json:"location,omitempty" tf:"location"`
	Name           *string `json:"name" tf:"name"`
	OwnerAccountID *string `json:"ownerAccountID" tf:"owner_account_id"`
	// +optional
	PartnerName *string `json:"partnerName,omitempty" tf:"partner_name"`
	// +optional
	ProviderName *string `json:"providerName,omitempty" tf:"provider_name"`
	// +optional
	Region *string `json:"region,omitempty" tf:"region"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	Vlan  *int64  `json:"vlan" tf:"vlan"`
}

type HostedConnectionStatus struct {
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

// HostedConnectionList is a list of HostedConnections
type HostedConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HostedConnection CRD objects
	Items []HostedConnection `json:"items,omitempty"`
}
