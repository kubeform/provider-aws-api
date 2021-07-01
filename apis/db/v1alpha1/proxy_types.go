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

type Proxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProxySpec   `json:"spec,omitempty"`
	Status            ProxyStatus `json:"status,omitempty"`
}

type ProxySpecAuth struct {
	// +optional
	AuthScheme *string `json:"authScheme,omitempty" tf:"auth_scheme"`
	// +optional
	Description *string `json:"description,omitempty" tf:"description"`
	// +optional
	IamAuth *string `json:"iamAuth,omitempty" tf:"iam_auth"`
	// +optional
	SecretArn *string `json:"secretArn,omitempty" tf:"secret_arn"`
}

type ProxySpec struct {
	KubeformOutput *ProxySpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource ProxySpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ProxySpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn  *string         `json:"arn,omitempty" tf:"arn"`
	Auth []ProxySpecAuth `json:"auth" tf:"auth"`
	// +optional
	DebugLogging *bool `json:"debugLogging,omitempty" tf:"debug_logging"`
	// +optional
	Endpoint     *string `json:"endpoint,omitempty" tf:"endpoint"`
	EngineFamily *string `json:"engineFamily" tf:"engine_family"`
	// +optional
	IdleClientTimeout *int64  `json:"idleClientTimeout,omitempty" tf:"idle_client_timeout"`
	Name              *string `json:"name" tf:"name"`
	// +optional
	RequireTls *bool   `json:"requireTls,omitempty" tf:"require_tls"`
	RoleArn    *string `json:"roleArn" tf:"role_arn"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	// +optional
	VpcSecurityGroupIDS []string `json:"vpcSecurityGroupIDS,omitempty" tf:"vpc_security_group_ids"`
	VpcSubnetIDS        []string `json:"vpcSubnetIDS" tf:"vpc_subnet_ids"`
}

type ProxyStatus struct {
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

// ProxyList is a list of Proxys
type ProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Proxy CRD objects
	Items []Proxy `json:"items,omitempty"`
}
