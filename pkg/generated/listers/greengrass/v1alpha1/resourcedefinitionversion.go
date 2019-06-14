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

// ResourceDefinitionVersionLister helps list ResourceDefinitionVersions.
type ResourceDefinitionVersionLister interface {
	// List lists all ResourceDefinitionVersions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceDefinitionVersion, err error)
	// ResourceDefinitionVersions returns an object that can list and get ResourceDefinitionVersions.
	ResourceDefinitionVersions(namespace string) ResourceDefinitionVersionNamespaceLister
	ResourceDefinitionVersionListerExpansion
}

// resourceDefinitionVersionLister implements the ResourceDefinitionVersionLister interface.
type resourceDefinitionVersionLister struct {
	indexer cache.Indexer
}

// NewResourceDefinitionVersionLister returns a new ResourceDefinitionVersionLister.
func NewResourceDefinitionVersionLister(indexer cache.Indexer) ResourceDefinitionVersionLister {
	return &resourceDefinitionVersionLister{indexer: indexer}
}

// List lists all ResourceDefinitionVersions in the indexer.
func (s *resourceDefinitionVersionLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceDefinitionVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceDefinitionVersion))
	})
	return ret, err
}

// ResourceDefinitionVersions returns an object that can list and get ResourceDefinitionVersions.
func (s *resourceDefinitionVersionLister) ResourceDefinitionVersions(namespace string) ResourceDefinitionVersionNamespaceLister {
	return resourceDefinitionVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceDefinitionVersionNamespaceLister helps list and get ResourceDefinitionVersions.
type ResourceDefinitionVersionNamespaceLister interface {
	// List lists all ResourceDefinitionVersions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceDefinitionVersion, err error)
	// Get retrieves the ResourceDefinitionVersion from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ResourceDefinitionVersion, error)
	ResourceDefinitionVersionNamespaceListerExpansion
}

// resourceDefinitionVersionNamespaceLister implements the ResourceDefinitionVersionNamespaceLister
// interface.
type resourceDefinitionVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceDefinitionVersions in the indexer for a given namespace.
func (s resourceDefinitionVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceDefinitionVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceDefinitionVersion))
	})
	return ret, err
}

// Get retrieves the ResourceDefinitionVersion from the indexer for a given namespace and name.
func (s resourceDefinitionVersionNamespaceLister) Get(name string) (*v1alpha1.ResourceDefinitionVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourcedefinitionversion"), name)
	}
	return obj.(*v1alpha1.ResourceDefinitionVersion), nil
}