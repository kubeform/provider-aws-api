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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/default/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkACLs implements NetworkACLInterface
type FakeNetworkACLs struct {
	Fake *FakeDefaultV1alpha1
	ns   string
}

var networkaclsResource = schema.GroupVersionResource{Group: "default.aws.kubeform.com", Version: "v1alpha1", Resource: "networkacls"}

var networkaclsKind = schema.GroupVersionKind{Group: "default.aws.kubeform.com", Version: "v1alpha1", Kind: "NetworkACL"}

// Get takes name of the networkACL, and returns the corresponding networkACL object, and an error if there is any.
func (c *FakeNetworkACLs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkaclsResource, c.ns, name), &v1alpha1.NetworkACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkACL), err
}

// List takes label and field selectors, and returns the list of NetworkACLs that match those selectors.
func (c *FakeNetworkACLs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkACLList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkaclsResource, networkaclsKind, c.ns, opts), &v1alpha1.NetworkACLList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkACLList{ListMeta: obj.(*v1alpha1.NetworkACLList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkACLList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkACLs.
func (c *FakeNetworkACLs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkaclsResource, c.ns, opts))

}

// Create takes the representation of a networkACL and creates it.  Returns the server's representation of the networkACL, and an error, if there is any.
func (c *FakeNetworkACLs) Create(ctx context.Context, networkACL *v1alpha1.NetworkACL, opts v1.CreateOptions) (result *v1alpha1.NetworkACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkaclsResource, c.ns, networkACL), &v1alpha1.NetworkACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkACL), err
}

// Update takes the representation of a networkACL and updates it. Returns the server's representation of the networkACL, and an error, if there is any.
func (c *FakeNetworkACLs) Update(ctx context.Context, networkACL *v1alpha1.NetworkACL, opts v1.UpdateOptions) (result *v1alpha1.NetworkACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkaclsResource, c.ns, networkACL), &v1alpha1.NetworkACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkACL), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkACLs) UpdateStatus(ctx context.Context, networkACL *v1alpha1.NetworkACL, opts v1.UpdateOptions) (*v1alpha1.NetworkACL, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkaclsResource, "status", c.ns, networkACL), &v1alpha1.NetworkACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkACL), err
}

// Delete takes name of the networkACL and deletes it. Returns an error if one occurs.
func (c *FakeNetworkACLs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkaclsResource, c.ns, name), &v1alpha1.NetworkACL{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkACLs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkaclsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkACLList{})
	return err
}

// Patch applies the patch and returns the patched networkACL.
func (c *FakeNetworkACLs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkACL, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkaclsResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkACL{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkACL), err
}
