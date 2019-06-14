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

// ModelsGetter has a method to return a ModelInterface.
// A group's client should implement this interface.
type ModelsGetter interface {
	Models(namespace string) ModelInterface
}

// ModelInterface has methods to work with Model resources.
type ModelInterface interface {
	Create(*v1alpha1.Model) (*v1alpha1.Model, error)
	Update(*v1alpha1.Model) (*v1alpha1.Model, error)
	UpdateStatus(*v1alpha1.Model) (*v1alpha1.Model, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Model, error)
	List(opts v1.ListOptions) (*v1alpha1.ModelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Model, err error)
	ModelExpansion
}

// models implements ModelInterface
type models struct {
	client rest.Interface
	ns     string
}

// newModels returns a Models
func newModels(c *Apigatewayv2V1alpha1Client, namespace string) *models {
	return &models{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the model, and returns the corresponding model object, and an error if there is any.
func (c *models) Get(name string, options v1.GetOptions) (result *v1alpha1.Model, err error) {
	result = &v1alpha1.Model{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("models").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Models that match those selectors.
func (c *models) List(opts v1.ListOptions) (result *v1alpha1.ModelList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ModelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("models").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested models.
func (c *models) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("models").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a model and creates it.  Returns the server's representation of the model, and an error, if there is any.
func (c *models) Create(model *v1alpha1.Model) (result *v1alpha1.Model, err error) {
	result = &v1alpha1.Model{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("models").
		Body(model).
		Do().
		Into(result)
	return
}

// Update takes the representation of a model and updates it. Returns the server's representation of the model, and an error, if there is any.
func (c *models) Update(model *v1alpha1.Model) (result *v1alpha1.Model, err error) {
	result = &v1alpha1.Model{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("models").
		Name(model.Name).
		Body(model).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *models) UpdateStatus(model *v1alpha1.Model) (result *v1alpha1.Model, err error) {
	result = &v1alpha1.Model{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("models").
		Name(model.Name).
		SubResource("status").
		Body(model).
		Do().
		Into(result)
	return
}

// Delete takes name of the model and deletes it. Returns an error if one occurs.
func (c *models) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("models").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *models) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("models").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched model.
func (c *models) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Model, err error) {
	result = &v1alpha1.Model{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("models").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}