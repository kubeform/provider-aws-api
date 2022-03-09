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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudcontrolapi/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCloudcontrolapiResources implements CloudcontrolapiResourceInterface
type FakeCloudcontrolapiResources struct {
	Fake *FakeCloudcontrolapiV1alpha1
	ns   string
}

var cloudcontrolapiresourcesResource = schema.GroupVersionResource{Group: "cloudcontrolapi.aws.kubeform.com", Version: "v1alpha1", Resource: "cloudcontrolapiresources"}

var cloudcontrolapiresourcesKind = schema.GroupVersionKind{Group: "cloudcontrolapi.aws.kubeform.com", Version: "v1alpha1", Kind: "CloudcontrolapiResource"}

// Get takes name of the cloudcontrolapiResource, and returns the corresponding cloudcontrolapiResource object, and an error if there is any.
func (c *FakeCloudcontrolapiResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloudcontrolapiResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cloudcontrolapiresourcesResource, c.ns, name), &v1alpha1.CloudcontrolapiResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudcontrolapiResource), err
}

// List takes label and field selectors, and returns the list of CloudcontrolapiResources that match those selectors.
func (c *FakeCloudcontrolapiResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloudcontrolapiResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cloudcontrolapiresourcesResource, cloudcontrolapiresourcesKind, c.ns, opts), &v1alpha1.CloudcontrolapiResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CloudcontrolapiResourceList{ListMeta: obj.(*v1alpha1.CloudcontrolapiResourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.CloudcontrolapiResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cloudcontrolapiResources.
func (c *FakeCloudcontrolapiResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cloudcontrolapiresourcesResource, c.ns, opts))

}

// Create takes the representation of a cloudcontrolapiResource and creates it.  Returns the server's representation of the cloudcontrolapiResource, and an error, if there is any.
func (c *FakeCloudcontrolapiResources) Create(ctx context.Context, cloudcontrolapiResource *v1alpha1.CloudcontrolapiResource, opts v1.CreateOptions) (result *v1alpha1.CloudcontrolapiResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cloudcontrolapiresourcesResource, c.ns, cloudcontrolapiResource), &v1alpha1.CloudcontrolapiResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudcontrolapiResource), err
}

// Update takes the representation of a cloudcontrolapiResource and updates it. Returns the server's representation of the cloudcontrolapiResource, and an error, if there is any.
func (c *FakeCloudcontrolapiResources) Update(ctx context.Context, cloudcontrolapiResource *v1alpha1.CloudcontrolapiResource, opts v1.UpdateOptions) (result *v1alpha1.CloudcontrolapiResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cloudcontrolapiresourcesResource, c.ns, cloudcontrolapiResource), &v1alpha1.CloudcontrolapiResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudcontrolapiResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCloudcontrolapiResources) UpdateStatus(ctx context.Context, cloudcontrolapiResource *v1alpha1.CloudcontrolapiResource, opts v1.UpdateOptions) (*v1alpha1.CloudcontrolapiResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cloudcontrolapiresourcesResource, "status", c.ns, cloudcontrolapiResource), &v1alpha1.CloudcontrolapiResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudcontrolapiResource), err
}

// Delete takes name of the cloudcontrolapiResource and deletes it. Returns an error if one occurs.
func (c *FakeCloudcontrolapiResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cloudcontrolapiresourcesResource, c.ns, name), &v1alpha1.CloudcontrolapiResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCloudcontrolapiResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cloudcontrolapiresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CloudcontrolapiResourceList{})
	return err
}

// Patch applies the patch and returns the patched cloudcontrolapiResource.
func (c *FakeCloudcontrolapiResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloudcontrolapiResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cloudcontrolapiresourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CloudcontrolapiResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CloudcontrolapiResource), err
}
