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
	// ConfigurationSets returns a ConfigurationSetInformer.
	ConfigurationSets() ConfigurationSetInformer
	// ConfigurationSetEventDestinations returns a ConfigurationSetEventDestinationInformer.
	ConfigurationSetEventDestinations() ConfigurationSetEventDestinationInformer
	// DedicatedIpPools returns a DedicatedIpPoolInformer.
	DedicatedIpPools() DedicatedIpPoolInformer
	// Identities returns a IdentityInformer.
	Identities() IdentityInformer
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

// ConfigurationSets returns a ConfigurationSetInformer.
func (v *version) ConfigurationSets() ConfigurationSetInformer {
	return &configurationSetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConfigurationSetEventDestinations returns a ConfigurationSetEventDestinationInformer.
func (v *version) ConfigurationSetEventDestinations() ConfigurationSetEventDestinationInformer {
	return &configurationSetEventDestinationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DedicatedIpPools returns a DedicatedIpPoolInformer.
func (v *version) DedicatedIpPools() DedicatedIpPoolInformer {
	return &dedicatedIpPoolInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Identities returns a IdentityInformer.
func (v *version) Identities() IdentityInformer {
	return &identityInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}