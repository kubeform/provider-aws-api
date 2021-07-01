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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"
)

// FakeTransitGatewayVpcAttachmentAccepters implements TransitGatewayVpcAttachmentAccepterInterface
type FakeTransitGatewayVpcAttachmentAccepters struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var transitgatewayvpcattachmentacceptersResource = schema.GroupVersionResource{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Resource: "transitgatewayvpcattachmentaccepters"}

var transitgatewayvpcattachmentacceptersKind = schema.GroupVersionKind{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Kind: "TransitGatewayVpcAttachmentAccepter"}

// Get takes name of the transitGatewayVpcAttachmentAccepter, and returns the corresponding transitGatewayVpcAttachmentAccepter object, and an error if there is any.
func (c *FakeTransitGatewayVpcAttachmentAccepters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transitgatewayvpcattachmentacceptersResource, c.ns, name), &v1alpha1.TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayVpcAttachmentAccepter), err
}

// List takes label and field selectors, and returns the list of TransitGatewayVpcAttachmentAccepters that match those selectors.
func (c *FakeTransitGatewayVpcAttachmentAccepters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransitGatewayVpcAttachmentAccepterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transitgatewayvpcattachmentacceptersResource, transitgatewayvpcattachmentacceptersKind, c.ns, opts), &v1alpha1.TransitGatewayVpcAttachmentAccepterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransitGatewayVpcAttachmentAccepterList{ListMeta: obj.(*v1alpha1.TransitGatewayVpcAttachmentAccepterList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransitGatewayVpcAttachmentAccepterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transitGatewayVpcAttachmentAccepters.
func (c *FakeTransitGatewayVpcAttachmentAccepters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transitgatewayvpcattachmentacceptersResource, c.ns, opts))

}

// Create takes the representation of a transitGatewayVpcAttachmentAccepter and creates it.  Returns the server's representation of the transitGatewayVpcAttachmentAccepter, and an error, if there is any.
func (c *FakeTransitGatewayVpcAttachmentAccepters) Create(ctx context.Context, transitGatewayVpcAttachmentAccepter *v1alpha1.TransitGatewayVpcAttachmentAccepter, opts v1.CreateOptions) (result *v1alpha1.TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transitgatewayvpcattachmentacceptersResource, c.ns, transitGatewayVpcAttachmentAccepter), &v1alpha1.TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayVpcAttachmentAccepter), err
}

// Update takes the representation of a transitGatewayVpcAttachmentAccepter and updates it. Returns the server's representation of the transitGatewayVpcAttachmentAccepter, and an error, if there is any.
func (c *FakeTransitGatewayVpcAttachmentAccepters) Update(ctx context.Context, transitGatewayVpcAttachmentAccepter *v1alpha1.TransitGatewayVpcAttachmentAccepter, opts v1.UpdateOptions) (result *v1alpha1.TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transitgatewayvpcattachmentacceptersResource, c.ns, transitGatewayVpcAttachmentAccepter), &v1alpha1.TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayVpcAttachmentAccepter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransitGatewayVpcAttachmentAccepters) UpdateStatus(ctx context.Context, transitGatewayVpcAttachmentAccepter *v1alpha1.TransitGatewayVpcAttachmentAccepter, opts v1.UpdateOptions) (*v1alpha1.TransitGatewayVpcAttachmentAccepter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transitgatewayvpcattachmentacceptersResource, "status", c.ns, transitGatewayVpcAttachmentAccepter), &v1alpha1.TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayVpcAttachmentAccepter), err
}

// Delete takes name of the transitGatewayVpcAttachmentAccepter and deletes it. Returns an error if one occurs.
func (c *FakeTransitGatewayVpcAttachmentAccepters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transitgatewayvpcattachmentacceptersResource, c.ns, name), &v1alpha1.TransitGatewayVpcAttachmentAccepter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransitGatewayVpcAttachmentAccepters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transitgatewayvpcattachmentacceptersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransitGatewayVpcAttachmentAccepterList{})
	return err
}

// Patch applies the patch and returns the patched transitGatewayVpcAttachmentAccepter.
func (c *FakeTransitGatewayVpcAttachmentAccepters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitGatewayVpcAttachmentAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transitgatewayvpcattachmentacceptersResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransitGatewayVpcAttachmentAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayVpcAttachmentAccepter), err
}
