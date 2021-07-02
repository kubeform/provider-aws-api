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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/sagemaker/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeImageVersions implements ImageVersionInterface
type FakeImageVersions struct {
	Fake *FakeSagemakerV1alpha1
	ns   string
}

var imageversionsResource = schema.GroupVersionResource{Group: "sagemaker.aws.kubeform.com", Version: "v1alpha1", Resource: "imageversions"}

var imageversionsKind = schema.GroupVersionKind{Group: "sagemaker.aws.kubeform.com", Version: "v1alpha1", Kind: "ImageVersion"}

// Get takes name of the imageVersion, and returns the corresponding imageVersion object, and an error if there is any.
func (c *FakeImageVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImageVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imageversionsResource, c.ns, name), &v1alpha1.ImageVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageVersion), err
}

// List takes label and field selectors, and returns the list of ImageVersions that match those selectors.
func (c *FakeImageVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImageVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imageversionsResource, imageversionsKind, c.ns, opts), &v1alpha1.ImageVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImageVersionList{ListMeta: obj.(*v1alpha1.ImageVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImageVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageVersions.
func (c *FakeImageVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(imageversionsResource, c.ns, opts))

}

// Create takes the representation of a imageVersion and creates it.  Returns the server's representation of the imageVersion, and an error, if there is any.
func (c *FakeImageVersions) Create(ctx context.Context, imageVersion *v1alpha1.ImageVersion, opts v1.CreateOptions) (result *v1alpha1.ImageVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imageversionsResource, c.ns, imageVersion), &v1alpha1.ImageVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageVersion), err
}

// Update takes the representation of a imageVersion and updates it. Returns the server's representation of the imageVersion, and an error, if there is any.
func (c *FakeImageVersions) Update(ctx context.Context, imageVersion *v1alpha1.ImageVersion, opts v1.UpdateOptions) (result *v1alpha1.ImageVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imageversionsResource, c.ns, imageVersion), &v1alpha1.ImageVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageVersions) UpdateStatus(ctx context.Context, imageVersion *v1alpha1.ImageVersion, opts v1.UpdateOptions) (*v1alpha1.ImageVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(imageversionsResource, "status", c.ns, imageVersion), &v1alpha1.ImageVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageVersion), err
}

// Delete takes name of the imageVersion and deletes it. Returns an error if one occurs.
func (c *FakeImageVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(imageversionsResource, c.ns, name), &v1alpha1.ImageVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(imageversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImageVersionList{})
	return err
}

// Patch applies the patch and returns the patched imageVersion.
func (c *FakeImageVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imageversionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ImageVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageVersion), err
}
