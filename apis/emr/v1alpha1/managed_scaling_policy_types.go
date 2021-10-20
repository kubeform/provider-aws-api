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

type ManagedScalingPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ManagedScalingPolicySpec   `json:"spec,omitempty"`
	Status            ManagedScalingPolicyStatus `json:"status,omitempty"`
}

type ManagedScalingPolicySpecComputeLimits struct {
	MaximumCapacityUnits *int64 `json:"maximumCapacityUnits" tf:"maximum_capacity_units"`
	// +optional
	MaximumCoreCapacityUnits *int64 `json:"maximumCoreCapacityUnits,omitempty" tf:"maximum_core_capacity_units"`
	// +optional
	MaximumOndemandCapacityUnits *int64  `json:"maximumOndemandCapacityUnits,omitempty" tf:"maximum_ondemand_capacity_units"`
	MinimumCapacityUnits         *int64  `json:"minimumCapacityUnits" tf:"minimum_capacity_units"`
	UnitType                     *string `json:"unitType" tf:"unit_type"`
}

type ManagedScalingPolicySpec struct {
	State *ManagedScalingPolicySpecResource `json:"state,omitempty" tf:"-"`

	Resource ManagedScalingPolicySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ManagedScalingPolicySpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClusterID     *string                                 `json:"clusterID" tf:"cluster_id"`
	ComputeLimits []ManagedScalingPolicySpecComputeLimits `json:"computeLimits" tf:"compute_limits"`
}

type ManagedScalingPolicyStatus struct {
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

// ManagedScalingPolicyList is a list of ManagedScalingPolicys
type ManagedScalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ManagedScalingPolicy CRD objects
	Items []ManagedScalingPolicy `json:"items,omitempty"`
}
