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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDefinition) DeepCopyInto(out *ReportDefinition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDefinition.
func (in *ReportDefinition) DeepCopy() *ReportDefinition {
	if in == nil {
		return nil
	}
	out := new(ReportDefinition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportDefinition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDefinitionList) DeepCopyInto(out *ReportDefinitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReportDefinition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDefinitionList.
func (in *ReportDefinitionList) DeepCopy() *ReportDefinitionList {
	if in == nil {
		return nil
	}
	out := new(ReportDefinitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportDefinitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDefinitionSpec) DeepCopyInto(out *ReportDefinitionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ReportDefinitionSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDefinitionSpec.
func (in *ReportDefinitionSpec) DeepCopy() *ReportDefinitionSpec {
	if in == nil {
		return nil
	}
	out := new(ReportDefinitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDefinitionSpecResource) DeepCopyInto(out *ReportDefinitionSpecResource) {
	*out = *in
	if in.AdditionalArtifacts != nil {
		in, out := &in.AdditionalArtifacts, &out.AdditionalArtifacts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalSchemaElements != nil {
		in, out := &in.AdditionalSchemaElements, &out.AdditionalSchemaElements
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Compression != nil {
		in, out := &in.Compression, &out.Compression
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.RefreshClosedReports != nil {
		in, out := &in.RefreshClosedReports, &out.RefreshClosedReports
		*out = new(bool)
		**out = **in
	}
	if in.ReportName != nil {
		in, out := &in.ReportName, &out.ReportName
		*out = new(string)
		**out = **in
	}
	if in.ReportVersioning != nil {
		in, out := &in.ReportVersioning, &out.ReportVersioning
		*out = new(string)
		**out = **in
	}
	if in.S3Bucket != nil {
		in, out := &in.S3Bucket, &out.S3Bucket
		*out = new(string)
		**out = **in
	}
	if in.S3Prefix != nil {
		in, out := &in.S3Prefix, &out.S3Prefix
		*out = new(string)
		**out = **in
	}
	if in.S3Region != nil {
		in, out := &in.S3Region, &out.S3Region
		*out = new(string)
		**out = **in
	}
	if in.TimeUnit != nil {
		in, out := &in.TimeUnit, &out.TimeUnit
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDefinitionSpecResource.
func (in *ReportDefinitionSpecResource) DeepCopy() *ReportDefinitionSpecResource {
	if in == nil {
		return nil
	}
	out := new(ReportDefinitionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportDefinitionStatus) DeepCopyInto(out *ReportDefinitionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportDefinitionStatus.
func (in *ReportDefinitionStatus) DeepCopy() *ReportDefinitionStatus {
	if in == nil {
		return nil
	}
	out := new(ReportDefinitionStatus)
	in.DeepCopyInto(out)
	return out
}
