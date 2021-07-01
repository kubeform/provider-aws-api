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

// FakeCaches implements CacheInterface
type FakeCaches struct {
	Fake *FakeStoragegatewayV1alpha1
	ns   string
}

var cachesResource = schema.GroupVersionResource{Group: "storagegateway.aws.kubeform.com", Version: "v1alpha1", Resource: "caches"}

var cachesKind = schema.GroupVersionKind{Group: "storagegateway.aws.kubeform.com", Version: "v1alpha1", Kind: "Cache"}

// Get takes name of the cache, and returns the corresponding cache object, and an error if there is any.
func (c *FakeCaches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Cache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cachesResource, c.ns, name), &v1alpha1.Cache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cache), err
}

// List takes label and field selectors, and returns the list of Caches that match those selectors.
func (c *FakeCaches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CacheList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cachesResource, cachesKind, c.ns, opts), &v1alpha1.CacheList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CacheList{ListMeta: obj.(*v1alpha1.CacheList).ListMeta}
	for _, item := range obj.(*v1alpha1.CacheList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested caches.
func (c *FakeCaches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cachesResource, c.ns, opts))

}

// Create takes the representation of a cache and creates it.  Returns the server's representation of the cache, and an error, if there is any.
func (c *FakeCaches) Create(ctx context.Context, cache *v1alpha1.Cache, opts v1.CreateOptions) (result *v1alpha1.Cache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cachesResource, c.ns, cache), &v1alpha1.Cache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cache), err
}

// Update takes the representation of a cache and updates it. Returns the server's representation of the cache, and an error, if there is any.
func (c *FakeCaches) Update(ctx context.Context, cache *v1alpha1.Cache, opts v1.UpdateOptions) (result *v1alpha1.Cache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cachesResource, c.ns, cache), &v1alpha1.Cache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cache), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCaches) UpdateStatus(ctx context.Context, cache *v1alpha1.Cache, opts v1.UpdateOptions) (*v1alpha1.Cache, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cachesResource, "status", c.ns, cache), &v1alpha1.Cache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cache), err
}

// Delete takes name of the cache and deletes it. Returns an error if one occurs.
func (c *FakeCaches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cachesResource, c.ns, name), &v1alpha1.Cache{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCaches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cachesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CacheList{})
	return err
}

// Patch applies the patch and returns the patched cache.
func (c *FakeCaches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Cache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cachesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Cache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cache), err
}
