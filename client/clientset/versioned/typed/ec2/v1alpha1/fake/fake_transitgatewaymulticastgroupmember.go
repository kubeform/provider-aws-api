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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTransitGatewayMulticastGroupMembers implements TransitGatewayMulticastGroupMemberInterface
type FakeTransitGatewayMulticastGroupMembers struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var transitgatewaymulticastgroupmembersResource = schema.GroupVersionResource{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Resource: "transitgatewaymulticastgroupmembers"}

var transitgatewaymulticastgroupmembersKind = schema.GroupVersionKind{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Kind: "TransitGatewayMulticastGroupMember"}

// Get takes name of the transitGatewayMulticastGroupMember, and returns the corresponding transitGatewayMulticastGroupMember object, and an error if there is any.
func (c *FakeTransitGatewayMulticastGroupMembers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TransitGatewayMulticastGroupMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transitgatewaymulticastgroupmembersResource, c.ns, name), &v1alpha1.TransitGatewayMulticastGroupMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayMulticastGroupMember), err
}

// List takes label and field selectors, and returns the list of TransitGatewayMulticastGroupMembers that match those selectors.
func (c *FakeTransitGatewayMulticastGroupMembers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransitGatewayMulticastGroupMemberList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transitgatewaymulticastgroupmembersResource, transitgatewaymulticastgroupmembersKind, c.ns, opts), &v1alpha1.TransitGatewayMulticastGroupMemberList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransitGatewayMulticastGroupMemberList{ListMeta: obj.(*v1alpha1.TransitGatewayMulticastGroupMemberList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransitGatewayMulticastGroupMemberList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transitGatewayMulticastGroupMembers.
func (c *FakeTransitGatewayMulticastGroupMembers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transitgatewaymulticastgroupmembersResource, c.ns, opts))

}

// Create takes the representation of a transitGatewayMulticastGroupMember and creates it.  Returns the server's representation of the transitGatewayMulticastGroupMember, and an error, if there is any.
func (c *FakeTransitGatewayMulticastGroupMembers) Create(ctx context.Context, transitGatewayMulticastGroupMember *v1alpha1.TransitGatewayMulticastGroupMember, opts v1.CreateOptions) (result *v1alpha1.TransitGatewayMulticastGroupMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transitgatewaymulticastgroupmembersResource, c.ns, transitGatewayMulticastGroupMember), &v1alpha1.TransitGatewayMulticastGroupMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayMulticastGroupMember), err
}

// Update takes the representation of a transitGatewayMulticastGroupMember and updates it. Returns the server's representation of the transitGatewayMulticastGroupMember, and an error, if there is any.
func (c *FakeTransitGatewayMulticastGroupMembers) Update(ctx context.Context, transitGatewayMulticastGroupMember *v1alpha1.TransitGatewayMulticastGroupMember, opts v1.UpdateOptions) (result *v1alpha1.TransitGatewayMulticastGroupMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transitgatewaymulticastgroupmembersResource, c.ns, transitGatewayMulticastGroupMember), &v1alpha1.TransitGatewayMulticastGroupMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayMulticastGroupMember), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransitGatewayMulticastGroupMembers) UpdateStatus(ctx context.Context, transitGatewayMulticastGroupMember *v1alpha1.TransitGatewayMulticastGroupMember, opts v1.UpdateOptions) (*v1alpha1.TransitGatewayMulticastGroupMember, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transitgatewaymulticastgroupmembersResource, "status", c.ns, transitGatewayMulticastGroupMember), &v1alpha1.TransitGatewayMulticastGroupMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayMulticastGroupMember), err
}

// Delete takes name of the transitGatewayMulticastGroupMember and deletes it. Returns an error if one occurs.
func (c *FakeTransitGatewayMulticastGroupMembers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transitgatewaymulticastgroupmembersResource, c.ns, name), &v1alpha1.TransitGatewayMulticastGroupMember{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransitGatewayMulticastGroupMembers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transitgatewaymulticastgroupmembersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransitGatewayMulticastGroupMemberList{})
	return err
}

// Patch applies the patch and returns the patched transitGatewayMulticastGroupMember.
func (c *FakeTransitGatewayMulticastGroupMembers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitGatewayMulticastGroupMember, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transitgatewaymulticastgroupmembersResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransitGatewayMulticastGroupMember{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayMulticastGroupMember), err
}
