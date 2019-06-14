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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// IdentityPools returns a IdentityPoolInformer.
	IdentityPools() IdentityPoolInformer
	// IdentityPoolRoleAttachments returns a IdentityPoolRoleAttachmentInformer.
	IdentityPoolRoleAttachments() IdentityPoolRoleAttachmentInformer
	// UserPools returns a UserPoolInformer.
	UserPools() UserPoolInformer
	// UserPoolClients returns a UserPoolClientInformer.
	UserPoolClients() UserPoolClientInformer
	// UserPoolGroups returns a UserPoolGroupInformer.
	UserPoolGroups() UserPoolGroupInformer
	// UserPoolUsers returns a UserPoolUserInformer.
	UserPoolUsers() UserPoolUserInformer
	// UserPoolUserToGroupAttachments returns a UserPoolUserToGroupAttachmentInformer.
	UserPoolUserToGroupAttachments() UserPoolUserToGroupAttachmentInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// IdentityPools returns a IdentityPoolInformer.
func (v *version) IdentityPools() IdentityPoolInformer {
	return &identityPoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// IdentityPoolRoleAttachments returns a IdentityPoolRoleAttachmentInformer.
func (v *version) IdentityPoolRoleAttachments() IdentityPoolRoleAttachmentInformer {
	return &identityPoolRoleAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UserPools returns a UserPoolInformer.
func (v *version) UserPools() UserPoolInformer {
	return &userPoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UserPoolClients returns a UserPoolClientInformer.
func (v *version) UserPoolClients() UserPoolClientInformer {
	return &userPoolClientInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UserPoolGroups returns a UserPoolGroupInformer.
func (v *version) UserPoolGroups() UserPoolGroupInformer {
	return &userPoolGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UserPoolUsers returns a UserPoolUserInformer.
func (v *version) UserPoolUsers() UserPoolUserInformer {
	return &userPoolUserInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UserPoolUserToGroupAttachments returns a UserPoolUserToGroupAttachmentInformer.
func (v *version) UserPoolUserToGroupAttachments() UserPoolUserToGroupAttachmentInformer {
	return &userPoolUserToGroupAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}