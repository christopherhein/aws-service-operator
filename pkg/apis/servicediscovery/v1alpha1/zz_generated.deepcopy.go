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
func (in *HttpNamespace) DeepCopyInto(out *HttpNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpNamespace.
func (in *HttpNamespace) DeepCopy() *HttpNamespace {
	if in == nil {
		return nil
	}
	out := new(HttpNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HttpNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpNamespaceList) DeepCopyInto(out *HttpNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HttpNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpNamespaceList.
func (in *HttpNamespaceList) DeepCopy() *HttpNamespaceList {
	if in == nil {
		return nil
	}
	out := new(HttpNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HttpNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpNamespaceOutputs) DeepCopyInto(out *HttpNamespaceOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpNamespaceOutputs.
func (in *HttpNamespaceOutputs) DeepCopy() *HttpNamespaceOutputs {
	if in == nil {
		return nil
	}
	out := new(HttpNamespaceOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpNamespaceSpec) DeepCopyInto(out *HttpNamespaceSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpNamespaceSpec.
func (in *HttpNamespaceSpec) DeepCopy() *HttpNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(HttpNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HttpNamespaceStatus) DeepCopyInto(out *HttpNamespaceStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HttpNamespaceStatus.
func (in *HttpNamespaceStatus) DeepCopy() *HttpNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(HttpNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceOutputs) DeepCopyInto(out *InstanceOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceOutputs.
func (in *InstanceOutputs) DeepCopy() *InstanceOutputs {
	if in == nil {
		return nil
	}
	out := new(InstanceOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.InstanceAttributes != nil {
		in, out := &in.InstanceAttributes, &out.InstanceAttributes
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsNamespace) DeepCopyInto(out *PrivateDnsNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsNamespace.
func (in *PrivateDnsNamespace) DeepCopy() *PrivateDnsNamespace {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateDnsNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsNamespaceList) DeepCopyInto(out *PrivateDnsNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrivateDnsNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsNamespaceList.
func (in *PrivateDnsNamespaceList) DeepCopy() *PrivateDnsNamespaceList {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrivateDnsNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsNamespaceOutputs) DeepCopyInto(out *PrivateDnsNamespaceOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsNamespaceOutputs.
func (in *PrivateDnsNamespaceOutputs) DeepCopy() *PrivateDnsNamespaceOutputs {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsNamespaceOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsNamespaceSpec) DeepCopyInto(out *PrivateDnsNamespaceSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsNamespaceSpec.
func (in *PrivateDnsNamespaceSpec) DeepCopy() *PrivateDnsNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateDnsNamespaceStatus) DeepCopyInto(out *PrivateDnsNamespaceStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateDnsNamespaceStatus.
func (in *PrivateDnsNamespaceStatus) DeepCopy() *PrivateDnsNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(PrivateDnsNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDnsNamespace) DeepCopyInto(out *PublicDnsNamespace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDnsNamespace.
func (in *PublicDnsNamespace) DeepCopy() *PublicDnsNamespace {
	if in == nil {
		return nil
	}
	out := new(PublicDnsNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicDnsNamespace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDnsNamespaceList) DeepCopyInto(out *PublicDnsNamespaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PublicDnsNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDnsNamespaceList.
func (in *PublicDnsNamespaceList) DeepCopy() *PublicDnsNamespaceList {
	if in == nil {
		return nil
	}
	out := new(PublicDnsNamespaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PublicDnsNamespaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDnsNamespaceOutputs) DeepCopyInto(out *PublicDnsNamespaceOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDnsNamespaceOutputs.
func (in *PublicDnsNamespaceOutputs) DeepCopy() *PublicDnsNamespaceOutputs {
	if in == nil {
		return nil
	}
	out := new(PublicDnsNamespaceOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDnsNamespaceSpec) DeepCopyInto(out *PublicDnsNamespaceSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDnsNamespaceSpec.
func (in *PublicDnsNamespaceSpec) DeepCopy() *PublicDnsNamespaceSpec {
	if in == nil {
		return nil
	}
	out := new(PublicDnsNamespaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicDnsNamespaceStatus) DeepCopyInto(out *PublicDnsNamespaceStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicDnsNamespaceStatus.
func (in *PublicDnsNamespaceStatus) DeepCopy() *PublicDnsNamespaceStatus {
	if in == nil {
		return nil
	}
	out := new(PublicDnsNamespaceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Service) DeepCopyInto(out *Service) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Service.
func (in *Service) DeepCopy() *Service {
	if in == nil {
		return nil
	}
	out := new(Service)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Service) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceList) DeepCopyInto(out *ServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Service, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceList.
func (in *ServiceList) DeepCopy() *ServiceList {
	if in == nil {
		return nil
	}
	out := new(ServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceOutputs) DeepCopyInto(out *ServiceOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceOutputs.
func (in *ServiceOutputs) DeepCopy() *ServiceOutputs {
	if in == nil {
		return nil
	}
	out := new(ServiceOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	in.DnsConfig.DeepCopyInto(&out.DnsConfig)
	out.HealthCheckConfig = in.HealthCheckConfig
	out.HealthCheckCustomConfig = in.HealthCheckCustomConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecDnsConfig) DeepCopyInto(out *ServiceSpecDnsConfig) {
	*out = *in
	if in.DnsRecords != nil {
		in, out := &in.DnsRecords, &out.DnsRecords
		*out = make([]ServiceSpecDnsConfigDnsRecord, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecDnsConfig.
func (in *ServiceSpecDnsConfig) DeepCopy() *ServiceSpecDnsConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecDnsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecDnsConfigDnsRecord) DeepCopyInto(out *ServiceSpecDnsConfigDnsRecord) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecDnsConfigDnsRecord.
func (in *ServiceSpecDnsConfigDnsRecord) DeepCopy() *ServiceSpecDnsConfigDnsRecord {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecDnsConfigDnsRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecHealthCheckConfig) DeepCopyInto(out *ServiceSpecHealthCheckConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecHealthCheckConfig.
func (in *ServiceSpecHealthCheckConfig) DeepCopy() *ServiceSpecHealthCheckConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecHealthCheckConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpecHealthCheckCustomConfig) DeepCopyInto(out *ServiceSpecHealthCheckCustomConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpecHealthCheckCustomConfig.
func (in *ServiceSpecHealthCheckCustomConfig) DeepCopy() *ServiceSpecHealthCheckCustomConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceSpecHealthCheckCustomConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}