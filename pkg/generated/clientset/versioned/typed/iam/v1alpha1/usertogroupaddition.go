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

package v1alpha1

import (
	"time"

	v1alpha1 "awsoperator.io/pkg/apis/iam/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UserToGroupAdditionsGetter has a method to return a UserToGroupAdditionInterface.
// A group's client should implement this interface.
type UserToGroupAdditionsGetter interface {
	UserToGroupAdditions(namespace string) UserToGroupAdditionInterface
}

// UserToGroupAdditionInterface has methods to work with UserToGroupAddition resources.
type UserToGroupAdditionInterface interface {
	Create(*v1alpha1.UserToGroupAddition) (*v1alpha1.UserToGroupAddition, error)
	Update(*v1alpha1.UserToGroupAddition) (*v1alpha1.UserToGroupAddition, error)
	UpdateStatus(*v1alpha1.UserToGroupAddition) (*v1alpha1.UserToGroupAddition, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.UserToGroupAddition, error)
	List(opts v1.ListOptions) (*v1alpha1.UserToGroupAdditionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UserToGroupAddition, err error)
	UserToGroupAdditionExpansion
}

// userToGroupAdditions implements UserToGroupAdditionInterface
type userToGroupAdditions struct {
	client rest.Interface
	ns     string
}

// newUserToGroupAdditions returns a UserToGroupAdditions
func newUserToGroupAdditions(c *IamV1alpha1Client, namespace string) *userToGroupAdditions {
	return &userToGroupAdditions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the userToGroupAddition, and returns the corresponding userToGroupAddition object, and an error if there is any.
func (c *userToGroupAdditions) Get(name string, options v1.GetOptions) (result *v1alpha1.UserToGroupAddition, err error) {
	result = &v1alpha1.UserToGroupAddition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("usertogroupadditions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UserToGroupAdditions that match those selectors.
func (c *userToGroupAdditions) List(opts v1.ListOptions) (result *v1alpha1.UserToGroupAdditionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UserToGroupAdditionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("usertogroupadditions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested userToGroupAdditions.
func (c *userToGroupAdditions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("usertogroupadditions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a userToGroupAddition and creates it.  Returns the server's representation of the userToGroupAddition, and an error, if there is any.
func (c *userToGroupAdditions) Create(userToGroupAddition *v1alpha1.UserToGroupAddition) (result *v1alpha1.UserToGroupAddition, err error) {
	result = &v1alpha1.UserToGroupAddition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("usertogroupadditions").
		Body(userToGroupAddition).
		Do().
		Into(result)
	return
}

// Update takes the representation of a userToGroupAddition and updates it. Returns the server's representation of the userToGroupAddition, and an error, if there is any.
func (c *userToGroupAdditions) Update(userToGroupAddition *v1alpha1.UserToGroupAddition) (result *v1alpha1.UserToGroupAddition, err error) {
	result = &v1alpha1.UserToGroupAddition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("usertogroupadditions").
		Name(userToGroupAddition.Name).
		Body(userToGroupAddition).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *userToGroupAdditions) UpdateStatus(userToGroupAddition *v1alpha1.UserToGroupAddition) (result *v1alpha1.UserToGroupAddition, err error) {
	result = &v1alpha1.UserToGroupAddition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("usertogroupadditions").
		Name(userToGroupAddition.Name).
		SubResource("status").
		Body(userToGroupAddition).
		Do().
		Into(result)
	return
}

// Delete takes name of the userToGroupAddition and deletes it. Returns an error if one occurs.
func (c *userToGroupAdditions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("usertogroupadditions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *userToGroupAdditions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("usertogroupadditions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched userToGroupAddition.
func (c *userToGroupAdditions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UserToGroupAddition, err error) {
	result = &v1alpha1.UserToGroupAddition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("usertogroupadditions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}