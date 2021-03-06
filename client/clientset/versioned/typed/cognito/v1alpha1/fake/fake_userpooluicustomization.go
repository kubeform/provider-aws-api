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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cognito/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUserPoolUiCustomizations implements UserPoolUiCustomizationInterface
type FakeUserPoolUiCustomizations struct {
	Fake *FakeCognitoV1alpha1
	ns   string
}

var userpooluicustomizationsResource = schema.GroupVersionResource{Group: "cognito.aws.kubeform.com", Version: "v1alpha1", Resource: "userpooluicustomizations"}

var userpooluicustomizationsKind = schema.GroupVersionKind{Group: "cognito.aws.kubeform.com", Version: "v1alpha1", Kind: "UserPoolUiCustomization"}

// Get takes name of the userPoolUiCustomization, and returns the corresponding userPoolUiCustomization object, and an error if there is any.
func (c *FakeUserPoolUiCustomizations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserPoolUiCustomization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(userpooluicustomizationsResource, c.ns, name), &v1alpha1.UserPoolUiCustomization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolUiCustomization), err
}

// List takes label and field selectors, and returns the list of UserPoolUiCustomizations that match those selectors.
func (c *FakeUserPoolUiCustomizations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserPoolUiCustomizationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(userpooluicustomizationsResource, userpooluicustomizationsKind, c.ns, opts), &v1alpha1.UserPoolUiCustomizationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UserPoolUiCustomizationList{ListMeta: obj.(*v1alpha1.UserPoolUiCustomizationList).ListMeta}
	for _, item := range obj.(*v1alpha1.UserPoolUiCustomizationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userPoolUiCustomizations.
func (c *FakeUserPoolUiCustomizations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(userpooluicustomizationsResource, c.ns, opts))

}

// Create takes the representation of a userPoolUiCustomization and creates it.  Returns the server's representation of the userPoolUiCustomization, and an error, if there is any.
func (c *FakeUserPoolUiCustomizations) Create(ctx context.Context, userPoolUiCustomization *v1alpha1.UserPoolUiCustomization, opts v1.CreateOptions) (result *v1alpha1.UserPoolUiCustomization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(userpooluicustomizationsResource, c.ns, userPoolUiCustomization), &v1alpha1.UserPoolUiCustomization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolUiCustomization), err
}

// Update takes the representation of a userPoolUiCustomization and updates it. Returns the server's representation of the userPoolUiCustomization, and an error, if there is any.
func (c *FakeUserPoolUiCustomizations) Update(ctx context.Context, userPoolUiCustomization *v1alpha1.UserPoolUiCustomization, opts v1.UpdateOptions) (result *v1alpha1.UserPoolUiCustomization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(userpooluicustomizationsResource, c.ns, userPoolUiCustomization), &v1alpha1.UserPoolUiCustomization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolUiCustomization), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserPoolUiCustomizations) UpdateStatus(ctx context.Context, userPoolUiCustomization *v1alpha1.UserPoolUiCustomization, opts v1.UpdateOptions) (*v1alpha1.UserPoolUiCustomization, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(userpooluicustomizationsResource, "status", c.ns, userPoolUiCustomization), &v1alpha1.UserPoolUiCustomization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolUiCustomization), err
}

// Delete takes name of the userPoolUiCustomization and deletes it. Returns an error if one occurs.
func (c *FakeUserPoolUiCustomizations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(userpooluicustomizationsResource, c.ns, name), &v1alpha1.UserPoolUiCustomization{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserPoolUiCustomizations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(userpooluicustomizationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UserPoolUiCustomizationList{})
	return err
}

// Patch applies the patch and returns the patched userPoolUiCustomization.
func (c *FakeUserPoolUiCustomizations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserPoolUiCustomization, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(userpooluicustomizationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.UserPoolUiCustomization{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolUiCustomization), err
}
