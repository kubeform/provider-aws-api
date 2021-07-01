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

// FakeRuleGroups implements RuleGroupInterface
type FakeRuleGroups struct {
	Fake *FakeNetworkfirewallV1alpha1
	ns   string
}

var rulegroupsResource = schema.GroupVersionResource{Group: "networkfirewall.aws.kubeform.com", Version: "v1alpha1", Resource: "rulegroups"}

var rulegroupsKind = schema.GroupVersionKind{Group: "networkfirewall.aws.kubeform.com", Version: "v1alpha1", Kind: "RuleGroup"}

// Get takes name of the ruleGroup, and returns the corresponding ruleGroup object, and an error if there is any.
func (c *FakeRuleGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rulegroupsResource, c.ns, name), &v1alpha1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroup), err
}

// List takes label and field selectors, and returns the list of RuleGroups that match those selectors.
func (c *FakeRuleGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RuleGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rulegroupsResource, rulegroupsKind, c.ns, opts), &v1alpha1.RuleGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RuleGroupList{ListMeta: obj.(*v1alpha1.RuleGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.RuleGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ruleGroups.
func (c *FakeRuleGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rulegroupsResource, c.ns, opts))

}

// Create takes the representation of a ruleGroup and creates it.  Returns the server's representation of the ruleGroup, and an error, if there is any.
func (c *FakeRuleGroups) Create(ctx context.Context, ruleGroup *v1alpha1.RuleGroup, opts v1.CreateOptions) (result *v1alpha1.RuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rulegroupsResource, c.ns, ruleGroup), &v1alpha1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroup), err
}

// Update takes the representation of a ruleGroup and updates it. Returns the server's representation of the ruleGroup, and an error, if there is any.
func (c *FakeRuleGroups) Update(ctx context.Context, ruleGroup *v1alpha1.RuleGroup, opts v1.UpdateOptions) (result *v1alpha1.RuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rulegroupsResource, c.ns, ruleGroup), &v1alpha1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRuleGroups) UpdateStatus(ctx context.Context, ruleGroup *v1alpha1.RuleGroup, opts v1.UpdateOptions) (*v1alpha1.RuleGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rulegroupsResource, "status", c.ns, ruleGroup), &v1alpha1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroup), err
}

// Delete takes name of the ruleGroup and deletes it. Returns an error if one occurs.
func (c *FakeRuleGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rulegroupsResource, c.ns, name), &v1alpha1.RuleGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRuleGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rulegroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RuleGroupList{})
	return err
}

// Patch applies the patch and returns the patched ruleGroup.
func (c *FakeRuleGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rulegroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroup), err
}
