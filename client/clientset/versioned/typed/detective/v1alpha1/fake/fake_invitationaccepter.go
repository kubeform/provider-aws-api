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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/detective/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeInvitationAccepters implements InvitationAccepterInterface
type FakeInvitationAccepters struct {
	Fake *FakeDetectiveV1alpha1
	ns   string
}

var invitationacceptersResource = schema.GroupVersionResource{Group: "detective.aws.kubeform.com", Version: "v1alpha1", Resource: "invitationaccepters"}

var invitationacceptersKind = schema.GroupVersionKind{Group: "detective.aws.kubeform.com", Version: "v1alpha1", Kind: "InvitationAccepter"}

// Get takes name of the invitationAccepter, and returns the corresponding invitationAccepter object, and an error if there is any.
func (c *FakeInvitationAccepters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InvitationAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(invitationacceptersResource, c.ns, name), &v1alpha1.InvitationAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InvitationAccepter), err
}

// List takes label and field selectors, and returns the list of InvitationAccepters that match those selectors.
func (c *FakeInvitationAccepters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InvitationAccepterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(invitationacceptersResource, invitationacceptersKind, c.ns, opts), &v1alpha1.InvitationAccepterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InvitationAccepterList{ListMeta: obj.(*v1alpha1.InvitationAccepterList).ListMeta}
	for _, item := range obj.(*v1alpha1.InvitationAccepterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested invitationAccepters.
func (c *FakeInvitationAccepters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(invitationacceptersResource, c.ns, opts))

}

// Create takes the representation of a invitationAccepter and creates it.  Returns the server's representation of the invitationAccepter, and an error, if there is any.
func (c *FakeInvitationAccepters) Create(ctx context.Context, invitationAccepter *v1alpha1.InvitationAccepter, opts v1.CreateOptions) (result *v1alpha1.InvitationAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(invitationacceptersResource, c.ns, invitationAccepter), &v1alpha1.InvitationAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InvitationAccepter), err
}

// Update takes the representation of a invitationAccepter and updates it. Returns the server's representation of the invitationAccepter, and an error, if there is any.
func (c *FakeInvitationAccepters) Update(ctx context.Context, invitationAccepter *v1alpha1.InvitationAccepter, opts v1.UpdateOptions) (result *v1alpha1.InvitationAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(invitationacceptersResource, c.ns, invitationAccepter), &v1alpha1.InvitationAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InvitationAccepter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInvitationAccepters) UpdateStatus(ctx context.Context, invitationAccepter *v1alpha1.InvitationAccepter, opts v1.UpdateOptions) (*v1alpha1.InvitationAccepter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(invitationacceptersResource, "status", c.ns, invitationAccepter), &v1alpha1.InvitationAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InvitationAccepter), err
}

// Delete takes name of the invitationAccepter and deletes it. Returns an error if one occurs.
func (c *FakeInvitationAccepters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(invitationacceptersResource, c.ns, name), &v1alpha1.InvitationAccepter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInvitationAccepters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(invitationacceptersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InvitationAccepterList{})
	return err
}

// Patch applies the patch and returns the patched invitationAccepter.
func (c *FakeInvitationAccepters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InvitationAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(invitationacceptersResource, c.ns, name, pt, data, subresources...), &v1alpha1.InvitationAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InvitationAccepter), err
}
