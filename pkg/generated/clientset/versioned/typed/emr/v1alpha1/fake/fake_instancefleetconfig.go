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
	v1alpha1 "awsoperator.io/pkg/apis/emr/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInstanceFleetConfigs implements InstanceFleetConfigInterface
type FakeInstanceFleetConfigs struct {
	Fake *FakeEmrV1alpha1
	ns   string
}

var instancefleetconfigsResource = schema.GroupVersionResource{Group: "emr.awsoperator.io", Version: "v1alpha1", Resource: "instancefleetconfigs"}

var instancefleetconfigsKind = schema.GroupVersionKind{Group: "emr.awsoperator.io", Version: "v1alpha1", Kind: "InstanceFleetConfig"}

// Get takes name of the instanceFleetConfig, and returns the corresponding instanceFleetConfig object, and an error if there is any.
func (c *FakeInstanceFleetConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.InstanceFleetConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(instancefleetconfigsResource, c.ns, name), &v1alpha1.InstanceFleetConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceFleetConfig), err
}

// List takes label and field selectors, and returns the list of InstanceFleetConfigs that match those selectors.
func (c *FakeInstanceFleetConfigs) List(opts v1.ListOptions) (result *v1alpha1.InstanceFleetConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(instancefleetconfigsResource, instancefleetconfigsKind, c.ns, opts), &v1alpha1.InstanceFleetConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstanceFleetConfigList{ListMeta: obj.(*v1alpha1.InstanceFleetConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstanceFleetConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instanceFleetConfigs.
func (c *FakeInstanceFleetConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(instancefleetconfigsResource, c.ns, opts))

}

// Create takes the representation of a instanceFleetConfig and creates it.  Returns the server's representation of the instanceFleetConfig, and an error, if there is any.
func (c *FakeInstanceFleetConfigs) Create(instanceFleetConfig *v1alpha1.InstanceFleetConfig) (result *v1alpha1.InstanceFleetConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(instancefleetconfigsResource, c.ns, instanceFleetConfig), &v1alpha1.InstanceFleetConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceFleetConfig), err
}

// Update takes the representation of a instanceFleetConfig and updates it. Returns the server's representation of the instanceFleetConfig, and an error, if there is any.
func (c *FakeInstanceFleetConfigs) Update(instanceFleetConfig *v1alpha1.InstanceFleetConfig) (result *v1alpha1.InstanceFleetConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(instancefleetconfigsResource, c.ns, instanceFleetConfig), &v1alpha1.InstanceFleetConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceFleetConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstanceFleetConfigs) UpdateStatus(instanceFleetConfig *v1alpha1.InstanceFleetConfig) (*v1alpha1.InstanceFleetConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(instancefleetconfigsResource, "status", c.ns, instanceFleetConfig), &v1alpha1.InstanceFleetConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceFleetConfig), err
}

// Delete takes name of the instanceFleetConfig and deletes it. Returns an error if one occurs.
func (c *FakeInstanceFleetConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(instancefleetconfigsResource, c.ns, name), &v1alpha1.InstanceFleetConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstanceFleetConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(instancefleetconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstanceFleetConfigList{})
	return err
}

// Patch applies the patch and returns the patched instanceFleetConfig.
func (c *FakeInstanceFleetConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstanceFleetConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(instancefleetconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.InstanceFleetConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceFleetConfig), err
}