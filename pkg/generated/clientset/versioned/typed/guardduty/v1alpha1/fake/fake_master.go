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
	v1alpha1 "awsoperator.io/pkg/apis/guardduty/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMasters implements MasterInterface
type FakeMasters struct {
	Fake *FakeGuarddutyV1alpha1
	ns   string
}

var mastersResource = schema.GroupVersionResource{Group: "guardduty.awsoperator.io", Version: "v1alpha1", Resource: "masters"}

var mastersKind = schema.GroupVersionKind{Group: "guardduty.awsoperator.io", Version: "v1alpha1", Kind: "Master"}

// Get takes name of the master, and returns the corresponding master object, and an error if there is any.
func (c *FakeMasters) Get(name string, options v1.GetOptions) (result *v1alpha1.Master, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mastersResource, c.ns, name), &v1alpha1.Master{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Master), err
}

// List takes label and field selectors, and returns the list of Masters that match those selectors.
func (c *FakeMasters) List(opts v1.ListOptions) (result *v1alpha1.MasterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mastersResource, mastersKind, c.ns, opts), &v1alpha1.MasterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MasterList{ListMeta: obj.(*v1alpha1.MasterList).ListMeta}
	for _, item := range obj.(*v1alpha1.MasterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested masters.
func (c *FakeMasters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mastersResource, c.ns, opts))

}

// Create takes the representation of a master and creates it.  Returns the server's representation of the master, and an error, if there is any.
func (c *FakeMasters) Create(master *v1alpha1.Master) (result *v1alpha1.Master, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mastersResource, c.ns, master), &v1alpha1.Master{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Master), err
}

// Update takes the representation of a master and updates it. Returns the server's representation of the master, and an error, if there is any.
func (c *FakeMasters) Update(master *v1alpha1.Master) (result *v1alpha1.Master, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mastersResource, c.ns, master), &v1alpha1.Master{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Master), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMasters) UpdateStatus(master *v1alpha1.Master) (*v1alpha1.Master, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mastersResource, "status", c.ns, master), &v1alpha1.Master{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Master), err
}

// Delete takes name of the master and deletes it. Returns an error if one occurs.
func (c *FakeMasters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mastersResource, c.ns, name), &v1alpha1.Master{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMasters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mastersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MasterList{})
	return err
}

// Patch applies the patch and returns the patched master.
func (c *FakeMasters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Master, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mastersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Master{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Master), err
}