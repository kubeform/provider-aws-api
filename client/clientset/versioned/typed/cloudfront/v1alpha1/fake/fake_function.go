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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudfront/v1alpha1"
)

// FakeFunctions implements FunctionInterface
type FakeFunctions struct {
	Fake *FakeCloudfrontV1alpha1
	ns   string
}

var functionsResource = schema.GroupVersionResource{Group: "cloudfront.aws.kubeform.com", Version: "v1alpha1", Resource: "functions"}

var functionsKind = schema.GroupVersionKind{Group: "cloudfront.aws.kubeform.com", Version: "v1alpha1", Kind: "Function"}

// Get takes name of the function, and returns the corresponding function object, and an error if there is any.
func (c *FakeFunctions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Function, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(functionsResource, c.ns, name), &v1alpha1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Function), err
}

// List takes label and field selectors, and returns the list of Functions that match those selectors.
func (c *FakeFunctions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FunctionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(functionsResource, functionsKind, c.ns, opts), &v1alpha1.FunctionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FunctionList{ListMeta: obj.(*v1alpha1.FunctionList).ListMeta}
	for _, item := range obj.(*v1alpha1.FunctionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested functions.
func (c *FakeFunctions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(functionsResource, c.ns, opts))

}

// Create takes the representation of a function and creates it.  Returns the server's representation of the function, and an error, if there is any.
func (c *FakeFunctions) Create(ctx context.Context, function *v1alpha1.Function, opts v1.CreateOptions) (result *v1alpha1.Function, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(functionsResource, c.ns, function), &v1alpha1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Function), err
}

// Update takes the representation of a function and updates it. Returns the server's representation of the function, and an error, if there is any.
func (c *FakeFunctions) Update(ctx context.Context, function *v1alpha1.Function, opts v1.UpdateOptions) (result *v1alpha1.Function, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(functionsResource, c.ns, function), &v1alpha1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Function), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFunctions) UpdateStatus(ctx context.Context, function *v1alpha1.Function, opts v1.UpdateOptions) (*v1alpha1.Function, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(functionsResource, "status", c.ns, function), &v1alpha1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Function), err
}

// Delete takes name of the function and deletes it. Returns an error if one occurs.
func (c *FakeFunctions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(functionsResource, c.ns, name), &v1alpha1.Function{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFunctions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(functionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FunctionList{})
	return err
}

// Patch applies the patch and returns the patched function.
func (c *FakeFunctions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Function, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(functionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Function{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Function), err
}
