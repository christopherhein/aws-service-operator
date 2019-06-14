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
	v1alpha1 "awsoperator.io/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigurationAggregators implements ConfigurationAggregatorInterface
type FakeConfigurationAggregators struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var configurationaggregatorsResource = schema.GroupVersionResource{Group: "config.awsoperator.io", Version: "v1alpha1", Resource: "configurationaggregators"}

var configurationaggregatorsKind = schema.GroupVersionKind{Group: "config.awsoperator.io", Version: "v1alpha1", Kind: "ConfigurationAggregator"}

// Get takes name of the configurationAggregator, and returns the corresponding configurationAggregator object, and an error if there is any.
func (c *FakeConfigurationAggregators) Get(name string, options v1.GetOptions) (result *v1alpha1.ConfigurationAggregator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configurationaggregatorsResource, c.ns, name), &v1alpha1.ConfigurationAggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationAggregator), err
}

// List takes label and field selectors, and returns the list of ConfigurationAggregators that match those selectors.
func (c *FakeConfigurationAggregators) List(opts v1.ListOptions) (result *v1alpha1.ConfigurationAggregatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configurationaggregatorsResource, configurationaggregatorsKind, c.ns, opts), &v1alpha1.ConfigurationAggregatorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConfigurationAggregatorList{ListMeta: obj.(*v1alpha1.ConfigurationAggregatorList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConfigurationAggregatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configurationAggregators.
func (c *FakeConfigurationAggregators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configurationaggregatorsResource, c.ns, opts))

}

// Create takes the representation of a configurationAggregator and creates it.  Returns the server's representation of the configurationAggregator, and an error, if there is any.
func (c *FakeConfigurationAggregators) Create(configurationAggregator *v1alpha1.ConfigurationAggregator) (result *v1alpha1.ConfigurationAggregator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configurationaggregatorsResource, c.ns, configurationAggregator), &v1alpha1.ConfigurationAggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationAggregator), err
}

// Update takes the representation of a configurationAggregator and updates it. Returns the server's representation of the configurationAggregator, and an error, if there is any.
func (c *FakeConfigurationAggregators) Update(configurationAggregator *v1alpha1.ConfigurationAggregator) (result *v1alpha1.ConfigurationAggregator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configurationaggregatorsResource, c.ns, configurationAggregator), &v1alpha1.ConfigurationAggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationAggregator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigurationAggregators) UpdateStatus(configurationAggregator *v1alpha1.ConfigurationAggregator) (*v1alpha1.ConfigurationAggregator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(configurationaggregatorsResource, "status", c.ns, configurationAggregator), &v1alpha1.ConfigurationAggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationAggregator), err
}

// Delete takes name of the configurationAggregator and deletes it. Returns an error if one occurs.
func (c *FakeConfigurationAggregators) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configurationaggregatorsResource, c.ns, name), &v1alpha1.ConfigurationAggregator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigurationAggregators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configurationaggregatorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConfigurationAggregatorList{})
	return err
}

// Patch applies the patch and returns the patched configurationAggregator.
func (c *FakeConfigurationAggregators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ConfigurationAggregator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configurationaggregatorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConfigurationAggregator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationAggregator), err
}