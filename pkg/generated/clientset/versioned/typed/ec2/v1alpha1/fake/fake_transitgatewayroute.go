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
	v1alpha1 "awsoperator.io/pkg/apis/ec2/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTransitGatewayRoutes implements TransitGatewayRouteInterface
type FakeTransitGatewayRoutes struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var transitgatewayroutesResource = schema.GroupVersionResource{Group: "ec2.awsoperator.io", Version: "v1alpha1", Resource: "transitgatewayroutes"}

var transitgatewayroutesKind = schema.GroupVersionKind{Group: "ec2.awsoperator.io", Version: "v1alpha1", Kind: "TransitGatewayRoute"}

// Get takes name of the transitGatewayRoute, and returns the corresponding transitGatewayRoute object, and an error if there is any.
func (c *FakeTransitGatewayRoutes) Get(name string, options v1.GetOptions) (result *v1alpha1.TransitGatewayRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transitgatewayroutesResource, c.ns, name), &v1alpha1.TransitGatewayRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRoute), err
}

// List takes label and field selectors, and returns the list of TransitGatewayRoutes that match those selectors.
func (c *FakeTransitGatewayRoutes) List(opts v1.ListOptions) (result *v1alpha1.TransitGatewayRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transitgatewayroutesResource, transitgatewayroutesKind, c.ns, opts), &v1alpha1.TransitGatewayRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransitGatewayRouteList{ListMeta: obj.(*v1alpha1.TransitGatewayRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransitGatewayRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transitGatewayRoutes.
func (c *FakeTransitGatewayRoutes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transitgatewayroutesResource, c.ns, opts))

}

// Create takes the representation of a transitGatewayRoute and creates it.  Returns the server's representation of the transitGatewayRoute, and an error, if there is any.
func (c *FakeTransitGatewayRoutes) Create(transitGatewayRoute *v1alpha1.TransitGatewayRoute) (result *v1alpha1.TransitGatewayRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transitgatewayroutesResource, c.ns, transitGatewayRoute), &v1alpha1.TransitGatewayRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRoute), err
}

// Update takes the representation of a transitGatewayRoute and updates it. Returns the server's representation of the transitGatewayRoute, and an error, if there is any.
func (c *FakeTransitGatewayRoutes) Update(transitGatewayRoute *v1alpha1.TransitGatewayRoute) (result *v1alpha1.TransitGatewayRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transitgatewayroutesResource, c.ns, transitGatewayRoute), &v1alpha1.TransitGatewayRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransitGatewayRoutes) UpdateStatus(transitGatewayRoute *v1alpha1.TransitGatewayRoute) (*v1alpha1.TransitGatewayRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transitgatewayroutesResource, "status", c.ns, transitGatewayRoute), &v1alpha1.TransitGatewayRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRoute), err
}

// Delete takes name of the transitGatewayRoute and deletes it. Returns an error if one occurs.
func (c *FakeTransitGatewayRoutes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transitgatewayroutesResource, c.ns, name), &v1alpha1.TransitGatewayRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransitGatewayRoutes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transitgatewayroutesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransitGatewayRouteList{})
	return err
}

// Patch applies the patch and returns the patched transitGatewayRoute.
func (c *FakeTransitGatewayRoutes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TransitGatewayRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transitgatewayroutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransitGatewayRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRoute), err
}