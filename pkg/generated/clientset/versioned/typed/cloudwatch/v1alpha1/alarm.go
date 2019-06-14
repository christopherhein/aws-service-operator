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

	v1alpha1 "awsoperator.io/pkg/apis/cloudwatch/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AlarmsGetter has a method to return a AlarmInterface.
// A group's client should implement this interface.
type AlarmsGetter interface {
	Alarms(namespace string) AlarmInterface
}

// AlarmInterface has methods to work with Alarm resources.
type AlarmInterface interface {
	Create(*v1alpha1.Alarm) (*v1alpha1.Alarm, error)
	Update(*v1alpha1.Alarm) (*v1alpha1.Alarm, error)
	UpdateStatus(*v1alpha1.Alarm) (*v1alpha1.Alarm, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Alarm, error)
	List(opts v1.ListOptions) (*v1alpha1.AlarmList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Alarm, err error)
	AlarmExpansion
}

// alarms implements AlarmInterface
type alarms struct {
	client rest.Interface
	ns     string
}

// newAlarms returns a Alarms
func newAlarms(c *CloudwatchV1alpha1Client, namespace string) *alarms {
	return &alarms{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the alarm, and returns the corresponding alarm object, and an error if there is any.
func (c *alarms) Get(name string, options v1.GetOptions) (result *v1alpha1.Alarm, err error) {
	result = &v1alpha1.Alarm{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alarms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Alarms that match those selectors.
func (c *alarms) List(opts v1.ListOptions) (result *v1alpha1.AlarmList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AlarmList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("alarms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested alarms.
func (c *alarms) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("alarms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a alarm and creates it.  Returns the server's representation of the alarm, and an error, if there is any.
func (c *alarms) Create(alarm *v1alpha1.Alarm) (result *v1alpha1.Alarm, err error) {
	result = &v1alpha1.Alarm{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("alarms").
		Body(alarm).
		Do().
		Into(result)
	return
}

// Update takes the representation of a alarm and updates it. Returns the server's representation of the alarm, and an error, if there is any.
func (c *alarms) Update(alarm *v1alpha1.Alarm) (result *v1alpha1.Alarm, err error) {
	result = &v1alpha1.Alarm{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alarms").
		Name(alarm.Name).
		Body(alarm).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *alarms) UpdateStatus(alarm *v1alpha1.Alarm) (result *v1alpha1.Alarm, err error) {
	result = &v1alpha1.Alarm{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("alarms").
		Name(alarm.Name).
		SubResource("status").
		Body(alarm).
		Do().
		Into(result)
	return
}

// Delete takes name of the alarm and deletes it. Returns an error if one occurs.
func (c *alarms) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alarms").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *alarms) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("alarms").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched alarm.
func (c *alarms) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Alarm, err error) {
	result = &v1alpha1.Alarm{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("alarms").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}