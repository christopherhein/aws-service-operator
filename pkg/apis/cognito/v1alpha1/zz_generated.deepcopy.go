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
func (in *IdentityPool) DeepCopyInto(out *IdentityPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPool.
func (in *IdentityPool) DeepCopy() *IdentityPool {
	if in == nil {
		return nil
	}
	out := new(IdentityPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolList) DeepCopyInto(out *IdentityPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolList.
func (in *IdentityPoolList) DeepCopy() *IdentityPoolList {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolOutputs) DeepCopyInto(out *IdentityPoolOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolOutputs.
func (in *IdentityPoolOutputs) DeepCopy() *IdentityPoolOutputs {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolRoleAttachment) DeepCopyInto(out *IdentityPoolRoleAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolRoleAttachment.
func (in *IdentityPoolRoleAttachment) DeepCopy() *IdentityPoolRoleAttachment {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolRoleAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPoolRoleAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolRoleAttachmentList) DeepCopyInto(out *IdentityPoolRoleAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IdentityPoolRoleAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolRoleAttachmentList.
func (in *IdentityPoolRoleAttachmentList) DeepCopy() *IdentityPoolRoleAttachmentList {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolRoleAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IdentityPoolRoleAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolRoleAttachmentOutputs) DeepCopyInto(out *IdentityPoolRoleAttachmentOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolRoleAttachmentOutputs.
func (in *IdentityPoolRoleAttachmentOutputs) DeepCopy() *IdentityPoolRoleAttachmentOutputs {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolRoleAttachmentOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolRoleAttachmentSpec) DeepCopyInto(out *IdentityPoolRoleAttachmentSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.RoleMappings = in.RoleMappings
	out.Roles = in.Roles
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolRoleAttachmentSpec.
func (in *IdentityPoolRoleAttachmentSpec) DeepCopy() *IdentityPoolRoleAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolRoleAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolRoleAttachmentSpecRoleMappings) DeepCopyInto(out *IdentityPoolRoleAttachmentSpecRoleMappings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolRoleAttachmentSpecRoleMappings.
func (in *IdentityPoolRoleAttachmentSpecRoleMappings) DeepCopy() *IdentityPoolRoleAttachmentSpecRoleMappings {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolRoleAttachmentSpecRoleMappings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolRoleAttachmentSpecRoles) DeepCopyInto(out *IdentityPoolRoleAttachmentSpecRoles) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolRoleAttachmentSpecRoles.
func (in *IdentityPoolRoleAttachmentSpecRoles) DeepCopy() *IdentityPoolRoleAttachmentSpecRoles {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolRoleAttachmentSpecRoles)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolRoleAttachmentStatus) DeepCopyInto(out *IdentityPoolRoleAttachmentStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolRoleAttachmentStatus.
func (in *IdentityPoolRoleAttachmentStatus) DeepCopy() *IdentityPoolRoleAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolRoleAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolSpec) DeepCopyInto(out *IdentityPoolSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.CognitoEvents = in.CognitoEvents
	if in.CognitoentityProverRef != nil {
		in, out := &in.CognitoentityProverRef, &out.CognitoentityProverRef
		*out = make([]IdentityPoolSpecCognitoentityProverRef, len(*in))
		copy(*out, *in)
	}
	out.CognitoStreams = in.CognitoStreams
	if in.OpenConnectProverARNRef != nil {
		in, out := &in.OpenConnectProverARNRef, &out.OpenConnectProverARNRef
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.PushSync.DeepCopyInto(&out.PushSync)
	if in.SamlProviderARNs != nil {
		in, out := &in.SamlProviderARNs, &out.SamlProviderARNs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.SupportedLoginProviders = in.SupportedLoginProviders
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolSpec.
func (in *IdentityPoolSpec) DeepCopy() *IdentityPoolSpec {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolSpecCognitoEvents) DeepCopyInto(out *IdentityPoolSpecCognitoEvents) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolSpecCognitoEvents.
func (in *IdentityPoolSpecCognitoEvents) DeepCopy() *IdentityPoolSpecCognitoEvents {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolSpecCognitoEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolSpecCognitoStreams) DeepCopyInto(out *IdentityPoolSpecCognitoStreams) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolSpecCognitoStreams.
func (in *IdentityPoolSpecCognitoStreams) DeepCopy() *IdentityPoolSpecCognitoStreams {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolSpecCognitoStreams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolSpecCognitoentityProverRef) DeepCopyInto(out *IdentityPoolSpecCognitoentityProverRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolSpecCognitoentityProverRef.
func (in *IdentityPoolSpecCognitoentityProverRef) DeepCopy() *IdentityPoolSpecCognitoentityProverRef {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolSpecCognitoentityProverRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolSpecPushSync) DeepCopyInto(out *IdentityPoolSpecPushSync) {
	*out = *in
	if in.ApplicationRef != nil {
		in, out := &in.ApplicationRef, &out.ApplicationRef
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolSpecPushSync.
func (in *IdentityPoolSpecPushSync) DeepCopy() *IdentityPoolSpecPushSync {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolSpecPushSync)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolSpecSupportedLoginProviders) DeepCopyInto(out *IdentityPoolSpecSupportedLoginProviders) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolSpecSupportedLoginProviders.
func (in *IdentityPoolSpecSupportedLoginProviders) DeepCopy() *IdentityPoolSpecSupportedLoginProviders {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolSpecSupportedLoginProviders)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IdentityPoolStatus) DeepCopyInto(out *IdentityPoolStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IdentityPoolStatus.
func (in *IdentityPoolStatus) DeepCopy() *IdentityPoolStatus {
	if in == nil {
		return nil
	}
	out := new(IdentityPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPool) DeepCopyInto(out *UserPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPool.
func (in *UserPool) DeepCopy() *UserPool {
	if in == nil {
		return nil
	}
	out := new(UserPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClient) DeepCopyInto(out *UserPoolClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClient.
func (in *UserPoolClient) DeepCopy() *UserPoolClient {
	if in == nil {
		return nil
	}
	out := new(UserPoolClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientList) DeepCopyInto(out *UserPoolClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserPoolClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientList.
func (in *UserPoolClientList) DeepCopy() *UserPoolClientList {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientOutputs) DeepCopyInto(out *UserPoolClientOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientOutputs.
func (in *UserPoolClientOutputs) DeepCopy() *UserPoolClientOutputs {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientSpec) DeepCopyInto(out *UserPoolClientSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.ExplicitAuthFlows != nil {
		in, out := &in.ExplicitAuthFlows, &out.ExplicitAuthFlows
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ReadAttributes != nil {
		in, out := &in.ReadAttributes, &out.ReadAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WriteAttributes != nil {
		in, out := &in.WriteAttributes, &out.WriteAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientSpec.
func (in *UserPoolClientSpec) DeepCopy() *UserPoolClientSpec {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolClientStatus) DeepCopyInto(out *UserPoolClientStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolClientStatus.
func (in *UserPoolClientStatus) DeepCopy() *UserPoolClientStatus {
	if in == nil {
		return nil
	}
	out := new(UserPoolClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolGroup) DeepCopyInto(out *UserPoolGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolGroup.
func (in *UserPoolGroup) DeepCopy() *UserPoolGroup {
	if in == nil {
		return nil
	}
	out := new(UserPoolGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolGroupList) DeepCopyInto(out *UserPoolGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserPoolGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolGroupList.
func (in *UserPoolGroupList) DeepCopy() *UserPoolGroupList {
	if in == nil {
		return nil
	}
	out := new(UserPoolGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolGroupOutputs) DeepCopyInto(out *UserPoolGroupOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolGroupOutputs.
func (in *UserPoolGroupOutputs) DeepCopy() *UserPoolGroupOutputs {
	if in == nil {
		return nil
	}
	out := new(UserPoolGroupOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolGroupSpec) DeepCopyInto(out *UserPoolGroupSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolGroupSpec.
func (in *UserPoolGroupSpec) DeepCopy() *UserPoolGroupSpec {
	if in == nil {
		return nil
	}
	out := new(UserPoolGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolGroupStatus) DeepCopyInto(out *UserPoolGroupStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolGroupStatus.
func (in *UserPoolGroupStatus) DeepCopy() *UserPoolGroupStatus {
	if in == nil {
		return nil
	}
	out := new(UserPoolGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolList) DeepCopyInto(out *UserPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolList.
func (in *UserPoolList) DeepCopy() *UserPoolList {
	if in == nil {
		return nil
	}
	out := new(UserPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolOutputs) DeepCopyInto(out *UserPoolOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolOutputs.
func (in *UserPoolOutputs) DeepCopy() *UserPoolOutputs {
	if in == nil {
		return nil
	}
	out := new(UserPoolOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpec) DeepCopyInto(out *UserPoolSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.AdminCreateUserConfig = in.AdminCreateUserConfig
	if in.AliasAttributes != nil {
		in, out := &in.AliasAttributes, &out.AliasAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AutoVerifiedAttributes != nil {
		in, out := &in.AutoVerifiedAttributes, &out.AutoVerifiedAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.DeviceConfiguration = in.DeviceConfiguration
	out.EmailConfiguration = in.EmailConfiguration
	out.LambdaConfig = in.LambdaConfig
	out.Policies = in.Policies
	if in.Schema != nil {
		in, out := &in.Schema, &out.Schema
		*out = make([]UserPoolSpecSchema, len(*in))
		copy(*out, *in)
	}
	out.SmsConfiguration = in.SmsConfiguration
	out.UserPoolTags = in.UserPoolTags
	if in.UsernameAttributes != nil {
		in, out := &in.UsernameAttributes, &out.UsernameAttributes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpec.
func (in *UserPoolSpec) DeepCopy() *UserPoolSpec {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecAdminCreateUserConfig) DeepCopyInto(out *UserPoolSpecAdminCreateUserConfig) {
	*out = *in
	out.InviteMessageTemplate = in.InviteMessageTemplate
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecAdminCreateUserConfig.
func (in *UserPoolSpecAdminCreateUserConfig) DeepCopy() *UserPoolSpecAdminCreateUserConfig {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecAdminCreateUserConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecAdminCreateUserConfigInviteMessageTemplate) DeepCopyInto(out *UserPoolSpecAdminCreateUserConfigInviteMessageTemplate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecAdminCreateUserConfigInviteMessageTemplate.
func (in *UserPoolSpecAdminCreateUserConfigInviteMessageTemplate) DeepCopy() *UserPoolSpecAdminCreateUserConfigInviteMessageTemplate {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecAdminCreateUserConfigInviteMessageTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecDeviceConfiguration) DeepCopyInto(out *UserPoolSpecDeviceConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecDeviceConfiguration.
func (in *UserPoolSpecDeviceConfiguration) DeepCopy() *UserPoolSpecDeviceConfiguration {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecDeviceConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecEmailConfiguration) DeepCopyInto(out *UserPoolSpecEmailConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecEmailConfiguration.
func (in *UserPoolSpecEmailConfiguration) DeepCopy() *UserPoolSpecEmailConfiguration {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecEmailConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecLambdaConfig) DeepCopyInto(out *UserPoolSpecLambdaConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecLambdaConfig.
func (in *UserPoolSpecLambdaConfig) DeepCopy() *UserPoolSpecLambdaConfig {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecLambdaConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecPolicies) DeepCopyInto(out *UserPoolSpecPolicies) {
	*out = *in
	out.PasswordPolicy = in.PasswordPolicy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecPolicies.
func (in *UserPoolSpecPolicies) DeepCopy() *UserPoolSpecPolicies {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecPolicies)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecPoliciesPasswordPolicy) DeepCopyInto(out *UserPoolSpecPoliciesPasswordPolicy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecPoliciesPasswordPolicy.
func (in *UserPoolSpecPoliciesPasswordPolicy) DeepCopy() *UserPoolSpecPoliciesPasswordPolicy {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecPoliciesPasswordPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecSchema) DeepCopyInto(out *UserPoolSpecSchema) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecSchema.
func (in *UserPoolSpecSchema) DeepCopy() *UserPoolSpecSchema {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecSchema)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecSchemaNumberAttributeConstraints) DeepCopyInto(out *UserPoolSpecSchemaNumberAttributeConstraints) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecSchemaNumberAttributeConstraints.
func (in *UserPoolSpecSchemaNumberAttributeConstraints) DeepCopy() *UserPoolSpecSchemaNumberAttributeConstraints {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecSchemaNumberAttributeConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecSchemaStringAttributeConstraints) DeepCopyInto(out *UserPoolSpecSchemaStringAttributeConstraints) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecSchemaStringAttributeConstraints.
func (in *UserPoolSpecSchemaStringAttributeConstraints) DeepCopy() *UserPoolSpecSchemaStringAttributeConstraints {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecSchemaStringAttributeConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecSmsConfiguration) DeepCopyInto(out *UserPoolSpecSmsConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecSmsConfiguration.
func (in *UserPoolSpecSmsConfiguration) DeepCopy() *UserPoolSpecSmsConfiguration {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecSmsConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolSpecUserPoolTags) DeepCopyInto(out *UserPoolSpecUserPoolTags) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolSpecUserPoolTags.
func (in *UserPoolSpecUserPoolTags) DeepCopy() *UserPoolSpecUserPoolTags {
	if in == nil {
		return nil
	}
	out := new(UserPoolSpecUserPoolTags)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolStatus) DeepCopyInto(out *UserPoolStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolStatus.
func (in *UserPoolStatus) DeepCopy() *UserPoolStatus {
	if in == nil {
		return nil
	}
	out := new(UserPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUser) DeepCopyInto(out *UserPoolUser) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUser.
func (in *UserPoolUser) DeepCopy() *UserPoolUser {
	if in == nil {
		return nil
	}
	out := new(UserPoolUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolUser) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserList) DeepCopyInto(out *UserPoolUserList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserPoolUser, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserList.
func (in *UserPoolUserList) DeepCopy() *UserPoolUserList {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolUserList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserOutputs) DeepCopyInto(out *UserPoolUserOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserOutputs.
func (in *UserPoolUserOutputs) DeepCopy() *UserPoolUserOutputs {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserSpec) DeepCopyInto(out *UserPoolUserSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.DesiredDeliveryMediums != nil {
		in, out := &in.DesiredDeliveryMediums, &out.DesiredDeliveryMediums
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.UserAttributes != nil {
		in, out := &in.UserAttributes, &out.UserAttributes
		*out = make([]UserPoolUserSpecUserAttribute, len(*in))
		copy(*out, *in)
	}
	if in.ValidationData != nil {
		in, out := &in.ValidationData, &out.ValidationData
		*out = make([]UserPoolUserSpecValidationDatum, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserSpec.
func (in *UserPoolUserSpec) DeepCopy() *UserPoolUserSpec {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserSpecUserAttribute) DeepCopyInto(out *UserPoolUserSpecUserAttribute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserSpecUserAttribute.
func (in *UserPoolUserSpecUserAttribute) DeepCopy() *UserPoolUserSpecUserAttribute {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserSpecUserAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserSpecValidationDatum) DeepCopyInto(out *UserPoolUserSpecValidationDatum) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserSpecValidationDatum.
func (in *UserPoolUserSpecValidationDatum) DeepCopy() *UserPoolUserSpecValidationDatum {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserSpecValidationDatum)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserStatus) DeepCopyInto(out *UserPoolUserStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserStatus.
func (in *UserPoolUserStatus) DeepCopy() *UserPoolUserStatus {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserToGroupAttachment) DeepCopyInto(out *UserPoolUserToGroupAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserToGroupAttachment.
func (in *UserPoolUserToGroupAttachment) DeepCopy() *UserPoolUserToGroupAttachment {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserToGroupAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolUserToGroupAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserToGroupAttachmentList) DeepCopyInto(out *UserPoolUserToGroupAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UserPoolUserToGroupAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserToGroupAttachmentList.
func (in *UserPoolUserToGroupAttachmentList) DeepCopy() *UserPoolUserToGroupAttachmentList {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserToGroupAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UserPoolUserToGroupAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserToGroupAttachmentOutputs) DeepCopyInto(out *UserPoolUserToGroupAttachmentOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserToGroupAttachmentOutputs.
func (in *UserPoolUserToGroupAttachmentOutputs) DeepCopy() *UserPoolUserToGroupAttachmentOutputs {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserToGroupAttachmentOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserToGroupAttachmentSpec) DeepCopyInto(out *UserPoolUserToGroupAttachmentSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserToGroupAttachmentSpec.
func (in *UserPoolUserToGroupAttachmentSpec) DeepCopy() *UserPoolUserToGroupAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserToGroupAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserPoolUserToGroupAttachmentStatus) DeepCopyInto(out *UserPoolUserToGroupAttachmentStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserPoolUserToGroupAttachmentStatus.
func (in *UserPoolUserToGroupAttachmentStatus) DeepCopy() *UserPoolUserToGroupAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(UserPoolUserToGroupAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}