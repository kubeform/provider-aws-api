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

// FakeNodejsAppLayers implements NodejsAppLayerInterface
type FakeNodejsAppLayers struct {
	Fake *FakeOpsworksV1alpha1
	ns   string
}

var nodejsapplayersResource = schema.GroupVersionResource{Group: "opsworks.aws.kubeform.com", Version: "v1alpha1", Resource: "nodejsapplayers"}

var nodejsapplayersKind = schema.GroupVersionKind{Group: "opsworks.aws.kubeform.com", Version: "v1alpha1", Kind: "NodejsAppLayer"}

// Get takes name of the nodejsAppLayer, and returns the corresponding nodejsAppLayer object, and an error if there is any.
func (c *FakeNodejsAppLayers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodejsAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodejsapplayersResource, c.ns, name), &v1alpha1.NodejsAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodejsAppLayer), err
}

// List takes label and field selectors, and returns the list of NodejsAppLayers that match those selectors.
func (c *FakeNodejsAppLayers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodejsAppLayerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodejsapplayersResource, nodejsapplayersKind, c.ns, opts), &v1alpha1.NodejsAppLayerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodejsAppLayerList{ListMeta: obj.(*v1alpha1.NodejsAppLayerList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodejsAppLayerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodejsAppLayers.
func (c *FakeNodejsAppLayers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodejsapplayersResource, c.ns, opts))

}

// Create takes the representation of a nodejsAppLayer and creates it.  Returns the server's representation of the nodejsAppLayer, and an error, if there is any.
func (c *FakeNodejsAppLayers) Create(ctx context.Context, nodejsAppLayer *v1alpha1.NodejsAppLayer, opts v1.CreateOptions) (result *v1alpha1.NodejsAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodejsapplayersResource, c.ns, nodejsAppLayer), &v1alpha1.NodejsAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodejsAppLayer), err
}

// Update takes the representation of a nodejsAppLayer and updates it. Returns the server's representation of the nodejsAppLayer, and an error, if there is any.
func (c *FakeNodejsAppLayers) Update(ctx context.Context, nodejsAppLayer *v1alpha1.NodejsAppLayer, opts v1.UpdateOptions) (result *v1alpha1.NodejsAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodejsapplayersResource, c.ns, nodejsAppLayer), &v1alpha1.NodejsAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodejsAppLayer), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodejsAppLayers) UpdateStatus(ctx context.Context, nodejsAppLayer *v1alpha1.NodejsAppLayer, opts v1.UpdateOptions) (*v1alpha1.NodejsAppLayer, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nodejsapplayersResource, "status", c.ns, nodejsAppLayer), &v1alpha1.NodejsAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodejsAppLayer), err
}

// Delete takes name of the nodejsAppLayer and deletes it. Returns an error if one occurs.
func (c *FakeNodejsAppLayers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodejsapplayersResource, c.ns, name), &v1alpha1.NodejsAppLayer{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodejsAppLayers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodejsapplayersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodejsAppLayerList{})
	return err
}

// Patch applies the patch and returns the patched nodejsAppLayer.
func (c *FakeNodejsAppLayers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodejsAppLayer, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodejsapplayersResource, c.ns, name, pt, data, subresources...), &v1alpha1.NodejsAppLayer{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodejsAppLayer), err
}
