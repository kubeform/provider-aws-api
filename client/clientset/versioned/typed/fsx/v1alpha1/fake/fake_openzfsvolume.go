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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/fsx/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOpenzfsVolumes implements OpenzfsVolumeInterface
type FakeOpenzfsVolumes struct {
	Fake *FakeFsxV1alpha1
	ns   string
}

var openzfsvolumesResource = schema.GroupVersionResource{Group: "fsx.aws.kubeform.com", Version: "v1alpha1", Resource: "openzfsvolumes"}

var openzfsvolumesKind = schema.GroupVersionKind{Group: "fsx.aws.kubeform.com", Version: "v1alpha1", Kind: "OpenzfsVolume"}

// Get takes name of the openzfsVolume, and returns the corresponding openzfsVolume object, and an error if there is any.
func (c *FakeOpenzfsVolumes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OpenzfsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(openzfsvolumesResource, c.ns, name), &v1alpha1.OpenzfsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenzfsVolume), err
}

// List takes label and field selectors, and returns the list of OpenzfsVolumes that match those selectors.
func (c *FakeOpenzfsVolumes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OpenzfsVolumeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(openzfsvolumesResource, openzfsvolumesKind, c.ns, opts), &v1alpha1.OpenzfsVolumeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpenzfsVolumeList{ListMeta: obj.(*v1alpha1.OpenzfsVolumeList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpenzfsVolumeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openzfsVolumes.
func (c *FakeOpenzfsVolumes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(openzfsvolumesResource, c.ns, opts))

}

// Create takes the representation of a openzfsVolume and creates it.  Returns the server's representation of the openzfsVolume, and an error, if there is any.
func (c *FakeOpenzfsVolumes) Create(ctx context.Context, openzfsVolume *v1alpha1.OpenzfsVolume, opts v1.CreateOptions) (result *v1alpha1.OpenzfsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(openzfsvolumesResource, c.ns, openzfsVolume), &v1alpha1.OpenzfsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenzfsVolume), err
}

// Update takes the representation of a openzfsVolume and updates it. Returns the server's representation of the openzfsVolume, and an error, if there is any.
func (c *FakeOpenzfsVolumes) Update(ctx context.Context, openzfsVolume *v1alpha1.OpenzfsVolume, opts v1.UpdateOptions) (result *v1alpha1.OpenzfsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(openzfsvolumesResource, c.ns, openzfsVolume), &v1alpha1.OpenzfsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenzfsVolume), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpenzfsVolumes) UpdateStatus(ctx context.Context, openzfsVolume *v1alpha1.OpenzfsVolume, opts v1.UpdateOptions) (*v1alpha1.OpenzfsVolume, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(openzfsvolumesResource, "status", c.ns, openzfsVolume), &v1alpha1.OpenzfsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenzfsVolume), err
}

// Delete takes name of the openzfsVolume and deletes it. Returns an error if one occurs.
func (c *FakeOpenzfsVolumes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(openzfsvolumesResource, c.ns, name), &v1alpha1.OpenzfsVolume{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenzfsVolumes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(openzfsvolumesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpenzfsVolumeList{})
	return err
}

// Patch applies the patch and returns the patched openzfsVolume.
func (c *FakeOpenzfsVolumes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpenzfsVolume, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(openzfsvolumesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpenzfsVolume{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenzfsVolume), err
}
