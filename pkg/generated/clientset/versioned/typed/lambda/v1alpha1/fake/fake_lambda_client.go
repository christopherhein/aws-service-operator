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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "awsoperator.io/pkg/generated/clientset/versioned/typed/lambda/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLambdaV1alpha1 struct {
	*testing.Fake
}

func (c *FakeLambdaV1alpha1) Aliases(namespace string) v1alpha1.AliasInterface {
	return &FakeAliases{c, namespace}
}

func (c *FakeLambdaV1alpha1) EventSourceMappings(namespace string) v1alpha1.EventSourceMappingInterface {
	return &FakeEventSourceMappings{c, namespace}
}

func (c *FakeLambdaV1alpha1) Functions(namespace string) v1alpha1.FunctionInterface {
	return &FakeFunctions{c, namespace}
}

func (c *FakeLambdaV1alpha1) LayerVersions(namespace string) v1alpha1.LayerVersionInterface {
	return &FakeLayerVersions{c, namespace}
}

func (c *FakeLambdaV1alpha1) LayerVersionPermissions(namespace string) v1alpha1.LayerVersionPermissionInterface {
	return &FakeLayerVersionPermissions{c, namespace}
}

func (c *FakeLambdaV1alpha1) Permissions(namespace string) v1alpha1.PermissionInterface {
	return &FakePermissions{c, namespace}
}

func (c *FakeLambdaV1alpha1) Versions(namespace string) v1alpha1.VersionInterface {
	return &FakeVersions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLambdaV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}