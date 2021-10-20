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

type Gateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec,omitempty"`
	Status            GatewayStatus `json:"status,omitempty"`
}

type GatewaySpecGatewayNetworkInterface struct {
	// +optional
	Ipv4Address *string `json:"ipv4Address,omitempty" tf:"ipv4_address"`
}

type GatewaySpecSmbActiveDirectorySettings struct {
	// +optional
	ActiveDirectoryStatus *string `json:"activeDirectoryStatus,omitempty" tf:"active_directory_status"`
	// +optional
	DomainControllers []string `json:"domainControllers,omitempty" tf:"domain_controllers"`
	DomainName        *string  `json:"domainName" tf:"domain_name"`
	// +optional
	OrganizationalUnit *string `json:"organizationalUnit,omitempty" tf:"organizational_unit"`
	Password           *string `json:"-" sensitive:"true" tf:"password"`
	// +optional
	TimeoutInSeconds *int64  `json:"timeoutInSeconds,omitempty" tf:"timeout_in_seconds"`
	Username         *string `json:"username" tf:"username"`
}

type GatewaySpec struct {
	State *GatewaySpecResource `json:"state,omitempty" tf:"-"`

	Resource GatewaySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GatewaySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ActivationKey *string `json:"activationKey,omitempty" tf:"activation_key"`
	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	AverageDownloadRateLimitInBitsPerSec *int64 `json:"averageDownloadRateLimitInBitsPerSec,omitempty" tf:"average_download_rate_limit_in_bits_per_sec"`
	// +optional
	AverageUploadRateLimitInBitsPerSec *int64 `json:"averageUploadRateLimitInBitsPerSec,omitempty" tf:"average_upload_rate_limit_in_bits_per_sec"`
	// +optional
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn,omitempty" tf:"cloudwatch_log_group_arn"`
	// +optional
	Ec2InstanceID *string `json:"ec2InstanceID,omitempty" tf:"ec2_instance_id"`
	// +optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type"`
	// +optional
	GatewayID *string `json:"gatewayID,omitempty" tf:"gateway_id"`
	// +optional
	GatewayIPAddress *string `json:"gatewayIPAddress,omitempty" tf:"gateway_ip_address"`
	GatewayName      *string `json:"gatewayName" tf:"gateway_name"`
	// +optional
	GatewayNetworkInterface []GatewaySpecGatewayNetworkInterface `json:"gatewayNetworkInterface,omitempty" tf:"gateway_network_interface"`
	GatewayTimezone         *string                              `json:"gatewayTimezone" tf:"gateway_timezone"`
	// +optional
	GatewayType *string `json:"gatewayType,omitempty" tf:"gateway_type"`
	// +optional
	GatewayVpcEndpoint *string `json:"gatewayVpcEndpoint,omitempty" tf:"gateway_vpc_endpoint"`
	// +optional
	HostEnvironment *string `json:"hostEnvironment,omitempty" tf:"host_environment"`
	// +optional
	MediumChangerType *string `json:"mediumChangerType,omitempty" tf:"medium_changer_type"`
	// +optional
	SmbActiveDirectorySettings *GatewaySpecSmbActiveDirectorySettings `json:"smbActiveDirectorySettings,omitempty" tf:"smb_active_directory_settings"`
	// +optional
	SmbFileShareVisibility *bool `json:"smbFileShareVisibility,omitempty" tf:"smb_file_share_visibility"`
	// +optional
	SmbGuestPassword *string `json:"-" sensitive:"true" tf:"smb_guest_password"`
	// +optional
	SmbSecurityStrategy *string `json:"smbSecurityStrategy,omitempty" tf:"smb_security_strategy"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	TapeDriveType *string `json:"tapeDriveType,omitempty" tf:"tape_drive_type"`
}

type GatewayStatus struct {
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

// GatewayList is a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Gateway CRD objects
	Items []Gateway `json:"items,omitempty"`
}
