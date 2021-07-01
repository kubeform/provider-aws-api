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

type GameSessionQueue struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GameSessionQueueSpec   `json:"spec,omitempty"`
	Status            GameSessionQueueStatus `json:"status,omitempty"`
}

type GameSessionQueueSpecPlayerLatencyPolicy struct {
	MaximumIndividualPlayerLatencyMilliseconds *int64 `json:"maximumIndividualPlayerLatencyMilliseconds" tf:"maximum_individual_player_latency_milliseconds"`
	// +optional
	PolicyDurationSeconds *int64 `json:"policyDurationSeconds,omitempty" tf:"policy_duration_seconds"`
}

type GameSessionQueueSpec struct {
	KubeformOutput *GameSessionQueueSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource GameSessionQueueSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type GameSessionQueueSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Destinations []string `json:"destinations,omitempty" tf:"destinations"`
	Name         *string  `json:"name" tf:"name"`
	// +optional
	PlayerLatencyPolicy []GameSessionQueueSpecPlayerLatencyPolicy `json:"playerLatencyPolicy,omitempty" tf:"player_latency_policy"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	TimeoutInSeconds *int64 `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds"`
}

type GameSessionQueueStatus struct {
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

// GameSessionQueueList is a list of GameSessionQueues
type GameSessionQueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GameSessionQueue CRD objects
	Items []GameSessionQueue `json:"items,omitempty"`
}
