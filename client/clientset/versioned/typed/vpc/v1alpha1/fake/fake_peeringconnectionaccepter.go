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

// FakePeeringConnectionAccepters implements PeeringConnectionAccepterInterface
type FakePeeringConnectionAccepters struct {
	Fake *FakeVpcV1alpha1
	ns   string
}

var peeringconnectionacceptersResource = schema.GroupVersionResource{Group: "vpc.aws.kubeform.com", Version: "v1alpha1", Resource: "peeringconnectionaccepters"}

var peeringconnectionacceptersKind = schema.GroupVersionKind{Group: "vpc.aws.kubeform.com", Version: "v1alpha1", Kind: "PeeringConnectionAccepter"}

// Get takes name of the peeringConnectionAccepter, and returns the corresponding peeringConnectionAccepter object, and an error if there is any.
func (c *FakePeeringConnectionAccepters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PeeringConnectionAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(peeringconnectionacceptersResource, c.ns, name), &v1alpha1.PeeringConnectionAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PeeringConnectionAccepter), err
}

// List takes label and field selectors, and returns the list of PeeringConnectionAccepters that match those selectors.
func (c *FakePeeringConnectionAccepters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PeeringConnectionAccepterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(peeringconnectionacceptersResource, peeringconnectionacceptersKind, c.ns, opts), &v1alpha1.PeeringConnectionAccepterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PeeringConnectionAccepterList{ListMeta: obj.(*v1alpha1.PeeringConnectionAccepterList).ListMeta}
	for _, item := range obj.(*v1alpha1.PeeringConnectionAccepterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested peeringConnectionAccepters.
func (c *FakePeeringConnectionAccepters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(peeringconnectionacceptersResource, c.ns, opts))

}

// Create takes the representation of a peeringConnectionAccepter and creates it.  Returns the server's representation of the peeringConnectionAccepter, and an error, if there is any.
func (c *FakePeeringConnectionAccepters) Create(ctx context.Context, peeringConnectionAccepter *v1alpha1.PeeringConnectionAccepter, opts v1.CreateOptions) (result *v1alpha1.PeeringConnectionAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(peeringconnectionacceptersResource, c.ns, peeringConnectionAccepter), &v1alpha1.PeeringConnectionAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PeeringConnectionAccepter), err
}

// Update takes the representation of a peeringConnectionAccepter and updates it. Returns the server's representation of the peeringConnectionAccepter, and an error, if there is any.
func (c *FakePeeringConnectionAccepters) Update(ctx context.Context, peeringConnectionAccepter *v1alpha1.PeeringConnectionAccepter, opts v1.UpdateOptions) (result *v1alpha1.PeeringConnectionAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(peeringconnectionacceptersResource, c.ns, peeringConnectionAccepter), &v1alpha1.PeeringConnectionAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PeeringConnectionAccepter), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePeeringConnectionAccepters) UpdateStatus(ctx context.Context, peeringConnectionAccepter *v1alpha1.PeeringConnectionAccepter, opts v1.UpdateOptions) (*v1alpha1.PeeringConnectionAccepter, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(peeringconnectionacceptersResource, "status", c.ns, peeringConnectionAccepter), &v1alpha1.PeeringConnectionAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PeeringConnectionAccepter), err
}

// Delete takes name of the peeringConnectionAccepter and deletes it. Returns an error if one occurs.
func (c *FakePeeringConnectionAccepters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(peeringconnectionacceptersResource, c.ns, name), &v1alpha1.PeeringConnectionAccepter{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePeeringConnectionAccepters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(peeringconnectionacceptersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PeeringConnectionAccepterList{})
	return err
}

// Patch applies the patch and returns the patched peeringConnectionAccepter.
func (c *FakePeeringConnectionAccepters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PeeringConnectionAccepter, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(peeringconnectionacceptersResource, c.ns, name, pt, data, subresources...), &v1alpha1.PeeringConnectionAccepter{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PeeringConnectionAccepter), err
}
