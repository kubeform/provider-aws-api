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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/dynamodb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTableItems implements TableItemInterface
type FakeTableItems struct {
	Fake *FakeDynamodbV1alpha1
	ns   string
}

var tableitemsResource = schema.GroupVersionResource{Group: "dynamodb.aws.kubeform.com", Version: "v1alpha1", Resource: "tableitems"}

var tableitemsKind = schema.GroupVersionKind{Group: "dynamodb.aws.kubeform.com", Version: "v1alpha1", Kind: "TableItem"}

// Get takes name of the tableItem, and returns the corresponding tableItem object, and an error if there is any.
func (c *FakeTableItems) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tableitemsResource, c.ns, name), &v1alpha1.TableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableItem), err
}

// List takes label and field selectors, and returns the list of TableItems that match those selectors.
func (c *FakeTableItems) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TableItemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tableitemsResource, tableitemsKind, c.ns, opts), &v1alpha1.TableItemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TableItemList{ListMeta: obj.(*v1alpha1.TableItemList).ListMeta}
	for _, item := range obj.(*v1alpha1.TableItemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tableItems.
func (c *FakeTableItems) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tableitemsResource, c.ns, opts))

}

// Create takes the representation of a tableItem and creates it.  Returns the server's representation of the tableItem, and an error, if there is any.
func (c *FakeTableItems) Create(ctx context.Context, tableItem *v1alpha1.TableItem, opts v1.CreateOptions) (result *v1alpha1.TableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tableitemsResource, c.ns, tableItem), &v1alpha1.TableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableItem), err
}

// Update takes the representation of a tableItem and updates it. Returns the server's representation of the tableItem, and an error, if there is any.
func (c *FakeTableItems) Update(ctx context.Context, tableItem *v1alpha1.TableItem, opts v1.UpdateOptions) (result *v1alpha1.TableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tableitemsResource, c.ns, tableItem), &v1alpha1.TableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableItem), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTableItems) UpdateStatus(ctx context.Context, tableItem *v1alpha1.TableItem, opts v1.UpdateOptions) (*v1alpha1.TableItem, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tableitemsResource, "status", c.ns, tableItem), &v1alpha1.TableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableItem), err
}

// Delete takes name of the tableItem and deletes it. Returns an error if one occurs.
func (c *FakeTableItems) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tableitemsResource, c.ns, name), &v1alpha1.TableItem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTableItems) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tableitemsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TableItemList{})
	return err
}

// Patch applies the patch and returns the patched tableItem.
func (c *FakeTableItems) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TableItem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tableitemsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TableItem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TableItem), err
}
