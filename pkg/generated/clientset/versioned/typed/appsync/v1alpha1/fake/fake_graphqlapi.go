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
	v1alpha1 "awsoperator.io/pkg/apis/appsync/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGraphQLApis implements GraphQLApiInterface
type FakeGraphQLApis struct {
	Fake *FakeAppsyncV1alpha1
	ns   string
}

var graphqlapisResource = schema.GroupVersionResource{Group: "appsync.awsoperator.io", Version: "v1alpha1", Resource: "graphqlapis"}

var graphqlapisKind = schema.GroupVersionKind{Group: "appsync.awsoperator.io", Version: "v1alpha1", Kind: "GraphQLApi"}

// Get takes name of the graphQLApi, and returns the corresponding graphQLApi object, and an error if there is any.
func (c *FakeGraphQLApis) Get(name string, options v1.GetOptions) (result *v1alpha1.GraphQLApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(graphqlapisResource, c.ns, name), &v1alpha1.GraphQLApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLApi), err
}

// List takes label and field selectors, and returns the list of GraphQLApis that match those selectors.
func (c *FakeGraphQLApis) List(opts v1.ListOptions) (result *v1alpha1.GraphQLApiList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(graphqlapisResource, graphqlapisKind, c.ns, opts), &v1alpha1.GraphQLApiList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GraphQLApiList{ListMeta: obj.(*v1alpha1.GraphQLApiList).ListMeta}
	for _, item := range obj.(*v1alpha1.GraphQLApiList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested graphQLApis.
func (c *FakeGraphQLApis) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(graphqlapisResource, c.ns, opts))

}

// Create takes the representation of a graphQLApi and creates it.  Returns the server's representation of the graphQLApi, and an error, if there is any.
func (c *FakeGraphQLApis) Create(graphQLApi *v1alpha1.GraphQLApi) (result *v1alpha1.GraphQLApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(graphqlapisResource, c.ns, graphQLApi), &v1alpha1.GraphQLApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLApi), err
}

// Update takes the representation of a graphQLApi and updates it. Returns the server's representation of the graphQLApi, and an error, if there is any.
func (c *FakeGraphQLApis) Update(graphQLApi *v1alpha1.GraphQLApi) (result *v1alpha1.GraphQLApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(graphqlapisResource, c.ns, graphQLApi), &v1alpha1.GraphQLApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLApi), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGraphQLApis) UpdateStatus(graphQLApi *v1alpha1.GraphQLApi) (*v1alpha1.GraphQLApi, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(graphqlapisResource, "status", c.ns, graphQLApi), &v1alpha1.GraphQLApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLApi), err
}

// Delete takes name of the graphQLApi and deletes it. Returns an error if one occurs.
func (c *FakeGraphQLApis) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(graphqlapisResource, c.ns, name), &v1alpha1.GraphQLApi{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGraphQLApis) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(graphqlapisResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GraphQLApiList{})
	return err
}

// Patch applies the patch and returns the patched graphQLApi.
func (c *FakeGraphQLApis) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GraphQLApi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(graphqlapisResource, c.ns, name, pt, data, subresources...), &v1alpha1.GraphQLApi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLApi), err
}