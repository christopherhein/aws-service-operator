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
	v1alpha1 "awsoperator.io/pkg/apis/cloudfront/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DistributionLister helps list Distributions.
type DistributionLister interface {
	// List lists all Distributions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Distribution, err error)
	// Distributions returns an object that can list and get Distributions.
	Distributions(namespace string) DistributionNamespaceLister
	DistributionListerExpansion
}

// distributionLister implements the DistributionLister interface.
type distributionLister struct {
	indexer cache.Indexer
}

// NewDistributionLister returns a new DistributionLister.
func NewDistributionLister(indexer cache.Indexer) DistributionLister {
	return &distributionLister{indexer: indexer}
}

// List lists all Distributions in the indexer.
func (s *distributionLister) List(selector labels.Selector) (ret []*v1alpha1.Distribution, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Distribution))
	})
	return ret, err
}

// Distributions returns an object that can list and get Distributions.
func (s *distributionLister) Distributions(namespace string) DistributionNamespaceLister {
	return distributionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DistributionNamespaceLister helps list and get Distributions.
type DistributionNamespaceLister interface {
	// List lists all Distributions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Distribution, err error)
	// Get retrieves the Distribution from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Distribution, error)
	DistributionNamespaceListerExpansion
}

// distributionNamespaceLister implements the DistributionNamespaceLister
// interface.
type distributionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Distributions in the indexer for a given namespace.
func (s distributionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Distribution, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Distribution))
	})
	return ret, err
}

// Get retrieves the Distribution from the indexer for a given namespace and name.
func (s distributionNamespaceLister) Get(name string) (*v1alpha1.Distribution, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("distribution"), name)
	}
	return obj.(*v1alpha1.Distribution), nil
}