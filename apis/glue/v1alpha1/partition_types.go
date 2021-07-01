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

type Partition struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PartitionSpec   `json:"spec,omitempty"`
	Status            PartitionStatus `json:"status,omitempty"`
}

type PartitionSpecStorageDescriptorColumns struct {
	// +optional
	Comment *string `json:"comment,omitempty" tf:"comment"`
	Name    *string `json:"name" tf:"name"`
	// +optional
	Type *string `json:"type,omitempty" tf:"type"`
}

type PartitionSpecStorageDescriptorSerDeInfo struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	SerializationLibrary *string `json:"serializationLibrary,omitempty" tf:"serialization_library"`
}

type PartitionSpecStorageDescriptorSkewedInfo struct {
	// +optional
	SkewedColumnNames []string `json:"skewedColumnNames,omitempty" tf:"skewed_column_names"`
	// +optional
	SkewedColumnValueLocationMaps *map[string]string `json:"skewedColumnValueLocationMaps,omitempty" tf:"skewed_column_value_location_maps"`
	// +optional
	SkewedColumnValues []string `json:"skewedColumnValues,omitempty" tf:"skewed_column_values"`
}

type PartitionSpecStorageDescriptorSortColumns struct {
	Column    *string `json:"column" tf:"column"`
	SortOrder *int64  `json:"sortOrder" tf:"sort_order"`
}

type PartitionSpecStorageDescriptor struct {
	// +optional
	BucketColumns []string `json:"bucketColumns,omitempty" tf:"bucket_columns"`
	// +optional
	Columns []PartitionSpecStorageDescriptorColumns `json:"columns,omitempty" tf:"columns"`
	// +optional
	Compressed *bool `json:"compressed,omitempty" tf:"compressed"`
	// +optional
	InputFormat *string `json:"inputFormat,omitempty" tf:"input_format"`
	// +optional
	Location *string `json:"location,omitempty" tf:"location"`
	// +optional
	NumberOfBuckets *int64 `json:"numberOfBuckets,omitempty" tf:"number_of_buckets"`
	// +optional
	OutputFormat *string `json:"outputFormat,omitempty" tf:"output_format"`
	// +optional
	Parameters *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	// +optional
	SerDeInfo *PartitionSpecStorageDescriptorSerDeInfo `json:"serDeInfo,omitempty" tf:"ser_de_info"`
	// +optional
	SkewedInfo *PartitionSpecStorageDescriptorSkewedInfo `json:"skewedInfo,omitempty" tf:"skewed_info"`
	// +optional
	SortColumns []PartitionSpecStorageDescriptorSortColumns `json:"sortColumns,omitempty" tf:"sort_columns"`
	// +optional
	StoredAsSubDirectories *bool `json:"storedAsSubDirectories,omitempty" tf:"stored_as_sub_directories"`
}

type PartitionSpec struct {
	KubeformOutput *PartitionSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource PartitionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type PartitionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CatalogID *string `json:"catalogID,omitempty" tf:"catalog_id"`
	// +optional
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time"`
	DatabaseName *string `json:"databaseName" tf:"database_name"`
	// +optional
	LastAccessedTime *string `json:"lastAccessedTime,omitempty" tf:"last_accessed_time"`
	// +optional
	LastAnalyzedTime *string `json:"lastAnalyzedTime,omitempty" tf:"last_analyzed_time"`
	// +optional
	Parameters      *map[string]string `json:"parameters,omitempty" tf:"parameters"`
	PartitionValues []string           `json:"partitionValues" tf:"partition_values"`
	// +optional
	StorageDescriptor *PartitionSpecStorageDescriptor `json:"storageDescriptor,omitempty" tf:"storage_descriptor"`
	TableName         *string                         `json:"tableName" tf:"table_name"`
}

type PartitionStatus struct {
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

// PartitionList is a list of Partitions
type PartitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Partition CRD objects
	Items []Partition `json:"items,omitempty"`
}
