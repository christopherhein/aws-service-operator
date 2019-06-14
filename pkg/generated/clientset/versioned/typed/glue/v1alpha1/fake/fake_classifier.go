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
	v1alpha1 "awsoperator.io/pkg/apis/glue/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClassifiers implements ClassifierInterface
type FakeClassifiers struct {
	Fake *FakeGlueV1alpha1
	ns   string
}

var classifiersResource = schema.GroupVersionResource{Group: "glue.awsoperator.io", Version: "v1alpha1", Resource: "classifiers"}

var classifiersKind = schema.GroupVersionKind{Group: "glue.awsoperator.io", Version: "v1alpha1", Kind: "Classifier"}

// Get takes name of the classifier, and returns the corresponding classifier object, and an error if there is any.
func (c *FakeClassifiers) Get(name string, options v1.GetOptions) (result *v1alpha1.Classifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(classifiersResource, c.ns, name), &v1alpha1.Classifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Classifier), err
}

// List takes label and field selectors, and returns the list of Classifiers that match those selectors.
func (c *FakeClassifiers) List(opts v1.ListOptions) (result *v1alpha1.ClassifierList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(classifiersResource, classifiersKind, c.ns, opts), &v1alpha1.ClassifierList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClassifierList{ListMeta: obj.(*v1alpha1.ClassifierList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClassifierList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested classifiers.
func (c *FakeClassifiers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(classifiersResource, c.ns, opts))

}

// Create takes the representation of a classifier and creates it.  Returns the server's representation of the classifier, and an error, if there is any.
func (c *FakeClassifiers) Create(classifier *v1alpha1.Classifier) (result *v1alpha1.Classifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(classifiersResource, c.ns, classifier), &v1alpha1.Classifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Classifier), err
}

// Update takes the representation of a classifier and updates it. Returns the server's representation of the classifier, and an error, if there is any.
func (c *FakeClassifiers) Update(classifier *v1alpha1.Classifier) (result *v1alpha1.Classifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(classifiersResource, c.ns, classifier), &v1alpha1.Classifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Classifier), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClassifiers) UpdateStatus(classifier *v1alpha1.Classifier) (*v1alpha1.Classifier, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(classifiersResource, "status", c.ns, classifier), &v1alpha1.Classifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Classifier), err
}

// Delete takes name of the classifier and deletes it. Returns an error if one occurs.
func (c *FakeClassifiers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(classifiersResource, c.ns, name), &v1alpha1.Classifier{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClassifiers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(classifiersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClassifierList{})
	return err
}

// Patch applies the patch and returns the patched classifier.
func (c *FakeClassifiers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Classifier, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(classifiersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Classifier{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Classifier), err
}