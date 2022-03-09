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

type ApiCache struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiCacheSpec   `json:"spec,omitempty"`
	Status            ApiCacheStatus `json:"status,omitempty"`
}

type ApiCacheSpec struct {
	State *ApiCacheSpecResource `json:"state,omitempty" tf:"-"`

	Resource ApiCacheSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ApiCacheSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiCachingBehavior *string `json:"apiCachingBehavior" tf:"api_caching_behavior"`
	ApiID              *string `json:"apiID" tf:"api_id"`
	// +optional
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled"`
	// +optional
	TransitEncryptionEnabled *bool   `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled"`
	Ttl                      *int64  `json:"ttl" tf:"ttl"`
	Type                     *string `json:"type" tf:"type"`
}

type ApiCacheStatus struct {
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

// ApiCacheList is a list of ApiCaches
type ApiCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiCache CRD objects
	Items []ApiCache `json:"items,omitempty"`
}
