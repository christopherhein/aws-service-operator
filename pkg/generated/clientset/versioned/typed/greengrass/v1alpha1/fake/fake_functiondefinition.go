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
	v1alpha1 "awsoperator.io/pkg/apis/greengrass/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFunctionDefinitions implements FunctionDefinitionInterface
type FakeFunctionDefinitions struct {
	Fake *FakeGreengrassV1alpha1
	ns   string
}

var functiondefinitionsResource = schema.GroupVersionResource{Group: "greengrass.awsoperator.io", Version: "v1alpha1", Resource: "functiondefinitions"}

var functiondefinitionsKind = schema.GroupVersionKind{Group: "greengrass.awsoperator.io", Version: "v1alpha1", Kind: "FunctionDefinition"}

// Get takes name of the functionDefinition, and returns the corresponding functionDefinition object, and an error if there is any.
func (c *FakeFunctionDefinitions) Get(name string, options v1.GetOptions) (result *v1alpha1.FunctionDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(functiondefinitionsResource, c.ns, name), &v1alpha1.FunctionDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionDefinition), err
}

// List takes label and field selectors, and returns the list of FunctionDefinitions that match those selectors.
func (c *FakeFunctionDefinitions) List(opts v1.ListOptions) (result *v1alpha1.FunctionDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(functiondefinitionsResource, functiondefinitionsKind, c.ns, opts), &v1alpha1.FunctionDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FunctionDefinitionList{ListMeta: obj.(*v1alpha1.FunctionDefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.FunctionDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested functionDefinitions.
func (c *FakeFunctionDefinitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(functiondefinitionsResource, c.ns, opts))

}

// Create takes the representation of a functionDefinition and creates it.  Returns the server's representation of the functionDefinition, and an error, if there is any.
func (c *FakeFunctionDefinitions) Create(functionDefinition *v1alpha1.FunctionDefinition) (result *v1alpha1.FunctionDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(functiondefinitionsResource, c.ns, functionDefinition), &v1alpha1.FunctionDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionDefinition), err
}

// Update takes the representation of a functionDefinition and updates it. Returns the server's representation of the functionDefinition, and an error, if there is any.
func (c *FakeFunctionDefinitions) Update(functionDefinition *v1alpha1.FunctionDefinition) (result *v1alpha1.FunctionDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(functiondefinitionsResource, c.ns, functionDefinition), &v1alpha1.FunctionDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFunctionDefinitions) UpdateStatus(functionDefinition *v1alpha1.FunctionDefinition) (*v1alpha1.FunctionDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(functiondefinitionsResource, "status", c.ns, functionDefinition), &v1alpha1.FunctionDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionDefinition), err
}

// Delete takes name of the functionDefinition and deletes it. Returns an error if one occurs.
func (c *FakeFunctionDefinitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(functiondefinitionsResource, c.ns, name), &v1alpha1.FunctionDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFunctionDefinitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(functiondefinitionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FunctionDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched functionDefinition.
func (c *FakeFunctionDefinitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FunctionDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(functiondefinitionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FunctionDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionDefinition), err
}