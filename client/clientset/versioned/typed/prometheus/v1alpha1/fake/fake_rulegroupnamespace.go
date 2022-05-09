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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/prometheus/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRuleGroupNamespaces implements RuleGroupNamespaceInterface
type FakeRuleGroupNamespaces struct {
	Fake *FakePrometheusV1alpha1
	ns   string
}

var rulegroupnamespacesResource = schema.GroupVersionResource{Group: "prometheus.aws.kubeform.com", Version: "v1alpha1", Resource: "rulegroupnamespaces"}

var rulegroupnamespacesKind = schema.GroupVersionKind{Group: "prometheus.aws.kubeform.com", Version: "v1alpha1", Kind: "RuleGroupNamespace"}

// Get takes name of the ruleGroupNamespace, and returns the corresponding ruleGroupNamespace object, and an error if there is any.
func (c *FakeRuleGroupNamespaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RuleGroupNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rulegroupnamespacesResource, c.ns, name), &v1alpha1.RuleGroupNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroupNamespace), err
}

// List takes label and field selectors, and returns the list of RuleGroupNamespaces that match those selectors.
func (c *FakeRuleGroupNamespaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RuleGroupNamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rulegroupnamespacesResource, rulegroupnamespacesKind, c.ns, opts), &v1alpha1.RuleGroupNamespaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RuleGroupNamespaceList{ListMeta: obj.(*v1alpha1.RuleGroupNamespaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.RuleGroupNamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ruleGroupNamespaces.
func (c *FakeRuleGroupNamespaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rulegroupnamespacesResource, c.ns, opts))

}

// Create takes the representation of a ruleGroupNamespace and creates it.  Returns the server's representation of the ruleGroupNamespace, and an error, if there is any.
func (c *FakeRuleGroupNamespaces) Create(ctx context.Context, ruleGroupNamespace *v1alpha1.RuleGroupNamespace, opts v1.CreateOptions) (result *v1alpha1.RuleGroupNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rulegroupnamespacesResource, c.ns, ruleGroupNamespace), &v1alpha1.RuleGroupNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroupNamespace), err
}

// Update takes the representation of a ruleGroupNamespace and updates it. Returns the server's representation of the ruleGroupNamespace, and an error, if there is any.
func (c *FakeRuleGroupNamespaces) Update(ctx context.Context, ruleGroupNamespace *v1alpha1.RuleGroupNamespace, opts v1.UpdateOptions) (result *v1alpha1.RuleGroupNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rulegroupnamespacesResource, c.ns, ruleGroupNamespace), &v1alpha1.RuleGroupNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroupNamespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRuleGroupNamespaces) UpdateStatus(ctx context.Context, ruleGroupNamespace *v1alpha1.RuleGroupNamespace, opts v1.UpdateOptions) (*v1alpha1.RuleGroupNamespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rulegroupnamespacesResource, "status", c.ns, ruleGroupNamespace), &v1alpha1.RuleGroupNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroupNamespace), err
}

// Delete takes name of the ruleGroupNamespace and deletes it. Returns an error if one occurs.
func (c *FakeRuleGroupNamespaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rulegroupnamespacesResource, c.ns, name), &v1alpha1.RuleGroupNamespace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRuleGroupNamespaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rulegroupnamespacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RuleGroupNamespaceList{})
	return err
}

// Patch applies the patch and returns the patched ruleGroupNamespace.
func (c *FakeRuleGroupNamespaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RuleGroupNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rulegroupnamespacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RuleGroupNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RuleGroupNamespace), err
}