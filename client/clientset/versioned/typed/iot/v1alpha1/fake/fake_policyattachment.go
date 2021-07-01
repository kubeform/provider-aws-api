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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/iot/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePolicyAttachments implements PolicyAttachmentInterface
type FakePolicyAttachments struct {
	Fake *FakeIotV1alpha1
	ns   string
}

var policyattachmentsResource = schema.GroupVersionResource{Group: "iot.aws.kubeform.com", Version: "v1alpha1", Resource: "policyattachments"}

var policyattachmentsKind = schema.GroupVersionKind{Group: "iot.aws.kubeform.com", Version: "v1alpha1", Kind: "PolicyAttachment"}

// Get takes name of the policyAttachment, and returns the corresponding policyAttachment object, and an error if there is any.
func (c *FakePolicyAttachments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(policyattachmentsResource, c.ns, name), &v1alpha1.PolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAttachment), err
}

// List takes label and field selectors, and returns the list of PolicyAttachments that match those selectors.
func (c *FakePolicyAttachments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PolicyAttachmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(policyattachmentsResource, policyattachmentsKind, c.ns, opts), &v1alpha1.PolicyAttachmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PolicyAttachmentList{ListMeta: obj.(*v1alpha1.PolicyAttachmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.PolicyAttachmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested policyAttachments.
func (c *FakePolicyAttachments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(policyattachmentsResource, c.ns, opts))

}

// Create takes the representation of a policyAttachment and creates it.  Returns the server's representation of the policyAttachment, and an error, if there is any.
func (c *FakePolicyAttachments) Create(ctx context.Context, policyAttachment *v1alpha1.PolicyAttachment, opts v1.CreateOptions) (result *v1alpha1.PolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(policyattachmentsResource, c.ns, policyAttachment), &v1alpha1.PolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAttachment), err
}

// Update takes the representation of a policyAttachment and updates it. Returns the server's representation of the policyAttachment, and an error, if there is any.
func (c *FakePolicyAttachments) Update(ctx context.Context, policyAttachment *v1alpha1.PolicyAttachment, opts v1.UpdateOptions) (result *v1alpha1.PolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(policyattachmentsResource, c.ns, policyAttachment), &v1alpha1.PolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAttachment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePolicyAttachments) UpdateStatus(ctx context.Context, policyAttachment *v1alpha1.PolicyAttachment, opts v1.UpdateOptions) (*v1alpha1.PolicyAttachment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(policyattachmentsResource, "status", c.ns, policyAttachment), &v1alpha1.PolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAttachment), err
}

// Delete takes name of the policyAttachment and deletes it. Returns an error if one occurs.
func (c *FakePolicyAttachments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(policyattachmentsResource, c.ns, name), &v1alpha1.PolicyAttachment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePolicyAttachments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(policyattachmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PolicyAttachmentList{})
	return err
}

// Patch applies the patch and returns the patched policyAttachment.
func (c *FakePolicyAttachments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PolicyAttachment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(policyattachmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PolicyAttachment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PolicyAttachment), err
}
