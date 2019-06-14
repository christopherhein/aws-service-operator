// +build !ignore_autogenerated

/*
Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License"). You may
not use this file except in compliance with the License. A copy of the
License is located at

     http://aws.amazon.com/apache2.0/

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2) DeepCopyInto(out *EnvironmentEC2) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2.
func (in *EnvironmentEC2) DeepCopy() *EnvironmentEC2 {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentEC2) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2List) DeepCopyInto(out *EnvironmentEC2List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvironmentEC2, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2List.
func (in *EnvironmentEC2List) DeepCopy() *EnvironmentEC2List {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentEC2List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2Outputs) DeepCopyInto(out *EnvironmentEC2Outputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2Outputs.
func (in *EnvironmentEC2Outputs) DeepCopy() *EnvironmentEC2Outputs {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2Outputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2Spec) DeepCopyInto(out *EnvironmentEC2Spec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]EnvironmentEC2SpecRepository, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2Spec.
func (in *EnvironmentEC2Spec) DeepCopy() *EnvironmentEC2Spec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2SpecRepository) DeepCopyInto(out *EnvironmentEC2SpecRepository) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2SpecRepository.
func (in *EnvironmentEC2SpecRepository) DeepCopy() *EnvironmentEC2SpecRepository {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2SpecRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentEC2Status) DeepCopyInto(out *EnvironmentEC2Status) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentEC2Status.
func (in *EnvironmentEC2Status) DeepCopy() *EnvironmentEC2Status {
	if in == nil {
		return nil
	}
	out := new(EnvironmentEC2Status)
	in.DeepCopyInto(out)
	return out
}