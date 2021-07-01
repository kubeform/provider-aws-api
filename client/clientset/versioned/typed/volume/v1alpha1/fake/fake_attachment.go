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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/volume/v1alpha1"
)

// FakeAttachments implements AttachmentInterface
type FakeAttachments struct {
	Fake *FakeVolumeV1alpha1
	ns   string
}

var attachmentsResource = schema.GroupVersionResource{Group: "volume.aws.kubeform.com", Version: "v1alpha1", Resource: "attachments"}

var attachmentsKind = schema.GroupVersionKind{Group: "volume.aws.kubeform.com", Version: "v1alpha1", Kind: "Attachment"}

// Get takes name of the attachment, and returns the corresponding attachment object, and an error if there is any.
func (c *FakeAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Attachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(attachmentsResource, c.ns, name), &v1alpha1.Attachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attachment), err
}

// List takes label and field selectors, and returns the list of Attachments that match those selectors.
func (c *FakeAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(attachmentsResource, attachmentsKind, c.ns, opts), &v1alpha1.AttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AttachmentList{ListMeta: obj.(*v1alpha1.AttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.AttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested attachments.
func (c *FakeAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(attachmentsResource, c.ns, opts))

}

// Create takes the representation of a attachment and creates it.  Returns the server's representation of the attachment, and an error, if there is any.
func (c *FakeAttachments) Create(ctx context.Context, attachment *v1alpha1.Attachment, opts v1.CreateOptions) (result *v1alpha1.Attachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(attachmentsResource, c.ns, attachment), &v1alpha1.Attachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attachment), err
}

// Update takes the representation of a attachment and updates it. Returns the server's representation of the attachment, and an error, if there is any.
func (c *FakeAttachments) Update(ctx context.Context, attachment *v1alpha1.Attachment, opts v1.UpdateOptions) (result *v1alpha1.Attachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(attachmentsResource, c.ns, attachment), &v1alpha1.Attachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAttachments) UpdateStatus(ctx context.Context, attachment *v1alpha1.Attachment, opts v1.UpdateOptions) (*v1alpha1.Attachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(attachmentsResource, "status", c.ns, attachment), &v1alpha1.Attachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attachment), err
}

// Delete takes name of the attachment and deletes it. Returns an error if one occurs.
func (c *FakeAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(attachmentsResource, c.ns, name), &v1alpha1.Attachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(attachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AttachmentList{})
	return err
}

// Patch applies the patch and returns the patched attachment.
func (c *FakeAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Attachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(attachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Attachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Attachment), err
}
