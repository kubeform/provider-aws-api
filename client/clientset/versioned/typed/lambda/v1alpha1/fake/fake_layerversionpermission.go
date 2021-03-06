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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/lambda/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLayerVersionPermissions implements LayerVersionPermissionInterface
type FakeLayerVersionPermissions struct {
	Fake *FakeLambdaV1alpha1
	ns   string
}

var layerversionpermissionsResource = schema.GroupVersionResource{Group: "lambda.aws.kubeform.com", Version: "v1alpha1", Resource: "layerversionpermissions"}

var layerversionpermissionsKind = schema.GroupVersionKind{Group: "lambda.aws.kubeform.com", Version: "v1alpha1", Kind: "LayerVersionPermission"}

// Get takes name of the layerVersionPermission, and returns the corresponding layerVersionPermission object, and an error if there is any.
func (c *FakeLayerVersionPermissions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LayerVersionPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(layerversionpermissionsResource, c.ns, name), &v1alpha1.LayerVersionPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersionPermission), err
}

// List takes label and field selectors, and returns the list of LayerVersionPermissions that match those selectors.
func (c *FakeLayerVersionPermissions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LayerVersionPermissionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(layerversionpermissionsResource, layerversionpermissionsKind, c.ns, opts), &v1alpha1.LayerVersionPermissionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LayerVersionPermissionList{ListMeta: obj.(*v1alpha1.LayerVersionPermissionList).ListMeta}
	for _, item := range obj.(*v1alpha1.LayerVersionPermissionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested layerVersionPermissions.
func (c *FakeLayerVersionPermissions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(layerversionpermissionsResource, c.ns, opts))

}

// Create takes the representation of a layerVersionPermission and creates it.  Returns the server's representation of the layerVersionPermission, and an error, if there is any.
func (c *FakeLayerVersionPermissions) Create(ctx context.Context, layerVersionPermission *v1alpha1.LayerVersionPermission, opts v1.CreateOptions) (result *v1alpha1.LayerVersionPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(layerversionpermissionsResource, c.ns, layerVersionPermission), &v1alpha1.LayerVersionPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersionPermission), err
}

// Update takes the representation of a layerVersionPermission and updates it. Returns the server's representation of the layerVersionPermission, and an error, if there is any.
func (c *FakeLayerVersionPermissions) Update(ctx context.Context, layerVersionPermission *v1alpha1.LayerVersionPermission, opts v1.UpdateOptions) (result *v1alpha1.LayerVersionPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(layerversionpermissionsResource, c.ns, layerVersionPermission), &v1alpha1.LayerVersionPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersionPermission), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLayerVersionPermissions) UpdateStatus(ctx context.Context, layerVersionPermission *v1alpha1.LayerVersionPermission, opts v1.UpdateOptions) (*v1alpha1.LayerVersionPermission, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(layerversionpermissionsResource, "status", c.ns, layerVersionPermission), &v1alpha1.LayerVersionPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersionPermission), err
}

// Delete takes name of the layerVersionPermission and deletes it. Returns an error if one occurs.
func (c *FakeLayerVersionPermissions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(layerversionpermissionsResource, c.ns, name), &v1alpha1.LayerVersionPermission{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLayerVersionPermissions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(layerversionpermissionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LayerVersionPermissionList{})
	return err
}

// Patch applies the patch and returns the patched layerVersionPermission.
func (c *FakeLayerVersionPermissions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LayerVersionPermission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(layerversionpermissionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LayerVersionPermission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersionPermission), err
}
