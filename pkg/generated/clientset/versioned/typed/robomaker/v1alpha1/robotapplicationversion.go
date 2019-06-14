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

	v1alpha1 "awsoperator.io/pkg/apis/robomaker/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RobotApplicationVersionsGetter has a method to return a RobotApplicationVersionInterface.
// A group's client should implement this interface.
type RobotApplicationVersionsGetter interface {
	RobotApplicationVersions(namespace string) RobotApplicationVersionInterface
}

// RobotApplicationVersionInterface has methods to work with RobotApplicationVersion resources.
type RobotApplicationVersionInterface interface {
	Create(*v1alpha1.RobotApplicationVersion) (*v1alpha1.RobotApplicationVersion, error)
	Update(*v1alpha1.RobotApplicationVersion) (*v1alpha1.RobotApplicationVersion, error)
	UpdateStatus(*v1alpha1.RobotApplicationVersion) (*v1alpha1.RobotApplicationVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RobotApplicationVersion, error)
	List(opts v1.ListOptions) (*v1alpha1.RobotApplicationVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RobotApplicationVersion, err error)
	RobotApplicationVersionExpansion
}

// robotApplicationVersions implements RobotApplicationVersionInterface
type robotApplicationVersions struct {
	client rest.Interface
	ns     string
}

// newRobotApplicationVersions returns a RobotApplicationVersions
func newRobotApplicationVersions(c *RobomakerV1alpha1Client, namespace string) *robotApplicationVersions {
	return &robotApplicationVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the robotApplicationVersion, and returns the corresponding robotApplicationVersion object, and an error if there is any.
func (c *robotApplicationVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.RobotApplicationVersion, err error) {
	result = &v1alpha1.RobotApplicationVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("robotapplicationversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RobotApplicationVersions that match those selectors.
func (c *robotApplicationVersions) List(opts v1.ListOptions) (result *v1alpha1.RobotApplicationVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RobotApplicationVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("robotapplicationversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested robotApplicationVersions.
func (c *robotApplicationVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("robotapplicationversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a robotApplicationVersion and creates it.  Returns the server's representation of the robotApplicationVersion, and an error, if there is any.
func (c *robotApplicationVersions) Create(robotApplicationVersion *v1alpha1.RobotApplicationVersion) (result *v1alpha1.RobotApplicationVersion, err error) {
	result = &v1alpha1.RobotApplicationVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("robotapplicationversions").
		Body(robotApplicationVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a robotApplicationVersion and updates it. Returns the server's representation of the robotApplicationVersion, and an error, if there is any.
func (c *robotApplicationVersions) Update(robotApplicationVersion *v1alpha1.RobotApplicationVersion) (result *v1alpha1.RobotApplicationVersion, err error) {
	result = &v1alpha1.RobotApplicationVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("robotapplicationversions").
		Name(robotApplicationVersion.Name).
		Body(robotApplicationVersion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *robotApplicationVersions) UpdateStatus(robotApplicationVersion *v1alpha1.RobotApplicationVersion) (result *v1alpha1.RobotApplicationVersion, err error) {
	result = &v1alpha1.RobotApplicationVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("robotapplicationversions").
		Name(robotApplicationVersion.Name).
		SubResource("status").
		Body(robotApplicationVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the robotApplicationVersion and deletes it. Returns an error if one occurs.
func (c *robotApplicationVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("robotapplicationversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *robotApplicationVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("robotapplicationversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched robotApplicationVersion.
func (c *robotApplicationVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RobotApplicationVersion, err error) {
	result = &v1alpha1.RobotApplicationVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("robotapplicationversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}