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
func (in *Access) DeepCopyInto(out *Access) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Access.
func (in *Access) DeepCopy() *Access {
	if in == nil {
		return nil
	}
	out := new(Access)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Access) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessList) DeepCopyInto(out *AccessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Access, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessList.
func (in *AccessList) DeepCopy() *AccessList {
	if in == nil {
		return nil
	}
	out := new(AccessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessSpec) DeepCopyInto(out *AccessSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AccessSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessSpec.
func (in *AccessSpec) DeepCopy() *AccessSpec {
	if in == nil {
		return nil
	}
	out := new(AccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessSpecHomeDirectoryMappings) DeepCopyInto(out *AccessSpecHomeDirectoryMappings) {
	*out = *in
	if in.Entry != nil {
		in, out := &in.Entry, &out.Entry
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessSpecHomeDirectoryMappings.
func (in *AccessSpecHomeDirectoryMappings) DeepCopy() *AccessSpecHomeDirectoryMappings {
	if in == nil {
		return nil
	}
	out := new(AccessSpecHomeDirectoryMappings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessSpecPosixProfile) DeepCopyInto(out *AccessSpecPosixProfile) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessSpecPosixProfile.
func (in *AccessSpecPosixProfile) DeepCopy() *AccessSpecPosixProfile {
	if in == nil {
		return nil
	}
	out := new(AccessSpecPosixProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessSpecResource) DeepCopyInto(out *AccessSpecResource) {
	*out = *in
	if in.ExternalID != nil {
		in, out := &in.ExternalID, &out.ExternalID
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectory != nil {
		in, out := &in.HomeDirectory, &out.HomeDirectory
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectoryMappings != nil {
		in, out := &in.HomeDirectoryMappings, &out.HomeDirectoryMappings
		*out = make([]AccessSpecHomeDirectoryMappings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HomeDirectoryType != nil {
		in, out := &in.HomeDirectoryType, &out.HomeDirectoryType
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.PosixProfile != nil {
		in, out := &in.PosixProfile, &out.PosixProfile
		*out = new(AccessSpecPosixProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessSpecResource.
func (in *AccessSpecResource) DeepCopy() *AccessSpecResource {
	if in == nil {
		return nil
	}
	out := new(AccessSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessStatus) DeepCopyInto(out *AccessStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessStatus.
func (in *AccessStatus) DeepCopy() *AccessStatus {
	if in == nil {
		return nil
	}
	out := new(AccessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Server) DeepCopyInto(out *Server) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Server.
func (in *Server) DeepCopy() *Server {
	if in == nil {
		return nil
	}
	out := new(Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Server) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerList) DeepCopyInto(out *ServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Server, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerList.
func (in *ServerList) DeepCopy() *ServerList {
	if in == nil {
		return nil
	}
	out := new(ServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpec) DeepCopyInto(out *ServerSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ServerSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpec.
func (in *ServerSpec) DeepCopy() *ServerSpec {
	if in == nil {
		return nil
	}
	out := new(ServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecEndpointDetails) DeepCopyInto(out *ServerSpecEndpointDetails) {
	*out = *in
	if in.AddressAllocationIDS != nil {
		in, out := &in.AddressAllocationIDS, &out.AddressAllocationIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
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
	if in.VpcEndpointID != nil {
		in, out := &in.VpcEndpointID, &out.VpcEndpointID
		*out = new(string)
		**out = **in
	}
	if in.VpcID != nil {
		in, out := &in.VpcID, &out.VpcID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecEndpointDetails.
func (in *ServerSpecEndpointDetails) DeepCopy() *ServerSpecEndpointDetails {
	if in == nil {
		return nil
	}
	out := new(ServerSpecEndpointDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSpecResource) DeepCopyInto(out *ServerSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Certificate != nil {
		in, out := &in.Certificate, &out.Certificate
		*out = new(string)
		**out = **in
	}
	if in.DirectoryID != nil {
		in, out := &in.DirectoryID, &out.DirectoryID
		*out = new(string)
		**out = **in
	}
	if in.Domain != nil {
		in, out := &in.Domain, &out.Domain
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(string)
		**out = **in
	}
	if in.EndpointDetails != nil {
		in, out := &in.EndpointDetails, &out.EndpointDetails
		*out = new(ServerSpecEndpointDetails)
		(*in).DeepCopyInto(*out)
	}
	if in.EndpointType != nil {
		in, out := &in.EndpointType, &out.EndpointType
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = new(string)
		**out = **in
	}
	if in.HostKey != nil {
		in, out := &in.HostKey, &out.HostKey
		*out = new(string)
		**out = **in
	}
	if in.HostKeyFingerprint != nil {
		in, out := &in.HostKeyFingerprint, &out.HostKeyFingerprint
		*out = new(string)
		**out = **in
	}
	if in.IdentityProviderType != nil {
		in, out := &in.IdentityProviderType, &out.IdentityProviderType
		*out = new(string)
		**out = **in
	}
	if in.InvocationRole != nil {
		in, out := &in.InvocationRole, &out.InvocationRole
		*out = new(string)
		**out = **in
	}
	if in.LoggingRole != nil {
		in, out := &in.LoggingRole, &out.LoggingRole
		*out = new(string)
		**out = **in
	}
	if in.Protocols != nil {
		in, out := &in.Protocols, &out.Protocols
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SecurityPolicyName != nil {
		in, out := &in.SecurityPolicyName, &out.SecurityPolicyName
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
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSpecResource.
func (in *ServerSpecResource) DeepCopy() *ServerSpecResource {
	if in == nil {
		return nil
	}
	out := new(ServerSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerStatus) DeepCopyInto(out *ServerStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerStatus.
func (in *ServerStatus) DeepCopy() *ServerStatus {
	if in == nil {
		return nil
	}
	out := new(ServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SshKey) DeepCopyInto(out *SshKey) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SshKey.
func (in *SshKey) DeepCopy() *SshKey {
	if in == nil {
		return nil
	}
	out := new(SshKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SshKey) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SshKeyList) DeepCopyInto(out *SshKeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SshKey, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SshKeyList.
func (in *SshKeyList) DeepCopy() *SshKeyList {
	if in == nil {
		return nil
	}
	out := new(SshKeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SshKeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SshKeySpec) DeepCopyInto(out *SshKeySpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SshKeySpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SshKeySpec.
func (in *SshKeySpec) DeepCopy() *SshKeySpec {
	if in == nil {
		return nil
	}
	out := new(SshKeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SshKeySpecResource) DeepCopyInto(out *SshKeySpecResource) {
	*out = *in
	if in.Body != nil {
		in, out := &in.Body, &out.Body
		*out = new(string)
		**out = **in
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
		*out = new(string)
		**out = **in
	}
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SshKeySpecResource.
func (in *SshKeySpecResource) DeepCopy() *SshKeySpecResource {
	if in == nil {
		return nil
	}
	out := new(SshKeySpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SshKeyStatus) DeepCopyInto(out *SshKeyStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SshKeyStatus.
func (in *SshKeyStatus) DeepCopy() *SshKeyStatus {
	if in == nil {
		return nil
	}
	out := new(SshKeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *User) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserList) DeepCopyInto(out *UserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]User, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserList.
func (in *UserList) DeepCopy() *UserList {
	if in == nil {
		return nil
	}
	out := new(UserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(UserSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecHomeDirectoryMappings) DeepCopyInto(out *UserSpecHomeDirectoryMappings) {
	*out = *in
	if in.Entry != nil {
		in, out := &in.Entry, &out.Entry
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecHomeDirectoryMappings.
func (in *UserSpecHomeDirectoryMappings) DeepCopy() *UserSpecHomeDirectoryMappings {
	if in == nil {
		return nil
	}
	out := new(UserSpecHomeDirectoryMappings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecPosixProfile) DeepCopyInto(out *UserSpecPosixProfile) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecPosixProfile.
func (in *UserSpecPosixProfile) DeepCopy() *UserSpecPosixProfile {
	if in == nil {
		return nil
	}
	out := new(UserSpecPosixProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpecResource) DeepCopyInto(out *UserSpecResource) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectory != nil {
		in, out := &in.HomeDirectory, &out.HomeDirectory
		*out = new(string)
		**out = **in
	}
	if in.HomeDirectoryMappings != nil {
		in, out := &in.HomeDirectoryMappings, &out.HomeDirectoryMappings
		*out = make([]UserSpecHomeDirectoryMappings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HomeDirectoryType != nil {
		in, out := &in.HomeDirectoryType, &out.HomeDirectoryType
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.PosixProfile != nil {
		in, out := &in.PosixProfile, &out.PosixProfile
		*out = new(UserSpecPosixProfile)
		(*in).DeepCopyInto(*out)
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.ServerID != nil {
		in, out := &in.ServerID, &out.ServerID
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
	if in.UserName != nil {
		in, out := &in.UserName, &out.UserName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpecResource.
func (in *UserSpecResource) DeepCopy() *UserSpecResource {
	if in == nil {
		return nil
	}
	out := new(UserSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}
