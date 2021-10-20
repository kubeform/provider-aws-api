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
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ClusterSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecCacheNodes) DeepCopyInto(out *ClusterSpecCacheNodes) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecCacheNodes.
func (in *ClusterSpecCacheNodes) DeepCopy() *ClusterSpecCacheNodes {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecCacheNodes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecResource) DeepCopyInto(out *ClusterSpecResource) {
	*out = *in
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.AzMode != nil {
		in, out := &in.AzMode, &out.AzMode
		*out = new(string)
		**out = **in
	}
	if in.CacheNodes != nil {
		in, out := &in.CacheNodes, &out.CacheNodes
		*out = make([]ClusterSpecCacheNodes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterAddress != nil {
		in, out := &in.ClusterAddress, &out.ClusterAddress
		*out = new(string)
		**out = **in
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.ConfigurationEndpoint != nil {
		in, out := &in.ConfigurationEndpoint, &out.ConfigurationEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.EngineVersionActual != nil {
		in, out := &in.EngineVersionActual, &out.EngineVersionActual
		*out = new(string)
		**out = **in
	}
	if in.FinalSnapshotIdentifier != nil {
		in, out := &in.FinalSnapshotIdentifier, &out.FinalSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.NotificationTopicArn != nil {
		in, out := &in.NotificationTopicArn, &out.NotificationTopicArn
		*out = new(string)
		**out = **in
	}
	if in.NumCacheNodes != nil {
		in, out := &in.NumCacheNodes, &out.NumCacheNodes
		*out = new(int64)
		**out = **in
	}
	if in.ParameterGroupName != nil {
		in, out := &in.ParameterGroupName, &out.ParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.PreferredAvailabilityZones != nil {
		in, out := &in.PreferredAvailabilityZones, &out.PreferredAvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReplicationGroupID != nil {
		in, out := &in.ReplicationGroupID, &out.ReplicationGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIDS != nil {
		in, out := &in.SecurityGroupIDS, &out.SecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupNames != nil {
		in, out := &in.SecurityGroupNames, &out.SecurityGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotArns != nil {
		in, out := &in.SnapshotArns, &out.SnapshotArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotName != nil {
		in, out := &in.SnapshotName, &out.SnapshotName
		*out = new(string)
		**out = **in
	}
	if in.SnapshotRetentionLimit != nil {
		in, out := &in.SnapshotRetentionLimit, &out.SnapshotRetentionLimit
		*out = new(int64)
		**out = **in
	}
	if in.SnapshotWindow != nil {
		in, out := &in.SnapshotWindow, &out.SnapshotWindow
		*out = new(string)
		**out = **in
	}
	if in.SubnetGroupName != nil {
		in, out := &in.SubnetGroupName, &out.SubnetGroupName
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecResource.
func (in *ClusterSpecResource) DeepCopy() *ClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalReplicationGroup) DeepCopyInto(out *GlobalReplicationGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalReplicationGroup.
func (in *GlobalReplicationGroup) DeepCopy() *GlobalReplicationGroup {
	if in == nil {
		return nil
	}
	out := new(GlobalReplicationGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalReplicationGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalReplicationGroupList) DeepCopyInto(out *GlobalReplicationGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GlobalReplicationGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalReplicationGroupList.
func (in *GlobalReplicationGroupList) DeepCopy() *GlobalReplicationGroupList {
	if in == nil {
		return nil
	}
	out := new(GlobalReplicationGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GlobalReplicationGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalReplicationGroupSpec) DeepCopyInto(out *GlobalReplicationGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(GlobalReplicationGroupSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalReplicationGroupSpec.
func (in *GlobalReplicationGroupSpec) DeepCopy() *GlobalReplicationGroupSpec {
	if in == nil {
		return nil
	}
	out := new(GlobalReplicationGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalReplicationGroupSpecResource) DeepCopyInto(out *GlobalReplicationGroupSpecResource) {
	*out = *in
	if in.ActualEngineVersion != nil {
		in, out := &in.ActualEngineVersion, &out.ActualEngineVersion
		*out = new(string)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AtRestEncryptionEnabled != nil {
		in, out := &in.AtRestEncryptionEnabled, &out.AtRestEncryptionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthTokenEnabled != nil {
		in, out := &in.AuthTokenEnabled, &out.AuthTokenEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CacheNodeType != nil {
		in, out := &in.CacheNodeType, &out.CacheNodeType
		*out = new(string)
		**out = **in
	}
	if in.ClusterEnabled != nil {
		in, out := &in.ClusterEnabled, &out.ClusterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersionActual != nil {
		in, out := &in.EngineVersionActual, &out.EngineVersionActual
		*out = new(string)
		**out = **in
	}
	if in.GlobalReplicationGroupDescription != nil {
		in, out := &in.GlobalReplicationGroupDescription, &out.GlobalReplicationGroupDescription
		*out = new(string)
		**out = **in
	}
	if in.GlobalReplicationGroupID != nil {
		in, out := &in.GlobalReplicationGroupID, &out.GlobalReplicationGroupID
		*out = new(string)
		**out = **in
	}
	if in.GlobalReplicationGroupIDSuffix != nil {
		in, out := &in.GlobalReplicationGroupIDSuffix, &out.GlobalReplicationGroupIDSuffix
		*out = new(string)
		**out = **in
	}
	if in.PrimaryReplicationGroupID != nil {
		in, out := &in.PrimaryReplicationGroupID, &out.PrimaryReplicationGroupID
		*out = new(string)
		**out = **in
	}
	if in.TransitEncryptionEnabled != nil {
		in, out := &in.TransitEncryptionEnabled, &out.TransitEncryptionEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalReplicationGroupSpecResource.
func (in *GlobalReplicationGroupSpecResource) DeepCopy() *GlobalReplicationGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(GlobalReplicationGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalReplicationGroupStatus) DeepCopyInto(out *GlobalReplicationGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalReplicationGroupStatus.
func (in *GlobalReplicationGroupStatus) DeepCopy() *GlobalReplicationGroupStatus {
	if in == nil {
		return nil
	}
	out := new(GlobalReplicationGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroup) DeepCopyInto(out *ParameterGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroup.
func (in *ParameterGroup) DeepCopy() *ParameterGroup {
	if in == nil {
		return nil
	}
	out := new(ParameterGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParameterGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupList) DeepCopyInto(out *ParameterGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ParameterGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupList.
func (in *ParameterGroupList) DeepCopy() *ParameterGroupList {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ParameterGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupSpec) DeepCopyInto(out *ParameterGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ParameterGroupSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupSpec.
func (in *ParameterGroupSpec) DeepCopy() *ParameterGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupSpecParameter) DeepCopyInto(out *ParameterGroupSpecParameter) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupSpecParameter.
func (in *ParameterGroupSpecParameter) DeepCopy() *ParameterGroupSpecParameter {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupSpecParameter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupSpecResource) DeepCopyInto(out *ParameterGroupSpecResource) {
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
	if in.Family != nil {
		in, out := &in.Family, &out.Family
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Parameter != nil {
		in, out := &in.Parameter, &out.Parameter
		*out = make([]ParameterGroupSpecParameter, len(*in))
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupSpecResource.
func (in *ParameterGroupSpecResource) DeepCopy() *ParameterGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterGroupStatus) DeepCopyInto(out *ParameterGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterGroupStatus.
func (in *ParameterGroupStatus) DeepCopy() *ParameterGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ParameterGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroup) DeepCopyInto(out *ReplicationGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroup.
func (in *ReplicationGroup) DeepCopy() *ReplicationGroup {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupList) DeepCopyInto(out *ReplicationGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ReplicationGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupList.
func (in *ReplicationGroupList) DeepCopy() *ReplicationGroupList {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReplicationGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupSpec) DeepCopyInto(out *ReplicationGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ReplicationGroupSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupSpec.
func (in *ReplicationGroupSpec) DeepCopy() *ReplicationGroupSpec {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupSpecClusterMode) DeepCopyInto(out *ReplicationGroupSpecClusterMode) {
	*out = *in
	if in.NumNodeGroups != nil {
		in, out := &in.NumNodeGroups, &out.NumNodeGroups
		*out = new(int64)
		**out = **in
	}
	if in.ReplicasPerNodeGroup != nil {
		in, out := &in.ReplicasPerNodeGroup, &out.ReplicasPerNodeGroup
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupSpecClusterMode.
func (in *ReplicationGroupSpecClusterMode) DeepCopy() *ReplicationGroupSpecClusterMode {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupSpecClusterMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupSpecResource) DeepCopyInto(out *ReplicationGroupSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.ApplyImmediately != nil {
		in, out := &in.ApplyImmediately, &out.ApplyImmediately
		*out = new(bool)
		**out = **in
	}
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.AtRestEncryptionEnabled != nil {
		in, out := &in.AtRestEncryptionEnabled, &out.AtRestEncryptionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AuthToken != nil {
		in, out := &in.AuthToken, &out.AuthToken
		*out = new(string)
		**out = **in
	}
	if in.AutoMinorVersionUpgrade != nil {
		in, out := &in.AutoMinorVersionUpgrade, &out.AutoMinorVersionUpgrade
		*out = new(bool)
		**out = **in
	}
	if in.AutomaticFailoverEnabled != nil {
		in, out := &in.AutomaticFailoverEnabled, &out.AutomaticFailoverEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterEnabled != nil {
		in, out := &in.ClusterEnabled, &out.ClusterEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ClusterMode != nil {
		in, out := &in.ClusterMode, &out.ClusterMode
		*out = new(ReplicationGroupSpecClusterMode)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigurationEndpointAddress != nil {
		in, out := &in.ConfigurationEndpointAddress, &out.ConfigurationEndpointAddress
		*out = new(string)
		**out = **in
	}
	if in.Engine != nil {
		in, out := &in.Engine, &out.Engine
		*out = new(string)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = new(string)
		**out = **in
	}
	if in.EngineVersionActual != nil {
		in, out := &in.EngineVersionActual, &out.EngineVersionActual
		*out = new(string)
		**out = **in
	}
	if in.FinalSnapshotIdentifier != nil {
		in, out := &in.FinalSnapshotIdentifier, &out.FinalSnapshotIdentifier
		*out = new(string)
		**out = **in
	}
	if in.GlobalReplicationGroupID != nil {
		in, out := &in.GlobalReplicationGroupID, &out.GlobalReplicationGroupID
		*out = new(string)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.MaintenanceWindow != nil {
		in, out := &in.MaintenanceWindow, &out.MaintenanceWindow
		*out = new(string)
		**out = **in
	}
	if in.MemberClusters != nil {
		in, out := &in.MemberClusters, &out.MemberClusters
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MultiAzEnabled != nil {
		in, out := &in.MultiAzEnabled, &out.MultiAzEnabled
		*out = new(bool)
		**out = **in
	}
	if in.NodeType != nil {
		in, out := &in.NodeType, &out.NodeType
		*out = new(string)
		**out = **in
	}
	if in.NotificationTopicArn != nil {
		in, out := &in.NotificationTopicArn, &out.NotificationTopicArn
		*out = new(string)
		**out = **in
	}
	if in.NumberCacheClusters != nil {
		in, out := &in.NumberCacheClusters, &out.NumberCacheClusters
		*out = new(int64)
		**out = **in
	}
	if in.ParameterGroupName != nil {
		in, out := &in.ParameterGroupName, &out.ParameterGroupName
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.PrimaryEndpointAddress != nil {
		in, out := &in.PrimaryEndpointAddress, &out.PrimaryEndpointAddress
		*out = new(string)
		**out = **in
	}
	if in.ReaderEndpointAddress != nil {
		in, out := &in.ReaderEndpointAddress, &out.ReaderEndpointAddress
		*out = new(string)
		**out = **in
	}
	if in.ReplicationGroupDescription != nil {
		in, out := &in.ReplicationGroupDescription, &out.ReplicationGroupDescription
		*out = new(string)
		**out = **in
	}
	if in.ReplicationGroupID != nil {
		in, out := &in.ReplicationGroupID, &out.ReplicationGroupID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupIDS != nil {
		in, out := &in.SecurityGroupIDS, &out.SecurityGroupIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupNames != nil {
		in, out := &in.SecurityGroupNames, &out.SecurityGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotArns != nil {
		in, out := &in.SnapshotArns, &out.SnapshotArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SnapshotName != nil {
		in, out := &in.SnapshotName, &out.SnapshotName
		*out = new(string)
		**out = **in
	}
	if in.SnapshotRetentionLimit != nil {
		in, out := &in.SnapshotRetentionLimit, &out.SnapshotRetentionLimit
		*out = new(int64)
		**out = **in
	}
	if in.SnapshotWindow != nil {
		in, out := &in.SnapshotWindow, &out.SnapshotWindow
		*out = new(string)
		**out = **in
	}
	if in.SubnetGroupName != nil {
		in, out := &in.SubnetGroupName, &out.SubnetGroupName
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
	if in.TransitEncryptionEnabled != nil {
		in, out := &in.TransitEncryptionEnabled, &out.TransitEncryptionEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupSpecResource.
func (in *ReplicationGroupSpecResource) DeepCopy() *ReplicationGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReplicationGroupStatus) DeepCopyInto(out *ReplicationGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReplicationGroupStatus.
func (in *ReplicationGroupStatus) DeepCopy() *ReplicationGroupStatus {
	if in == nil {
		return nil
	}
	out := new(ReplicationGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroup.
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupList) DeepCopyInto(out *SecurityGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SecurityGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupList.
func (in *SecurityGroupList) DeepCopy() *SecurityGroupList {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SecurityGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupSpec) DeepCopyInto(out *SecurityGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SecurityGroupSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupSpec.
func (in *SecurityGroupSpec) DeepCopy() *SecurityGroupSpec {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupSpecResource) DeepCopyInto(out *SecurityGroupSpecResource) {
	*out = *in
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
	if in.SecurityGroupNames != nil {
		in, out := &in.SecurityGroupNames, &out.SecurityGroupNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupSpecResource.
func (in *SecurityGroupSpecResource) DeepCopy() *SecurityGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityGroupStatus) DeepCopyInto(out *SecurityGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityGroupStatus.
func (in *SecurityGroupStatus) DeepCopy() *SecurityGroupStatus {
	if in == nil {
		return nil
	}
	out := new(SecurityGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroup) DeepCopyInto(out *SubnetGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroup.
func (in *SubnetGroup) DeepCopy() *SubnetGroup {
	if in == nil {
		return nil
	}
	out := new(SubnetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupList) DeepCopyInto(out *SubnetGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubnetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupList.
func (in *SubnetGroupList) DeepCopy() *SubnetGroupList {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubnetGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupSpec) DeepCopyInto(out *SubnetGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SubnetGroupSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupSpec.
func (in *SubnetGroupSpec) DeepCopy() *SubnetGroupSpec {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupSpecResource) DeepCopyInto(out *SubnetGroupSpecResource) {
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
	if in.SubnetIDS != nil {
		in, out := &in.SubnetIDS, &out.SubnetIDS
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupSpecResource.
func (in *SubnetGroupSpecResource) DeepCopy() *SubnetGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetGroupStatus) DeepCopyInto(out *SubnetGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetGroupStatus.
func (in *SubnetGroupStatus) DeepCopy() *SubnetGroupStatus {
	if in == nil {
		return nil
	}
	out := new(SubnetGroupStatus)
	in.DeepCopyInto(out)
	return out
}
