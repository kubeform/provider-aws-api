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

type Workspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorkspaceSpec   `json:"spec,omitempty"`
	Status            WorkspaceStatus `json:"status,omitempty"`
}

type WorkspaceSpecWorkspaceProperties struct {
	// +optional
	ComputeTypeName *string `json:"computeTypeName,omitempty" tf:"compute_type_name"`
	// +optional
	RootVolumeSizeGib *int64 `json:"rootVolumeSizeGib,omitempty" tf:"root_volume_size_gib"`
	// +optional
	RunningMode *string `json:"runningMode,omitempty" tf:"running_mode"`
	// +optional
	RunningModeAutoStopTimeoutInMinutes *int64 `json:"runningModeAutoStopTimeoutInMinutes,omitempty" tf:"running_mode_auto_stop_timeout_in_minutes"`
	// +optional
	UserVolumeSizeGib *int64 `json:"userVolumeSizeGib,omitempty" tf:"user_volume_size_gib"`
}

type WorkspaceSpec struct {
	State *WorkspaceSpecResource `json:"state,omitempty" tf:"-"`

	Resource WorkspaceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type WorkspaceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BundleID *string `json:"bundleID" tf:"bundle_id"`
	// +optional
	ComputerName *string `json:"computerName,omitempty" tf:"computer_name"`
	DirectoryID  *string `json:"directoryID" tf:"directory_id"`
	// +optional
	IpAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`
	// +optional
	RootVolumeEncryptionEnabled *bool `json:"rootVolumeEncryptionEnabled,omitempty" tf:"root_volume_encryption_enabled"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll  *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	UserName *string            `json:"userName" tf:"user_name"`
	// +optional
	UserVolumeEncryptionEnabled *bool `json:"userVolumeEncryptionEnabled,omitempty" tf:"user_volume_encryption_enabled"`
	// +optional
	VolumeEncryptionKey *string `json:"volumeEncryptionKey,omitempty" tf:"volume_encryption_key"`
	// +optional
	WorkspaceProperties *WorkspaceSpecWorkspaceProperties `json:"workspaceProperties,omitempty" tf:"workspace_properties"`
}

type WorkspaceStatus struct {
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

// WorkspaceList is a list of Workspaces
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Workspace CRD objects
	Items []Workspace `json:"items,omitempty"`
}
