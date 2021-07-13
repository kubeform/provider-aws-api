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

type Route struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteSpec   `json:"spec,omitempty"`
	Status            RouteStatus `json:"status,omitempty"`
}

type RouteSpecSpecGrpcRouteActionWeightedTarget struct {
	VirtualNode *string `json:"virtualNode" tf:"virtual_node"`
	Weight      *int64  `json:"weight" tf:"weight"`
}

type RouteSpecSpecGrpcRouteAction struct {
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	WeightedTarget []RouteSpecSpecGrpcRouteActionWeightedTarget `json:"weightedTarget" tf:"weighted_target"`
}

type RouteSpecSpecGrpcRouteMatchMetadataMatchRange struct {
	End   *int64 `json:"end" tf:"end"`
	Start *int64 `json:"start" tf:"start"`
}

type RouteSpecSpecGrpcRouteMatchMetadataMatch struct {
	// +optional
	Exact *string `json:"exact,omitempty" tf:"exact"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	Range *RouteSpecSpecGrpcRouteMatchMetadataMatchRange `json:"range,omitempty" tf:"range"`
	// +optional
	Regex *string `json:"regex,omitempty" tf:"regex"`
	// +optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix"`
}

type RouteSpecSpecGrpcRouteMatchMetadata struct {
	// +optional
	Invert *bool `json:"invert,omitempty" tf:"invert"`
	// +optional
	Match *RouteSpecSpecGrpcRouteMatchMetadataMatch `json:"match,omitempty" tf:"match"`
	Name  *string                                   `json:"name" tf:"name"`
}

type RouteSpecSpecGrpcRouteMatch struct {
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Metadata []RouteSpecSpecGrpcRouteMatchMetadata `json:"metadata,omitempty" tf:"metadata"`
	// +optional
	MethodName *string `json:"methodName,omitempty" tf:"method_name"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	ServiceName *string `json:"serviceName,omitempty" tf:"service_name"`
}

type RouteSpecSpecGrpcRouteRetryPolicyPerRetryTimeout struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecGrpcRouteRetryPolicy struct {
	// +optional
	GrpcRetryEvents []string `json:"grpcRetryEvents,omitempty" tf:"grpc_retry_events"`
	// +optional
	HttpRetryEvents []string                                          `json:"httpRetryEvents,omitempty" tf:"http_retry_events"`
	MaxRetries      *int64                                            `json:"maxRetries" tf:"max_retries"`
	PerRetryTimeout *RouteSpecSpecGrpcRouteRetryPolicyPerRetryTimeout `json:"perRetryTimeout" tf:"per_retry_timeout"`
	// +optional
	TcpRetryEvents []string `json:"tcpRetryEvents,omitempty" tf:"tcp_retry_events"`
}

type RouteSpecSpecGrpcRouteTimeoutIdle struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecGrpcRouteTimeoutPerRequest struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecGrpcRouteTimeout struct {
	// +optional
	Idle *RouteSpecSpecGrpcRouteTimeoutIdle `json:"idle,omitempty" tf:"idle"`
	// +optional
	PerRequest *RouteSpecSpecGrpcRouteTimeoutPerRequest `json:"perRequest,omitempty" tf:"per_request"`
}

type RouteSpecSpecGrpcRoute struct {
	Action *RouteSpecSpecGrpcRouteAction `json:"action" tf:"action"`
	// +optional
	Match *RouteSpecSpecGrpcRouteMatch `json:"match,omitempty" tf:"match"`
	// +optional
	RetryPolicy *RouteSpecSpecGrpcRouteRetryPolicy `json:"retryPolicy,omitempty" tf:"retry_policy"`
	// +optional
	Timeout *RouteSpecSpecGrpcRouteTimeout `json:"timeout,omitempty" tf:"timeout"`
}

type RouteSpecSpecHttp2RouteActionWeightedTarget struct {
	VirtualNode *string `json:"virtualNode" tf:"virtual_node"`
	Weight      *int64  `json:"weight" tf:"weight"`
}

type RouteSpecSpecHttp2RouteAction struct {
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	WeightedTarget []RouteSpecSpecHttp2RouteActionWeightedTarget `json:"weightedTarget" tf:"weighted_target"`
}

type RouteSpecSpecHttp2RouteMatchHeaderMatchRange struct {
	End   *int64 `json:"end" tf:"end"`
	Start *int64 `json:"start" tf:"start"`
}

type RouteSpecSpecHttp2RouteMatchHeaderMatch struct {
	// +optional
	Exact *string `json:"exact,omitempty" tf:"exact"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	Range *RouteSpecSpecHttp2RouteMatchHeaderMatchRange `json:"range,omitempty" tf:"range"`
	// +optional
	Regex *string `json:"regex,omitempty" tf:"regex"`
	// +optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix"`
}

