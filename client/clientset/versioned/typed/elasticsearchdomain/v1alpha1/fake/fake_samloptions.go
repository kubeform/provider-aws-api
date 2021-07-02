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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/elasticsearchdomain/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSamlOptionses implements SamlOptionsInterface
type FakeSamlOptionses struct {
	Fake *FakeElasticsearchdomainV1alpha1
	ns   string
}

var samloptionsesResource = schema.GroupVersionResource{Group: "elasticsearchdomain.aws.kubeform.com", Version: "v1alpha1", Resource: "samloptionses"}

var samloptionsesKind = schema.GroupVersionKind{Group: "elasticsearchdomain.aws.kubeform.com", Version: "v1alpha1", Kind: "SamlOptions"}

// Get takes name of the samlOptions, and returns the corresponding samlOptions object, and an error if there is any.
func (c *FakeSamlOptionses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SamlOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(samloptionsesResource, c.ns, name), &v1alpha1.SamlOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlOptions), err
}

// List takes label and field selectors, and returns the list of SamlOptionses that match those selectors.
func (c *FakeSamlOptionses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SamlOptionsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(samloptionsesResource, samloptionsesKind, c.ns, opts), &v1alpha1.SamlOptionsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SamlOptionsList{ListMeta: obj.(*v1alpha1.SamlOptionsList).ListMeta}
	for _, item := range obj.(*v1alpha1.SamlOptionsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested samlOptionses.
func (c *FakeSamlOptionses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(samloptionsesResource, c.ns, opts))

}

// Create takes the representation of a samlOptions and creates it.  Returns the server's representation of the samlOptions, and an error, if there is any.
func (c *FakeSamlOptionses) Create(ctx context.Context, samlOptions *v1alpha1.SamlOptions, opts v1.CreateOptions) (result *v1alpha1.SamlOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(samloptionsesResource, c.ns, samlOptions), &v1alpha1.SamlOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlOptions), err
}

// Update takes the representation of a samlOptions and updates it. Returns the server's representation of the samlOptions, and an error, if there is any.
func (c *FakeSamlOptionses) Update(ctx context.Context, samlOptions *v1alpha1.SamlOptions, opts v1.UpdateOptions) (result *v1alpha1.SamlOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(samloptionsesResource, c.ns, samlOptions), &v1alpha1.SamlOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlOptions), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSamlOptionses) UpdateStatus(ctx context.Context, samlOptions *v1alpha1.SamlOptions, opts v1.UpdateOptions) (*v1alpha1.SamlOptions, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(samloptionsesResource, "status", c.ns, samlOptions), &v1alpha1.SamlOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlOptions), err
}

// Delete takes name of the samlOptions and deletes it. Returns an error if one occurs.
func (c *FakeSamlOptionses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(samloptionsesResource, c.ns, name), &v1alpha1.SamlOptions{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSamlOptionses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(samloptionsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SamlOptionsList{})
	return err
}

// Patch applies the patch and returns the patched samlOptions.
func (c *FakeSamlOptionses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SamlOptions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(samloptionsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SamlOptions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlOptions), err
}
