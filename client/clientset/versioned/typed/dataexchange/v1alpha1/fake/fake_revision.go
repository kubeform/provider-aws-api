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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/dataexchange/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRevisions implements RevisionInterface
type FakeRevisions struct {
	Fake *FakeDataexchangeV1alpha1
	ns   string
}

var revisionsResource = schema.GroupVersionResource{Group: "dataexchange.aws.kubeform.com", Version: "v1alpha1", Resource: "revisions"}

var revisionsKind = schema.GroupVersionKind{Group: "dataexchange.aws.kubeform.com", Version: "v1alpha1", Kind: "Revision"}

// Get takes name of the revision, and returns the corresponding revision object, and an error if there is any.
func (c *FakeRevisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Revision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(revisionsResource, c.ns, name), &v1alpha1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Revision), err
}

// List takes label and field selectors, and returns the list of Revisions that match those selectors.
func (c *FakeRevisions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RevisionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(revisionsResource, revisionsKind, c.ns, opts), &v1alpha1.RevisionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RevisionList{ListMeta: obj.(*v1alpha1.RevisionList).ListMeta}
	for _, item := range obj.(*v1alpha1.RevisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested revisions.
func (c *FakeRevisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(revisionsResource, c.ns, opts))

}

// Create takes the representation of a revision and creates it.  Returns the server's representation of the revision, and an error, if there is any.
func (c *FakeRevisions) Create(ctx context.Context, revision *v1alpha1.Revision, opts v1.CreateOptions) (result *v1alpha1.Revision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(revisionsResource, c.ns, revision), &v1alpha1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Revision), err
}

// Update takes the representation of a revision and updates it. Returns the server's representation of the revision, and an error, if there is any.
func (c *FakeRevisions) Update(ctx context.Context, revision *v1alpha1.Revision, opts v1.UpdateOptions) (result *v1alpha1.Revision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(revisionsResource, c.ns, revision), &v1alpha1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Revision), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRevisions) UpdateStatus(ctx context.Context, revision *v1alpha1.Revision, opts v1.UpdateOptions) (*v1alpha1.Revision, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(revisionsResource, "status", c.ns, revision), &v1alpha1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Revision), err
}

// Delete takes name of the revision and deletes it. Returns an error if one occurs.
func (c *FakeRevisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(revisionsResource, c.ns, name), &v1alpha1.Revision{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRevisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(revisionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RevisionList{})
	return err
}

// Patch applies the patch and returns the patched revision.
func (c *FakeRevisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Revision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(revisionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Revision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Revision), err
}