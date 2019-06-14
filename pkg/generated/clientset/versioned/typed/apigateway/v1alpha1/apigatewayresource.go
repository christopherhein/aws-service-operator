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

	v1alpha1 "awsoperator.io/pkg/apis/apigateway/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApiGatewayResourcesGetter has a method to return a ApiGatewayResourceInterface.
// A group's client should implement this interface.
type ApiGatewayResourcesGetter interface {
	ApiGatewayResources(namespace string) ApiGatewayResourceInterface
}

// ApiGatewayResourceInterface has methods to work with ApiGatewayResource resources.
type ApiGatewayResourceInterface interface {
	Create(*v1alpha1.ApiGatewayResource) (*v1alpha1.ApiGatewayResource, error)
	Update(*v1alpha1.ApiGatewayResource) (*v1alpha1.ApiGatewayResource, error)
	UpdateStatus(*v1alpha1.ApiGatewayResource) (*v1alpha1.ApiGatewayResource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiGatewayResource, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiGatewayResourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayResource, err error)
	ApiGatewayResourceExpansion
}

// apiGatewayResources implements ApiGatewayResourceInterface
type apiGatewayResources struct {
	client rest.Interface
	ns     string
}

// newApiGatewayResources returns a ApiGatewayResources
func newApiGatewayResources(c *ApigatewayV1alpha1Client, namespace string) *apiGatewayResources {
	return &apiGatewayResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiGatewayResource, and returns the corresponding apiGatewayResource object, and an error if there is any.
func (c *apiGatewayResources) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiGatewayResource, err error) {
	result = &v1alpha1.ApiGatewayResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiGatewayResources that match those selectors.
func (c *apiGatewayResources) List(opts v1.ListOptions) (result *v1alpha1.ApiGatewayResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiGatewayResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiGatewayResources.
func (c *apiGatewayResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apigatewayresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiGatewayResource and creates it.  Returns the server's representation of the apiGatewayResource, and an error, if there is any.
func (c *apiGatewayResources) Create(apiGatewayResource *v1alpha1.ApiGatewayResource) (result *v1alpha1.ApiGatewayResource, err error) {
	result = &v1alpha1.ApiGatewayResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apigatewayresources").
		Body(apiGatewayResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiGatewayResource and updates it. Returns the server's representation of the apiGatewayResource, and an error, if there is any.
func (c *apiGatewayResources) Update(apiGatewayResource *v1alpha1.ApiGatewayResource) (result *v1alpha1.ApiGatewayResource, err error) {
	result = &v1alpha1.ApiGatewayResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewayresources").
		Name(apiGatewayResource.Name).
		Body(apiGatewayResource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiGatewayResources) UpdateStatus(apiGatewayResource *v1alpha1.ApiGatewayResource) (result *v1alpha1.ApiGatewayResource, err error) {
	result = &v1alpha1.ApiGatewayResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apigatewayresources").
		Name(apiGatewayResource.Name).
		SubResource("status").
		Body(apiGatewayResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiGatewayResource and deletes it. Returns an error if one occurs.
func (c *apiGatewayResources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewayresources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiGatewayResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apigatewayresources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiGatewayResource.
func (c *apiGatewayResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiGatewayResource, err error) {
	result = &v1alpha1.ApiGatewayResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apigatewayresources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}