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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/config/v1alpha1"
)

// FakeOrganizationManagedRules implements OrganizationManagedRuleInterface
type FakeOrganizationManagedRules struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var organizationmanagedrulesResource = schema.GroupVersionResource{Group: "config.aws.kubeform.com", Version: "v1alpha1", Resource: "organizationmanagedrules"}

var organizationmanagedrulesKind = schema.GroupVersionKind{Group: "config.aws.kubeform.com", Version: "v1alpha1", Kind: "OrganizationManagedRule"}

// Get takes name of the organizationManagedRule, and returns the corresponding organizationManagedRule object, and an error if there is any.
func (c *FakeOrganizationManagedRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OrganizationManagedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(organizationmanagedrulesResource, c.ns, name), &v1alpha1.OrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationManagedRule), err
}

// List takes label and field selectors, and returns the list of OrganizationManagedRules that match those selectors.
func (c *FakeOrganizationManagedRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrganizationManagedRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(organizationmanagedrulesResource, organizationmanagedrulesKind, c.ns, opts), &v1alpha1.OrganizationManagedRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OrganizationManagedRuleList{ListMeta: obj.(*v1alpha1.OrganizationManagedRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.OrganizationManagedRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested organizationManagedRules.
func (c *FakeOrganizationManagedRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(organizationmanagedrulesResource, c.ns, opts))

}

// Create takes the representation of a organizationManagedRule and creates it.  Returns the server's representation of the organizationManagedRule, and an error, if there is any.
func (c *FakeOrganizationManagedRules) Create(ctx context.Context, organizationManagedRule *v1alpha1.OrganizationManagedRule, opts v1.CreateOptions) (result *v1alpha1.OrganizationManagedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(organizationmanagedrulesResource, c.ns, organizationManagedRule), &v1alpha1.OrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationManagedRule), err
}

// Update takes the representation of a organizationManagedRule and updates it. Returns the server's representation of the organizationManagedRule, and an error, if there is any.
func (c *FakeOrganizationManagedRules) Update(ctx context.Context, organizationManagedRule *v1alpha1.OrganizationManagedRule, opts v1.UpdateOptions) (result *v1alpha1.OrganizationManagedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(organizationmanagedrulesResource, c.ns, organizationManagedRule), &v1alpha1.OrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationManagedRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOrganizationManagedRules) UpdateStatus(ctx context.Context, organizationManagedRule *v1alpha1.OrganizationManagedRule, opts v1.UpdateOptions) (*v1alpha1.OrganizationManagedRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(organizationmanagedrulesResource, "status", c.ns, organizationManagedRule), &v1alpha1.OrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationManagedRule), err
}

// Delete takes name of the organizationManagedRule and deletes it. Returns an error if one occurs.
func (c *FakeOrganizationManagedRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(organizationmanagedrulesResource, c.ns, name), &v1alpha1.OrganizationManagedRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOrganizationManagedRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(organizationmanagedrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OrganizationManagedRuleList{})
	return err
}

// Patch applies the patch and returns the patched organizationManagedRule.
func (c *FakeOrganizationManagedRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrganizationManagedRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(organizationmanagedrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OrganizationManagedRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OrganizationManagedRule), err
}
