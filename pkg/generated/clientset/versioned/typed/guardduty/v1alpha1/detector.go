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

	v1alpha1 "awsoperator.io/pkg/apis/guardduty/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DetectorsGetter has a method to return a DetectorInterface.
// A group's client should implement this interface.
type DetectorsGetter interface {
	Detectors(namespace string) DetectorInterface
}

// DetectorInterface has methods to work with Detector resources.
type DetectorInterface interface {
	Create(*v1alpha1.Detector) (*v1alpha1.Detector, error)
	Update(*v1alpha1.Detector) (*v1alpha1.Detector, error)
	UpdateStatus(*v1alpha1.Detector) (*v1alpha1.Detector, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Detector, error)
	List(opts v1.ListOptions) (*v1alpha1.DetectorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Detector, err error)
	DetectorExpansion
}

// detectors implements DetectorInterface
type detectors struct {
	client rest.Interface
	ns     string
}

// newDetectors returns a Detectors
func newDetectors(c *GuarddutyV1alpha1Client, namespace string) *detectors {
	return &detectors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the detector, and returns the corresponding detector object, and an error if there is any.
func (c *detectors) Get(name string, options v1.GetOptions) (result *v1alpha1.Detector, err error) {
	result = &v1alpha1.Detector{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("detectors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Detectors that match those selectors.
func (c *detectors) List(opts v1.ListOptions) (result *v1alpha1.DetectorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DetectorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("detectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested detectors.
func (c *detectors) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("detectors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a detector and creates it.  Returns the server's representation of the detector, and an error, if there is any.
func (c *detectors) Create(detector *v1alpha1.Detector) (result *v1alpha1.Detector, err error) {
	result = &v1alpha1.Detector{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("detectors").
		Body(detector).
		Do().
		Into(result)
	return
}

// Update takes the representation of a detector and updates it. Returns the server's representation of the detector, and an error, if there is any.
func (c *detectors) Update(detector *v1alpha1.Detector) (result *v1alpha1.Detector, err error) {
	result = &v1alpha1.Detector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("detectors").
		Name(detector.Name).
		Body(detector).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *detectors) UpdateStatus(detector *v1alpha1.Detector) (result *v1alpha1.Detector, err error) {
	result = &v1alpha1.Detector{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("detectors").
		Name(detector.Name).
		SubResource("status").
		Body(detector).
		Do().
		Into(result)
	return
}

// Delete takes name of the detector and deletes it. Returns an error if one occurs.
func (c *detectors) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("detectors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *detectors) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("detectors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched detector.
func (c *detectors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Detector, err error) {
	result = &v1alpha1.Detector{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("detectors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}