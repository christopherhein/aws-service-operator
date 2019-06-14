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

	v1alpha1 "awsoperator.io/pkg/apis/servicecatalog/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloudFormationProvisionedProductsGetter has a method to return a CloudFormationProvisionedProductInterface.
// A group's client should implement this interface.
type CloudFormationProvisionedProductsGetter interface {
	CloudFormationProvisionedProducts(namespace string) CloudFormationProvisionedProductInterface
}

// CloudFormationProvisionedProductInterface has methods to work with CloudFormationProvisionedProduct resources.
type CloudFormationProvisionedProductInterface interface {
	Create(*v1alpha1.CloudFormationProvisionedProduct) (*v1alpha1.CloudFormationProvisionedProduct, error)
	Update(*v1alpha1.CloudFormationProvisionedProduct) (*v1alpha1.CloudFormationProvisionedProduct, error)
	UpdateStatus(*v1alpha1.CloudFormationProvisionedProduct) (*v1alpha1.CloudFormationProvisionedProduct, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.CloudFormationProvisionedProduct, error)
	List(opts v1.ListOptions) (*v1alpha1.CloudFormationProvisionedProductList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudFormationProvisionedProduct, err error)
	CloudFormationProvisionedProductExpansion
}

// cloudFormationProvisionedProducts implements CloudFormationProvisionedProductInterface
type cloudFormationProvisionedProducts struct {
	client rest.Interface
	ns     string
}

// newCloudFormationProvisionedProducts returns a CloudFormationProvisionedProducts
func newCloudFormationProvisionedProducts(c *ServicecatalogV1alpha1Client, namespace string) *cloudFormationProvisionedProducts {
	return &cloudFormationProvisionedProducts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloudFormationProvisionedProduct, and returns the corresponding cloudFormationProvisionedProduct object, and an error if there is any.
func (c *cloudFormationProvisionedProducts) Get(name string, options v1.GetOptions) (result *v1alpha1.CloudFormationProvisionedProduct, err error) {
	result = &v1alpha1.CloudFormationProvisionedProduct{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloudFormationProvisionedProducts that match those selectors.
func (c *cloudFormationProvisionedProducts) List(opts v1.ListOptions) (result *v1alpha1.CloudFormationProvisionedProductList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloudFormationProvisionedProductList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloudFormationProvisionedProducts.
func (c *cloudFormationProvisionedProducts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cloudFormationProvisionedProduct and creates it.  Returns the server's representation of the cloudFormationProvisionedProduct, and an error, if there is any.
func (c *cloudFormationProvisionedProducts) Create(cloudFormationProvisionedProduct *v1alpha1.CloudFormationProvisionedProduct) (result *v1alpha1.CloudFormationProvisionedProduct, err error) {
	result = &v1alpha1.CloudFormationProvisionedProduct{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		Body(cloudFormationProvisionedProduct).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cloudFormationProvisionedProduct and updates it. Returns the server's representation of the cloudFormationProvisionedProduct, and an error, if there is any.
func (c *cloudFormationProvisionedProducts) Update(cloudFormationProvisionedProduct *v1alpha1.CloudFormationProvisionedProduct) (result *v1alpha1.CloudFormationProvisionedProduct, err error) {
	result = &v1alpha1.CloudFormationProvisionedProduct{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		Name(cloudFormationProvisionedProduct.Name).
		Body(cloudFormationProvisionedProduct).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cloudFormationProvisionedProducts) UpdateStatus(cloudFormationProvisionedProduct *v1alpha1.CloudFormationProvisionedProduct) (result *v1alpha1.CloudFormationProvisionedProduct, err error) {
	result = &v1alpha1.CloudFormationProvisionedProduct{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		Name(cloudFormationProvisionedProduct.Name).
		SubResource("status").
		Body(cloudFormationProvisionedProduct).
		Do().
		Into(result)
	return
}

// Delete takes name of the cloudFormationProvisionedProduct and deletes it. Returns an error if one occurs.
func (c *cloudFormationProvisionedProducts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloudFormationProvisionedProducts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cloudFormationProvisionedProduct.
func (c *cloudFormationProvisionedProducts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CloudFormationProvisionedProduct, err error) {
	result = &v1alpha1.CloudFormationProvisionedProduct{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("cloudformationprovisionedproducts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}