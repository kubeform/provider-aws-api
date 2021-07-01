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

// FakeResolverFirewallRuleGroupAssociations implements ResolverFirewallRuleGroupAssociationInterface
type FakeResolverFirewallRuleGroupAssociations struct {
	Fake *FakeRoute53V1alpha1
	ns   string
}

var resolverfirewallrulegroupassociationsResource = schema.GroupVersionResource{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Resource: "resolverfirewallrulegroupassociations"}

var resolverfirewallrulegroupassociationsKind = schema.GroupVersionKind{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Kind: "ResolverFirewallRuleGroupAssociation"}

// Get takes name of the resolverFirewallRuleGroupAssociation, and returns the corresponding resolverFirewallRuleGroupAssociation object, and an error if there is any.
func (c *FakeResolverFirewallRuleGroupAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResolverFirewallRuleGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resolverfirewallrulegroupassociationsResource, c.ns, name), &v1alpha1.ResolverFirewallRuleGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRuleGroupAssociation), err
}

// List takes label and field selectors, and returns the list of ResolverFirewallRuleGroupAssociations that match those selectors.
func (c *FakeResolverFirewallRuleGroupAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResolverFirewallRuleGroupAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resolverfirewallrulegroupassociationsResource, resolverfirewallrulegroupassociationsKind, c.ns, opts), &v1alpha1.ResolverFirewallRuleGroupAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResolverFirewallRuleGroupAssociationList{ListMeta: obj.(*v1alpha1.ResolverFirewallRuleGroupAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResolverFirewallRuleGroupAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resolverFirewallRuleGroupAssociations.
func (c *FakeResolverFirewallRuleGroupAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resolverfirewallrulegroupassociationsResource, c.ns, opts))

}

// Create takes the representation of a resolverFirewallRuleGroupAssociation and creates it.  Returns the server's representation of the resolverFirewallRuleGroupAssociation, and an error, if there is any.
func (c *FakeResolverFirewallRuleGroupAssociations) Create(ctx context.Context, resolverFirewallRuleGroupAssociation *v1alpha1.ResolverFirewallRuleGroupAssociation, opts v1.CreateOptions) (result *v1alpha1.ResolverFirewallRuleGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resolverfirewallrulegroupassociationsResource, c.ns, resolverFirewallRuleGroupAssociation), &v1alpha1.ResolverFirewallRuleGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRuleGroupAssociation), err
}

// Update takes the representation of a resolverFirewallRuleGroupAssociation and updates it. Returns the server's representation of the resolverFirewallRuleGroupAssociation, and an error, if there is any.
func (c *FakeResolverFirewallRuleGroupAssociations) Update(ctx context.Context, resolverFirewallRuleGroupAssociation *v1alpha1.ResolverFirewallRuleGroupAssociation, opts v1.UpdateOptions) (result *v1alpha1.ResolverFirewallRuleGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resolverfirewallrulegroupassociationsResource, c.ns, resolverFirewallRuleGroupAssociation), &v1alpha1.ResolverFirewallRuleGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRuleGroupAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResolverFirewallRuleGroupAssociations) UpdateStatus(ctx context.Context, resolverFirewallRuleGroupAssociation *v1alpha1.ResolverFirewallRuleGroupAssociation, opts v1.UpdateOptions) (*v1alpha1.ResolverFirewallRuleGroupAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resolverfirewallrulegroupassociationsResource, "status", c.ns, resolverFirewallRuleGroupAssociation), &v1alpha1.ResolverFirewallRuleGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRuleGroupAssociation), err
}

// Delete takes name of the resolverFirewallRuleGroupAssociation and deletes it. Returns an error if one occurs.
func (c *FakeResolverFirewallRuleGroupAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resolverfirewallrulegroupassociationsResource, c.ns, name), &v1alpha1.ResolverFirewallRuleGroupAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResolverFirewallRuleGroupAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resolverfirewallrulegroupassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResolverFirewallRuleGroupAssociationList{})
	return err
}

// Patch applies the patch and returns the patched resolverFirewallRuleGroupAssociation.
func (c *FakeResolverFirewallRuleGroupAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResolverFirewallRuleGroupAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resolverfirewallrulegroupassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResolverFirewallRuleGroupAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverFirewallRuleGroupAssociation), err
}
