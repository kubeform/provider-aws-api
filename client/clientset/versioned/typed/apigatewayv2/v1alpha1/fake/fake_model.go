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

// FakeModels implements ModelInterface
type FakeModels struct {
	Fake *FakeApigatewayv2V1alpha1
	ns   string
}

var modelsResource = schema.GroupVersionResource{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Resource: "models"}

var modelsKind = schema.GroupVersionKind{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Kind: "Model"}

// Get takes name of the model, and returns the corresponding model object, and an error if there is any.
func (c *FakeModels) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Model, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(modelsResource, c.ns, name), &v1alpha1.Model{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Model), err
}

// List takes label and field selectors, and returns the list of Models that match those selectors.
func (c *FakeModels) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ModelList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(modelsResource, modelsKind, c.ns, opts), &v1alpha1.ModelList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ModelList{ListMeta: obj.(*v1alpha1.ModelList).ListMeta}
	for _, item := range obj.(*v1alpha1.ModelList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested models.
func (c *FakeModels) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(modelsResource, c.ns, opts))

}

// Create takes the representation of a model and creates it.  Returns the server's representation of the model, and an error, if there is any.
func (c *FakeModels) Create(ctx context.Context, model *v1alpha1.Model, opts v1.CreateOptions) (result *v1alpha1.Model, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(modelsResource, c.ns, model), &v1alpha1.Model{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Model), err
}

// Update takes the representation of a model and updates it. Returns the server's representation of the model, and an error, if there is any.
func (c *FakeModels) Update(ctx context.Context, model *v1alpha1.Model, opts v1.UpdateOptions) (result *v1alpha1.Model, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(modelsResource, c.ns, model), &v1alpha1.Model{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Model), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeModels) UpdateStatus(ctx context.Context, model *v1alpha1.Model, opts v1.UpdateOptions) (*v1alpha1.Model, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(modelsResource, "status", c.ns, model), &v1alpha1.Model{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Model), err
}

// Delete takes name of the model and deletes it. Returns an error if one occurs.
func (c *FakeModels) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(modelsResource, c.ns, name), &v1alpha1.Model{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeModels) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(modelsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ModelList{})
	return err
}

// Patch applies the patch and returns the patched model.
func (c *FakeModels) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Model, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(modelsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Model{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Model), err
}
