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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudwatch/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEventTargets implements EventTargetInterface
type FakeEventTargets struct {
	Fake *FakeCloudwatchV1alpha1
	ns   string
}

var eventtargetsResource = schema.GroupVersionResource{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Resource: "eventtargets"}

var eventtargetsKind = schema.GroupVersionKind{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Kind: "EventTarget"}

// Get takes name of the eventTarget, and returns the corresponding eventTarget object, and an error if there is any.
func (c *FakeEventTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventtargetsResource, c.ns, name), &v1alpha1.EventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventTarget), err
}

// List takes label and field selectors, and returns the list of EventTargets that match those selectors.
func (c *FakeEventTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventtargetsResource, eventtargetsKind, c.ns, opts), &v1alpha1.EventTargetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventTargetList{ListMeta: obj.(*v1alpha1.EventTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventTargets.
func (c *FakeEventTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventtargetsResource, c.ns, opts))

}

// Create takes the representation of a eventTarget and creates it.  Returns the server's representation of the eventTarget, and an error, if there is any.
func (c *FakeEventTargets) Create(ctx context.Context, eventTarget *v1alpha1.EventTarget, opts v1.CreateOptions) (result *v1alpha1.EventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventtargetsResource, c.ns, eventTarget), &v1alpha1.EventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventTarget), err
}

// Update takes the representation of a eventTarget and updates it. Returns the server's representation of the eventTarget, and an error, if there is any.
func (c *FakeEventTargets) Update(ctx context.Context, eventTarget *v1alpha1.EventTarget, opts v1.UpdateOptions) (result *v1alpha1.EventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventtargetsResource, c.ns, eventTarget), &v1alpha1.EventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventTargets) UpdateStatus(ctx context.Context, eventTarget *v1alpha1.EventTarget, opts v1.UpdateOptions) (*v1alpha1.EventTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventtargetsResource, "status", c.ns, eventTarget), &v1alpha1.EventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventTarget), err
}

// Delete takes name of the eventTarget and deletes it. Returns an error if one occurs.
func (c *FakeEventTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventtargetsResource, c.ns, name), &v1alpha1.EventTarget{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventtargetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventTargetList{})
	return err
}

// Patch applies the patch and returns the patched eventTarget.
func (c *FakeEventTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventtargetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventTarget{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventTarget), err
}
