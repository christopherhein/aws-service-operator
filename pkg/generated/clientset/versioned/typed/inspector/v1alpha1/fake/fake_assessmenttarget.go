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
	v1alpha1 "awsoperator.io/pkg/apis/inspector/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAssessmentTargets implements AssessmentTargetInterface
type FakeAssessmentTargets struct {
	Fake *FakeInspectorV1alpha1
	ns   string
}

var assessmenttargetsResource = schema.GroupVersionResource{Group: "inspector.awsoperator.io", Version: "v1alpha1", Resource: "assessmenttargets"}

var assessmenttargetsKind = schema.GroupVersionKind{Group: "inspector.awsoperator.io", Version: "v1alpha1", Kind: "AssessmentTarget"}

// Get takes name of the assessmentTarget, and returns the corresponding assessmentTarget object, and an error if there is any.
func (c *FakeAssessmentTargets) Get(name string, options v1.GetOptions) (result *v1alpha1.AssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(assessmenttargetsResource, c.ns, name), &v1alpha1.AssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssessmentTarget), err
}

// List takes label and field selectors, and returns the list of AssessmentTargets that match those selectors.
func (c *FakeAssessmentTargets) List(opts v1.ListOptions) (result *v1alpha1.AssessmentTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(assessmenttargetsResource, assessmenttargetsKind, c.ns, opts), &v1alpha1.AssessmentTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AssessmentTargetList{ListMeta: obj.(*v1alpha1.AssessmentTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.AssessmentTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested assessmentTargets.
func (c *FakeAssessmentTargets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(assessmenttargetsResource, c.ns, opts))

}

// Create takes the representation of a assessmentTarget and creates it.  Returns the server's representation of the assessmentTarget, and an error, if there is any.
func (c *FakeAssessmentTargets) Create(assessmentTarget *v1alpha1.AssessmentTarget) (result *v1alpha1.AssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(assessmenttargetsResource, c.ns, assessmentTarget), &v1alpha1.AssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssessmentTarget), err
}

// Update takes the representation of a assessmentTarget and updates it. Returns the server's representation of the assessmentTarget, and an error, if there is any.
func (c *FakeAssessmentTargets) Update(assessmentTarget *v1alpha1.AssessmentTarget) (result *v1alpha1.AssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(assessmenttargetsResource, c.ns, assessmentTarget), &v1alpha1.AssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssessmentTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAssessmentTargets) UpdateStatus(assessmentTarget *v1alpha1.AssessmentTarget) (*v1alpha1.AssessmentTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(assessmenttargetsResource, "status", c.ns, assessmentTarget), &v1alpha1.AssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssessmentTarget), err
}

// Delete takes name of the assessmentTarget and deletes it. Returns an error if one occurs.
func (c *FakeAssessmentTargets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(assessmenttargetsResource, c.ns, name), &v1alpha1.AssessmentTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAssessmentTargets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(assessmenttargetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AssessmentTargetList{})
	return err
}

// Patch applies the patch and returns the patched assessmentTarget.
func (c *FakeAssessmentTargets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AssessmentTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(assessmenttargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AssessmentTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AssessmentTarget), err
}