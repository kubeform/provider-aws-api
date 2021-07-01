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

type Firewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallSpec   `json:"spec,omitempty"`
	Status            FirewallStatus `json:"status,omitempty"`
}

type FirewallSpecFirewallStatusSyncStatesAttachment struct {
	// +optional
	EndpointID *string `json:"endpointID,omitempty" tf:"endpoint_id"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
}

type FirewallSpecFirewallStatusSyncStates struct {
	// +optional
	Attachment []FirewallSpecFirewallStatusSyncStatesAttachment `json:"attachment,omitempty" tf:"attachment"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
}

type FirewallSpecFirewallStatus struct {
	// +optional
	SyncStates []FirewallSpecFirewallStatusSyncStates `json:"syncStates,omitempty" tf:"sync_states"`
}

type FirewallSpecSubnetMapping struct {
	SubnetID *string `json:"subnetID" tf:"subnet_id"`
}

type FirewallSpec struct {
	KubeformOutput *FirewallSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource FirewallSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type FirewallSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	DeleteProtection *bool `json:"deleteProtection,omitempty" tf:"delete_protection"`
	// +optional
	Description       *string `json:"description,omitempty" tf:"description"`
	FirewallPolicyArn *string `json:"firewallPolicyArn" tf:"firewall_policy_arn"`
	// +optional
	FirewallPolicyChangeProtection *bool `json:"firewallPolicyChangeProtection,omitempty" tf:"firewall_policy_change_protection"`
	// +optional
	FirewallStatus []FirewallSpecFirewallStatus `json:"firewallStatus,omitempty" tf:"firewall_status"`
	Name           *string                      `json:"name" tf:"name"`
	// +optional
	SubnetChangeProtection *bool                       `json:"subnetChangeProtection,omitempty" tf:"subnet_change_protection"`
	SubnetMapping          []FirewallSpecSubnetMapping `json:"subnetMapping" tf:"subnet_mapping"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	UpdateToken *string `json:"updateToken,omitempty" tf:"update_token"`
	VpcID       *string `json:"vpcID" tf:"vpc_id"`
}

type FirewallStatus struct {
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

// FirewallList is a list of Firewalls
type FirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Firewall CRD objects
	Items []Firewall `json:"items,omitempty"`
}
