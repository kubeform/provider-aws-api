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

type Bot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BotSpec   `json:"spec,omitempty"`
	Status            BotStatus `json:"status,omitempty"`
}

type BotSpecAbortStatementMessage struct {
	Content     *string `json:"content" tf:"content"`
	ContentType *string `json:"contentType" tf:"content_type"`
	// +optional
	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type BotSpecAbortStatement struct {
	// +kubebuilder:validation:MaxItems=15
	// +kubebuilder:validation:MinItems=1
	Message []BotSpecAbortStatementMessage `json:"message" tf:"message"`
	// +optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type BotSpecClarificationPromptMessage struct {
	Content     *string `json:"content" tf:"content"`
	ContentType *string `json:"contentType" tf:"content_type"`
	// +optional
	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type BotSpecClarificationPrompt struct {
	MaxAttempts *int64 `json:"maxAttempts" tf:"max_attempts"`
	// +kubebuilder:validation:MaxItems=15
	// +kubebuilder:validation:MinItems=1
	Message []BotSpecClarificationPromptMessage `json:"message" tf:"message"`
	// +optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type BotSpecIntent struct {
	IntentName    *string `json:"intentName" tf:"intent_name"`
	IntentVersion *string `json:"intentVersion" tf:"intent_version"`
}

type BotSpec struct {
	KubeformOutput *BotSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource BotSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BotSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AbortStatement *BotSpecAbortStatement `json:"abortStatement" tf:"abort_statement"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	Checksum      *string `json:"checksum,omitempty" tf:"checksum"`
	ChildDirected *bool   `json:"childDirected" tf:"child_directed"`
	// +optional
	ClarificationPrompt *BotSpecClarificationPrompt `json:"clarificationPrompt,omitempty" tf:"clarification_prompt"`
	// +optional
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version"`
	// +optional
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	DetectSentiment *bool `json:"detectSentiment,omitempty" tf:"detect_sentiment"`
	// +optional
	EnableModelImprovements *bool `json:"enableModelImprovements,omitempty" tf:"enable_model_improvements"`
	// +optional
	FailureReason *string `json:"failureReason,omitempty" tf:"failure_reason"`
	// +optional
	IdleSessionTtlInSeconds *int64 `json:"idleSessionTtlInSeconds,omitempty" tf:"idle_session_ttl_in_seconds"`
	// +kubebuilder:validation:MaxItems=100
	// +kubebuilder:validation:MinItems=1
	Intent []BotSpecIntent `json:"intent" tf:"intent"`
	// +optional
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date"`
	// +optional
	Locale *string `json:"locale,omitempty" tf:"locale"`
	Name   *string `json:"name" tf:"name"`
	// +optional
	NluIntentConfidenceThreshold *float64 `json:"nluIntentConfidenceThreshold,omitempty" tf:"nlu_intent_confidence_threshold"`
	// +optional
	ProcessBehavior *string `json:"processBehavior,omitempty" tf:"process_behavior"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	Version *string `json:"version,omitempty" tf:"version"`
	// +optional
	VoiceID *string `json:"voiceID,omitempty" tf:"voice_id"`
}

type BotStatus struct {
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

// BotList is a list of Bots
type BotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Bot CRD objects
	Items []Bot `json:"items,omitempty"`
}
