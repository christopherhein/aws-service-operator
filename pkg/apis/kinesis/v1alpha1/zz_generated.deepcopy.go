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
func (in *Stream) DeepCopyInto(out *Stream) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Stream.
func (in *Stream) DeepCopy() *Stream {
	if in == nil {
		return nil
	}
	out := new(Stream)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Stream) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamConsumer) DeepCopyInto(out *StreamConsumer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamConsumer.
func (in *StreamConsumer) DeepCopy() *StreamConsumer {
	if in == nil {
		return nil
	}
	out := new(StreamConsumer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamConsumer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamConsumerList) DeepCopyInto(out *StreamConsumerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StreamConsumer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamConsumerList.
func (in *StreamConsumerList) DeepCopy() *StreamConsumerList {
	if in == nil {
		return nil
	}
	out := new(StreamConsumerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamConsumerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamConsumerOutputs) DeepCopyInto(out *StreamConsumerOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamConsumerOutputs.
func (in *StreamConsumerOutputs) DeepCopy() *StreamConsumerOutputs {
	if in == nil {
		return nil
	}
	out := new(StreamConsumerOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamConsumerSpec) DeepCopyInto(out *StreamConsumerSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamConsumerSpec.
func (in *StreamConsumerSpec) DeepCopy() *StreamConsumerSpec {
	if in == nil {
		return nil
	}
	out := new(StreamConsumerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamConsumerStatus) DeepCopyInto(out *StreamConsumerStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamConsumerStatus.
func (in *StreamConsumerStatus) DeepCopy() *StreamConsumerStatus {
	if in == nil {
		return nil
	}
	out := new(StreamConsumerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamList) DeepCopyInto(out *StreamList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Stream, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamList.
func (in *StreamList) DeepCopy() *StreamList {
	if in == nil {
		return nil
	}
	out := new(StreamList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StreamList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamOutputs) DeepCopyInto(out *StreamOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamOutputs.
func (in *StreamOutputs) DeepCopy() *StreamOutputs {
	if in == nil {
		return nil
	}
	out := new(StreamOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSpec) DeepCopyInto(out *StreamSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.StreamEncryption = in.StreamEncryption
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]StreamSpecTag, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSpec.
func (in *StreamSpec) DeepCopy() *StreamSpec {
	if in == nil {
		return nil
	}
	out := new(StreamSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSpecStreamEncryption) DeepCopyInto(out *StreamSpecStreamEncryption) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSpecStreamEncryption.
func (in *StreamSpecStreamEncryption) DeepCopy() *StreamSpecStreamEncryption {
	if in == nil {
		return nil
	}
	out := new(StreamSpecStreamEncryption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamSpecTag) DeepCopyInto(out *StreamSpecTag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamSpecTag.
func (in *StreamSpecTag) DeepCopy() *StreamSpecTag {
	if in == nil {
		return nil
	}
	out := new(StreamSpecTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StreamStatus) DeepCopyInto(out *StreamStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StreamStatus.
func (in *StreamStatus) DeepCopy() *StreamStatus {
	if in == nil {
		return nil
	}
	out := new(StreamStatus)
	in.DeepCopyInto(out)
	return out
}