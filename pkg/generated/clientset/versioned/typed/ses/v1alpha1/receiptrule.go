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

	v1alpha1 "awsoperator.io/pkg/apis/ses/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ReceiptRulesGetter has a method to return a ReceiptRuleInterface.
// A group's client should implement this interface.
type ReceiptRulesGetter interface {
	ReceiptRules(namespace string) ReceiptRuleInterface
}

// ReceiptRuleInterface has methods to work with ReceiptRule resources.
type ReceiptRuleInterface interface {
	Create(*v1alpha1.ReceiptRule) (*v1alpha1.ReceiptRule, error)
	Update(*v1alpha1.ReceiptRule) (*v1alpha1.ReceiptRule, error)
	UpdateStatus(*v1alpha1.ReceiptRule) (*v1alpha1.ReceiptRule, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ReceiptRule, error)
	List(opts v1.ListOptions) (*v1alpha1.ReceiptRuleList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReceiptRule, err error)
	ReceiptRuleExpansion
}

// receiptRules implements ReceiptRuleInterface
type receiptRules struct {
	client rest.Interface
	ns     string
}

// newReceiptRules returns a ReceiptRules
func newReceiptRules(c *SesV1alpha1Client, namespace string) *receiptRules {
	return &receiptRules{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the receiptRule, and returns the corresponding receiptRule object, and an error if there is any.
func (c *receiptRules) Get(name string, options v1.GetOptions) (result *v1alpha1.ReceiptRule, err error) {
	result = &v1alpha1.ReceiptRule{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("receiptrules").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ReceiptRules that match those selectors.
func (c *receiptRules) List(opts v1.ListOptions) (result *v1alpha1.ReceiptRuleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ReceiptRuleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("receiptrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested receiptRules.
func (c *receiptRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("receiptrules").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a receiptRule and creates it.  Returns the server's representation of the receiptRule, and an error, if there is any.
func (c *receiptRules) Create(receiptRule *v1alpha1.ReceiptRule) (result *v1alpha1.ReceiptRule, err error) {
	result = &v1alpha1.ReceiptRule{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("receiptrules").
		Body(receiptRule).
		Do().
		Into(result)
	return
}

// Update takes the representation of a receiptRule and updates it. Returns the server's representation of the receiptRule, and an error, if there is any.
func (c *receiptRules) Update(receiptRule *v1alpha1.ReceiptRule) (result *v1alpha1.ReceiptRule, err error) {
	result = &v1alpha1.ReceiptRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("receiptrules").
		Name(receiptRule.Name).
		Body(receiptRule).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *receiptRules) UpdateStatus(receiptRule *v1alpha1.ReceiptRule) (result *v1alpha1.ReceiptRule, err error) {
	result = &v1alpha1.ReceiptRule{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("receiptrules").
		Name(receiptRule.Name).
		SubResource("status").
		Body(receiptRule).
		Do().
		Into(result)
	return
}

// Delete takes name of the receiptRule and deletes it. Returns an error if one occurs.
func (c *receiptRules) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("receiptrules").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *receiptRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("receiptrules").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched receiptRule.
func (c *receiptRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReceiptRule, err error) {
	result = &v1alpha1.ReceiptRule{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("receiptrules").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}