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

// FakeFieldLevelEncryptionConfigs implements FieldLevelEncryptionConfigInterface
type FakeFieldLevelEncryptionConfigs struct {
	Fake *FakeCloudfrontV1alpha1
	ns   string
}

var fieldlevelencryptionconfigsResource = schema.GroupVersionResource{Group: "cloudfront.aws.kubeform.com", Version: "v1alpha1", Resource: "fieldlevelencryptionconfigs"}

var fieldlevelencryptionconfigsKind = schema.GroupVersionKind{Group: "cloudfront.aws.kubeform.com", Version: "v1alpha1", Kind: "FieldLevelEncryptionConfig"}

// Get takes name of the fieldLevelEncryptionConfig, and returns the corresponding fieldLevelEncryptionConfig object, and an error if there is any.
func (c *FakeFieldLevelEncryptionConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FieldLevelEncryptionConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fieldlevelencryptionconfigsResource, c.ns, name), &v1alpha1.FieldLevelEncryptionConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FieldLevelEncryptionConfig), err
}

// List takes label and field selectors, and returns the list of FieldLevelEncryptionConfigs that match those selectors.
func (c *FakeFieldLevelEncryptionConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FieldLevelEncryptionConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fieldlevelencryptionconfigsResource, fieldlevelencryptionconfigsKind, c.ns, opts), &v1alpha1.FieldLevelEncryptionConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FieldLevelEncryptionConfigList{ListMeta: obj.(*v1alpha1.FieldLevelEncryptionConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.FieldLevelEncryptionConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fieldLevelEncryptionConfigs.
func (c *FakeFieldLevelEncryptionConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fieldlevelencryptionconfigsResource, c.ns, opts))

}

// Create takes the representation of a fieldLevelEncryptionConfig and creates it.  Returns the server's representation of the fieldLevelEncryptionConfig, and an error, if there is any.
func (c *FakeFieldLevelEncryptionConfigs) Create(ctx context.Context, fieldLevelEncryptionConfig *v1alpha1.FieldLevelEncryptionConfig, opts v1.CreateOptions) (result *v1alpha1.FieldLevelEncryptionConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fieldlevelencryptionconfigsResource, c.ns, fieldLevelEncryptionConfig), &v1alpha1.FieldLevelEncryptionConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FieldLevelEncryptionConfig), err
}

// Update takes the representation of a fieldLevelEncryptionConfig and updates it. Returns the server's representation of the fieldLevelEncryptionConfig, and an error, if there is any.
func (c *FakeFieldLevelEncryptionConfigs) Update(ctx context.Context, fieldLevelEncryptionConfig *v1alpha1.FieldLevelEncryptionConfig, opts v1.UpdateOptions) (result *v1alpha1.FieldLevelEncryptionConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fieldlevelencryptionconfigsResource, c.ns, fieldLevelEncryptionConfig), &v1alpha1.FieldLevelEncryptionConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FieldLevelEncryptionConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFieldLevelEncryptionConfigs) UpdateStatus(ctx context.Context, fieldLevelEncryptionConfig *v1alpha1.FieldLevelEncryptionConfig, opts v1.UpdateOptions) (*v1alpha1.FieldLevelEncryptionConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fieldlevelencryptionconfigsResource, "status", c.ns, fieldLevelEncryptionConfig), &v1alpha1.FieldLevelEncryptionConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FieldLevelEncryptionConfig), err
}

// Delete takes name of the fieldLevelEncryptionConfig and deletes it. Returns an error if one occurs.
func (c *FakeFieldLevelEncryptionConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fieldlevelencryptionconfigsResource, c.ns, name), &v1alpha1.FieldLevelEncryptionConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFieldLevelEncryptionConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fieldlevelencryptionconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FieldLevelEncryptionConfigList{})
	return err
}

// Patch applies the patch and returns the patched fieldLevelEncryptionConfig.
func (c *FakeFieldLevelEncryptionConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FieldLevelEncryptionConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fieldlevelencryptionconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FieldLevelEncryptionConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FieldLevelEncryptionConfig), err
}
