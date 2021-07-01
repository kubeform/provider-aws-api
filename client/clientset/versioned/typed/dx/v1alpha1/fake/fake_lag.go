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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dx/v1alpha1"
)

// FakeLags implements LagInterface
type FakeLags struct {
	Fake *FakeDxV1alpha1
	ns   string
}

var lagsResource = schema.GroupVersionResource{Group: "dx.aws.kubeform.com", Version: "v1alpha1", Resource: "lags"}

var lagsKind = schema.GroupVersionKind{Group: "dx.aws.kubeform.com", Version: "v1alpha1", Kind: "Lag"}

// Get takes name of the lag, and returns the corresponding lag object, and an error if there is any.
func (c *FakeLags) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Lag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lagsResource, c.ns, name), &v1alpha1.Lag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lag), err
}

// List takes label and field selectors, and returns the list of Lags that match those selectors.
func (c *FakeLags) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LagList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lagsResource, lagsKind, c.ns, opts), &v1alpha1.LagList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LagList{ListMeta: obj.(*v1alpha1.LagList).ListMeta}
	for _, item := range obj.(*v1alpha1.LagList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lags.
func (c *FakeLags) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lagsResource, c.ns, opts))

}

// Create takes the representation of a lag and creates it.  Returns the server's representation of the lag, and an error, if there is any.
func (c *FakeLags) Create(ctx context.Context, lag *v1alpha1.Lag, opts v1.CreateOptions) (result *v1alpha1.Lag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lagsResource, c.ns, lag), &v1alpha1.Lag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lag), err
}

// Update takes the representation of a lag and updates it. Returns the server's representation of the lag, and an error, if there is any.
func (c *FakeLags) Update(ctx context.Context, lag *v1alpha1.Lag, opts v1.UpdateOptions) (result *v1alpha1.Lag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lagsResource, c.ns, lag), &v1alpha1.Lag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lag), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLags) UpdateStatus(ctx context.Context, lag *v1alpha1.Lag, opts v1.UpdateOptions) (*v1alpha1.Lag, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lagsResource, "status", c.ns, lag), &v1alpha1.Lag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lag), err
}

// Delete takes name of the lag and deletes it. Returns an error if one occurs.
func (c *FakeLags) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lagsResource, c.ns, name), &v1alpha1.Lag{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLags) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lagsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LagList{})
	return err
}

// Patch applies the patch and returns the patched lag.
func (c *FakeLags) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lag, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lagsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Lag{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Lag), err
}
