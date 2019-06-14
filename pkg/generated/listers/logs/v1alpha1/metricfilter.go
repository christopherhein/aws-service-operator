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
	v1alpha1 "awsoperator.io/pkg/apis/logs/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MetricFilterLister helps list MetricFilters.
type MetricFilterLister interface {
	// List lists all MetricFilters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.MetricFilter, err error)
	// MetricFilters returns an object that can list and get MetricFilters.
	MetricFilters(namespace string) MetricFilterNamespaceLister
	MetricFilterListerExpansion
}

// metricFilterLister implements the MetricFilterLister interface.
type metricFilterLister struct {
	indexer cache.Indexer
}

// NewMetricFilterLister returns a new MetricFilterLister.
func NewMetricFilterLister(indexer cache.Indexer) MetricFilterLister {
	return &metricFilterLister{indexer: indexer}
}

// List lists all MetricFilters in the indexer.
func (s *metricFilterLister) List(selector labels.Selector) (ret []*v1alpha1.MetricFilter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetricFilter))
	})
	return ret, err
}

// MetricFilters returns an object that can list and get MetricFilters.
func (s *metricFilterLister) MetricFilters(namespace string) MetricFilterNamespaceLister {
	return metricFilterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MetricFilterNamespaceLister helps list and get MetricFilters.
type MetricFilterNamespaceLister interface {
	// List lists all MetricFilters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.MetricFilter, err error)
	// Get retrieves the MetricFilter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.MetricFilter, error)
	MetricFilterNamespaceListerExpansion
}

// metricFilterNamespaceLister implements the MetricFilterNamespaceLister
// interface.
type metricFilterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MetricFilters in the indexer for a given namespace.
func (s metricFilterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MetricFilter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetricFilter))
	})
	return ret, err
}

// Get retrieves the MetricFilter from the indexer for a given namespace and name.
func (s metricFilterNamespaceLister) Get(name string) (*v1alpha1.MetricFilter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("metricfilter"), name)
	}
	return obj.(*v1alpha1.MetricFilter), nil
}