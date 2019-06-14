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

// FakeAutoScalingGroups implements AutoScalingGroupInterface
type FakeAutoScalingGroups struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var autoscalinggroupsResource = schema.GroupVersionResource{Group: "autoscaling.awsoperator.io", Version: "v1alpha1", Resource: "autoscalinggroups"}

var autoscalinggroupsKind = schema.GroupVersionKind{Group: "autoscaling.awsoperator.io", Version: "v1alpha1", Kind: "AutoScalingGroup"}

// Get takes name of the autoScalingGroup, and returns the corresponding autoScalingGroup object, and an error if there is any.
func (c *FakeAutoScalingGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.AutoScalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(autoscalinggroupsResource, c.ns, name), &v1alpha1.AutoScalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScalingGroup), err
}

// List takes label and field selectors, and returns the list of AutoScalingGroups that match those selectors.
func (c *FakeAutoScalingGroups) List(opts v1.ListOptions) (result *v1alpha1.AutoScalingGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(autoscalinggroupsResource, autoscalinggroupsKind, c.ns, opts), &v1alpha1.AutoScalingGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AutoScalingGroupList{ListMeta: obj.(*v1alpha1.AutoScalingGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.AutoScalingGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested autoScalingGroups.
func (c *FakeAutoScalingGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(autoscalinggroupsResource, c.ns, opts))

}

// Create takes the representation of a autoScalingGroup and creates it.  Returns the server's representation of the autoScalingGroup, and an error, if there is any.
func (c *FakeAutoScalingGroups) Create(autoScalingGroup *v1alpha1.AutoScalingGroup) (result *v1alpha1.AutoScalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(autoscalinggroupsResource, c.ns, autoScalingGroup), &v1alpha1.AutoScalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScalingGroup), err
}

// Update takes the representation of a autoScalingGroup and updates it. Returns the server's representation of the autoScalingGroup, and an error, if there is any.
func (c *FakeAutoScalingGroups) Update(autoScalingGroup *v1alpha1.AutoScalingGroup) (result *v1alpha1.AutoScalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(autoscalinggroupsResource, c.ns, autoScalingGroup), &v1alpha1.AutoScalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScalingGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAutoScalingGroups) UpdateStatus(autoScalingGroup *v1alpha1.AutoScalingGroup) (*v1alpha1.AutoScalingGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(autoscalinggroupsResource, "status", c.ns, autoScalingGroup), &v1alpha1.AutoScalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScalingGroup), err
}

// Delete takes name of the autoScalingGroup and deletes it. Returns an error if one occurs.
func (c *FakeAutoScalingGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(autoscalinggroupsResource, c.ns, name), &v1alpha1.AutoScalingGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAutoScalingGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(autoscalinggroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AutoScalingGroupList{})
	return err
}

// Patch applies the patch and returns the patched autoScalingGroup.
func (c *FakeAutoScalingGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AutoScalingGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(autoscalinggroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AutoScalingGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AutoScalingGroup), err
}