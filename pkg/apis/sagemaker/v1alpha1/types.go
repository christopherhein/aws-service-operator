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

// Endpoint is a specification for a Endpoint resource
type Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec EndpointSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status EndpointStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs EndpointOutputs `json:"outputs"`
}

// EndpointSpec is the spec for the Endpoint resource
type EndpointSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// EndpointConfigName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-endpointconfigname
	EndpointConfigName string `json:"endpointConfigName,omitempty" cloudformation:"EndpointConfigName,Parameter"`

	// EndpointName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-endpointname
	EndpointName string `json:"endpointName,omitempty" cloudformation:"EndpointName,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpoint.html#cfn-sagemaker-endpoint-tags
	Tags []EndpointSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`
}

// EndpointSpecTag is the definition for a Endpoint resource
type EndpointSpecTag struct {
}

// EndpointStatus is the status for a  resource
type EndpointStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// EndpointOutputs is the status for a  resource
type EndpointOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EndpointList is a list of Endpoint resources
type EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Endpoint `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EndpointConfig is a specification for a EndpointConfig resource
type EndpointConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec EndpointConfigSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status EndpointConfigStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs EndpointConfigOutputs `json:"outputs"`
}

// EndpointConfigSpec is the spec for the EndpointConfig resource
type EndpointConfigSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// EndpointConfigName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-endpointconfigname
	EndpointConfigName string `json:"endpointConfigName,omitempty" cloudformation:"EndpointConfigName,Parameter"`

	// KmsKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-kmskeyid
	KmsKeyRef string `json:"kmsKeyRef,omitempty" cloudformation:"KmsKeyId,Parameter"`

	// ProductionVariants http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-productionvariants
	ProductionVariants []EndpointConfigSpecProductionVariant `json:"productionVariants,omitempty" cloudformation:"ProductionVariants,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-endpointconfig.html#cfn-sagemaker-endpointconfig-tags
	Tags []EndpointConfigSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`
}

// EndpointConfigSpecProductionVariant is the definition for a EndpointConfig resource
type EndpointConfigSpecProductionVariant struct {
}

// EndpointConfigSpecTag is the definition for a EndpointConfig resource
type EndpointConfigSpecTag struct {
}

// EndpointConfigStatus is the status for a  resource
type EndpointConfigStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// EndpointConfigOutputs is the status for a  resource
type EndpointConfigOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EndpointConfigList is a list of EndpointConfig resources
type EndpointConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EndpointConfig `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Model is a specification for a Model resource
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ModelSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ModelStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ModelOutputs `json:"outputs"`
}

// ModelSpec is the spec for the Model resource
type ModelSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Containers http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html#cfn-sagemaker-model-containers
	Containers []ModelSpecContainer `json:"containers,omitempty" cloudformation:"Containers,Parameter"`

	// ExecutionRoleRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html#cfn-sagemaker-model-executionrolearn
	ExecutionRoleRef string `json:"executionRoleRef,omitempty" cloudformation:"ExecutionRoleArn,Parameter"`

	// ModelName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html#cfn-sagemaker-model-modelname
	ModelName string `json:"modelName,omitempty" cloudformation:"ModelName,Parameter"`

	// PrimaryContainer http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html#cfn-sagemaker-model-primarycontainer
	PrimaryContainer ModelSpecPrimaryContainer `json:"primaryContainer,omitempty"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html#cfn-sagemaker-model-tags
	Tags []ModelSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`

	// VpcConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-model.html#cfn-sagemaker-model-vpcconfig
	VpcConfig ModelSpecVpcConfig `json:"vpcConfig,omitempty"`
}

// ModelSpecContainer is the definition for a Model resource
type ModelSpecContainer struct {
}

// ModelSpecContainerEnvironment is the definition for a Model resource
type ModelSpecContainerEnvironment struct {
}

// ModelSpecPrimaryContainer is the definition for a Model resource
type ModelSpecPrimaryContainer struct {
	// ContainerHostname http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition.html#cfn-sagemaker-model-containerdefinition-containerhostname
	ContainerHostname string `json:"containerHostname,omitempty" cloudformation:"ContainerHostname,Parameter"`

	// Environment http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition.html#cfn-sagemaker-model-containerdefinition-environment
	Environment ModelSpecPrimaryContainerEnvironment `json:"environment,omitempty" cloudformation:"Environment,Parameter"`

	// Image http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition.html#cfn-sagemaker-model-containerdefinition-image
	Image string `json:"image,omitempty" cloudformation:"Image,Parameter"`

	// ModelDataUrl http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition.html#cfn-sagemaker-model-containerdefinition-modeldataurl
	ModelDataUrl string `json:"modelDataUrl,omitempty" cloudformation:"ModelDataUrl,Parameter"`
}

// ModelSpecPrimaryContainerEnvironment is the definition for a Model resource
type ModelSpecPrimaryContainerEnvironment struct {
}

// ModelSpecTag is the definition for a Model resource
type ModelSpecTag struct {
}

