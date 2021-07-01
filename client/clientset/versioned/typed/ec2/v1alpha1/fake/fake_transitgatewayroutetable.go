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

// FakeTransitGatewayRouteTables implements TransitGatewayRouteTableInterface
type FakeTransitGatewayRouteTables struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var transitgatewayroutetablesResource = schema.GroupVersionResource{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Resource: "transitgatewayroutetables"}

var transitgatewayroutetablesKind = schema.GroupVersionKind{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Kind: "TransitGatewayRouteTable"}

// Get takes name of the transitGatewayRouteTable, and returns the corresponding transitGatewayRouteTable object, and an error if there is any.
func (c *FakeTransitGatewayRouteTables) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(transitgatewayroutetablesResource, c.ns, name), &v1alpha1.TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTable), err
}

// List takes label and field selectors, and returns the list of TransitGatewayRouteTables that match those selectors.
func (c *FakeTransitGatewayRouteTables) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransitGatewayRouteTableList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(transitgatewayroutetablesResource, transitgatewayroutetablesKind, c.ns, opts), &v1alpha1.TransitGatewayRouteTableList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TransitGatewayRouteTableList{ListMeta: obj.(*v1alpha1.TransitGatewayRouteTableList).ListMeta}
	for _, item := range obj.(*v1alpha1.TransitGatewayRouteTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested transitGatewayRouteTables.
func (c *FakeTransitGatewayRouteTables) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(transitgatewayroutetablesResource, c.ns, opts))

}

// Create takes the representation of a transitGatewayRouteTable and creates it.  Returns the server's representation of the transitGatewayRouteTable, and an error, if there is any.
func (c *FakeTransitGatewayRouteTables) Create(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.CreateOptions) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(transitgatewayroutetablesResource, c.ns, transitGatewayRouteTable), &v1alpha1.TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTable), err
}

// Update takes the representation of a transitGatewayRouteTable and updates it. Returns the server's representation of the transitGatewayRouteTable, and an error, if there is any.
func (c *FakeTransitGatewayRouteTables) Update(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.UpdateOptions) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(transitgatewayroutetablesResource, c.ns, transitGatewayRouteTable), &v1alpha1.TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTransitGatewayRouteTables) UpdateStatus(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.UpdateOptions) (*v1alpha1.TransitGatewayRouteTable, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(transitgatewayroutetablesResource, "status", c.ns, transitGatewayRouteTable), &v1alpha1.TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTable), err
}

// Delete takes name of the transitGatewayRouteTable and deletes it. Returns an error if one occurs.
func (c *FakeTransitGatewayRouteTables) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(transitgatewayroutetablesResource, c.ns, name), &v1alpha1.TransitGatewayRouteTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTransitGatewayRouteTables) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(transitgatewayroutetablesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TransitGatewayRouteTableList{})
	return err
}

// Patch applies the patch and returns the patched transitGatewayRouteTable.
func (c *FakeTransitGatewayRouteTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(transitgatewayroutetablesResource, c.ns, name, pt, data, subresources...), &v1alpha1.TransitGatewayRouteTable{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TransitGatewayRouteTable), err
}
