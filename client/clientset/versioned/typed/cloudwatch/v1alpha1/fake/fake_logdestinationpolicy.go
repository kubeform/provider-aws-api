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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudwatch/v1alpha1"
)

// FakeLogDestinationPolicies implements LogDestinationPolicyInterface
type FakeLogDestinationPolicies struct {
	Fake *FakeCloudwatchV1alpha1
	ns   string
}

var logdestinationpoliciesResource = schema.GroupVersionResource{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Resource: "logdestinationpolicies"}

var logdestinationpoliciesKind = schema.GroupVersionKind{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Kind: "LogDestinationPolicy"}

// Get takes name of the logDestinationPolicy, and returns the corresponding logDestinationPolicy object, and an error if there is any.
func (c *FakeLogDestinationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LogDestinationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logdestinationpoliciesResource, c.ns, name), &v1alpha1.LogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestinationPolicy), err
}

// List takes label and field selectors, and returns the list of LogDestinationPolicies that match those selectors.
func (c *FakeLogDestinationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogDestinationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logdestinationpoliciesResource, logdestinationpoliciesKind, c.ns, opts), &v1alpha1.LogDestinationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogDestinationPolicyList{ListMeta: obj.(*v1alpha1.LogDestinationPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogDestinationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logDestinationPolicies.
func (c *FakeLogDestinationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logdestinationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a logDestinationPolicy and creates it.  Returns the server's representation of the logDestinationPolicy, and an error, if there is any.
func (c *FakeLogDestinationPolicies) Create(ctx context.Context, logDestinationPolicy *v1alpha1.LogDestinationPolicy, opts v1.CreateOptions) (result *v1alpha1.LogDestinationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logdestinationpoliciesResource, c.ns, logDestinationPolicy), &v1alpha1.LogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestinationPolicy), err
}

// Update takes the representation of a logDestinationPolicy and updates it. Returns the server's representation of the logDestinationPolicy, and an error, if there is any.
func (c *FakeLogDestinationPolicies) Update(ctx context.Context, logDestinationPolicy *v1alpha1.LogDestinationPolicy, opts v1.UpdateOptions) (result *v1alpha1.LogDestinationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logdestinationpoliciesResource, c.ns, logDestinationPolicy), &v1alpha1.LogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestinationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogDestinationPolicies) UpdateStatus(ctx context.Context, logDestinationPolicy *v1alpha1.LogDestinationPolicy, opts v1.UpdateOptions) (*v1alpha1.LogDestinationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logdestinationpoliciesResource, "status", c.ns, logDestinationPolicy), &v1alpha1.LogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestinationPolicy), err
}

// Delete takes name of the logDestinationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeLogDestinationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(logdestinationpoliciesResource, c.ns, name), &v1alpha1.LogDestinationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogDestinationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logdestinationpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogDestinationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched logDestinationPolicy.
func (c *FakeLogDestinationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LogDestinationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logdestinationpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogDestinationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogDestinationPolicy), err
}
