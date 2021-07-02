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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/route53/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRecords implements RecordInterface
type FakeRecords struct {
	Fake *FakeRoute53V1alpha1
	ns   string
}

var recordsResource = schema.GroupVersionResource{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Resource: "records"}

var recordsKind = schema.GroupVersionKind{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Kind: "Record"}

// Get takes name of the record, and returns the corresponding record object, and an error if there is any.
func (c *FakeRecords) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Record, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(recordsResource, c.ns, name), &v1alpha1.Record{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}

// List takes label and field selectors, and returns the list of Records that match those selectors.
func (c *FakeRecords) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RecordList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(recordsResource, recordsKind, c.ns, opts), &v1alpha1.RecordList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RecordList{ListMeta: obj.(*v1alpha1.RecordList).ListMeta}
	for _, item := range obj.(*v1alpha1.RecordList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested records.
func (c *FakeRecords) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(recordsResource, c.ns, opts))

}

// Create takes the representation of a record and creates it.  Returns the server's representation of the record, and an error, if there is any.
func (c *FakeRecords) Create(ctx context.Context, record *v1alpha1.Record, opts v1.CreateOptions) (result *v1alpha1.Record, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(recordsResource, c.ns, record), &v1alpha1.Record{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}

// Update takes the representation of a record and updates it. Returns the server's representation of the record, and an error, if there is any.
func (c *FakeRecords) Update(ctx context.Context, record *v1alpha1.Record, opts v1.UpdateOptions) (result *v1alpha1.Record, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(recordsResource, c.ns, record), &v1alpha1.Record{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRecords) UpdateStatus(ctx context.Context, record *v1alpha1.Record, opts v1.UpdateOptions) (*v1alpha1.Record, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(recordsResource, "status", c.ns, record), &v1alpha1.Record{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}

// Delete takes name of the record and deletes it. Returns an error if one occurs.
func (c *FakeRecords) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(recordsResource, c.ns, name), &v1alpha1.Record{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRecords) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(recordsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RecordList{})
	return err
}

// Patch applies the patch and returns the patched record.
func (c *FakeRecords) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Record, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(recordsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Record{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Record), err
}
