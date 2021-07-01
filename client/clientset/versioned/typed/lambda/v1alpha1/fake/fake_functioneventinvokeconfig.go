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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/lambda/v1alpha1"
)

// FakeFunctionEventInvokeConfigs implements FunctionEventInvokeConfigInterface
type FakeFunctionEventInvokeConfigs struct {
	Fake *FakeLambdaV1alpha1
	ns   string
}

var functioneventinvokeconfigsResource = schema.GroupVersionResource{Group: "lambda.aws.kubeform.com", Version: "v1alpha1", Resource: "functioneventinvokeconfigs"}

var functioneventinvokeconfigsKind = schema.GroupVersionKind{Group: "lambda.aws.kubeform.com", Version: "v1alpha1", Kind: "FunctionEventInvokeConfig"}

// Get takes name of the functionEventInvokeConfig, and returns the corresponding functionEventInvokeConfig object, and an error if there is any.
func (c *FakeFunctionEventInvokeConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FunctionEventInvokeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(functioneventinvokeconfigsResource, c.ns, name), &v1alpha1.FunctionEventInvokeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionEventInvokeConfig), err
}

// List takes label and field selectors, and returns the list of FunctionEventInvokeConfigs that match those selectors.
func (c *FakeFunctionEventInvokeConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FunctionEventInvokeConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(functioneventinvokeconfigsResource, functioneventinvokeconfigsKind, c.ns, opts), &v1alpha1.FunctionEventInvokeConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FunctionEventInvokeConfigList{ListMeta: obj.(*v1alpha1.FunctionEventInvokeConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.FunctionEventInvokeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested functionEventInvokeConfigs.
func (c *FakeFunctionEventInvokeConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(functioneventinvokeconfigsResource, c.ns, opts))

}

// Create takes the representation of a functionEventInvokeConfig and creates it.  Returns the server's representation of the functionEventInvokeConfig, and an error, if there is any.
func (c *FakeFunctionEventInvokeConfigs) Create(ctx context.Context, functionEventInvokeConfig *v1alpha1.FunctionEventInvokeConfig, opts v1.CreateOptions) (result *v1alpha1.FunctionEventInvokeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(functioneventinvokeconfigsResource, c.ns, functionEventInvokeConfig), &v1alpha1.FunctionEventInvokeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionEventInvokeConfig), err
}

// Update takes the representation of a functionEventInvokeConfig and updates it. Returns the server's representation of the functionEventInvokeConfig, and an error, if there is any.
func (c *FakeFunctionEventInvokeConfigs) Update(ctx context.Context, functionEventInvokeConfig *v1alpha1.FunctionEventInvokeConfig, opts v1.UpdateOptions) (result *v1alpha1.FunctionEventInvokeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(functioneventinvokeconfigsResource, c.ns, functionEventInvokeConfig), &v1alpha1.FunctionEventInvokeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionEventInvokeConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFunctionEventInvokeConfigs) UpdateStatus(ctx context.Context, functionEventInvokeConfig *v1alpha1.FunctionEventInvokeConfig, opts v1.UpdateOptions) (*v1alpha1.FunctionEventInvokeConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(functioneventinvokeconfigsResource, "status", c.ns, functionEventInvokeConfig), &v1alpha1.FunctionEventInvokeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionEventInvokeConfig), err
}

// Delete takes name of the functionEventInvokeConfig and deletes it. Returns an error if one occurs.
func (c *FakeFunctionEventInvokeConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(functioneventinvokeconfigsResource, c.ns, name), &v1alpha1.FunctionEventInvokeConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFunctionEventInvokeConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(functioneventinvokeconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FunctionEventInvokeConfigList{})
	return err
}

// Patch applies the patch and returns the patched functionEventInvokeConfig.
func (c *FakeFunctionEventInvokeConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FunctionEventInvokeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(functioneventinvokeconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FunctionEventInvokeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FunctionEventInvokeConfig), err
}
