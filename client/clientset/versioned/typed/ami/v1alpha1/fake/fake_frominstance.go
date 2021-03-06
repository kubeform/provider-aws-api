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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ami/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFromInstances implements FromInstanceInterface
type FakeFromInstances struct {
	Fake *FakeAmiV1alpha1
	ns   string
}

var frominstancesResource = schema.GroupVersionResource{Group: "ami.aws.kubeform.com", Version: "v1alpha1", Resource: "frominstances"}

var frominstancesKind = schema.GroupVersionKind{Group: "ami.aws.kubeform.com", Version: "v1alpha1", Kind: "FromInstance"}

// Get takes name of the fromInstance, and returns the corresponding fromInstance object, and an error if there is any.
func (c *FakeFromInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FromInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(frominstancesResource, c.ns, name), &v1alpha1.FromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromInstance), err
}

// List takes label and field selectors, and returns the list of FromInstances that match those selectors.
func (c *FakeFromInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FromInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(frominstancesResource, frominstancesKind, c.ns, opts), &v1alpha1.FromInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FromInstanceList{ListMeta: obj.(*v1alpha1.FromInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.FromInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fromInstances.
func (c *FakeFromInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(frominstancesResource, c.ns, opts))

}

// Create takes the representation of a fromInstance and creates it.  Returns the server's representation of the fromInstance, and an error, if there is any.
func (c *FakeFromInstances) Create(ctx context.Context, fromInstance *v1alpha1.FromInstance, opts v1.CreateOptions) (result *v1alpha1.FromInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(frominstancesResource, c.ns, fromInstance), &v1alpha1.FromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromInstance), err
}

// Update takes the representation of a fromInstance and updates it. Returns the server's representation of the fromInstance, and an error, if there is any.
func (c *FakeFromInstances) Update(ctx context.Context, fromInstance *v1alpha1.FromInstance, opts v1.UpdateOptions) (result *v1alpha1.FromInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(frominstancesResource, c.ns, fromInstance), &v1alpha1.FromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFromInstances) UpdateStatus(ctx context.Context, fromInstance *v1alpha1.FromInstance, opts v1.UpdateOptions) (*v1alpha1.FromInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(frominstancesResource, "status", c.ns, fromInstance), &v1alpha1.FromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromInstance), err
}

// Delete takes name of the fromInstance and deletes it. Returns an error if one occurs.
func (c *FakeFromInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(frominstancesResource, c.ns, name), &v1alpha1.FromInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFromInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(frominstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FromInstanceList{})
	return err
}

// Patch applies the patch and returns the patched fromInstance.
func (c *FakeFromInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FromInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(frominstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FromInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FromInstance), err
}
