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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/apigatewayv2/v1alpha1"
)

// FakeApiMappings implements ApiMappingInterface
type FakeApiMappings struct {
	Fake *FakeApigatewayv2V1alpha1
	ns   string
}

var apimappingsResource = schema.GroupVersionResource{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Resource: "apimappings"}

var apimappingsKind = schema.GroupVersionKind{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Kind: "ApiMapping"}

// Get takes name of the apiMapping, and returns the corresponding apiMapping object, and an error if there is any.
func (c *FakeApiMappings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApiMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(apimappingsResource, c.ns, name), &v1alpha1.ApiMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiMapping), err
}

// List takes label and field selectors, and returns the list of ApiMappings that match those selectors.
func (c *FakeApiMappings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApiMappingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(apimappingsResource, apimappingsKind, c.ns, opts), &v1alpha1.ApiMappingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApiMappingList{ListMeta: obj.(*v1alpha1.ApiMappingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApiMappingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested apiMappings.
func (c *FakeApiMappings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(apimappingsResource, c.ns, opts))

}

// Create takes the representation of a apiMapping and creates it.  Returns the server's representation of the apiMapping, and an error, if there is any.
func (c *FakeApiMappings) Create(ctx context.Context, apiMapping *v1alpha1.ApiMapping, opts v1.CreateOptions) (result *v1alpha1.ApiMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(apimappingsResource, c.ns, apiMapping), &v1alpha1.ApiMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiMapping), err
}

// Update takes the representation of a apiMapping and updates it. Returns the server's representation of the apiMapping, and an error, if there is any.
func (c *FakeApiMappings) Update(ctx context.Context, apiMapping *v1alpha1.ApiMapping, opts v1.UpdateOptions) (result *v1alpha1.ApiMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(apimappingsResource, c.ns, apiMapping), &v1alpha1.ApiMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiMapping), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApiMappings) UpdateStatus(ctx context.Context, apiMapping *v1alpha1.ApiMapping, opts v1.UpdateOptions) (*v1alpha1.ApiMapping, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(apimappingsResource, "status", c.ns, apiMapping), &v1alpha1.ApiMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiMapping), err
}

// Delete takes name of the apiMapping and deletes it. Returns an error if one occurs.
func (c *FakeApiMappings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(apimappingsResource, c.ns, name), &v1alpha1.ApiMapping{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApiMappings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(apimappingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApiMappingList{})
	return err
}

// Patch applies the patch and returns the patched apiMapping.
func (c *FakeApiMappings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApiMapping, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(apimappingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApiMapping{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApiMapping), err
}
