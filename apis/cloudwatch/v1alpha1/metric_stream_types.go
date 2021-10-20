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

type MetricStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MetricStreamSpec   `json:"spec,omitempty"`
	Status            MetricStreamStatus `json:"status,omitempty"`
}

type MetricStreamSpecExcludeFilter struct {
	Namespace *string `json:"namespace" tf:"namespace"`
}

type MetricStreamSpecIncludeFilter struct {
	Namespace *string `json:"namespace" tf:"namespace"`
}

type MetricStreamSpec struct {
	State *MetricStreamSpecResource `json:"state,omitempty" tf:"-"`

	Resource MetricStreamSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MetricStreamSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CreationDate *string `json:"creationDate,omitempty" tf:"creation_date"`
	// +optional
	ExcludeFilter []MetricStreamSpecExcludeFilter `json:"excludeFilter,omitempty" tf:"exclude_filter"`
	FirehoseArn   *string                         `json:"firehoseArn" tf:"firehose_arn"`
	// +optional
	IncludeFilter []MetricStreamSpecIncludeFilter `json:"includeFilter,omitempty" tf:"include_filter"`
	// +optional
	LastUpdateDate *string `json:"lastUpdateDate,omitempty" tf:"last_update_date"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix   *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	OutputFormat *string `json:"outputFormat" tf:"output_format"`
	RoleArn      *string `json:"roleArn" tf:"role_arn"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type MetricStreamStatus struct {
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

// MetricStreamList is a list of MetricStreams
type MetricStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MetricStream CRD objects
	Items []MetricStream `json:"items,omitempty"`
}
