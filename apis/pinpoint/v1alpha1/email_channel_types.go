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

type EmailChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmailChannelSpec   `json:"spec,omitempty"`
	Status            EmailChannelStatus `json:"status,omitempty"`
}

type EmailChannelSpec struct {
	State *EmailChannelSpecResource `json:"state,omitempty" tf:"-"`

	Resource EmailChannelSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type EmailChannelSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApplicationID *string `json:"applicationID" tf:"application_id"`
	// +optional
	ConfigurationSet *string `json:"configurationSet,omitempty" tf:"configuration_set"`
	// +optional
	Enabled     *bool   `json:"enabled,omitempty" tf:"enabled"`
	FromAddress *string `json:"fromAddress" tf:"from_address"`
	Identity    *string `json:"identity" tf:"identity"`
	// +optional
	MessagesPerSecond *int64 `json:"messagesPerSecond,omitempty" tf:"messages_per_second"`
	// +optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn"`
}

type EmailChannelStatus struct {
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

// EmailChannelList is a list of EmailChannels
type EmailChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EmailChannel CRD objects
	Items []EmailChannel `json:"items,omitempty"`
}
