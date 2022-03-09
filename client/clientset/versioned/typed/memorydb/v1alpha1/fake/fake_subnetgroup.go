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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/memorydb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSubnetGroups implements SubnetGroupInterface
type FakeSubnetGroups struct {
	Fake *FakeMemorydbV1alpha1
	ns   string
}

var subnetgroupsResource = schema.GroupVersionResource{Group: "memorydb.aws.kubeform.com", Version: "v1alpha1", Resource: "subnetgroups"}

var subnetgroupsKind = schema.GroupVersionKind{Group: "memorydb.aws.kubeform.com", Version: "v1alpha1", Kind: "SubnetGroup"}

// Get takes name of the subnetGroup, and returns the corresponding subnetGroup object, and an error if there is any.
func (c *FakeSubnetGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subnetgroupsResource, c.ns, name), &v1alpha1.SubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetGroup), err
}

// List takes label and field selectors, and returns the list of SubnetGroups that match those selectors.
func (c *FakeSubnetGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubnetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subnetgroupsResource, subnetgroupsKind, c.ns, opts), &v1alpha1.SubnetGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubnetGroupList{ListMeta: obj.(*v1alpha1.SubnetGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubnetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subnetGroups.
func (c *FakeSubnetGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subnetgroupsResource, c.ns, opts))

}

// Create takes the representation of a subnetGroup and creates it.  Returns the server's representation of the subnetGroup, and an error, if there is any.
func (c *FakeSubnetGroups) Create(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.CreateOptions) (result *v1alpha1.SubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subnetgroupsResource, c.ns, subnetGroup), &v1alpha1.SubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetGroup), err
}

// Update takes the representation of a subnetGroup and updates it. Returns the server's representation of the subnetGroup, and an error, if there is any.
func (c *FakeSubnetGroups) Update(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.UpdateOptions) (result *v1alpha1.SubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subnetgroupsResource, c.ns, subnetGroup), &v1alpha1.SubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubnetGroups) UpdateStatus(ctx context.Context, subnetGroup *v1alpha1.SubnetGroup, opts v1.UpdateOptions) (*v1alpha1.SubnetGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(subnetgroupsResource, "status", c.ns, subnetGroup), &v1alpha1.SubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetGroup), err
}

// Delete takes name of the subnetGroup and deletes it. Returns an error if one occurs.
func (c *FakeSubnetGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(subnetgroupsResource, c.ns, name), &v1alpha1.SubnetGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubnetGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subnetgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubnetGroupList{})
	return err
}

// Patch applies the patch and returns the patched subnetGroup.
func (c *FakeSubnetGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subnetgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetGroup), err
}
