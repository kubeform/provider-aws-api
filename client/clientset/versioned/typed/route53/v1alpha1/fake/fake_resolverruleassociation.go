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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/route53/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResolverRuleAssociations implements ResolverRuleAssociationInterface
type FakeResolverRuleAssociations struct {
	Fake *FakeRoute53V1alpha1
	ns   string
}

var resolverruleassociationsResource = schema.GroupVersionResource{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Resource: "resolverruleassociations"}

var resolverruleassociationsKind = schema.GroupVersionKind{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Kind: "ResolverRuleAssociation"}

// Get takes name of the resolverRuleAssociation, and returns the corresponding resolverRuleAssociation object, and an error if there is any.
func (c *FakeResolverRuleAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ResolverRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resolverruleassociationsResource, c.ns, name), &v1alpha1.ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRuleAssociation), err
}

// List takes label and field selectors, and returns the list of ResolverRuleAssociations that match those selectors.
func (c *FakeResolverRuleAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResolverRuleAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resolverruleassociationsResource, resolverruleassociationsKind, c.ns, opts), &v1alpha1.ResolverRuleAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResolverRuleAssociationList{ListMeta: obj.(*v1alpha1.ResolverRuleAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResolverRuleAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resolverRuleAssociations.
func (c *FakeResolverRuleAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resolverruleassociationsResource, c.ns, opts))

}

// Create takes the representation of a resolverRuleAssociation and creates it.  Returns the server's representation of the resolverRuleAssociation, and an error, if there is any.
func (c *FakeResolverRuleAssociations) Create(ctx context.Context, resolverRuleAssociation *v1alpha1.ResolverRuleAssociation, opts v1.CreateOptions) (result *v1alpha1.ResolverRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resolverruleassociationsResource, c.ns, resolverRuleAssociation), &v1alpha1.ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRuleAssociation), err
}

// Update takes the representation of a resolverRuleAssociation and updates it. Returns the server's representation of the resolverRuleAssociation, and an error, if there is any.
func (c *FakeResolverRuleAssociations) Update(ctx context.Context, resolverRuleAssociation *v1alpha1.ResolverRuleAssociation, opts v1.UpdateOptions) (result *v1alpha1.ResolverRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resolverruleassociationsResource, c.ns, resolverRuleAssociation), &v1alpha1.ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRuleAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResolverRuleAssociations) UpdateStatus(ctx context.Context, resolverRuleAssociation *v1alpha1.ResolverRuleAssociation, opts v1.UpdateOptions) (*v1alpha1.ResolverRuleAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resolverruleassociationsResource, "status", c.ns, resolverRuleAssociation), &v1alpha1.ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRuleAssociation), err
}

// Delete takes name of the resolverRuleAssociation and deletes it. Returns an error if one occurs.
func (c *FakeResolverRuleAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resolverruleassociationsResource, c.ns, name), &v1alpha1.ResolverRuleAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResolverRuleAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resolverruleassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResolverRuleAssociationList{})
	return err
}

// Patch applies the patch and returns the patched resolverRuleAssociation.
func (c *FakeResolverRuleAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ResolverRuleAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resolverruleassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ResolverRuleAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResolverRuleAssociation), err
}
