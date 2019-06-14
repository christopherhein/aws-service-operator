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

// Application is a specification for a Application resource
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ApplicationSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ApplicationStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ApplicationOutputs `json:"outputs"`
}

// ApplicationSpec is the spec for the Application resource
type ApplicationSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ApplicationName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-name
	ApplicationName string `json:"applicationName,omitempty" cloudformation:"ApplicationName,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// ResourceLifecycleConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-resourcelifecycleconfig
	ResourceLifecycleConfig ApplicationSpecResourceLifecycleConfig `json:"resourceLifecycleConfig,omitempty"`
}

// ApplicationSpecResourceLifecycleConfig is the definition for a Application resource
type ApplicationSpecResourceLifecycleConfig struct {
	// ServiceRole http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationresourcelifecycleconfig.html#cfn-elasticbeanstalk-application-applicationresourcelifecycleconfig-servicerole
	ServiceRole string `json:"serviceRole,omitempty" cloudformation:"ServiceRole,Parameter"`

	// VersionLifecycleConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationresourcelifecycleconfig.html#cfn-elasticbeanstalk-application-applicationresourcelifecycleconfig-versionlifecycleconfig
	VersionLifecycleConfig ApplicationSpecResourceLifecycleConfigVersionLifecycleConfig `json:"versionLifecycleConfig,omitempty"`
}

// ApplicationSpecResourceLifecycleConfigVersionLifecycleConfig is the definition for a Application resource
type ApplicationSpecResourceLifecycleConfigVersionLifecycleConfig struct {
	// MaxAgeRule http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html#cfn-elasticbeanstalk-application-applicationversionlifecycleconfig-maxagerule
	MaxAgeRule ApplicationSpecResourceLifecycleConfigVersionLifecycleConfigMaxAgeRule `json:"maxAgeRule,omitempty"`

	// MaxCountRule http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-applicationversionlifecycleconfig.html#cfn-elasticbeanstalk-application-applicationversionlifecycleconfig-maxcountrule
	MaxCountRule ApplicationSpecResourceLifecycleConfigVersionLifecycleConfigMaxCountRule `json:"maxCountRule,omitempty"`
}

// ApplicationSpecResourceLifecycleConfigVersionLifecycleConfigMaxAgeRule is the definition for a Application resource
type ApplicationSpecResourceLifecycleConfigVersionLifecycleConfigMaxAgeRule struct {
	// DeleteSourceFromS3 http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxagerule.html#cfn-elasticbeanstalk-application-maxagerule-deletesourcefroms3
	DeleteSourceFromS3 bool `json:"deleteSourceFromS3,omitempty" cloudformation:"DeleteSourceFromS3,Parameter"`

	// Enabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxagerule.html#cfn-elasticbeanstalk-application-maxagerule-enabled
	Enabled bool `json:"enabled,omitempty" cloudformation:"Enabled,Parameter"`

	// MaxAgeInDays http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxagerule.html#cfn-elasticbeanstalk-application-maxagerule-maxageindays
	MaxAgeInDays int32 `json:"maxAgeInDays,omitempty" cloudformation:"MaxAgeInDays,Parameter"`
}

// ApplicationSpecResourceLifecycleConfigVersionLifecycleConfigMaxCountRule is the definition for a Application resource
type ApplicationSpecResourceLifecycleConfigVersionLifecycleConfigMaxCountRule struct {
	// DeleteSourceFromS3 http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-deletesourcefroms3
	DeleteSourceFromS3 bool `json:"deleteSourceFromS3,omitempty" cloudformation:"DeleteSourceFromS3,Parameter"`

	// Enabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-enabled
	Enabled bool `json:"enabled,omitempty" cloudformation:"Enabled,Parameter"`

	// MaxCount http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-maxcount
	MaxCount int32 `json:"maxCount,omitempty" cloudformation:"MaxCount,Parameter"`
}

// ApplicationStatus is the status for a  resource
type ApplicationStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ApplicationOutputs is the status for a  resource
type ApplicationOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApplicationList is a list of Application resources
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Application `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApplicationVersion is a specification for a ApplicationVersion resource
type ApplicationVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ApplicationVersionSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ApplicationVersionStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ApplicationVersionOutputs `json:"outputs"`
}

// ApplicationVersionSpec is the spec for the ApplicationVersion resource
type ApplicationVersionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ApplicationName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-applicationname
	ApplicationName string `json:"applicationName,omitempty" cloudformation:"ApplicationName,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// SourceBundle http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	SourceBundle ApplicationVersionSpecSourceBundle `json:"sourceBundle,omitempty"`
}

// ApplicationVersionSpecSourceBundle is the definition for a ApplicationVersion resource
type ApplicationVersionSpecSourceBundle struct {
	// S3Bucket http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html#cfn-beanstalk-sourcebundle-s3bucket
	S3Bucket string `json:"s3Bucket,omitempty" cloudformation:"S3Bucket,Parameter"`

	// S3Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-sourcebundle.html#cfn-beanstalk-sourcebundle-s3key
	S3Key string `json:"s3Key,omitempty" cloudformation:"S3Key,Parameter"`
}

// ApplicationVersionStatus is the status for a  resource
type ApplicationVersionStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ApplicationVersionOutputs is the status for a  resource
type ApplicationVersionOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ApplicationVersionList is a list of ApplicationVersion resources
type ApplicationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ApplicationVersion `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigurationTemplate is a specification for a ConfigurationTemplate resource
type ConfigurationTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ConfigurationTemplateSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ConfigurationTemplateStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ConfigurationTemplateOutputs `json:"outputs"`
}

