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

// Subscription is a specification for a Subscription resource
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec SubscriptionSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status SubscriptionStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs SubscriptionOutputs `json:"outputs"`
}

// SubscriptionSpec is the spec for the Subscription resource
type SubscriptionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DeliveryPolicy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-deliverypolicy
	DeliveryPolicy SubscriptionSpecDeliveryPolicy `json:"deliveryPolicy,omitempty" cloudformation:"DeliveryPolicy,Parameter"`

	// Endpoint http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-endpoint
	Endpoint string `json:"endpoint,omitempty" cloudformation:"Endpoint,Parameter"`

	// FilterPolicy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-filterpolicy
	FilterPolicy SubscriptionSpecFilterPolicy `json:"filterPolicy,omitempty" cloudformation:"FilterPolicy,Parameter"`

	// Protocol http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-protocol
	Protocol string `json:"protocol,omitempty" cloudformation:"Protocol,Parameter"`

	// RawMessageDelivery http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-rawmessagedelivery
	RawMessageDelivery bool `json:"rawMessageDelivery,omitempty" cloudformation:"RawMessageDelivery,Parameter"`

	// Region http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-region
	Region string `json:"region,omitempty" cloudformation:"Region,Parameter"`

	// TopicRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#topicarn
	TopicRef string `json:"topicRef,omitempty" cloudformation:"TopicArn,Parameter"`
}

// SubscriptionSpecDeliveryPolicy is the definition for a Subscription resource
type SubscriptionSpecDeliveryPolicy struct {
}

// SubscriptionSpecFilterPolicy is the definition for a Subscription resource
type SubscriptionSpecFilterPolicy struct {
}

// SubscriptionStatus is the status for a  resource
type SubscriptionStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// SubscriptionOutputs is the status for a  resource
type SubscriptionOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SubscriptionList is a list of Subscription resources
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Subscription `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Topic is a specification for a Topic resource
type Topic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec TopicSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status TopicStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs TopicOutputs `json:"outputs"`
}

// TopicSpec is the spec for the Topic resource
type TopicSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DisplayName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-displayname
	DisplayName string `json:"displayName,omitempty" cloudformation:"DisplayName,Parameter"`

	// KmsMasterKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-kmsmasterkeyid
	KmsMasterKeyRef string `json:"kmsMasterKeyRef,omitempty" cloudformation:"KmsMasterKeyId,Parameter"`

	// Subscription http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-subscription
	Subscription []TopicSpecSubscription `json:"subscription,omitempty" cloudformation:"Subscription,Parameter"`

	// TopicName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-topicname
	TopicName string `json:"topicName,omitempty" cloudformation:"TopicName,Parameter"`
}

// TopicSpecSubscription is the definition for a Topic resource
type TopicSpecSubscription struct {
}

// TopicStatus is the status for a  resource
type TopicStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// TopicOutputs is the status for a  resource
type TopicOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TopicList is a list of Topic resources
type TopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Topic `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TopicPolicy is a specification for a TopicPolicy resource
type TopicPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec TopicPolicySpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status TopicPolicyStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs TopicPolicyOutputs `json:"outputs"`
}

// TopicPolicySpec is the spec for the TopicPolicy resource
type TopicPolicySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// PolicyDocument http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-policydocument
	PolicyDocument TopicPolicySpecPolicyDocument `json:"policyDocument,omitempty" cloudformation:"PolicyDocument,Parameter"`

	// Topics http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-topics
	Topics []string `json:"topics,omitempty" cloudformation:"Topics,Parameter"`
}

// TopicPolicySpecPolicyDocument is the definition for a TopicPolicy resource
type TopicPolicySpecPolicyDocument struct {
}

// TopicPolicyStatus is the status for a  resource
type TopicPolicyStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// TopicPolicyOutputs is the status for a  resource
type TopicPolicyOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TopicPolicyList is a list of TopicPolicy resources
type TopicPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TopicPolicy `json:"items"`
}