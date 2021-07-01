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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"
)

// FakeClientVPNNetworkAssociations implements ClientVPNNetworkAssociationInterface
type FakeClientVPNNetworkAssociations struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var clientvpnnetworkassociationsResource = schema.GroupVersionResource{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Resource: "clientvpnnetworkassociations"}

var clientvpnnetworkassociationsKind = schema.GroupVersionKind{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Kind: "ClientVPNNetworkAssociation"}

// Get takes name of the clientVPNNetworkAssociation, and returns the corresponding clientVPNNetworkAssociation object, and an error if there is any.
func (c *FakeClientVPNNetworkAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClientVPNNetworkAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clientvpnnetworkassociationsResource, c.ns, name), &v1alpha1.ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClientVPNNetworkAssociation), err
}

// List takes label and field selectors, and returns the list of ClientVPNNetworkAssociations that match those selectors.
func (c *FakeClientVPNNetworkAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClientVPNNetworkAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clientvpnnetworkassociationsResource, clientvpnnetworkassociationsKind, c.ns, opts), &v1alpha1.ClientVPNNetworkAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClientVPNNetworkAssociationList{ListMeta: obj.(*v1alpha1.ClientVPNNetworkAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClientVPNNetworkAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clientVPNNetworkAssociations.
func (c *FakeClientVPNNetworkAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clientvpnnetworkassociationsResource, c.ns, opts))

}

// Create takes the representation of a clientVPNNetworkAssociation and creates it.  Returns the server's representation of the clientVPNNetworkAssociation, and an error, if there is any.
func (c *FakeClientVPNNetworkAssociations) Create(ctx context.Context, clientVPNNetworkAssociation *v1alpha1.ClientVPNNetworkAssociation, opts v1.CreateOptions) (result *v1alpha1.ClientVPNNetworkAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clientvpnnetworkassociationsResource, c.ns, clientVPNNetworkAssociation), &v1alpha1.ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClientVPNNetworkAssociation), err
}

// Update takes the representation of a clientVPNNetworkAssociation and updates it. Returns the server's representation of the clientVPNNetworkAssociation, and an error, if there is any.
func (c *FakeClientVPNNetworkAssociations) Update(ctx context.Context, clientVPNNetworkAssociation *v1alpha1.ClientVPNNetworkAssociation, opts v1.UpdateOptions) (result *v1alpha1.ClientVPNNetworkAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clientvpnnetworkassociationsResource, c.ns, clientVPNNetworkAssociation), &v1alpha1.ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClientVPNNetworkAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClientVPNNetworkAssociations) UpdateStatus(ctx context.Context, clientVPNNetworkAssociation *v1alpha1.ClientVPNNetworkAssociation, opts v1.UpdateOptions) (*v1alpha1.ClientVPNNetworkAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clientvpnnetworkassociationsResource, "status", c.ns, clientVPNNetworkAssociation), &v1alpha1.ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClientVPNNetworkAssociation), err
}

// Delete takes name of the clientVPNNetworkAssociation and deletes it. Returns an error if one occurs.
func (c *FakeClientVPNNetworkAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clientvpnnetworkassociationsResource, c.ns, name), &v1alpha1.ClientVPNNetworkAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClientVPNNetworkAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clientvpnnetworkassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClientVPNNetworkAssociationList{})
	return err
}

// Patch applies the patch and returns the patched clientVPNNetworkAssociation.
func (c *FakeClientVPNNetworkAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClientVPNNetworkAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clientvpnnetworkassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClientVPNNetworkAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClientVPNNetworkAssociation), err
}
