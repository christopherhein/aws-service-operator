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
	v1alpha1 "awsoperator.io/pkg/apis/dms/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReplicationSubnetGroupLister helps list ReplicationSubnetGroups.
type ReplicationSubnetGroupLister interface {
	// List lists all ReplicationSubnetGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ReplicationSubnetGroup, err error)
	// ReplicationSubnetGroups returns an object that can list and get ReplicationSubnetGroups.
	ReplicationSubnetGroups(namespace string) ReplicationSubnetGroupNamespaceLister
	ReplicationSubnetGroupListerExpansion
}

// replicationSubnetGroupLister implements the ReplicationSubnetGroupLister interface.
type replicationSubnetGroupLister struct {
	indexer cache.Indexer
}

// NewReplicationSubnetGroupLister returns a new ReplicationSubnetGroupLister.
func NewReplicationSubnetGroupLister(indexer cache.Indexer) ReplicationSubnetGroupLister {
	return &replicationSubnetGroupLister{indexer: indexer}
}

// List lists all ReplicationSubnetGroups in the indexer.
func (s *replicationSubnetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ReplicationSubnetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReplicationSubnetGroup))
	})
	return ret, err
}

// ReplicationSubnetGroups returns an object that can list and get ReplicationSubnetGroups.
func (s *replicationSubnetGroupLister) ReplicationSubnetGroups(namespace string) ReplicationSubnetGroupNamespaceLister {
	return replicationSubnetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReplicationSubnetGroupNamespaceLister helps list and get ReplicationSubnetGroups.
type ReplicationSubnetGroupNamespaceLister interface {
	// List lists all ReplicationSubnetGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ReplicationSubnetGroup, err error)
	// Get retrieves the ReplicationSubnetGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ReplicationSubnetGroup, error)
	ReplicationSubnetGroupNamespaceListerExpansion
}

// replicationSubnetGroupNamespaceLister implements the ReplicationSubnetGroupNamespaceLister
// interface.
type replicationSubnetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReplicationSubnetGroups in the indexer for a given namespace.
func (s replicationSubnetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ReplicationSubnetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReplicationSubnetGroup))
	})
	return ret, err
}

// Get retrieves the ReplicationSubnetGroup from the indexer for a given namespace and name.
func (s replicationSubnetGroupNamespaceLister) Get(name string) (*v1alpha1.ReplicationSubnetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("replicationsubnetgroup"), name)
	}
	return obj.(*v1alpha1.ReplicationSubnetGroup), nil
}