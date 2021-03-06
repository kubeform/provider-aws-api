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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/sns/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTopicSubscriptions implements TopicSubscriptionInterface
type FakeTopicSubscriptions struct {
	Fake *FakeSnsV1alpha1
	ns   string
}

var topicsubscriptionsResource = schema.GroupVersionResource{Group: "sns.aws.kubeform.com", Version: "v1alpha1", Resource: "topicsubscriptions"}

var topicsubscriptionsKind = schema.GroupVersionKind{Group: "sns.aws.kubeform.com", Version: "v1alpha1", Kind: "TopicSubscription"}

// Get takes name of the topicSubscription, and returns the corresponding topicSubscription object, and an error if there is any.
func (c *FakeTopicSubscriptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(topicsubscriptionsResource, c.ns, name), &v1alpha1.TopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TopicSubscription), err
}

// List takes label and field selectors, and returns the list of TopicSubscriptions that match those selectors.
func (c *FakeTopicSubscriptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TopicSubscriptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(topicsubscriptionsResource, topicsubscriptionsKind, c.ns, opts), &v1alpha1.TopicSubscriptionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TopicSubscriptionList{ListMeta: obj.(*v1alpha1.TopicSubscriptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.TopicSubscriptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested topicSubscriptions.
func (c *FakeTopicSubscriptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(topicsubscriptionsResource, c.ns, opts))

}

// Create takes the representation of a topicSubscription and creates it.  Returns the server's representation of the topicSubscription, and an error, if there is any.
func (c *FakeTopicSubscriptions) Create(ctx context.Context, topicSubscription *v1alpha1.TopicSubscription, opts v1.CreateOptions) (result *v1alpha1.TopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(topicsubscriptionsResource, c.ns, topicSubscription), &v1alpha1.TopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TopicSubscription), err
}

// Update takes the representation of a topicSubscription and updates it. Returns the server's representation of the topicSubscription, and an error, if there is any.
func (c *FakeTopicSubscriptions) Update(ctx context.Context, topicSubscription *v1alpha1.TopicSubscription, opts v1.UpdateOptions) (result *v1alpha1.TopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(topicsubscriptionsResource, c.ns, topicSubscription), &v1alpha1.TopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TopicSubscription), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTopicSubscriptions) UpdateStatus(ctx context.Context, topicSubscription *v1alpha1.TopicSubscription, opts v1.UpdateOptions) (*v1alpha1.TopicSubscription, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(topicsubscriptionsResource, "status", c.ns, topicSubscription), &v1alpha1.TopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TopicSubscription), err
}

// Delete takes name of the topicSubscription and deletes it. Returns an error if one occurs.
func (c *FakeTopicSubscriptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(topicsubscriptionsResource, c.ns, name), &v1alpha1.TopicSubscription{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTopicSubscriptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(topicsubscriptionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TopicSubscriptionList{})
	return err
}

// Patch applies the patch and returns the patched topicSubscription.
func (c *FakeTopicSubscriptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TopicSubscription, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(topicsubscriptionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TopicSubscription{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TopicSubscription), err
}
