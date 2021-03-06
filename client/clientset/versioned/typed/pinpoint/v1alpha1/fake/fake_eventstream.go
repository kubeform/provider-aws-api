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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/pinpoint/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEventStreams implements EventStreamInterface
type FakeEventStreams struct {
	Fake *FakePinpointV1alpha1
	ns   string
}

var eventstreamsResource = schema.GroupVersionResource{Group: "pinpoint.aws.kubeform.com", Version: "v1alpha1", Resource: "eventstreams"}

var eventstreamsKind = schema.GroupVersionKind{Group: "pinpoint.aws.kubeform.com", Version: "v1alpha1", Kind: "EventStream"}

// Get takes name of the eventStream, and returns the corresponding eventStream object, and an error if there is any.
func (c *FakeEventStreams) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EventStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eventstreamsResource, c.ns, name), &v1alpha1.EventStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventStream), err
}

// List takes label and field selectors, and returns the list of EventStreams that match those selectors.
func (c *FakeEventStreams) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EventStreamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eventstreamsResource, eventstreamsKind, c.ns, opts), &v1alpha1.EventStreamList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EventStreamList{ListMeta: obj.(*v1alpha1.EventStreamList).ListMeta}
	for _, item := range obj.(*v1alpha1.EventStreamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eventStreams.
func (c *FakeEventStreams) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eventstreamsResource, c.ns, opts))

}

// Create takes the representation of a eventStream and creates it.  Returns the server's representation of the eventStream, and an error, if there is any.
func (c *FakeEventStreams) Create(ctx context.Context, eventStream *v1alpha1.EventStream, opts v1.CreateOptions) (result *v1alpha1.EventStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eventstreamsResource, c.ns, eventStream), &v1alpha1.EventStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventStream), err
}

// Update takes the representation of a eventStream and updates it. Returns the server's representation of the eventStream, and an error, if there is any.
func (c *FakeEventStreams) Update(ctx context.Context, eventStream *v1alpha1.EventStream, opts v1.UpdateOptions) (result *v1alpha1.EventStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eventstreamsResource, c.ns, eventStream), &v1alpha1.EventStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventStream), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEventStreams) UpdateStatus(ctx context.Context, eventStream *v1alpha1.EventStream, opts v1.UpdateOptions) (*v1alpha1.EventStream, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eventstreamsResource, "status", c.ns, eventStream), &v1alpha1.EventStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventStream), err
}

// Delete takes name of the eventStream and deletes it. Returns an error if one occurs.
func (c *FakeEventStreams) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eventstreamsResource, c.ns, name), &v1alpha1.EventStream{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEventStreams) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eventstreamsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EventStreamList{})
	return err
}

// Patch applies the patch and returns the patched eventStream.
func (c *FakeEventStreams) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EventStream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eventstreamsResource, c.ns, name, pt, data, subresources...), &v1alpha1.EventStream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EventStream), err
}
