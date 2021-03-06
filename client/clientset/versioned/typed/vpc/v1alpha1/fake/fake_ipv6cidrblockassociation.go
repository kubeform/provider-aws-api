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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/vpc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIpv6CIDRBlockAssociations implements Ipv6CIDRBlockAssociationInterface
type FakeIpv6CIDRBlockAssociations struct {
	Fake *FakeVpcV1alpha1
	ns   string
}

var ipv6cidrblockassociationsResource = schema.GroupVersionResource{Group: "vpc.aws.kubeform.com", Version: "v1alpha1", Resource: "ipv6cidrblockassociations"}

var ipv6cidrblockassociationsKind = schema.GroupVersionKind{Group: "vpc.aws.kubeform.com", Version: "v1alpha1", Kind: "Ipv6CIDRBlockAssociation"}

// Get takes name of the ipv6CIDRBlockAssociation, and returns the corresponding ipv6CIDRBlockAssociation object, and an error if there is any.
func (c *FakeIpv6CIDRBlockAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipv6cidrblockassociationsResource, c.ns, name), &v1alpha1.Ipv6CIDRBlockAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6CIDRBlockAssociation), err
}

// List takes label and field selectors, and returns the list of Ipv6CIDRBlockAssociations that match those selectors.
func (c *FakeIpv6CIDRBlockAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.Ipv6CIDRBlockAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipv6cidrblockassociationsResource, ipv6cidrblockassociationsKind, c.ns, opts), &v1alpha1.Ipv6CIDRBlockAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.Ipv6CIDRBlockAssociationList{ListMeta: obj.(*v1alpha1.Ipv6CIDRBlockAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.Ipv6CIDRBlockAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ipv6CIDRBlockAssociations.
func (c *FakeIpv6CIDRBlockAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipv6cidrblockassociationsResource, c.ns, opts))

}

// Create takes the representation of a ipv6CIDRBlockAssociation and creates it.  Returns the server's representation of the ipv6CIDRBlockAssociation, and an error, if there is any.
func (c *FakeIpv6CIDRBlockAssociations) Create(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.CreateOptions) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipv6cidrblockassociationsResource, c.ns, ipv6CIDRBlockAssociation), &v1alpha1.Ipv6CIDRBlockAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6CIDRBlockAssociation), err
}

// Update takes the representation of a ipv6CIDRBlockAssociation and updates it. Returns the server's representation of the ipv6CIDRBlockAssociation, and an error, if there is any.
func (c *FakeIpv6CIDRBlockAssociations) Update(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.UpdateOptions) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipv6cidrblockassociationsResource, c.ns, ipv6CIDRBlockAssociation), &v1alpha1.Ipv6CIDRBlockAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6CIDRBlockAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIpv6CIDRBlockAssociations) UpdateStatus(ctx context.Context, ipv6CIDRBlockAssociation *v1alpha1.Ipv6CIDRBlockAssociation, opts v1.UpdateOptions) (*v1alpha1.Ipv6CIDRBlockAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipv6cidrblockassociationsResource, "status", c.ns, ipv6CIDRBlockAssociation), &v1alpha1.Ipv6CIDRBlockAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6CIDRBlockAssociation), err
}

// Delete takes name of the ipv6CIDRBlockAssociation and deletes it. Returns an error if one occurs.
func (c *FakeIpv6CIDRBlockAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ipv6cidrblockassociationsResource, c.ns, name), &v1alpha1.Ipv6CIDRBlockAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIpv6CIDRBlockAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipv6cidrblockassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.Ipv6CIDRBlockAssociationList{})
	return err
}

// Patch applies the patch and returns the patched ipv6CIDRBlockAssociation.
func (c *FakeIpv6CIDRBlockAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ipv6CIDRBlockAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipv6cidrblockassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ipv6CIDRBlockAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipv6CIDRBlockAssociation), err
}
