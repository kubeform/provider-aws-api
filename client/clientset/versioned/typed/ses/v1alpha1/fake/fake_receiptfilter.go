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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReceiptFilters implements ReceiptFilterInterface
type FakeReceiptFilters struct {
	Fake *FakeSesV1alpha1
	ns   string
}

var receiptfiltersResource = schema.GroupVersionResource{Group: "ses.aws.kubeform.com", Version: "v1alpha1", Resource: "receiptfilters"}

var receiptfiltersKind = schema.GroupVersionKind{Group: "ses.aws.kubeform.com", Version: "v1alpha1", Kind: "ReceiptFilter"}

// Get takes name of the receiptFilter, and returns the corresponding receiptFilter object, and an error if there is any.
func (c *FakeReceiptFilters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ReceiptFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(receiptfiltersResource, c.ns, name), &v1alpha1.ReceiptFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReceiptFilter), err
}

// List takes label and field selectors, and returns the list of ReceiptFilters that match those selectors.
func (c *FakeReceiptFilters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReceiptFilterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(receiptfiltersResource, receiptfiltersKind, c.ns, opts), &v1alpha1.ReceiptFilterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReceiptFilterList{ListMeta: obj.(*v1alpha1.ReceiptFilterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReceiptFilterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested receiptFilters.
func (c *FakeReceiptFilters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(receiptfiltersResource, c.ns, opts))

}

// Create takes the representation of a receiptFilter and creates it.  Returns the server's representation of the receiptFilter, and an error, if there is any.
func (c *FakeReceiptFilters) Create(ctx context.Context, receiptFilter *v1alpha1.ReceiptFilter, opts v1.CreateOptions) (result *v1alpha1.ReceiptFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(receiptfiltersResource, c.ns, receiptFilter), &v1alpha1.ReceiptFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReceiptFilter), err
}

// Update takes the representation of a receiptFilter and updates it. Returns the server's representation of the receiptFilter, and an error, if there is any.
func (c *FakeReceiptFilters) Update(ctx context.Context, receiptFilter *v1alpha1.ReceiptFilter, opts v1.UpdateOptions) (result *v1alpha1.ReceiptFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(receiptfiltersResource, c.ns, receiptFilter), &v1alpha1.ReceiptFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReceiptFilter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReceiptFilters) UpdateStatus(ctx context.Context, receiptFilter *v1alpha1.ReceiptFilter, opts v1.UpdateOptions) (*v1alpha1.ReceiptFilter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(receiptfiltersResource, "status", c.ns, receiptFilter), &v1alpha1.ReceiptFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReceiptFilter), err
}

// Delete takes name of the receiptFilter and deletes it. Returns an error if one occurs.
func (c *FakeReceiptFilters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(receiptfiltersResource, c.ns, name), &v1alpha1.ReceiptFilter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReceiptFilters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(receiptfiltersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReceiptFilterList{})
	return err
}

// Patch applies the patch and returns the patched receiptFilter.
func (c *FakeReceiptFilters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ReceiptFilter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(receiptfiltersResource, c.ns, name, pt, data, subresources...), &v1alpha1.ReceiptFilter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReceiptFilter), err
}
