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

type Fleet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FleetSpec   `json:"spec,omitempty"`
	Status            FleetStatus `json:"status,omitempty"`
}

type FleetSpecLaunchTemplateConfigLaunchTemplateSpecification struct {
	// +optional
	LaunchTemplateID *string `json:"launchTemplateID,omitempty" tf:"launch_template_id"`
	// +optional
	LaunchTemplateName *string `json:"launchTemplateName,omitempty" tf:"launch_template_name"`
	Version            *string `json:"version" tf:"version"`
}

type FleetSpecLaunchTemplateConfigOverride struct {
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	InstanceType *string `json:"instanceType,omitempty" tf:"instance_type"`
	// +optional
	MaxPrice *string `json:"maxPrice,omitempty" tf:"max_price"`
	// +optional
	Priority *float64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	SubnetID *string `json:"subnetID,omitempty" tf:"subnet_id"`
	// +optional
	WeightedCapacity *float64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type FleetSpecLaunchTemplateConfig struct {
	LaunchTemplateSpecification *FleetSpecLaunchTemplateConfigLaunchTemplateSpecification `json:"launchTemplateSpecification" tf:"launch_template_specification"`
	// +optional
	// +kubebuilder:validation:MaxItems=50
	Override []FleetSpecLaunchTemplateConfigOverride `json:"override,omitempty" tf:"override"`
}

type FleetSpecOnDemandOptions struct {
	// +optional
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy"`
}

type FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance struct {
	// +optional
	ReplacementStrategy *string `json:"replacementStrategy,omitempty" tf:"replacement_strategy"`
}

type FleetSpecSpotOptionsMaintenanceStrategies struct {
	// +optional
	CapacityRebalance *FleetSpecSpotOptionsMaintenanceStrategiesCapacityRebalance `json:"capacityRebalance,omitempty" tf:"capacity_rebalance"`
}

type FleetSpecSpotOptions struct {
	// +optional
	AllocationStrategy *string `json:"allocationStrategy,omitempty" tf:"allocation_strategy"`
	// +optional
	InstanceInterruptionBehavior *string `json:"instanceInterruptionBehavior,omitempty" tf:"instance_interruption_behavior"`
	// +optional
	InstancePoolsToUseCount *int64 `json:"instancePoolsToUseCount,omitempty" tf:"instance_pools_to_use_count"`
	// +optional
	MaintenanceStrategies *FleetSpecSpotOptionsMaintenanceStrategies `json:"maintenanceStrategies,omitempty" tf:"maintenance_strategies"`
}

type FleetSpecTargetCapacitySpecification struct {
	DefaultTargetCapacityType *string `json:"defaultTargetCapacityType" tf:"default_target_capacity_type"`
	// +optional
	OnDemandTargetCapacity *int64 `json:"onDemandTargetCapacity,omitempty" tf:"on_demand_target_capacity"`
	// +optional
	SpotTargetCapacity  *int64 `json:"spotTargetCapacity,omitempty" tf:"spot_target_capacity"`
	TotalTargetCapacity *int64 `json:"totalTargetCapacity" tf:"total_target_capacity"`
}

type FleetSpec struct {
	State *FleetSpecResource `json:"state,omitempty" tf:"-"`

	Resource FleetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FleetSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ExcessCapacityTerminationPolicy *string                        `json:"excessCapacityTerminationPolicy,omitempty" tf:"excess_capacity_termination_policy"`
	LaunchTemplateConfig            *FleetSpecLaunchTemplateConfig `json:"launchTemplateConfig" tf:"launch_template_config"`
	// +optional
	OnDemandOptions *FleetSpecOnDemandOptions `json:"onDemandOptions,omitempty" tf:"on_demand_options"`
	// +optional
	ReplaceUnhealthyInstances *bool `json:"replaceUnhealthyInstances,omitempty" tf:"replace_unhealthy_instances"`
	// +optional
	SpotOptions *FleetSpecSpotOptions `json:"spotOptions,omitempty" tf:"spot_options"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll                     *map[string]string                    `json:"tagsAll,omitempty" tf:"tags_all"`
	TargetCapacitySpecification *FleetSpecTargetCapacitySpecification `json:"targetCapacitySpecification" tf:"target_capacity_specification"`
	// +optional
	TerminateInstances *bool `json:"terminateInstances,omitempty" tf:"terminate_instances"`
	// +optional
	TerminateInstancesWithExpiration *bool `json:"terminateInstancesWithExpiration,omitempty" tf:"terminate_instances_with_expiration"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type FleetStatus struct {
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

// FleetList is a list of Fleets
type FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Fleet CRD objects
	Items []Fleet `json:"items,omitempty"`
}
