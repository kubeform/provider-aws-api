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

type ObjectLambdaAccessPoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ObjectLambdaAccessPointSpec   `json:"spec,omitempty"`
	Status            ObjectLambdaAccessPointStatus `json:"status,omitempty"`
}

type ObjectLambdaAccessPointSpecConfigurationTransformationConfigurationContentTransformationAwsLambda struct {
	FunctionArn *string `json:"functionArn" tf:"function_arn"`
	// +optional
	FunctionPayload *string `json:"functionPayload,omitempty" tf:"function_payload"`
}

type ObjectLambdaAccessPointSpecConfigurationTransformationConfigurationContentTransformation struct {
	AwsLambda *ObjectLambdaAccessPointSpecConfigurationTransformationConfigurationContentTransformationAwsLambda `json:"awsLambda" tf:"aws_lambda"`
}

type ObjectLambdaAccessPointSpecConfigurationTransformationConfiguration struct {
	Actions               []string                                                                                  `json:"actions" tf:"actions"`
	ContentTransformation *ObjectLambdaAccessPointSpecConfigurationTransformationConfigurationContentTransformation `json:"contentTransformation" tf:"content_transformation"`
}

type ObjectLambdaAccessPointSpecConfiguration struct {
	// +optional
	AllowedFeatures []string `json:"allowedFeatures,omitempty" tf:"allowed_features"`
	// +optional
	CloudWatchMetricsEnabled    *bool                                                                 `json:"cloudWatchMetricsEnabled,omitempty" tf:"cloud_watch_metrics_enabled"`
	SupportingAccessPoint       *string                                                               `json:"supportingAccessPoint" tf:"supporting_access_point"`
	TransformationConfiguration []ObjectLambdaAccessPointSpecConfigurationTransformationConfiguration `json:"transformationConfiguration" tf:"transformation_configuration"`
}

type ObjectLambdaAccessPointSpec struct {
	State *ObjectLambdaAccessPointSpecResource `json:"state,omitempty" tf:"-"`

	Resource ObjectLambdaAccessPointSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type ObjectLambdaAccessPointSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID *string `json:"accountID,omitempty" tf:"account_id"`
	// +optional
	Arn           *string                                   `json:"arn,omitempty" tf:"arn"`
	Configuration *ObjectLambdaAccessPointSpecConfiguration `json:"configuration" tf:"configuration"`
	Name          *string                                   `json:"name" tf:"name"`
}

type ObjectLambdaAccessPointStatus struct {
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

// ObjectLambdaAccessPointList is a list of ObjectLambdaAccessPoints
type ObjectLambdaAccessPointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ObjectLambdaAccessPoint CRD objects
	Items []ObjectLambdaAccessPoint `json:"items,omitempty"`
}
