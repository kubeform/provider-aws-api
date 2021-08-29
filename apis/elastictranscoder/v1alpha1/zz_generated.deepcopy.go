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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineList) DeepCopyInto(out *PipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineList.
func (in *PipelineList) DeepCopy() *PipelineList {
	if in == nil {
		return nil
	}
	out := new(PipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PipelineSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecContentConfig) DeepCopyInto(out *PipelineSpecContentConfig) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecContentConfig.
func (in *PipelineSpecContentConfig) DeepCopy() *PipelineSpecContentConfig {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecContentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecContentConfigPermissions) DeepCopyInto(out *PipelineSpecContentConfigPermissions) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Grantee != nil {
		in, out := &in.Grantee, &out.Grantee
		*out = new(string)
		**out = **in
	}
	if in.GranteeType != nil {
		in, out := &in.GranteeType, &out.GranteeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecContentConfigPermissions.
func (in *PipelineSpecContentConfigPermissions) DeepCopy() *PipelineSpecContentConfigPermissions {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecContentConfigPermissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecNotifications) DeepCopyInto(out *PipelineSpecNotifications) {
	*out = *in
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = new(string)
		**out = **in
	}
	if in.Error != nil {
		in, out := &in.Error, &out.Error
		*out = new(string)
		**out = **in
	}
	if in.Progressing != nil {
		in, out := &in.Progressing, &out.Progressing
		*out = new(string)
		**out = **in
	}
	if in.Warning != nil {
		in, out := &in.Warning, &out.Warning
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecNotifications.
func (in *PipelineSpecNotifications) DeepCopy() *PipelineSpecNotifications {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecNotifications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecResource) DeepCopyInto(out *PipelineSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AwsKmsKeyArn != nil {
		in, out := &in.AwsKmsKeyArn, &out.AwsKmsKeyArn
		*out = new(string)
		**out = **in
	}
	if in.ContentConfig != nil {
		in, out := &in.ContentConfig, &out.ContentConfig
		*out = new(PipelineSpecContentConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ContentConfigPermissions != nil {
		in, out := &in.ContentConfigPermissions, &out.ContentConfigPermissions
		*out = make([]PipelineSpecContentConfigPermissions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InputBucket != nil {
		in, out := &in.InputBucket, &out.InputBucket
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Notifications != nil {
		in, out := &in.Notifications, &out.Notifications
		*out = new(PipelineSpecNotifications)
		(*in).DeepCopyInto(*out)
	}
	if in.OutputBucket != nil {
		in, out := &in.OutputBucket, &out.OutputBucket
		*out = new(string)
		**out = **in
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.ThumbnailConfig != nil {
		in, out := &in.ThumbnailConfig, &out.ThumbnailConfig
		*out = new(PipelineSpecThumbnailConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ThumbnailConfigPermissions != nil {
		in, out := &in.ThumbnailConfigPermissions, &out.ThumbnailConfigPermissions
		*out = make([]PipelineSpecThumbnailConfigPermissions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecResource.
func (in *PipelineSpecResource) DeepCopy() *PipelineSpecResource {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecThumbnailConfig) DeepCopyInto(out *PipelineSpecThumbnailConfig) {
	*out = *in
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.StorageClass != nil {
		in, out := &in.StorageClass, &out.StorageClass
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecThumbnailConfig.
func (in *PipelineSpecThumbnailConfig) DeepCopy() *PipelineSpecThumbnailConfig {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecThumbnailConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecThumbnailConfigPermissions) DeepCopyInto(out *PipelineSpecThumbnailConfigPermissions) {
	*out = *in
	if in.Access != nil {
		in, out := &in.Access, &out.Access
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Grantee != nil {
		in, out := &in.Grantee, &out.Grantee
		*out = new(string)
		**out = **in
	}
	if in.GranteeType != nil {
		in, out := &in.GranteeType, &out.GranteeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecThumbnailConfigPermissions.
func (in *PipelineSpecThumbnailConfigPermissions) DeepCopy() *PipelineSpecThumbnailConfigPermissions {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecThumbnailConfigPermissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preset) DeepCopyInto(out *Preset) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preset.
func (in *Preset) DeepCopy() *Preset {
	if in == nil {
		return nil
	}
	out := new(Preset)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Preset) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetList) DeepCopyInto(out *PresetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Preset, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetList.
func (in *PresetList) DeepCopy() *PresetList {
	if in == nil {
		return nil
	}
	out := new(PresetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PresetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetSpec) DeepCopyInto(out *PresetSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PresetSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetSpec.
func (in *PresetSpec) DeepCopy() *PresetSpec {
	if in == nil {
		return nil
	}
	out := new(PresetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetSpecAudio) DeepCopyInto(out *PresetSpecAudio) {
	*out = *in
	if in.AudioPackingMode != nil {
		in, out := &in.AudioPackingMode, &out.AudioPackingMode
		*out = new(string)
		**out = **in
	}
	if in.BitRate != nil {
		in, out := &in.BitRate, &out.BitRate
		*out = new(string)
		**out = **in
	}
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = new(string)
		**out = **in
	}
	if in.Codec != nil {
		in, out := &in.Codec, &out.Codec
		*out = new(string)
		**out = **in
	}
	if in.SampleRate != nil {
		in, out := &in.SampleRate, &out.SampleRate
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetSpecAudio.
func (in *PresetSpecAudio) DeepCopy() *PresetSpecAudio {
	if in == nil {
		return nil
	}
	out := new(PresetSpecAudio)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetSpecAudioCodecOptions) DeepCopyInto(out *PresetSpecAudioCodecOptions) {
	*out = *in
	if in.BitDepth != nil {
		in, out := &in.BitDepth, &out.BitDepth
		*out = new(string)
		**out = **in
	}
	if in.BitOrder != nil {
		in, out := &in.BitOrder, &out.BitOrder
		*out = new(string)
		**out = **in
	}
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = new(string)
		**out = **in
	}
	if in.Signed != nil {
		in, out := &in.Signed, &out.Signed
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetSpecAudioCodecOptions.
func (in *PresetSpecAudioCodecOptions) DeepCopy() *PresetSpecAudioCodecOptions {
	if in == nil {
		return nil
	}
	out := new(PresetSpecAudioCodecOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetSpecResource) DeepCopyInto(out *PresetSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Audio != nil {
		in, out := &in.Audio, &out.Audio
		*out = new(PresetSpecAudio)
		(*in).DeepCopyInto(*out)
	}
	if in.AudioCodecOptions != nil {
		in, out := &in.AudioCodecOptions, &out.AudioCodecOptions
		*out = new(PresetSpecAudioCodecOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Container != nil {
		in, out := &in.Container, &out.Container
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
	if in.Thumbnails != nil {
		in, out := &in.Thumbnails, &out.Thumbnails
		*out = new(PresetSpecThumbnails)
		(*in).DeepCopyInto(*out)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Video != nil {
		in, out := &in.Video, &out.Video
		*out = new(PresetSpecVideo)
		(*in).DeepCopyInto(*out)
	}
	if in.VideoCodecOptions != nil {
		in, out := &in.VideoCodecOptions, &out.VideoCodecOptions
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.VideoWatermarks != nil {
		in, out := &in.VideoWatermarks, &out.VideoWatermarks
		*out = make([]PresetSpecVideoWatermarks, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetSpecResource.
func (in *PresetSpecResource) DeepCopy() *PresetSpecResource {
	if in == nil {
		return nil
	}
	out := new(PresetSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetSpecThumbnails) DeepCopyInto(out *PresetSpecThumbnails) {
	*out = *in
	if in.AspectRatio != nil {
		in, out := &in.AspectRatio, &out.AspectRatio
		*out = new(string)
		**out = **in
	}
	if in.Format != nil {
		in, out := &in.Format, &out.Format
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.MaxHeight != nil {
		in, out := &in.MaxHeight, &out.MaxHeight
		*out = new(string)
		**out = **in
	}
	if in.MaxWidth != nil {
		in, out := &in.MaxWidth, &out.MaxWidth
		*out = new(string)
		**out = **in
	}
	if in.PaddingPolicy != nil {
		in, out := &in.PaddingPolicy, &out.PaddingPolicy
		*out = new(string)
		**out = **in
	}
	if in.Resolution != nil {
		in, out := &in.Resolution, &out.Resolution
		*out = new(string)
		**out = **in
	}
	if in.SizingPolicy != nil {
		in, out := &in.SizingPolicy, &out.SizingPolicy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetSpecThumbnails.
func (in *PresetSpecThumbnails) DeepCopy() *PresetSpecThumbnails {
	if in == nil {
		return nil
	}
	out := new(PresetSpecThumbnails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetSpecVideo) DeepCopyInto(out *PresetSpecVideo) {
	*out = *in
	if in.AspectRatio != nil {
		in, out := &in.AspectRatio, &out.AspectRatio
		*out = new(string)
		**out = **in
	}
	if in.BitRate != nil {
		in, out := &in.BitRate, &out.BitRate
		*out = new(string)
		**out = **in
	}
	if in.Codec != nil {
		in, out := &in.Codec, &out.Codec
		*out = new(string)
		**out = **in
	}
	if in.DisplayAspectRatio != nil {
		in, out := &in.DisplayAspectRatio, &out.DisplayAspectRatio
		*out = new(string)
		**out = **in
	}
	if in.FixedGop != nil {
		in, out := &in.FixedGop, &out.FixedGop
		*out = new(string)
		**out = **in
	}
	if in.FrameRate != nil {
		in, out := &in.FrameRate, &out.FrameRate
		*out = new(string)
		**out = **in
	}
	if in.KeyframesMaxDist != nil {
		in, out := &in.KeyframesMaxDist, &out.KeyframesMaxDist
		*out = new(string)
		**out = **in
	}
	if in.MaxFrameRate != nil {
		in, out := &in.MaxFrameRate, &out.MaxFrameRate
		*out = new(string)
		**out = **in
	}
	if in.MaxHeight != nil {
		in, out := &in.MaxHeight, &out.MaxHeight
		*out = new(string)
		**out = **in
	}
	if in.MaxWidth != nil {
		in, out := &in.MaxWidth, &out.MaxWidth
		*out = new(string)
		**out = **in
	}
	if in.PaddingPolicy != nil {
		in, out := &in.PaddingPolicy, &out.PaddingPolicy
		*out = new(string)
		**out = **in
	}
	if in.Resolution != nil {
		in, out := &in.Resolution, &out.Resolution
		*out = new(string)
		**out = **in
	}
	if in.SizingPolicy != nil {
		in, out := &in.SizingPolicy, &out.SizingPolicy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetSpecVideo.
func (in *PresetSpecVideo) DeepCopy() *PresetSpecVideo {
	if in == nil {
		return nil
	}
	out := new(PresetSpecVideo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetSpecVideoWatermarks) DeepCopyInto(out *PresetSpecVideoWatermarks) {
	*out = *in
	if in.HorizontalAlign != nil {
		in, out := &in.HorizontalAlign, &out.HorizontalAlign
		*out = new(string)
		**out = **in
	}
	if in.HorizontalOffset != nil {
		in, out := &in.HorizontalOffset, &out.HorizontalOffset
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.MaxHeight != nil {
		in, out := &in.MaxHeight, &out.MaxHeight
		*out = new(string)
		**out = **in
	}
	if in.MaxWidth != nil {
		in, out := &in.MaxWidth, &out.MaxWidth
		*out = new(string)
		**out = **in
	}
	if in.Opacity != nil {
		in, out := &in.Opacity, &out.Opacity
		*out = new(string)
		**out = **in
	}
	if in.SizingPolicy != nil {
		in, out := &in.SizingPolicy, &out.SizingPolicy
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	if in.VerticalAlign != nil {
		in, out := &in.VerticalAlign, &out.VerticalAlign
		*out = new(string)
		**out = **in
	}
	if in.VerticalOffset != nil {
		in, out := &in.VerticalOffset, &out.VerticalOffset
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetSpecVideoWatermarks.
func (in *PresetSpecVideoWatermarks) DeepCopy() *PresetSpecVideoWatermarks {
	if in == nil {
		return nil
	}
	out := new(PresetSpecVideoWatermarks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PresetStatus) DeepCopyInto(out *PresetStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PresetStatus.
func (in *PresetStatus) DeepCopy() *PresetStatus {
	if in == nil {
		return nil
	}
	out := new(PresetStatus)
	in.DeepCopyInto(out)
	return out
}
