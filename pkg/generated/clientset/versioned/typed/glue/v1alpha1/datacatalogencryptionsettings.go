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

	v1alpha1 "awsoperator.io/pkg/apis/glue/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DataCatalogEncryptionSettingsesGetter has a method to return a DataCatalogEncryptionSettingsInterface.
// A group's client should implement this interface.
type DataCatalogEncryptionSettingsesGetter interface {
	DataCatalogEncryptionSettingses(namespace string) DataCatalogEncryptionSettingsInterface
}

// DataCatalogEncryptionSettingsInterface has methods to work with DataCatalogEncryptionSettings resources.
type DataCatalogEncryptionSettingsInterface interface {
	Create(*v1alpha1.DataCatalogEncryptionSettings) (*v1alpha1.DataCatalogEncryptionSettings, error)
	Update(*v1alpha1.DataCatalogEncryptionSettings) (*v1alpha1.DataCatalogEncryptionSettings, error)
	UpdateStatus(*v1alpha1.DataCatalogEncryptionSettings) (*v1alpha1.DataCatalogEncryptionSettings, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DataCatalogEncryptionSettings, error)
	List(opts v1.ListOptions) (*v1alpha1.DataCatalogEncryptionSettingsList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataCatalogEncryptionSettings, err error)
	DataCatalogEncryptionSettingsExpansion
}

// dataCatalogEncryptionSettingses implements DataCatalogEncryptionSettingsInterface
type dataCatalogEncryptionSettingses struct {
	client rest.Interface
	ns     string
}

// newDataCatalogEncryptionSettingses returns a DataCatalogEncryptionSettingses
func newDataCatalogEncryptionSettingses(c *GlueV1alpha1Client, namespace string) *dataCatalogEncryptionSettingses {
	return &dataCatalogEncryptionSettingses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataCatalogEncryptionSettings, and returns the corresponding dataCatalogEncryptionSettings object, and an error if there is any.
func (c *dataCatalogEncryptionSettingses) Get(name string, options v1.GetOptions) (result *v1alpha1.DataCatalogEncryptionSettings, err error) {
	result = &v1alpha1.DataCatalogEncryptionSettings{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DataCatalogEncryptionSettingses that match those selectors.
func (c *dataCatalogEncryptionSettingses) List(opts v1.ListOptions) (result *v1alpha1.DataCatalogEncryptionSettingsList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DataCatalogEncryptionSettingsList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dataCatalogEncryptionSettingses.
func (c *dataCatalogEncryptionSettingses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataCatalogEncryptionSettings and creates it.  Returns the server's representation of the dataCatalogEncryptionSettings, and an error, if there is any.
func (c *dataCatalogEncryptionSettingses) Create(dataCatalogEncryptionSettings *v1alpha1.DataCatalogEncryptionSettings) (result *v1alpha1.DataCatalogEncryptionSettings, err error) {
	result = &v1alpha1.DataCatalogEncryptionSettings{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		Body(dataCatalogEncryptionSettings).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataCatalogEncryptionSettings and updates it. Returns the server's representation of the dataCatalogEncryptionSettings, and an error, if there is any.
func (c *dataCatalogEncryptionSettingses) Update(dataCatalogEncryptionSettings *v1alpha1.DataCatalogEncryptionSettings) (result *v1alpha1.DataCatalogEncryptionSettings, err error) {
	result = &v1alpha1.DataCatalogEncryptionSettings{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		Name(dataCatalogEncryptionSettings.Name).
		Body(dataCatalogEncryptionSettings).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *dataCatalogEncryptionSettingses) UpdateStatus(dataCatalogEncryptionSettings *v1alpha1.DataCatalogEncryptionSettings) (result *v1alpha1.DataCatalogEncryptionSettings, err error) {
	result = &v1alpha1.DataCatalogEncryptionSettings{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		Name(dataCatalogEncryptionSettings.Name).
		SubResource("status").
		Body(dataCatalogEncryptionSettings).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataCatalogEncryptionSettings and deletes it. Returns an error if one occurs.
func (c *dataCatalogEncryptionSettingses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dataCatalogEncryptionSettingses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataCatalogEncryptionSettings.
func (c *dataCatalogEncryptionSettingses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DataCatalogEncryptionSettings, err error) {
	result = &v1alpha1.DataCatalogEncryptionSettings{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datacatalogencryptionsettingses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}