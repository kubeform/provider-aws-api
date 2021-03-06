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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecr/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePullThroughCacheRules implements PullThroughCacheRuleInterface
type FakePullThroughCacheRules struct {
	Fake *FakeEcrV1alpha1
	ns   string
}

var pullthroughcacherulesResource = schema.GroupVersionResource{Group: "ecr.aws.kubeform.com", Version: "v1alpha1", Resource: "pullthroughcacherules"}

var pullthroughcacherulesKind = schema.GroupVersionKind{Group: "ecr.aws.kubeform.com", Version: "v1alpha1", Kind: "PullThroughCacheRule"}

// Get takes name of the pullThroughCacheRule, and returns the corresponding pullThroughCacheRule object, and an error if there is any.
func (c *FakePullThroughCacheRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PullThroughCacheRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pullthroughcacherulesResource, c.ns, name), &v1alpha1.PullThroughCacheRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PullThroughCacheRule), err
}

// List takes label and field selectors, and returns the list of PullThroughCacheRules that match those selectors.
func (c *FakePullThroughCacheRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PullThroughCacheRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pullthroughcacherulesResource, pullthroughcacherulesKind, c.ns, opts), &v1alpha1.PullThroughCacheRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PullThroughCacheRuleList{ListMeta: obj.(*v1alpha1.PullThroughCacheRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.PullThroughCacheRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pullThroughCacheRules.
func (c *FakePullThroughCacheRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pullthroughcacherulesResource, c.ns, opts))

}

// Create takes the representation of a pullThroughCacheRule and creates it.  Returns the server's representation of the pullThroughCacheRule, and an error, if there is any.
func (c *FakePullThroughCacheRules) Create(ctx context.Context, pullThroughCacheRule *v1alpha1.PullThroughCacheRule, opts v1.CreateOptions) (result *v1alpha1.PullThroughCacheRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pullthroughcacherulesResource, c.ns, pullThroughCacheRule), &v1alpha1.PullThroughCacheRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PullThroughCacheRule), err
}

// Update takes the representation of a pullThroughCacheRule and updates it. Returns the server's representation of the pullThroughCacheRule, and an error, if there is any.
func (c *FakePullThroughCacheRules) Update(ctx context.Context, pullThroughCacheRule *v1alpha1.PullThroughCacheRule, opts v1.UpdateOptions) (result *v1alpha1.PullThroughCacheRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pullthroughcacherulesResource, c.ns, pullThroughCacheRule), &v1alpha1.PullThroughCacheRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PullThroughCacheRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePullThroughCacheRules) UpdateStatus(ctx context.Context, pullThroughCacheRule *v1alpha1.PullThroughCacheRule, opts v1.UpdateOptions) (*v1alpha1.PullThroughCacheRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pullthroughcacherulesResource, "status", c.ns, pullThroughCacheRule), &v1alpha1.PullThroughCacheRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PullThroughCacheRule), err
}

// Delete takes name of the pullThroughCacheRule and deletes it. Returns an error if one occurs.
func (c *FakePullThroughCacheRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pullthroughcacherulesResource, c.ns, name), &v1alpha1.PullThroughCacheRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePullThroughCacheRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pullthroughcacherulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PullThroughCacheRuleList{})
	return err
}

// Patch applies the patch and returns the patched pullThroughCacheRule.
func (c *FakePullThroughCacheRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PullThroughCacheRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pullthroughcacherulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PullThroughCacheRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PullThroughCacheRule), err
}
