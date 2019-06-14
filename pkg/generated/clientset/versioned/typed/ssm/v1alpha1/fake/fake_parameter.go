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
	v1alpha1 "awsoperator.io/pkg/apis/ssm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeParameters implements ParameterInterface
type FakeParameters struct {
	Fake *FakeSsmV1alpha1
	ns   string
}

var parametersResource = schema.GroupVersionResource{Group: "ssm.awsoperator.io", Version: "v1alpha1", Resource: "parameters"}

var parametersKind = schema.GroupVersionKind{Group: "ssm.awsoperator.io", Version: "v1alpha1", Kind: "Parameter"}

// Get takes name of the parameter, and returns the corresponding parameter object, and an error if there is any.
func (c *FakeParameters) Get(name string, options v1.GetOptions) (result *v1alpha1.Parameter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(parametersResource, c.ns, name), &v1alpha1.Parameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Parameter), err
}

// List takes label and field selectors, and returns the list of Parameters that match those selectors.
func (c *FakeParameters) List(opts v1.ListOptions) (result *v1alpha1.ParameterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(parametersResource, parametersKind, c.ns, opts), &v1alpha1.ParameterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ParameterList{ListMeta: obj.(*v1alpha1.ParameterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ParameterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested parameters.
func (c *FakeParameters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(parametersResource, c.ns, opts))

}

// Create takes the representation of a parameter and creates it.  Returns the server's representation of the parameter, and an error, if there is any.
func (c *FakeParameters) Create(parameter *v1alpha1.Parameter) (result *v1alpha1.Parameter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(parametersResource, c.ns, parameter), &v1alpha1.Parameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Parameter), err
}

// Update takes the representation of a parameter and updates it. Returns the server's representation of the parameter, and an error, if there is any.
func (c *FakeParameters) Update(parameter *v1alpha1.Parameter) (result *v1alpha1.Parameter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(parametersResource, c.ns, parameter), &v1alpha1.Parameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Parameter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeParameters) UpdateStatus(parameter *v1alpha1.Parameter) (*v1alpha1.Parameter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(parametersResource, "status", c.ns, parameter), &v1alpha1.Parameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Parameter), err
}

// Delete takes name of the parameter and deletes it. Returns an error if one occurs.
func (c *FakeParameters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(parametersResource, c.ns, name), &v1alpha1.Parameter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeParameters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(parametersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ParameterList{})
	return err
}

// Patch applies the patch and returns the patched parameter.
func (c *FakeParameters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Parameter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(parametersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Parameter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Parameter), err
}