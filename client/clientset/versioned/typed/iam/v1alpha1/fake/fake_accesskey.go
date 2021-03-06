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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/iam/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccessKeys implements AccessKeyInterface
type FakeAccessKeys struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var accesskeysResource = schema.GroupVersionResource{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Resource: "accesskeys"}

var accesskeysKind = schema.GroupVersionKind{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Kind: "AccessKey"}

// Get takes name of the accessKey, and returns the corresponding accessKey object, and an error if there is any.
func (c *FakeAccessKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accesskeysResource, c.ns, name), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}

// List takes label and field selectors, and returns the list of AccessKeys that match those selectors.
func (c *FakeAccessKeys) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccessKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accesskeysResource, accesskeysKind, c.ns, opts), &v1alpha1.AccessKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccessKeyList{ListMeta: obj.(*v1alpha1.AccessKeyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccessKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accessKeys.
func (c *FakeAccessKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accesskeysResource, c.ns, opts))

}

// Create takes the representation of a accessKey and creates it.  Returns the server's representation of the accessKey, and an error, if there is any.
func (c *FakeAccessKeys) Create(ctx context.Context, accessKey *v1alpha1.AccessKey, opts v1.CreateOptions) (result *v1alpha1.AccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accesskeysResource, c.ns, accessKey), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}

// Update takes the representation of a accessKey and updates it. Returns the server's representation of the accessKey, and an error, if there is any.
func (c *FakeAccessKeys) Update(ctx context.Context, accessKey *v1alpha1.AccessKey, opts v1.UpdateOptions) (result *v1alpha1.AccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accesskeysResource, c.ns, accessKey), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccessKeys) UpdateStatus(ctx context.Context, accessKey *v1alpha1.AccessKey, opts v1.UpdateOptions) (*v1alpha1.AccessKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accesskeysResource, "status", c.ns, accessKey), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}

// Delete takes name of the accessKey and deletes it. Returns an error if one occurs.
func (c *FakeAccessKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accesskeysResource, c.ns, name), &v1alpha1.AccessKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccessKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accesskeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccessKeyList{})
	return err
}

// Patch applies the patch and returns the patched accessKey.
func (c *FakeAccessKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accesskeysResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccessKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessKey), err
}
