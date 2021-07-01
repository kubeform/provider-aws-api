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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/networkfirewall/v1alpha1"
)

// FakeFirewallPolicies implements FirewallPolicyInterface
type FakeFirewallPolicies struct {
	Fake *FakeNetworkfirewallV1alpha1
	ns   string
}

var firewallpoliciesResource = schema.GroupVersionResource{Group: "networkfirewall.aws.kubeform.com", Version: "v1alpha1", Resource: "firewallpolicies"}

var firewallpoliciesKind = schema.GroupVersionKind{Group: "networkfirewall.aws.kubeform.com", Version: "v1alpha1", Kind: "FirewallPolicy"}

// Get takes name of the firewallPolicy, and returns the corresponding firewallPolicy object, and an error if there is any.
func (c *FakeFirewallPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(firewallpoliciesResource, c.ns, name), &v1alpha1.FirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallPolicy), err
}

// List takes label and field selectors, and returns the list of FirewallPolicies that match those selectors.
func (c *FakeFirewallPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FirewallPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(firewallpoliciesResource, firewallpoliciesKind, c.ns, opts), &v1alpha1.FirewallPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FirewallPolicyList{ListMeta: obj.(*v1alpha1.FirewallPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.FirewallPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested firewallPolicies.
func (c *FakeFirewallPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(firewallpoliciesResource, c.ns, opts))

}

// Create takes the representation of a firewallPolicy and creates it.  Returns the server's representation of the firewallPolicy, and an error, if there is any.
func (c *FakeFirewallPolicies) Create(ctx context.Context, firewallPolicy *v1alpha1.FirewallPolicy, opts v1.CreateOptions) (result *v1alpha1.FirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(firewallpoliciesResource, c.ns, firewallPolicy), &v1alpha1.FirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallPolicy), err
}

// Update takes the representation of a firewallPolicy and updates it. Returns the server's representation of the firewallPolicy, and an error, if there is any.
func (c *FakeFirewallPolicies) Update(ctx context.Context, firewallPolicy *v1alpha1.FirewallPolicy, opts v1.UpdateOptions) (result *v1alpha1.FirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(firewallpoliciesResource, c.ns, firewallPolicy), &v1alpha1.FirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFirewallPolicies) UpdateStatus(ctx context.Context, firewallPolicy *v1alpha1.FirewallPolicy, opts v1.UpdateOptions) (*v1alpha1.FirewallPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(firewallpoliciesResource, "status", c.ns, firewallPolicy), &v1alpha1.FirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallPolicy), err
}

// Delete takes name of the firewallPolicy and deletes it. Returns an error if one occurs.
func (c *FakeFirewallPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(firewallpoliciesResource, c.ns, name), &v1alpha1.FirewallPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFirewallPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(firewallpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FirewallPolicyList{})
	return err
}

// Patch applies the patch and returns the patched firewallPolicy.
func (c *FakeFirewallPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FirewallPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(firewallpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FirewallPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FirewallPolicy), err
}
