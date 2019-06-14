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

// FakeRoles implements RoleInterface
type FakeRoles struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var rolesResource = schema.GroupVersionResource{Group: "iam.awsoperator.io", Version: "v1alpha1", Resource: "roles"}

var rolesKind = schema.GroupVersionKind{Group: "iam.awsoperator.io", Version: "v1alpha1", Kind: "Role"}

// Get takes name of the role, and returns the corresponding role object, and an error if there is any.
func (c *FakeRoles) Get(name string, options v1.GetOptions) (result *v1alpha1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rolesResource, c.ns, name), &v1alpha1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Role), err
}

// List takes label and field selectors, and returns the list of Roles that match those selectors.
func (c *FakeRoles) List(opts v1.ListOptions) (result *v1alpha1.RoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rolesResource, rolesKind, c.ns, opts), &v1alpha1.RoleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RoleList{ListMeta: obj.(*v1alpha1.RoleList).ListMeta}
	for _, item := range obj.(*v1alpha1.RoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roles.
func (c *FakeRoles) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rolesResource, c.ns, opts))

}

// Create takes the representation of a role and creates it.  Returns the server's representation of the role, and an error, if there is any.
func (c *FakeRoles) Create(role *v1alpha1.Role) (result *v1alpha1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rolesResource, c.ns, role), &v1alpha1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Role), err
}

// Update takes the representation of a role and updates it. Returns the server's representation of the role, and an error, if there is any.
func (c *FakeRoles) Update(role *v1alpha1.Role) (result *v1alpha1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rolesResource, c.ns, role), &v1alpha1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Role), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoles) UpdateStatus(role *v1alpha1.Role) (*v1alpha1.Role, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rolesResource, "status", c.ns, role), &v1alpha1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Role), err
}

// Delete takes name of the role and deletes it. Returns an error if one occurs.
func (c *FakeRoles) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rolesResource, c.ns, name), &v1alpha1.Role{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoles) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rolesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RoleList{})
	return err
}

// Patch applies the patch and returns the patched role.
func (c *FakeRoles) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Role, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rolesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Role{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Role), err
}