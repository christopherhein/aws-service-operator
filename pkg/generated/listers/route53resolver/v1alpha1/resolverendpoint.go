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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "awsoperator.io/pkg/apis/route53resolver/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResolverEndpointLister helps list ResolverEndpoints.
type ResolverEndpointLister interface {
	// List lists all ResolverEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ResolverEndpoint, err error)
	// ResolverEndpoints returns an object that can list and get ResolverEndpoints.
	ResolverEndpoints(namespace string) ResolverEndpointNamespaceLister
	ResolverEndpointListerExpansion
}

// resolverEndpointLister implements the ResolverEndpointLister interface.
type resolverEndpointLister struct {
	indexer cache.Indexer
}

// NewResolverEndpointLister returns a new ResolverEndpointLister.
func NewResolverEndpointLister(indexer cache.Indexer) ResolverEndpointLister {
	return &resolverEndpointLister{indexer: indexer}
}

// List lists all ResolverEndpoints in the indexer.
func (s *resolverEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.ResolverEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResolverEndpoint))
	})
	return ret, err
}

// ResolverEndpoints returns an object that can list and get ResolverEndpoints.
func (s *resolverEndpointLister) ResolverEndpoints(namespace string) ResolverEndpointNamespaceLister {
	return resolverEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResolverEndpointNamespaceLister helps list and get ResolverEndpoints.
type ResolverEndpointNamespaceLister interface {
	// List lists all ResolverEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ResolverEndpoint, err error)
	// Get retrieves the ResolverEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ResolverEndpoint, error)
	ResolverEndpointNamespaceListerExpansion
}

// resolverEndpointNamespaceLister implements the ResolverEndpointNamespaceLister
// interface.
type resolverEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResolverEndpoints in the indexer for a given namespace.
func (s resolverEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResolverEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResolverEndpoint))
	})
	return ret, err
}

// Get retrieves the ResolverEndpoint from the indexer for a given namespace and name.
func (s resolverEndpointNamespaceLister) Get(name string) (*v1alpha1.ResolverEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resolverendpoint"), name)
	}
	return obj.(*v1alpha1.ResolverEndpoint), nil
}