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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/lb/v1alpha1"
)

// FakeListenerRules implements ListenerRuleInterface
type FakeListenerRules struct {
	Fake *FakeLbV1alpha1
	ns   string
}

var listenerrulesResource = schema.GroupVersionResource{Group: "lb.aws.kubeform.com", Version: "v1alpha1", Resource: "listenerrules"}

var listenerrulesKind = schema.GroupVersionKind{Group: "lb.aws.kubeform.com", Version: "v1alpha1", Kind: "ListenerRule"}

// Get takes name of the listenerRule, and returns the corresponding listenerRule object, and an error if there is any.
func (c *FakeListenerRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(listenerrulesResource, c.ns, name), &v1alpha1.ListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerRule), err
}

// List takes label and field selectors, and returns the list of ListenerRules that match those selectors.
func (c *FakeListenerRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ListenerRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(listenerrulesResource, listenerrulesKind, c.ns, opts), &v1alpha1.ListenerRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ListenerRuleList{ListMeta: obj.(*v1alpha1.ListenerRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ListenerRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested listenerRules.
func (c *FakeListenerRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(listenerrulesResource, c.ns, opts))

}

// Create takes the representation of a listenerRule and creates it.  Returns the server's representation of the listenerRule, and an error, if there is any.
func (c *FakeListenerRules) Create(ctx context.Context, listenerRule *v1alpha1.ListenerRule, opts v1.CreateOptions) (result *v1alpha1.ListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(listenerrulesResource, c.ns, listenerRule), &v1alpha1.ListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerRule), err
}

// Update takes the representation of a listenerRule and updates it. Returns the server's representation of the listenerRule, and an error, if there is any.
func (c *FakeListenerRules) Update(ctx context.Context, listenerRule *v1alpha1.ListenerRule, opts v1.UpdateOptions) (result *v1alpha1.ListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(listenerrulesResource, c.ns, listenerRule), &v1alpha1.ListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeListenerRules) UpdateStatus(ctx context.Context, listenerRule *v1alpha1.ListenerRule, opts v1.UpdateOptions) (*v1alpha1.ListenerRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(listenerrulesResource, "status", c.ns, listenerRule), &v1alpha1.ListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerRule), err
}

// Delete takes name of the listenerRule and deletes it. Returns an error if one occurs.
func (c *FakeListenerRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(listenerrulesResource, c.ns, name), &v1alpha1.ListenerRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeListenerRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(listenerrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ListenerRuleList{})
	return err
}

// Patch applies the patch and returns the patched listenerRule.
func (c *FakeListenerRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ListenerRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(listenerrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ListenerRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerRule), err
}
