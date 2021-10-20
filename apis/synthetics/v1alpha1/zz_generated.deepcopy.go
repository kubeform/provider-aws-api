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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Canary) DeepCopyInto(out *Canary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Canary.
func (in *Canary) DeepCopy() *Canary {
	if in == nil {
		return nil
	}
	out := new(Canary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Canary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryList) DeepCopyInto(out *CanaryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Canary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryList.
func (in *CanaryList) DeepCopy() *CanaryList {
	if in == nil {
		return nil
	}
	out := new(CanaryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CanaryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpec) DeepCopyInto(out *CanarySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(CanarySpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpec.
func (in *CanarySpec) DeepCopy() *CanarySpec {
	if in == nil {
		return nil
	}
	out := new(CanarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpecResource) DeepCopyInto(out *CanarySpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ArtifactS3Location != nil {
		in, out := &in.ArtifactS3Location, &out.ArtifactS3Location
		*out = new(string)
		**out = **in
	}
	if in.EngineArn != nil {
		in, out := &in.EngineArn, &out.EngineArn
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleArn != nil {
		in, out := &in.ExecutionRoleArn, &out.ExecutionRoleArn
		*out = new(string)
		**out = **in
	}
	if in.FailureRetentionPeriod != nil {
		in, out := &in.FailureRetentionPeriod, &out.FailureRetentionPeriod
		*out = new(int64)
		**out = **in
	}
	if in.Handler != nil {
		in, out := &in.Handler, &out.Handler
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RunConfig != nil {
		in, out := &in.RunConfig, &out.RunConfig
		*out = new(CanarySpecRunConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.RuntimeVersion != nil {
		in, out := &in.RuntimeVersion, &out.RuntimeVersion
		*out = new(string)
		**out = **in
	}
	if in.S3Bucket != nil {
		in, out := &in.S3Bucket, &out.S3Bucket
		*out = new(string)
		**out = **in
	}
	if in.S3Key != nil {
		in, out := &in.S3Key, &out.S3Key
		*out = new(string)
		**out = **in
	}
	if in.S3Version != nil {
		in, out := &in.S3Version, &out.S3Version
		*out = new(string)
		**out = **in
	}
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(CanarySpecSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.SourceLocationArn != nil {
		in, out := &in.SourceLocationArn, &out.SourceLocationArn
		*out = new(string)
		**out = **in
	}
	if in.StartCanary != nil {
		in, out := &in.StartCanary, &out.StartCanary
		*out = new(bool)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.SuccessRetentionPeriod != nil {
		in, out := &in.SuccessRetentionPeriod, &out.SuccessRetentionPeriod
		*out = new(int64)
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
	if in.Timeline != nil {
		in, out := &in.Timeline, &out.Timeline
		*out = make([]CanarySpecTimeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VpcConfig != nil {
		in, out := &in.VpcConfig, &out.VpcConfig
		*out = new(CanarySpecVpcConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ZipFile != nil {
		in, out := &in.ZipFile, &out.ZipFile
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpecResource.
func (in *CanarySpecResource) DeepCopy() *CanarySpecResource {
	if in == nil {
		return nil
	}
	out := new(CanarySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpecRunConfig) DeepCopyInto(out *CanarySpecRunConfig) {
	*out = *in
	if in.ActiveTracing != nil {
		in, out := &in.ActiveTracing, &out.ActiveTracing
		*out = new(bool)
		**out = **in
	}
	if in.MemoryInMb != nil {
		in, out := &in.MemoryInMb, &out.MemoryInMb
		*out = new(int64)
		**out = **in
	}
	if in.TimeoutInSeconds != nil {
		in, out := &in.TimeoutInSeconds, &out.TimeoutInSeconds
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpecRunConfig.
func (in *CanarySpecRunConfig) DeepCopy() *CanarySpecRunConfig {
	if in == nil {
		return nil
	}
	out := new(CanarySpecRunConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpecSchedule) DeepCopyInto(out *CanarySpecSchedule) {
	*out = *in
	if in.DurationInSeconds != nil {
		in, out := &in.DurationInSeconds, &out.DurationInSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpecSchedule.
func (in *CanarySpecSchedule) DeepCopy() *CanarySpecSchedule {
	if in == nil {
		return nil
	}
	out := new(CanarySpecSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpecTimeline) DeepCopyInto(out *CanarySpecTimeline) {
	*out = *in
	if in.Created != nil {
		in, out := &in.Created, &out.Created
		*out = new(string)
		**out = **in
	}
	if in.LastModified != nil {
		in, out := &in.LastModified, &out.LastModified
		*out = new(string)
		**out = **in
	}
	if in.LastStarted != nil {
		in, out := &in.LastStarted, &out.LastStarted
		*out = new(string)
		**out = **in
	}
	if in.LastStopped != nil {
		in, out := &in.LastStopped, &out.LastStopped
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpecTimeline.
func (in *CanarySpecTimeline) DeepCopy() *CanarySpecTimeline {
	if in == nil {
		return nil
	}
	out := new(CanarySpecTimeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpecVpcConfig) DeepCopyInto(out *CanarySpecVpcConfig) {
	*out = *in
	if in.SecurityGroupIDS != nil {
		in, out := &in.SecurityGroupIDS, &out.SecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpecVpcConfig.
func (in *CanarySpecVpcConfig) DeepCopy() *CanarySpecVpcConfig {
	if in == nil {
		return nil
	}
	out := new(CanarySpecVpcConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStatus) DeepCopyInto(out *CanaryStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStatus.
func (in *CanaryStatus) DeepCopy() *CanaryStatus {
	if in == nil {
		return nil
	}
	out := new(CanaryStatus)
	in.DeepCopyInto(out)
	return out
}
