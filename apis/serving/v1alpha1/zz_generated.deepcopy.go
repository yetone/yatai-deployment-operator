//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/autoscaling/v2beta2"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeployment) DeepCopyInto(out *BentoDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeployment.
func (in *BentoDeployment) DeepCopy() *BentoDeployment {
	if in == nil {
		return nil
	}
	out := new(BentoDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BentoDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentAutoscalingSpec) DeepCopyInto(out *BentoDeploymentAutoscalingSpec) {
	*out = *in
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2beta2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentAutoscalingSpec.
func (in *BentoDeploymentAutoscalingSpec) DeepCopy() *BentoDeploymentAutoscalingSpec {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentAutoscalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentList) DeepCopyInto(out *BentoDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BentoDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentList.
func (in *BentoDeploymentList) DeepCopy() *BentoDeploymentList {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BentoDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentRunnerSpec) DeepCopyInto(out *BentoDeploymentRunnerSpec) {
	*out = *in
	in.Autoscaling.DeepCopyInto(&out.Autoscaling)
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentRunnerSpec.
func (in *BentoDeploymentRunnerSpec) DeepCopy() *BentoDeploymentRunnerSpec {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentRunnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentSpec) DeepCopyInto(out *BentoDeploymentSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(DeploymentTargetResources)
		(*in).DeepCopyInto(*out)
	}
	if in.Autoscaling != nil {
		in, out := &in.Autoscaling, &out.Autoscaling
		*out = new(DeploymentTargetHPAConf)
		(*in).DeepCopyInto(*out)
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = new([]LabelItemSchema)
		if **in != nil {
			in, out := *in, *out
			*out = make([]LabelItemSchema, len(*in))
			copy(*out, *in)
		}
	}
	if in.Runners != nil {
		in, out := &in.Runners, &out.Runners
		*out = make([]BentoDeploymentRunnerSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentSpec.
func (in *BentoDeploymentSpec) DeepCopy() *BentoDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BentoDeploymentStatus) DeepCopyInto(out *BentoDeploymentStatus) {
	*out = *in
	if in.PodSelector != nil {
		in, out := &in.PodSelector, &out.PodSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BentoDeploymentStatus.
func (in *BentoDeploymentStatus) DeepCopy() *BentoDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(BentoDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTargetHPAConf) DeepCopyInto(out *DeploymentTargetHPAConf) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(int32)
		**out = **in
	}
	if in.GPU != nil {
		in, out := &in.GPU, &out.GPU
		*out = new(int32)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(string)
		**out = **in
	}
	if in.QPS != nil {
		in, out := &in.QPS, &out.QPS
		*out = new(int64)
		**out = **in
	}
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTargetHPAConf.
func (in *DeploymentTargetHPAConf) DeepCopy() *DeploymentTargetHPAConf {
	if in == nil {
		return nil
	}
	out := new(DeploymentTargetHPAConf)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTargetResourceItem) DeepCopyInto(out *DeploymentTargetResourceItem) {
	*out = *in
	if in.Custom != nil {
		in, out := &in.Custom, &out.Custom
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTargetResourceItem.
func (in *DeploymentTargetResourceItem) DeepCopy() *DeploymentTargetResourceItem {
	if in == nil {
		return nil
	}
	out := new(DeploymentTargetResourceItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentTargetResources) DeepCopyInto(out *DeploymentTargetResources) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = new(DeploymentTargetResourceItem)
		(*in).DeepCopyInto(*out)
	}
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = new(DeploymentTargetResourceItem)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentTargetResources.
func (in *DeploymentTargetResources) DeepCopy() *DeploymentTargetResources {
	if in == nil {
		return nil
	}
	out := new(DeploymentTargetResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelItemSchema) DeepCopyInto(out *LabelItemSchema) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelItemSchema.
func (in *LabelItemSchema) DeepCopy() *LabelItemSchema {
	if in == nil {
		return nil
	}
	out := new(LabelItemSchema)
	in.DeepCopyInto(out)
	return out
}
