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

// CustomActionType is a specification for a CustomActionType resource
type CustomActionType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec CustomActionTypeSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status CustomActionTypeStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs CustomActionTypeOutputs `json:"outputs"`
}

// CustomActionTypeSpec is the spec for the CustomActionType resource
type CustomActionTypeSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Category http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-category
	Category string `json:"category,omitempty" cloudformation:"Category,Parameter"`

	// ConfigurationProperties http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-configurationproperties
	ConfigurationProperties []CustomActionTypeSpecConfigurationProperty `json:"configurationProperties,omitempty" cloudformation:"ConfigurationProperties,Parameter"`

	// InputArtifactDetails http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-inputartifactdetails
	InputArtifactDetails CustomActionTypeSpecInputArtifactDetails `json:"inputArtifactDetails,omitempty"`

	// OutputArtifactDetails http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-outputartifactdetails
	OutputArtifactDetails CustomActionTypeSpecOutputArtifactDetails `json:"outputArtifactDetails,omitempty"`

	// Provider http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-provider
	Provider string `json:"provider,omitempty" cloudformation:"Provider,Parameter"`

	// Settings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-settings
	Settings CustomActionTypeSpecSettings `json:"settings,omitempty"`

	// Version http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-version
	Version string `json:"version,omitempty" cloudformation:"Version,Parameter"`
}

// CustomActionTypeSpecConfigurationProperty is the definition for a CustomActionType resource
type CustomActionTypeSpecConfigurationProperty struct {
}

// CustomActionTypeSpecInputArtifactDetails is the definition for a CustomActionType resource
type CustomActionTypeSpecInputArtifactDetails struct {
	// MaximumCount http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-maximumcount
	MaximumCount int32 `json:"maximumCount,omitempty" cloudformation:"MaximumCount,Parameter"`

	// MinimumCount http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-minimumcount
	MinimumCount int32 `json:"minimumCount,omitempty" cloudformation:"MinimumCount,Parameter"`
}

// CustomActionTypeSpecOutputArtifactDetails is the definition for a CustomActionType resource
type CustomActionTypeSpecOutputArtifactDetails struct {
	// MaximumCount http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-maximumcount
	MaximumCount int32 `json:"maximumCount,omitempty" cloudformation:"MaximumCount,Parameter"`

	// MinimumCount http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-minimumcount
	MinimumCount int32 `json:"minimumCount,omitempty" cloudformation:"MinimumCount,Parameter"`
}

// CustomActionTypeSpecSettings is the definition for a CustomActionType resource
type CustomActionTypeSpecSettings struct {
	// EntityUrlTemplate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-entityurltemplate
	EntityUrlTemplate string `json:"entityUrlTemplate,omitempty" cloudformation:"EntityUrlTemplate,Parameter"`

	// ExecutionUrlTemplate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-executionurltemplate
	ExecutionUrlTemplate string `json:"executionUrlTemplate,omitempty" cloudformation:"ExecutionUrlTemplate,Parameter"`

	// RevisionUrlTemplate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-revisionurltemplate
	RevisionUrlTemplate string `json:"revisionUrlTemplate,omitempty" cloudformation:"RevisionUrlTemplate,Parameter"`

	// ThirdPartyConfigurationUrl http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-thirdpartyconfigurationurl
	ThirdPartyConfigurationUrl string `json:"thirdPartyConfigurationUrl,omitempty" cloudformation:"ThirdPartyConfigurationUrl,Parameter"`
}

// CustomActionTypeStatus is the status for a  resource
type CustomActionTypeStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// CustomActionTypeOutputs is the status for a  resource
type CustomActionTypeOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomActionTypeList is a list of CustomActionType resources
type CustomActionTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CustomActionType `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Pipeline is a specification for a Pipeline resource
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec PipelineSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status PipelineStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs PipelineOutputs `json:"outputs"`
}

// PipelineSpec is the spec for the Pipeline resource
type PipelineSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ArtifactStore http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstore
	ArtifactStore PipelineSpecArtifactStore `json:"artifactStore,omitempty"`

	// ArtifactStores http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstores
	ArtifactStores []PipelineSpecArtifactStore `json:"artifactStores,omitempty" cloudformation:"ArtifactStores,Parameter"`

	// DisableInboundStageTransitions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-disableinboundstagetransitions
	DisableInboundStageTransitions []PipelineSpecDisableInboundStageTransition `json:"disableInboundStageTransitions,omitempty" cloudformation:"DisableInboundStageTransitions,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// RestartExecutionOnUpdate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-restartexecutiononupdate
	RestartExecutionOnUpdate bool `json:"restartExecutionOnUpdate,omitempty" cloudformation:"RestartExecutionOnUpdate,Parameter"`

	// RoleRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-rolearn
	RoleRef string `json:"roleRef,omitempty" cloudformation:"RoleArn,Parameter"`

	// Stages http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-stages
	Stages []PipelineSpecStage `json:"stages,omitempty" cloudformation:"Stages,Parameter"`
}

// PipelineSpecArtifactStore is the definition for a Pipeline resource
type PipelineSpecArtifactStore struct {
}

