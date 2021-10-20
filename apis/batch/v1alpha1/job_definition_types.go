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

type JobDefinition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobDefinitionSpec   `json:"spec,omitempty"`
	Status            JobDefinitionStatus `json:"status,omitempty"`
}

type JobDefinitionSpecRetryStrategyEvaluateOnExit struct {
	Action *string `json:"action" tf:"action"`
	// +optional
	OnExitCode *string `json:"onExitCode,omitempty" tf:"on_exit_code"`
	// +optional
	OnReason *string `json:"onReason,omitempty" tf:"on_reason"`
	// +optional
	OnStatusReason *string `json:"onStatusReason,omitempty" tf:"on_status_reason"`
}

type JobDefinitionSpecRetryStrategy struct {
	// +optional
	Attempts *int64 `json:"attempts,omitempty" tf:"attempts"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	EvaluateOnExit []JobDefinitionSpecRetryStrategyEvaluateOnExit `json:"evaluateOnExit,omitempty" tf:"evaluate_on_exit"`
}

type JobDefinitionSpecTimeout struct {
	// +optional
	AttemptDurationSeconds *int64 `json:"attemptDurationSeconds,omitempty" tf:"attempt_duration_seconds"`
}

type JobDefinitionSpec struct {
	State *JobDefinitionSpecResource `json:"state,omitempty" tf:"-"`

	Resource JobDefinitionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type JobDefinitionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ContainerProperties *string `json:"containerProperties,omitempty" tf:"container_properties"`
	Name                *string `json:"name" tf:"name"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	PlatformCapabilities []string `json:"platformCapabilities,omitempty" tf:"platform_capabilities"`
	// +optional
	PropagateTags *bool `json:"propagateTags,omitempty" tf:"propagate_tags"`
	// +optional
	RetryStrategy *JobDefinitionSpecRetryStrategy `json:"retryStrategy,omitempty" tf:"retry_strategy"`
	// +optional
	Revision *int64 `json:"revision,omitempty" tf:"revision"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Timeout *JobDefinitionSpecTimeout `json:"timeout,omitempty" tf:"timeout"`
	Type    *string                   `json:"type" tf:"type"`
}

type JobDefinitionStatus struct {
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

// JobDefinitionList is a list of JobDefinitions
type JobDefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of JobDefinition CRD objects
	Items []JobDefinition `json:"items,omitempty"`
}
