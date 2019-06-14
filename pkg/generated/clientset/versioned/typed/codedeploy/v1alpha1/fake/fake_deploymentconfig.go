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
	v1alpha1 "awsoperator.io/pkg/apis/codedeploy/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDeploymentConfigs implements DeploymentConfigInterface
type FakeDeploymentConfigs struct {
	Fake *FakeCodedeployV1alpha1
	ns   string
}

var deploymentconfigsResource = schema.GroupVersionResource{Group: "codedeploy.awsoperator.io", Version: "v1alpha1", Resource: "deploymentconfigs"}

var deploymentconfigsKind = schema.GroupVersionKind{Group: "codedeploy.awsoperator.io", Version: "v1alpha1", Kind: "DeploymentConfig"}

// Get takes name of the deploymentConfig, and returns the corresponding deploymentConfig object, and an error if there is any.
func (c *FakeDeploymentConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(deploymentconfigsResource, c.ns, name), &v1alpha1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentConfig), err
}

// List takes label and field selectors, and returns the list of DeploymentConfigs that match those selectors.
func (c *FakeDeploymentConfigs) List(opts v1.ListOptions) (result *v1alpha1.DeploymentConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(deploymentconfigsResource, deploymentconfigsKind, c.ns, opts), &v1alpha1.DeploymentConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DeploymentConfigList{ListMeta: obj.(*v1alpha1.DeploymentConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.DeploymentConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deploymentConfigs.
func (c *FakeDeploymentConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(deploymentconfigsResource, c.ns, opts))

}

// Create takes the representation of a deploymentConfig and creates it.  Returns the server's representation of the deploymentConfig, and an error, if there is any.
func (c *FakeDeploymentConfigs) Create(deploymentConfig *v1alpha1.DeploymentConfig) (result *v1alpha1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(deploymentconfigsResource, c.ns, deploymentConfig), &v1alpha1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentConfig), err
}

// Update takes the representation of a deploymentConfig and updates it. Returns the server's representation of the deploymentConfig, and an error, if there is any.
func (c *FakeDeploymentConfigs) Update(deploymentConfig *v1alpha1.DeploymentConfig) (result *v1alpha1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(deploymentconfigsResource, c.ns, deploymentConfig), &v1alpha1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDeploymentConfigs) UpdateStatus(deploymentConfig *v1alpha1.DeploymentConfig) (*v1alpha1.DeploymentConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(deploymentconfigsResource, "status", c.ns, deploymentConfig), &v1alpha1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentConfig), err
}

// Delete takes name of the deploymentConfig and deletes it. Returns an error if one occurs.
func (c *FakeDeploymentConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(deploymentconfigsResource, c.ns, name), &v1alpha1.DeploymentConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDeploymentConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(deploymentconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DeploymentConfigList{})
	return err
}

// Patch applies the patch and returns the patched deploymentConfig.
func (c *FakeDeploymentConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeploymentConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(deploymentconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DeploymentConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DeploymentConfig), err
}