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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/glue/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCrawlers implements CrawlerInterface
type FakeCrawlers struct {
	Fake *FakeGlueV1alpha1
	ns   string
}

var crawlersResource = schema.GroupVersionResource{Group: "glue.aws.kubeform.com", Version: "v1alpha1", Resource: "crawlers"}

var crawlersKind = schema.GroupVersionKind{Group: "glue.aws.kubeform.com", Version: "v1alpha1", Kind: "Crawler"}

// Get takes name of the crawler, and returns the corresponding crawler object, and an error if there is any.
func (c *FakeCrawlers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Crawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(crawlersResource, c.ns, name), &v1alpha1.Crawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Crawler), err
}

// List takes label and field selectors, and returns the list of Crawlers that match those selectors.
func (c *FakeCrawlers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CrawlerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(crawlersResource, crawlersKind, c.ns, opts), &v1alpha1.CrawlerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CrawlerList{ListMeta: obj.(*v1alpha1.CrawlerList).ListMeta}
	for _, item := range obj.(*v1alpha1.CrawlerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested crawlers.
func (c *FakeCrawlers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(crawlersResource, c.ns, opts))

}

// Create takes the representation of a crawler and creates it.  Returns the server's representation of the crawler, and an error, if there is any.
func (c *FakeCrawlers) Create(ctx context.Context, crawler *v1alpha1.Crawler, opts v1.CreateOptions) (result *v1alpha1.Crawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(crawlersResource, c.ns, crawler), &v1alpha1.Crawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Crawler), err
}

// Update takes the representation of a crawler and updates it. Returns the server's representation of the crawler, and an error, if there is any.
func (c *FakeCrawlers) Update(ctx context.Context, crawler *v1alpha1.Crawler, opts v1.UpdateOptions) (result *v1alpha1.Crawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(crawlersResource, c.ns, crawler), &v1alpha1.Crawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Crawler), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCrawlers) UpdateStatus(ctx context.Context, crawler *v1alpha1.Crawler, opts v1.UpdateOptions) (*v1alpha1.Crawler, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(crawlersResource, "status", c.ns, crawler), &v1alpha1.Crawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Crawler), err
}

// Delete takes name of the crawler and deletes it. Returns an error if one occurs.
func (c *FakeCrawlers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(crawlersResource, c.ns, name), &v1alpha1.Crawler{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCrawlers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(crawlersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CrawlerList{})
	return err
}

// Patch applies the patch and returns the patched crawler.
func (c *FakeCrawlers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Crawler, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(crawlersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Crawler{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Crawler), err
}
