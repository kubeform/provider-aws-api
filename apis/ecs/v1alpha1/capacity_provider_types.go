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

type CapacityProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CapacityProviderSpec   `json:"spec,omitempty"`
	Status            CapacityProviderStatus `json:"status,omitempty"`
}

type CapacityProviderSpecAutoScalingGroupProviderManagedScaling struct {
	// +optional
	InstanceWarmupPeriod *int64 `json:"instanceWarmupPeriod,omitempty" tf:"instance_warmup_period"`
	// +optional
	MaximumScalingStepSize *int64 `json:"maximumScalingStepSize,omitempty" tf:"maximum_scaling_step_size"`
	// +optional
	MinimumScalingStepSize *int64 `json:"minimumScalingStepSize,omitempty" tf:"minimum_scaling_step_size"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	TargetCapacity *int64 `json:"targetCapacity,omitempty" tf:"target_capacity"`
}

type CapacityProviderSpecAutoScalingGroupProvider struct {
	AutoScalingGroupArn *string `json:"autoScalingGroupArn" tf:"auto_scaling_group_arn"`
	// +optional
	ManagedScaling *CapacityProviderSpecAutoScalingGroupProviderManagedScaling `json:"managedScaling,omitempty" tf:"managed_scaling"`
	// +optional
	ManagedTerminationProtection *string `json:"managedTerminationProtection,omitempty" tf:"managed_termination_protection"`
}

type CapacityProviderSpec struct {
	State *CapacityProviderSpecResource `json:"state,omitempty" tf:"-"`

	Resource CapacityProviderSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type CapacityProviderSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                      *string                                       `json:"arn,omitempty" tf:"arn"`
	AutoScalingGroupProvider *CapacityProviderSpecAutoScalingGroupProvider `json:"autoScalingGroupProvider" tf:"auto_scaling_group_provider"`
	Name                     *string                                       `json:"name" tf:"name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type CapacityProviderStatus struct {
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

// CapacityProviderList is a list of CapacityProviders
type CapacityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CapacityProvider CRD objects
	Items []CapacityProvider `json:"items,omitempty"`
}
