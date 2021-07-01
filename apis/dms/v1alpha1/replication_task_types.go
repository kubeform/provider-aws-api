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

type ReplicationTask struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationTaskSpec   `json:"spec,omitempty"`
	Status            ReplicationTaskStatus `json:"status,omitempty"`
}

type ReplicationTaskSpec struct {
	KubeformOutput *ReplicationTaskSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ReplicationTaskSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ReplicationTaskSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CdcStartTime           *string `json:"cdcStartTime,omitempty" tf:"cdc_start_time"`
	MigrationType          *string `json:"migrationType" tf:"migration_type"`
	ReplicationInstanceArn *string `json:"replicationInstanceArn" tf:"replication_instance_arn"`
	// +optional
	ReplicationTaskArn *string `json:"replicationTaskArn,omitempty" tf:"replication_task_arn"`
	ReplicationTaskID  *string `json:"replicationTaskID" tf:"replication_task_id"`
	// +optional
	ReplicationTaskSettings *string `json:"replicationTaskSettings,omitempty" tf:"replication_task_settings"`
	SourceEndpointArn       *string `json:"sourceEndpointArn" tf:"source_endpoint_arn"`
	TableMappings           *string `json:"tableMappings" tf:"table_mappings"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll           *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	TargetEndpointArn *string            `json:"targetEndpointArn" tf:"target_endpoint_arn"`
}

type ReplicationTaskStatus struct {
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

// ReplicationTaskList is a list of ReplicationTasks
type ReplicationTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ReplicationTask CRD objects
	Items []ReplicationTask `json:"items,omitempty"`
}
