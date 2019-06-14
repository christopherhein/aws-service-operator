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
	v1alpha1 "awsoperator.io/pkg/apis/iam/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccessKeys implements AccessKeyInterface
type FakeAccessKeys struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var accesskeysResource = schema.GroupVersionResource{Group: "iam.awsoperator.io", Version: "v1alpha1", Resource: "accesskeys"}

var accesskeysKind = schema.GroupVersionKind{Group: "iam.awsoperator.io", Version: "v1alpha1", Kind: "AccessKey"}

// Get takes name of the accessKey, and returns the corresponding accessKey object, and an error if there is any.
func (c *FakeAccessKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.AccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accesskeysResource, c.ns, name), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}

// List takes label and field selectors, and returns the list of AccessKeys that match those selectors.
func (c *FakeAccessKeys) List(opts v1.ListOptions) (result *v1alpha1.AccessKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accesskeysResource, accesskeysKind, c.ns, opts), &v1alpha1.AccessKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccessKeyList{ListMeta: obj.(*v1alpha1.AccessKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccessKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accessKeys.
func (c *FakeAccessKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accesskeysResource, c.ns, opts))

}

// Create takes the representation of a accessKey and creates it.  Returns the server's representation of the accessKey, and an error, if there is any.
func (c *FakeAccessKeys) Create(accessKey *v1alpha1.AccessKey) (result *v1alpha1.AccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accesskeysResource, c.ns, accessKey), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}

// Update takes the representation of a accessKey and updates it. Returns the server's representation of the accessKey, and an error, if there is any.
func (c *FakeAccessKeys) Update(accessKey *v1alpha1.AccessKey) (result *v1alpha1.AccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accesskeysResource, c.ns, accessKey), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccessKeys) UpdateStatus(accessKey *v1alpha1.AccessKey) (*v1alpha1.AccessKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accesskeysResource, "status", c.ns, accessKey), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}

// Delete takes name of the accessKey and deletes it. Returns an error if one occurs.
func (c *FakeAccessKeys) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accesskeysResource, c.ns, name), &v1alpha1.AccessKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccessKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accesskeysResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccessKeyList{})
	return err
}

// Patch applies the patch and returns the patched accessKey.
func (c *FakeAccessKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accesskeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}