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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/opsworks/v1alpha1"
)

// FakeMysqlLayers implements MysqlLayerInterface
type FakeMysqlLayers struct {
	Fake *FakeOpsworksV1alpha1
	ns   string
}

var mysqllayersResource = schema.GroupVersionResource{Group: "opsworks.aws.kubeform.com", Version: "v1alpha1", Resource: "mysqllayers"}

var mysqllayersKind = schema.GroupVersionKind{Group: "opsworks.aws.kubeform.com", Version: "v1alpha1", Kind: "MysqlLayer"}

// Get takes name of the mysqlLayer, and returns the corresponding mysqlLayer object, and an error if there is any.
func (c *FakeMysqlLayers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MysqlLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqllayersResource, c.ns, name), &v1alpha1.MysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlLayer), err
}

// List takes label and field selectors, and returns the list of MysqlLayers that match those selectors.
func (c *FakeMysqlLayers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MysqlLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqllayersResource, mysqllayersKind, c.ns, opts), &v1alpha1.MysqlLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlLayerList{ListMeta: obj.(*v1alpha1.MysqlLayerList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mysqlLayers.
func (c *FakeMysqlLayers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqllayersResource, c.ns, opts))

}

// Create takes the representation of a mysqlLayer and creates it.  Returns the server's representation of the mysqlLayer, and an error, if there is any.
func (c *FakeMysqlLayers) Create(ctx context.Context, mysqlLayer *v1alpha1.MysqlLayer, opts v1.CreateOptions) (result *v1alpha1.MysqlLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqllayersResource, c.ns, mysqlLayer), &v1alpha1.MysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlLayer), err
}

// Update takes the representation of a mysqlLayer and updates it. Returns the server's representation of the mysqlLayer, and an error, if there is any.
func (c *FakeMysqlLayers) Update(ctx context.Context, mysqlLayer *v1alpha1.MysqlLayer, opts v1.UpdateOptions) (result *v1alpha1.MysqlLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqllayersResource, c.ns, mysqlLayer), &v1alpha1.MysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlLayer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMysqlLayers) UpdateStatus(ctx context.Context, mysqlLayer *v1alpha1.MysqlLayer, opts v1.UpdateOptions) (*v1alpha1.MysqlLayer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mysqllayersResource, "status", c.ns, mysqlLayer), &v1alpha1.MysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlLayer), err
}

// Delete takes name of the mysqlLayer and deletes it. Returns an error if one occurs.
func (c *FakeMysqlLayers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqllayersResource, c.ns, name), &v1alpha1.MysqlLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMysqlLayers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqllayersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlLayerList{})
	return err
}

// Patch applies the patch and returns the patched mysqlLayer.
func (c *FakeMysqlLayers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MysqlLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqllayersResource, c.ns, name, pt, data, subresources...), &v1alpha1.MysqlLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MysqlLayer), err
}
