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
func (in *AccessPoint) DeepCopyInto(out *AccessPoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPoint.
func (in *AccessPoint) DeepCopy() *AccessPoint {
	if in == nil {
		return nil
	}
	out := new(AccessPoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointList) DeepCopyInto(out *AccessPointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccessPoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointList.
func (in *AccessPointList) DeepCopy() *AccessPointList {
	if in == nil {
		return nil
	}
	out := new(AccessPointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessPointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointSpec) DeepCopyInto(out *AccessPointSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AccessPointSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointSpec.
func (in *AccessPointSpec) DeepCopy() *AccessPointSpec {
	if in == nil {
		return nil
	}
	out := new(AccessPointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointSpecPosixUser) DeepCopyInto(out *AccessPointSpecPosixUser) {
	*out = *in
	if in.Gid != nil {
		in, out := &in.Gid, &out.Gid
		*out = new(int64)
		**out = **in
	}
	if in.SecondaryGids != nil {
		in, out := &in.SecondaryGids, &out.SecondaryGids
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.Uid != nil {
		in, out := &in.Uid, &out.Uid
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointSpecPosixUser.
func (in *AccessPointSpecPosixUser) DeepCopy() *AccessPointSpecPosixUser {
	if in == nil {
		return nil
	}
	out := new(AccessPointSpecPosixUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointSpecResource) DeepCopyInto(out *AccessPointSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.FileSystemArn != nil {
		in, out := &in.FileSystemArn, &out.FileSystemArn
		*out = new(string)
		**out = **in
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.PosixUser != nil {
		in, out := &in.PosixUser, &out.PosixUser
		*out = new(AccessPointSpecPosixUser)
		(*in).DeepCopyInto(*out)
	}
	if in.RootDirectory != nil {
		in, out := &in.RootDirectory, &out.RootDirectory
		*out = new(AccessPointSpecRootDirectory)
		(*in).DeepCopyInto(*out)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointSpecResource.
func (in *AccessPointSpecResource) DeepCopy() *AccessPointSpecResource {
	if in == nil {
		return nil
	}
	out := new(AccessPointSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointSpecRootDirectory) DeepCopyInto(out *AccessPointSpecRootDirectory) {
	*out = *in
	if in.CreationInfo != nil {
		in, out := &in.CreationInfo, &out.CreationInfo
		*out = new(AccessPointSpecRootDirectoryCreationInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointSpecRootDirectory.
func (in *AccessPointSpecRootDirectory) DeepCopy() *AccessPointSpecRootDirectory {
	if in == nil {
		return nil
	}
	out := new(AccessPointSpecRootDirectory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointSpecRootDirectoryCreationInfo) DeepCopyInto(out *AccessPointSpecRootDirectoryCreationInfo) {
	*out = *in
	if in.OwnerGid != nil {
		in, out := &in.OwnerGid, &out.OwnerGid
		*out = new(int64)
		**out = **in
	}
	if in.OwnerUid != nil {
		in, out := &in.OwnerUid, &out.OwnerUid
		*out = new(int64)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointSpecRootDirectoryCreationInfo.
func (in *AccessPointSpecRootDirectoryCreationInfo) DeepCopy() *AccessPointSpecRootDirectoryCreationInfo {
	if in == nil {
		return nil
	}
	out := new(AccessPointSpecRootDirectoryCreationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPointStatus) DeepCopyInto(out *AccessPointStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPointStatus.
func (in *AccessPointStatus) DeepCopy() *AccessPointStatus {
	if in == nil {
		return nil
	}
	out := new(AccessPointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicy) DeepCopyInto(out *BackupPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicy.
func (in *BackupPolicy) DeepCopy() *BackupPolicy {
	if in == nil {
		return nil
	}
	out := new(BackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyList) DeepCopyInto(out *BackupPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackupPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyList.
func (in *BackupPolicyList) DeepCopy() *BackupPolicyList {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackupPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicySpec) DeepCopyInto(out *BackupPolicySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BackupPolicySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicySpec.
func (in *BackupPolicySpec) DeepCopy() *BackupPolicySpec {
	if in == nil {
		return nil
	}
	out := new(BackupPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicySpecBackupPolicy) DeepCopyInto(out *BackupPolicySpecBackupPolicy) {
	*out = *in
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicySpecBackupPolicy.
func (in *BackupPolicySpecBackupPolicy) DeepCopy() *BackupPolicySpecBackupPolicy {
	if in == nil {
		return nil
	}
	out := new(BackupPolicySpecBackupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicySpecResource) DeepCopyInto(out *BackupPolicySpecResource) {
	*out = *in
	if in.BackupPolicy != nil {
		in, out := &in.BackupPolicy, &out.BackupPolicy
		*out = new(BackupPolicySpecBackupPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicySpecResource.
func (in *BackupPolicySpecResource) DeepCopy() *BackupPolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(BackupPolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPolicyStatus) DeepCopyInto(out *BackupPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPolicyStatus.
func (in *BackupPolicyStatus) DeepCopy() *BackupPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(BackupPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystem) DeepCopyInto(out *FileSystem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystem.
func (in *FileSystem) DeepCopy() *FileSystem {
	if in == nil {
		return nil
	}
	out := new(FileSystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemList) DeepCopyInto(out *FileSystemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FileSystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemList.
func (in *FileSystemList) DeepCopy() *FileSystemList {
	if in == nil {
		return nil
	}
	out := new(FileSystemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemPolicy) DeepCopyInto(out *FileSystemPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemPolicy.
func (in *FileSystemPolicy) DeepCopy() *FileSystemPolicy {
	if in == nil {
		return nil
	}
	out := new(FileSystemPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystemPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemPolicyList) DeepCopyInto(out *FileSystemPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FileSystemPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemPolicyList.
func (in *FileSystemPolicyList) DeepCopy() *FileSystemPolicyList {
	if in == nil {
		return nil
	}
	out := new(FileSystemPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FileSystemPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemPolicySpec) DeepCopyInto(out *FileSystemPolicySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(FileSystemPolicySpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemPolicySpec.
func (in *FileSystemPolicySpec) DeepCopy() *FileSystemPolicySpec {
	if in == nil {
		return nil
	}
	out := new(FileSystemPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemPolicySpecResource) DeepCopyInto(out *FileSystemPolicySpecResource) {
	*out = *in
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemPolicySpecResource.
func (in *FileSystemPolicySpecResource) DeepCopy() *FileSystemPolicySpecResource {
	if in == nil {
		return nil
	}
	out := new(FileSystemPolicySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemPolicyStatus) DeepCopyInto(out *FileSystemPolicyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemPolicyStatus.
func (in *FileSystemPolicyStatus) DeepCopy() *FileSystemPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(FileSystemPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemSpec) DeepCopyInto(out *FileSystemSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(FileSystemSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemSpec.
func (in *FileSystemSpec) DeepCopy() *FileSystemSpec {
	if in == nil {
		return nil
	}
	out := new(FileSystemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemSpecLifecyclePolicy) DeepCopyInto(out *FileSystemSpecLifecyclePolicy) {
	*out = *in
	if in.TransitionToIa != nil {
		in, out := &in.TransitionToIa, &out.TransitionToIa
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemSpecLifecyclePolicy.
func (in *FileSystemSpecLifecyclePolicy) DeepCopy() *FileSystemSpecLifecyclePolicy {
	if in == nil {
		return nil
	}
	out := new(FileSystemSpecLifecyclePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemSpecResource) DeepCopyInto(out *FileSystemSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZoneID != nil {
		in, out := &in.AvailabilityZoneID, &out.AvailabilityZoneID
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZoneName != nil {
		in, out := &in.AvailabilityZoneName, &out.AvailabilityZoneName
		*out = new(string)
		**out = **in
	}
	if in.CreationToken != nil {
		in, out := &in.CreationToken, &out.CreationToken
		*out = new(string)
		**out = **in
	}
	if in.DnsName != nil {
		in, out := &in.DnsName, &out.DnsName
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.LifecyclePolicy != nil {
		in, out := &in.LifecyclePolicy, &out.LifecyclePolicy
		*out = new(FileSystemSpecLifecyclePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.NumberOfMountTargets != nil {
		in, out := &in.NumberOfMountTargets, &out.NumberOfMountTargets
		*out = new(int64)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.PerformanceMode != nil {
		in, out := &in.PerformanceMode, &out.PerformanceMode
		*out = new(string)
		**out = **in
	}
	if in.ProvisionedThroughputInMibps != nil {
		in, out := &in.ProvisionedThroughputInMibps, &out.ProvisionedThroughputInMibps
		*out = new(float64)
		**out = **in
	}
	if in.SizeInBytes != nil {
		in, out := &in.SizeInBytes, &out.SizeInBytes
		*out = make([]FileSystemSpecSizeInBytes, len(*in))
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
	if in.ThroughputMode != nil {
		in, out := &in.ThroughputMode, &out.ThroughputMode
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemSpecResource.
func (in *FileSystemSpecResource) DeepCopy() *FileSystemSpecResource {
	if in == nil {
		return nil
	}
	out := new(FileSystemSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemSpecSizeInBytes) DeepCopyInto(out *FileSystemSpecSizeInBytes) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(int64)
		**out = **in
	}
	if in.ValueInIa != nil {
		in, out := &in.ValueInIa, &out.ValueInIa
		*out = new(int64)
		**out = **in
	}
	if in.ValueInStandard != nil {
		in, out := &in.ValueInStandard, &out.ValueInStandard
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemSpecSizeInBytes.
func (in *FileSystemSpecSizeInBytes) DeepCopy() *FileSystemSpecSizeInBytes {
	if in == nil {
		return nil
	}
	out := new(FileSystemSpecSizeInBytes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemStatus) DeepCopyInto(out *FileSystemStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemStatus.
func (in *FileSystemStatus) DeepCopy() *FileSystemStatus {
	if in == nil {
		return nil
	}
	out := new(FileSystemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTarget) DeepCopyInto(out *MountTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTarget.
func (in *MountTarget) DeepCopy() *MountTarget {
	if in == nil {
		return nil
	}
	out := new(MountTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MountTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetList) DeepCopyInto(out *MountTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MountTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetList.
func (in *MountTargetList) DeepCopy() *MountTargetList {
	if in == nil {
		return nil
	}
	out := new(MountTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MountTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetSpec) DeepCopyInto(out *MountTargetSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(MountTargetSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetSpec.
func (in *MountTargetSpec) DeepCopy() *MountTargetSpec {
	if in == nil {
		return nil
	}
	out := new(MountTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetSpecResource) DeepCopyInto(out *MountTargetSpecResource) {
	*out = *in
	if in.AvailabilityZoneID != nil {
		in, out := &in.AvailabilityZoneID, &out.AvailabilityZoneID
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZoneName != nil {
		in, out := &in.AvailabilityZoneName, &out.AvailabilityZoneName
		*out = new(string)
		**out = **in
	}
	if in.DnsName != nil {
		in, out := &in.DnsName, &out.DnsName
		*out = new(string)
		**out = **in
	}
	if in.FileSystemArn != nil {
		in, out := &in.FileSystemArn, &out.FileSystemArn
		*out = new(string)
		**out = **in
	}
	if in.FileSystemID != nil {
		in, out := &in.FileSystemID, &out.FileSystemID
		*out = new(string)
		**out = **in
	}
	if in.IpAddress != nil {
		in, out := &in.IpAddress, &out.IpAddress
		*out = new(string)
		**out = **in
	}
	if in.MountTargetDNSName != nil {
		in, out := &in.MountTargetDNSName, &out.MountTargetDNSName
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.OwnerID != nil {
		in, out := &in.OwnerID, &out.OwnerID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetSpecResource.
func (in *MountTargetSpecResource) DeepCopy() *MountTargetSpecResource {
	if in == nil {
		return nil
	}
	out := new(MountTargetSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MountTargetStatus) DeepCopyInto(out *MountTargetStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MountTargetStatus.
func (in *MountTargetStatus) DeepCopy() *MountTargetStatus {
	if in == nil {
		return nil
	}
	out := new(MountTargetStatus)
	in.DeepCopyInto(out)
	return out
}
