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

// FakeTransitGatewayRouteTableAssociations implements TransitGatewayRouteTableAssociationInterface
type FakeTransitGatewayRouteTableAssociations struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var transitgatewayroutetableassociationsResource = schema.GroupVersionResource{Group: "ec2.awsoperator.io", Version: "v1alpha1", Resource: "transitgatewayroutetableassociations"}

var transitgatewayroutetableassociationsKind = schema.GroupVersionKind{Group: "ec2.awsoperator.io", Version: "v1alpha1", Kind: "TransitGatewayRouteTableAssociation"}

// Get takes name of the transitGatewayRouteTableAssociation, and returns the corresponding transitGatewayRouteTableAssociation object, and an error if there is any.
func (c *FakeTransitGatewayRouteTableAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.TransitGatewayRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transitgatewayroutetableassociationsResource, c.ns, name), &v1alpha1.TransitGatewayRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTableAssociation), err
}

// List takes label and field selectors, and returns the list of TransitGatewayRouteTableAssociations that match those selectors.
func (c *FakeTransitGatewayRouteTableAssociations) List(opts v1.ListOptions) (result *v1alpha1.TransitGatewayRouteTableAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transitgatewayroutetableassociationsResource, transitgatewayroutetableassociationsKind, c.ns, opts), &v1alpha1.TransitGatewayRouteTableAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransitGatewayRouteTableAssociationList{ListMeta: obj.(*v1alpha1.TransitGatewayRouteTableAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransitGatewayRouteTableAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transitGatewayRouteTableAssociations.
func (c *FakeTransitGatewayRouteTableAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transitgatewayroutetableassociationsResource, c.ns, opts))

}

// Create takes the representation of a transitGatewayRouteTableAssociation and creates it.  Returns the server's representation of the transitGatewayRouteTableAssociation, and an error, if there is any.
func (c *FakeTransitGatewayRouteTableAssociations) Create(transitGatewayRouteTableAssociation *v1alpha1.TransitGatewayRouteTableAssociation) (result *v1alpha1.TransitGatewayRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transitgatewayroutetableassociationsResource, c.ns, transitGatewayRouteTableAssociation), &v1alpha1.TransitGatewayRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTableAssociation), err
}

// Update takes the representation of a transitGatewayRouteTableAssociation and updates it. Returns the server's representation of the transitGatewayRouteTableAssociation, and an error, if there is any.
func (c *FakeTransitGatewayRouteTableAssociations) Update(transitGatewayRouteTableAssociation *v1alpha1.TransitGatewayRouteTableAssociation) (result *v1alpha1.TransitGatewayRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transitgatewayroutetableassociationsResource, c.ns, transitGatewayRouteTableAssociation), &v1alpha1.TransitGatewayRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTableAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransitGatewayRouteTableAssociations) UpdateStatus(transitGatewayRouteTableAssociation *v1alpha1.TransitGatewayRouteTableAssociation) (*v1alpha1.TransitGatewayRouteTableAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transitgatewayroutetableassociationsResource, "status", c.ns, transitGatewayRouteTableAssociation), &v1alpha1.TransitGatewayRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTableAssociation), err
}

// Delete takes name of the transitGatewayRouteTableAssociation and deletes it. Returns an error if one occurs.
func (c *FakeTransitGatewayRouteTableAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transitgatewayroutetableassociationsResource, c.ns, name), &v1alpha1.TransitGatewayRouteTableAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransitGatewayRouteTableAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transitgatewayroutetableassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransitGatewayRouteTableAssociationList{})
	return err
}

// Patch applies the patch and returns the patched transitGatewayRouteTableAssociation.
func (c *FakeTransitGatewayRouteTableAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TransitGatewayRouteTableAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transitgatewayroutetableassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransitGatewayRouteTableAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTableAssociation), err
}