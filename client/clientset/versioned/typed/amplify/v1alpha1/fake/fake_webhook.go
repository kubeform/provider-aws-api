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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/amplify/v1alpha1"
)

// FakeWebhooks implements WebhookInterface
type FakeWebhooks struct {
	Fake *FakeAmplifyV1alpha1
	ns   string
}

var webhooksResource = schema.GroupVersionResource{Group: "amplify.aws.kubeform.com", Version: "v1alpha1", Resource: "webhooks"}

var webhooksKind = schema.GroupVersionKind{Group: "amplify.aws.kubeform.com", Version: "v1alpha1", Kind: "Webhook"}

// Get takes name of the webhook, and returns the corresponding webhook object, and an error if there is any.
func (c *FakeWebhooks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Webhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(webhooksResource, c.ns, name), &v1alpha1.Webhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}

// List takes label and field selectors, and returns the list of Webhooks that match those selectors.
func (c *FakeWebhooks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WebhookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(webhooksResource, webhooksKind, c.ns, opts), &v1alpha1.WebhookList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WebhookList{ListMeta: obj.(*v1alpha1.WebhookList).ListMeta}
	for _, item := range obj.(*v1alpha1.WebhookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested webhooks.
func (c *FakeWebhooks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(webhooksResource, c.ns, opts))

}

// Create takes the representation of a webhook and creates it.  Returns the server's representation of the webhook, and an error, if there is any.
func (c *FakeWebhooks) Create(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.CreateOptions) (result *v1alpha1.Webhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(webhooksResource, c.ns, webhook), &v1alpha1.Webhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}

// Update takes the representation of a webhook and updates it. Returns the server's representation of the webhook, and an error, if there is any.
func (c *FakeWebhooks) Update(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.UpdateOptions) (result *v1alpha1.Webhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(webhooksResource, c.ns, webhook), &v1alpha1.Webhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWebhooks) UpdateStatus(ctx context.Context, webhook *v1alpha1.Webhook, opts v1.UpdateOptions) (*v1alpha1.Webhook, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(webhooksResource, "status", c.ns, webhook), &v1alpha1.Webhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}

// Delete takes name of the webhook and deletes it. Returns an error if one occurs.
func (c *FakeWebhooks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(webhooksResource, c.ns, name), &v1alpha1.Webhook{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWebhooks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(webhooksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WebhookList{})
	return err
}

// Patch applies the patch and returns the patched webhook.
func (c *FakeWebhooks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Webhook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(webhooksResource, c.ns, name, pt, data, subresources...), &v1alpha1.Webhook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Webhook), err
}
