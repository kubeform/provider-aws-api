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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/vpn/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGatewayRoutePropagations implements GatewayRoutePropagationInterface
type FakeGatewayRoutePropagations struct {
	Fake *FakeVpnV1alpha1
	ns   string
}

var gatewayroutepropagationsResource = schema.GroupVersionResource{Group: "vpn.aws.kubeform.com", Version: "v1alpha1", Resource: "gatewayroutepropagations"}

var gatewayroutepropagationsKind = schema.GroupVersionKind{Group: "vpn.aws.kubeform.com", Version: "v1alpha1", Kind: "GatewayRoutePropagation"}

// Get takes name of the gatewayRoutePropagation, and returns the corresponding gatewayRoutePropagation object, and an error if there is any.
func (c *FakeGatewayRoutePropagations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GatewayRoutePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(gatewayroutepropagationsResource, c.ns, name), &v1alpha1.GatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayRoutePropagation), err
}

// List takes label and field selectors, and returns the list of GatewayRoutePropagations that match those selectors.
func (c *FakeGatewayRoutePropagations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GatewayRoutePropagationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(gatewayroutepropagationsResource, gatewayroutepropagationsKind, c.ns, opts), &v1alpha1.GatewayRoutePropagationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.GatewayRoutePropagationList{ListMeta: obj.(*v1alpha1.GatewayRoutePropagationList).ListMeta}
	for _, item := range obj.(*v1alpha1.GatewayRoutePropagationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gatewayRoutePropagations.
func (c *FakeGatewayRoutePropagations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(gatewayroutepropagationsResource, c.ns, opts))

}

// Create takes the representation of a gatewayRoutePropagation and creates it.  Returns the server's representation of the gatewayRoutePropagation, and an error, if there is any.
func (c *FakeGatewayRoutePropagations) Create(ctx context.Context, gatewayRoutePropagation *v1alpha1.GatewayRoutePropagation, opts v1.CreateOptions) (result *v1alpha1.GatewayRoutePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(gatewayroutepropagationsResource, c.ns, gatewayRoutePropagation), &v1alpha1.GatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayRoutePropagation), err
}

// Update takes the representation of a gatewayRoutePropagation and updates it. Returns the server's representation of the gatewayRoutePropagation, and an error, if there is any.
func (c *FakeGatewayRoutePropagations) Update(ctx context.Context, gatewayRoutePropagation *v1alpha1.GatewayRoutePropagation, opts v1.UpdateOptions) (result *v1alpha1.GatewayRoutePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(gatewayroutepropagationsResource, c.ns, gatewayRoutePropagation), &v1alpha1.GatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayRoutePropagation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGatewayRoutePropagations) UpdateStatus(ctx context.Context, gatewayRoutePropagation *v1alpha1.GatewayRoutePropagation, opts v1.UpdateOptions) (*v1alpha1.GatewayRoutePropagation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(gatewayroutepropagationsResource, "status", c.ns, gatewayRoutePropagation), &v1alpha1.GatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayRoutePropagation), err
}

// Delete takes name of the gatewayRoutePropagation and deletes it. Returns an error if one occurs.
func (c *FakeGatewayRoutePropagations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(gatewayroutepropagationsResource, c.ns, name), &v1alpha1.GatewayRoutePropagation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGatewayRoutePropagations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(gatewayroutepropagationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.GatewayRoutePropagationList{})
	return err
}

// Patch applies the patch and returns the patched gatewayRoutePropagation.
func (c *FakeGatewayRoutePropagations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GatewayRoutePropagation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(gatewayroutepropagationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.GatewayRoutePropagation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.GatewayRoutePropagation), err
}
