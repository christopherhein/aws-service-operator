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
	v1alpha1 "awsoperator.io/pkg/apis/emr/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SecurityConfigurationLister helps list SecurityConfigurations.
type SecurityConfigurationLister interface {
	// List lists all SecurityConfigurations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityConfiguration, err error)
	// SecurityConfigurations returns an object that can list and get SecurityConfigurations.
	SecurityConfigurations(namespace string) SecurityConfigurationNamespaceLister
	SecurityConfigurationListerExpansion
}

// securityConfigurationLister implements the SecurityConfigurationLister interface.
type securityConfigurationLister struct {
	indexer cache.Indexer
}

// NewSecurityConfigurationLister returns a new SecurityConfigurationLister.
func NewSecurityConfigurationLister(indexer cache.Indexer) SecurityConfigurationLister {
	return &securityConfigurationLister{indexer: indexer}
}

// List lists all SecurityConfigurations in the indexer.
func (s *securityConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityConfiguration))
	})
	return ret, err
}

// SecurityConfigurations returns an object that can list and get SecurityConfigurations.
func (s *securityConfigurationLister) SecurityConfigurations(namespace string) SecurityConfigurationNamespaceLister {
	return securityConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecurityConfigurationNamespaceLister helps list and get SecurityConfigurations.
type SecurityConfigurationNamespaceLister interface {
	// List lists all SecurityConfigurations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SecurityConfiguration, err error)
	// Get retrieves the SecurityConfiguration from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SecurityConfiguration, error)
	SecurityConfigurationNamespaceListerExpansion
}

// securityConfigurationNamespaceLister implements the SecurityConfigurationNamespaceLister
// interface.
type securityConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecurityConfigurations in the indexer for a given namespace.
func (s securityConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecurityConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecurityConfiguration))
	})
	return ret, err
}

// Get retrieves the SecurityConfiguration from the indexer for a given namespace and name.
func (s securityConfigurationNamespaceLister) Get(name string) (*v1alpha1.SecurityConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("securityconfiguration"), name)
	}
	return obj.(*v1alpha1.SecurityConfiguration), nil
}