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

// Code generated by make generate. DO NOT EDIT.

// Package v1alpha1 is the v1alpha1 version of the awsoperator.io api.
package v1alpha1

import (
	metav1alpha1 "awsoperator.io/pkg/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FileSystem is a specification for a FileSystem resource
type FileSystem struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec FileSystemSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status FileSystemStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs FileSystemOutputs `json:"outputs"`
}

// FileSystemSpec is the spec for the FileSystem resource
type FileSystemSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Encrypted http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-encrypted
	Encrypted bool `json:"encrypted,omitempty" cloudformation:"Encrypted,Parameter"`

	// FileSystemTags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-filesystemtags
	FileSystemTags []FileSystemSpecFileSystemTag `json:"fileSystemTags,omitempty" cloudformation:"FileSystemTags,Parameter"`

	// KmsKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-kmskeyid
	KmsKeyRef string `json:"kmsKeyRef,omitempty" cloudformation:"KmsKeyId,Parameter"`

	// PerformanceMode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-efs-filesystem-performancemode
	PerformanceMode string `json:"performanceMode,omitempty" cloudformation:"PerformanceMode,Parameter"`

	// ProvisionedThroughputInMibps http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-elasticfilesystem-filesystem-provisionedthroughputinmibps
	ProvisionedThroughputInMibps float64 `json:"provisionedThroughputInMibps,omitempty" cloudformation:"ProvisionedThroughputInMibps,Parameter"`

	// ThroughputMode http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-filesystem.html#cfn-elasticfilesystem-filesystem-throughputmode
	ThroughputMode string `json:"throughputMode,omitempty" cloudformation:"ThroughputMode,Parameter"`
}

// FileSystemSpecFileSystemTag is the definition for a FileSystem resource
type FileSystemSpecFileSystemTag struct {
}

// FileSystemStatus is the status for a  resource
type FileSystemStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// FileSystemOutputs is the status for a  resource
type FileSystemOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FileSystemList is a list of FileSystem resources
type FileSystemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []FileSystem `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MountTarget is a specification for a MountTarget resource
type MountTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec MountTargetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status MountTargetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs MountTargetOutputs `json:"outputs"`
}

// MountTargetSpec is the spec for the MountTarget resource
type MountTargetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// FileSystemRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-filesystemid
	FileSystemRef string `json:"fileSystemRef,omitempty" cloudformation:"FileSystemId,Parameter"`

	// IpAddress http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-ipaddress
	IpAddress string `json:"ipAddress,omitempty" cloudformation:"IpAddress,Parameter"`

	// SecurityGroupsRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-securitygroups
	SecurityGroupsRef []string `json:"securityGroupsRef,omitempty" cloudformation:"SecurityGroups,Parameter"`

	// SubnetRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-efs-mounttarget.html#cfn-efs-mounttarget-subnetid
	SubnetRef string `json:"subnetRef,omitempty" cloudformation:"SubnetId,Parameter"`
}

// MountTargetStatus is the status for a  resource
type MountTargetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// MountTargetOutputs is the status for a  resource
type MountTargetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MountTargetList is a list of MountTarget resources
type MountTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []MountTarget `json:"items"`
}