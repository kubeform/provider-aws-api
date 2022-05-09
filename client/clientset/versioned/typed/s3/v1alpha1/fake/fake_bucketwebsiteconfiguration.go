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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBucketWebsiteConfigurations implements BucketWebsiteConfigurationInterface
type FakeBucketWebsiteConfigurations struct {
	Fake *FakeS3V1alpha1
	ns   string
}

var bucketwebsiteconfigurationsResource = schema.GroupVersionResource{Group: "s3.aws.kubeform.com", Version: "v1alpha1", Resource: "bucketwebsiteconfigurations"}

var bucketwebsiteconfigurationsKind = schema.GroupVersionKind{Group: "s3.aws.kubeform.com", Version: "v1alpha1", Kind: "BucketWebsiteConfiguration"}

// Get takes name of the bucketWebsiteConfiguration, and returns the corresponding bucketWebsiteConfiguration object, and an error if there is any.
func (c *FakeBucketWebsiteConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BucketWebsiteConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bucketwebsiteconfigurationsResource, c.ns, name), &v1alpha1.BucketWebsiteConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketWebsiteConfiguration), err
}

// List takes label and field selectors, and returns the list of BucketWebsiteConfigurations that match those selectors.
func (c *FakeBucketWebsiteConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BucketWebsiteConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bucketwebsiteconfigurationsResource, bucketwebsiteconfigurationsKind, c.ns, opts), &v1alpha1.BucketWebsiteConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BucketWebsiteConfigurationList{ListMeta: obj.(*v1alpha1.BucketWebsiteConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.BucketWebsiteConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bucketWebsiteConfigurations.
func (c *FakeBucketWebsiteConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bucketwebsiteconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a bucketWebsiteConfiguration and creates it.  Returns the server's representation of the bucketWebsiteConfiguration, and an error, if there is any.
func (c *FakeBucketWebsiteConfigurations) Create(ctx context.Context, bucketWebsiteConfiguration *v1alpha1.BucketWebsiteConfiguration, opts v1.CreateOptions) (result *v1alpha1.BucketWebsiteConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bucketwebsiteconfigurationsResource, c.ns, bucketWebsiteConfiguration), &v1alpha1.BucketWebsiteConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketWebsiteConfiguration), err
}

// Update takes the representation of a bucketWebsiteConfiguration and updates it. Returns the server's representation of the bucketWebsiteConfiguration, and an error, if there is any.
func (c *FakeBucketWebsiteConfigurations) Update(ctx context.Context, bucketWebsiteConfiguration *v1alpha1.BucketWebsiteConfiguration, opts v1.UpdateOptions) (result *v1alpha1.BucketWebsiteConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bucketwebsiteconfigurationsResource, c.ns, bucketWebsiteConfiguration), &v1alpha1.BucketWebsiteConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketWebsiteConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBucketWebsiteConfigurations) UpdateStatus(ctx context.Context, bucketWebsiteConfiguration *v1alpha1.BucketWebsiteConfiguration, opts v1.UpdateOptions) (*v1alpha1.BucketWebsiteConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bucketwebsiteconfigurationsResource, "status", c.ns, bucketWebsiteConfiguration), &v1alpha1.BucketWebsiteConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketWebsiteConfiguration), err
}

// Delete takes name of the bucketWebsiteConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeBucketWebsiteConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bucketwebsiteconfigurationsResource, c.ns, name), &v1alpha1.BucketWebsiteConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBucketWebsiteConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bucketwebsiteconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BucketWebsiteConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched bucketWebsiteConfiguration.
func (c *FakeBucketWebsiteConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BucketWebsiteConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bucketwebsiteconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BucketWebsiteConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BucketWebsiteConfiguration), err
}