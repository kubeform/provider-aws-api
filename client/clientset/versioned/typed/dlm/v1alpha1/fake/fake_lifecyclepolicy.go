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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dlm/v1alpha1"
)

// FakeLifecyclePolicies implements LifecyclePolicyInterface
type FakeLifecyclePolicies struct {
	Fake *FakeDlmV1alpha1
	ns   string
}

var lifecyclepoliciesResource = schema.GroupVersionResource{Group: "dlm.aws.kubeform.com", Version: "v1alpha1", Resource: "lifecyclepolicies"}

var lifecyclepoliciesKind = schema.GroupVersionKind{Group: "dlm.aws.kubeform.com", Version: "v1alpha1", Kind: "LifecyclePolicy"}

// Get takes name of the lifecyclePolicy, and returns the corresponding lifecyclePolicy object, and an error if there is any.
func (c *FakeLifecyclePolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LifecyclePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lifecyclepoliciesResource, c.ns, name), &v1alpha1.LifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecyclePolicy), err
}

// List takes label and field selectors, and returns the list of LifecyclePolicies that match those selectors.
func (c *FakeLifecyclePolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LifecyclePolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lifecyclepoliciesResource, lifecyclepoliciesKind, c.ns, opts), &v1alpha1.LifecyclePolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LifecyclePolicyList{ListMeta: obj.(*v1alpha1.LifecyclePolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.LifecyclePolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lifecyclePolicies.
func (c *FakeLifecyclePolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lifecyclepoliciesResource, c.ns, opts))

}

// Create takes the representation of a lifecyclePolicy and creates it.  Returns the server's representation of the lifecyclePolicy, and an error, if there is any.
func (c *FakeLifecyclePolicies) Create(ctx context.Context, lifecyclePolicy *v1alpha1.LifecyclePolicy, opts v1.CreateOptions) (result *v1alpha1.LifecyclePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lifecyclepoliciesResource, c.ns, lifecyclePolicy), &v1alpha1.LifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecyclePolicy), err
}

// Update takes the representation of a lifecyclePolicy and updates it. Returns the server's representation of the lifecyclePolicy, and an error, if there is any.
func (c *FakeLifecyclePolicies) Update(ctx context.Context, lifecyclePolicy *v1alpha1.LifecyclePolicy, opts v1.UpdateOptions) (result *v1alpha1.LifecyclePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lifecyclepoliciesResource, c.ns, lifecyclePolicy), &v1alpha1.LifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecyclePolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLifecyclePolicies) UpdateStatus(ctx context.Context, lifecyclePolicy *v1alpha1.LifecyclePolicy, opts v1.UpdateOptions) (*v1alpha1.LifecyclePolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lifecyclepoliciesResource, "status", c.ns, lifecyclePolicy), &v1alpha1.LifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecyclePolicy), err
}

// Delete takes name of the lifecyclePolicy and deletes it. Returns an error if one occurs.
func (c *FakeLifecyclePolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lifecyclepoliciesResource, c.ns, name), &v1alpha1.LifecyclePolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLifecyclePolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lifecyclepoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LifecyclePolicyList{})
	return err
}

// Patch applies the patch and returns the patched lifecyclePolicy.
func (c *FakeLifecyclePolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LifecyclePolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lifecyclepoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LifecyclePolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecyclePolicy), err
}
