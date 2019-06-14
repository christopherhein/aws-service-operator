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

	v1alpha1 "awsoperator.io/pkg/apis/servicediscovery/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PrivateDnsNamespacesGetter has a method to return a PrivateDnsNamespaceInterface.
// A group's client should implement this interface.
type PrivateDnsNamespacesGetter interface {
	PrivateDnsNamespaces(namespace string) PrivateDnsNamespaceInterface
}

// PrivateDnsNamespaceInterface has methods to work with PrivateDnsNamespace resources.
type PrivateDnsNamespaceInterface interface {
	Create(*v1alpha1.PrivateDnsNamespace) (*v1alpha1.PrivateDnsNamespace, error)
	Update(*v1alpha1.PrivateDnsNamespace) (*v1alpha1.PrivateDnsNamespace, error)
	UpdateStatus(*v1alpha1.PrivateDnsNamespace) (*v1alpha1.PrivateDnsNamespace, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PrivateDnsNamespace, error)
	List(opts v1.ListOptions) (*v1alpha1.PrivateDnsNamespaceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrivateDnsNamespace, err error)
	PrivateDnsNamespaceExpansion
}

// privateDnsNamespaces implements PrivateDnsNamespaceInterface
type privateDnsNamespaces struct {
	client rest.Interface
	ns     string
}

// newPrivateDnsNamespaces returns a PrivateDnsNamespaces
func newPrivateDnsNamespaces(c *ServicediscoveryV1alpha1Client, namespace string) *privateDnsNamespaces {
	return &privateDnsNamespaces{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the privateDnsNamespace, and returns the corresponding privateDnsNamespace object, and an error if there is any.
func (c *privateDnsNamespaces) Get(name string, options v1.GetOptions) (result *v1alpha1.PrivateDnsNamespace, err error) {
	result = &v1alpha1.PrivateDnsNamespace{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PrivateDnsNamespaces that match those selectors.
func (c *privateDnsNamespaces) List(opts v1.ListOptions) (result *v1alpha1.PrivateDnsNamespaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PrivateDnsNamespaceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested privateDnsNamespaces.
func (c *privateDnsNamespaces) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a privateDnsNamespace and creates it.  Returns the server's representation of the privateDnsNamespace, and an error, if there is any.
func (c *privateDnsNamespaces) Create(privateDnsNamespace *v1alpha1.PrivateDnsNamespace) (result *v1alpha1.PrivateDnsNamespace, err error) {
	result = &v1alpha1.PrivateDnsNamespace{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		Body(privateDnsNamespace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a privateDnsNamespace and updates it. Returns the server's representation of the privateDnsNamespace, and an error, if there is any.
func (c *privateDnsNamespaces) Update(privateDnsNamespace *v1alpha1.PrivateDnsNamespace) (result *v1alpha1.PrivateDnsNamespace, err error) {
	result = &v1alpha1.PrivateDnsNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		Name(privateDnsNamespace.Name).
		Body(privateDnsNamespace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *privateDnsNamespaces) UpdateStatus(privateDnsNamespace *v1alpha1.PrivateDnsNamespace) (result *v1alpha1.PrivateDnsNamespace, err error) {
	result = &v1alpha1.PrivateDnsNamespace{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		Name(privateDnsNamespace.Name).
		SubResource("status").
		Body(privateDnsNamespace).
		Do().
		Into(result)
	return
}

// Delete takes name of the privateDnsNamespace and deletes it. Returns an error if one occurs.
func (c *privateDnsNamespaces) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *privateDnsNamespaces) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched privateDnsNamespace.
func (c *privateDnsNamespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PrivateDnsNamespace, err error) {
	result = &v1alpha1.PrivateDnsNamespace{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("privatednsnamespaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}