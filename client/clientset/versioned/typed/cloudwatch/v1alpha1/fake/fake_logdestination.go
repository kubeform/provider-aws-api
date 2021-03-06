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

// FakeLogDestinations implements LogDestinationInterface
type FakeLogDestinations struct {
	Fake *FakeCloudwatchV1alpha1
	ns   string
}

var logdestinationsResource = schema.GroupVersionResource{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Resource: "logdestinations"}

var logdestinationsKind = schema.GroupVersionKind{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Kind: "LogDestination"}

// Get takes name of the logDestination, and returns the corresponding logDestination object, and an error if there is any.
func (c *FakeLogDestinations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LogDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logdestinationsResource, c.ns, name), &v1alpha1.LogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestination), err
}

// List takes label and field selectors, and returns the list of LogDestinations that match those selectors.
func (c *FakeLogDestinations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogDestinationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logdestinationsResource, logdestinationsKind, c.ns, opts), &v1alpha1.LogDestinationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogDestinationList{ListMeta: obj.(*v1alpha1.LogDestinationList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogDestinationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logDestinations.
func (c *FakeLogDestinations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logdestinationsResource, c.ns, opts))

}

// Create takes the representation of a logDestination and creates it.  Returns the server's representation of the logDestination, and an error, if there is any.
func (c *FakeLogDestinations) Create(ctx context.Context, logDestination *v1alpha1.LogDestination, opts v1.CreateOptions) (result *v1alpha1.LogDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logdestinationsResource, c.ns, logDestination), &v1alpha1.LogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestination), err
}

// Update takes the representation of a logDestination and updates it. Returns the server's representation of the logDestination, and an error, if there is any.
func (c *FakeLogDestinations) Update(ctx context.Context, logDestination *v1alpha1.LogDestination, opts v1.UpdateOptions) (result *v1alpha1.LogDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logdestinationsResource, c.ns, logDestination), &v1alpha1.LogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestination), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogDestinations) UpdateStatus(ctx context.Context, logDestination *v1alpha1.LogDestination, opts v1.UpdateOptions) (*v1alpha1.LogDestination, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logdestinationsResource, "status", c.ns, logDestination), &v1alpha1.LogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestination), err
}

// Delete takes name of the logDestination and deletes it. Returns an error if one occurs.
func (c *FakeLogDestinations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logdestinationsResource, c.ns, name), &v1alpha1.LogDestination{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogDestinations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logdestinationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogDestinationList{})
	return err
}

// Patch applies the patch and returns the patched logDestination.
func (c *FakeLogDestinations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LogDestination, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logdestinationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogDestination{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestination), err
}
