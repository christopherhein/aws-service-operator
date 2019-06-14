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
	v1alpha1 "awsoperator.io/pkg/apis/mediastore/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ContainerLister helps list Containers.
type ContainerLister interface {
	// List lists all Containers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Container, err error)
	// Containers returns an object that can list and get Containers.
	Containers(namespace string) ContainerNamespaceLister
	ContainerListerExpansion
}

// containerLister implements the ContainerLister interface.
type containerLister struct {
	indexer cache.Indexer
}

// NewContainerLister returns a new ContainerLister.
func NewContainerLister(indexer cache.Indexer) ContainerLister {
	return &containerLister{indexer: indexer}
}

// List lists all Containers in the indexer.
func (s *containerLister) List(selector labels.Selector) (ret []*v1alpha1.Container, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Container))
	})
	return ret, err
}

// Containers returns an object that can list and get Containers.
func (s *containerLister) Containers(namespace string) ContainerNamespaceLister {
	return containerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ContainerNamespaceLister helps list and get Containers.
type ContainerNamespaceLister interface {
	// List lists all Containers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Container, err error)
	// Get retrieves the Container from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Container, error)
	ContainerNamespaceListerExpansion
}

// containerNamespaceLister implements the ContainerNamespaceLister
// interface.
type containerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Containers in the indexer for a given namespace.
func (s containerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Container, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Container))
	})
	return ret, err
}

// Get retrieves the Container from the indexer for a given namespace and name.
func (s containerNamespaceLister) Get(name string) (*v1alpha1.Container, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("container"), name)
	}
	return obj.(*v1alpha1.Container), nil
}