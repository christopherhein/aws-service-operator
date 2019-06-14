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
	v1alpha1 "awsoperator.io/pkg/apis/config/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAggregationAuthorizations implements AggregationAuthorizationInterface
type FakeAggregationAuthorizations struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var aggregationauthorizationsResource = schema.GroupVersionResource{Group: "config.awsoperator.io", Version: "v1alpha1", Resource: "aggregationauthorizations"}

var aggregationauthorizationsKind = schema.GroupVersionKind{Group: "config.awsoperator.io", Version: "v1alpha1", Kind: "AggregationAuthorization"}

// Get takes name of the aggregationAuthorization, and returns the corresponding aggregationAuthorization object, and an error if there is any.
func (c *FakeAggregationAuthorizations) Get(name string, options v1.GetOptions) (result *v1alpha1.AggregationAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(aggregationauthorizationsResource, c.ns, name), &v1alpha1.AggregationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregationAuthorization), err
}

// List takes label and field selectors, and returns the list of AggregationAuthorizations that match those selectors.
func (c *FakeAggregationAuthorizations) List(opts v1.ListOptions) (result *v1alpha1.AggregationAuthorizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(aggregationauthorizationsResource, aggregationauthorizationsKind, c.ns, opts), &v1alpha1.AggregationAuthorizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AggregationAuthorizationList{ListMeta: obj.(*v1alpha1.AggregationAuthorizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.AggregationAuthorizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aggregationAuthorizations.
func (c *FakeAggregationAuthorizations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(aggregationauthorizationsResource, c.ns, opts))

}

// Create takes the representation of a aggregationAuthorization and creates it.  Returns the server's representation of the aggregationAuthorization, and an error, if there is any.
func (c *FakeAggregationAuthorizations) Create(aggregationAuthorization *v1alpha1.AggregationAuthorization) (result *v1alpha1.AggregationAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(aggregationauthorizationsResource, c.ns, aggregationAuthorization), &v1alpha1.AggregationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregationAuthorization), err
}

// Update takes the representation of a aggregationAuthorization and updates it. Returns the server's representation of the aggregationAuthorization, and an error, if there is any.
func (c *FakeAggregationAuthorizations) Update(aggregationAuthorization *v1alpha1.AggregationAuthorization) (result *v1alpha1.AggregationAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(aggregationauthorizationsResource, c.ns, aggregationAuthorization), &v1alpha1.AggregationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregationAuthorization), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAggregationAuthorizations) UpdateStatus(aggregationAuthorization *v1alpha1.AggregationAuthorization) (*v1alpha1.AggregationAuthorization, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(aggregationauthorizationsResource, "status", c.ns, aggregationAuthorization), &v1alpha1.AggregationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregationAuthorization), err
}

// Delete takes name of the aggregationAuthorization and deletes it. Returns an error if one occurs.
func (c *FakeAggregationAuthorizations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(aggregationauthorizationsResource, c.ns, name), &v1alpha1.AggregationAuthorization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAggregationAuthorizations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(aggregationauthorizationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AggregationAuthorizationList{})
	return err
}

// Patch applies the patch and returns the patched aggregationAuthorization.
func (c *FakeAggregationAuthorizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AggregationAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(aggregationauthorizationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AggregationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AggregationAuthorization), err
}