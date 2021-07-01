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

type DomainName struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainNameSpec   `json:"spec,omitempty"`
	Status            DomainNameStatus `json:"status,omitempty"`
}

type DomainNameSpecDomainNameConfiguration struct {
	CertificateArn *string `json:"certificateArn" tf:"certificate_arn"`
	EndpointType   *string `json:"endpointType" tf:"endpoint_type"`
	// +optional
	HostedZoneID   *string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id"`
	SecurityPolicy *string `json:"securityPolicy" tf:"security_policy"`
	// +optional
	TargetDomainName *string `json:"targetDomainName,omitempty" tf:"target_domain_name"`
}

type DomainNameSpecMutualTlsAuthentication struct {
	TruststoreURI *string `json:"truststoreURI" tf:"truststore_uri"`
	// +optional
	TruststoreVersion *string `json:"truststoreVersion,omitempty" tf:"truststore_version"`
}

type DomainNameSpec struct {
	KubeformOutput *DomainNameSpecResource `json:"kubeformOutput,omitempty" tf:"-"`

	Resource DomainNameSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DomainNameSpecResource struct {
	Timeouts *base.ResourceTimeout `json:"timeouts,omitempty" tf:"timeouts"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApiMappingSelectionExpression *string `json:"apiMappingSelectionExpression,omitempty" tf:"api_mapping_selection_expression"`
	// +optional
	Arn                     *string                                `json:"arn,omitempty" tf:"arn"`
	DomainName              *string                                `json:"domainName" tf:"domain_name"`
	DomainNameConfiguration *DomainNameSpecDomainNameConfiguration `json:"domainNameConfiguration" tf:"domain_name_configuration"`
	// +optional
	MutualTlsAuthentication *DomainNameSpecMutualTlsAuthentication `json:"mutualTlsAuthentication,omitempty" tf:"mutual_tls_authentication"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type DomainNameStatus struct {
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

// DomainNameList is a list of DomainNames
type DomainNameList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DomainName CRD objects
	Items []DomainName `json:"items,omitempty"`
}
