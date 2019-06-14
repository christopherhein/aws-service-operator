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
	v1alpha1 "awsoperator.io/pkg/apis/wafregional/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRateBasedRules implements RateBasedRuleInterface
type FakeRateBasedRules struct {
	Fake *FakeWafregionalV1alpha1
	ns   string
}

var ratebasedrulesResource = schema.GroupVersionResource{Group: "wafregional.awsoperator.io", Version: "v1alpha1", Resource: "ratebasedrules"}

var ratebasedrulesKind = schema.GroupVersionKind{Group: "wafregional.awsoperator.io", Version: "v1alpha1", Kind: "RateBasedRule"}

// Get takes name of the rateBasedRule, and returns the corresponding rateBasedRule object, and an error if there is any.
func (c *FakeRateBasedRules) Get(name string, options v1.GetOptions) (result *v1alpha1.RateBasedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ratebasedrulesResource, c.ns, name), &v1alpha1.RateBasedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RateBasedRule), err
}

// List takes label and field selectors, and returns the list of RateBasedRules that match those selectors.
func (c *FakeRateBasedRules) List(opts v1.ListOptions) (result *v1alpha1.RateBasedRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ratebasedrulesResource, ratebasedrulesKind, c.ns, opts), &v1alpha1.RateBasedRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RateBasedRuleList{ListMeta: obj.(*v1alpha1.RateBasedRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.RateBasedRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rateBasedRules.
func (c *FakeRateBasedRules) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ratebasedrulesResource, c.ns, opts))

}

// Create takes the representation of a rateBasedRule and creates it.  Returns the server's representation of the rateBasedRule, and an error, if there is any.
func (c *FakeRateBasedRules) Create(rateBasedRule *v1alpha1.RateBasedRule) (result *v1alpha1.RateBasedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ratebasedrulesResource, c.ns, rateBasedRule), &v1alpha1.RateBasedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RateBasedRule), err
}

// Update takes the representation of a rateBasedRule and updates it. Returns the server's representation of the rateBasedRule, and an error, if there is any.
func (c *FakeRateBasedRules) Update(rateBasedRule *v1alpha1.RateBasedRule) (result *v1alpha1.RateBasedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ratebasedrulesResource, c.ns, rateBasedRule), &v1alpha1.RateBasedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RateBasedRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRateBasedRules) UpdateStatus(rateBasedRule *v1alpha1.RateBasedRule) (*v1alpha1.RateBasedRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ratebasedrulesResource, "status", c.ns, rateBasedRule), &v1alpha1.RateBasedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RateBasedRule), err
}

// Delete takes name of the rateBasedRule and deletes it. Returns an error if one occurs.
func (c *FakeRateBasedRules) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ratebasedrulesResource, c.ns, name), &v1alpha1.RateBasedRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRateBasedRules) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ratebasedrulesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RateBasedRuleList{})
	return err
}

// Patch applies the patch and returns the patched rateBasedRule.
func (c *FakeRateBasedRules) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RateBasedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ratebasedrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RateBasedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RateBasedRule), err
}