// ModelSpecVpcConfig is the definition for a Model resource
type ModelSpecVpcConfig struct {
	// SecurityGroupRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-vpcconfig.html#cfn-sagemaker-model-vpcconfig-securitygroupids
	SecurityGroupRef []string `json:"securityGroupRef,omitempty" cloudformation:"SecurityGroupIds,Parameter"`

	// Subnets http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-vpcconfig.html#cfn-sagemaker-model-vpcconfig-subnets
	Subnets []string `json:"subnets,omitempty" cloudformation:"Subnets,Parameter"`
}

// ModelStatus is the status for a  resource
type ModelStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ModelOutputs is the status for a  resource
type ModelOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ModelList is a list of Model resources
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Model `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotebookInstance is a specification for a NotebookInstance resource
type NotebookInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec NotebookInstanceSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status NotebookInstanceStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs NotebookInstanceOutputs `json:"outputs"`
}

// NotebookInstanceSpec is the spec for the NotebookInstance resource
type NotebookInstanceSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DirectInternetAccess http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-directinternetaccess
	DirectInternetAccess string `json:"directInternetAccess,omitempty" cloudformation:"DirectInternetAccess,Parameter"`

	// InstanceType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-instancetype
	InstanceType string `json:"instanceType,omitempty" cloudformation:"InstanceType,Parameter"`

	// KmsKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-kmskeyid
	KmsKeyRef string `json:"kmsKeyRef,omitempty" cloudformation:"KmsKeyId,Parameter"`

	// LifecycleConfigName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-lifecycleconfigname
	LifecycleConfigName string `json:"lifecycleConfigName,omitempty" cloudformation:"LifecycleConfigName,Parameter"`

	// NotebookInstanceName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-notebookinstancename
	NotebookInstanceName string `json:"notebookInstanceName,omitempty" cloudformation:"NotebookInstanceName,Parameter"`

	// RoleRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-rolearn
	RoleRef string `json:"roleRef,omitempty" cloudformation:"RoleArn,Parameter"`

	// RootAccess http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-rootaccess
	RootAccess string `json:"rootAccess,omitempty" cloudformation:"RootAccess,Parameter"`

	// SecurityGroupRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-securitygroupids
	SecurityGroupRef []string `json:"securityGroupRef,omitempty" cloudformation:"SecurityGroupIds,Parameter"`

	// SubnetRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-subnetid
	SubnetRef string `json:"subnetRef,omitempty" cloudformation:"SubnetId,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-tags
	Tags []NotebookInstanceSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`

	// VolumeSizeInGB http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html#cfn-sagemaker-notebookinstance-volumesizeingb
	VolumeSizeInGB int32 `json:"volumeSizeInGB,omitempty" cloudformation:"VolumeSizeInGB,Parameter"`
}

// NotebookInstanceSpecTag is the definition for a NotebookInstance resource
type NotebookInstanceSpecTag struct {
}

// NotebookInstanceStatus is the status for a  resource
type NotebookInstanceStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// NotebookInstanceOutputs is the status for a  resource
type NotebookInstanceOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotebookInstanceList is a list of NotebookInstance resources
type NotebookInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NotebookInstance `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotebookInstanceLifecycleConfig is a specification for a NotebookInstanceLifecycleConfig resource
type NotebookInstanceLifecycleConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec NotebookInstanceLifecycleConfigSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status NotebookInstanceLifecycleConfigStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs NotebookInstanceLifecycleConfigOutputs `json:"outputs"`
}

// NotebookInstanceLifecycleConfigSpec is the spec for the NotebookInstanceLifecycleConfig resource
type NotebookInstanceLifecycleConfigSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// NotebookInstanceLifecycleConfigName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecycleconfigname
	NotebookInstanceLifecycleConfigName string `json:"notebookInstanceLifecycleConfigName,omitempty" cloudformation:"NotebookInstanceLifecycleConfigName,Parameter"`

	// OnCreate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-oncreate
	OnCreate []NotebookInstanceLifecycleConfigSpecOnCreate `json:"onCreate,omitempty" cloudformation:"OnCreate,Parameter"`

	// OnStart http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html#cfn-sagemaker-notebookinstancelifecycleconfig-onstart
	OnStart []NotebookInstanceLifecycleConfigSpecOnStart `json:"onStart,omitempty" cloudformation:"OnStart,Parameter"`
}

// NotebookInstanceLifecycleConfigSpecOnCreate is the definition for a NotebookInstanceLifecycleConfig resource
type NotebookInstanceLifecycleConfigSpecOnCreate struct {
}

// NotebookInstanceLifecycleConfigSpecOnStart is the definition for a NotebookInstanceLifecycleConfig resource
type NotebookInstanceLifecycleConfigSpecOnStart struct {
}

// NotebookInstanceLifecycleConfigStatus is the status for a  resource
type NotebookInstanceLifecycleConfigStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// NotebookInstanceLifecycleConfigOutputs is the status for a  resource
type NotebookInstanceLifecycleConfigOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NotebookInstanceLifecycleConfigList is a list of NotebookInstanceLifecycleConfig resources
type NotebookInstanceLifecycleConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []NotebookInstanceLifecycleConfig `json:"items"`
}