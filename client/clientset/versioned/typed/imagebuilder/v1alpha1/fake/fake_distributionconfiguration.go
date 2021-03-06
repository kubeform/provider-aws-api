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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/imagebuilder/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDistributionConfigurations implements DistributionConfigurationInterface
type FakeDistributionConfigurations struct {
	Fake *FakeImagebuilderV1alpha1
	ns   string
}

var distributionconfigurationsResource = schema.GroupVersionResource{Group: "imagebuilder.aws.kubeform.com", Version: "v1alpha1", Resource: "distributionconfigurations"}

var distributionconfigurationsKind = schema.GroupVersionKind{Group: "imagebuilder.aws.kubeform.com", Version: "v1alpha1", Kind: "DistributionConfiguration"}

// Get takes name of the distributionConfiguration, and returns the corresponding distributionConfiguration object, and an error if there is any.
func (c *FakeDistributionConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DistributionConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(distributionconfigurationsResource, c.ns, name), &v1alpha1.DistributionConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DistributionConfiguration), err
}

// List takes label and field selectors, and returns the list of DistributionConfigurations that match those selectors.
func (c *FakeDistributionConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DistributionConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(distributionconfigurationsResource, distributionconfigurationsKind, c.ns, opts), &v1alpha1.DistributionConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DistributionConfigurationList{ListMeta: obj.(*v1alpha1.DistributionConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.DistributionConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested distributionConfigurations.
func (c *FakeDistributionConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(distributionconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a distributionConfiguration and creates it.  Returns the server's representation of the distributionConfiguration, and an error, if there is any.
func (c *FakeDistributionConfigurations) Create(ctx context.Context, distributionConfiguration *v1alpha1.DistributionConfiguration, opts v1.CreateOptions) (result *v1alpha1.DistributionConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(distributionconfigurationsResource, c.ns, distributionConfiguration), &v1alpha1.DistributionConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DistributionConfiguration), err
}

// Update takes the representation of a distributionConfiguration and updates it. Returns the server's representation of the distributionConfiguration, and an error, if there is any.
func (c *FakeDistributionConfigurations) Update(ctx context.Context, distributionConfiguration *v1alpha1.DistributionConfiguration, opts v1.UpdateOptions) (result *v1alpha1.DistributionConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(distributionconfigurationsResource, c.ns, distributionConfiguration), &v1alpha1.DistributionConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DistributionConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDistributionConfigurations) UpdateStatus(ctx context.Context, distributionConfiguration *v1alpha1.DistributionConfiguration, opts v1.UpdateOptions) (*v1alpha1.DistributionConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(distributionconfigurationsResource, "status", c.ns, distributionConfiguration), &v1alpha1.DistributionConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DistributionConfiguration), err
}

// Delete takes name of the distributionConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeDistributionConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(distributionconfigurationsResource, c.ns, name), &v1alpha1.DistributionConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDistributionConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(distributionconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DistributionConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched distributionConfiguration.
func (c *FakeDistributionConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DistributionConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(distributionconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DistributionConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DistributionConfiguration), err
}
