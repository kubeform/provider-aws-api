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

type ClusterInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterInstanceSpec   `json:"spec,omitempty"`
	Status            ClusterInstanceStatus `json:"status,omitempty"`
}

type ClusterInstanceSpec struct {
	KubeformOutput *ClusterInstanceSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ClusterInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ClusterInstanceSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade"`
	// +optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`
	// +optional
	CaCertIdentifier  *string `json:"caCertIdentifier,omitempty" tf:"ca_cert_identifier"`
	ClusterIdentifier *string `json:"clusterIdentifier" tf:"cluster_identifier"`
	// +optional
	CopyTagsToSnapshot *bool `json:"copyTagsToSnapshot,omitempty" tf:"copy_tags_to_snapshot"`
	// +optional
	DbParameterGroupName *string `json:"dbParameterGroupName,omitempty" tf:"db_parameter_group_name"`
	// +optional
	DbSubnetGroupName *string `json:"dbSubnetGroupName,omitempty" tf:"db_subnet_group_name"`
	// +optional
	DbiResourceID *string `json:"dbiResourceID,omitempty" tf:"dbi_resource_id"`
	// +optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint"`
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// +optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`
	// +optional
	Identifier *string `json:"identifier,omitempty" tf:"identifier"`
	// +optional
	IdentifierPrefix *string `json:"identifierPrefix,omitempty" tf:"identifier_prefix"`
	InstanceClass    *string `json:"instanceClass" tf:"instance_class"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	MonitoringInterval *int64 `json:"monitoringInterval,omitempty" tf:"monitoring_interval"`
	// +optional
	MonitoringRoleArn *string `json:"monitoringRoleArn,omitempty" tf:"monitoring_role_arn"`
	// +optional
	PerformanceInsightsEnabled *bool `json:"performanceInsightsEnabled,omitempty" tf:"performance_insights_enabled"`
	// +optional
	PerformanceInsightsKmsKeyID *string `json:"performanceInsightsKmsKeyID,omitempty" tf:"performance_insights_kms_key_id"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PreferredBackupWindow *string `json:"preferredBackupWindow,omitempty" tf:"preferred_backup_window"`
	// +optional
	PreferredMaintenanceWindow *string `json:"preferredMaintenanceWindow,omitempty" tf:"preferred_maintenance_window"`
	// +optional
	PromotionTier *int64 `json:"promotionTier,omitempty" tf:"promotion_tier"`
	// +optional
	PubliclyAccessible *bool `json:"publiclyAccessible,omitempty" tf:"publicly_accessible"`
	// +optional
	StorageEncrypted *bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	Writer *bool `json:"writer,omitempty" tf:"writer"`
}

type ClusterInstanceStatus struct {
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

// ClusterInstanceList is a list of ClusterInstances
type ClusterInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ClusterInstance CRD objects
	Items []ClusterInstance `json:"items,omitempty"`
}
