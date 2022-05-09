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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/appstream/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUserStackAssociations implements UserStackAssociationInterface
type FakeUserStackAssociations struct {
	Fake *FakeAppstreamV1alpha1
	ns   string
}

var userstackassociationsResource = schema.GroupVersionResource{Group: "appstream.aws.kubeform.com", Version: "v1alpha1", Resource: "userstackassociations"}

var userstackassociationsKind = schema.GroupVersionKind{Group: "appstream.aws.kubeform.com", Version: "v1alpha1", Kind: "UserStackAssociation"}

// Get takes name of the userStackAssociation, and returns the corresponding userStackAssociation object, and an error if there is any.
func (c *FakeUserStackAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserStackAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(userstackassociationsResource, c.ns, name), &v1alpha1.UserStackAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserStackAssociation), err
}

// List takes label and field selectors, and returns the list of UserStackAssociations that match those selectors.
func (c *FakeUserStackAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserStackAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(userstackassociationsResource, userstackassociationsKind, c.ns, opts), &v1alpha1.UserStackAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UserStackAssociationList{ListMeta: obj.(*v1alpha1.UserStackAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.UserStackAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userStackAssociations.
func (c *FakeUserStackAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(userstackassociationsResource, c.ns, opts))

}

// Create takes the representation of a userStackAssociation and creates it.  Returns the server's representation of the userStackAssociation, and an error, if there is any.
func (c *FakeUserStackAssociations) Create(ctx context.Context, userStackAssociation *v1alpha1.UserStackAssociation, opts v1.CreateOptions) (result *v1alpha1.UserStackAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(userstackassociationsResource, c.ns, userStackAssociation), &v1alpha1.UserStackAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserStackAssociation), err
}

// Update takes the representation of a userStackAssociation and updates it. Returns the server's representation of the userStackAssociation, and an error, if there is any.
func (c *FakeUserStackAssociations) Update(ctx context.Context, userStackAssociation *v1alpha1.UserStackAssociation, opts v1.UpdateOptions) (result *v1alpha1.UserStackAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(userstackassociationsResource, c.ns, userStackAssociation), &v1alpha1.UserStackAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserStackAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserStackAssociations) UpdateStatus(ctx context.Context, userStackAssociation *v1alpha1.UserStackAssociation, opts v1.UpdateOptions) (*v1alpha1.UserStackAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(userstackassociationsResource, "status", c.ns, userStackAssociation), &v1alpha1.UserStackAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserStackAssociation), err
}

// Delete takes name of the userStackAssociation and deletes it. Returns an error if one occurs.
func (c *FakeUserStackAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(userstackassociationsResource, c.ns, name), &v1alpha1.UserStackAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserStackAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(userstackassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UserStackAssociationList{})
	return err
}

// Patch applies the patch and returns the patched userStackAssociation.
func (c *FakeUserStackAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserStackAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(userstackassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.UserStackAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserStackAssociation), err
}