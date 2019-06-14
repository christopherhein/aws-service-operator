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

	v1alpha1 "awsoperator.io/pkg/apis/apigatewayv2/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApisGetter has a method to return a ApiInterface.
// A group's client should implement this interface.
type ApisGetter interface {
	Apis(namespace string) ApiInterface
}

// ApiInterface has methods to work with Api resources.
type ApiInterface interface {
	Create(*v1alpha1.Api) (*v1alpha1.Api, error)
	Update(*v1alpha1.Api) (*v1alpha1.Api, error)
	UpdateStatus(*v1alpha1.Api) (*v1alpha1.Api, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Api, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Api, err error)
	ApiExpansion
}

// apis implements ApiInterface
type apis struct {
	client rest.Interface
	ns     string
}

// newApis returns a Apis
func newApis(c *Apigatewayv2V1alpha1Client, namespace string) *apis {
	return &apis{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the api, and returns the corresponding api object, and an error if there is any.
func (c *apis) Get(name string, options v1.GetOptions) (result *v1alpha1.Api, err error) {
	result = &v1alpha1.Api{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apis").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Apis that match those selectors.
func (c *apis) List(opts v1.ListOptions) (result *v1alpha1.ApiList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apis.
func (c *apis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apis").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a api and creates it.  Returns the server's representation of the api, and an error, if there is any.
func (c *apis) Create(api *v1alpha1.Api) (result *v1alpha1.Api, err error) {
	result = &v1alpha1.Api{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apis").
		Body(api).
		Do().
		Into(result)
	return
}

// Update takes the representation of a api and updates it. Returns the server's representation of the api, and an error, if there is any.
func (c *apis) Update(api *v1alpha1.Api) (result *v1alpha1.Api, err error) {
	result = &v1alpha1.Api{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apis").
		Name(api.Name).
		Body(api).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apis) UpdateStatus(api *v1alpha1.Api) (result *v1alpha1.Api, err error) {
	result = &v1alpha1.Api{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apis").
		Name(api.Name).
		SubResource("status").
		Body(api).
		Do().
		Into(result)
	return
}

// Delete takes name of the api and deletes it. Returns an error if one occurs.
func (c *apis) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apis").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apis").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched api.
func (c *apis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Api, err error) {
	result = &v1alpha1.Api{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apis").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}