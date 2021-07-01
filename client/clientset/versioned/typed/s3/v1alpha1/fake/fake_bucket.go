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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"
)

// FakeBuckets implements BucketInterface
type FakeBuckets struct {
	Fake *FakeS3V1alpha1
	ns   string
}

var bucketsResource = schema.GroupVersionResource{Group: "s3.aws.kubeform.com", Version: "v1alpha1", Resource: "buckets"}

var bucketsKind = schema.GroupVersionKind{Group: "s3.aws.kubeform.com", Version: "v1alpha1", Kind: "Bucket"}

// Get takes name of the bucket, and returns the corresponding bucket object, and an error if there is any.
func (c *FakeBuckets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Bucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bucketsResource, c.ns, name), &v1alpha1.Bucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bucket), err
}

// List takes label and field selectors, and returns the list of Buckets that match those selectors.
func (c *FakeBuckets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BucketList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bucketsResource, bucketsKind, c.ns, opts), &v1alpha1.BucketList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BucketList{ListMeta: obj.(*v1alpha1.BucketList).ListMeta}
	for _, item := range obj.(*v1alpha1.BucketList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buckets.
func (c *FakeBuckets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bucketsResource, c.ns, opts))

}

// Create takes the representation of a bucket and creates it.  Returns the server's representation of the bucket, and an error, if there is any.
func (c *FakeBuckets) Create(ctx context.Context, bucket *v1alpha1.Bucket, opts v1.CreateOptions) (result *v1alpha1.Bucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bucketsResource, c.ns, bucket), &v1alpha1.Bucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bucket), err
}

// Update takes the representation of a bucket and updates it. Returns the server's representation of the bucket, and an error, if there is any.
func (c *FakeBuckets) Update(ctx context.Context, bucket *v1alpha1.Bucket, opts v1.UpdateOptions) (result *v1alpha1.Bucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bucketsResource, c.ns, bucket), &v1alpha1.Bucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bucket), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuckets) UpdateStatus(ctx context.Context, bucket *v1alpha1.Bucket, opts v1.UpdateOptions) (*v1alpha1.Bucket, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bucketsResource, "status", c.ns, bucket), &v1alpha1.Bucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bucket), err
}

// Delete takes name of the bucket and deletes it. Returns an error if one occurs.
func (c *FakeBuckets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bucketsResource, c.ns, name), &v1alpha1.Bucket{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuckets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bucketsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BucketList{})
	return err
}

// Patch applies the patch and returns the patched bucket.
func (c *FakeBuckets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Bucket, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bucketsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Bucket{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Bucket), err
}