type RouteSpecSpecHttp2RouteMatchHeader struct {
	// +optional
	Invert *bool `json:"invert,omitempty" tf:"invert"`
	// +optional
	Match *RouteSpecSpecHttp2RouteMatchHeaderMatch `json:"match,omitempty" tf:"match"`
	Name  *string                                  `json:"name" tf:"name"`
}

type RouteSpecSpecHttp2RouteMatch struct {
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Header []RouteSpecSpecHttp2RouteMatchHeader `json:"header,omitempty" tf:"header"`
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	Prefix *string `json:"prefix" tf:"prefix"`
	// +optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme"`
}

type RouteSpecSpecHttp2RouteRetryPolicyPerRetryTimeout struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecHttp2RouteRetryPolicy struct {
	// +optional
	HttpRetryEvents []string                                           `json:"httpRetryEvents,omitempty" tf:"http_retry_events"`
	MaxRetries      *int64                                             `json:"maxRetries" tf:"max_retries"`
	PerRetryTimeout *RouteSpecSpecHttp2RouteRetryPolicyPerRetryTimeout `json:"perRetryTimeout" tf:"per_retry_timeout"`
	// +optional
	TcpRetryEvents []string `json:"tcpRetryEvents,omitempty" tf:"tcp_retry_events"`
}

type RouteSpecSpecHttp2RouteTimeoutIdle struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecHttp2RouteTimeoutPerRequest struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecHttp2RouteTimeout struct {
	// +optional
	Idle *RouteSpecSpecHttp2RouteTimeoutIdle `json:"idle,omitempty" tf:"idle"`
	// +optional
	PerRequest *RouteSpecSpecHttp2RouteTimeoutPerRequest `json:"perRequest,omitempty" tf:"per_request"`
}

type RouteSpecSpecHttp2Route struct {
	Action *RouteSpecSpecHttp2RouteAction `json:"action" tf:"action"`
	Match  *RouteSpecSpecHttp2RouteMatch  `json:"match" tf:"match"`
	// +optional
	RetryPolicy *RouteSpecSpecHttp2RouteRetryPolicy `json:"retryPolicy,omitempty" tf:"retry_policy"`
	// +optional
	Timeout *RouteSpecSpecHttp2RouteTimeout `json:"timeout,omitempty" tf:"timeout"`
}

type RouteSpecSpecHttpRouteActionWeightedTarget struct {
	VirtualNode *string `json:"virtualNode" tf:"virtual_node"`
	Weight      *int64  `json:"weight" tf:"weight"`
}

type RouteSpecSpecHttpRouteAction struct {
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	WeightedTarget []RouteSpecSpecHttpRouteActionWeightedTarget `json:"weightedTarget" tf:"weighted_target"`
}

type RouteSpecSpecHttpRouteMatchHeaderMatchRange struct {
	End   *int64 `json:"end" tf:"end"`
	Start *int64 `json:"start" tf:"start"`
}

type RouteSpecSpecHttpRouteMatchHeaderMatch struct {
	// +optional
	Exact *string `json:"exact,omitempty" tf:"exact"`
	// +optional
	Prefix *string `json:"prefix,omitempty" tf:"prefix"`
	// +optional
	Range *RouteSpecSpecHttpRouteMatchHeaderMatchRange `json:"range,omitempty" tf:"range"`
	// +optional
	Regex *string `json:"regex,omitempty" tf:"regex"`
	// +optional
	Suffix *string `json:"suffix,omitempty" tf:"suffix"`
}

type RouteSpecSpecHttpRouteMatchHeader struct {
	// +optional
	Invert *bool `json:"invert,omitempty" tf:"invert"`
	// +optional
	Match *RouteSpecSpecHttpRouteMatchHeaderMatch `json:"match,omitempty" tf:"match"`
	Name  *string                                 `json:"name" tf:"name"`
}

type RouteSpecSpecHttpRouteMatch struct {
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Header []RouteSpecSpecHttpRouteMatchHeader `json:"header,omitempty" tf:"header"`
	// +optional
	Method *string `json:"method,omitempty" tf:"method"`
	Prefix *string `json:"prefix" tf:"prefix"`
	// +optional
	Scheme *string `json:"scheme,omitempty" tf:"scheme"`
}

