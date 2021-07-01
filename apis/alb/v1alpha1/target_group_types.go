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

type TargetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TargetGroupSpec   `json:"spec,omitempty"`
	Status            TargetGroupStatus `json:"status,omitempty"`
}

type TargetGroupSpecHealthCheck struct {
	// +optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled"`
	// +optional
	HealthyThreshold *int64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold"`
	// +optional
	Interval *int64 `json:"interval,omitempty" tf:"interval"`
	// +optional
	Matcher *string `json:"matcher,omitempty" tf:"matcher"`
	// +optional
	Path *string `json:"path,omitempty" tf:"path"`
	// +optional
	Port *string `json:"port,omitempty" tf:"port"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`
	// +optional
	UnhealthyThreshold *int64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold"`
}

type TargetGroupSpecStickiness struct {
	// +optional
	CookieDuration *int64 `json:"cookieDuration,omitempty" tf:"cookie_duration"`
	// +optional
	CookieName *string `json:"cookieName,omitempty" tf:"cookie_name"`
	// +optional
	Enabled *bool   `json:"enabled,omitempty" tf:"enabled"`
	Type    *string `json:"type" tf:"type"`
}

type TargetGroupSpec struct {
	KubeformOutput *TargetGroupSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource TargetGroupSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type TargetGroupSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	ArnSuffix *string `json:"arnSuffix,omitempty" tf:"arn_suffix"`
	// +optional
	DeregistrationDelay *int64 `json:"deregistrationDelay,omitempty" tf:"deregistration_delay"`
	// +optional
	HealthCheck *TargetGroupSpecHealthCheck `json:"healthCheck,omitempty" tf:"health_check"`
	// +optional
	LambdaMultiValueHeadersEnabled *bool `json:"lambdaMultiValueHeadersEnabled,omitempty" tf:"lambda_multi_value_headers_enabled"`
	// +optional
	LoadBalancingAlgorithmType *string `json:"loadBalancingAlgorithmType,omitempty" tf:"load_balancing_algorithm_type"`
	// +optional
	Name *string `json:"name,omitempty" tf:"name"`
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`
	// +optional
	Port *int64 `json:"port,omitempty" tf:"port"`
	// +optional
	PreserveClientIP *string `json:"preserveClientIP,omitempty" tf:"preserve_client_ip"`
	// +optional
	Protocol *string `json:"protocol,omitempty" tf:"protocol"`
	// +optional
	ProtocolVersion *string `json:"protocolVersion,omitempty" tf:"protocol_version"`
	// +optional
	ProxyProtocolV2 *bool `json:"proxyProtocolV2,omitempty" tf:"proxy_protocol_v2"`
	// +optional
	SlowStart *int64 `json:"slowStart,omitempty" tf:"slow_start"`
	// +optional
	Stickiness *TargetGroupSpecStickiness `json:"stickiness,omitempty" tf:"stickiness"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	TargetType *string `json:"targetType,omitempty" tf:"target_type"`
	// +optional
	VpcID *string `json:"vpcID,omitempty" tf:"vpc_id"`
}

type TargetGroupStatus struct {
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

// TargetGroupList is a list of TargetGroups
type TargetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TargetGroup CRD objects
	Items []TargetGroup `json:"items,omitempty"`
}
