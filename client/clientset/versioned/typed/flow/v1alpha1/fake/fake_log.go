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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/flow/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLogs implements LogInterface
type FakeLogs struct {
	Fake *FakeFlowV1alpha1
	ns   string
}

var logsResource = schema.GroupVersionResource{Group: "flow.aws.kubeform.com", Version: "v1alpha1", Resource: "logs"}

var logsKind = schema.GroupVersionKind{Group: "flow.aws.kubeform.com", Version: "v1alpha1", Kind: "Log"}

// Get takes name of the log, and returns the corresponding log object, and an error if there is any.
func (c *FakeLogs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Log, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logsResource, c.ns, name), &v1alpha1.Log{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Log), err
}

// List takes label and field selectors, and returns the list of Logs that match those selectors.
func (c *FakeLogs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logsResource, logsKind, c.ns, opts), &v1alpha1.LogList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogList{ListMeta: obj.(*v1alpha1.LogList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logs.
func (c *FakeLogs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logsResource, c.ns, opts))

}

// Create takes the representation of a log and creates it.  Returns the server's representation of the log, and an error, if there is any.
func (c *FakeLogs) Create(ctx context.Context, log *v1alpha1.Log, opts v1.CreateOptions) (result *v1alpha1.Log, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logsResource, c.ns, log), &v1alpha1.Log{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Log), err
}

// Update takes the representation of a log and updates it. Returns the server's representation of the log, and an error, if there is any.
func (c *FakeLogs) Update(ctx context.Context, log *v1alpha1.Log, opts v1.UpdateOptions) (result *v1alpha1.Log, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logsResource, c.ns, log), &v1alpha1.Log{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Log), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogs) UpdateStatus(ctx context.Context, log *v1alpha1.Log, opts v1.UpdateOptions) (*v1alpha1.Log, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logsResource, "status", c.ns, log), &v1alpha1.Log{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Log), err
}

// Delete takes name of the log and deletes it. Returns an error if one occurs.
func (c *FakeLogs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logsResource, c.ns, name), &v1alpha1.Log{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogList{})
	return err
}

// Patch applies the patch and returns the patched log.
func (c *FakeLogs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Log, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Log{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Log), err
}
