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

	v1alpha1 "awsoperator.io/pkg/apis/applicationautoscaling/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ScalableTargetsGetter has a method to return a ScalableTargetInterface.
// A group's client should implement this interface.
type ScalableTargetsGetter interface {
	ScalableTargets(namespace string) ScalableTargetInterface
}

// ScalableTargetInterface has methods to work with ScalableTarget resources.
type ScalableTargetInterface interface {
	Create(*v1alpha1.ScalableTarget) (*v1alpha1.ScalableTarget, error)
	Update(*v1alpha1.ScalableTarget) (*v1alpha1.ScalableTarget, error)
	UpdateStatus(*v1alpha1.ScalableTarget) (*v1alpha1.ScalableTarget, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ScalableTarget, error)
	List(opts v1.ListOptions) (*v1alpha1.ScalableTargetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScalableTarget, err error)
	ScalableTargetExpansion
}

// scalableTargets implements ScalableTargetInterface
type scalableTargets struct {
	client rest.Interface
	ns     string
}

// newScalableTargets returns a ScalableTargets
func newScalableTargets(c *ApplicationautoscalingV1alpha1Client, namespace string) *scalableTargets {
	return &scalableTargets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the scalableTarget, and returns the corresponding scalableTarget object, and an error if there is any.
func (c *scalableTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.ScalableTarget, err error) {
	result = &v1alpha1.ScalableTarget{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scalabletargets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ScalableTargets that match those selectors.
func (c *scalableTargets) List(opts v1.ListOptions) (result *v1alpha1.ScalableTargetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ScalableTargetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("scalabletargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested scalableTargets.
func (c *scalableTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("scalabletargets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a scalableTarget and creates it.  Returns the server's representation of the scalableTarget, and an error, if there is any.
func (c *scalableTargets) Create(scalableTarget *v1alpha1.ScalableTarget) (result *v1alpha1.ScalableTarget, err error) {
	result = &v1alpha1.ScalableTarget{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("scalabletargets").
		Body(scalableTarget).
		Do().
		Into(result)
	return
}

// Update takes the representation of a scalableTarget and updates it. Returns the server's representation of the scalableTarget, and an error, if there is any.
func (c *scalableTargets) Update(scalableTarget *v1alpha1.ScalableTarget) (result *v1alpha1.ScalableTarget, err error) {
	result = &v1alpha1.ScalableTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scalabletargets").
		Name(scalableTarget.Name).
		Body(scalableTarget).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *scalableTargets) UpdateStatus(scalableTarget *v1alpha1.ScalableTarget) (result *v1alpha1.ScalableTarget, err error) {
	result = &v1alpha1.ScalableTarget{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("scalabletargets").
		Name(scalableTarget.Name).
		SubResource("status").
		Body(scalableTarget).
		Do().
		Into(result)
	return
}

// Delete takes name of the scalableTarget and deletes it. Returns an error if one occurs.
func (c *scalableTargets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scalabletargets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *scalableTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("scalabletargets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched scalableTarget.
func (c *scalableTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ScalableTarget, err error) {
	result = &v1alpha1.ScalableTarget{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("scalabletargets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}