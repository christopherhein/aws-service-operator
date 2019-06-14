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

package fake

import (
	v1alpha1 "awsoperator.io/pkg/apis/iotanalytics/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDatasets implements DatasetInterface
type FakeDatasets struct {
	Fake *FakeIotanalyticsV1alpha1
	ns   string
}

var datasetsResource = schema.GroupVersionResource{Group: "iotanalytics.awsoperator.io", Version: "v1alpha1", Resource: "datasets"}

var datasetsKind = schema.GroupVersionKind{Group: "iotanalytics.awsoperator.io", Version: "v1alpha1", Kind: "Dataset"}

// Get takes name of the dataset, and returns the corresponding dataset object, and an error if there is any.
func (c *FakeDatasets) Get(name string, options v1.GetOptions) (result *v1alpha1.Dataset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(datasetsResource, c.ns, name), &v1alpha1.Dataset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dataset), err
}

// List takes label and field selectors, and returns the list of Datasets that match those selectors.
func (c *FakeDatasets) List(opts v1.ListOptions) (result *v1alpha1.DatasetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(datasetsResource, datasetsKind, c.ns, opts), &v1alpha1.DatasetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DatasetList{ListMeta: obj.(*v1alpha1.DatasetList).ListMeta}
	for _, item := range obj.(*v1alpha1.DatasetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested datasets.
func (c *FakeDatasets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(datasetsResource, c.ns, opts))

}

// Create takes the representation of a dataset and creates it.  Returns the server's representation of the dataset, and an error, if there is any.
func (c *FakeDatasets) Create(dataset *v1alpha1.Dataset) (result *v1alpha1.Dataset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(datasetsResource, c.ns, dataset), &v1alpha1.Dataset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dataset), err
}

// Update takes the representation of a dataset and updates it. Returns the server's representation of the dataset, and an error, if there is any.
func (c *FakeDatasets) Update(dataset *v1alpha1.Dataset) (result *v1alpha1.Dataset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(datasetsResource, c.ns, dataset), &v1alpha1.Dataset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dataset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDatasets) UpdateStatus(dataset *v1alpha1.Dataset) (*v1alpha1.Dataset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(datasetsResource, "status", c.ns, dataset), &v1alpha1.Dataset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dataset), err
}

// Delete takes name of the dataset and deletes it. Returns an error if one occurs.
func (c *FakeDatasets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(datasetsResource, c.ns, name), &v1alpha1.Dataset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDatasets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(datasetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DatasetList{})
	return err
}

// Patch applies the patch and returns the patched dataset.
func (c *FakeDatasets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Dataset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(datasetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Dataset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Dataset), err
}