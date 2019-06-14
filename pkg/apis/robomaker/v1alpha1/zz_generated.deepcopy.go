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
func (in *Fleet) DeepCopyInto(out *Fleet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet.
func (in *Fleet) DeepCopy() *Fleet {
	if in == nil {
		return nil
	}
	out := new(Fleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fleet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetList) DeepCopyInto(out *FleetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fleet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetList.
func (in *FleetList) DeepCopy() *FleetList {
	if in == nil {
		return nil
	}
	out := new(FleetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetOutputs) DeepCopyInto(out *FleetOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetOutputs.
func (in *FleetOutputs) DeepCopy() *FleetOutputs {
	if in == nil {
		return nil
	}
	out := new(FleetOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpec) DeepCopyInto(out *FleetSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.Tags = in.Tags
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpec.
func (in *FleetSpec) DeepCopy() *FleetSpec {
	if in == nil {
		return nil
	}
	out := new(FleetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpecTags) DeepCopyInto(out *FleetSpecTags) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpecTags.
func (in *FleetSpecTags) DeepCopy() *FleetSpecTags {
	if in == nil {
		return nil
	}
	out := new(FleetSpecTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetStatus) DeepCopyInto(out *FleetStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetStatus.
func (in *FleetStatus) DeepCopy() *FleetStatus {
	if in == nil {
		return nil
	}
	out := new(FleetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Robot) DeepCopyInto(out *Robot) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Robot.
func (in *Robot) DeepCopy() *Robot {
	if in == nil {
		return nil
	}
	out := new(Robot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Robot) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplication) DeepCopyInto(out *RobotApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplication.
func (in *RobotApplication) DeepCopy() *RobotApplication {
	if in == nil {
		return nil
	}
	out := new(RobotApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationList) DeepCopyInto(out *RobotApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RobotApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationList.
func (in *RobotApplicationList) DeepCopy() *RobotApplicationList {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationOutputs) DeepCopyInto(out *RobotApplicationOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationOutputs.
func (in *RobotApplicationOutputs) DeepCopy() *RobotApplicationOutputs {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationSpec) DeepCopyInto(out *RobotApplicationSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.RobotSoftwareSuite = in.RobotSoftwareSuite
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]RobotApplicationSpecSource, len(*in))
		copy(*out, *in)
	}
	out.Tags = in.Tags
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationSpec.
func (in *RobotApplicationSpec) DeepCopy() *RobotApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationSpecRobotSoftwareSuite) DeepCopyInto(out *RobotApplicationSpecRobotSoftwareSuite) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationSpecRobotSoftwareSuite.
func (in *RobotApplicationSpecRobotSoftwareSuite) DeepCopy() *RobotApplicationSpecRobotSoftwareSuite {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationSpecRobotSoftwareSuite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationSpecSource) DeepCopyInto(out *RobotApplicationSpecSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationSpecSource.
func (in *RobotApplicationSpecSource) DeepCopy() *RobotApplicationSpecSource {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationSpecSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationSpecTags) DeepCopyInto(out *RobotApplicationSpecTags) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationSpecTags.
func (in *RobotApplicationSpecTags) DeepCopy() *RobotApplicationSpecTags {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationSpecTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationStatus) DeepCopyInto(out *RobotApplicationStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationStatus.
func (in *RobotApplicationStatus) DeepCopy() *RobotApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationVersion) DeepCopyInto(out *RobotApplicationVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationVersion.
func (in *RobotApplicationVersion) DeepCopy() *RobotApplicationVersion {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotApplicationVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationVersionList) DeepCopyInto(out *RobotApplicationVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RobotApplicationVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationVersionList.
func (in *RobotApplicationVersionList) DeepCopy() *RobotApplicationVersionList {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotApplicationVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationVersionOutputs) DeepCopyInto(out *RobotApplicationVersionOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationVersionOutputs.
func (in *RobotApplicationVersionOutputs) DeepCopy() *RobotApplicationVersionOutputs {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationVersionOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationVersionSpec) DeepCopyInto(out *RobotApplicationVersionSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationVersionSpec.
func (in *RobotApplicationVersionSpec) DeepCopy() *RobotApplicationVersionSpec {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotApplicationVersionStatus) DeepCopyInto(out *RobotApplicationVersionStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotApplicationVersionStatus.
func (in *RobotApplicationVersionStatus) DeepCopy() *RobotApplicationVersionStatus {
	if in == nil {
		return nil
	}
	out := new(RobotApplicationVersionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotList) DeepCopyInto(out *RobotList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Robot, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotList.
func (in *RobotList) DeepCopy() *RobotList {
	if in == nil {
		return nil
	}
	out := new(RobotList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RobotList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotOutputs) DeepCopyInto(out *RobotOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotOutputs.
func (in *RobotOutputs) DeepCopy() *RobotOutputs {
	if in == nil {
		return nil
	}
	out := new(RobotOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotSpec) DeepCopyInto(out *RobotSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.Tags = in.Tags
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotSpec.
func (in *RobotSpec) DeepCopy() *RobotSpec {
	if in == nil {
		return nil
	}
	out := new(RobotSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotSpecTags) DeepCopyInto(out *RobotSpecTags) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotSpecTags.
func (in *RobotSpecTags) DeepCopy() *RobotSpecTags {
	if in == nil {
		return nil
	}
	out := new(RobotSpecTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RobotStatus) DeepCopyInto(out *RobotStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RobotStatus.
func (in *RobotStatus) DeepCopy() *RobotStatus {
	if in == nil {
		return nil
	}
	out := new(RobotStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplication) DeepCopyInto(out *SimulationApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplication.
func (in *SimulationApplication) DeepCopy() *SimulationApplication {
	if in == nil {
		return nil
	}
	out := new(SimulationApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimulationApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationList) DeepCopyInto(out *SimulationApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SimulationApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationList.
func (in *SimulationApplicationList) DeepCopy() *SimulationApplicationList {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimulationApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationOutputs) DeepCopyInto(out *SimulationApplicationOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationOutputs.
func (in *SimulationApplicationOutputs) DeepCopy() *SimulationApplicationOutputs {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationSpec) DeepCopyInto(out *SimulationApplicationSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.RenderingEngine = in.RenderingEngine
	out.RobotSoftwareSuite = in.RobotSoftwareSuite
	out.SimulationSoftwareSuite = in.SimulationSoftwareSuite
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make([]SimulationApplicationSpecSource, len(*in))
		copy(*out, *in)
	}
	out.Tags = in.Tags
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationSpec.
func (in *SimulationApplicationSpec) DeepCopy() *SimulationApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationSpecRenderingEngine) DeepCopyInto(out *SimulationApplicationSpecRenderingEngine) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationSpecRenderingEngine.
func (in *SimulationApplicationSpecRenderingEngine) DeepCopy() *SimulationApplicationSpecRenderingEngine {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationSpecRenderingEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationSpecRobotSoftwareSuite) DeepCopyInto(out *SimulationApplicationSpecRobotSoftwareSuite) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationSpecRobotSoftwareSuite.
func (in *SimulationApplicationSpecRobotSoftwareSuite) DeepCopy() *SimulationApplicationSpecRobotSoftwareSuite {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationSpecRobotSoftwareSuite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationSpecSimulationSoftwareSuite) DeepCopyInto(out *SimulationApplicationSpecSimulationSoftwareSuite) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationSpecSimulationSoftwareSuite.
func (in *SimulationApplicationSpecSimulationSoftwareSuite) DeepCopy() *SimulationApplicationSpecSimulationSoftwareSuite {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationSpecSimulationSoftwareSuite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationSpecSource) DeepCopyInto(out *SimulationApplicationSpecSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationSpecSource.
func (in *SimulationApplicationSpecSource) DeepCopy() *SimulationApplicationSpecSource {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationSpecSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationSpecTags) DeepCopyInto(out *SimulationApplicationSpecTags) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationSpecTags.
func (in *SimulationApplicationSpecTags) DeepCopy() *SimulationApplicationSpecTags {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationSpecTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationStatus) DeepCopyInto(out *SimulationApplicationStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationStatus.
func (in *SimulationApplicationStatus) DeepCopy() *SimulationApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationVersion) DeepCopyInto(out *SimulationApplicationVersion) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationVersion.
func (in *SimulationApplicationVersion) DeepCopy() *SimulationApplicationVersion {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimulationApplicationVersion) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationVersionList) DeepCopyInto(out *SimulationApplicationVersionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SimulationApplicationVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationVersionList.
func (in *SimulationApplicationVersionList) DeepCopy() *SimulationApplicationVersionList {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationVersionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimulationApplicationVersionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationVersionOutputs) DeepCopyInto(out *SimulationApplicationVersionOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationVersionOutputs.
func (in *SimulationApplicationVersionOutputs) DeepCopy() *SimulationApplicationVersionOutputs {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationVersionOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationVersionSpec) DeepCopyInto(out *SimulationApplicationVersionSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationVersionSpec.
func (in *SimulationApplicationVersionSpec) DeepCopy() *SimulationApplicationVersionSpec {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationVersionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimulationApplicationVersionStatus) DeepCopyInto(out *SimulationApplicationVersionStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimulationApplicationVersionStatus.
func (in *SimulationApplicationVersionStatus) DeepCopy() *SimulationApplicationVersionStatus {
	if in == nil {
		return nil
	}
	out := new(SimulationApplicationVersionStatus)
	in.DeepCopyInto(out)
	return out
}