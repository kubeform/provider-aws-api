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
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationList) DeepCopyInto(out *ConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Configuration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationList.
func (in *ConfigurationList) DeepCopy() *ConfigurationList {
	if in == nil {
		return nil
	}
	out := new(ConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpec) DeepCopyInto(out *ConfigurationSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ConfigurationSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpec.
func (in *ConfigurationSpec) DeepCopy() *ConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecEbsBlockDevice) DeepCopyInto(out *ConfigurationSpecEbsBlockDevice) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(bool)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecEbsBlockDevice.
func (in *ConfigurationSpecEbsBlockDevice) DeepCopy() *ConfigurationSpecEbsBlockDevice {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecEbsBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecEphemeralBlockDevice) DeepCopyInto(out *ConfigurationSpecEphemeralBlockDevice) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecEphemeralBlockDevice.
func (in *ConfigurationSpecEphemeralBlockDevice) DeepCopy() *ConfigurationSpecEphemeralBlockDevice {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecEphemeralBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecMetadataOptions) DeepCopyInto(out *ConfigurationSpecMetadataOptions) {
	*out = *in
	if in.HttpEndpoint != nil {
		in, out := &in.HttpEndpoint, &out.HttpEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HttpPutResponseHopLimit != nil {
		in, out := &in.HttpPutResponseHopLimit, &out.HttpPutResponseHopLimit
		*out = new(int64)
		**out = **in
	}
	if in.HttpTokens != nil {
		in, out := &in.HttpTokens, &out.HttpTokens
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecMetadataOptions.
func (in *ConfigurationSpecMetadataOptions) DeepCopy() *ConfigurationSpecMetadataOptions {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecMetadataOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecResource) DeepCopyInto(out *ConfigurationSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AssociatePublicIPAddress != nil {
		in, out := &in.AssociatePublicIPAddress, &out.AssociatePublicIPAddress
		*out = new(bool)
		**out = **in
	}
	if in.EbsBlockDevice != nil {
		in, out := &in.EbsBlockDevice, &out.EbsBlockDevice
		*out = make([]ConfigurationSpecEbsBlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EbsOptimized != nil {
		in, out := &in.EbsOptimized, &out.EbsOptimized
		*out = new(bool)
		**out = **in
	}
	if in.EnableMonitoring != nil {
		in, out := &in.EnableMonitoring, &out.EnableMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.EphemeralBlockDevice != nil {
		in, out := &in.EphemeralBlockDevice, &out.EphemeralBlockDevice
		*out = make([]ConfigurationSpecEphemeralBlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IamInstanceProfile != nil {
		in, out := &in.IamInstanceProfile, &out.IamInstanceProfile
		*out = new(string)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.MetadataOptions != nil {
		in, out := &in.MetadataOptions, &out.MetadataOptions
		*out = new(ConfigurationSpecMetadataOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.PlacementTenancy != nil {
		in, out := &in.PlacementTenancy, &out.PlacementTenancy
		*out = new(string)
		**out = **in
	}
	if in.RootBlockDevice != nil {
		in, out := &in.RootBlockDevice, &out.RootBlockDevice
		*out = new(ConfigurationSpecRootBlockDevice)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SpotPrice != nil {
		in, out := &in.SpotPrice, &out.SpotPrice
		*out = new(string)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.UserDataBase64 != nil {
		in, out := &in.UserDataBase64, &out.UserDataBase64
		*out = new(string)
		**out = **in
	}
	if in.VpcClassicLinkID != nil {
		in, out := &in.VpcClassicLinkID, &out.VpcClassicLinkID
		*out = new(string)
		**out = **in
	}
	if in.VpcClassicLinkSecurityGroups != nil {
		in, out := &in.VpcClassicLinkSecurityGroups, &out.VpcClassicLinkSecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecResource.
func (in *ConfigurationSpecResource) DeepCopy() *ConfigurationSpecResource {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationSpecRootBlockDevice) DeepCopyInto(out *ConfigurationSpecRootBlockDevice) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationSpecRootBlockDevice.
func (in *ConfigurationSpecRootBlockDevice) DeepCopy() *ConfigurationSpecRootBlockDevice {
	if in == nil {
		return nil
	}
	out := new(ConfigurationSpecRootBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationStatus) DeepCopyInto(out *ConfigurationStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationStatus.
func (in *ConfigurationStatus) DeepCopy() *ConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(ConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Template) DeepCopyInto(out *Template) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Template.
func (in *Template) DeepCopy() *Template {
	if in == nil {
		return nil
	}
	out := new(Template)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Template) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateList) DeepCopyInto(out *TemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Template, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateList.
func (in *TemplateList) DeepCopy() *TemplateList {
	if in == nil {
		return nil
	}
	out := new(TemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpec) DeepCopyInto(out *TemplateSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(TemplateSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpec.
func (in *TemplateSpec) DeepCopy() *TemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecBlockDeviceMappings) DeepCopyInto(out *TemplateSpecBlockDeviceMappings) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.Ebs != nil {
		in, out := &in.Ebs, &out.Ebs
		*out = new(TemplateSpecBlockDeviceMappingsEbs)
		(*in).DeepCopyInto(*out)
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecBlockDeviceMappings.
func (in *TemplateSpecBlockDeviceMappings) DeepCopy() *TemplateSpecBlockDeviceMappings {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecBlockDeviceMappings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecBlockDeviceMappingsEbs) DeepCopyInto(out *TemplateSpecBlockDeviceMappingsEbs) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(string)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(string)
		**out = **in
	}
	if in.Iops != nil {
		in, out := &in.Iops, &out.Iops
		*out = new(int64)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.Throughput != nil {
		in, out := &in.Throughput, &out.Throughput
		*out = new(int64)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecBlockDeviceMappingsEbs.
func (in *TemplateSpecBlockDeviceMappingsEbs) DeepCopy() *TemplateSpecBlockDeviceMappingsEbs {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecBlockDeviceMappingsEbs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecCapacityReservationSpecification) DeepCopyInto(out *TemplateSpecCapacityReservationSpecification) {
	*out = *in
	if in.CapacityReservationPreference != nil {
		in, out := &in.CapacityReservationPreference, &out.CapacityReservationPreference
		*out = new(string)
		**out = **in
	}
	if in.CapacityReservationTarget != nil {
		in, out := &in.CapacityReservationTarget, &out.CapacityReservationTarget
		*out = new(TemplateSpecCapacityReservationSpecificationCapacityReservationTarget)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecCapacityReservationSpecification.
func (in *TemplateSpecCapacityReservationSpecification) DeepCopy() *TemplateSpecCapacityReservationSpecification {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecCapacityReservationSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecCapacityReservationSpecificationCapacityReservationTarget) DeepCopyInto(out *TemplateSpecCapacityReservationSpecificationCapacityReservationTarget) {
	*out = *in
	if in.CapacityReservationID != nil {
		in, out := &in.CapacityReservationID, &out.CapacityReservationID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecCapacityReservationSpecificationCapacityReservationTarget.
func (in *TemplateSpecCapacityReservationSpecificationCapacityReservationTarget) DeepCopy() *TemplateSpecCapacityReservationSpecificationCapacityReservationTarget {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecCapacityReservationSpecificationCapacityReservationTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecCpuOptions) DeepCopyInto(out *TemplateSpecCpuOptions) {
	*out = *in
	if in.CoreCount != nil {
		in, out := &in.CoreCount, &out.CoreCount
		*out = new(int64)
		**out = **in
	}
	if in.ThreadsPerCore != nil {
		in, out := &in.ThreadsPerCore, &out.ThreadsPerCore
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecCpuOptions.
func (in *TemplateSpecCpuOptions) DeepCopy() *TemplateSpecCpuOptions {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecCpuOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecCreditSpecification) DeepCopyInto(out *TemplateSpecCreditSpecification) {
	*out = *in
	if in.CpuCredits != nil {
		in, out := &in.CpuCredits, &out.CpuCredits
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecCreditSpecification.
func (in *TemplateSpecCreditSpecification) DeepCopy() *TemplateSpecCreditSpecification {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecCreditSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecElasticGpuSpecifications) DeepCopyInto(out *TemplateSpecElasticGpuSpecifications) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecElasticGpuSpecifications.
func (in *TemplateSpecElasticGpuSpecifications) DeepCopy() *TemplateSpecElasticGpuSpecifications {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecElasticGpuSpecifications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecElasticInferenceAccelerator) DeepCopyInto(out *TemplateSpecElasticInferenceAccelerator) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecElasticInferenceAccelerator.
func (in *TemplateSpecElasticInferenceAccelerator) DeepCopy() *TemplateSpecElasticInferenceAccelerator {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecElasticInferenceAccelerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecEnclaveOptions) DeepCopyInto(out *TemplateSpecEnclaveOptions) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecEnclaveOptions.
func (in *TemplateSpecEnclaveOptions) DeepCopy() *TemplateSpecEnclaveOptions {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecEnclaveOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecHibernationOptions) DeepCopyInto(out *TemplateSpecHibernationOptions) {
	*out = *in
	if in.Configured != nil {
		in, out := &in.Configured, &out.Configured
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecHibernationOptions.
func (in *TemplateSpecHibernationOptions) DeepCopy() *TemplateSpecHibernationOptions {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecHibernationOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecIamInstanceProfile) DeepCopyInto(out *TemplateSpecIamInstanceProfile) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecIamInstanceProfile.
func (in *TemplateSpecIamInstanceProfile) DeepCopy() *TemplateSpecIamInstanceProfile {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecIamInstanceProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecInstanceMarketOptions) DeepCopyInto(out *TemplateSpecInstanceMarketOptions) {
	*out = *in
	if in.MarketType != nil {
		in, out := &in.MarketType, &out.MarketType
		*out = new(string)
		**out = **in
	}
	if in.SpotOptions != nil {
		in, out := &in.SpotOptions, &out.SpotOptions
		*out = new(TemplateSpecInstanceMarketOptionsSpotOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecInstanceMarketOptions.
func (in *TemplateSpecInstanceMarketOptions) DeepCopy() *TemplateSpecInstanceMarketOptions {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecInstanceMarketOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecInstanceMarketOptionsSpotOptions) DeepCopyInto(out *TemplateSpecInstanceMarketOptionsSpotOptions) {
	*out = *in
	if in.BlockDurationMinutes != nil {
		in, out := &in.BlockDurationMinutes, &out.BlockDurationMinutes
		*out = new(int64)
		**out = **in
	}
	if in.InstanceInterruptionBehavior != nil {
		in, out := &in.InstanceInterruptionBehavior, &out.InstanceInterruptionBehavior
		*out = new(string)
		**out = **in
	}
	if in.MaxPrice != nil {
		in, out := &in.MaxPrice, &out.MaxPrice
		*out = new(string)
		**out = **in
	}
	if in.SpotInstanceType != nil {
		in, out := &in.SpotInstanceType, &out.SpotInstanceType
		*out = new(string)
		**out = **in
	}
	if in.ValidUntil != nil {
		in, out := &in.ValidUntil, &out.ValidUntil
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecInstanceMarketOptionsSpotOptions.
func (in *TemplateSpecInstanceMarketOptionsSpotOptions) DeepCopy() *TemplateSpecInstanceMarketOptionsSpotOptions {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecInstanceMarketOptionsSpotOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecLicenseSpecification) DeepCopyInto(out *TemplateSpecLicenseSpecification) {
	*out = *in
	if in.LicenseConfigurationArn != nil {
		in, out := &in.LicenseConfigurationArn, &out.LicenseConfigurationArn
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecLicenseSpecification.
func (in *TemplateSpecLicenseSpecification) DeepCopy() *TemplateSpecLicenseSpecification {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecLicenseSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecMetadataOptions) DeepCopyInto(out *TemplateSpecMetadataOptions) {
	*out = *in
	if in.HttpEndpoint != nil {
		in, out := &in.HttpEndpoint, &out.HttpEndpoint
		*out = new(string)
		**out = **in
	}
	if in.HttpPutResponseHopLimit != nil {
		in, out := &in.HttpPutResponseHopLimit, &out.HttpPutResponseHopLimit
		*out = new(int64)
		**out = **in
	}
	if in.HttpTokens != nil {
		in, out := &in.HttpTokens, &out.HttpTokens
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecMetadataOptions.
func (in *TemplateSpecMetadataOptions) DeepCopy() *TemplateSpecMetadataOptions {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecMetadataOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecMonitoring) DeepCopyInto(out *TemplateSpecMonitoring) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecMonitoring.
func (in *TemplateSpecMonitoring) DeepCopy() *TemplateSpecMonitoring {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecMonitoring)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecNetworkInterfaces) DeepCopyInto(out *TemplateSpecNetworkInterfaces) {
	*out = *in
	if in.AssociateCarrierIPAddress != nil {
		in, out := &in.AssociateCarrierIPAddress, &out.AssociateCarrierIPAddress
		*out = new(string)
		**out = **in
	}
	if in.AssociatePublicIPAddress != nil {
		in, out := &in.AssociatePublicIPAddress, &out.AssociatePublicIPAddress
		*out = new(string)
		**out = **in
	}
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DeviceIndex != nil {
		in, out := &in.DeviceIndex, &out.DeviceIndex
		*out = new(int64)
		**out = **in
	}
	if in.InterfaceType != nil {
		in, out := &in.InterfaceType, &out.InterfaceType
		*out = new(string)
		**out = **in
	}
	if in.Ipv4AddressCount != nil {
		in, out := &in.Ipv4AddressCount, &out.Ipv4AddressCount
		*out = new(int64)
		**out = **in
	}
	if in.Ipv4Addresses != nil {
		in, out := &in.Ipv4Addresses, &out.Ipv4Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ipv6AddressCount != nil {
		in, out := &in.Ipv6AddressCount, &out.Ipv6AddressCount
		*out = new(int64)
		**out = **in
	}
	if in.Ipv6Addresses != nil {
		in, out := &in.Ipv6Addresses, &out.Ipv6Addresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecNetworkInterfaces.
func (in *TemplateSpecNetworkInterfaces) DeepCopy() *TemplateSpecNetworkInterfaces {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecNetworkInterfaces)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecPlacement) DeepCopyInto(out *TemplateSpecPlacement) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.GroupName != nil {
		in, out := &in.GroupName, &out.GroupName
		*out = new(string)
		**out = **in
	}
	if in.HostID != nil {
		in, out := &in.HostID, &out.HostID
		*out = new(string)
		**out = **in
	}
	if in.HostResourceGroupArn != nil {
		in, out := &in.HostResourceGroupArn, &out.HostResourceGroupArn
		*out = new(string)
		**out = **in
	}
	if in.PartitionNumber != nil {
		in, out := &in.PartitionNumber, &out.PartitionNumber
		*out = new(int64)
		**out = **in
	}
	if in.SpreadDomain != nil {
		in, out := &in.SpreadDomain, &out.SpreadDomain
		*out = new(string)
		**out = **in
	}
	if in.Tenancy != nil {
		in, out := &in.Tenancy, &out.Tenancy
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecPlacement.
func (in *TemplateSpecPlacement) DeepCopy() *TemplateSpecPlacement {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecPlacement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecResource) DeepCopyInto(out *TemplateSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.BlockDeviceMappings != nil {
		in, out := &in.BlockDeviceMappings, &out.BlockDeviceMappings
		*out = make([]TemplateSpecBlockDeviceMappings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CapacityReservationSpecification != nil {
		in, out := &in.CapacityReservationSpecification, &out.CapacityReservationSpecification
		*out = new(TemplateSpecCapacityReservationSpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.CpuOptions != nil {
		in, out := &in.CpuOptions, &out.CpuOptions
		*out = new(TemplateSpecCpuOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.CreditSpecification != nil {
		in, out := &in.CreditSpecification, &out.CreditSpecification
		*out = new(TemplateSpecCreditSpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultVersion != nil {
		in, out := &in.DefaultVersion, &out.DefaultVersion
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisableAPITermination != nil {
		in, out := &in.DisableAPITermination, &out.DisableAPITermination
		*out = new(bool)
		**out = **in
	}
	if in.EbsOptimized != nil {
		in, out := &in.EbsOptimized, &out.EbsOptimized
		*out = new(string)
		**out = **in
	}
	if in.ElasticGpuSpecifications != nil {
		in, out := &in.ElasticGpuSpecifications, &out.ElasticGpuSpecifications
		*out = make([]TemplateSpecElasticGpuSpecifications, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ElasticInferenceAccelerator != nil {
		in, out := &in.ElasticInferenceAccelerator, &out.ElasticInferenceAccelerator
		*out = new(TemplateSpecElasticInferenceAccelerator)
		(*in).DeepCopyInto(*out)
	}
	if in.EnclaveOptions != nil {
		in, out := &in.EnclaveOptions, &out.EnclaveOptions
		*out = new(TemplateSpecEnclaveOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.HibernationOptions != nil {
		in, out := &in.HibernationOptions, &out.HibernationOptions
		*out = new(TemplateSpecHibernationOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.IamInstanceProfile != nil {
		in, out := &in.IamInstanceProfile, &out.IamInstanceProfile
		*out = new(TemplateSpecIamInstanceProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.InstanceInitiatedShutdownBehavior != nil {
		in, out := &in.InstanceInitiatedShutdownBehavior, &out.InstanceInitiatedShutdownBehavior
		*out = new(string)
		**out = **in
	}
	if in.InstanceMarketOptions != nil {
		in, out := &in.InstanceMarketOptions, &out.InstanceMarketOptions
		*out = new(TemplateSpecInstanceMarketOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.KernelID != nil {
		in, out := &in.KernelID, &out.KernelID
		*out = new(string)
		**out = **in
	}
	if in.KeyName != nil {
		in, out := &in.KeyName, &out.KeyName
		*out = new(string)
		**out = **in
	}
	if in.LatestVersion != nil {
		in, out := &in.LatestVersion, &out.LatestVersion
		*out = new(int64)
		**out = **in
	}
	if in.LicenseSpecification != nil {
		in, out := &in.LicenseSpecification, &out.LicenseSpecification
		*out = make([]TemplateSpecLicenseSpecification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MetadataOptions != nil {
		in, out := &in.MetadataOptions, &out.MetadataOptions
		*out = new(TemplateSpecMetadataOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(TemplateSpecMonitoring)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NamePrefix != nil {
		in, out := &in.NamePrefix, &out.NamePrefix
		*out = new(string)
		**out = **in
	}
	if in.NetworkInterfaces != nil {
		in, out := &in.NetworkInterfaces, &out.NetworkInterfaces
		*out = make([]TemplateSpecNetworkInterfaces, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = new(TemplateSpecPlacement)
		(*in).DeepCopyInto(*out)
	}
	if in.RamDiskID != nil {
		in, out := &in.RamDiskID, &out.RamDiskID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupNames != nil {
		in, out := &in.SecurityGroupNames, &out.SecurityGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TagSpecifications != nil {
		in, out := &in.TagSpecifications, &out.TagSpecifications
		*out = make([]TemplateSpecTagSpecifications, len(*in))
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
	if in.UpdateDefaultVersion != nil {
		in, out := &in.UpdateDefaultVersion, &out.UpdateDefaultVersion
		*out = new(bool)
		**out = **in
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
	if in.VpcSecurityGroupIDS != nil {
		in, out := &in.VpcSecurityGroupIDS, &out.VpcSecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecResource.
func (in *TemplateSpecResource) DeepCopy() *TemplateSpecResource {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpecTagSpecifications) DeepCopyInto(out *TemplateSpecTagSpecifications) {
	*out = *in
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpecTagSpecifications.
func (in *TemplateSpecTagSpecifications) DeepCopy() *TemplateSpecTagSpecifications {
	if in == nil {
		return nil
	}
	out := new(TemplateSpecTagSpecifications)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateStatus) DeepCopyInto(out *TemplateStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateStatus.
func (in *TemplateStatus) DeepCopy() *TemplateStatus {
	if in == nil {
		return nil
	}
	out := new(TemplateStatus)
	in.DeepCopyInto(out)
	return out
}
