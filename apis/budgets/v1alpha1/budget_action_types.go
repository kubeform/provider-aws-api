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

type BudgetAction struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetActionSpec   `json:"spec,omitempty"`
	Status            BudgetActionStatus `json:"status,omitempty"`
}

type BudgetActionSpecActionThreshold struct {
	ActionThresholdType  *string  `json:"actionThresholdType" tf:"action_threshold_type"`
	ActionThresholdValue *float64 `json:"actionThresholdValue" tf:"action_threshold_value"`
}

type BudgetActionSpecDefinitionIamActionDefinition struct {
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Groups    []string `json:"groups,omitempty" tf:"groups"`
	PolicyArn *string  `json:"policyArn" tf:"policy_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Roles []string `json:"roles,omitempty" tf:"roles"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Users []string `json:"users,omitempty" tf:"users"`
}

type BudgetActionSpecDefinitionScpActionDefinition struct {
	PolicyID *string `json:"policyID" tf:"policy_id"`
	// +kubebuilder:validation:MaxItems=100
	TargetIDS []string `json:"targetIDS" tf:"target_ids"`
}

type BudgetActionSpecDefinitionSsmActionDefinition struct {
	ActionSubType *string `json:"actionSubType" tf:"action_sub_type"`
	// +kubebuilder:validation:MaxItems=100
	InstanceIDS []string `json:"instanceIDS" tf:"instance_ids"`
	Region      *string  `json:"region" tf:"region"`
}

type BudgetActionSpecDefinition struct {
	// +optional
	IamActionDefinition *BudgetActionSpecDefinitionIamActionDefinition `json:"iamActionDefinition,omitempty" tf:"iam_action_definition"`
	// +optional
	ScpActionDefinition *BudgetActionSpecDefinitionScpActionDefinition `json:"scpActionDefinition,omitempty" tf:"scp_action_definition"`
	// +optional
	SsmActionDefinition *BudgetActionSpecDefinitionSsmActionDefinition `json:"ssmActionDefinition,omitempty" tf:"ssm_action_definition"`
}

type BudgetActionSpecSubscriber struct {
	Address          *string `json:"address" tf:"address"`
	SubscriptionType *string `json:"subscriptionType" tf:"subscription_type"`
}

type BudgetActionSpec struct {
	KubeformOutput *BudgetActionSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource BudgetActionSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type BudgetActionSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// +optional
	ActionID        *string                          `json:"actionID,omitempty" tf:"action_id"`
	ActionThreshold *BudgetActionSpecActionThreshold `json:"actionThreshold" tf:"action_threshold"`
	ActionType      *string                          `json:"actionType" tf:"action_type"`
	ApprovalModel   *string                          `json:"approvalModel" tf:"approval_model"`
	// +optional
	Arn              *string                     `json:"arn,omitempty" tf:"arn"`
	BudgetName       *string                     `json:"budgetName" tf:"budget_name"`
	Definition       *BudgetActionSpecDefinition `json:"definition" tf:"definition"`
	ExecutionRoleArn *string                     `json:"executionRoleArn" tf:"execution_role_arn"`
	NotificationType *string                     `json:"notificationType" tf:"notification_type"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +kubebuilder:validation:MaxItems=11
	Subscriber []BudgetActionSpecSubscriber `json:"subscriber" tf:"subscriber"`
}

type BudgetActionStatus struct {
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

// BudgetActionList is a list of BudgetActions
type BudgetActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BudgetAction CRD objects
	Items []BudgetAction `json:"items,omitempty"`
}
