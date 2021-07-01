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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/organizations/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDelegatedAdministrators implements DelegatedAdministratorInterface
type FakeDelegatedAdministrators struct {
	Fake *FakeOrganizationsV1alpha1
	ns   string
}

var delegatedadministratorsResource = schema.GroupVersionResource{Group: "organizations.aws.kubeform.com", Version: "v1alpha1", Resource: "delegatedadministrators"}

var delegatedadministratorsKind = schema.GroupVersionKind{Group: "organizations.aws.kubeform.com", Version: "v1alpha1", Kind: "DelegatedAdministrator"}

// Get takes name of the delegatedAdministrator, and returns the corresponding delegatedAdministrator object, and an error if there is any.
func (c *FakeDelegatedAdministrators) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DelegatedAdministrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(delegatedadministratorsResource, c.ns, name), &v1alpha1.DelegatedAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DelegatedAdministrator), err
}

// List takes label and field selectors, and returns the list of DelegatedAdministrators that match those selectors.
func (c *FakeDelegatedAdministrators) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DelegatedAdministratorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(delegatedadministratorsResource, delegatedadministratorsKind, c.ns, opts), &v1alpha1.DelegatedAdministratorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DelegatedAdministratorList{ListMeta: obj.(*v1alpha1.DelegatedAdministratorList).ListMeta}
	for _, item := range obj.(*v1alpha1.DelegatedAdministratorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested delegatedAdministrators.
func (c *FakeDelegatedAdministrators) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(delegatedadministratorsResource, c.ns, opts))

}

// Create takes the representation of a delegatedAdministrator and creates it.  Returns the server's representation of the delegatedAdministrator, and an error, if there is any.
func (c *FakeDelegatedAdministrators) Create(ctx context.Context, delegatedAdministrator *v1alpha1.DelegatedAdministrator, opts v1.CreateOptions) (result *v1alpha1.DelegatedAdministrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(delegatedadministratorsResource, c.ns, delegatedAdministrator), &v1alpha1.DelegatedAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DelegatedAdministrator), err
}

// Update takes the representation of a delegatedAdministrator and updates it. Returns the server's representation of the delegatedAdministrator, and an error, if there is any.
func (c *FakeDelegatedAdministrators) Update(ctx context.Context, delegatedAdministrator *v1alpha1.DelegatedAdministrator, opts v1.UpdateOptions) (result *v1alpha1.DelegatedAdministrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(delegatedadministratorsResource, c.ns, delegatedAdministrator), &v1alpha1.DelegatedAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DelegatedAdministrator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDelegatedAdministrators) UpdateStatus(ctx context.Context, delegatedAdministrator *v1alpha1.DelegatedAdministrator, opts v1.UpdateOptions) (*v1alpha1.DelegatedAdministrator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(delegatedadministratorsResource, "status", c.ns, delegatedAdministrator), &v1alpha1.DelegatedAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DelegatedAdministrator), err
}

// Delete takes name of the delegatedAdministrator and deletes it. Returns an error if one occurs.
func (c *FakeDelegatedAdministrators) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(delegatedadministratorsResource, c.ns, name), &v1alpha1.DelegatedAdministrator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDelegatedAdministrators) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(delegatedadministratorsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DelegatedAdministratorList{})
	return err
}

// Patch applies the patch and returns the patched delegatedAdministrator.
func (c *FakeDelegatedAdministrators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DelegatedAdministrator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(delegatedadministratorsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DelegatedAdministrator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DelegatedAdministrator), err
}
