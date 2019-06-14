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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "awsoperator.io/pkg/apis/dms/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicationTasks implements ReplicationTaskInterface
type FakeReplicationTasks struct {
	Fake *FakeDmsV1alpha1
	ns   string
}

var replicationtasksResource = schema.GroupVersionResource{Group: "dms.awsoperator.io", Version: "v1alpha1", Resource: "replicationtasks"}

var replicationtasksKind = schema.GroupVersionKind{Group: "dms.awsoperator.io", Version: "v1alpha1", Kind: "ReplicationTask"}

// Get takes name of the replicationTask, and returns the corresponding replicationTask object, and an error if there is any.
func (c *FakeReplicationTasks) Get(name string, options v1.GetOptions) (result *v1alpha1.ReplicationTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicationtasksResource, c.ns, name), &v1alpha1.ReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationTask), err
}

// List takes label and field selectors, and returns the list of ReplicationTasks that match those selectors.
func (c *FakeReplicationTasks) List(opts v1.ListOptions) (result *v1alpha1.ReplicationTaskList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicationtasksResource, replicationtasksKind, c.ns, opts), &v1alpha1.ReplicationTaskList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReplicationTaskList{ListMeta: obj.(*v1alpha1.ReplicationTaskList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReplicationTaskList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicationTasks.
func (c *FakeReplicationTasks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicationtasksResource, c.ns, opts))

}

// Create takes the representation of a replicationTask and creates it.  Returns the server's representation of the replicationTask, and an error, if there is any.
func (c *FakeReplicationTasks) Create(replicationTask *v1alpha1.ReplicationTask) (result *v1alpha1.ReplicationTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicationtasksResource, c.ns, replicationTask), &v1alpha1.ReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationTask), err
}

// Update takes the representation of a replicationTask and updates it. Returns the server's representation of the replicationTask, and an error, if there is any.
func (c *FakeReplicationTasks) Update(replicationTask *v1alpha1.ReplicationTask) (result *v1alpha1.ReplicationTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicationtasksResource, c.ns, replicationTask), &v1alpha1.ReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationTask), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicationTasks) UpdateStatus(replicationTask *v1alpha1.ReplicationTask) (*v1alpha1.ReplicationTask, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicationtasksResource, "status", c.ns, replicationTask), &v1alpha1.ReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationTask), err
}

// Delete takes name of the replicationTask and deletes it. Returns an error if one occurs.
func (c *FakeReplicationTasks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicationtasksResource, c.ns, name), &v1alpha1.ReplicationTask{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicationTasks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicationtasksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReplicationTaskList{})
	return err
}

// Patch applies the patch and returns the patched replicationTask.
func (c *FakeReplicationTasks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReplicationTask, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicationtasksResource, c.ns, name, pt, data, subresources...), &v1alpha1.ReplicationTask{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationTask), err
}