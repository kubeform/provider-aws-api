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

type ClassificationJob struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClassificationJobSpec   `json:"spec,omitempty"`
	Status            ClassificationJobStatus `json:"status,omitempty"`
}

type ClassificationJobSpecS3JobDefinitionBucketDefinitions struct {
	AccountID *string  `json:"accountID" tf:"account_id"`
	Buckets   []string `json:"buckets" tf:"buckets"`
}

type ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm struct {
	// +optional
	Comparator *string `json:"comparator,omitempty" tf:"comparator"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTermTagValues struct {
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm struct {
	// +optional
	Comparator *string `json:"comparator,omitempty" tf:"comparator"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	TagValues []ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTermTagValues `json:"tagValues,omitempty" tf:"tag_values"`
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
}

type ClassificationJobSpecS3JobDefinitionScopingExcludesAnd struct {
	// +optional
	SimpleScopeTerm *ClassificationJobSpecS3JobDefinitionScopingExcludesAndSimpleScopeTerm `json:"simpleScopeTerm,omitempty" tf:"simple_scope_term"`
	// +optional
	TagScopeTerm *ClassificationJobSpecS3JobDefinitionScopingExcludesAndTagScopeTerm `json:"tagScopeTerm,omitempty" tf:"tag_scope_term"`
}

type ClassificationJobSpecS3JobDefinitionScopingExcludes struct {
	// +optional
	And []ClassificationJobSpecS3JobDefinitionScopingExcludesAnd `json:"and,omitempty" tf:"and"`
}

type ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm struct {
	// +optional
	Comparator *string `json:"comparator,omitempty" tf:"comparator"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Values []string `json:"values,omitempty" tf:"values"`
}

type ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTermTagValues struct {
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm struct {
	// +optional
	Comparator *string `json:"comparator,omitempty" tf:"comparator"`
	// +optional
	Key *string `json:"key,omitempty" tf:"key"`
	// +optional
	TagValues []ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTermTagValues `json:"tagValues,omitempty" tf:"tag_values"`
	// +optional
	Target *string `json:"target,omitempty" tf:"target"`
}

type ClassificationJobSpecS3JobDefinitionScopingIncludesAnd struct {
	// +optional
	SimpleScopeTerm *ClassificationJobSpecS3JobDefinitionScopingIncludesAndSimpleScopeTerm `json:"simpleScopeTerm,omitempty" tf:"simple_scope_term"`
	// +optional
	TagScopeTerm *ClassificationJobSpecS3JobDefinitionScopingIncludesAndTagScopeTerm `json:"tagScopeTerm,omitempty" tf:"tag_scope_term"`
}

type ClassificationJobSpecS3JobDefinitionScopingIncludes struct {
	// +optional
	And []ClassificationJobSpecS3JobDefinitionScopingIncludesAnd `json:"and,omitempty" tf:"and"`
}

type ClassificationJobSpecS3JobDefinitionScoping struct {
	// +optional
	Excludes *ClassificationJobSpecS3JobDefinitionScopingExcludes `json:"excludes,omitempty" tf:"excludes"`
	// +optional
	Includes *ClassificationJobSpecS3JobDefinitionScopingIncludes `json:"includes,omitempty" tf:"includes"`
}

type ClassificationJobSpecS3JobDefinition struct {
	// +optional
	BucketDefinitions []ClassificationJobSpecS3JobDefinitionBucketDefinitions `json:"bucketDefinitions,omitempty" tf:"bucket_definitions"`
	// +optional
	Scoping *ClassificationJobSpecS3JobDefinitionScoping `json:"scoping,omitempty" tf:"scoping"`
}

type ClassificationJobSpecScheduleFrequency struct {
	// +optional
	DailySchedule *bool `json:"dailySchedule,omitempty" tf:"daily_schedule"`
	// +optional
	MonthlySchedule *int64 `json:"monthlySchedule,omitempty" tf:"monthly_schedule"`
	// +optional
	WeeklySchedule *string `json:"weeklySchedule,omitempty" tf:"weekly_schedule"`
}

type ClassificationJobSpecUserPausedDetails struct {
	// +optional
	JobExpiresAt *string `json:"jobExpiresAt,omitempty" tf:"job_expires_at"`
	// +optional
	JobImminentExpirationHealthEventArn *string `json:"jobImminentExpirationHealthEventArn,omitempty" tf:"job_imminent_expiration_health_event_arn"`
	// +optional
	JobPausedAt *string `json:"jobPausedAt,omitempty" tf:"job_paused_at"`
}

type ClassificationJobSpec struct {
	State *ClassificationJobSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClassificationJobSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ClassificationJobSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at"`
	// +optional
	CustomDataIdentifierIDS []string `json:"customDataIdentifierIDS,omitempty" tf:"custom_data_identifier_ids"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	InitialRun *bool `json:"initialRun,omitempty" tf:"initial_run"`
	// +optional
	JobArn *string `json:"jobArn,omitempty" tf:"job_arn"`
	// +optional
	JobID *string `json:"jobID,omitempty" tf:"job_id"`
	// +optional
	JobStatus *string `json:"jobStatus,omitempty" tf:"job_status"`
	JobType   *string `json:"jobType" tf:"job_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix      *string                               `json:"namePrefix,omitempty" tf:"name_prefix"`
	S3JobDefinition *ClassificationJobSpecS3JobDefinition `json:"s3JobDefinition" tf:"s3_job_definition"`
	// +optional
	SamplingPercentage *int64 `json:"samplingPercentage,omitempty" tf:"sampling_percentage"`
	// +optional
	ScheduleFrequency *ClassificationJobSpecScheduleFrequency `json:"scheduleFrequency,omitempty" tf:"schedule_frequency"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	UserPausedDetails []ClassificationJobSpecUserPausedDetails `json:"userPausedDetails,omitempty" tf:"user_paused_details"`
}

type ClassificationJobStatus struct {
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

// ClassificationJobList is a list of ClassificationJobs
type ClassificationJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClassificationJob CRD objects
	Items []ClassificationJob `json:"items,omitempty"`
}
