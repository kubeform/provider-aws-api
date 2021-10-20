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

type FargateProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FargateProfileSpec   `json:"spec,omitempty"`
	Status            FargateProfileStatus `json:"status,omitempty"`
}

type FargateProfileSpecSelector struct {
	// +optional
	Labels    *map[string]string `json:"labels,omitempty" tf:"labels"`
	Namespace *string            `json:"namespace" tf:"namespace"`
}

type FargateProfileSpec struct {
	State *FargateProfileSpecResource `json:"state,omitempty" tf:"-"`

	Resource FargateProfileSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type FargateProfileSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn                 *string `json:"arn,omitempty" tf:"arn"`
	ClusterName         *string `json:"clusterName" tf:"cluster_name"`
	FargateProfileName  *string `json:"fargateProfileName" tf:"fargate_profile_name"`
	PodExecutionRoleArn *string `json:"podExecutionRoleArn" tf:"pod_execution_role_arn"`
	// +kubebuilder:validation:MinItems=1
	Selector []FargateProfileSpecSelector `json:"selector" tf:"selector"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type FargateProfileStatus struct {
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

// FargateProfileList is a list of FargateProfiles
type FargateProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FargateProfile CRD objects
	Items []FargateProfile `json:"items,omitempty"`
}
