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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/shield/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProtectionHealthCheckAssociations implements ProtectionHealthCheckAssociationInterface
type FakeProtectionHealthCheckAssociations struct {
	Fake *FakeShieldV1alpha1
	ns   string
}

var protectionhealthcheckassociationsResource = schema.GroupVersionResource{Group: "shield.aws.kubeform.com", Version: "v1alpha1", Resource: "protectionhealthcheckassociations"}

var protectionhealthcheckassociationsKind = schema.GroupVersionKind{Group: "shield.aws.kubeform.com", Version: "v1alpha1", Kind: "ProtectionHealthCheckAssociation"}

// Get takes name of the protectionHealthCheckAssociation, and returns the corresponding protectionHealthCheckAssociation object, and an error if there is any.
func (c *FakeProtectionHealthCheckAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProtectionHealthCheckAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(protectionhealthcheckassociationsResource, c.ns, name), &v1alpha1.ProtectionHealthCheckAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionHealthCheckAssociation), err
}

// List takes label and field selectors, and returns the list of ProtectionHealthCheckAssociations that match those selectors.
func (c *FakeProtectionHealthCheckAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProtectionHealthCheckAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(protectionhealthcheckassociationsResource, protectionhealthcheckassociationsKind, c.ns, opts), &v1alpha1.ProtectionHealthCheckAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProtectionHealthCheckAssociationList{ListMeta: obj.(*v1alpha1.ProtectionHealthCheckAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProtectionHealthCheckAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested protectionHealthCheckAssociations.
func (c *FakeProtectionHealthCheckAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(protectionhealthcheckassociationsResource, c.ns, opts))

}

// Create takes the representation of a protectionHealthCheckAssociation and creates it.  Returns the server's representation of the protectionHealthCheckAssociation, and an error, if there is any.
func (c *FakeProtectionHealthCheckAssociations) Create(ctx context.Context, protectionHealthCheckAssociation *v1alpha1.ProtectionHealthCheckAssociation, opts v1.CreateOptions) (result *v1alpha1.ProtectionHealthCheckAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(protectionhealthcheckassociationsResource, c.ns, protectionHealthCheckAssociation), &v1alpha1.ProtectionHealthCheckAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionHealthCheckAssociation), err
}

// Update takes the representation of a protectionHealthCheckAssociation and updates it. Returns the server's representation of the protectionHealthCheckAssociation, and an error, if there is any.
func (c *FakeProtectionHealthCheckAssociations) Update(ctx context.Context, protectionHealthCheckAssociation *v1alpha1.ProtectionHealthCheckAssociation, opts v1.UpdateOptions) (result *v1alpha1.ProtectionHealthCheckAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(protectionhealthcheckassociationsResource, c.ns, protectionHealthCheckAssociation), &v1alpha1.ProtectionHealthCheckAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionHealthCheckAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProtectionHealthCheckAssociations) UpdateStatus(ctx context.Context, protectionHealthCheckAssociation *v1alpha1.ProtectionHealthCheckAssociation, opts v1.UpdateOptions) (*v1alpha1.ProtectionHealthCheckAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(protectionhealthcheckassociationsResource, "status", c.ns, protectionHealthCheckAssociation), &v1alpha1.ProtectionHealthCheckAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionHealthCheckAssociation), err
}

// Delete takes name of the protectionHealthCheckAssociation and deletes it. Returns an error if one occurs.
func (c *FakeProtectionHealthCheckAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(protectionhealthcheckassociationsResource, c.ns, name), &v1alpha1.ProtectionHealthCheckAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProtectionHealthCheckAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(protectionhealthcheckassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProtectionHealthCheckAssociationList{})
	return err
}

// Patch applies the patch and returns the patched protectionHealthCheckAssociation.
func (c *FakeProtectionHealthCheckAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProtectionHealthCheckAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(protectionhealthcheckassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProtectionHealthCheckAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtectionHealthCheckAssociation), err
}