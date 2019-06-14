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

// FakeGraphQLSchemas implements GraphQLSchemaInterface
type FakeGraphQLSchemas struct {
	Fake *FakeAppsyncV1alpha1
	ns   string
}

var graphqlschemasResource = schema.GroupVersionResource{Group: "appsync.awsoperator.io", Version: "v1alpha1", Resource: "graphqlschemas"}

var graphqlschemasKind = schema.GroupVersionKind{Group: "appsync.awsoperator.io", Version: "v1alpha1", Kind: "GraphQLSchema"}

// Get takes name of the graphQLSchema, and returns the corresponding graphQLSchema object, and an error if there is any.
func (c *FakeGraphQLSchemas) Get(name string, options v1.GetOptions) (result *v1alpha1.GraphQLSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(graphqlschemasResource, c.ns, name), &v1alpha1.GraphQLSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLSchema), err
}

// List takes label and field selectors, and returns the list of GraphQLSchemas that match those selectors.
func (c *FakeGraphQLSchemas) List(opts v1.ListOptions) (result *v1alpha1.GraphQLSchemaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(graphqlschemasResource, graphqlschemasKind, c.ns, opts), &v1alpha1.GraphQLSchemaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GraphQLSchemaList{ListMeta: obj.(*v1alpha1.GraphQLSchemaList).ListMeta}
	for _, item := range obj.(*v1alpha1.GraphQLSchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested graphQLSchemas.
func (c *FakeGraphQLSchemas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(graphqlschemasResource, c.ns, opts))

}

// Create takes the representation of a graphQLSchema and creates it.  Returns the server's representation of the graphQLSchema, and an error, if there is any.
func (c *FakeGraphQLSchemas) Create(graphQLSchema *v1alpha1.GraphQLSchema) (result *v1alpha1.GraphQLSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(graphqlschemasResource, c.ns, graphQLSchema), &v1alpha1.GraphQLSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLSchema), err
}

// Update takes the representation of a graphQLSchema and updates it. Returns the server's representation of the graphQLSchema, and an error, if there is any.
func (c *FakeGraphQLSchemas) Update(graphQLSchema *v1alpha1.GraphQLSchema) (result *v1alpha1.GraphQLSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(graphqlschemasResource, c.ns, graphQLSchema), &v1alpha1.GraphQLSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLSchema), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGraphQLSchemas) UpdateStatus(graphQLSchema *v1alpha1.GraphQLSchema) (*v1alpha1.GraphQLSchema, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(graphqlschemasResource, "status", c.ns, graphQLSchema), &v1alpha1.GraphQLSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLSchema), err
}

// Delete takes name of the graphQLSchema and deletes it. Returns an error if one occurs.
func (c *FakeGraphQLSchemas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(graphqlschemasResource, c.ns, name), &v1alpha1.GraphQLSchema{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGraphQLSchemas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(graphqlschemasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.GraphQLSchemaList{})
	return err
}

// Patch applies the patch and returns the patched graphQLSchema.
func (c *FakeGraphQLSchemas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GraphQLSchema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(graphqlschemasResource, c.ns, name, pt, data, subresources...), &v1alpha1.GraphQLSchema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GraphQLSchema), err
}