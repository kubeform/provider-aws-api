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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/internet/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGatewayAttachments implements GatewayAttachmentInterface
type FakeGatewayAttachments struct {
	Fake *FakeInternetV1alpha1
	ns   string
}

var gatewayattachmentsResource = schema.GroupVersionResource{Group: "internet.aws.kubeform.com", Version: "v1alpha1", Resource: "gatewayattachments"}

var gatewayattachmentsKind = schema.GroupVersionKind{Group: "internet.aws.kubeform.com", Version: "v1alpha1", Kind: "GatewayAttachment"}

// Get takes name of the gatewayAttachment, and returns the corresponding gatewayAttachment object, and an error if there is any.
func (c *FakeGatewayAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gatewayattachmentsResource, c.ns, name), &v1alpha1.GatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAttachment), err
}

// List takes label and field selectors, and returns the list of GatewayAttachments that match those selectors.
func (c *FakeGatewayAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GatewayAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gatewayattachmentsResource, gatewayattachmentsKind, c.ns, opts), &v1alpha1.GatewayAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GatewayAttachmentList{ListMeta: obj.(*v1alpha1.GatewayAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.GatewayAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gatewayAttachments.
func (c *FakeGatewayAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gatewayattachmentsResource, c.ns, opts))

}

// Create takes the representation of a gatewayAttachment and creates it.  Returns the server's representation of the gatewayAttachment, and an error, if there is any.
func (c *FakeGatewayAttachments) Create(ctx context.Context, gatewayAttachment *v1alpha1.GatewayAttachment, opts v1.CreateOptions) (result *v1alpha1.GatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gatewayattachmentsResource, c.ns, gatewayAttachment), &v1alpha1.GatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAttachment), err
}

// Update takes the representation of a gatewayAttachment and updates it. Returns the server's representation of the gatewayAttachment, and an error, if there is any.
func (c *FakeGatewayAttachments) Update(ctx context.Context, gatewayAttachment *v1alpha1.GatewayAttachment, opts v1.UpdateOptions) (result *v1alpha1.GatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gatewayattachmentsResource, c.ns, gatewayAttachment), &v1alpha1.GatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGatewayAttachments) UpdateStatus(ctx context.Context, gatewayAttachment *v1alpha1.GatewayAttachment, opts v1.UpdateOptions) (*v1alpha1.GatewayAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gatewayattachmentsResource, "status", c.ns, gatewayAttachment), &v1alpha1.GatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAttachment), err
}

// Delete takes name of the gatewayAttachment and deletes it. Returns an error if one occurs.
func (c *FakeGatewayAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gatewayattachmentsResource, c.ns, name), &v1alpha1.GatewayAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGatewayAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gatewayattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GatewayAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched gatewayAttachment.
func (c *FakeGatewayAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GatewayAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gatewayattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GatewayAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayAttachment), err
}
