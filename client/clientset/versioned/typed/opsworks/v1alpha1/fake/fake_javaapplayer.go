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

// FakeJavaAppLayers implements JavaAppLayerInterface
type FakeJavaAppLayers struct {
	Fake *FakeOpsworksV1alpha1
	ns   string
}

var javaapplayersResource = schema.GroupVersionResource{Group: "opsworks.aws.kubeform.com", Version: "v1alpha1", Resource: "javaapplayers"}

var javaapplayersKind = schema.GroupVersionKind{Group: "opsworks.aws.kubeform.com", Version: "v1alpha1", Kind: "JavaAppLayer"}

// Get takes name of the javaAppLayer, and returns the corresponding javaAppLayer object, and an error if there is any.
func (c *FakeJavaAppLayers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JavaAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(javaapplayersResource, c.ns, name), &v1alpha1.JavaAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JavaAppLayer), err
}

// List takes label and field selectors, and returns the list of JavaAppLayers that match those selectors.
func (c *FakeJavaAppLayers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JavaAppLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(javaapplayersResource, javaapplayersKind, c.ns, opts), &v1alpha1.JavaAppLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.JavaAppLayerList{ListMeta: obj.(*v1alpha1.JavaAppLayerList).ListMeta}
	for _, item := range obj.(*v1alpha1.JavaAppLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested javaAppLayers.
func (c *FakeJavaAppLayers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(javaapplayersResource, c.ns, opts))

}

// Create takes the representation of a javaAppLayer and creates it.  Returns the server's representation of the javaAppLayer, and an error, if there is any.
func (c *FakeJavaAppLayers) Create(ctx context.Context, javaAppLayer *v1alpha1.JavaAppLayer, opts v1.CreateOptions) (result *v1alpha1.JavaAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(javaapplayersResource, c.ns, javaAppLayer), &v1alpha1.JavaAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JavaAppLayer), err
}

// Update takes the representation of a javaAppLayer and updates it. Returns the server's representation of the javaAppLayer, and an error, if there is any.
func (c *FakeJavaAppLayers) Update(ctx context.Context, javaAppLayer *v1alpha1.JavaAppLayer, opts v1.UpdateOptions) (result *v1alpha1.JavaAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(javaapplayersResource, c.ns, javaAppLayer), &v1alpha1.JavaAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JavaAppLayer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJavaAppLayers) UpdateStatus(ctx context.Context, javaAppLayer *v1alpha1.JavaAppLayer, opts v1.UpdateOptions) (*v1alpha1.JavaAppLayer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(javaapplayersResource, "status", c.ns, javaAppLayer), &v1alpha1.JavaAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JavaAppLayer), err
}

// Delete takes name of the javaAppLayer and deletes it. Returns an error if one occurs.
func (c *FakeJavaAppLayers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(javaapplayersResource, c.ns, name), &v1alpha1.JavaAppLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJavaAppLayers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(javaapplayersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.JavaAppLayerList{})
	return err
}

// Patch applies the patch and returns the patched javaAppLayer.
func (c *FakeJavaAppLayers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JavaAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(javaapplayersResource, c.ns, name, pt, data, subresources...), &v1alpha1.JavaAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JavaAppLayer), err
}
