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
	v1alpha1 "awsoperator.io/pkg/apis/greengrass/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LoggerDefinitionLister helps list LoggerDefinitions.
type LoggerDefinitionLister interface {
	// List lists all LoggerDefinitions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.LoggerDefinition, err error)
	// LoggerDefinitions returns an object that can list and get LoggerDefinitions.
	LoggerDefinitions(namespace string) LoggerDefinitionNamespaceLister
	LoggerDefinitionListerExpansion
}

// loggerDefinitionLister implements the LoggerDefinitionLister interface.
type loggerDefinitionLister struct {
	indexer cache.Indexer
}

// NewLoggerDefinitionLister returns a new LoggerDefinitionLister.
func NewLoggerDefinitionLister(indexer cache.Indexer) LoggerDefinitionLister {
	return &loggerDefinitionLister{indexer: indexer}
}

// List lists all LoggerDefinitions in the indexer.
func (s *loggerDefinitionLister) List(selector labels.Selector) (ret []*v1alpha1.LoggerDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LoggerDefinition))
	})
	return ret, err
}

// LoggerDefinitions returns an object that can list and get LoggerDefinitions.
func (s *loggerDefinitionLister) LoggerDefinitions(namespace string) LoggerDefinitionNamespaceLister {
	return loggerDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LoggerDefinitionNamespaceLister helps list and get LoggerDefinitions.
type LoggerDefinitionNamespaceLister interface {
	// List lists all LoggerDefinitions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.LoggerDefinition, err error)
	// Get retrieves the LoggerDefinition from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.LoggerDefinition, error)
	LoggerDefinitionNamespaceListerExpansion
}

// loggerDefinitionNamespaceLister implements the LoggerDefinitionNamespaceLister
// interface.
type loggerDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LoggerDefinitions in the indexer for a given namespace.
func (s loggerDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LoggerDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LoggerDefinition))
	})
	return ret, err
}

// Get retrieves the LoggerDefinition from the indexer for a given namespace and name.
func (s loggerDefinitionNamespaceLister) Get(name string) (*v1alpha1.LoggerDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("loggerdefinition"), name)
	}
	return obj.(*v1alpha1.LoggerDefinition), nil
}