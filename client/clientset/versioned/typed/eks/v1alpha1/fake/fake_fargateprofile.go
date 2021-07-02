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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/eks/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFargateProfiles implements FargateProfileInterface
type FakeFargateProfiles struct {
	Fake *FakeEksV1alpha1
	ns   string
}

var fargateprofilesResource = schema.GroupVersionResource{Group: "eks.aws.kubeform.com", Version: "v1alpha1", Resource: "fargateprofiles"}

var fargateprofilesKind = schema.GroupVersionKind{Group: "eks.aws.kubeform.com", Version: "v1alpha1", Kind: "FargateProfile"}

// Get takes name of the fargateProfile, and returns the corresponding fargateProfile object, and an error if there is any.
func (c *FakeFargateProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FargateProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fargateprofilesResource, c.ns, name), &v1alpha1.FargateProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FargateProfile), err
}

// List takes label and field selectors, and returns the list of FargateProfiles that match those selectors.
func (c *FakeFargateProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FargateProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fargateprofilesResource, fargateprofilesKind, c.ns, opts), &v1alpha1.FargateProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FargateProfileList{ListMeta: obj.(*v1alpha1.FargateProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.FargateProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fargateProfiles.
func (c *FakeFargateProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fargateprofilesResource, c.ns, opts))

}

// Create takes the representation of a fargateProfile and creates it.  Returns the server's representation of the fargateProfile, and an error, if there is any.
func (c *FakeFargateProfiles) Create(ctx context.Context, fargateProfile *v1alpha1.FargateProfile, opts v1.CreateOptions) (result *v1alpha1.FargateProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fargateprofilesResource, c.ns, fargateProfile), &v1alpha1.FargateProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FargateProfile), err
}

// Update takes the representation of a fargateProfile and updates it. Returns the server's representation of the fargateProfile, and an error, if there is any.
func (c *FakeFargateProfiles) Update(ctx context.Context, fargateProfile *v1alpha1.FargateProfile, opts v1.UpdateOptions) (result *v1alpha1.FargateProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fargateprofilesResource, c.ns, fargateProfile), &v1alpha1.FargateProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FargateProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFargateProfiles) UpdateStatus(ctx context.Context, fargateProfile *v1alpha1.FargateProfile, opts v1.UpdateOptions) (*v1alpha1.FargateProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fargateprofilesResource, "status", c.ns, fargateProfile), &v1alpha1.FargateProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FargateProfile), err
}

// Delete takes name of the fargateProfile and deletes it. Returns an error if one occurs.
func (c *FakeFargateProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fargateprofilesResource, c.ns, name), &v1alpha1.FargateProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFargateProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fargateprofilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FargateProfileList{})
	return err
}

// Patch applies the patch and returns the patched fargateProfile.
func (c *FakeFargateProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FargateProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fargateprofilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.FargateProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FargateProfile), err
}
