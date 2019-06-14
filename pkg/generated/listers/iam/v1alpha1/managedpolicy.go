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
	v1alpha1 "awsoperator.io/pkg/apis/iam/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagedPolicyLister helps list ManagedPolicies.
type ManagedPolicyLister interface {
	// List lists all ManagedPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedPolicy, err error)
	// ManagedPolicies returns an object that can list and get ManagedPolicies.
	ManagedPolicies(namespace string) ManagedPolicyNamespaceLister
	ManagedPolicyListerExpansion
}

// managedPolicyLister implements the ManagedPolicyLister interface.
type managedPolicyLister struct {
	indexer cache.Indexer
}

// NewManagedPolicyLister returns a new ManagedPolicyLister.
func NewManagedPolicyLister(indexer cache.Indexer) ManagedPolicyLister {
	return &managedPolicyLister{indexer: indexer}
}

// List lists all ManagedPolicies in the indexer.
func (s *managedPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedPolicy))
	})
	return ret, err
}

// ManagedPolicies returns an object that can list and get ManagedPolicies.
func (s *managedPolicyLister) ManagedPolicies(namespace string) ManagedPolicyNamespaceLister {
	return managedPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagedPolicyNamespaceLister helps list and get ManagedPolicies.
type ManagedPolicyNamespaceLister interface {
	// List lists all ManagedPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedPolicy, err error)
	// Get retrieves the ManagedPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ManagedPolicy, error)
	ManagedPolicyNamespaceListerExpansion
}

// managedPolicyNamespaceLister implements the ManagedPolicyNamespaceLister
// interface.
type managedPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagedPolicies in the indexer for a given namespace.
func (s managedPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedPolicy))
	})
	return ret, err
}

// Get retrieves the ManagedPolicy from the indexer for a given namespace and name.
func (s managedPolicyNamespaceLister) Get(name string) (*v1alpha1.ManagedPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managedpolicy"), name)
	}
	return obj.(*v1alpha1.ManagedPolicy), nil
}