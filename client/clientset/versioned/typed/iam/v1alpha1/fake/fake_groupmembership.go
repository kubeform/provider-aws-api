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

// FakeGroupMemberships implements GroupMembershipInterface
type FakeGroupMemberships struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var groupmembershipsResource = schema.GroupVersionResource{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Resource: "groupmemberships"}

var groupmembershipsKind = schema.GroupVersionKind{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Kind: "GroupMembership"}

// Get takes name of the groupMembership, and returns the corresponding groupMembership object, and an error if there is any.
func (c *FakeGroupMemberships) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(groupmembershipsResource, c.ns, name), &v1alpha1.GroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMembership), err
}

// List takes label and field selectors, and returns the list of GroupMemberships that match those selectors.
func (c *FakeGroupMemberships) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GroupMembershipList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(groupmembershipsResource, groupmembershipsKind, c.ns, opts), &v1alpha1.GroupMembershipList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GroupMembershipList{ListMeta: obj.(*v1alpha1.GroupMembershipList).ListMeta}
	for _, item := range obj.(*v1alpha1.GroupMembershipList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groupMemberships.
func (c *FakeGroupMemberships) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(groupmembershipsResource, c.ns, opts))

}

// Create takes the representation of a groupMembership and creates it.  Returns the server's representation of the groupMembership, and an error, if there is any.
func (c *FakeGroupMemberships) Create(ctx context.Context, groupMembership *v1alpha1.GroupMembership, opts v1.CreateOptions) (result *v1alpha1.GroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(groupmembershipsResource, c.ns, groupMembership), &v1alpha1.GroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMembership), err
}

// Update takes the representation of a groupMembership and updates it. Returns the server's representation of the groupMembership, and an error, if there is any.
func (c *FakeGroupMemberships) Update(ctx context.Context, groupMembership *v1alpha1.GroupMembership, opts v1.UpdateOptions) (result *v1alpha1.GroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(groupmembershipsResource, c.ns, groupMembership), &v1alpha1.GroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMembership), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGroupMemberships) UpdateStatus(ctx context.Context, groupMembership *v1alpha1.GroupMembership, opts v1.UpdateOptions) (*v1alpha1.GroupMembership, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(groupmembershipsResource, "status", c.ns, groupMembership), &v1alpha1.GroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMembership), err
}

// Delete takes name of the groupMembership and deletes it. Returns an error if one occurs.
func (c *FakeGroupMemberships) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(groupmembershipsResource, c.ns, name), &v1alpha1.GroupMembership{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroupMemberships) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(groupmembershipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GroupMembershipList{})
	return err
}

// Patch applies the patch and returns the patched groupMembership.
func (c *FakeGroupMemberships) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GroupMembership, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(groupmembershipsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GroupMembership{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GroupMembership), err
}
