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

	v1alpha1 "awsoperator.io/pkg/apis/elasticloadbalancingv2/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TargetGroupsGetter has a method to return a TargetGroupInterface.
// A group's client should implement this interface.
type TargetGroupsGetter interface {
	TargetGroups(namespace string) TargetGroupInterface
}

// TargetGroupInterface has methods to work with TargetGroup resources.
type TargetGroupInterface interface {
	Create(*v1alpha1.TargetGroup) (*v1alpha1.TargetGroup, error)
	Update(*v1alpha1.TargetGroup) (*v1alpha1.TargetGroup, error)
	UpdateStatus(*v1alpha1.TargetGroup) (*v1alpha1.TargetGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TargetGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.TargetGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TargetGroup, err error)
	TargetGroupExpansion
}

// targetGroups implements TargetGroupInterface
type targetGroups struct {
	client rest.Interface
	ns     string
}

// newTargetGroups returns a TargetGroups
func newTargetGroups(c *Elasticloadbalancingv2V1alpha1Client, namespace string) *targetGroups {
	return &targetGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the targetGroup, and returns the corresponding targetGroup object, and an error if there is any.
func (c *targetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.TargetGroup, err error) {
	result = &v1alpha1.TargetGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("targetgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TargetGroups that match those selectors.
func (c *targetGroups) List(opts v1.ListOptions) (result *v1alpha1.TargetGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TargetGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("targetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested targetGroups.
func (c *targetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("targetgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a targetGroup and creates it.  Returns the server's representation of the targetGroup, and an error, if there is any.
func (c *targetGroups) Create(targetGroup *v1alpha1.TargetGroup) (result *v1alpha1.TargetGroup, err error) {
	result = &v1alpha1.TargetGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("targetgroups").
		Body(targetGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a targetGroup and updates it. Returns the server's representation of the targetGroup, and an error, if there is any.
func (c *targetGroups) Update(targetGroup *v1alpha1.TargetGroup) (result *v1alpha1.TargetGroup, err error) {
	result = &v1alpha1.TargetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("targetgroups").
		Name(targetGroup.Name).
		Body(targetGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *targetGroups) UpdateStatus(targetGroup *v1alpha1.TargetGroup) (result *v1alpha1.TargetGroup, err error) {
	result = &v1alpha1.TargetGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("targetgroups").
		Name(targetGroup.Name).
		SubResource("status").
		Body(targetGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the targetGroup and deletes it. Returns an error if one occurs.
func (c *targetGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("targetgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *targetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("targetgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched targetGroup.
func (c *targetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TargetGroup, err error) {
	result = &v1alpha1.TargetGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("targetgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}