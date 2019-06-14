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
	v1alpha1 "awsoperator.io/pkg/apis/servicediscovery/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrivateDnsNamespaces implements PrivateDnsNamespaceInterface
type FakePrivateDnsNamespaces struct {
	Fake *FakeServicediscoveryV1alpha1
	ns   string
}

var privatednsnamespacesResource = schema.GroupVersionResource{Group: "servicediscovery.awsoperator.io", Version: "v1alpha1", Resource: "privatednsnamespaces"}

var privatednsnamespacesKind = schema.GroupVersionKind{Group: "servicediscovery.awsoperator.io", Version: "v1alpha1", Kind: "PrivateDnsNamespace"}

// Get takes name of the privateDnsNamespace, and returns the corresponding privateDnsNamespace object, and an error if there is any.
func (c *FakePrivateDnsNamespaces) Get(name string, options v1.GetOptions) (result *v1alpha1.PrivateDnsNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(privatednsnamespacesResource, c.ns, name), &v1alpha1.PrivateDnsNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDnsNamespace), err
}

// List takes label and field selectors, and returns the list of PrivateDnsNamespaces that match those selectors.
func (c *FakePrivateDnsNamespaces) List(opts v1.ListOptions) (result *v1alpha1.PrivateDnsNamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(privatednsnamespacesResource, privatednsnamespacesKind, c.ns, opts), &v1alpha1.PrivateDnsNamespaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrivateDnsNamespaceList{ListMeta: obj.(*v1alpha1.PrivateDnsNamespaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrivateDnsNamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested privateDnsNamespaces.
func (c *FakePrivateDnsNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(privatednsnamespacesResource, c.ns, opts))

}

// Create takes the representation of a privateDnsNamespace and creates it.  Returns the server's representation of the privateDnsNamespace, and an error, if there is any.
func (c *FakePrivateDnsNamespaces) Create(privateDnsNamespace *v1alpha1.PrivateDnsNamespace) (result *v1alpha1.PrivateDnsNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(privatednsnamespacesResource, c.ns, privateDnsNamespace), &v1alpha1.PrivateDnsNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDnsNamespace), err
}

// Update takes the representation of a privateDnsNamespace and updates it. Returns the server's representation of the privateDnsNamespace, and an error, if there is any.
func (c *FakePrivateDnsNamespaces) Update(privateDnsNamespace *v1alpha1.PrivateDnsNamespace) (result *v1alpha1.PrivateDnsNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(privatednsnamespacesResource, c.ns, privateDnsNamespace), &v1alpha1.PrivateDnsNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDnsNamespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrivateDnsNamespaces) UpdateStatus(privateDnsNamespace *v1alpha1.PrivateDnsNamespace) (*v1alpha1.PrivateDnsNamespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(privatednsnamespacesResource, "status", c.ns, privateDnsNamespace), &v1alpha1.PrivateDnsNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDnsNamespace), err
}

// Delete takes name of the privateDnsNamespace and deletes it. Returns an error if one occurs.
func (c *FakePrivateDnsNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(privatednsnamespacesResource, c.ns, name), &v1alpha1.PrivateDnsNamespace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrivateDnsNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(privatednsnamespacesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrivateDnsNamespaceList{})
	return err
}

// Patch applies the patch and returns the patched privateDnsNamespace.
func (c *FakePrivateDnsNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrivateDnsNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(privatednsnamespacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrivateDnsNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrivateDnsNamespace), err
}