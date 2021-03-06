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

type Insight struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InsightSpec   `json:"spec,omitempty"`
	Status            InsightStatus `json:"status,omitempty"`
}

type InsightSpecFiltersAwsAccountID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersCompanyName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersComplianceStatus struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersConfidence struct {
	// +optional
	Eq *string `json:"eq,omitempty" tf:"eq"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type InsightSpecFiltersCreatedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersCreatedAt struct {
	// +optional
	DateRange *InsightSpecFiltersCreatedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersCriticality struct {
	// +optional
	Eq *string `json:"eq,omitempty" tf:"eq"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type InsightSpecFiltersDescription struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersFindingProviderFieldsConfidence struct {
	// +optional
	Eq *string `json:"eq,omitempty" tf:"eq"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type InsightSpecFiltersFindingProviderFieldsCriticality struct {
	// +optional
	Eq *string `json:"eq,omitempty" tf:"eq"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type InsightSpecFiltersFindingProviderFieldsRelatedFindingsID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersFindingProviderFieldsRelatedFindingsProductArn struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersFindingProviderFieldsSeverityLabel struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersFindingProviderFieldsSeverityOriginal struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersFindingProviderFieldsTypes struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersFirstObservedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersFirstObservedAt struct {
	// +optional
	DateRange *InsightSpecFiltersFirstObservedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersGeneratorID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersKeyword struct {
	Value *string `json:"value" tf:"value"`
}

type InsightSpecFiltersLastObservedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersLastObservedAt struct {
	// +optional
	DateRange *InsightSpecFiltersLastObservedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersMalwareName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersMalwarePath struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersMalwareState struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersMalwareType struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersNetworkDestinationDomain struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersNetworkDestinationIpv4 struct {
	Cidr *string `json:"cidr" tf:"cidr"`
}

type InsightSpecFiltersNetworkDestinationIpv6 struct {
	Cidr *string `json:"cidr" tf:"cidr"`
}

type InsightSpecFiltersNetworkDestinationPort struct {
	// +optional
	Eq *string `json:"eq,omitempty" tf:"eq"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type InsightSpecFiltersNetworkDirection struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersNetworkProtocol struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersNetworkSourceDomain struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersNetworkSourceIpv4 struct {
	Cidr *string `json:"cidr" tf:"cidr"`
}

type InsightSpecFiltersNetworkSourceIpv6 struct {
	Cidr *string `json:"cidr" tf:"cidr"`
}

type InsightSpecFiltersNetworkSourceMAC struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersNetworkSourcePort struct {
	// +optional
	Eq *string `json:"eq,omitempty" tf:"eq"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type InsightSpecFiltersNoteText struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersNoteUpdatedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersNoteUpdatedAt struct {
	// +optional
	DateRange *InsightSpecFiltersNoteUpdatedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersNoteUpdatedBy struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersProcessLaunchedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersProcessLaunchedAt struct {
	// +optional
	DateRange *InsightSpecFiltersProcessLaunchedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersProcessName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersProcessParentPid struct {
	// +optional
	Eq *string `json:"eq,omitempty" tf:"eq"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type InsightSpecFiltersProcessPath struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersProcessPid struct {
	// +optional
	Eq *string `json:"eq,omitempty" tf:"eq"`
	// +optional
	Gte *string `json:"gte,omitempty" tf:"gte"`
	// +optional
	Lte *string `json:"lte,omitempty" tf:"lte"`
}

type InsightSpecFiltersProcessTerminatedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersProcessTerminatedAt struct {
	// +optional
	DateRange *InsightSpecFiltersProcessTerminatedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersProductArn struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersProductFields struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Key        *string `json:"key" tf:"key"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersProductName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersRecommendationText struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersRecordState struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersRelatedFindingsID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersRelatedFindingsProductArn struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsEc2InstanceIamInstanceProfileArn struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsEc2InstanceImageID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsEc2InstanceIpv4Addresses struct {
	Cidr *string `json:"cidr" tf:"cidr"`
}

type InsightSpecFiltersResourceAwsEc2InstanceIpv6Addresses struct {
	Cidr *string `json:"cidr" tf:"cidr"`
}

type InsightSpecFiltersResourceAwsEc2InstanceKeyName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsEc2InstanceLaunchedAt struct {
	// +optional
	DateRange *InsightSpecFiltersResourceAwsEc2InstanceLaunchedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersResourceAwsEc2InstanceSubnetID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsEc2InstanceType struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsEc2InstanceVpcID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsIamAccessKeyCreatedAt struct {
	// +optional
	DateRange *InsightSpecFiltersResourceAwsIamAccessKeyCreatedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersResourceAwsIamAccessKeyStatus struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsIamAccessKeyUserName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsS3BucketOwnerID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceAwsS3BucketOwnerName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceContainerImageID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceContainerImageName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceContainerLaunchedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceContainerLaunchedAt struct {
	// +optional
	DateRange *InsightSpecFiltersResourceContainerLaunchedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersResourceContainerName struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceDetailsOther struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Key        *string `json:"key" tf:"key"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceID struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourcePartition struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceRegion struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceTags struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Key        *string `json:"key" tf:"key"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersResourceType struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersSeverityLabel struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersSourceURL struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersThreatIntelIndicatorCategory struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersThreatIntelIndicatorLastObservedAt struct {
	// +optional
	DateRange *InsightSpecFiltersThreatIntelIndicatorLastObservedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersThreatIntelIndicatorSource struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersThreatIntelIndicatorSourceURL struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersThreatIntelIndicatorType struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersThreatIntelIndicatorValue struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersTitle struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersType struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersUpdatedAtDateRange struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type InsightSpecFiltersUpdatedAt struct {
	// +optional
	DateRange *InsightSpecFiltersUpdatedAtDateRange `json:"dateRange,omitempty" tf:"date_range"`
	// +optional
	End *string `json:"end,omitempty" tf:"end"`
	// +optional
	Start *string `json:"start,omitempty" tf:"start"`
}

type InsightSpecFiltersUserDefinedValues struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Key        *string `json:"key" tf:"key"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersVerificationState struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFiltersWorkflowStatus struct {
	Comparison *string `json:"comparison" tf:"comparison"`
	Value      *string `json:"value" tf:"value"`
}

type InsightSpecFilters struct {
	// +optional
	// +kubebuilder:validation:MaxItems=20
	AwsAccountID []InsightSpecFiltersAwsAccountID `json:"awsAccountID,omitempty" tf:"aws_account_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	CompanyName []InsightSpecFiltersCompanyName `json:"companyName,omitempty" tf:"company_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ComplianceStatus []InsightSpecFiltersComplianceStatus `json:"complianceStatus,omitempty" tf:"compliance_status"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	Confidence []InsightSpecFiltersConfidence `json:"confidence,omitempty" tf:"confidence"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	CreatedAt []InsightSpecFiltersCreatedAt `json:"createdAt,omitempty" tf:"created_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	Criticality []InsightSpecFiltersCriticality `json:"criticality,omitempty" tf:"criticality"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	Description []InsightSpecFiltersDescription `json:"description,omitempty" tf:"description"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	FindingProviderFieldsConfidence []InsightSpecFiltersFindingProviderFieldsConfidence `json:"findingProviderFieldsConfidence,omitempty" tf:"finding_provider_fields_confidence"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	FindingProviderFieldsCriticality []InsightSpecFiltersFindingProviderFieldsCriticality `json:"findingProviderFieldsCriticality,omitempty" tf:"finding_provider_fields_criticality"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	FindingProviderFieldsRelatedFindingsID []InsightSpecFiltersFindingProviderFieldsRelatedFindingsID `json:"findingProviderFieldsRelatedFindingsID,omitempty" tf:"finding_provider_fields_related_findings_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	FindingProviderFieldsRelatedFindingsProductArn []InsightSpecFiltersFindingProviderFieldsRelatedFindingsProductArn `json:"findingProviderFieldsRelatedFindingsProductArn,omitempty" tf:"finding_provider_fields_related_findings_product_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	FindingProviderFieldsSeverityLabel []InsightSpecFiltersFindingProviderFieldsSeverityLabel `json:"findingProviderFieldsSeverityLabel,omitempty" tf:"finding_provider_fields_severity_label"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	FindingProviderFieldsSeverityOriginal []InsightSpecFiltersFindingProviderFieldsSeverityOriginal `json:"findingProviderFieldsSeverityOriginal,omitempty" tf:"finding_provider_fields_severity_original"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	FindingProviderFieldsTypes []InsightSpecFiltersFindingProviderFieldsTypes `json:"findingProviderFieldsTypes,omitempty" tf:"finding_provider_fields_types"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	FirstObservedAt []InsightSpecFiltersFirstObservedAt `json:"firstObservedAt,omitempty" tf:"first_observed_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	GeneratorID []InsightSpecFiltersGeneratorID `json:"generatorID,omitempty" tf:"generator_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ID []InsightSpecFiltersID `json:"ID,omitempty" tf:"id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	Keyword []InsightSpecFiltersKeyword `json:"keyword,omitempty" tf:"keyword"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	LastObservedAt []InsightSpecFiltersLastObservedAt `json:"lastObservedAt,omitempty" tf:"last_observed_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	MalwareName []InsightSpecFiltersMalwareName `json:"malwareName,omitempty" tf:"malware_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	MalwarePath []InsightSpecFiltersMalwarePath `json:"malwarePath,omitempty" tf:"malware_path"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	MalwareState []InsightSpecFiltersMalwareState `json:"malwareState,omitempty" tf:"malware_state"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	MalwareType []InsightSpecFiltersMalwareType `json:"malwareType,omitempty" tf:"malware_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkDestinationDomain []InsightSpecFiltersNetworkDestinationDomain `json:"networkDestinationDomain,omitempty" tf:"network_destination_domain"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkDestinationIpv4 []InsightSpecFiltersNetworkDestinationIpv4 `json:"networkDestinationIpv4,omitempty" tf:"network_destination_ipv4"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkDestinationIpv6 []InsightSpecFiltersNetworkDestinationIpv6 `json:"networkDestinationIpv6,omitempty" tf:"network_destination_ipv6"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkDestinationPort []InsightSpecFiltersNetworkDestinationPort `json:"networkDestinationPort,omitempty" tf:"network_destination_port"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkDirection []InsightSpecFiltersNetworkDirection `json:"networkDirection,omitempty" tf:"network_direction"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkProtocol []InsightSpecFiltersNetworkProtocol `json:"networkProtocol,omitempty" tf:"network_protocol"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkSourceDomain []InsightSpecFiltersNetworkSourceDomain `json:"networkSourceDomain,omitempty" tf:"network_source_domain"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkSourceIpv4 []InsightSpecFiltersNetworkSourceIpv4 `json:"networkSourceIpv4,omitempty" tf:"network_source_ipv4"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkSourceIpv6 []InsightSpecFiltersNetworkSourceIpv6 `json:"networkSourceIpv6,omitempty" tf:"network_source_ipv6"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkSourceMAC []InsightSpecFiltersNetworkSourceMAC `json:"networkSourceMAC,omitempty" tf:"network_source_mac"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NetworkSourcePort []InsightSpecFiltersNetworkSourcePort `json:"networkSourcePort,omitempty" tf:"network_source_port"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NoteText []InsightSpecFiltersNoteText `json:"noteText,omitempty" tf:"note_text"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NoteUpdatedAt []InsightSpecFiltersNoteUpdatedAt `json:"noteUpdatedAt,omitempty" tf:"note_updated_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	NoteUpdatedBy []InsightSpecFiltersNoteUpdatedBy `json:"noteUpdatedBy,omitempty" tf:"note_updated_by"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProcessLaunchedAt []InsightSpecFiltersProcessLaunchedAt `json:"processLaunchedAt,omitempty" tf:"process_launched_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProcessName []InsightSpecFiltersProcessName `json:"processName,omitempty" tf:"process_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProcessParentPid []InsightSpecFiltersProcessParentPid `json:"processParentPid,omitempty" tf:"process_parent_pid"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProcessPath []InsightSpecFiltersProcessPath `json:"processPath,omitempty" tf:"process_path"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProcessPid []InsightSpecFiltersProcessPid `json:"processPid,omitempty" tf:"process_pid"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProcessTerminatedAt []InsightSpecFiltersProcessTerminatedAt `json:"processTerminatedAt,omitempty" tf:"process_terminated_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProductArn []InsightSpecFiltersProductArn `json:"productArn,omitempty" tf:"product_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProductFields []InsightSpecFiltersProductFields `json:"productFields,omitempty" tf:"product_fields"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ProductName []InsightSpecFiltersProductName `json:"productName,omitempty" tf:"product_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	RecommendationText []InsightSpecFiltersRecommendationText `json:"recommendationText,omitempty" tf:"recommendation_text"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	RecordState []InsightSpecFiltersRecordState `json:"recordState,omitempty" tf:"record_state"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	RelatedFindingsID []InsightSpecFiltersRelatedFindingsID `json:"relatedFindingsID,omitempty" tf:"related_findings_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	RelatedFindingsProductArn []InsightSpecFiltersRelatedFindingsProductArn `json:"relatedFindingsProductArn,omitempty" tf:"related_findings_product_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceIamInstanceProfileArn []InsightSpecFiltersResourceAwsEc2InstanceIamInstanceProfileArn `json:"resourceAwsEc2InstanceIamInstanceProfileArn,omitempty" tf:"resource_aws_ec2_instance_iam_instance_profile_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceImageID []InsightSpecFiltersResourceAwsEc2InstanceImageID `json:"resourceAwsEc2InstanceImageID,omitempty" tf:"resource_aws_ec2_instance_image_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceIpv4Addresses []InsightSpecFiltersResourceAwsEc2InstanceIpv4Addresses `json:"resourceAwsEc2InstanceIpv4Addresses,omitempty" tf:"resource_aws_ec2_instance_ipv4_addresses"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceIpv6Addresses []InsightSpecFiltersResourceAwsEc2InstanceIpv6Addresses `json:"resourceAwsEc2InstanceIpv6Addresses,omitempty" tf:"resource_aws_ec2_instance_ipv6_addresses"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceKeyName []InsightSpecFiltersResourceAwsEc2InstanceKeyName `json:"resourceAwsEc2InstanceKeyName,omitempty" tf:"resource_aws_ec2_instance_key_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceLaunchedAt []InsightSpecFiltersResourceAwsEc2InstanceLaunchedAt `json:"resourceAwsEc2InstanceLaunchedAt,omitempty" tf:"resource_aws_ec2_instance_launched_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceSubnetID []InsightSpecFiltersResourceAwsEc2InstanceSubnetID `json:"resourceAwsEc2InstanceSubnetID,omitempty" tf:"resource_aws_ec2_instance_subnet_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceType []InsightSpecFiltersResourceAwsEc2InstanceType `json:"resourceAwsEc2InstanceType,omitempty" tf:"resource_aws_ec2_instance_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsEc2InstanceVpcID []InsightSpecFiltersResourceAwsEc2InstanceVpcID `json:"resourceAwsEc2InstanceVpcID,omitempty" tf:"resource_aws_ec2_instance_vpc_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsIamAccessKeyCreatedAt []InsightSpecFiltersResourceAwsIamAccessKeyCreatedAt `json:"resourceAwsIamAccessKeyCreatedAt,omitempty" tf:"resource_aws_iam_access_key_created_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsIamAccessKeyStatus []InsightSpecFiltersResourceAwsIamAccessKeyStatus `json:"resourceAwsIamAccessKeyStatus,omitempty" tf:"resource_aws_iam_access_key_status"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsIamAccessKeyUserName []InsightSpecFiltersResourceAwsIamAccessKeyUserName `json:"resourceAwsIamAccessKeyUserName,omitempty" tf:"resource_aws_iam_access_key_user_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsS3BucketOwnerID []InsightSpecFiltersResourceAwsS3BucketOwnerID `json:"resourceAwsS3BucketOwnerID,omitempty" tf:"resource_aws_s3_bucket_owner_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceAwsS3BucketOwnerName []InsightSpecFiltersResourceAwsS3BucketOwnerName `json:"resourceAwsS3BucketOwnerName,omitempty" tf:"resource_aws_s3_bucket_owner_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceContainerImageID []InsightSpecFiltersResourceContainerImageID `json:"resourceContainerImageID,omitempty" tf:"resource_container_image_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceContainerImageName []InsightSpecFiltersResourceContainerImageName `json:"resourceContainerImageName,omitempty" tf:"resource_container_image_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceContainerLaunchedAt []InsightSpecFiltersResourceContainerLaunchedAt `json:"resourceContainerLaunchedAt,omitempty" tf:"resource_container_launched_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceContainerName []InsightSpecFiltersResourceContainerName `json:"resourceContainerName,omitempty" tf:"resource_container_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceDetailsOther []InsightSpecFiltersResourceDetailsOther `json:"resourceDetailsOther,omitempty" tf:"resource_details_other"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceID []InsightSpecFiltersResourceID `json:"resourceID,omitempty" tf:"resource_id"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourcePartition []InsightSpecFiltersResourcePartition `json:"resourcePartition,omitempty" tf:"resource_partition"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceRegion []InsightSpecFiltersResourceRegion `json:"resourceRegion,omitempty" tf:"resource_region"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceTags []InsightSpecFiltersResourceTags `json:"resourceTags,omitempty" tf:"resource_tags"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ResourceType []InsightSpecFiltersResourceType `json:"resourceType,omitempty" tf:"resource_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	SeverityLabel []InsightSpecFiltersSeverityLabel `json:"severityLabel,omitempty" tf:"severity_label"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	SourceURL []InsightSpecFiltersSourceURL `json:"sourceURL,omitempty" tf:"source_url"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ThreatIntelIndicatorCategory []InsightSpecFiltersThreatIntelIndicatorCategory `json:"threatIntelIndicatorCategory,omitempty" tf:"threat_intel_indicator_category"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ThreatIntelIndicatorLastObservedAt []InsightSpecFiltersThreatIntelIndicatorLastObservedAt `json:"threatIntelIndicatorLastObservedAt,omitempty" tf:"threat_intel_indicator_last_observed_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ThreatIntelIndicatorSource []InsightSpecFiltersThreatIntelIndicatorSource `json:"threatIntelIndicatorSource,omitempty" tf:"threat_intel_indicator_source"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ThreatIntelIndicatorSourceURL []InsightSpecFiltersThreatIntelIndicatorSourceURL `json:"threatIntelIndicatorSourceURL,omitempty" tf:"threat_intel_indicator_source_url"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ThreatIntelIndicatorType []InsightSpecFiltersThreatIntelIndicatorType `json:"threatIntelIndicatorType,omitempty" tf:"threat_intel_indicator_type"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	ThreatIntelIndicatorValue []InsightSpecFiltersThreatIntelIndicatorValue `json:"threatIntelIndicatorValue,omitempty" tf:"threat_intel_indicator_value"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	Title []InsightSpecFiltersTitle `json:"title,omitempty" tf:"title"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	Type []InsightSpecFiltersType `json:"type,omitempty" tf:"type"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	UpdatedAt []InsightSpecFiltersUpdatedAt `json:"updatedAt,omitempty" tf:"updated_at"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	UserDefinedValues []InsightSpecFiltersUserDefinedValues `json:"userDefinedValues,omitempty" tf:"user_defined_values"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	VerificationState []InsightSpecFiltersVerificationState `json:"verificationState,omitempty" tf:"verification_state"`
	// +optional
	// +kubebuilder:validation:MaxItems=20
	WorkflowStatus []InsightSpecFiltersWorkflowStatus `json:"workflowStatus,omitempty" tf:"workflow_status"`
}

type InsightSpec struct {
	State *InsightSpecResource `json:"state,omitempty" tf:"-"`

	Resource InsightSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type InsightSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn              *string             `json:"arn,omitempty" tf:"arn"`
	Filters          *InsightSpecFilters `json:"filters" tf:"filters"`
	GroupByAttribute *string             `json:"groupByAttribute" tf:"group_by_attribute"`
	Name             *string             `json:"name" tf:"name"`
}

type InsightStatus struct {
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

// InsightList is a list of Insights
type InsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Insight CRD objects
	Items []Insight `json:"items,omitempty"`
}