type RouteSpecSpecHttpRouteRetryPolicyPerRetryTimeout struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecHttpRouteRetryPolicy struct {
	// +optional
	HttpRetryEvents []string                                          `json:"httpRetryEvents,omitempty" tf:"http_retry_events"`
	MaxRetries      *int64                                            `json:"maxRetries" tf:"max_retries"`
	PerRetryTimeout *RouteSpecSpecHttpRouteRetryPolicyPerRetryTimeout `json:"perRetryTimeout" tf:"per_retry_timeout"`
	// +optional
	TcpRetryEvents []string `json:"tcpRetryEvents,omitempty" tf:"tcp_retry_events"`
}

type RouteSpecSpecHttpRouteTimeoutIdle struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecHttpRouteTimeoutPerRequest struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecHttpRouteTimeout struct {
	// +optional
	Idle *RouteSpecSpecHttpRouteTimeoutIdle `json:"idle,omitempty" tf:"idle"`
	// +optional
	PerRequest *RouteSpecSpecHttpRouteTimeoutPerRequest `json:"perRequest,omitempty" tf:"per_request"`
}

type RouteSpecSpecHttpRoute struct {
	Action *RouteSpecSpecHttpRouteAction `json:"action" tf:"action"`
	Match  *RouteSpecSpecHttpRouteMatch  `json:"match" tf:"match"`
	// +optional
	RetryPolicy *RouteSpecSpecHttpRouteRetryPolicy `json:"retryPolicy,omitempty" tf:"retry_policy"`
	// +optional
	Timeout *RouteSpecSpecHttpRouteTimeout `json:"timeout,omitempty" tf:"timeout"`
}

type RouteSpecSpecTcpRouteActionWeightedTarget struct {
	VirtualNode *string `json:"virtualNode" tf:"virtual_node"`
	Weight      *int64  `json:"weight" tf:"weight"`
}

type RouteSpecSpecTcpRouteAction struct {
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	WeightedTarget []RouteSpecSpecTcpRouteActionWeightedTarget `json:"weightedTarget" tf:"weighted_target"`
}

type RouteSpecSpecTcpRouteTimeoutIdle struct {
	Unit  *string `json:"unit" tf:"unit"`
	Value *int64  `json:"value" tf:"value"`
}

type RouteSpecSpecTcpRouteTimeout struct {
	// +optional
	Idle *RouteSpecSpecTcpRouteTimeoutIdle `json:"idle,omitempty" tf:"idle"`
}

type RouteSpecSpecTcpRoute struct {
	Action *RouteSpecSpecTcpRouteAction `json:"action" tf:"action"`
	// +optional
	Timeout *RouteSpecSpecTcpRouteTimeout `json:"timeout,omitempty" tf:"timeout"`
}

type RouteSpecSpec struct {
	// +optional
	GrpcRoute *RouteSpecSpecGrpcRoute `json:"grpcRoute,omitempty" tf:"grpc_route"`
	// +optional
	Http2Route *RouteSpecSpecHttp2Route `json:"http2Route,omitempty" tf:"http2_route"`
	// +optional
	HttpRoute *RouteSpecSpecHttpRoute `json:"httpRoute,omitempty" tf:"http_route"`
	// +optional
	Priority *int64 `json:"priority,omitempty" tf:"priority"`
	// +optional
	TcpRoute *RouteSpecSpecTcpRoute `json:"tcpRoute,omitempty" tf:"tcp_route"`
}

type RouteSpec struct {
	State *RouteSpecResource `json:"state,omitempty" tf:"-"`

	Resource RouteSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type RouteSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CreatedDate *string `json:"createdDate,omitempty" tf:"created_date"`
	// +optional
	LastUpdatedDate *string `json:"lastUpdatedDate,omitempty" tf:"last_updated_date"`
	MeshName        *string `json:"meshName" tf:"mesh_name"`
	// +optional
	MeshOwner *string `json:"meshOwner,omitempty" tf:"mesh_owner"`
	Name      *string `json:"name" tf:"name"`
	// +optional
	ResourceOwner *string        `json:"resourceOwner,omitempty" tf:"resource_owner"`
	Spec          *RouteSpecSpec `json:"spec" tf:"spec"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TagsAll           *map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
	VirtualRouterName *string            `json:"virtualRouterName" tf:"virtual_router_name"`
}

type RouteStatus struct {
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

// RouteList is a list of Routes
type RouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Route CRD objects
	Items []Route `json:"items,omitempty"`
}
