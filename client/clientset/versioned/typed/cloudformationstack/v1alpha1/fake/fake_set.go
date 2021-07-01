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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudformationstack/v1alpha1"
)

// FakeSets implements SetInterface
type FakeSets struct {
	Fake *FakeCloudformationstackV1alpha1
	ns   string
}

var setsResource = schema.GroupVersionResource{Group: "cloudformationstack.aws.kubeform.com", Version: "v1alpha1", Resource: "sets"}

var setsKind = schema.GroupVersionKind{Group: "cloudformationstack.aws.kubeform.com", Version: "v1alpha1", Kind: "Set"}

// Get takes name of the set, and returns the corresponding set object, and an error if there is any.
func (c *FakeSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Set, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(setsResource, c.ns, name), &v1alpha1.Set{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Set), err
}

// List takes label and field selectors, and returns the list of Sets that match those selectors.
func (c *FakeSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(setsResource, setsKind, c.ns, opts), &v1alpha1.SetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SetList{ListMeta: obj.(*v1alpha1.SetList).ListMeta}
	for _, item := range obj.(*v1alpha1.SetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sets.
func (c *FakeSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(setsResource, c.ns, opts))

}

// Create takes the representation of a set and creates it.  Returns the server's representation of the set, and an error, if there is any.
func (c *FakeSets) Create(ctx context.Context, set *v1alpha1.Set, opts v1.CreateOptions) (result *v1alpha1.Set, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(setsResource, c.ns, set), &v1alpha1.Set{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Set), err
}

// Update takes the representation of a set and updates it. Returns the server's representation of the set, and an error, if there is any.
func (c *FakeSets) Update(ctx context.Context, set *v1alpha1.Set, opts v1.UpdateOptions) (result *v1alpha1.Set, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(setsResource, c.ns, set), &v1alpha1.Set{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Set), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSets) UpdateStatus(ctx context.Context, set *v1alpha1.Set, opts v1.UpdateOptions) (*v1alpha1.Set, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(setsResource, "status", c.ns, set), &v1alpha1.Set{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Set), err
}

// Delete takes name of the set and deletes it. Returns an error if one occurs.
func (c *FakeSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(setsResource, c.ns, name), &v1alpha1.Set{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(setsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SetList{})
	return err
}

// Patch applies the patch and returns the patched set.
func (c *FakeSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Set, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(setsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Set{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Set), err
}
