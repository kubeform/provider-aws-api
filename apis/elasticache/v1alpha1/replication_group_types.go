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

type ReplicationGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReplicationGroupSpec   `json:"spec,omitempty"`
	Status            ReplicationGroupStatus `json:"status,omitempty"`
}

type ReplicationGroupSpecClusterMode struct {
	NumNodeGroups        *int64 `json:"numNodeGroups" tf:"num_node_groups"`
	ReplicasPerNodeGroup *int64 `json:"replicasPerNodeGroup" tf:"replicas_per_node_group"`
}

type ReplicationGroupSpec struct {
	State *ReplicationGroupSpecResource `json:"state,omitempty" tf:"-"`

	Resource ReplicationGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ReplicationGroupSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApplyImmediately *bool `json:"applyImmediately,omitempty" tf:"apply_immediately"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AtRestEncryptionEnabled *bool `json:"atRestEncryptionEnabled,omitempty" tf:"at_rest_encryption_enabled"`
	// +optional
	AuthToken *string `json:"-" sensitive:"true" tf:"auth_token"`
	// +optional
	AutoMinorVersionUpgrade *bool `json:"autoMinorVersionUpgrade,omitempty" tf:"auto_minor_version_upgrade"`
	// +optional
	AutomaticFailoverEnabled *bool `json:"automaticFailoverEnabled,omitempty" tf:"automatic_failover_enabled"`
	// +optional
	AvailabilityZones []string `json:"availabilityZones,omitempty" tf:"availability_zones"`
	// +optional
	ClusterEnabled *bool `json:"clusterEnabled,omitempty" tf:"cluster_enabled"`
	// +optional
	ClusterMode *ReplicationGroupSpecClusterMode `json:"clusterMode,omitempty" tf:"cluster_mode"`
	// +optional
	ConfigurationEndpointAddress *string `json:"configurationEndpointAddress,omitempty" tf:"configuration_endpoint_address"`
	// +optional
	Engine *string `json:"engine,omitempty" tf:"engine"`
	// +optional
	EngineVersion *string `json:"engineVersion,omitempty" tf:"engine_version"`
	// +optional
	EngineVersionActual *string `json:"engineVersionActual,omitempty" tf:"engine_version_actual"`
	// +optional
	FinalSnapshotIdentifier *string `json:"finalSnapshotIdentifier,omitempty" tf:"final_snapshot_identifier"`
	// +optional
	GlobalReplicationGroupID *string `json:"globalReplicationGroupID,omitempty" tf:"global_replication_group_id"`
	// +optional
	KmsKeyID *string `json:"kmsKeyID,omitempty" tf:"kms_key_id"`
	// +optional
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty" tf:"maintenance_window"`
	// +optional
	MemberClusters []string `json:"memberClusters,omitempty" tf:"member_clusters"`
	// +optional
	MultiAzEnabled *bool `json:"multiAzEnabled,omitempty" tf:"multi_az_enabled"`
	// +optional
	NodeType *string `json:"nodeType,omitempty" tf:"node_type"`
	// +optional
	NotificationTopicArn *string `json:"notificationTopicArn,omitempty" tf:"notification_topic_arn"`
	// +optional
	NumberCacheClusters *int64 `json:"numberCacheClusters,omitempty" tf:"number_cache_clusters"`
	// +optional
	ParameterGroupName *string `json:"parameterGroupName,omitempty" tf:"parameter_group_name"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PrimaryEndpointAddress *string `json:"primaryEndpointAddress,omitempty" tf:"primary_endpoint_address"`
	// +optional
	ReaderEndpointAddress       *string `json:"readerEndpointAddress,omitempty" tf:"reader_endpoint_address"`
	ReplicationGroupDescription *string `json:"replicationGroupDescription" tf:"replication_group_description"`
	ReplicationGroupID          *string `json:"replicationGroupID" tf:"replication_group_id"`
	// +optional
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids"`
	// +optional
	SecurityGroupNames []string `json:"securityGroupNames,omitempty" tf:"security_group_names"`
	// +optional
	SnapshotArns []string `json:"snapshotArns,omitempty" tf:"snapshot_arns"`
	// +optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name"`
	// +optional
	SnapshotRetentionLimit *int64 `json:"snapshotRetentionLimit,omitempty" tf:"snapshot_retention_limit"`
	// +optional
	SnapshotWindow *string `json:"snapshotWindow,omitempty" tf:"snapshot_window"`
	// +optional
	SubnetGroupName *string `json:"subnetGroupName,omitempty" tf:"subnet_group_name"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	TransitEncryptionEnabled *bool `json:"transitEncryptionEnabled,omitempty" tf:"transit_encryption_enabled"`
}

type ReplicationGroupStatus struct {
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

// ReplicationGroupList is a list of ReplicationGroups
type ReplicationGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ReplicationGroup CRD objects
	Items []ReplicationGroup `json:"items,omitempty"`
}
