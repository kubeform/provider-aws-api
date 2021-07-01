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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dx/v1alpha1"
)

// FakeHostedPrivateVirtualInterfaces implements HostedPrivateVirtualInterfaceInterface
type FakeHostedPrivateVirtualInterfaces struct {
	Fake *FakeDxV1alpha1
	ns   string
}

var hostedprivatevirtualinterfacesResource = schema.GroupVersionResource{Group: "dx.aws.kubeform.com", Version: "v1alpha1", Resource: "hostedprivatevirtualinterfaces"}

var hostedprivatevirtualinterfacesKind = schema.GroupVersionKind{Group: "dx.aws.kubeform.com", Version: "v1alpha1", Kind: "HostedPrivateVirtualInterface"}

// Get takes name of the hostedPrivateVirtualInterface, and returns the corresponding hostedPrivateVirtualInterface object, and an error if there is any.
func (c *FakeHostedPrivateVirtualInterfaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HostedPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hostedprivatevirtualinterfacesResource, c.ns, name), &v1alpha1.HostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostedPrivateVirtualInterface), err
}

// List takes label and field selectors, and returns the list of HostedPrivateVirtualInterfaces that match those selectors.
func (c *FakeHostedPrivateVirtualInterfaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HostedPrivateVirtualInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hostedprivatevirtualinterfacesResource, hostedprivatevirtualinterfacesKind, c.ns, opts), &v1alpha1.HostedPrivateVirtualInterfaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HostedPrivateVirtualInterfaceList{ListMeta: obj.(*v1alpha1.HostedPrivateVirtualInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.HostedPrivateVirtualInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hostedPrivateVirtualInterfaces.
func (c *FakeHostedPrivateVirtualInterfaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hostedprivatevirtualinterfacesResource, c.ns, opts))

}

// Create takes the representation of a hostedPrivateVirtualInterface and creates it.  Returns the server's representation of the hostedPrivateVirtualInterface, and an error, if there is any.
func (c *FakeHostedPrivateVirtualInterfaces) Create(ctx context.Context, hostedPrivateVirtualInterface *v1alpha1.HostedPrivateVirtualInterface, opts v1.CreateOptions) (result *v1alpha1.HostedPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hostedprivatevirtualinterfacesResource, c.ns, hostedPrivateVirtualInterface), &v1alpha1.HostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostedPrivateVirtualInterface), err
}

// Update takes the representation of a hostedPrivateVirtualInterface and updates it. Returns the server's representation of the hostedPrivateVirtualInterface, and an error, if there is any.
func (c *FakeHostedPrivateVirtualInterfaces) Update(ctx context.Context, hostedPrivateVirtualInterface *v1alpha1.HostedPrivateVirtualInterface, opts v1.UpdateOptions) (result *v1alpha1.HostedPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hostedprivatevirtualinterfacesResource, c.ns, hostedPrivateVirtualInterface), &v1alpha1.HostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostedPrivateVirtualInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHostedPrivateVirtualInterfaces) UpdateStatus(ctx context.Context, hostedPrivateVirtualInterface *v1alpha1.HostedPrivateVirtualInterface, opts v1.UpdateOptions) (*v1alpha1.HostedPrivateVirtualInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hostedprivatevirtualinterfacesResource, "status", c.ns, hostedPrivateVirtualInterface), &v1alpha1.HostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostedPrivateVirtualInterface), err
}

// Delete takes name of the hostedPrivateVirtualInterface and deletes it. Returns an error if one occurs.
func (c *FakeHostedPrivateVirtualInterfaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hostedprivatevirtualinterfacesResource, c.ns, name), &v1alpha1.HostedPrivateVirtualInterface{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHostedPrivateVirtualInterfaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hostedprivatevirtualinterfacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HostedPrivateVirtualInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched hostedPrivateVirtualInterface.
func (c *FakeHostedPrivateVirtualInterfaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HostedPrivateVirtualInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hostedprivatevirtualinterfacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.HostedPrivateVirtualInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HostedPrivateVirtualInterface), err
}
