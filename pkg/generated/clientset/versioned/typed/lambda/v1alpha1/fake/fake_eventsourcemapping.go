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
	v1alpha1 "awsoperator.io/pkg/apis/lambda/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEventSourceMappings implements EventSourceMappingInterface
type FakeEventSourceMappings struct {
	Fake *FakeLambdaV1alpha1
	ns   string
}

var eventsourcemappingsResource = schema.GroupVersionResource{Group: "lambda.awsoperator.io", Version: "v1alpha1", Resource: "eventsourcemappings"}

var eventsourcemappingsKind = schema.GroupVersionKind{Group: "lambda.awsoperator.io", Version: "v1alpha1", Kind: "EventSourceMapping"}

// Get takes name of the eventSourceMapping, and returns the corresponding eventSourceMapping object, and an error if there is any.
func (c *FakeEventSourceMappings) Get(name string, options v1.GetOptions) (result *v1alpha1.EventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventsourcemappingsResource, c.ns, name), &v1alpha1.EventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventSourceMapping), err
}

// List takes label and field selectors, and returns the list of EventSourceMappings that match those selectors.
func (c *FakeEventSourceMappings) List(opts v1.ListOptions) (result *v1alpha1.EventSourceMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventsourcemappingsResource, eventsourcemappingsKind, c.ns, opts), &v1alpha1.EventSourceMappingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventSourceMappingList{ListMeta: obj.(*v1alpha1.EventSourceMappingList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventSourceMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventSourceMappings.
func (c *FakeEventSourceMappings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventsourcemappingsResource, c.ns, opts))

}

// Create takes the representation of a eventSourceMapping and creates it.  Returns the server's representation of the eventSourceMapping, and an error, if there is any.
func (c *FakeEventSourceMappings) Create(eventSourceMapping *v1alpha1.EventSourceMapping) (result *v1alpha1.EventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventsourcemappingsResource, c.ns, eventSourceMapping), &v1alpha1.EventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventSourceMapping), err
}

// Update takes the representation of a eventSourceMapping and updates it. Returns the server's representation of the eventSourceMapping, and an error, if there is any.
func (c *FakeEventSourceMappings) Update(eventSourceMapping *v1alpha1.EventSourceMapping) (result *v1alpha1.EventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventsourcemappingsResource, c.ns, eventSourceMapping), &v1alpha1.EventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventSourceMapping), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventSourceMappings) UpdateStatus(eventSourceMapping *v1alpha1.EventSourceMapping) (*v1alpha1.EventSourceMapping, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventsourcemappingsResource, "status", c.ns, eventSourceMapping), &v1alpha1.EventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventSourceMapping), err
}

// Delete takes name of the eventSourceMapping and deletes it. Returns an error if one occurs.
func (c *FakeEventSourceMappings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventsourcemappingsResource, c.ns, name), &v1alpha1.EventSourceMapping{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventSourceMappings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventsourcemappingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventSourceMappingList{})
	return err
}

// Patch applies the patch and returns the patched eventSourceMapping.
func (c *FakeEventSourceMappings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.EventSourceMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventsourcemappingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventSourceMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventSourceMapping), err
}