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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/amplify/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBackendEnvironments implements BackendEnvironmentInterface
type FakeBackendEnvironments struct {
	Fake *FakeAmplifyV1alpha1
	ns   string
}

var backendenvironmentsResource = schema.GroupVersionResource{Group: "amplify.aws.kubeform.com", Version: "v1alpha1", Resource: "backendenvironments"}

var backendenvironmentsKind = schema.GroupVersionKind{Group: "amplify.aws.kubeform.com", Version: "v1alpha1", Kind: "BackendEnvironment"}

// Get takes name of the backendEnvironment, and returns the corresponding backendEnvironment object, and an error if there is any.
func (c *FakeBackendEnvironments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackendEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(backendenvironmentsResource, c.ns, name), &v1alpha1.BackendEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendEnvironment), err
}

// List takes label and field selectors, and returns the list of BackendEnvironments that match those selectors.
func (c *FakeBackendEnvironments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackendEnvironmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(backendenvironmentsResource, backendenvironmentsKind, c.ns, opts), &v1alpha1.BackendEnvironmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BackendEnvironmentList{ListMeta: obj.(*v1alpha1.BackendEnvironmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.BackendEnvironmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested backendEnvironments.
func (c *FakeBackendEnvironments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(backendenvironmentsResource, c.ns, opts))

}

// Create takes the representation of a backendEnvironment and creates it.  Returns the server's representation of the backendEnvironment, and an error, if there is any.
func (c *FakeBackendEnvironments) Create(ctx context.Context, backendEnvironment *v1alpha1.BackendEnvironment, opts v1.CreateOptions) (result *v1alpha1.BackendEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(backendenvironmentsResource, c.ns, backendEnvironment), &v1alpha1.BackendEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendEnvironment), err
}

// Update takes the representation of a backendEnvironment and updates it. Returns the server's representation of the backendEnvironment, and an error, if there is any.
func (c *FakeBackendEnvironments) Update(ctx context.Context, backendEnvironment *v1alpha1.BackendEnvironment, opts v1.UpdateOptions) (result *v1alpha1.BackendEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(backendenvironmentsResource, c.ns, backendEnvironment), &v1alpha1.BackendEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendEnvironment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBackendEnvironments) UpdateStatus(ctx context.Context, backendEnvironment *v1alpha1.BackendEnvironment, opts v1.UpdateOptions) (*v1alpha1.BackendEnvironment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(backendenvironmentsResource, "status", c.ns, backendEnvironment), &v1alpha1.BackendEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendEnvironment), err
}

// Delete takes name of the backendEnvironment and deletes it. Returns an error if one occurs.
func (c *FakeBackendEnvironments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(backendenvironmentsResource, c.ns, name), &v1alpha1.BackendEnvironment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBackendEnvironments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(backendenvironmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BackendEnvironmentList{})
	return err
}

// Patch applies the patch and returns the patched backendEnvironment.
func (c *FakeBackendEnvironments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackendEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(backendenvironmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BackendEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BackendEnvironment), err
}
