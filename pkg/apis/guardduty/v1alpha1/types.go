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

// Detector is a specification for a Detector resource
type Detector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec DetectorSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status DetectorStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs DetectorOutputs `json:"outputs"`
}

// DetectorSpec is the spec for the Detector resource
type DetectorSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Enable http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-enable
	Enable bool `json:"enable,omitempty" cloudformation:"Enable,Parameter"`

	// FindingPublishingFrequency http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-detector.html#cfn-guardduty-detector-findingpublishingfrequency
	FindingPublishingFrequency string `json:"findingPublishingFrequency,omitempty" cloudformation:"FindingPublishingFrequency,Parameter"`
}

// DetectorStatus is the status for a  resource
type DetectorStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// DetectorOutputs is the status for a  resource
type DetectorOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DetectorList is a list of Detector resources
type DetectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Detector `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Filter is a specification for a Filter resource
type Filter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec FilterSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status FilterStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs FilterOutputs `json:"outputs"`
}

// FilterSpec is the spec for the Filter resource
type FilterSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Action http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-action
	Action string `json:"action,omitempty" cloudformation:"Action,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// DetectorRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-detectorid
	DetectorRef string `json:"detectorRef,omitempty" cloudformation:"DetectorId,Parameter"`

	// FindingCriteria http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-findingcriteria
	FindingCriteria FilterSpecFindingCriteria `json:"findingCriteria,omitempty"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Rank http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-filter.html#cfn-guardduty-filter-rank
	Rank int32 `json:"rank,omitempty" cloudformation:"Rank,Parameter"`
}

// FilterSpecFindingCriteria is the definition for a Filter resource
type FilterSpecFindingCriteria struct {
	// Criterion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-findingcriteria.html#cfn-guardduty-filter-findingcriteria-criterion
	Criterion FilterSpecFindingCriteriaCriterion `json:"criterion,omitempty" cloudformation:"Criterion,Parameter"`

	// ItemType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-findingcriteria.html#cfn-guardduty-filter-findingcriteria-itemtype
	ItemType FilterSpecFindingCriteriaItemType `json:"itemType,omitempty"`
}

// FilterSpecFindingCriteriaCriterion is the definition for a Filter resource
type FilterSpecFindingCriteriaCriterion struct {
}

// FilterSpecFindingCriteriaItemType is the definition for a Filter resource
type FilterSpecFindingCriteriaItemType struct {
	// Eq http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-condition.html#cfn-guardduty-filter-condition-eq
	Eq []string `json:"eq,omitempty" cloudformation:"Eq,Parameter"`

	// Gte http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-condition.html#cfn-guardduty-filter-condition-gte
	Gte int32 `json:"gte,omitempty" cloudformation:"Gte,Parameter"`

	// Lt http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-condition.html#cfn-guardduty-filter-condition-lt
	Lt int32 `json:"lt,omitempty" cloudformation:"Lt,Parameter"`

	// Lte http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-condition.html#cfn-guardduty-filter-condition-lte
	Lte int32 `json:"lte,omitempty" cloudformation:"Lte,Parameter"`

	// Neq http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-guardduty-filter-condition.html#cfn-guardduty-filter-condition-neq
	Neq []string `json:"neq,omitempty" cloudformation:"Neq,Parameter"`
}

// FilterStatus is the status for a  resource
type FilterStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// FilterOutputs is the status for a  resource
type FilterOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FilterList is a list of Filter resources
type FilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Filter `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IPSet is a specification for a IPSet resource
type IPSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec IPSetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status IPSetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs IPSetOutputs `json:"outputs"`
}

// IPSetSpec is the spec for the IPSet resource
type IPSetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Activate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-activate
	Activate bool `json:"activate,omitempty" cloudformation:"Activate,Parameter"`

	// DetectorRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-detectorid
	DetectorRef string `json:"detectorRef,omitempty" cloudformation:"DetectorId,Parameter"`

	// Format http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-format
	Format string `json:"format,omitempty" cloudformation:"Format,Parameter"`

	// Location http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-location
	Location string `json:"location,omitempty" cloudformation:"Location,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-ipset.html#cfn-guardduty-ipset-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`
}

// IPSetStatus is the status for a  resource
type IPSetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// IPSetOutputs is the status for a  resource
type IPSetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IPSetList is a list of IPSet resources
type IPSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []IPSet `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Master is a specification for a Master resource
type Master struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec MasterSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status MasterStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs MasterOutputs `json:"outputs"`
}

// MasterSpec is the spec for the Master resource
type MasterSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DetectorRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-detectorid
	DetectorRef string `json:"detectorRef,omitempty" cloudformation:"DetectorId,Parameter"`

	// InvitationRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-invitationid
	InvitationRef string `json:"invitationRef,omitempty" cloudformation:"InvitationId,Parameter"`

	// MasterRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-master.html#cfn-guardduty-master-masterid
	MasterRef string `json:"masterRef,omitempty" cloudformation:"MasterId,Parameter"`
}

// MasterStatus is the status for a  resource
type MasterStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// MasterOutputs is the status for a  resource
type MasterOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MasterList is a list of Master resources
type MasterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Master `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Member is a specification for a Member resource
type Member struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec MemberSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status MemberStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs MemberOutputs `json:"outputs"`
}

// MemberSpec is the spec for the Member resource
type MemberSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DetectorRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-detectorid
	DetectorRef string `json:"detectorRef,omitempty" cloudformation:"DetectorId,Parameter"`

	// DisableEmailNotification http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-disableemailnotification
	DisableEmailNotification bool `json:"disableEmailNotification,omitempty" cloudformation:"DisableEmailNotification,Parameter"`

	// Email http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-email
	Email string `json:"email,omitempty" cloudformation:"Email,Parameter"`

	// MemberRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-memberid
	MemberRef string `json:"memberRef,omitempty" cloudformation:"MemberId,Parameter"`

	// Message http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-message
	Message string `json:"message,omitempty" cloudformation:"Message,Parameter"`

	// Status http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-member.html#cfn-guardduty-member-status
	Status string `json:"status,omitempty" cloudformation:"Status,Parameter"`
}

// MemberStatus is the status for a  resource
type MemberStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// MemberOutputs is the status for a  resource
type MemberOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MemberList is a list of Member resources
type MemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Member `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ThreatIntelSet is a specification for a ThreatIntelSet resource
type ThreatIntelSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ThreatIntelSetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ThreatIntelSetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ThreatIntelSetOutputs `json:"outputs"`
}

// ThreatIntelSetSpec is the spec for the ThreatIntelSet resource
type ThreatIntelSetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Activate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatintelset.html#cfn-guardduty-threatintelset-activate
	Activate bool `json:"activate,omitempty" cloudformation:"Activate,Parameter"`

	// DetectorRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatintelset.html#cfn-guardduty-threatintelset-detectorid
	DetectorRef string `json:"detectorRef,omitempty" cloudformation:"DetectorId,Parameter"`

	// Format http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatintelset.html#cfn-guardduty-threatintelset-format
	Format string `json:"format,omitempty" cloudformation:"Format,Parameter"`

	// Location http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatintelset.html#cfn-guardduty-threatintelset-location
	Location string `json:"location,omitempty" cloudformation:"Location,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-guardduty-threatintelset.html#cfn-guardduty-threatintelset-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`
}

// ThreatIntelSetStatus is the status for a  resource
type ThreatIntelSetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ThreatIntelSetOutputs is the status for a  resource
type ThreatIntelSetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ThreatIntelSetList is a list of ThreatIntelSet resources
type ThreatIntelSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ThreatIntelSet `json:"items"`
}