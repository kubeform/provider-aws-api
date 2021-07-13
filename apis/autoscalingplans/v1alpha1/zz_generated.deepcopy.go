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
func (in *ScalingPlan) DeepCopyInto(out *ScalingPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlan.
func (in *ScalingPlan) DeepCopy() *ScalingPlan {
	if in == nil {
		return nil
	}
	out := new(ScalingPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanList) DeepCopyInto(out *ScalingPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalingPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanList.
func (in *ScalingPlanList) DeepCopy() *ScalingPlanList {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpec) DeepCopyInto(out *ScalingPlanSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ScalingPlanSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpec.
func (in *ScalingPlanSpec) DeepCopy() *ScalingPlanSpec {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecApplicationSource) DeepCopyInto(out *ScalingPlanSpecApplicationSource) {
	*out = *in
	if in.CloudformationStackArn != nil {
		in, out := &in.CloudformationStackArn, &out.CloudformationStackArn
		*out = new(string)
		**out = **in
	}
	if in.TagFilter != nil {
		in, out := &in.TagFilter, &out.TagFilter
		*out = make([]ScalingPlanSpecApplicationSourceTagFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecApplicationSource.
func (in *ScalingPlanSpecApplicationSource) DeepCopy() *ScalingPlanSpecApplicationSource {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecApplicationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecApplicationSourceTagFilter) DeepCopyInto(out *ScalingPlanSpecApplicationSourceTagFilter) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecApplicationSourceTagFilter.
func (in *ScalingPlanSpecApplicationSourceTagFilter) DeepCopy() *ScalingPlanSpecApplicationSourceTagFilter {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecApplicationSourceTagFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecResource) DeepCopyInto(out *ScalingPlanSpecResource) {
	*out = *in
	if in.ApplicationSource != nil {
		in, out := &in.ApplicationSource, &out.ApplicationSource
		*out = new(ScalingPlanSpecApplicationSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ScalingInstruction != nil {
		in, out := &in.ScalingInstruction, &out.ScalingInstruction
		*out = make([]ScalingPlanSpecScalingInstruction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScalingPlanVersion != nil {
		in, out := &in.ScalingPlanVersion, &out.ScalingPlanVersion
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecResource.
func (in *ScalingPlanSpecResource) DeepCopy() *ScalingPlanSpecResource {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstruction) DeepCopyInto(out *ScalingPlanSpecScalingInstruction) {
	*out = *in
	if in.CustomizedLoadMetricSpecification != nil {
		in, out := &in.CustomizedLoadMetricSpecification, &out.CustomizedLoadMetricSpecification
		*out = new(ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableDynamicScaling != nil {
		in, out := &in.DisableDynamicScaling, &out.DisableDynamicScaling
		*out = new(bool)
		**out = **in
	}
	if in.MaxCapacity != nil {
		in, out := &in.MaxCapacity, &out.MaxCapacity
		*out = new(int64)
		**out = **in
	}
	if in.MinCapacity != nil {
		in, out := &in.MinCapacity, &out.MinCapacity
		*out = new(int64)
		**out = **in
	}
	if in.PredefinedLoadMetricSpecification != nil {
		in, out := &in.PredefinedLoadMetricSpecification, &out.PredefinedLoadMetricSpecification
		*out = new(ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.PredictiveScalingMaxCapacityBehavior != nil {
		in, out := &in.PredictiveScalingMaxCapacityBehavior, &out.PredictiveScalingMaxCapacityBehavior
		*out = new(string)
		**out = **in
	}
	if in.PredictiveScalingMaxCapacityBuffer != nil {
		in, out := &in.PredictiveScalingMaxCapacityBuffer, &out.PredictiveScalingMaxCapacityBuffer
		*out = new(int64)
		**out = **in
	}
	if in.PredictiveScalingMode != nil {
		in, out := &in.PredictiveScalingMode, &out.PredictiveScalingMode
		*out = new(string)
		**out = **in
	}
	if in.ResourceID != nil {
		in, out := &in.ResourceID, &out.ResourceID
		*out = new(string)
		**out = **in
	}
	if in.ScalableDimension != nil {
		in, out := &in.ScalableDimension, &out.ScalableDimension
		*out = new(string)
		**out = **in
	}
	if in.ScalingPolicyUpdateBehavior != nil {
		in, out := &in.ScalingPolicyUpdateBehavior, &out.ScalingPolicyUpdateBehavior
		*out = new(string)
		**out = **in
	}
	if in.ScheduledActionBufferTime != nil {
		in, out := &in.ScheduledActionBufferTime, &out.ScheduledActionBufferTime
		*out = new(int64)
		**out = **in
	}
	if in.ServiceNamespace != nil {
		in, out := &in.ServiceNamespace, &out.ServiceNamespace
		*out = new(string)
		**out = **in
	}
	if in.TargetTrackingConfiguration != nil {
		in, out := &in.TargetTrackingConfiguration, &out.TargetTrackingConfiguration
		*out = make([]ScalingPlanSpecScalingInstructionTargetTrackingConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstruction.
func (in *ScalingPlanSpecScalingInstruction) DeepCopy() *ScalingPlanSpecScalingInstruction {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstruction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification) DeepCopyInto(out *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification.
func (in *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification) DeepCopy() *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification) DeepCopyInto(out *ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification) {
	*out = *in
	if in.PredefinedLoadMetricType != nil {
		in, out := &in.PredefinedLoadMetricType, &out.PredefinedLoadMetricType
		*out = new(string)
		**out = **in
	}
	if in.ResourceLabel != nil {
		in, out := &in.ResourceLabel, &out.ResourceLabel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification.
func (in *ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification) DeepCopy() *ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfiguration) DeepCopyInto(out *ScalingPlanSpecScalingInstructionTargetTrackingConfiguration) {
	*out = *in
	if in.CustomizedScalingMetricSpecification != nil {
		in, out := &in.CustomizedScalingMetricSpecification, &out.CustomizedScalingMetricSpecification
		*out = new(ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.DisableScaleIn != nil {
		in, out := &in.DisableScaleIn, &out.DisableScaleIn
		*out = new(bool)
		**out = **in
	}
	if in.EstimatedInstanceWarmup != nil {
		in, out := &in.EstimatedInstanceWarmup, &out.EstimatedInstanceWarmup
		*out = new(int64)
		**out = **in
	}
	if in.PredefinedScalingMetricSpecification != nil {
		in, out := &in.PredefinedScalingMetricSpecification, &out.PredefinedScalingMetricSpecification
		*out = new(ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification)
		(*in).DeepCopyInto(*out)
	}
	if in.ScaleInCooldown != nil {
		in, out := &in.ScaleInCooldown, &out.ScaleInCooldown
		*out = new(int64)
		**out = **in
	}
	if in.ScaleOutCooldown != nil {
		in, out := &in.ScaleOutCooldown, &out.ScaleOutCooldown
		*out = new(int64)
		**out = **in
	}
	if in.TargetValue != nil {
		in, out := &in.TargetValue, &out.TargetValue
		*out = new(float64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionTargetTrackingConfiguration.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfiguration) DeepCopy() *ScalingPlanSpecScalingInstructionTargetTrackingConfiguration {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionTargetTrackingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification) DeepCopyInto(out *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	if in.MetricName != nil {
		in, out := &in.MetricName, &out.MetricName
		*out = new(string)
		**out = **in
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Statistic != nil {
		in, out := &in.Statistic, &out.Statistic
		*out = new(string)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification) DeepCopy() *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification) DeepCopyInto(out *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification) {
	*out = *in
	if in.PredefinedScalingMetricType != nil {
		in, out := &in.PredefinedScalingMetricType, &out.PredefinedScalingMetricType
		*out = new(string)
		**out = **in
	}
	if in.ResourceLabel != nil {
		in, out := &in.ResourceLabel, &out.ResourceLabel
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification) DeepCopy() *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanStatus) DeepCopyInto(out *ScalingPlanStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanStatus.
func (in *ScalingPlanStatus) DeepCopy() *ScalingPlanStatus {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanStatus)
	in.DeepCopyInto(out)
	return out
}
