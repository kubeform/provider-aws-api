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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/route53/v1alpha1"
)

// FakeResolverRules implements ResolverRuleInterface
type FakeResolverRules struct {
	Fake *FakeRoute53V1alpha1
	ns   string
}

var resolverrulesResource = schema.GroupVersionResource{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Resource: "resolverrules"}

var resolverrulesKind = schema.GroupVersionKind{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Kind: "ResolverRule"}

// Get takes name of the resolverRule, and returns the corresponding resolverRule object, and an error if there is any.
func (c *FakeResolverRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResolverRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resolverrulesResource, c.ns, name), &v1alpha1.ResolverRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRule), err
}

// List takes label and field selectors, and returns the list of ResolverRules that match those selectors.
func (c *FakeResolverRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResolverRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resolverrulesResource, resolverrulesKind, c.ns, opts), &v1alpha1.ResolverRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResolverRuleList{ListMeta: obj.(*v1alpha1.ResolverRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResolverRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resolverRules.
func (c *FakeResolverRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resolverrulesResource, c.ns, opts))

}

// Create takes the representation of a resolverRule and creates it.  Returns the server's representation of the resolverRule, and an error, if there is any.
func (c *FakeResolverRules) Create(ctx context.Context, resolverRule *v1alpha1.ResolverRule, opts v1.CreateOptions) (result *v1alpha1.ResolverRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resolverrulesResource, c.ns, resolverRule), &v1alpha1.ResolverRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRule), err
}

// Update takes the representation of a resolverRule and updates it. Returns the server's representation of the resolverRule, and an error, if there is any.
func (c *FakeResolverRules) Update(ctx context.Context, resolverRule *v1alpha1.ResolverRule, opts v1.UpdateOptions) (result *v1alpha1.ResolverRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resolverrulesResource, c.ns, resolverRule), &v1alpha1.ResolverRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResolverRules) UpdateStatus(ctx context.Context, resolverRule *v1alpha1.ResolverRule, opts v1.UpdateOptions) (*v1alpha1.ResolverRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resolverrulesResource, "status", c.ns, resolverRule), &v1alpha1.ResolverRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRule), err
}

// Delete takes name of the resolverRule and deletes it. Returns an error if one occurs.
func (c *FakeResolverRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resolverrulesResource, c.ns, name), &v1alpha1.ResolverRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResolverRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resolverrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResolverRuleList{})
	return err
}

// Patch applies the patch and returns the patched resolverRule.
func (c *FakeResolverRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResolverRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resolverrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResolverRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRule), err
}
