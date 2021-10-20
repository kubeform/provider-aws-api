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

type ClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            ClusterSnapshotStatus `json:"status,omitempty"`
}

type ClusterSnapshotSpec struct {
	State *ClusterSnapshotSpecResource `json:"state,omitempty" tf:"-"`

	Resource ClusterSnapshotSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ClusterSnapshotSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AvailabilityZones   []string `json:"availabilityZones,omitempty" tf:"availability_zones"`
	DbClusterIdentifier *string  `json:"dbClusterIdentifier" tf:"db_cluster_identifier"`
	// +optional
	DbClusterSnapshotArn        *string `json:"dbClusterSnapshotArn,omitempty" tf:"db_cluster_snapshot_arn"`
	DbClusterSnapshotIdentifier *string `json:"dbClusterSnapshotIdentifier" tf:"db_cluster_snapshot_identifier"`
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// +optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	SnapshotType *string `json:"snapshotType,omitempty" tf:"snapshot_type"`
	// +optional
	SourceDbClusterSnapshotArn *string `json:"sourceDbClusterSnapshotArn,omitempty" tf:"source_db_cluster_snapshot_arn"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type ClusterSnapshotStatus struct {
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

// ClusterSnapshotList is a list of ClusterSnapshots
type ClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterSnapshot CRD objects
	Items []ClusterSnapshot `json:"items,omitempty"`
}
