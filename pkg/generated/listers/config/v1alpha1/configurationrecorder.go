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
	v1alpha1 "awsoperator.io/pkg/apis/config/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConfigurationRecorderLister helps list ConfigurationRecorders.
type ConfigurationRecorderLister interface {
	// List lists all ConfigurationRecorders in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigurationRecorder, err error)
	// ConfigurationRecorders returns an object that can list and get ConfigurationRecorders.
	ConfigurationRecorders(namespace string) ConfigurationRecorderNamespaceLister
	ConfigurationRecorderListerExpansion
}

// configurationRecorderLister implements the ConfigurationRecorderLister interface.
type configurationRecorderLister struct {
	indexer cache.Indexer
}

// NewConfigurationRecorderLister returns a new ConfigurationRecorderLister.
func NewConfigurationRecorderLister(indexer cache.Indexer) ConfigurationRecorderLister {
	return &configurationRecorderLister{indexer: indexer}
}

// List lists all ConfigurationRecorders in the indexer.
func (s *configurationRecorderLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigurationRecorder, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigurationRecorder))
	})
	return ret, err
}

// ConfigurationRecorders returns an object that can list and get ConfigurationRecorders.
func (s *configurationRecorderLister) ConfigurationRecorders(namespace string) ConfigurationRecorderNamespaceLister {
	return configurationRecorderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigurationRecorderNamespaceLister helps list and get ConfigurationRecorders.
type ConfigurationRecorderNamespaceLister interface {
	// List lists all ConfigurationRecorders in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigurationRecorder, err error)
	// Get retrieves the ConfigurationRecorder from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ConfigurationRecorder, error)
	ConfigurationRecorderNamespaceListerExpansion
}

// configurationRecorderNamespaceLister implements the ConfigurationRecorderNamespaceLister
// interface.
type configurationRecorderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigurationRecorders in the indexer for a given namespace.
func (s configurationRecorderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigurationRecorder, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigurationRecorder))
	})
	return ret, err
}

// Get retrieves the ConfigurationRecorder from the indexer for a given namespace and name.
func (s configurationRecorderNamespaceLister) Get(name string) (*v1alpha1.ConfigurationRecorder, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("configurationrecorder"), name)
	}
	return obj.(*v1alpha1.ConfigurationRecorder), nil
}