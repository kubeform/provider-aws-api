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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/iam/v1alpha1"
)

// FakeSamlProviders implements SamlProviderInterface
type FakeSamlProviders struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var samlprovidersResource = schema.GroupVersionResource{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Resource: "samlproviders"}

var samlprovidersKind = schema.GroupVersionKind{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Kind: "SamlProvider"}

// Get takes name of the samlProvider, and returns the corresponding samlProvider object, and an error if there is any.
func (c *FakeSamlProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SamlProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(samlprovidersResource, c.ns, name), &v1alpha1.SamlProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlProvider), err
}

// List takes label and field selectors, and returns the list of SamlProviders that match those selectors.
func (c *FakeSamlProviders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SamlProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(samlprovidersResource, samlprovidersKind, c.ns, opts), &v1alpha1.SamlProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SamlProviderList{ListMeta: obj.(*v1alpha1.SamlProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.SamlProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested samlProviders.
func (c *FakeSamlProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(samlprovidersResource, c.ns, opts))

}

// Create takes the representation of a samlProvider and creates it.  Returns the server's representation of the samlProvider, and an error, if there is any.
func (c *FakeSamlProviders) Create(ctx context.Context, samlProvider *v1alpha1.SamlProvider, opts v1.CreateOptions) (result *v1alpha1.SamlProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(samlprovidersResource, c.ns, samlProvider), &v1alpha1.SamlProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlProvider), err
}

// Update takes the representation of a samlProvider and updates it. Returns the server's representation of the samlProvider, and an error, if there is any.
func (c *FakeSamlProviders) Update(ctx context.Context, samlProvider *v1alpha1.SamlProvider, opts v1.UpdateOptions) (result *v1alpha1.SamlProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(samlprovidersResource, c.ns, samlProvider), &v1alpha1.SamlProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSamlProviders) UpdateStatus(ctx context.Context, samlProvider *v1alpha1.SamlProvider, opts v1.UpdateOptions) (*v1alpha1.SamlProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(samlprovidersResource, "status", c.ns, samlProvider), &v1alpha1.SamlProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlProvider), err
}

// Delete takes name of the samlProvider and deletes it. Returns an error if one occurs.
func (c *FakeSamlProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(samlprovidersResource, c.ns, name), &v1alpha1.SamlProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSamlProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(samlprovidersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SamlProviderList{})
	return err
}

// Patch applies the patch and returns the patched samlProvider.
func (c *FakeSamlProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SamlProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(samlprovidersResource, c.ns, name, pt, data, subresources...), &v1alpha1.SamlProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SamlProvider), err
}
