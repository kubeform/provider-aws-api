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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/lex/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIntents implements IntentInterface
type FakeIntents struct {
	Fake *FakeLexV1alpha1
	ns   string
}

var intentsResource = schema.GroupVersionResource{Group: "lex.aws.kubeform.com", Version: "v1alpha1", Resource: "intents"}

var intentsKind = schema.GroupVersionKind{Group: "lex.aws.kubeform.com", Version: "v1alpha1", Kind: "Intent"}

// Get takes name of the intent, and returns the corresponding intent object, and an error if there is any.
func (c *FakeIntents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Intent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(intentsResource, c.ns, name), &v1alpha1.Intent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Intent), err
}

// List takes label and field selectors, and returns the list of Intents that match those selectors.
func (c *FakeIntents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IntentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(intentsResource, intentsKind, c.ns, opts), &v1alpha1.IntentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IntentList{ListMeta: obj.(*v1alpha1.IntentList).ListMeta}
	for _, item := range obj.(*v1alpha1.IntentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested intents.
func (c *FakeIntents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(intentsResource, c.ns, opts))

}

// Create takes the representation of a intent and creates it.  Returns the server's representation of the intent, and an error, if there is any.
func (c *FakeIntents) Create(ctx context.Context, intent *v1alpha1.Intent, opts v1.CreateOptions) (result *v1alpha1.Intent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(intentsResource, c.ns, intent), &v1alpha1.Intent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Intent), err
}

// Update takes the representation of a intent and updates it. Returns the server's representation of the intent, and an error, if there is any.
func (c *FakeIntents) Update(ctx context.Context, intent *v1alpha1.Intent, opts v1.UpdateOptions) (result *v1alpha1.Intent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(intentsResource, c.ns, intent), &v1alpha1.Intent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Intent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIntents) UpdateStatus(ctx context.Context, intent *v1alpha1.Intent, opts v1.UpdateOptions) (*v1alpha1.Intent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(intentsResource, "status", c.ns, intent), &v1alpha1.Intent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Intent), err
}

// Delete takes name of the intent and deletes it. Returns an error if one occurs.
func (c *FakeIntents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(intentsResource, c.ns, name), &v1alpha1.Intent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIntents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(intentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IntentList{})
	return err
}

// Patch applies the patch and returns the patched intent.
func (c *FakeIntents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Intent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(intentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Intent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Intent), err
}