// PipelineSpecArtifactStoreEncryptionKey is the definition for a Pipeline resource
type PipelineSpecArtifactStoreEncryptionKey struct {
	// Id http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-id
	Id string `json:"id,omitempty" cloudformation:"Id,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// PipelineSpecArtifactStoreArtifactStore is the definition for a Pipeline resource
type PipelineSpecArtifactStoreArtifactStore struct {
	// EncryptionKey http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey
	EncryptionKey PipelineSpecArtifactStoreArtifactStoreEncryptionKey `json:"encryptionKey,omitempty"`

	// Location http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-location
	Location string `json:"location,omitempty" cloudformation:"Location,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// PipelineSpecArtifactStoreArtifactStoreEncryptionKey is the definition for a Pipeline resource
type PipelineSpecArtifactStoreArtifactStoreEncryptionKey struct {
	// Id http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-id
	Id string `json:"id,omitempty" cloudformation:"Id,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// PipelineSpecDisableInboundStageTransition is the definition for a Pipeline resource
type PipelineSpecDisableInboundStageTransition struct {
}

// PipelineSpecStage is the definition for a Pipeline resource
type PipelineSpecStage struct {
}

// PipelineSpecStageAction is the definition for a Pipeline resource
type PipelineSpecStageAction struct {
}

// PipelineSpecStageActionActionTypeRef is the definition for a Pipeline resource
type PipelineSpecStageActionActionTypeRef struct {
	// Category http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid-category
	Category string `json:"category,omitempty" cloudformation:"Category,Parameter"`

	// Owner http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid-owner
	Owner string `json:"owner,omitempty" cloudformation:"Owner,Parameter"`

	// Provider http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid-provider
	Provider string `json:"provider,omitempty" cloudformation:"Provider,Parameter"`

	// Version http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-actiontypeid.html#cfn-codepipeline-pipeline-stages-actions-actiontypeid-version
	Version string `json:"version,omitempty" cloudformation:"Version,Parameter"`
}

// PipelineSpecStageActionConfiguration is the definition for a Pipeline resource
type PipelineSpecStageActionConfiguration struct {
}

// PipelineSpecStageActionInputArtifact is the definition for a Pipeline resource
type PipelineSpecStageActionInputArtifact struct {
}

// PipelineSpecStageActionOutputArtifact is the definition for a Pipeline resource
type PipelineSpecStageActionOutputArtifact struct {
}

// PipelineSpecStageBlocker is the definition for a Pipeline resource
type PipelineSpecStageBlocker struct {
}

// PipelineStatus is the status for a  resource
type PipelineStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// PipelineOutputs is the status for a  resource
type PipelineOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PipelineList is a list of Pipeline resources
type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Pipeline `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Webhook is a specification for a Webhook resource
type Webhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec WebhookSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status WebhookStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs WebhookOutputs `json:"outputs"`
}

// WebhookSpec is the spec for the Webhook resource
type WebhookSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Authentication http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-authentication
	Authentication string `json:"authentication,omitempty" cloudformation:"Authentication,Parameter"`

	// AuthenticationConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-authenticationconfiguration
	AuthenticationConfiguration WebhookSpecAuthenticationConfiguration `json:"authenticationConfiguration,omitempty"`

	// Filters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-filters
	Filters []WebhookSpecFilter `json:"filters,omitempty" cloudformation:"Filters,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// RegisterWithThirdParty http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-registerwiththirdparty
	RegisterWithThirdParty bool `json:"registerWithThirdParty,omitempty" cloudformation:"RegisterWithThirdParty,Parameter"`

	// TargetAction http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetaction
	TargetAction string `json:"targetAction,omitempty" cloudformation:"TargetAction,Parameter"`

	// TargetPipeline http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetpipeline
	TargetPipeline string `json:"targetPipeline,omitempty" cloudformation:"TargetPipeline,Parameter"`

	// TargetPipelineVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-webhook.html#cfn-codepipeline-webhook-targetpipelineversion
	TargetPipelineVersion int32 `json:"targetPipelineVersion,omitempty" cloudformation:"TargetPipelineVersion,Parameter"`
}

// WebhookSpecAuthenticationConfiguration is the definition for a Webhook resource
type WebhookSpecAuthenticationConfiguration struct {
	// AllowedIPRange http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookauthconfiguration.html#cfn-codepipeline-webhook-webhookauthconfiguration-allowediprange
	AllowedIPRange string `json:"allowedIPRange,omitempty" cloudformation:"AllowedIPRange,Parameter"`

	// SecretToken http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-webhook-webhookauthconfiguration.html#cfn-codepipeline-webhook-webhookauthconfiguration-secrettoken
	SecretToken string `json:"secretToken,omitempty" cloudformation:"SecretToken,Parameter"`
}

// WebhookSpecFilter is the definition for a Webhook resource
type WebhookSpecFilter struct {
}

// WebhookStatus is the status for a  resource
type WebhookStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// WebhookOutputs is the status for a  resource
type WebhookOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebhookList is a list of Webhook resources
type WebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Webhook `json:"items"`
}