// ConfigurationTemplateSpec is the spec for the ConfigurationTemplate resource
type ConfigurationTemplateSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ApplicationName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-applicationname
	ApplicationName string `json:"applicationName,omitempty" cloudformation:"ApplicationName,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// EnvironmentRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-environmentid
	EnvironmentRef string `json:"environmentRef,omitempty" cloudformation:"EnvironmentId,Parameter"`

	// OptionSettings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-optionsettings
	OptionSettings []ConfigurationTemplateSpecOptionSetting `json:"optionSettings,omitempty" cloudformation:"OptionSettings,Parameter"`

	// PlatformRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-platformarn
	PlatformRef string `json:"platformRef,omitempty" cloudformation:"PlatformArn,Parameter"`

	// SolutionStackName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-solutionstackname
	SolutionStackName string `json:"solutionStackName,omitempty" cloudformation:"SolutionStackName,Parameter"`

	// SourceConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticbeanstalk-configurationtemplate.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration
	SourceConfiguration ConfigurationTemplateSpecSourceConfiguration `json:"sourceConfiguration,omitempty"`
}

// ConfigurationTemplateSpecOptionSetting is the definition for a ConfigurationTemplate resource
type ConfigurationTemplateSpecOptionSetting struct {
}

// ConfigurationTemplateSpecSourceConfiguration is the definition for a ConfigurationTemplate resource
type ConfigurationTemplateSpecSourceConfiguration struct {
	// ApplicationName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration-applicationname
	ApplicationName string `json:"applicationName,omitempty" cloudformation:"ApplicationName,Parameter"`

	// TemplateName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-configurationtemplate-sourceconfiguration.html#cfn-elasticbeanstalk-configurationtemplate-sourceconfiguration-templatename
	TemplateName string `json:"templateName,omitempty" cloudformation:"TemplateName,Parameter"`
}

// ConfigurationTemplateStatus is the status for a  resource
type ConfigurationTemplateStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ConfigurationTemplateOutputs is the status for a  resource
type ConfigurationTemplateOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigurationTemplateList is a list of ConfigurationTemplate resources
type ConfigurationTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ConfigurationTemplate `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Environment is a specification for a Environment resource
type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec EnvironmentSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status EnvironmentStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs EnvironmentOutputs `json:"outputs"`
}

// EnvironmentSpec is the spec for the Environment resource
type EnvironmentSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ApplicationName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-applicationname
	ApplicationName string `json:"applicationName,omitempty" cloudformation:"ApplicationName,Parameter"`

	// CNAMEPrefix http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-cnameprefix
	CNAMEPrefix string `json:"cNAMEPrefix,omitempty" cloudformation:"CNAMEPrefix,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// EnvironmentName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-name
	EnvironmentName string `json:"environmentName,omitempty" cloudformation:"EnvironmentName,Parameter"`

	// OptionSettings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-optionsettings
	OptionSettings []EnvironmentSpecOptionSetting `json:"optionSettings,omitempty" cloudformation:"OptionSettings,Parameter"`

	// PlatformRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-platformarn
	PlatformRef string `json:"platformRef,omitempty" cloudformation:"PlatformArn,Parameter"`

	// SolutionStackName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-solutionstackname
	SolutionStackName string `json:"solutionStackName,omitempty" cloudformation:"SolutionStackName,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-elasticbeanstalk-environment-tags
	Tags []EnvironmentSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`

	// TemplateName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-templatename
	TemplateName string `json:"templateName,omitempty" cloudformation:"TemplateName,Parameter"`

	// Tier http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-tier
	Tier EnvironmentSpecTier `json:"tier,omitempty"`

	// VersionLabel http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-versionlabel
	VersionLabel string `json:"versionLabel,omitempty" cloudformation:"VersionLabel,Parameter"`
}

// EnvironmentSpecOptionSetting is the definition for a Environment resource
type EnvironmentSpecOptionSetting struct {
}

// EnvironmentSpecTag is the definition for a Environment resource
type EnvironmentSpecTag struct {
}

// EnvironmentSpecTier is the definition for a Environment resource
type EnvironmentSpecTier struct {
	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`

	// Version http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-version
	Version string `json:"version,omitempty" cloudformation:"Version,Parameter"`
}

// EnvironmentStatus is the status for a  resource
type EnvironmentStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// EnvironmentOutputs is the status for a  resource
type EnvironmentOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EnvironmentList is a list of Environment resources
type EnvironmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Environment `json:"items"`
}