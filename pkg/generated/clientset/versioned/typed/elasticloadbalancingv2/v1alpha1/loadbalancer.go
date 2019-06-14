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

// LoadBalancersGetter has a method to return a LoadBalancerInterface.
// A group's client should implement this interface.
type LoadBalancersGetter interface {
	LoadBalancers(namespace string) LoadBalancerInterface
}

// LoadBalancerInterface has methods to work with LoadBalancer resources.
type LoadBalancerInterface interface {
	Create(*v1alpha1.LoadBalancer) (*v1alpha1.LoadBalancer, error)
	Update(*v1alpha1.LoadBalancer) (*v1alpha1.LoadBalancer, error)
	UpdateStatus(*v1alpha1.LoadBalancer) (*v1alpha1.LoadBalancer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LoadBalancer, error)
	List(opts v1.ListOptions) (*v1alpha1.LoadBalancerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoadBalancer, err error)
	LoadBalancerExpansion
}

// loadBalancers implements LoadBalancerInterface
type loadBalancers struct {
	client rest.Interface
	ns     string
}

// newLoadBalancers returns a LoadBalancers
func newLoadBalancers(c *Elasticloadbalancingv2V1alpha1Client, namespace string) *loadBalancers {
	return &loadBalancers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the loadBalancer, and returns the corresponding loadBalancer object, and an error if there is any.
func (c *loadBalancers) Get(name string, options v1.GetOptions) (result *v1alpha1.LoadBalancer, err error) {
	result = &v1alpha1.LoadBalancer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LoadBalancers that match those selectors.
func (c *loadBalancers) List(opts v1.ListOptions) (result *v1alpha1.LoadBalancerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LoadBalancerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loadBalancers.
func (c *loadBalancers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loadbalancers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a loadBalancer and creates it.  Returns the server's representation of the loadBalancer, and an error, if there is any.
func (c *loadBalancers) Create(loadBalancer *v1alpha1.LoadBalancer) (result *v1alpha1.LoadBalancer, err error) {
	result = &v1alpha1.LoadBalancer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loadbalancers").
		Body(loadBalancer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a loadBalancer and updates it. Returns the server's representation of the loadBalancer, and an error, if there is any.
func (c *loadBalancers) Update(loadBalancer *v1alpha1.LoadBalancer) (result *v1alpha1.LoadBalancer, err error) {
	result = &v1alpha1.LoadBalancer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loadbalancers").
		Name(loadBalancer.Name).
		Body(loadBalancer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *loadBalancers) UpdateStatus(loadBalancer *v1alpha1.LoadBalancer) (result *v1alpha1.LoadBalancer, err error) {
	result = &v1alpha1.LoadBalancer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loadbalancers").
		Name(loadBalancer.Name).
		SubResource("status").
		Body(loadBalancer).
		Do().
		Into(result)
	return
}

// Delete takes name of the loadBalancer and deletes it. Returns an error if one occurs.
func (c *loadBalancers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loadbalancers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loadBalancers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loadbalancers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched loadBalancer.
func (c *loadBalancers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LoadBalancer, err error) {
	result = &v1alpha1.LoadBalancer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loadbalancers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}