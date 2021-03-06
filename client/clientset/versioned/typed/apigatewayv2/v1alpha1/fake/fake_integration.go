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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/apigatewayv2/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIntegrations implements IntegrationInterface
type FakeIntegrations struct {
	Fake *FakeApigatewayv2V1alpha1
	ns   string
}

var integrationsResource = schema.GroupVersionResource{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Resource: "integrations"}

var integrationsKind = schema.GroupVersionKind{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Kind: "Integration"}

// Get takes name of the integration, and returns the corresponding integration object, and an error if there is any.
func (c *FakeIntegrations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Integration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(integrationsResource, c.ns, name), &v1alpha1.Integration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Integration), err
}

// List takes label and field selectors, and returns the list of Integrations that match those selectors.
func (c *FakeIntegrations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IntegrationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(integrationsResource, integrationsKind, c.ns, opts), &v1alpha1.IntegrationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IntegrationList{ListMeta: obj.(*v1alpha1.IntegrationList).ListMeta}
	for _, item := range obj.(*v1alpha1.IntegrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested integrations.
func (c *FakeIntegrations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(integrationsResource, c.ns, opts))

}

// Create takes the representation of a integration and creates it.  Returns the server's representation of the integration, and an error, if there is any.
func (c *FakeIntegrations) Create(ctx context.Context, integration *v1alpha1.Integration, opts v1.CreateOptions) (result *v1alpha1.Integration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(integrationsResource, c.ns, integration), &v1alpha1.Integration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Integration), err
}

// Update takes the representation of a integration and updates it. Returns the server's representation of the integration, and an error, if there is any.
func (c *FakeIntegrations) Update(ctx context.Context, integration *v1alpha1.Integration, opts v1.UpdateOptions) (result *v1alpha1.Integration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(integrationsResource, c.ns, integration), &v1alpha1.Integration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Integration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIntegrations) UpdateStatus(ctx context.Context, integration *v1alpha1.Integration, opts v1.UpdateOptions) (*v1alpha1.Integration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(integrationsResource, "status", c.ns, integration), &v1alpha1.Integration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Integration), err
}

// Delete takes name of the integration and deletes it. Returns an error if one occurs.
func (c *FakeIntegrations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(integrationsResource, c.ns, name), &v1alpha1.Integration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIntegrations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(integrationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IntegrationList{})
	return err
}

// Patch applies the patch and returns the patched integration.
func (c *FakeIntegrations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Integration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(integrationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Integration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Integration), err
}
