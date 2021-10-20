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

type Table struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TableSpec   `json:"spec,omitempty"`
	Status            TableStatus `json:"status,omitempty"`
}

type TableSpecAttribute struct {
	Name *string `json:"name" tf:"name"`
	Type *string `json:"type" tf:"type"`
}

type TableSpecGlobalSecondaryIndex struct {
	HashKey *string `json:"hashKey" tf:"hash_key"`
	Name    *string `json:"name" tf:"name"`
	// +optional
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes"`
	ProjectionType   *string  `json:"projectionType" tf:"projection_type"`
	// +optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key"`
	// +optional
	ReadCapacity *int64 `json:"readCapacity,omitempty" tf:"read_capacity"`
	// +optional
	WriteCapacity *int64 `json:"writeCapacity,omitempty" tf:"write_capacity"`
}

type TableSpecLocalSecondaryIndex struct {
	Name *string `json:"name" tf:"name"`
	// +optional
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" tf:"non_key_attributes"`
	ProjectionType   *string  `json:"projectionType" tf:"projection_type"`
	RangeKey         *string  `json:"rangeKey" tf:"range_key"`
}

type TableSpecPointInTimeRecovery struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
}

type TableSpecReplica struct {
	// +optional
	KmsKeyArn  *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
	RegionName *string `json:"regionName" tf:"region_name"`
}

type TableSpecServerSideEncryption struct {
	Enabled *bool `json:"enabled" tf:"enabled"`
	// +optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
}

type TableSpecTtl struct {
	AttributeName *string `json:"attributeName" tf:"attribute_name"`
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	KmsKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn"`
}

type TableSpec struct {
	State *TableSpecResource `json:"state,omitempty" tf:"-"`

	Resource TableSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TableSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn       *string              `json:"arn,omitempty" tf:"arn"`
	Attribute []TableSpecAttribute `json:"attribute" tf:"attribute"`
	// +optional
	BillingMode *string `json:"billingMode,omitempty" tf:"billing_mode"`
	// +optional
	GlobalSecondaryIndex []TableSpecGlobalSecondaryIndex `json:"globalSecondaryIndex,omitempty" tf:"global_secondary_index"`
	HashKey              *string                         `json:"hashKey" tf:"hash_key"`
	// +optional
	LocalSecondaryIndex []TableSpecLocalSecondaryIndex `json:"localSecondaryIndex,omitempty" tf:"local_secondary_index"`
	Name                *string                        `json:"name" tf:"name"`
	// +optional
	PointInTimeRecovery *TableSpecPointInTimeRecovery `json:"pointInTimeRecovery,omitempty" tf:"point_in_time_recovery"`
	// +optional
	RangeKey *string `json:"rangeKey,omitempty" tf:"range_key"`
	// +optional
	ReadCapacity *int64 `json:"readCapacity,omitempty" tf:"read_capacity"`
	// +optional
	Replica []TableSpecReplica `json:"replica,omitempty" tf:"replica"`
	// +optional
	ServerSideEncryption *TableSpecServerSideEncryption `json:"serverSideEncryption,omitempty" tf:"server_side_encryption"`
	// +optional
	StreamArn *string `json:"streamArn,omitempty" tf:"stream_arn"`
	// +optional
	StreamEnabled *bool `json:"streamEnabled,omitempty" tf:"stream_enabled"`
	// +optional
	StreamLabel *string `json:"streamLabel,omitempty" tf:"stream_label"`
	// +optional
	StreamViewType *string `json:"streamViewType,omitempty" tf:"stream_view_type"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Ttl *TableSpecTtl `json:"ttl,omitempty" tf:"ttl"`
	// +optional
	WriteCapacity *int64 `json:"writeCapacity,omitempty" tf:"write_capacity"`
}

type TableStatus struct {
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

// TableList is a list of Tables
type TableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Table CRD objects
	Items []Table `json:"items,omitempty"`
}
