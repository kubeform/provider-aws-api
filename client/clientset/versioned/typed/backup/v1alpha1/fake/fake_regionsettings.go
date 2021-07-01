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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/backup/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRegionSettingses implements RegionSettingsInterface
type FakeRegionSettingses struct {
	Fake *FakeBackupV1alpha1
	ns   string
}

var regionsettingsesResource = schema.GroupVersionResource{Group: "backup.aws.kubeform.com", Version: "v1alpha1", Resource: "regionsettingses"}

var regionsettingsesKind = schema.GroupVersionKind{Group: "backup.aws.kubeform.com", Version: "v1alpha1", Kind: "RegionSettings"}

// Get takes name of the regionSettings, and returns the corresponding regionSettings object, and an error if there is any.
func (c *FakeRegionSettingses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RegionSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(regionsettingsesResource, c.ns, name), &v1alpha1.RegionSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionSettings), err
}

// List takes label and field selectors, and returns the list of RegionSettingses that match those selectors.
func (c *FakeRegionSettingses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RegionSettingsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(regionsettingsesResource, regionsettingsesKind, c.ns, opts), &v1alpha1.RegionSettingsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RegionSettingsList{ListMeta: obj.(*v1alpha1.RegionSettingsList).ListMeta}
	for _, item := range obj.(*v1alpha1.RegionSettingsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested regionSettingses.
func (c *FakeRegionSettingses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(regionsettingsesResource, c.ns, opts))

}

// Create takes the representation of a regionSettings and creates it.  Returns the server's representation of the regionSettings, and an error, if there is any.
func (c *FakeRegionSettingses) Create(ctx context.Context, regionSettings *v1alpha1.RegionSettings, opts v1.CreateOptions) (result *v1alpha1.RegionSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(regionsettingsesResource, c.ns, regionSettings), &v1alpha1.RegionSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionSettings), err
}

// Update takes the representation of a regionSettings and updates it. Returns the server's representation of the regionSettings, and an error, if there is any.
func (c *FakeRegionSettingses) Update(ctx context.Context, regionSettings *v1alpha1.RegionSettings, opts v1.UpdateOptions) (result *v1alpha1.RegionSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(regionsettingsesResource, c.ns, regionSettings), &v1alpha1.RegionSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionSettings), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRegionSettingses) UpdateStatus(ctx context.Context, regionSettings *v1alpha1.RegionSettings, opts v1.UpdateOptions) (*v1alpha1.RegionSettings, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(regionsettingsesResource, "status", c.ns, regionSettings), &v1alpha1.RegionSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionSettings), err
}

// Delete takes name of the regionSettings and deletes it. Returns an error if one occurs.
func (c *FakeRegionSettingses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(regionsettingsesResource, c.ns, name), &v1alpha1.RegionSettings{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRegionSettingses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(regionsettingsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RegionSettingsList{})
	return err
}

// Patch applies the patch and returns the patched regionSettings.
func (c *FakeRegionSettingses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RegionSettings, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(regionsettingsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.RegionSettings{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RegionSettings), err
}
