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
	v1alpha1 "awsoperator.io/pkg/apis/autoscaling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScheduledActions implements ScheduledActionInterface
type FakeScheduledActions struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var scheduledactionsResource = schema.GroupVersionResource{Group: "autoscaling.awsoperator.io", Version: "v1alpha1", Resource: "scheduledactions"}

var scheduledactionsKind = schema.GroupVersionKind{Group: "autoscaling.awsoperator.io", Version: "v1alpha1", Kind: "ScheduledAction"}

// Get takes name of the scheduledAction, and returns the corresponding scheduledAction object, and an error if there is any.
func (c *FakeScheduledActions) Get(name string, options v1.GetOptions) (result *v1alpha1.ScheduledAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scheduledactionsResource, c.ns, name), &v1alpha1.ScheduledAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledAction), err
}

// List takes label and field selectors, and returns the list of ScheduledActions that match those selectors.
func (c *FakeScheduledActions) List(opts v1.ListOptions) (result *v1alpha1.ScheduledActionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scheduledactionsResource, scheduledactionsKind, c.ns, opts), &v1alpha1.ScheduledActionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScheduledActionList{ListMeta: obj.(*v1alpha1.ScheduledActionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScheduledActionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scheduledActions.
func (c *FakeScheduledActions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scheduledactionsResource, c.ns, opts))

}

// Create takes the representation of a scheduledAction and creates it.  Returns the server's representation of the scheduledAction, and an error, if there is any.
func (c *FakeScheduledActions) Create(scheduledAction *v1alpha1.ScheduledAction) (result *v1alpha1.ScheduledAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scheduledactionsResource, c.ns, scheduledAction), &v1alpha1.ScheduledAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledAction), err
}

// Update takes the representation of a scheduledAction and updates it. Returns the server's representation of the scheduledAction, and an error, if there is any.
func (c *FakeScheduledActions) Update(scheduledAction *v1alpha1.ScheduledAction) (result *v1alpha1.ScheduledAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scheduledactionsResource, c.ns, scheduledAction), &v1alpha1.ScheduledAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledAction), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScheduledActions) UpdateStatus(scheduledAction *v1alpha1.ScheduledAction) (*v1alpha1.ScheduledAction, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(scheduledactionsResource, "status", c.ns, scheduledAction), &v1alpha1.ScheduledAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledAction), err
}

// Delete takes name of the scheduledAction and deletes it. Returns an error if one occurs.
func (c *FakeScheduledActions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scheduledactionsResource, c.ns, name), &v1alpha1.ScheduledAction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScheduledActions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scheduledactionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScheduledActionList{})
	return err
}

// Patch applies the patch and returns the patched scheduledAction.
func (c *FakeScheduledActions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScheduledAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scheduledactionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ScheduledAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScheduledAction), err
}