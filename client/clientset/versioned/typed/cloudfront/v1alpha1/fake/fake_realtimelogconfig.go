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

// FakeRealtimeLogConfigs implements RealtimeLogConfigInterface
type FakeRealtimeLogConfigs struct {
	Fake *FakeCloudfrontV1alpha1
	ns   string
}

var realtimelogconfigsResource = schema.GroupVersionResource{Group: "cloudfront.aws.kubeform.com", Version: "v1alpha1", Resource: "realtimelogconfigs"}

var realtimelogconfigsKind = schema.GroupVersionKind{Group: "cloudfront.aws.kubeform.com", Version: "v1alpha1", Kind: "RealtimeLogConfig"}

// Get takes name of the realtimeLogConfig, and returns the corresponding realtimeLogConfig object, and an error if there is any.
func (c *FakeRealtimeLogConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RealtimeLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(realtimelogconfigsResource, c.ns, name), &v1alpha1.RealtimeLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealtimeLogConfig), err
}

// List takes label and field selectors, and returns the list of RealtimeLogConfigs that match those selectors.
func (c *FakeRealtimeLogConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RealtimeLogConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(realtimelogconfigsResource, realtimelogconfigsKind, c.ns, opts), &v1alpha1.RealtimeLogConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RealtimeLogConfigList{ListMeta: obj.(*v1alpha1.RealtimeLogConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.RealtimeLogConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested realtimeLogConfigs.
func (c *FakeRealtimeLogConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(realtimelogconfigsResource, c.ns, opts))

}

// Create takes the representation of a realtimeLogConfig and creates it.  Returns the server's representation of the realtimeLogConfig, and an error, if there is any.
func (c *FakeRealtimeLogConfigs) Create(ctx context.Context, realtimeLogConfig *v1alpha1.RealtimeLogConfig, opts v1.CreateOptions) (result *v1alpha1.RealtimeLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(realtimelogconfigsResource, c.ns, realtimeLogConfig), &v1alpha1.RealtimeLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealtimeLogConfig), err
}

// Update takes the representation of a realtimeLogConfig and updates it. Returns the server's representation of the realtimeLogConfig, and an error, if there is any.
func (c *FakeRealtimeLogConfigs) Update(ctx context.Context, realtimeLogConfig *v1alpha1.RealtimeLogConfig, opts v1.UpdateOptions) (result *v1alpha1.RealtimeLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(realtimelogconfigsResource, c.ns, realtimeLogConfig), &v1alpha1.RealtimeLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealtimeLogConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRealtimeLogConfigs) UpdateStatus(ctx context.Context, realtimeLogConfig *v1alpha1.RealtimeLogConfig, opts v1.UpdateOptions) (*v1alpha1.RealtimeLogConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(realtimelogconfigsResource, "status", c.ns, realtimeLogConfig), &v1alpha1.RealtimeLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealtimeLogConfig), err
}

// Delete takes name of the realtimeLogConfig and deletes it. Returns an error if one occurs.
func (c *FakeRealtimeLogConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(realtimelogconfigsResource, c.ns, name), &v1alpha1.RealtimeLogConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRealtimeLogConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(realtimelogconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RealtimeLogConfigList{})
	return err
}

// Patch applies the patch and returns the patched realtimeLogConfig.
func (c *FakeRealtimeLogConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RealtimeLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(realtimelogconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.RealtimeLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RealtimeLogConfig), err
}
