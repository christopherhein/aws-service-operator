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
	v1alpha1 "awsoperator.io/pkg/apis/kinesisanalytics/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApplicationOutputLister helps list ApplicationOutputs.
type ApplicationOutputLister interface {
	// List lists all ApplicationOutputs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationOutput, err error)
	// ApplicationOutputs returns an object that can list and get ApplicationOutputs.
	ApplicationOutputs(namespace string) ApplicationOutputNamespaceLister
	ApplicationOutputListerExpansion
}

// applicationOutputLister implements the ApplicationOutputLister interface.
type applicationOutputLister struct {
	indexer cache.Indexer
}

// NewApplicationOutputLister returns a new ApplicationOutputLister.
func NewApplicationOutputLister(indexer cache.Indexer) ApplicationOutputLister {
	return &applicationOutputLister{indexer: indexer}
}

// List lists all ApplicationOutputs in the indexer.
func (s *applicationOutputLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationOutput, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationOutput))
	})
	return ret, err
}

// ApplicationOutputs returns an object that can list and get ApplicationOutputs.
func (s *applicationOutputLister) ApplicationOutputs(namespace string) ApplicationOutputNamespaceLister {
	return applicationOutputNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationOutputNamespaceLister helps list and get ApplicationOutputs.
type ApplicationOutputNamespaceLister interface {
	// List lists all ApplicationOutputs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationOutput, err error)
	// Get retrieves the ApplicationOutput from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApplicationOutput, error)
	ApplicationOutputNamespaceListerExpansion
}

// applicationOutputNamespaceLister implements the ApplicationOutputNamespaceLister
// interface.
type applicationOutputNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationOutputs in the indexer for a given namespace.
func (s applicationOutputNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationOutput, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationOutput))
	})
	return ret, err
}

// Get retrieves the ApplicationOutput from the indexer for a given namespace and name.
func (s applicationOutputNamespaceLister) Get(name string) (*v1alpha1.ApplicationOutput, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationoutput"), name)
	}
	return obj.(*v1alpha1.ApplicationOutput), nil
}