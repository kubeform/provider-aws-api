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

type RouteTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec,omitempty"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

type RouteTableSpecRoute struct {
	// +optional
	CidrBlock *string `json:"cidrBlock,omitempty" tf:"cidr_block"`
	// +optional
	DestinationPrefixListID *string `json:"destinationPrefixListID,omitempty" tf:"destination_prefix_list_id"`
	// +optional
	EgressOnlyGatewayID *string `json:"egressOnlyGatewayID,omitempty" tf:"egress_only_gateway_id"`
	// +optional
	GatewayID *string `json:"gatewayID,omitempty" tf:"gateway_id"`
	// +optional
	InstanceID *string `json:"instanceID,omitempty" tf:"instance_id"`
	// +optional
	Ipv6CIDRBlock *string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block"`
	// +optional
	NatGatewayID *string `json:"natGatewayID,omitempty" tf:"nat_gateway_id"`
	// +optional
	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty" tf:"network_interface_id"`
	// +optional
	TransitGatewayID *string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id"`
	// +optional
	VpcEndpointID *string `json:"vpcEndpointID,omitempty" tf:"vpc_endpoint_id"`
	// +optional
	VpcPeeringConnectionID *string `json:"vpcPeeringConnectionID,omitempty" tf:"vpc_peering_connection_id"`
}

type RouteTableSpec struct {
	State *RouteTableSpecResource `json:"state,omitempty" tf:"-"`

	Resource RouteTableSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type RouteTableSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                 *string `json:"arn,omitempty" tf:"arn"`
	DefaultRouteTableID *string `json:"defaultRouteTableID" tf:"default_route_table_id"`
	// +optional
	OwnerID *string `json:"ownerID,omitempty" tf:"owner_id"`
	// +optional
	PropagatingVgws []string `json:"propagatingVgws,omitempty" tf:"propagating_vgws"`
	// +optional
	Route []RouteTableSpecRoute `json:"route,omitempty" tf:"route"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type RouteTableStatus struct {
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

// RouteTableList is a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RouteTable CRD objects
	Items []RouteTable `json:"items,omitempty"`
}
