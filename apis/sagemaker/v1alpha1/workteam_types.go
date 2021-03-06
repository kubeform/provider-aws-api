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

type Workteam struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkteamSpec   `json:"spec,omitempty"`
	Status            WorkteamStatus `json:"status,omitempty"`
}

type WorkteamSpecMemberDefinitionCognitoMemberDefinition struct {
	ClientID  *string `json:"clientID" tf:"client_id"`
	UserGroup *string `json:"userGroup" tf:"user_group"`
	UserPool  *string `json:"userPool" tf:"user_pool"`
}

type WorkteamSpecMemberDefinitionOidcMemberDefinition struct {
	// +kubebuilder:validation:MaxItems=10
	Groups []string `json:"groups" tf:"groups"`
}

type WorkteamSpecMemberDefinition struct {
	// +optional
	CognitoMemberDefinition *WorkteamSpecMemberDefinitionCognitoMemberDefinition `json:"cognitoMemberDefinition,omitempty" tf:"cognito_member_definition"`
	// +optional
	OidcMemberDefinition *WorkteamSpecMemberDefinitionOidcMemberDefinition `json:"oidcMemberDefinition,omitempty" tf:"oidc_member_definition"`
}

type WorkteamSpecNotificationConfiguration struct {
	// +optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn"`
}

type WorkteamSpec struct {
	State *WorkteamSpecResource `json:"state,omitempty" tf:"-"`

	Resource WorkteamSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type WorkteamSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn         *string `json:"arn,omitempty" tf:"arn"`
	Description *string `json:"description" tf:"description"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	MemberDefinition []WorkteamSpecMemberDefinition `json:"memberDefinition" tf:"member_definition"`
	// +optional
	NotificationConfiguration *WorkteamSpecNotificationConfiguration `json:"notificationConfiguration,omitempty" tf:"notification_configuration"`
	// +optional
	Subdomain *string `json:"subdomain,omitempty" tf:"subdomain"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll       *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	WorkforceName *string            `json:"workforceName" tf:"workforce_name"`
	WorkteamName  *string            `json:"workteamName" tf:"workteam_name"`
}

type WorkteamStatus struct {
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

// WorkteamList is a list of Workteams
type WorkteamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Workteam CRD objects
	Items []Workteam `json:"items,omitempty"`
}
