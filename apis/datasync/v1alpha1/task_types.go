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

type Task struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TaskSpec   `json:"spec,omitempty"`
	Status            TaskStatus `json:"status,omitempty"`
}

type TaskSpecExcludes struct {
	// +optional
	FilterType *string `json:"filterType,omitempty" tf:"filter_type"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type TaskSpecOptions struct {
	// +optional
	Atime *string `json:"atime,omitempty" tf:"atime"`
	// +optional
	BytesPerSecond *int64 `json:"bytesPerSecond,omitempty" tf:"bytes_per_second"`
	// +optional
	Gid *string `json:"gid,omitempty" tf:"gid"`
	// +optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`
	// +optional
	Mtime *string `json:"mtime,omitempty" tf:"mtime"`
	// +optional
	OverwriteMode *string `json:"overwriteMode,omitempty" tf:"overwrite_mode"`
	// +optional
	PosixPermissions *string `json:"posixPermissions,omitempty" tf:"posix_permissions"`
	// +optional
	PreserveDeletedFiles *string `json:"preserveDeletedFiles,omitempty" tf:"preserve_deleted_files"`
	// +optional
	PreserveDevices *string `json:"preserveDevices,omitempty" tf:"preserve_devices"`
	// +optional
	TaskQueueing *string `json:"taskQueueing,omitempty" tf:"task_queueing"`
	// +optional
	TransferMode *string `json:"transferMode,omitempty" tf:"transfer_mode"`
	// +optional
	Uid *string `json:"uid,omitempty" tf:"uid"`
	// +optional
	VerifyMode *string `json:"verifyMode,omitempty" tf:"verify_mode"`
}

type TaskSpecSchedule struct {
	ScheduleExpression *string `json:"scheduleExpression" tf:"schedule_expression"`
}

type TaskSpec struct {
	State *TaskSpecResource `json:"state,omitempty" tf:"-"`

	Resource TaskSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type TaskSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CloudwatchLogGroupArn  *string `json:"cloudwatchLogGroupArn,omitempty" tf:"cloudwatch_log_group_arn"`
	DestinationLocationArn *string `json:"destinationLocationArn" tf:"destination_location_arn"`
	// +optional
	Excludes *TaskSpecExcludes `json:"excludes,omitempty" tf:"excludes"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	Options *TaskSpecOptions `json:"options,omitempty" tf:"options"`
	// +optional
	Schedule          *TaskSpecSchedule `json:"schedule,omitempty" tf:"schedule"`
	SourceLocationArn *string           `json:"sourceLocationArn" tf:"source_location_arn"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type TaskStatus struct {
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

// TaskList is a list of Tasks
type TaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Task CRD objects
	Items []Task `json:"items,omitempty"`
}
