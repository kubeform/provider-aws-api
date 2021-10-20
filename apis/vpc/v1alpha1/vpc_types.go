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

type Vpc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcSpec   `json:"spec,omitempty"`
	Status            VpcStatus `json:"status,omitempty"`
}

type VpcSpec struct {
	State *VpcSpecResource `json:"state,omitempty" tf:"-"`

	Resource VpcSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type VpcSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AssignGeneratedIpv6CIDRBlock *bool   `json:"assignGeneratedIpv6CIDRBlock,omitempty" tf:"assign_generated_ipv6_cidr_block"`
	CidrBlock                    *string `json:"cidrBlock" tf:"cidr_block"`
	// +optional
	DefaultNetworkACLID *string `json:"defaultNetworkACLID,omitempty" tf:"default_network_acl_id"`
	// +optional
	DefaultRouteTableID *string `json:"defaultRouteTableID,omitempty" tf:"default_route_table_id"`
	// +optional
	DefaultSecurityGroupID *string `json:"defaultSecurityGroupID,omitempty" tf:"default_security_group_id"`
	// +optional
	DhcpOptionsID *string `json:"dhcpOptionsID,omitempty" tf:"dhcp_options_id"`
	// +optional
	EnableClassiclink *bool `json:"enableClassiclink,omitempty" tf:"enable_classiclink"`
	// +optional
	EnableClassiclinkDNSSupport *bool `json:"enableClassiclinkDNSSupport,omitempty" tf:"enable_classiclink_dns_support"`
	// +optional
	EnableDNSHostnames *bool `json:"enableDNSHostnames,omitempty" tf:"enable_dns_hostnames"`
	// +optional
	EnableDNSSupport *bool `json:"enableDNSSupport,omitempty" tf:"enable_dns_support"`
	// +optional
	InstanceTenancy *string `json:"instanceTenancy,omitempty" tf:"instance_tenancy"`
	// +optional
	Ipv6AssociationID *string `json:"ipv6AssociationID,omitempty" tf:"ipv6_association_id"`
	// +optional
	Ipv6CIDRBlock *string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block"`
	// +optional
	MainRouteTableID *string `json:"mainRouteTableID,omitempty" tf:"main_route_table_id"`
	// +optional
	OwnerID *string `json:"ownerID,omitempty" tf:"owner_id"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type VpcStatus struct {
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

// VpcList is a list of Vpcs
type VpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Vpc CRD objects
	Items []Vpc `json:"items,omitempty"`
}
