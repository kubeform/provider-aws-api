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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/network/v1alpha1"
)

// FakeAcls implements AclInterface
type FakeAcls struct {
	Fake *FakeNetworkV1alpha1
	ns   string
}

var aclsResource = schema.GroupVersionResource{Group: "network.aws.kubeform.com", Version: "v1alpha1", Resource: "acls"}

var aclsKind = schema.GroupVersionKind{Group: "network.aws.kubeform.com", Version: "v1alpha1", Kind: "Acl"}

// Get takes name of the acl, and returns the corresponding acl object, and an error if there is any.
func (c *FakeAcls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Acl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(aclsResource, c.ns, name), &v1alpha1.Acl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Acl), err
}

// List takes label and field selectors, and returns the list of Acls that match those selectors.
func (c *FakeAcls) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AclList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(aclsResource, aclsKind, c.ns, opts), &v1alpha1.AclList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AclList{ListMeta: obj.(*v1alpha1.AclList).ListMeta}
	for _, item := range obj.(*v1alpha1.AclList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested acls.
func (c *FakeAcls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(aclsResource, c.ns, opts))

}

// Create takes the representation of a acl and creates it.  Returns the server's representation of the acl, and an error, if there is any.
func (c *FakeAcls) Create(ctx context.Context, acl *v1alpha1.Acl, opts v1.CreateOptions) (result *v1alpha1.Acl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(aclsResource, c.ns, acl), &v1alpha1.Acl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Acl), err
}

// Update takes the representation of a acl and updates it. Returns the server's representation of the acl, and an error, if there is any.
func (c *FakeAcls) Update(ctx context.Context, acl *v1alpha1.Acl, opts v1.UpdateOptions) (result *v1alpha1.Acl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(aclsResource, c.ns, acl), &v1alpha1.Acl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Acl), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAcls) UpdateStatus(ctx context.Context, acl *v1alpha1.Acl, opts v1.UpdateOptions) (*v1alpha1.Acl, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(aclsResource, "status", c.ns, acl), &v1alpha1.Acl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Acl), err
}

// Delete takes name of the acl and deletes it. Returns an error if one occurs.
func (c *FakeAcls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(aclsResource, c.ns, name), &v1alpha1.Acl{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAcls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(aclsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AclList{})
	return err
}

// Patch applies the patch and returns the patched acl.
func (c *FakeAcls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Acl, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(aclsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Acl{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Acl), err
}
