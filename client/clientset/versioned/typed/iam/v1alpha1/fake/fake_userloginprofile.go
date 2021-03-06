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

// FakeUserLoginProfiles implements UserLoginProfileInterface
type FakeUserLoginProfiles struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var userloginprofilesResource = schema.GroupVersionResource{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Resource: "userloginprofiles"}

var userloginprofilesKind = schema.GroupVersionKind{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Kind: "UserLoginProfile"}

// Get takes name of the userLoginProfile, and returns the corresponding userLoginProfile object, and an error if there is any.
func (c *FakeUserLoginProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserLoginProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(userloginprofilesResource, c.ns, name), &v1alpha1.UserLoginProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserLoginProfile), err
}

// List takes label and field selectors, and returns the list of UserLoginProfiles that match those selectors.
func (c *FakeUserLoginProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserLoginProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(userloginprofilesResource, userloginprofilesKind, c.ns, opts), &v1alpha1.UserLoginProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UserLoginProfileList{ListMeta: obj.(*v1alpha1.UserLoginProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.UserLoginProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userLoginProfiles.
func (c *FakeUserLoginProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(userloginprofilesResource, c.ns, opts))

}

// Create takes the representation of a userLoginProfile and creates it.  Returns the server's representation of the userLoginProfile, and an error, if there is any.
func (c *FakeUserLoginProfiles) Create(ctx context.Context, userLoginProfile *v1alpha1.UserLoginProfile, opts v1.CreateOptions) (result *v1alpha1.UserLoginProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(userloginprofilesResource, c.ns, userLoginProfile), &v1alpha1.UserLoginProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserLoginProfile), err
}

// Update takes the representation of a userLoginProfile and updates it. Returns the server's representation of the userLoginProfile, and an error, if there is any.
func (c *FakeUserLoginProfiles) Update(ctx context.Context, userLoginProfile *v1alpha1.UserLoginProfile, opts v1.UpdateOptions) (result *v1alpha1.UserLoginProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(userloginprofilesResource, c.ns, userLoginProfile), &v1alpha1.UserLoginProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserLoginProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserLoginProfiles) UpdateStatus(ctx context.Context, userLoginProfile *v1alpha1.UserLoginProfile, opts v1.UpdateOptions) (*v1alpha1.UserLoginProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(userloginprofilesResource, "status", c.ns, userLoginProfile), &v1alpha1.UserLoginProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserLoginProfile), err
}

// Delete takes name of the userLoginProfile and deletes it. Returns an error if one occurs.
func (c *FakeUserLoginProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(userloginprofilesResource, c.ns, name), &v1alpha1.UserLoginProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserLoginProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(userloginprofilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UserLoginProfileList{})
	return err
}

// Patch applies the patch and returns the patched userLoginProfile.
func (c *FakeUserLoginProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserLoginProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(userloginprofilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.UserLoginProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserLoginProfile), err
}
