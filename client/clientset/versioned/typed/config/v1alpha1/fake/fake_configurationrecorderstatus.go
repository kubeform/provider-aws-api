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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/config/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConfigurationRecorderStatuses implements ConfigurationRecorderStatusInterface
type FakeConfigurationRecorderStatuses struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var configurationrecorderstatusesResource = schema.GroupVersionResource{Group: "config.aws.kubeform.com", Version: "v1alpha1", Resource: "configurationrecorderstatuses"}

var configurationrecorderstatusesKind = schema.GroupVersionKind{Group: "config.aws.kubeform.com", Version: "v1alpha1", Kind: "ConfigurationRecorderStatus"}

// Get takes name of the configurationRecorderStatus, and returns the corresponding configurationRecorderStatus object, and an error if there is any.
func (c *FakeConfigurationRecorderStatuses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConfigurationRecorderStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(configurationrecorderstatusesResource, c.ns, name), &v1alpha1.ConfigurationRecorderStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationRecorderStatus), err
}

// List takes label and field selectors, and returns the list of ConfigurationRecorderStatuses that match those selectors.
func (c *FakeConfigurationRecorderStatuses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigurationRecorderStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(configurationrecorderstatusesResource, configurationrecorderstatusesKind, c.ns, opts), &v1alpha1.ConfigurationRecorderStatusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConfigurationRecorderStatusList{ListMeta: obj.(*v1alpha1.ConfigurationRecorderStatusList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConfigurationRecorderStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested configurationRecorderStatuses.
func (c *FakeConfigurationRecorderStatuses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(configurationrecorderstatusesResource, c.ns, opts))

}

// Create takes the representation of a configurationRecorderStatus and creates it.  Returns the server's representation of the configurationRecorderStatus, and an error, if there is any.
func (c *FakeConfigurationRecorderStatuses) Create(ctx context.Context, configurationRecorderStatus *v1alpha1.ConfigurationRecorderStatus, opts v1.CreateOptions) (result *v1alpha1.ConfigurationRecorderStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(configurationrecorderstatusesResource, c.ns, configurationRecorderStatus), &v1alpha1.ConfigurationRecorderStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationRecorderStatus), err
}

// Update takes the representation of a configurationRecorderStatus and updates it. Returns the server's representation of the configurationRecorderStatus, and an error, if there is any.
func (c *FakeConfigurationRecorderStatuses) Update(ctx context.Context, configurationRecorderStatus *v1alpha1.ConfigurationRecorderStatus, opts v1.UpdateOptions) (result *v1alpha1.ConfigurationRecorderStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(configurationrecorderstatusesResource, c.ns, configurationRecorderStatus), &v1alpha1.ConfigurationRecorderStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationRecorderStatus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConfigurationRecorderStatuses) UpdateStatus(ctx context.Context, configurationRecorderStatus *v1alpha1.ConfigurationRecorderStatus, opts v1.UpdateOptions) (*v1alpha1.ConfigurationRecorderStatus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(configurationrecorderstatusesResource, "status", c.ns, configurationRecorderStatus), &v1alpha1.ConfigurationRecorderStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationRecorderStatus), err
}

// Delete takes name of the configurationRecorderStatus and deletes it. Returns an error if one occurs.
func (c *FakeConfigurationRecorderStatuses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(configurationrecorderstatusesResource, c.ns, name), &v1alpha1.ConfigurationRecorderStatus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConfigurationRecorderStatuses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(configurationrecorderstatusesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConfigurationRecorderStatusList{})
	return err
}

// Patch applies the patch and returns the patched configurationRecorderStatus.
func (c *FakeConfigurationRecorderStatuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConfigurationRecorderStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(configurationrecorderstatusesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConfigurationRecorderStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConfigurationRecorderStatus), err
}
