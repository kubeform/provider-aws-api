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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/glue/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDevEndpoints implements DevEndpointInterface
type FakeDevEndpoints struct {
	Fake *FakeGlueV1alpha1
	ns   string
}

var devendpointsResource = schema.GroupVersionResource{Group: "glue.aws.kubeform.com", Version: "v1alpha1", Resource: "devendpoints"}

var devendpointsKind = schema.GroupVersionKind{Group: "glue.aws.kubeform.com", Version: "v1alpha1", Kind: "DevEndpoint"}

// Get takes name of the devEndpoint, and returns the corresponding devEndpoint object, and an error if there is any.
func (c *FakeDevEndpoints) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DevEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(devendpointsResource, c.ns, name), &v1alpha1.DevEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevEndpoint), err
}

// List takes label and field selectors, and returns the list of DevEndpoints that match those selectors.
func (c *FakeDevEndpoints) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DevEndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(devendpointsResource, devendpointsKind, c.ns, opts), &v1alpha1.DevEndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DevEndpointList{ListMeta: obj.(*v1alpha1.DevEndpointList).ListMeta}
	for _, item := range obj.(*v1alpha1.DevEndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested devEndpoints.
func (c *FakeDevEndpoints) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(devendpointsResource, c.ns, opts))

}

// Create takes the representation of a devEndpoint and creates it.  Returns the server's representation of the devEndpoint, and an error, if there is any.
func (c *FakeDevEndpoints) Create(ctx context.Context, devEndpoint *v1alpha1.DevEndpoint, opts v1.CreateOptions) (result *v1alpha1.DevEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(devendpointsResource, c.ns, devEndpoint), &v1alpha1.DevEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevEndpoint), err
}

// Update takes the representation of a devEndpoint and updates it. Returns the server's representation of the devEndpoint, and an error, if there is any.
func (c *FakeDevEndpoints) Update(ctx context.Context, devEndpoint *v1alpha1.DevEndpoint, opts v1.UpdateOptions) (result *v1alpha1.DevEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(devendpointsResource, c.ns, devEndpoint), &v1alpha1.DevEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevEndpoint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDevEndpoints) UpdateStatus(ctx context.Context, devEndpoint *v1alpha1.DevEndpoint, opts v1.UpdateOptions) (*v1alpha1.DevEndpoint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(devendpointsResource, "status", c.ns, devEndpoint), &v1alpha1.DevEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevEndpoint), err
}

// Delete takes name of the devEndpoint and deletes it. Returns an error if one occurs.
func (c *FakeDevEndpoints) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(devendpointsResource, c.ns, name), &v1alpha1.DevEndpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDevEndpoints) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(devendpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DevEndpointList{})
	return err
}

// Patch applies the patch and returns the patched devEndpoint.
func (c *FakeDevEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DevEndpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(devendpointsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DevEndpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DevEndpoint), err
}
