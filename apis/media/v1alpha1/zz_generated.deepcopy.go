// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConvertQueue) DeepCopyInto(out *ConvertQueue) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConvertQueue.
func (in *ConvertQueue) DeepCopy() *ConvertQueue {
	if in == nil {
		return nil
	}
	out := new(ConvertQueue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConvertQueue) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConvertQueueList) DeepCopyInto(out *ConvertQueueList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConvertQueue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConvertQueueList.
func (in *ConvertQueueList) DeepCopy() *ConvertQueueList {
	if in == nil {
		return nil
	}
	out := new(ConvertQueueList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConvertQueueList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConvertQueueSpec) DeepCopyInto(out *ConvertQueueSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ConvertQueueSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConvertQueueSpec.
func (in *ConvertQueueSpec) DeepCopy() *ConvertQueueSpec {
	if in == nil {
		return nil
	}
	out := new(ConvertQueueSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConvertQueueSpecReservationPlanSettings) DeepCopyInto(out *ConvertQueueSpecReservationPlanSettings) {
	*out = *in
	if in.Commitment != nil {
		in, out := &in.Commitment, &out.Commitment
		*out = new(string)
		**out = **in
	}
	if in.RenewalType != nil {
		in, out := &in.RenewalType, &out.RenewalType
		*out = new(string)
		**out = **in
	}
	if in.ReservedSlots != nil {
		in, out := &in.ReservedSlots, &out.ReservedSlots
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConvertQueueSpecReservationPlanSettings.
func (in *ConvertQueueSpecReservationPlanSettings) DeepCopy() *ConvertQueueSpecReservationPlanSettings {
	if in == nil {
		return nil
	}
	out := new(ConvertQueueSpecReservationPlanSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConvertQueueSpecResource) DeepCopyInto(out *ConvertQueueSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.PricingPlan != nil {
		in, out := &in.PricingPlan, &out.PricingPlan
		*out = new(string)
		**out = **in
	}
	if in.ReservationPlanSettings != nil {
		in, out := &in.ReservationPlanSettings, &out.ReservationPlanSettings
		*out = new(ConvertQueueSpecReservationPlanSettings)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConvertQueueSpecResource.
func (in *ConvertQueueSpecResource) DeepCopy() *ConvertQueueSpecResource {
	if in == nil {
		return nil
	}
	out := new(ConvertQueueSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConvertQueueStatus) DeepCopyInto(out *ConvertQueueStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConvertQueueStatus.
func (in *ConvertQueueStatus) DeepCopy() *ConvertQueueStatus {
	if in == nil {
		return nil
	}
	out := new(ConvertQueueStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageChannel) DeepCopyInto(out *PackageChannel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageChannel.
func (in *PackageChannel) DeepCopy() *PackageChannel {
	if in == nil {
		return nil
	}
	out := new(PackageChannel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageChannel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageChannelList) DeepCopyInto(out *PackageChannelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PackageChannel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageChannelList.
func (in *PackageChannelList) DeepCopy() *PackageChannelList {
	if in == nil {
		return nil
	}
	out := new(PackageChannelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageChannelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageChannelSpec) DeepCopyInto(out *PackageChannelSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PackageChannelSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageChannelSpec.
func (in *PackageChannelSpec) DeepCopy() *PackageChannelSpec {
	if in == nil {
		return nil
	}
	out := new(PackageChannelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageChannelSpecHlsIngest) DeepCopyInto(out *PackageChannelSpecHlsIngest) {
	*out = *in
	if in.IngestEndpoints != nil {
		in, out := &in.IngestEndpoints, &out.IngestEndpoints
		*out = make([]PackageChannelSpecHlsIngestIngestEndpoints, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageChannelSpecHlsIngest.
func (in *PackageChannelSpecHlsIngest) DeepCopy() *PackageChannelSpecHlsIngest {
	if in == nil {
		return nil
	}
	out := new(PackageChannelSpecHlsIngest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageChannelSpecHlsIngestIngestEndpoints) DeepCopyInto(out *PackageChannelSpecHlsIngestIngestEndpoints) {
	*out = *in
	if in.Password != nil {
		in, out := &in.Password, &out.Password
		*out = new(string)
		**out = **in
	}
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	if in.Username != nil {
		in, out := &in.Username, &out.Username
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageChannelSpecHlsIngestIngestEndpoints.
func (in *PackageChannelSpecHlsIngestIngestEndpoints) DeepCopy() *PackageChannelSpecHlsIngestIngestEndpoints {
	if in == nil {
		return nil
	}
	out := new(PackageChannelSpecHlsIngestIngestEndpoints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageChannelSpecResource) DeepCopyInto(out *PackageChannelSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ChannelID != nil {
		in, out := &in.ChannelID, &out.ChannelID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.HlsIngest != nil {
		in, out := &in.HlsIngest, &out.HlsIngest
		*out = make([]PackageChannelSpecHlsIngest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageChannelSpecResource.
func (in *PackageChannelSpecResource) DeepCopy() *PackageChannelSpecResource {
	if in == nil {
		return nil
	}
	out := new(PackageChannelSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageChannelStatus) DeepCopyInto(out *PackageChannelStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageChannelStatus.
func (in *PackageChannelStatus) DeepCopy() *PackageChannelStatus {
	if in == nil {
		return nil
	}
	out := new(PackageChannelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainer) DeepCopyInto(out *StoreContainer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainer.
func (in *StoreContainer) DeepCopy() *StoreContainer {
	if in == nil {
		return nil
	}
	out := new(StoreContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreContainer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerList) DeepCopyInto(out *StoreContainerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoreContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerList.
func (in *StoreContainerList) DeepCopy() *StoreContainerList {
	if in == nil {
		return nil
	}
	out := new(StoreContainerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreContainerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerPolicy) DeepCopyInto(out *StoreContainerPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerPolicy.
func (in *StoreContainerPolicy) DeepCopy() *StoreContainerPolicy {
	if in == nil {
		return nil
	}
	out := new(StoreContainerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreContainerPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerPolicyList) DeepCopyInto(out *StoreContainerPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StoreContainerPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerPolicyList.
func (in *StoreContainerPolicyList) DeepCopy() *StoreContainerPolicyList {
	if in == nil {
		return nil
	}
	out := new(StoreContainerPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StoreContainerPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerPolicySpec) DeepCopyInto(out *StoreContainerPolicySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(StoreContainerPolicySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerPolicySpec.
func (in *StoreContainerPolicySpec) DeepCopy() *StoreContainerPolicySpec {
	if in == nil {
		return nil
	}
	out := new(StoreContainerPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerPolicySpecResource) DeepCopyInto(out *StoreContainerPolicySpecResource) {
	*out = *in
	if in.ContainerName != nil {
		in, out := &in.ContainerName, &out.ContainerName
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerPolicySpecResource.
func (in *StoreContainerPolicySpecResource) DeepCopy() *StoreContainerPolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(StoreContainerPolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerPolicyStatus) DeepCopyInto(out *StoreContainerPolicyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerPolicyStatus.
func (in *StoreContainerPolicyStatus) DeepCopy() *StoreContainerPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(StoreContainerPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerSpec) DeepCopyInto(out *StoreContainerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(StoreContainerSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerSpec.
func (in *StoreContainerSpec) DeepCopy() *StoreContainerSpec {
	if in == nil {
		return nil
	}
	out := new(StoreContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerSpecResource) DeepCopyInto(out *StoreContainerSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerSpecResource.
func (in *StoreContainerSpecResource) DeepCopy() *StoreContainerSpecResource {
	if in == nil {
		return nil
	}
	out := new(StoreContainerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreContainerStatus) DeepCopyInto(out *StoreContainerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreContainerStatus.
func (in *StoreContainerStatus) DeepCopy() *StoreContainerStatus {
	if in == nil {
		return nil
	}
	out := new(StoreContainerStatus)
	in.DeepCopyInto(out)
	return out
}
