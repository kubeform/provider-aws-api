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

// FakeResolverFirewallRules implements ResolverFirewallRuleInterface
type FakeResolverFirewallRules struct {
	Fake *FakeRoute53V1alpha1
	ns   string
}

var resolverfirewallrulesResource = schema.GroupVersionResource{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Resource: "resolverfirewallrules"}

var resolverfirewallrulesKind = schema.GroupVersionKind{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Kind: "ResolverFirewallRule"}

// Get takes name of the resolverFirewallRule, and returns the corresponding resolverFirewallRule object, and an error if there is any.
func (c *FakeResolverFirewallRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResolverFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resolverfirewallrulesResource, c.ns, name), &v1alpha1.ResolverFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRule), err
}

// List takes label and field selectors, and returns the list of ResolverFirewallRules that match those selectors.
func (c *FakeResolverFirewallRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResolverFirewallRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resolverfirewallrulesResource, resolverfirewallrulesKind, c.ns, opts), &v1alpha1.ResolverFirewallRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResolverFirewallRuleList{ListMeta: obj.(*v1alpha1.ResolverFirewallRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResolverFirewallRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resolverFirewallRules.
func (c *FakeResolverFirewallRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resolverfirewallrulesResource, c.ns, opts))

}

// Create takes the representation of a resolverFirewallRule and creates it.  Returns the server's representation of the resolverFirewallRule, and an error, if there is any.
func (c *FakeResolverFirewallRules) Create(ctx context.Context, resolverFirewallRule *v1alpha1.ResolverFirewallRule, opts v1.CreateOptions) (result *v1alpha1.ResolverFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resolverfirewallrulesResource, c.ns, resolverFirewallRule), &v1alpha1.ResolverFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRule), err
}

// Update takes the representation of a resolverFirewallRule and updates it. Returns the server's representation of the resolverFirewallRule, and an error, if there is any.
func (c *FakeResolverFirewallRules) Update(ctx context.Context, resolverFirewallRule *v1alpha1.ResolverFirewallRule, opts v1.UpdateOptions) (result *v1alpha1.ResolverFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resolverfirewallrulesResource, c.ns, resolverFirewallRule), &v1alpha1.ResolverFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResolverFirewallRules) UpdateStatus(ctx context.Context, resolverFirewallRule *v1alpha1.ResolverFirewallRule, opts v1.UpdateOptions) (*v1alpha1.ResolverFirewallRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resolverfirewallrulesResource, "status", c.ns, resolverFirewallRule), &v1alpha1.ResolverFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRule), err
}

// Delete takes name of the resolverFirewallRule and deletes it. Returns an error if one occurs.
func (c *FakeResolverFirewallRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resolverfirewallrulesResource, c.ns, name), &v1alpha1.ResolverFirewallRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResolverFirewallRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resolverfirewallrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResolverFirewallRuleList{})
	return err
}

// Patch applies the patch and returns the patched resolverFirewallRule.
func (c *FakeResolverFirewallRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResolverFirewallRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resolverfirewallrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResolverFirewallRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRule), err
}
