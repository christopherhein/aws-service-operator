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

	v1alpha1 "awsoperator.io/pkg/apis/redshift/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterParameterGroupsGetter has a method to return a ClusterParameterGroupInterface.
// A group's client should implement this interface.
type ClusterParameterGroupsGetter interface {
	ClusterParameterGroups(namespace string) ClusterParameterGroupInterface
}

// ClusterParameterGroupInterface has methods to work with ClusterParameterGroup resources.
type ClusterParameterGroupInterface interface {
	Create(*v1alpha1.ClusterParameterGroup) (*v1alpha1.ClusterParameterGroup, error)
	Update(*v1alpha1.ClusterParameterGroup) (*v1alpha1.ClusterParameterGroup, error)
	UpdateStatus(*v1alpha1.ClusterParameterGroup) (*v1alpha1.ClusterParameterGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterParameterGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterParameterGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterParameterGroup, err error)
	ClusterParameterGroupExpansion
}

// clusterParameterGroups implements ClusterParameterGroupInterface
type clusterParameterGroups struct {
	client rest.Interface
	ns     string
}

// newClusterParameterGroups returns a ClusterParameterGroups
func newClusterParameterGroups(c *RedshiftV1alpha1Client, namespace string) *clusterParameterGroups {
	return &clusterParameterGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterParameterGroup, and returns the corresponding clusterParameterGroup object, and an error if there is any.
func (c *clusterParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterParameterGroup, err error) {
	result = &v1alpha1.ClusterParameterGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterparametergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterParameterGroups that match those selectors.
func (c *clusterParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.ClusterParameterGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterParameterGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterParameterGroups.
func (c *clusterParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterparametergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clusterParameterGroup and creates it.  Returns the server's representation of the clusterParameterGroup, and an error, if there is any.
func (c *clusterParameterGroups) Create(clusterParameterGroup *v1alpha1.ClusterParameterGroup) (result *v1alpha1.ClusterParameterGroup, err error) {
	result = &v1alpha1.ClusterParameterGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterparametergroups").
		Body(clusterParameterGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterParameterGroup and updates it. Returns the server's representation of the clusterParameterGroup, and an error, if there is any.
func (c *clusterParameterGroups) Update(clusterParameterGroup *v1alpha1.ClusterParameterGroup) (result *v1alpha1.ClusterParameterGroup, err error) {
	result = &v1alpha1.ClusterParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterparametergroups").
		Name(clusterParameterGroup.Name).
		Body(clusterParameterGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterParameterGroups) UpdateStatus(clusterParameterGroup *v1alpha1.ClusterParameterGroup) (result *v1alpha1.ClusterParameterGroup, err error) {
	result = &v1alpha1.ClusterParameterGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterparametergroups").
		Name(clusterParameterGroup.Name).
		SubResource("status").
		Body(clusterParameterGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterParameterGroup and deletes it. Returns an error if one occurs.
func (c *clusterParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterparametergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterparametergroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterParameterGroup.
func (c *clusterParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterParameterGroup, err error) {
	result = &v1alpha1.ClusterParameterGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterparametergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}