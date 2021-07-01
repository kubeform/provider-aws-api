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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/securityhub/v1alpha1"
)

// FakeInsights implements InsightInterface
type FakeInsights struct {
	Fake *FakeSecurityhubV1alpha1
	ns   string
}

var insightsResource = schema.GroupVersionResource{Group: "securityhub.aws.kubeform.com", Version: "v1alpha1", Resource: "insights"}

var insightsKind = schema.GroupVersionKind{Group: "securityhub.aws.kubeform.com", Version: "v1alpha1", Kind: "Insight"}

// Get takes name of the insight, and returns the corresponding insight object, and an error if there is any.
func (c *FakeInsights) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Insight, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(insightsResource, c.ns, name), &v1alpha1.Insight{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Insight), err
}

// List takes label and field selectors, and returns the list of Insights that match those selectors.
func (c *FakeInsights) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InsightList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(insightsResource, insightsKind, c.ns, opts), &v1alpha1.InsightList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InsightList{ListMeta: obj.(*v1alpha1.InsightList).ListMeta}
	for _, item := range obj.(*v1alpha1.InsightList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested insights.
func (c *FakeInsights) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(insightsResource, c.ns, opts))

}

// Create takes the representation of a insight and creates it.  Returns the server's representation of the insight, and an error, if there is any.
func (c *FakeInsights) Create(ctx context.Context, insight *v1alpha1.Insight, opts v1.CreateOptions) (result *v1alpha1.Insight, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(insightsResource, c.ns, insight), &v1alpha1.Insight{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Insight), err
}

// Update takes the representation of a insight and updates it. Returns the server's representation of the insight, and an error, if there is any.
func (c *FakeInsights) Update(ctx context.Context, insight *v1alpha1.Insight, opts v1.UpdateOptions) (result *v1alpha1.Insight, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(insightsResource, c.ns, insight), &v1alpha1.Insight{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Insight), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInsights) UpdateStatus(ctx context.Context, insight *v1alpha1.Insight, opts v1.UpdateOptions) (*v1alpha1.Insight, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(insightsResource, "status", c.ns, insight), &v1alpha1.Insight{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Insight), err
}

// Delete takes name of the insight and deletes it. Returns an error if one occurs.
func (c *FakeInsights) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(insightsResource, c.ns, name), &v1alpha1.Insight{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInsights) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(insightsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InsightList{})
	return err
}

// Patch applies the patch and returns the patched insight.
func (c *FakeInsights) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Insight, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(insightsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Insight{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Insight), err
}
