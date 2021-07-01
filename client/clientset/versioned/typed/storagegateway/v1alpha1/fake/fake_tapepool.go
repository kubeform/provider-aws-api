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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/storagegateway/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTapePools implements TapePoolInterface
type FakeTapePools struct {
	Fake *FakeStoragegatewayV1alpha1
	ns   string
}

var tapepoolsResource = schema.GroupVersionResource{Group: "storagegateway.aws.kubeform.com", Version: "v1alpha1", Resource: "tapepools"}

var tapepoolsKind = schema.GroupVersionKind{Group: "storagegateway.aws.kubeform.com", Version: "v1alpha1", Kind: "TapePool"}

// Get takes name of the tapePool, and returns the corresponding tapePool object, and an error if there is any.
func (c *FakeTapePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TapePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(tapepoolsResource, c.ns, name), &v1alpha1.TapePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TapePool), err
}

// List takes label and field selectors, and returns the list of TapePools that match those selectors.
func (c *FakeTapePools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TapePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(tapepoolsResource, tapepoolsKind, c.ns, opts), &v1alpha1.TapePoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TapePoolList{ListMeta: obj.(*v1alpha1.TapePoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.TapePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested tapePools.
func (c *FakeTapePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(tapepoolsResource, c.ns, opts))

}

// Create takes the representation of a tapePool and creates it.  Returns the server's representation of the tapePool, and an error, if there is any.
func (c *FakeTapePools) Create(ctx context.Context, tapePool *v1alpha1.TapePool, opts v1.CreateOptions) (result *v1alpha1.TapePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(tapepoolsResource, c.ns, tapePool), &v1alpha1.TapePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TapePool), err
}

// Update takes the representation of a tapePool and updates it. Returns the server's representation of the tapePool, and an error, if there is any.
func (c *FakeTapePools) Update(ctx context.Context, tapePool *v1alpha1.TapePool, opts v1.UpdateOptions) (result *v1alpha1.TapePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(tapepoolsResource, c.ns, tapePool), &v1alpha1.TapePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TapePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTapePools) UpdateStatus(ctx context.Context, tapePool *v1alpha1.TapePool, opts v1.UpdateOptions) (*v1alpha1.TapePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(tapepoolsResource, "status", c.ns, tapePool), &v1alpha1.TapePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TapePool), err
}

// Delete takes name of the tapePool and deletes it. Returns an error if one occurs.
func (c *FakeTapePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(tapepoolsResource, c.ns, name), &v1alpha1.TapePool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTapePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(tapepoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TapePoolList{})
	return err
}

// Patch applies the patch and returns the patched tapePool.
func (c *FakeTapePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TapePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(tapepoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TapePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TapePool), err
}
