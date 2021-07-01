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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudfront/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKeyGroups implements KeyGroupInterface
type FakeKeyGroups struct {
	Fake *FakeCloudfrontV1alpha1
	ns   string
}

var keygroupsResource = schema.GroupVersionResource{Group: "cloudfront.aws.kubeform.com", Version: "v1alpha1", Resource: "keygroups"}

var keygroupsKind = schema.GroupVersionKind{Group: "cloudfront.aws.kubeform.com", Version: "v1alpha1", Kind: "KeyGroup"}

// Get takes name of the keyGroup, and returns the corresponding keyGroup object, and an error if there is any.
func (c *FakeKeyGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KeyGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keygroupsResource, c.ns, name), &v1alpha1.KeyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyGroup), err
}

// List takes label and field selectors, and returns the list of KeyGroups that match those selectors.
func (c *FakeKeyGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KeyGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keygroupsResource, keygroupsKind, c.ns, opts), &v1alpha1.KeyGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KeyGroupList{ListMeta: obj.(*v1alpha1.KeyGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.KeyGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keyGroups.
func (c *FakeKeyGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keygroupsResource, c.ns, opts))

}

// Create takes the representation of a keyGroup and creates it.  Returns the server's representation of the keyGroup, and an error, if there is any.
func (c *FakeKeyGroups) Create(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.CreateOptions) (result *v1alpha1.KeyGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keygroupsResource, c.ns, keyGroup), &v1alpha1.KeyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyGroup), err
}

// Update takes the representation of a keyGroup and updates it. Returns the server's representation of the keyGroup, and an error, if there is any.
func (c *FakeKeyGroups) Update(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.UpdateOptions) (result *v1alpha1.KeyGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keygroupsResource, c.ns, keyGroup), &v1alpha1.KeyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKeyGroups) UpdateStatus(ctx context.Context, keyGroup *v1alpha1.KeyGroup, opts v1.UpdateOptions) (*v1alpha1.KeyGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(keygroupsResource, "status", c.ns, keyGroup), &v1alpha1.KeyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyGroup), err
}

// Delete takes name of the keyGroup and deletes it. Returns an error if one occurs.
func (c *FakeKeyGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(keygroupsResource, c.ns, name), &v1alpha1.KeyGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeyGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keygroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KeyGroupList{})
	return err
}

// Patch applies the patch and returns the patched keyGroup.
func (c *FakeKeyGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KeyGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keygroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KeyGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KeyGroup), err
}
