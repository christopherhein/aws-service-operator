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
	v1alpha1 "awsoperator.io/pkg/apis/ec2/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHosts implements HostInterface
type FakeHosts struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var hostsResource = schema.GroupVersionResource{Group: "ec2.awsoperator.io", Version: "v1alpha1", Resource: "hosts"}

var hostsKind = schema.GroupVersionKind{Group: "ec2.awsoperator.io", Version: "v1alpha1", Kind: "Host"}

// Get takes name of the host, and returns the corresponding host object, and an error if there is any.
func (c *FakeHosts) Get(name string, options v1.GetOptions) (result *v1alpha1.Host, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hostsResource, c.ns, name), &v1alpha1.Host{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Host), err
}

// List takes label and field selectors, and returns the list of Hosts that match those selectors.
func (c *FakeHosts) List(opts v1.ListOptions) (result *v1alpha1.HostList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hostsResource, hostsKind, c.ns, opts), &v1alpha1.HostList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HostList{ListMeta: obj.(*v1alpha1.HostList).ListMeta}
	for _, item := range obj.(*v1alpha1.HostList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hosts.
func (c *FakeHosts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hostsResource, c.ns, opts))

}

// Create takes the representation of a host and creates it.  Returns the server's representation of the host, and an error, if there is any.
func (c *FakeHosts) Create(host *v1alpha1.Host) (result *v1alpha1.Host, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hostsResource, c.ns, host), &v1alpha1.Host{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Host), err
}

// Update takes the representation of a host and updates it. Returns the server's representation of the host, and an error, if there is any.
func (c *FakeHosts) Update(host *v1alpha1.Host) (result *v1alpha1.Host, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hostsResource, c.ns, host), &v1alpha1.Host{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Host), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHosts) UpdateStatus(host *v1alpha1.Host) (*v1alpha1.Host, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hostsResource, "status", c.ns, host), &v1alpha1.Host{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Host), err
}

// Delete takes name of the host and deletes it. Returns an error if one occurs.
func (c *FakeHosts) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hostsResource, c.ns, name), &v1alpha1.Host{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHosts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hostsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HostList{})
	return err
}

// Patch applies the patch and returns the patched host.
func (c *FakeHosts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Host, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Host{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Host), err
}