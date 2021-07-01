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

// FakeManagedPrefixLists implements ManagedPrefixListInterface
type FakeManagedPrefixLists struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var managedprefixlistsResource = schema.GroupVersionResource{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Resource: "managedprefixlists"}

var managedprefixlistsKind = schema.GroupVersionKind{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Kind: "ManagedPrefixList"}

// Get takes name of the managedPrefixList, and returns the corresponding managedPrefixList object, and an error if there is any.
func (c *FakeManagedPrefixLists) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagedPrefixList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedprefixlistsResource, c.ns, name), &v1alpha1.ManagedPrefixList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedPrefixList), err
}

// List takes label and field selectors, and returns the list of ManagedPrefixLists that match those selectors.
func (c *FakeManagedPrefixLists) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagedPrefixListList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedprefixlistsResource, managedprefixlistsKind, c.ns, opts), &v1alpha1.ManagedPrefixListList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagedPrefixListList{ListMeta: obj.(*v1alpha1.ManagedPrefixListList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagedPrefixListList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedPrefixLists.
func (c *FakeManagedPrefixLists) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedprefixlistsResource, c.ns, opts))

}

// Create takes the representation of a managedPrefixList and creates it.  Returns the server's representation of the managedPrefixList, and an error, if there is any.
func (c *FakeManagedPrefixLists) Create(ctx context.Context, managedPrefixList *v1alpha1.ManagedPrefixList, opts v1.CreateOptions) (result *v1alpha1.ManagedPrefixList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedprefixlistsResource, c.ns, managedPrefixList), &v1alpha1.ManagedPrefixList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedPrefixList), err
}

// Update takes the representation of a managedPrefixList and updates it. Returns the server's representation of the managedPrefixList, and an error, if there is any.
func (c *FakeManagedPrefixLists) Update(ctx context.Context, managedPrefixList *v1alpha1.ManagedPrefixList, opts v1.UpdateOptions) (result *v1alpha1.ManagedPrefixList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedprefixlistsResource, c.ns, managedPrefixList), &v1alpha1.ManagedPrefixList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedPrefixList), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedPrefixLists) UpdateStatus(ctx context.Context, managedPrefixList *v1alpha1.ManagedPrefixList, opts v1.UpdateOptions) (*v1alpha1.ManagedPrefixList, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedprefixlistsResource, "status", c.ns, managedPrefixList), &v1alpha1.ManagedPrefixList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedPrefixList), err
}

// Delete takes name of the managedPrefixList and deletes it. Returns an error if one occurs.
func (c *FakeManagedPrefixLists) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managedprefixlistsResource, c.ns, name), &v1alpha1.ManagedPrefixList{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedPrefixLists) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedprefixlistsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagedPrefixListList{})
	return err
}

// Patch applies the patch and returns the patched managedPrefixList.
func (c *FakeManagedPrefixLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagedPrefixList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedprefixlistsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagedPrefixList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagedPrefixList), err
}
