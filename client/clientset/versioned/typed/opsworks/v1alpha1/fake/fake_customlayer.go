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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/opsworks/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCustomLayers implements CustomLayerInterface
type FakeCustomLayers struct {
	Fake *FakeOpsworksV1alpha1
	ns   string
}

var customlayersResource = schema.GroupVersionResource{Group: "opsworks.aws.kubeform.com", Version: "v1alpha1", Resource: "customlayers"}

var customlayersKind = schema.GroupVersionKind{Group: "opsworks.aws.kubeform.com", Version: "v1alpha1", Kind: "CustomLayer"}

// Get takes name of the customLayer, and returns the corresponding customLayer object, and an error if there is any.
func (c *FakeCustomLayers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CustomLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(customlayersResource, c.ns, name), &v1alpha1.CustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomLayer), err
}

// List takes label and field selectors, and returns the list of CustomLayers that match those selectors.
func (c *FakeCustomLayers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CustomLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(customlayersResource, customlayersKind, c.ns, opts), &v1alpha1.CustomLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CustomLayerList{ListMeta: obj.(*v1alpha1.CustomLayerList).ListMeta}
	for _, item := range obj.(*v1alpha1.CustomLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested customLayers.
func (c *FakeCustomLayers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(customlayersResource, c.ns, opts))

}

// Create takes the representation of a customLayer and creates it.  Returns the server's representation of the customLayer, and an error, if there is any.
func (c *FakeCustomLayers) Create(ctx context.Context, customLayer *v1alpha1.CustomLayer, opts v1.CreateOptions) (result *v1alpha1.CustomLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(customlayersResource, c.ns, customLayer), &v1alpha1.CustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomLayer), err
}

// Update takes the representation of a customLayer and updates it. Returns the server's representation of the customLayer, and an error, if there is any.
func (c *FakeCustomLayers) Update(ctx context.Context, customLayer *v1alpha1.CustomLayer, opts v1.UpdateOptions) (result *v1alpha1.CustomLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(customlayersResource, c.ns, customLayer), &v1alpha1.CustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomLayer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCustomLayers) UpdateStatus(ctx context.Context, customLayer *v1alpha1.CustomLayer, opts v1.UpdateOptions) (*v1alpha1.CustomLayer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(customlayersResource, "status", c.ns, customLayer), &v1alpha1.CustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomLayer), err
}

// Delete takes name of the customLayer and deletes it. Returns an error if one occurs.
func (c *FakeCustomLayers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(customlayersResource, c.ns, name), &v1alpha1.CustomLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCustomLayers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(customlayersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CustomLayerList{})
	return err
}

// Patch applies the patch and returns the patched customLayer.
func (c *FakeCustomLayers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CustomLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(customlayersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CustomLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CustomLayer), err
}
