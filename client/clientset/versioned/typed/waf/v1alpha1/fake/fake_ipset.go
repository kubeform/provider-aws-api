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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/waf/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIpsets implements IpsetInterface
type FakeIpsets struct {
	Fake *FakeWafV1alpha1
	ns   string
}

var ipsetsResource = schema.GroupVersionResource{Group: "waf.aws.kubeform.com", Version: "v1alpha1", Resource: "ipsets"}

var ipsetsKind = schema.GroupVersionKind{Group: "waf.aws.kubeform.com", Version: "v1alpha1", Kind: "Ipset"}

// Get takes name of the ipset, and returns the corresponding ipset object, and an error if there is any.
func (c *FakeIpsets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Ipset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipsetsResource, c.ns, name), &v1alpha1.Ipset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipset), err
}

// List takes label and field selectors, and returns the list of Ipsets that match those selectors.
func (c *FakeIpsets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IpsetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipsetsResource, ipsetsKind, c.ns, opts), &v1alpha1.IpsetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IpsetList{ListMeta: obj.(*v1alpha1.IpsetList).ListMeta}
	for _, item := range obj.(*v1alpha1.IpsetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ipsets.
func (c *FakeIpsets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipsetsResource, c.ns, opts))

}

// Create takes the representation of a ipset and creates it.  Returns the server's representation of the ipset, and an error, if there is any.
func (c *FakeIpsets) Create(ctx context.Context, ipset *v1alpha1.Ipset, opts v1.CreateOptions) (result *v1alpha1.Ipset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipsetsResource, c.ns, ipset), &v1alpha1.Ipset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipset), err
}

// Update takes the representation of a ipset and updates it. Returns the server's representation of the ipset, and an error, if there is any.
func (c *FakeIpsets) Update(ctx context.Context, ipset *v1alpha1.Ipset, opts v1.UpdateOptions) (result *v1alpha1.Ipset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipsetsResource, c.ns, ipset), &v1alpha1.Ipset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipset), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIpsets) UpdateStatus(ctx context.Context, ipset *v1alpha1.Ipset, opts v1.UpdateOptions) (*v1alpha1.Ipset, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipsetsResource, "status", c.ns, ipset), &v1alpha1.Ipset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipset), err
}

// Delete takes name of the ipset and deletes it. Returns an error if one occurs.
func (c *FakeIpsets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ipsetsResource, c.ns, name), &v1alpha1.Ipset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIpsets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IpsetList{})
	return err
}

// Patch applies the patch and returns the patched ipset.
func (c *FakeIpsets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Ipset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ipset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ipset), err
}
