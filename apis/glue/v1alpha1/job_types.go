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

type Job struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              JobSpec   `json:"spec,omitempty"`
	Status            JobStatus `json:"status,omitempty"`
}

type JobSpecCommand struct {
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	PythonVersion  *string `json:"pythonVersion,omitempty" tf:"python_version"`
	ScriptLocation *string `json:"scriptLocation" tf:"script_location"`
}

type JobSpecExecutionProperty struct {
	// +optional
	MaxConcurrentRuns *int64 `json:"maxConcurrentRuns,omitempty" tf:"max_concurrent_runs"`
}

type JobSpecNotificationProperty struct {
	// +optional
	NotifyDelayAfter *int64 `json:"notifyDelayAfter,omitempty" tf:"notify_delay_after"`
}

type JobSpec struct {
	KubeformOutput *JobSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource JobSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type JobSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn     *string         `json:"arn,omitempty" tf:"arn"`
	Command *JobSpecCommand `json:"command" tf:"command"`
	// +optional
	Connections []string `json:"connections,omitempty" tf:"connections"`
	// +optional
	DefaultArguments *map[string]string `json:"defaultArguments,omitempty" tf:"default_arguments"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	ExecutionProperty *JobSpecExecutionProperty `json:"executionProperty,omitempty" tf:"execution_property"`
	// +optional
	GlueVersion *string `json:"glueVersion,omitempty" tf:"glue_version"`
	// +optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity"`
	// +optional
	MaxRetries *int64  `json:"maxRetries,omitempty" tf:"max_retries"`
	Name       *string `json:"name" tf:"name"`
	// +optional
	NonOverridableArguments *map[string]string `json:"nonOverridableArguments,omitempty" tf:"non_overridable_arguments"`
	// +optional
	NotificationProperty *JobSpecNotificationProperty `json:"notificationProperty,omitempty" tf:"notification_property"`
	// +optional
	NumberOfWorkers *int64  `json:"numberOfWorkers,omitempty" tf:"number_of_workers"`
	RoleArn         *string `json:"roleArn" tf:"role_arn"`
	// +optional
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// +optional
	WorkerType *string `json:"workerType,omitempty" tf:"worker_type"`
}

type JobStatus struct {
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

// JobList is a list of Jobs
type JobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Job CRD objects
	Items []Job `json:"items,omitempty"`
}
