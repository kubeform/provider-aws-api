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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/codecommit/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApprovalRuleTemplateAssociations implements ApprovalRuleTemplateAssociationInterface
type FakeApprovalRuleTemplateAssociations struct {
	Fake *FakeCodecommitV1alpha1
	ns   string
}

var approvalruletemplateassociationsResource = schema.GroupVersionResource{Group: "codecommit.aws.kubeform.com", Version: "v1alpha1", Resource: "approvalruletemplateassociations"}

var approvalruletemplateassociationsKind = schema.GroupVersionKind{Group: "codecommit.aws.kubeform.com", Version: "v1alpha1", Kind: "ApprovalRuleTemplateAssociation"}

// Get takes name of the approvalRuleTemplateAssociation, and returns the corresponding approvalRuleTemplateAssociation object, and an error if there is any.
func (c *FakeApprovalRuleTemplateAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApprovalRuleTemplateAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(approvalruletemplateassociationsResource, c.ns, name), &v1alpha1.ApprovalRuleTemplateAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApprovalRuleTemplateAssociation), err
}

// List takes label and field selectors, and returns the list of ApprovalRuleTemplateAssociations that match those selectors.
func (c *FakeApprovalRuleTemplateAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApprovalRuleTemplateAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(approvalruletemplateassociationsResource, approvalruletemplateassociationsKind, c.ns, opts), &v1alpha1.ApprovalRuleTemplateAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApprovalRuleTemplateAssociationList{ListMeta: obj.(*v1alpha1.ApprovalRuleTemplateAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApprovalRuleTemplateAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested approvalRuleTemplateAssociations.
func (c *FakeApprovalRuleTemplateAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(approvalruletemplateassociationsResource, c.ns, opts))

}

// Create takes the representation of a approvalRuleTemplateAssociation and creates it.  Returns the server's representation of the approvalRuleTemplateAssociation, and an error, if there is any.
func (c *FakeApprovalRuleTemplateAssociations) Create(ctx context.Context, approvalRuleTemplateAssociation *v1alpha1.ApprovalRuleTemplateAssociation, opts v1.CreateOptions) (result *v1alpha1.ApprovalRuleTemplateAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(approvalruletemplateassociationsResource, c.ns, approvalRuleTemplateAssociation), &v1alpha1.ApprovalRuleTemplateAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApprovalRuleTemplateAssociation), err
}

// Update takes the representation of a approvalRuleTemplateAssociation and updates it. Returns the server's representation of the approvalRuleTemplateAssociation, and an error, if there is any.
func (c *FakeApprovalRuleTemplateAssociations) Update(ctx context.Context, approvalRuleTemplateAssociation *v1alpha1.ApprovalRuleTemplateAssociation, opts v1.UpdateOptions) (result *v1alpha1.ApprovalRuleTemplateAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(approvalruletemplateassociationsResource, c.ns, approvalRuleTemplateAssociation), &v1alpha1.ApprovalRuleTemplateAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApprovalRuleTemplateAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApprovalRuleTemplateAssociations) UpdateStatus(ctx context.Context, approvalRuleTemplateAssociation *v1alpha1.ApprovalRuleTemplateAssociation, opts v1.UpdateOptions) (*v1alpha1.ApprovalRuleTemplateAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(approvalruletemplateassociationsResource, "status", c.ns, approvalRuleTemplateAssociation), &v1alpha1.ApprovalRuleTemplateAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApprovalRuleTemplateAssociation), err
}

// Delete takes name of the approvalRuleTemplateAssociation and deletes it. Returns an error if one occurs.
func (c *FakeApprovalRuleTemplateAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(approvalruletemplateassociationsResource, c.ns, name), &v1alpha1.ApprovalRuleTemplateAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApprovalRuleTemplateAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(approvalruletemplateassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApprovalRuleTemplateAssociationList{})
	return err
}

// Patch applies the patch and returns the patched approvalRuleTemplateAssociation.
func (c *FakeApprovalRuleTemplateAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApprovalRuleTemplateAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(approvalruletemplateassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApprovalRuleTemplateAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApprovalRuleTemplateAssociation), err
}
