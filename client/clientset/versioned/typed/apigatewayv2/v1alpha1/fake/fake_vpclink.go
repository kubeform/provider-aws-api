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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/apigatewayv2/v1alpha1"
)

// FakeVpcLinks implements VpcLinkInterface
type FakeVpcLinks struct {
	Fake *FakeApigatewayv2V1alpha1
	ns   string
}

var vpclinksResource = schema.GroupVersionResource{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Resource: "vpclinks"}

var vpclinksKind = schema.GroupVersionKind{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Kind: "VpcLink"}

// Get takes name of the vpcLink, and returns the corresponding vpcLink object, and an error if there is any.
func (c *FakeVpcLinks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpcLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpclinksResource, c.ns, name), &v1alpha1.VpcLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcLink), err
}

// List takes label and field selectors, and returns the list of VpcLinks that match those selectors.
func (c *FakeVpcLinks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpcLinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpclinksResource, vpclinksKind, c.ns, opts), &v1alpha1.VpcLinkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcLinkList{ListMeta: obj.(*v1alpha1.VpcLinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcLinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcLinks.
func (c *FakeVpcLinks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpclinksResource, c.ns, opts))

}

// Create takes the representation of a vpcLink and creates it.  Returns the server's representation of the vpcLink, and an error, if there is any.
func (c *FakeVpcLinks) Create(ctx context.Context, vpcLink *v1alpha1.VpcLink, opts v1.CreateOptions) (result *v1alpha1.VpcLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpclinksResource, c.ns, vpcLink), &v1alpha1.VpcLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcLink), err
}

// Update takes the representation of a vpcLink and updates it. Returns the server's representation of the vpcLink, and an error, if there is any.
func (c *FakeVpcLinks) Update(ctx context.Context, vpcLink *v1alpha1.VpcLink, opts v1.UpdateOptions) (result *v1alpha1.VpcLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpclinksResource, c.ns, vpcLink), &v1alpha1.VpcLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcLink), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcLinks) UpdateStatus(ctx context.Context, vpcLink *v1alpha1.VpcLink, opts v1.UpdateOptions) (*v1alpha1.VpcLink, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpclinksResource, "status", c.ns, vpcLink), &v1alpha1.VpcLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcLink), err
}

// Delete takes name of the vpcLink and deletes it. Returns an error if one occurs.
func (c *FakeVpcLinks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpclinksResource, c.ns, name), &v1alpha1.VpcLink{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcLinks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpclinksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcLinkList{})
	return err
}

// Patch applies the patch and returns the patched vpcLink.
func (c *FakeVpcLinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpcLink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpclinksResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcLink{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcLink), err
}
