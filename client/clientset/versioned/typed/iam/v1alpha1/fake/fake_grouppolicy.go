/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-aws-api/apis/iam/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGroupPolicies implements GroupPolicyInterface
type FakeGroupPolicies struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var grouppoliciesResource = schema.GroupVersionResource{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Resource: "grouppolicies"}

var grouppoliciesKind = schema.GroupVersionKind{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Kind: "GroupPolicy"}

// Get takes name of the groupPolicy, and returns the corresponding groupPolicy object, and an error if there is any.
func (c *FakeGroupPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(grouppoliciesResource, c.ns, name), &v1alpha1.GroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicy), err
}

// List takes label and field selectors, and returns the list of GroupPolicies that match those selectors.
func (c *FakeGroupPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GroupPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(grouppoliciesResource, grouppoliciesKind, c.ns, opts), &v1alpha1.GroupPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GroupPolicyList{ListMeta: obj.(*v1alpha1.GroupPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.GroupPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groupPolicies.
func (c *FakeGroupPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(grouppoliciesResource, c.ns, opts))

}

// Create takes the representation of a groupPolicy and creates it.  Returns the server's representation of the groupPolicy, and an error, if there is any.
func (c *FakeGroupPolicies) Create(ctx context.Context, groupPolicy *v1alpha1.GroupPolicy, opts v1.CreateOptions) (result *v1alpha1.GroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(grouppoliciesResource, c.ns, groupPolicy), &v1alpha1.GroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicy), err
}

// Update takes the representation of a groupPolicy and updates it. Returns the server's representation of the groupPolicy, and an error, if there is any.
func (c *FakeGroupPolicies) Update(ctx context.Context, groupPolicy *v1alpha1.GroupPolicy, opts v1.UpdateOptions) (result *v1alpha1.GroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(grouppoliciesResource, c.ns, groupPolicy), &v1alpha1.GroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGroupPolicies) UpdateStatus(ctx context.Context, groupPolicy *v1alpha1.GroupPolicy, opts v1.UpdateOptions) (*v1alpha1.GroupPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(grouppoliciesResource, "status", c.ns, groupPolicy), &v1alpha1.GroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicy), err
}

// Delete takes name of the groupPolicy and deletes it. Returns an error if one occurs.
func (c *FakeGroupPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(grouppoliciesResource, c.ns, name), &v1alpha1.GroupPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroupPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(grouppoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GroupPolicyList{})
	return err
}

// Patch applies the patch and returns the patched groupPolicy.
func (c *FakeGroupPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GroupPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(grouppoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.GroupPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupPolicy), err
}
