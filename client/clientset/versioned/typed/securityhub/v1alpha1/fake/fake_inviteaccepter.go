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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/securityhub/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInviteAccepters implements InviteAccepterInterface
type FakeInviteAccepters struct {
	Fake *FakeSecurityhubV1alpha1
	ns   string
}

var inviteacceptersResource = schema.GroupVersionResource{Group: "securityhub.aws.kubeform.com", Version: "v1alpha1", Resource: "inviteaccepters"}

var inviteacceptersKind = schema.GroupVersionKind{Group: "securityhub.aws.kubeform.com", Version: "v1alpha1", Kind: "InviteAccepter"}

// Get takes name of the inviteAccepter, and returns the corresponding inviteAccepter object, and an error if there is any.
func (c *FakeInviteAccepters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InviteAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(inviteacceptersResource, c.ns, name), &v1alpha1.InviteAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InviteAccepter), err
}

// List takes label and field selectors, and returns the list of InviteAccepters that match those selectors.
func (c *FakeInviteAccepters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InviteAccepterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(inviteacceptersResource, inviteacceptersKind, c.ns, opts), &v1alpha1.InviteAccepterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InviteAccepterList{ListMeta: obj.(*v1alpha1.InviteAccepterList).ListMeta}
	for _, item := range obj.(*v1alpha1.InviteAccepterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inviteAccepters.
func (c *FakeInviteAccepters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(inviteacceptersResource, c.ns, opts))

}

// Create takes the representation of a inviteAccepter and creates it.  Returns the server's representation of the inviteAccepter, and an error, if there is any.
func (c *FakeInviteAccepters) Create(ctx context.Context, inviteAccepter *v1alpha1.InviteAccepter, opts v1.CreateOptions) (result *v1alpha1.InviteAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(inviteacceptersResource, c.ns, inviteAccepter), &v1alpha1.InviteAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InviteAccepter), err
}

// Update takes the representation of a inviteAccepter and updates it. Returns the server's representation of the inviteAccepter, and an error, if there is any.
func (c *FakeInviteAccepters) Update(ctx context.Context, inviteAccepter *v1alpha1.InviteAccepter, opts v1.UpdateOptions) (result *v1alpha1.InviteAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(inviteacceptersResource, c.ns, inviteAccepter), &v1alpha1.InviteAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InviteAccepter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInviteAccepters) UpdateStatus(ctx context.Context, inviteAccepter *v1alpha1.InviteAccepter, opts v1.UpdateOptions) (*v1alpha1.InviteAccepter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(inviteacceptersResource, "status", c.ns, inviteAccepter), &v1alpha1.InviteAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InviteAccepter), err
}

// Delete takes name of the inviteAccepter and deletes it. Returns an error if one occurs.
func (c *FakeInviteAccepters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(inviteacceptersResource, c.ns, name), &v1alpha1.InviteAccepter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInviteAccepters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(inviteacceptersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InviteAccepterList{})
	return err
}

// Patch applies the patch and returns the patched inviteAccepter.
func (c *FakeInviteAccepters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InviteAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(inviteacceptersResource, c.ns, name, pt, data, subresources...), &v1alpha1.InviteAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InviteAccepter), err
}
