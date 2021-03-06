//go:build !ignore_autogenerated
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
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cell) DeepCopyInto(out *Cell) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cell.
func (in *Cell) DeepCopy() *Cell {
	if in == nil {
		return nil
	}
	out := new(Cell)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cell) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CellList) DeepCopyInto(out *CellList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cell, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CellList.
func (in *CellList) DeepCopy() *CellList {
	if in == nil {
		return nil
	}
	out := new(CellList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CellList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CellSpec) DeepCopyInto(out *CellSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CellSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CellSpec.
func (in *CellSpec) DeepCopy() *CellSpec {
	if in == nil {
		return nil
	}
	out := new(CellSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CellSpecResource) DeepCopyInto(out *CellSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.CellName != nil {
		in, out := &in.CellName, &out.CellName
		*out = new(string)
		**out = **in
	}
	if in.Cells != nil {
		in, out := &in.Cells, &out.Cells
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ParentReadinessScopes != nil {
		in, out := &in.ParentReadinessScopes, &out.ParentReadinessScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CellSpecResource.
func (in *CellSpecResource) DeepCopy() *CellSpecResource {
	if in == nil {
		return nil
	}
	out := new(CellSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CellStatus) DeepCopyInto(out *CellStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CellStatus.
func (in *CellStatus) DeepCopy() *CellStatus {
	if in == nil {
		return nil
	}
	out := new(CellStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheck) DeepCopyInto(out *ReadinessCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheck.
func (in *ReadinessCheck) DeepCopy() *ReadinessCheck {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReadinessCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheckList) DeepCopyInto(out *ReadinessCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReadinessCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheckList.
func (in *ReadinessCheckList) DeepCopy() *ReadinessCheckList {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReadinessCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheckSpec) DeepCopyInto(out *ReadinessCheckSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ReadinessCheckSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheckSpec.
func (in *ReadinessCheckSpec) DeepCopy() *ReadinessCheckSpec {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheckSpecResource) DeepCopyInto(out *ReadinessCheckSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ReadinessCheckName != nil {
		in, out := &in.ReadinessCheckName, &out.ReadinessCheckName
		*out = new(string)
		**out = **in
	}
	if in.ResourceSetName != nil {
		in, out := &in.ResourceSetName, &out.ResourceSetName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheckSpecResource.
func (in *ReadinessCheckSpecResource) DeepCopy() *ReadinessCheckSpecResource {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheckSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessCheckStatus) DeepCopyInto(out *ReadinessCheckStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessCheckStatus.
func (in *ReadinessCheckStatus) DeepCopy() *ReadinessCheckStatus {
	if in == nil {
		return nil
	}
	out := new(ReadinessCheckStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryGroup) DeepCopyInto(out *RecoveryGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryGroup.
func (in *RecoveryGroup) DeepCopy() *RecoveryGroup {
	if in == nil {
		return nil
	}
	out := new(RecoveryGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecoveryGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryGroupList) DeepCopyInto(out *RecoveryGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RecoveryGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryGroupList.
func (in *RecoveryGroupList) DeepCopy() *RecoveryGroupList {
	if in == nil {
		return nil
	}
	out := new(RecoveryGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RecoveryGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryGroupSpec) DeepCopyInto(out *RecoveryGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(RecoveryGroupSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryGroupSpec.
func (in *RecoveryGroupSpec) DeepCopy() *RecoveryGroupSpec {
	if in == nil {
		return nil
	}
	out := new(RecoveryGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryGroupSpecResource) DeepCopyInto(out *RecoveryGroupSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Cells != nil {
		in, out := &in.Cells, &out.Cells
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RecoveryGroupName != nil {
		in, out := &in.RecoveryGroupName, &out.RecoveryGroupName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryGroupSpecResource.
func (in *RecoveryGroupSpecResource) DeepCopy() *RecoveryGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(RecoveryGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecoveryGroupStatus) DeepCopyInto(out *RecoveryGroupStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecoveryGroupStatus.
func (in *RecoveryGroupStatus) DeepCopy() *RecoveryGroupStatus {
	if in == nil {
		return nil
	}
	out := new(RecoveryGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSet) DeepCopyInto(out *ResourceSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSet.
func (in *ResourceSet) DeepCopy() *ResourceSet {
	if in == nil {
		return nil
	}
	out := new(ResourceSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetList) DeepCopyInto(out *ResourceSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetList.
func (in *ResourceSetList) DeepCopy() *ResourceSetList {
	if in == nil {
		return nil
	}
	out := new(ResourceSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetSpec) DeepCopyInto(out *ResourceSetSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ResourceSetSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetSpec.
func (in *ResourceSetSpec) DeepCopy() *ResourceSetSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetSpecResource) DeepCopyInto(out *ResourceSetSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ResourceSetName != nil {
		in, out := &in.ResourceSetName, &out.ResourceSetName
		*out = new(string)
		**out = **in
	}
	if in.ResourceSetType != nil {
		in, out := &in.ResourceSetType, &out.ResourceSetType
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceSetSpecResources, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetSpecResource.
func (in *ResourceSetSpecResource) DeepCopy() *ResourceSetSpecResource {
	if in == nil {
		return nil
	}
	out := new(ResourceSetSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetSpecResources) DeepCopyInto(out *ResourceSetSpecResources) {
	*out = *in
	if in.ComponentID != nil {
		in, out := &in.ComponentID, &out.ComponentID
		*out = new(string)
		**out = **in
	}
	if in.DnsTargetResource != nil {
		in, out := &in.DnsTargetResource, &out.DnsTargetResource
		*out = new(ResourceSetSpecResourcesDnsTargetResource)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessScopes != nil {
		in, out := &in.ReadinessScopes, &out.ReadinessScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ResourceArn != nil {
		in, out := &in.ResourceArn, &out.ResourceArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetSpecResources.
func (in *ResourceSetSpecResources) DeepCopy() *ResourceSetSpecResources {
	if in == nil {
		return nil
	}
	out := new(ResourceSetSpecResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetSpecResourcesDnsTargetResource) DeepCopyInto(out *ResourceSetSpecResourcesDnsTargetResource) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.HostedZoneArn != nil {
		in, out := &in.HostedZoneArn, &out.HostedZoneArn
		*out = new(string)
		**out = **in
	}
	if in.RecordSetID != nil {
		in, out := &in.RecordSetID, &out.RecordSetID
		*out = new(string)
		**out = **in
	}
	if in.RecordType != nil {
		in, out := &in.RecordType, &out.RecordType
		*out = new(string)
		**out = **in
	}
	if in.TargetResource != nil {
		in, out := &in.TargetResource, &out.TargetResource
		*out = new(ResourceSetSpecResourcesDnsTargetResourceTargetResource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetSpecResourcesDnsTargetResource.
func (in *ResourceSetSpecResourcesDnsTargetResource) DeepCopy() *ResourceSetSpecResourcesDnsTargetResource {
	if in == nil {
		return nil
	}
	out := new(ResourceSetSpecResourcesDnsTargetResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetSpecResourcesDnsTargetResourceTargetResource) DeepCopyInto(out *ResourceSetSpecResourcesDnsTargetResourceTargetResource) {
	*out = *in
	if in.NlbResource != nil {
		in, out := &in.NlbResource, &out.NlbResource
		*out = new(ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)
		(*in).DeepCopyInto(*out)
	}
	if in.R53Resource != nil {
		in, out := &in.R53Resource, &out.R53Resource
		*out = new(ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetSpecResourcesDnsTargetResourceTargetResource.
func (in *ResourceSetSpecResourcesDnsTargetResourceTargetResource) DeepCopy() *ResourceSetSpecResourcesDnsTargetResourceTargetResource {
	if in == nil {
		return nil
	}
	out := new(ResourceSetSpecResourcesDnsTargetResourceTargetResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource) DeepCopyInto(out *ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource.
func (in *ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource) DeepCopy() *ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource {
	if in == nil {
		return nil
	}
	out := new(ResourceSetSpecResourcesDnsTargetResourceTargetResourceNlbResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource) DeepCopyInto(out *ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource) {
	*out = *in
	if in.DomainName != nil {
		in, out := &in.DomainName, &out.DomainName
		*out = new(string)
		**out = **in
	}
	if in.RecordSetID != nil {
		in, out := &in.RecordSetID, &out.RecordSetID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource.
func (in *ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource) DeepCopy() *ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource {
	if in == nil {
		return nil
	}
	out := new(ResourceSetSpecResourcesDnsTargetResourceTargetResourceR53Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSetStatus) DeepCopyInto(out *ResourceSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSetStatus.
func (in *ResourceSetStatus) DeepCopy() *ResourceSetStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceSetStatus)
	in.DeepCopyInto(out)
	return out
}
