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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/iot/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRoleAliases implements RoleAliasInterface
type FakeRoleAliases struct {
	Fake *FakeIotV1alpha1
	ns   string
}

var rolealiasesResource = schema.GroupVersionResource{Group: "iot.aws.kubeform.com", Version: "v1alpha1", Resource: "rolealiases"}

var rolealiasesKind = schema.GroupVersionKind{Group: "iot.aws.kubeform.com", Version: "v1alpha1", Kind: "RoleAlias"}

// Get takes name of the roleAlias, and returns the corresponding roleAlias object, and an error if there is any.
func (c *FakeRoleAliases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RoleAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rolealiasesResource, c.ns, name), &v1alpha1.RoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAlias), err
}

// List takes label and field selectors, and returns the list of RoleAliases that match those selectors.
func (c *FakeRoleAliases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RoleAliasList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rolealiasesResource, rolealiasesKind, c.ns, opts), &v1alpha1.RoleAliasList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RoleAliasList{ListMeta: obj.(*v1alpha1.RoleAliasList).ListMeta}
	for _, item := range obj.(*v1alpha1.RoleAliasList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested roleAliases.
func (c *FakeRoleAliases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rolealiasesResource, c.ns, opts))

}

// Create takes the representation of a roleAlias and creates it.  Returns the server's representation of the roleAlias, and an error, if there is any.
func (c *FakeRoleAliases) Create(ctx context.Context, roleAlias *v1alpha1.RoleAlias, opts v1.CreateOptions) (result *v1alpha1.RoleAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rolealiasesResource, c.ns, roleAlias), &v1alpha1.RoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAlias), err
}

// Update takes the representation of a roleAlias and updates it. Returns the server's representation of the roleAlias, and an error, if there is any.
func (c *FakeRoleAliases) Update(ctx context.Context, roleAlias *v1alpha1.RoleAlias, opts v1.UpdateOptions) (result *v1alpha1.RoleAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rolealiasesResource, c.ns, roleAlias), &v1alpha1.RoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAlias), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRoleAliases) UpdateStatus(ctx context.Context, roleAlias *v1alpha1.RoleAlias, opts v1.UpdateOptions) (*v1alpha1.RoleAlias, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rolealiasesResource, "status", c.ns, roleAlias), &v1alpha1.RoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAlias), err
}

// Delete takes name of the roleAlias and deletes it. Returns an error if one occurs.
func (c *FakeRoleAliases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(rolealiasesResource, c.ns, name), &v1alpha1.RoleAlias{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRoleAliases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rolealiasesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RoleAliasList{})
	return err
}

// Patch applies the patch and returns the patched roleAlias.
func (c *FakeRoleAliases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RoleAlias, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rolealiasesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RoleAlias{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RoleAlias), err
}
