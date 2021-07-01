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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/route53/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVpcAssociationAuthorizations implements VpcAssociationAuthorizationInterface
type FakeVpcAssociationAuthorizations struct {
	Fake *FakeRoute53V1alpha1
	ns   string
}

var vpcassociationauthorizationsResource = schema.GroupVersionResource{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Resource: "vpcassociationauthorizations"}

var vpcassociationauthorizationsKind = schema.GroupVersionKind{Group: "route53.aws.kubeform.com", Version: "v1alpha1", Kind: "VpcAssociationAuthorization"}

// Get takes name of the vpcAssociationAuthorization, and returns the corresponding vpcAssociationAuthorization object, and an error if there is any.
func (c *FakeVpcAssociationAuthorizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VpcAssociationAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcassociationauthorizationsResource, c.ns, name), &v1alpha1.VpcAssociationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcAssociationAuthorization), err
}

// List takes label and field selectors, and returns the list of VpcAssociationAuthorizations that match those selectors.
func (c *FakeVpcAssociationAuthorizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VpcAssociationAuthorizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcassociationauthorizationsResource, vpcassociationauthorizationsKind, c.ns, opts), &v1alpha1.VpcAssociationAuthorizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VpcAssociationAuthorizationList{ListMeta: obj.(*v1alpha1.VpcAssociationAuthorizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.VpcAssociationAuthorizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vpcAssociationAuthorizations.
func (c *FakeVpcAssociationAuthorizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcassociationauthorizationsResource, c.ns, opts))

}

// Create takes the representation of a vpcAssociationAuthorization and creates it.  Returns the server's representation of the vpcAssociationAuthorization, and an error, if there is any.
func (c *FakeVpcAssociationAuthorizations) Create(ctx context.Context, vpcAssociationAuthorization *v1alpha1.VpcAssociationAuthorization, opts v1.CreateOptions) (result *v1alpha1.VpcAssociationAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcassociationauthorizationsResource, c.ns, vpcAssociationAuthorization), &v1alpha1.VpcAssociationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcAssociationAuthorization), err
}

// Update takes the representation of a vpcAssociationAuthorization and updates it. Returns the server's representation of the vpcAssociationAuthorization, and an error, if there is any.
func (c *FakeVpcAssociationAuthorizations) Update(ctx context.Context, vpcAssociationAuthorization *v1alpha1.VpcAssociationAuthorization, opts v1.UpdateOptions) (result *v1alpha1.VpcAssociationAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcassociationauthorizationsResource, c.ns, vpcAssociationAuthorization), &v1alpha1.VpcAssociationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcAssociationAuthorization), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVpcAssociationAuthorizations) UpdateStatus(ctx context.Context, vpcAssociationAuthorization *v1alpha1.VpcAssociationAuthorization, opts v1.UpdateOptions) (*v1alpha1.VpcAssociationAuthorization, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcassociationauthorizationsResource, "status", c.ns, vpcAssociationAuthorization), &v1alpha1.VpcAssociationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcAssociationAuthorization), err
}

// Delete takes name of the vpcAssociationAuthorization and deletes it. Returns an error if one occurs.
func (c *FakeVpcAssociationAuthorizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcassociationauthorizationsResource, c.ns, name), &v1alpha1.VpcAssociationAuthorization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVpcAssociationAuthorizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcassociationauthorizationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.VpcAssociationAuthorizationList{})
	return err
}

// Patch applies the patch and returns the patched vpcAssociationAuthorization.
func (c *FakeVpcAssociationAuthorizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VpcAssociationAuthorization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcassociationauthorizationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VpcAssociationAuthorization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VpcAssociationAuthorization), err
}
