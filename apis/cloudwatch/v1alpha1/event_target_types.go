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

type EventTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventTargetSpec   `json:"spec,omitempty"`
	Status            EventTargetStatus `json:"status,omitempty"`
}

type EventTargetSpecBatchTarget struct {
	// +optional
	ArraySize *int64 `json:"arraySize,omitempty" tf:"array_size"`
	// +optional
	JobAttempts   *int64  `json:"jobAttempts,omitempty" tf:"job_attempts"`
	JobDefinition *string `json:"jobDefinition" tf:"job_definition"`
	JobName       *string `json:"jobName" tf:"job_name"`
}

type EventTargetSpecDeadLetterConfig struct {
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
}

type EventTargetSpecEcsTargetNetworkConfiguration struct {
	// +optional
	AssignPublicIP *bool `json:"assignPublicIP,omitempty" tf:"assign_public_ip"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`
	Subnets        []string `json:"subnets" tf:"subnets"`
}

type EventTargetSpecEcsTargetPlacementConstraint struct {
	// +optional
	Expression *string `json:"expression,omitempty" tf:"expression"`
	Type       *string `json:"type" tf:"type"`
}

type EventTargetSpecEcsTarget struct {
	// +optional
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags"`
	// +optional
	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command"`
	// +optional
	Group *string `json:"group,omitempty" tf:"group"`
	// +optional
	LaunchType *string `json:"launchType,omitempty" tf:"launch_type"`
	// +optional
	NetworkConfiguration *EventTargetSpecEcsTargetNetworkConfiguration `json:"networkConfiguration,omitempty" tf:"network_configuration"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	PlacementConstraint []EventTargetSpecEcsTargetPlacementConstraint `json:"placementConstraint,omitempty" tf:"placement_constraint"`
	// +optional
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version"`
	// +optional
	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TaskCount         *int64  `json:"taskCount,omitempty" tf:"task_count"`
	TaskDefinitionArn *string `json:"taskDefinitionArn" tf:"task_definition_arn"`
}

type EventTargetSpecHttpTarget struct {
	// +optional
	HeaderParameters *map[string]string `json:"headerParameters,omitempty" tf:"header_parameters"`
	// +optional
	PathParameterValues []string `json:"pathParameterValues,omitempty" tf:"path_parameter_values"`
	// +optional
	QueryStringParameters *map[string]string `json:"queryStringParameters,omitempty" tf:"query_string_parameters"`
}

type EventTargetSpecInputTransformer struct {
	// +optional
	InputPaths    *map[string]string `json:"inputPaths,omitempty" tf:"input_paths"`
	InputTemplate *string            `json:"inputTemplate" tf:"input_template"`
}

type EventTargetSpecKinesisTarget struct {
	// +optional
	PartitionKeyPath *string `json:"partitionKeyPath,omitempty" tf:"partition_key_path"`
}

type EventTargetSpecRedshiftTarget struct {
	Database *string `json:"database" tf:"database"`
	// +optional
	DbUser *string `json:"dbUser,omitempty" tf:"db_user"`
	// +optional
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn"`
	// +optional
	Sql *string `json:"sql,omitempty" tf:"sql"`
	// +optional
	StatementName *string `json:"statementName,omitempty" tf:"statement_name"`
	// +optional
	WithEvent *bool `json:"withEvent,omitempty" tf:"with_event"`
}

type EventTargetSpecRetryPolicy struct {
	// +optional
	MaximumEventAgeInSeconds *int64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds"`
	// +optional
	MaximumRetryAttempts *int64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts"`
}

type EventTargetSpecRunCommandTargets struct {
	Key *string `json:"key" tf:"key"`
	// +kubebuilder:validation:MaxItems=50
	Values []string `json:"values" tf:"values"`
}

type EventTargetSpecSqsTarget struct {
	// +optional
	MessageGroupID *string `json:"messageGroupID,omitempty" tf:"message_group_id"`
}

type EventTargetSpec struct {
	State *EventTargetSpecResource `json:"state,omitempty" tf:"-"`

	Resource EventTargetSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type EventTargetSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Arn *string `json:"arn" tf:"arn"`
	// +optional
	BatchTarget *EventTargetSpecBatchTarget `json:"batchTarget,omitempty" tf:"batch_target"`
	// +optional
	DeadLetterConfig *EventTargetSpecDeadLetterConfig `json:"deadLetterConfig,omitempty" tf:"dead_letter_config"`
	// +optional
	EcsTarget *EventTargetSpecEcsTarget `json:"ecsTarget,omitempty" tf:"ecs_target"`
	// +optional
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name"`
	// +optional
	HttpTarget *EventTargetSpecHttpTarget `json:"httpTarget,omitempty" tf:"http_target"`
	// +optional
	Input *string `json:"input,omitempty" tf:"input"`
	// +optional
	InputPath *string `json:"inputPath,omitempty" tf:"input_path"`
	// +optional
	InputTransformer *EventTargetSpecInputTransformer `json:"inputTransformer,omitempty" tf:"input_transformer"`
	// +optional
	KinesisTarget *EventTargetSpecKinesisTarget `json:"kinesisTarget,omitempty" tf:"kinesis_target"`
	// +optional
	RedshiftTarget *EventTargetSpecRedshiftTarget `json:"redshiftTarget,omitempty" tf:"redshift_target"`
	// +optional
	RetryPolicy *EventTargetSpecRetryPolicy `json:"retryPolicy,omitempty" tf:"retry_policy"`
	// +optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn"`
	Rule    *string `json:"rule" tf:"rule"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	RunCommandTargets []EventTargetSpecRunCommandTargets `json:"runCommandTargets,omitempty" tf:"run_command_targets"`
	// +optional
	SqsTarget *EventTargetSpecSqsTarget `json:"sqsTarget,omitempty" tf:"sqs_target"`
	// +optional
	TargetID *string `json:"targetID,omitempty" tf:"target_id"`
}

type EventTargetStatus struct {
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

// EventTargetList is a list of EventTargets
type EventTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventTarget CRD objects
	Items []EventTarget `json:"items,omitempty"`
}
