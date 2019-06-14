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
	v1alpha1 "awsoperator.io/pkg/apis/servicecatalog/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTagOptions implements TagOptionInterface
type FakeTagOptions struct {
	Fake *FakeServicecatalogV1alpha1
	ns   string
}

var tagoptionsResource = schema.GroupVersionResource{Group: "servicecatalog.awsoperator.io", Version: "v1alpha1", Resource: "tagoptions"}

var tagoptionsKind = schema.GroupVersionKind{Group: "servicecatalog.awsoperator.io", Version: "v1alpha1", Kind: "TagOption"}

// Get takes name of the tagOption, and returns the corresponding tagOption object, and an error if there is any.
func (c *FakeTagOptions) Get(name string, options v1.GetOptions) (result *v1alpha1.TagOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tagoptionsResource, c.ns, name), &v1alpha1.TagOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagOption), err
}

// List takes label and field selectors, and returns the list of TagOptions that match those selectors.
func (c *FakeTagOptions) List(opts v1.ListOptions) (result *v1alpha1.TagOptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tagoptionsResource, tagoptionsKind, c.ns, opts), &v1alpha1.TagOptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TagOptionList{ListMeta: obj.(*v1alpha1.TagOptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.TagOptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tagOptions.
func (c *FakeTagOptions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tagoptionsResource, c.ns, opts))

}

// Create takes the representation of a tagOption and creates it.  Returns the server's representation of the tagOption, and an error, if there is any.
func (c *FakeTagOptions) Create(tagOption *v1alpha1.TagOption) (result *v1alpha1.TagOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tagoptionsResource, c.ns, tagOption), &v1alpha1.TagOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagOption), err
}

// Update takes the representation of a tagOption and updates it. Returns the server's representation of the tagOption, and an error, if there is any.
func (c *FakeTagOptions) Update(tagOption *v1alpha1.TagOption) (result *v1alpha1.TagOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tagoptionsResource, c.ns, tagOption), &v1alpha1.TagOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagOption), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTagOptions) UpdateStatus(tagOption *v1alpha1.TagOption) (*v1alpha1.TagOption, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tagoptionsResource, "status", c.ns, tagOption), &v1alpha1.TagOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagOption), err
}

// Delete takes name of the tagOption and deletes it. Returns an error if one occurs.
func (c *FakeTagOptions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tagoptionsResource, c.ns, name), &v1alpha1.TagOption{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTagOptions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tagoptionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TagOptionList{})
	return err
}

// Patch applies the patch and returns the patched tagOption.
func (c *FakeTagOptions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TagOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tagoptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TagOption{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TagOption), err
}