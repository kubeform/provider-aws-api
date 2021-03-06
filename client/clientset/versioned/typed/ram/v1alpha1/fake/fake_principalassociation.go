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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ram/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrincipalAssociations implements PrincipalAssociationInterface
type FakePrincipalAssociations struct {
	Fake *FakeRamV1alpha1
	ns   string
}

var principalassociationsResource = schema.GroupVersionResource{Group: "ram.aws.kubeform.com", Version: "v1alpha1", Resource: "principalassociations"}

var principalassociationsKind = schema.GroupVersionKind{Group: "ram.aws.kubeform.com", Version: "v1alpha1", Kind: "PrincipalAssociation"}

// Get takes name of the principalAssociation, and returns the corresponding principalAssociation object, and an error if there is any.
func (c *FakePrincipalAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PrincipalAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(principalassociationsResource, c.ns, name), &v1alpha1.PrincipalAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrincipalAssociation), err
}

// List takes label and field selectors, and returns the list of PrincipalAssociations that match those selectors.
func (c *FakePrincipalAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PrincipalAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(principalassociationsResource, principalassociationsKind, c.ns, opts), &v1alpha1.PrincipalAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PrincipalAssociationList{ListMeta: obj.(*v1alpha1.PrincipalAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.PrincipalAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested principalAssociations.
func (c *FakePrincipalAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(principalassociationsResource, c.ns, opts))

}

// Create takes the representation of a principalAssociation and creates it.  Returns the server's representation of the principalAssociation, and an error, if there is any.
func (c *FakePrincipalAssociations) Create(ctx context.Context, principalAssociation *v1alpha1.PrincipalAssociation, opts v1.CreateOptions) (result *v1alpha1.PrincipalAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(principalassociationsResource, c.ns, principalAssociation), &v1alpha1.PrincipalAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrincipalAssociation), err
}

// Update takes the representation of a principalAssociation and updates it. Returns the server's representation of the principalAssociation, and an error, if there is any.
func (c *FakePrincipalAssociations) Update(ctx context.Context, principalAssociation *v1alpha1.PrincipalAssociation, opts v1.UpdateOptions) (result *v1alpha1.PrincipalAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(principalassociationsResource, c.ns, principalAssociation), &v1alpha1.PrincipalAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrincipalAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrincipalAssociations) UpdateStatus(ctx context.Context, principalAssociation *v1alpha1.PrincipalAssociation, opts v1.UpdateOptions) (*v1alpha1.PrincipalAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(principalassociationsResource, "status", c.ns, principalAssociation), &v1alpha1.PrincipalAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrincipalAssociation), err
}

// Delete takes name of the principalAssociation and deletes it. Returns an error if one occurs.
func (c *FakePrincipalAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(principalassociationsResource, c.ns, name), &v1alpha1.PrincipalAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrincipalAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(principalassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PrincipalAssociationList{})
	return err
}

// Patch applies the patch and returns the patched principalAssociation.
func (c *FakePrincipalAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PrincipalAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(principalassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.PrincipalAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PrincipalAssociation), err
}
