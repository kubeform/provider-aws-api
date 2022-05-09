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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/appsync/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApiCaches implements ApiCacheInterface
type FakeApiCaches struct {
	Fake *FakeAppsyncV1alpha1
	ns   string
}

var apicachesResource = schema.GroupVersionResource{Group: "appsync.aws.kubeform.com", Version: "v1alpha1", Resource: "apicaches"}

var apicachesKind = schema.GroupVersionKind{Group: "appsync.aws.kubeform.com", Version: "v1alpha1", Kind: "ApiCache"}

// Get takes name of the apiCache, and returns the corresponding apiCache object, and an error if there is any.
func (c *FakeApiCaches) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apicachesResource, c.ns, name), &v1alpha1.ApiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiCache), err
}

// List takes label and field selectors, and returns the list of ApiCaches that match those selectors.
func (c *FakeApiCaches) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiCacheList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apicachesResource, apicachesKind, c.ns, opts), &v1alpha1.ApiCacheList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiCacheList{ListMeta: obj.(*v1alpha1.ApiCacheList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiCacheList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiCaches.
func (c *FakeApiCaches) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apicachesResource, c.ns, opts))

}

// Create takes the representation of a apiCache and creates it.  Returns the server's representation of the apiCache, and an error, if there is any.
func (c *FakeApiCaches) Create(ctx context.Context, apiCache *v1alpha1.ApiCache, opts v1.CreateOptions) (result *v1alpha1.ApiCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apicachesResource, c.ns, apiCache), &v1alpha1.ApiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiCache), err
}

// Update takes the representation of a apiCache and updates it. Returns the server's representation of the apiCache, and an error, if there is any.
func (c *FakeApiCaches) Update(ctx context.Context, apiCache *v1alpha1.ApiCache, opts v1.UpdateOptions) (result *v1alpha1.ApiCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apicachesResource, c.ns, apiCache), &v1alpha1.ApiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiCache), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiCaches) UpdateStatus(ctx context.Context, apiCache *v1alpha1.ApiCache, opts v1.UpdateOptions) (*v1alpha1.ApiCache, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apicachesResource, "status", c.ns, apiCache), &v1alpha1.ApiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiCache), err
}

// Delete takes name of the apiCache and deletes it. Returns an error if one occurs.
func (c *FakeApiCaches) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apicachesResource, c.ns, name), &v1alpha1.ApiCache{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiCaches) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apicachesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiCacheList{})
	return err
}

// Patch applies the patch and returns the patched apiCache.
func (c *FakeApiCaches) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apicachesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiCache), err
}