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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/vpn/v1alpha1"
)

// FakeConnectionRoutes implements ConnectionRouteInterface
type FakeConnectionRoutes struct {
	Fake *FakeVpnV1alpha1
	ns   string
}

var connectionroutesResource = schema.GroupVersionResource{Group: "vpn.aws.kubeform.com", Version: "v1alpha1", Resource: "connectionroutes"}

var connectionroutesKind = schema.GroupVersionKind{Group: "vpn.aws.kubeform.com", Version: "v1alpha1", Kind: "ConnectionRoute"}

// Get takes name of the connectionRoute, and returns the corresponding connectionRoute object, and an error if there is any.
func (c *FakeConnectionRoutes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConnectionRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(connectionroutesResource, c.ns, name), &v1alpha1.ConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectionRoute), err
}

// List takes label and field selectors, and returns the list of ConnectionRoutes that match those selectors.
func (c *FakeConnectionRoutes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConnectionRouteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(connectionroutesResource, connectionroutesKind, c.ns, opts), &v1alpha1.ConnectionRouteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConnectionRouteList{ListMeta: obj.(*v1alpha1.ConnectionRouteList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConnectionRouteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested connectionRoutes.
func (c *FakeConnectionRoutes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(connectionroutesResource, c.ns, opts))

}

// Create takes the representation of a connectionRoute and creates it.  Returns the server's representation of the connectionRoute, and an error, if there is any.
func (c *FakeConnectionRoutes) Create(ctx context.Context, connectionRoute *v1alpha1.ConnectionRoute, opts v1.CreateOptions) (result *v1alpha1.ConnectionRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(connectionroutesResource, c.ns, connectionRoute), &v1alpha1.ConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectionRoute), err
}

// Update takes the representation of a connectionRoute and updates it. Returns the server's representation of the connectionRoute, and an error, if there is any.
func (c *FakeConnectionRoutes) Update(ctx context.Context, connectionRoute *v1alpha1.ConnectionRoute, opts v1.UpdateOptions) (result *v1alpha1.ConnectionRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(connectionroutesResource, c.ns, connectionRoute), &v1alpha1.ConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectionRoute), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConnectionRoutes) UpdateStatus(ctx context.Context, connectionRoute *v1alpha1.ConnectionRoute, opts v1.UpdateOptions) (*v1alpha1.ConnectionRoute, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(connectionroutesResource, "status", c.ns, connectionRoute), &v1alpha1.ConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectionRoute), err
}

// Delete takes name of the connectionRoute and deletes it. Returns an error if one occurs.
func (c *FakeConnectionRoutes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(connectionroutesResource, c.ns, name), &v1alpha1.ConnectionRoute{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConnectionRoutes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(connectionroutesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConnectionRouteList{})
	return err
}

// Patch applies the patch and returns the patched connectionRoute.
func (c *FakeConnectionRoutes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConnectionRoute, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(connectionroutesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConnectionRoute{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConnectionRoute), err
}
