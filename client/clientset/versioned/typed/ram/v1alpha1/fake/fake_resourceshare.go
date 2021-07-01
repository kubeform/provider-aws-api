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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ram/v1alpha1"
)

// FakeResourceShares implements ResourceShareInterface
type FakeResourceShares struct {
	Fake *FakeRamV1alpha1
	ns   string
}

var resourcesharesResource = schema.GroupVersionResource{Group: "ram.aws.kubeform.com", Version: "v1alpha1", Resource: "resourceshares"}

var resourcesharesKind = schema.GroupVersionKind{Group: "ram.aws.kubeform.com", Version: "v1alpha1", Kind: "ResourceShare"}

// Get takes name of the resourceShare, and returns the corresponding resourceShare object, and an error if there is any.
func (c *FakeResourceShares) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResourceShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcesharesResource, c.ns, name), &v1alpha1.ResourceShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceShare), err
}

// List takes label and field selectors, and returns the list of ResourceShares that match those selectors.
func (c *FakeResourceShares) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResourceShareList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcesharesResource, resourcesharesKind, c.ns, opts), &v1alpha1.ResourceShareList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceShareList{ListMeta: obj.(*v1alpha1.ResourceShareList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceShareList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceShares.
func (c *FakeResourceShares) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcesharesResource, c.ns, opts))

}

// Create takes the representation of a resourceShare and creates it.  Returns the server's representation of the resourceShare, and an error, if there is any.
func (c *FakeResourceShares) Create(ctx context.Context, resourceShare *v1alpha1.ResourceShare, opts v1.CreateOptions) (result *v1alpha1.ResourceShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcesharesResource, c.ns, resourceShare), &v1alpha1.ResourceShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceShare), err
}

// Update takes the representation of a resourceShare and updates it. Returns the server's representation of the resourceShare, and an error, if there is any.
func (c *FakeResourceShares) Update(ctx context.Context, resourceShare *v1alpha1.ResourceShare, opts v1.UpdateOptions) (result *v1alpha1.ResourceShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcesharesResource, c.ns, resourceShare), &v1alpha1.ResourceShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceShare), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceShares) UpdateStatus(ctx context.Context, resourceShare *v1alpha1.ResourceShare, opts v1.UpdateOptions) (*v1alpha1.ResourceShare, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcesharesResource, "status", c.ns, resourceShare), &v1alpha1.ResourceShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceShare), err
}

// Delete takes name of the resourceShare and deletes it. Returns an error if one occurs.
func (c *FakeResourceShares) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resourcesharesResource, c.ns, name), &v1alpha1.ResourceShare{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceShares) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcesharesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceShareList{})
	return err
}

// Patch applies the patch and returns the patched resourceShare.
func (c *FakeResourceShares) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResourceShare, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcesharesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResourceShare{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceShare), err
}
