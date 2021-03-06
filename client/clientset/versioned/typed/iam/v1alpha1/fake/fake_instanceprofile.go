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

// FakeInstanceProfiles implements InstanceProfileInterface
type FakeInstanceProfiles struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var instanceprofilesResource = schema.GroupVersionResource{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Resource: "instanceprofiles"}

var instanceprofilesKind = schema.GroupVersionKind{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Kind: "InstanceProfile"}

// Get takes name of the instanceProfile, and returns the corresponding instanceProfile object, and an error if there is any.
func (c *FakeInstanceProfiles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(instanceprofilesResource, c.ns, name), &v1alpha1.InstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceProfile), err
}

// List takes label and field selectors, and returns the list of InstanceProfiles that match those selectors.
func (c *FakeInstanceProfiles) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstanceProfileList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(instanceprofilesResource, instanceprofilesKind, c.ns, opts), &v1alpha1.InstanceProfileList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstanceProfileList{ListMeta: obj.(*v1alpha1.InstanceProfileList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstanceProfileList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instanceProfiles.
func (c *FakeInstanceProfiles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(instanceprofilesResource, c.ns, opts))

}

// Create takes the representation of a instanceProfile and creates it.  Returns the server's representation of the instanceProfile, and an error, if there is any.
func (c *FakeInstanceProfiles) Create(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.CreateOptions) (result *v1alpha1.InstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(instanceprofilesResource, c.ns, instanceProfile), &v1alpha1.InstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceProfile), err
}

// Update takes the representation of a instanceProfile and updates it. Returns the server's representation of the instanceProfile, and an error, if there is any.
func (c *FakeInstanceProfiles) Update(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.UpdateOptions) (result *v1alpha1.InstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(instanceprofilesResource, c.ns, instanceProfile), &v1alpha1.InstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceProfile), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstanceProfiles) UpdateStatus(ctx context.Context, instanceProfile *v1alpha1.InstanceProfile, opts v1.UpdateOptions) (*v1alpha1.InstanceProfile, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(instanceprofilesResource, "status", c.ns, instanceProfile), &v1alpha1.InstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceProfile), err
}

// Delete takes name of the instanceProfile and deletes it. Returns an error if one occurs.
func (c *FakeInstanceProfiles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(instanceprofilesResource, c.ns, name), &v1alpha1.InstanceProfile{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstanceProfiles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(instanceprofilesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstanceProfileList{})
	return err
}

// Patch applies the patch and returns the patched instanceProfile.
func (c *FakeInstanceProfiles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstanceProfile, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(instanceprofilesResource, c.ns, name, pt, data, subresources...), &v1alpha1.InstanceProfile{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceProfile), err
}
