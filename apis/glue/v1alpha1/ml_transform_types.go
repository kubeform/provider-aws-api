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

type MlTransform struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MlTransformSpec   `json:"spec,omitempty"`
	Status            MlTransformStatus `json:"status,omitempty"`
}

type MlTransformSpecInputRecordTables struct {
	// +optional
	CatalogID *string `json:"catalogID,omitempty" tf:"catalog_id"`
	// +optional
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name"`
	DatabaseName   *string `json:"databaseName" tf:"database_name"`
	TableName      *string `json:"tableName" tf:"table_name"`
}

type MlTransformSpecParametersFindMatchesParameters struct {
	// +optional
	AccuracyCostTradeOff *float64 `json:"accuracyCostTradeOff,omitempty" tf:"accuracy_cost_trade_off"`
	// +optional
	EnforceProvidedLabels *bool `json:"enforceProvidedLabels,omitempty" tf:"enforce_provided_labels"`
	// +optional
	PrecisionRecallTradeOff *float64 `json:"precisionRecallTradeOff,omitempty" tf:"precision_recall_trade_off"`
	// +optional
	PrimaryKeyColumnName *string `json:"primaryKeyColumnName,omitempty" tf:"primary_key_column_name"`
}

type MlTransformSpecParameters struct {
	FindMatchesParameters *MlTransformSpecParametersFindMatchesParameters `json:"findMatchesParameters" tf:"find_matches_parameters"`
	TransformType         *string                                         `json:"transformType" tf:"transform_type"`
}

type MlTransformSpecSchema struct {
	// +optional
	DataType *string `json:"dataType,omitempty" tf:"data_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type MlTransformSpec struct {
	State *MlTransformSpecResource `json:"state,omitempty" tf:"-"`

	Resource MlTransformSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type MlTransformSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	GlueVersion       *string                            `json:"glueVersion,omitempty" tf:"glue_version"`
	InputRecordTables []MlTransformSpecInputRecordTables `json:"inputRecordTables" tf:"input_record_tables"`
	// +optional
	LabelCount *int64 `json:"labelCount,omitempty" tf:"label_count"`
	// +optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity"`
	// +optional
	MaxRetries *int64  `json:"maxRetries,omitempty" tf:"max_retries"`
	Name       *string `json:"name" tf:"name"`
	// +optional
	NumberOfWorkers *int64                     `json:"numberOfWorkers,omitempty" tf:"number_of_workers"`
	Parameters      *MlTransformSpecParameters `json:"parameters" tf:"parameters"`
	RoleArn         *string                    `json:"roleArn" tf:"role_arn"`
	// +optional
	Schema []MlTransformSpecSchema `json:"schema,omitempty" tf:"schema"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// +optional
	WorkerType *string `json:"workerType,omitempty" tf:"worker_type"`
}

type MlTransformStatus struct {
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

// MlTransformList is a list of MlTransforms
type MlTransformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MlTransform CRD objects
	Items []MlTransform `json:"items,omitempty"`
}
