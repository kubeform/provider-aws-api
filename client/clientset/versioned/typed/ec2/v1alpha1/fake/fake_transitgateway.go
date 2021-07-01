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

// FakeTransitGateways implements TransitGatewayInterface
type FakeTransitGateways struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var transitgatewaysResource = schema.GroupVersionResource{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Resource: "transitgateways"}

var transitgatewaysKind = schema.GroupVersionKind{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Kind: "TransitGateway"}

// Get takes name of the transitGateway, and returns the corresponding transitGateway object, and an error if there is any.
func (c *FakeTransitGateways) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TransitGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transitgatewaysResource, c.ns, name), &v1alpha1.TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGateway), err
}

// List takes label and field selectors, and returns the list of TransitGateways that match those selectors.
func (c *FakeTransitGateways) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransitGatewayList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transitgatewaysResource, transitgatewaysKind, c.ns, opts), &v1alpha1.TransitGatewayList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransitGatewayList{ListMeta: obj.(*v1alpha1.TransitGatewayList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransitGatewayList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transitGateways.
func (c *FakeTransitGateways) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transitgatewaysResource, c.ns, opts))

}

// Create takes the representation of a transitGateway and creates it.  Returns the server's representation of the transitGateway, and an error, if there is any.
func (c *FakeTransitGateways) Create(ctx context.Context, transitGateway *v1alpha1.TransitGateway, opts v1.CreateOptions) (result *v1alpha1.TransitGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transitgatewaysResource, c.ns, transitGateway), &v1alpha1.TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGateway), err
}

// Update takes the representation of a transitGateway and updates it. Returns the server's representation of the transitGateway, and an error, if there is any.
func (c *FakeTransitGateways) Update(ctx context.Context, transitGateway *v1alpha1.TransitGateway, opts v1.UpdateOptions) (result *v1alpha1.TransitGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transitgatewaysResource, c.ns, transitGateway), &v1alpha1.TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGateway), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransitGateways) UpdateStatus(ctx context.Context, transitGateway *v1alpha1.TransitGateway, opts v1.UpdateOptions) (*v1alpha1.TransitGateway, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transitgatewaysResource, "status", c.ns, transitGateway), &v1alpha1.TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGateway), err
}

// Delete takes name of the transitGateway and deletes it. Returns an error if one occurs.
func (c *FakeTransitGateways) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transitgatewaysResource, c.ns, name), &v1alpha1.TransitGateway{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransitGateways) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transitgatewaysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransitGatewayList{})
	return err
}

// Patch applies the patch and returns the patched transitGateway.
func (c *FakeTransitGateways) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitGateway, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transitgatewaysResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransitGateway{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGateway), err
}
