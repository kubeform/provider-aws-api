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

// FakeIpamPoolCIDRs implements IpamPoolCIDRInterface
type FakeIpamPoolCIDRs struct {
	Fake *FakeVpcV1alpha1
	ns   string
}

var ipampoolcidrsResource = schema.GroupVersionResource{Group: "vpc.aws.kubeform.com", Version: "v1alpha1", Resource: "ipampoolcidrs"}

var ipampoolcidrsKind = schema.GroupVersionKind{Group: "vpc.aws.kubeform.com", Version: "v1alpha1", Kind: "IpamPoolCIDR"}

// Get takes name of the ipamPoolCIDR, and returns the corresponding ipamPoolCIDR object, and an error if there is any.
func (c *FakeIpamPoolCIDRs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IpamPoolCIDR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipampoolcidrsResource, c.ns, name), &v1alpha1.IpamPoolCIDR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpamPoolCIDR), err
}

// List takes label and field selectors, and returns the list of IpamPoolCIDRs that match those selectors.
func (c *FakeIpamPoolCIDRs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IpamPoolCIDRList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipampoolcidrsResource, ipampoolcidrsKind, c.ns, opts), &v1alpha1.IpamPoolCIDRList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IpamPoolCIDRList{ListMeta: obj.(*v1alpha1.IpamPoolCIDRList).ListMeta}
	for _, item := range obj.(*v1alpha1.IpamPoolCIDRList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ipamPoolCIDRs.
func (c *FakeIpamPoolCIDRs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipampoolcidrsResource, c.ns, opts))

}

// Create takes the representation of a ipamPoolCIDR and creates it.  Returns the server's representation of the ipamPoolCIDR, and an error, if there is any.
func (c *FakeIpamPoolCIDRs) Create(ctx context.Context, ipamPoolCIDR *v1alpha1.IpamPoolCIDR, opts v1.CreateOptions) (result *v1alpha1.IpamPoolCIDR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipampoolcidrsResource, c.ns, ipamPoolCIDR), &v1alpha1.IpamPoolCIDR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpamPoolCIDR), err
}

// Update takes the representation of a ipamPoolCIDR and updates it. Returns the server's representation of the ipamPoolCIDR, and an error, if there is any.
func (c *FakeIpamPoolCIDRs) Update(ctx context.Context, ipamPoolCIDR *v1alpha1.IpamPoolCIDR, opts v1.UpdateOptions) (result *v1alpha1.IpamPoolCIDR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipampoolcidrsResource, c.ns, ipamPoolCIDR), &v1alpha1.IpamPoolCIDR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpamPoolCIDR), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIpamPoolCIDRs) UpdateStatus(ctx context.Context, ipamPoolCIDR *v1alpha1.IpamPoolCIDR, opts v1.UpdateOptions) (*v1alpha1.IpamPoolCIDR, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipampoolcidrsResource, "status", c.ns, ipamPoolCIDR), &v1alpha1.IpamPoolCIDR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpamPoolCIDR), err
}

// Delete takes name of the ipamPoolCIDR and deletes it. Returns an error if one occurs.
func (c *FakeIpamPoolCIDRs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ipampoolcidrsResource, c.ns, name), &v1alpha1.IpamPoolCIDR{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIpamPoolCIDRs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipampoolcidrsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IpamPoolCIDRList{})
	return err
}

// Patch applies the patch and returns the patched ipamPoolCIDR.
func (c *FakeIpamPoolCIDRs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IpamPoolCIDR, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipampoolcidrsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IpamPoolCIDR{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IpamPoolCIDR), err
}
