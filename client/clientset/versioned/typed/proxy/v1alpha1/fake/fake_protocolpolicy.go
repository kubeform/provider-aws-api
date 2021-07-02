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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/proxy/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProtocolPolicies implements ProtocolPolicyInterface
type FakeProtocolPolicies struct {
	Fake *FakeProxyV1alpha1
	ns   string
}

var protocolpoliciesResource = schema.GroupVersionResource{Group: "proxy.aws.kubeform.com", Version: "v1alpha1", Resource: "protocolpolicies"}

var protocolpoliciesKind = schema.GroupVersionKind{Group: "proxy.aws.kubeform.com", Version: "v1alpha1", Kind: "ProtocolPolicy"}

// Get takes name of the protocolPolicy, and returns the corresponding protocolPolicy object, and an error if there is any.
func (c *FakeProtocolPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ProtocolPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(protocolpoliciesResource, c.ns, name), &v1alpha1.ProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtocolPolicy), err
}

// List takes label and field selectors, and returns the list of ProtocolPolicies that match those selectors.
func (c *FakeProtocolPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ProtocolPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(protocolpoliciesResource, protocolpoliciesKind, c.ns, opts), &v1alpha1.ProtocolPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ProtocolPolicyList{ListMeta: obj.(*v1alpha1.ProtocolPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ProtocolPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested protocolPolicies.
func (c *FakeProtocolPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(protocolpoliciesResource, c.ns, opts))

}

// Create takes the representation of a protocolPolicy and creates it.  Returns the server's representation of the protocolPolicy, and an error, if there is any.
func (c *FakeProtocolPolicies) Create(ctx context.Context, protocolPolicy *v1alpha1.ProtocolPolicy, opts v1.CreateOptions) (result *v1alpha1.ProtocolPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(protocolpoliciesResource, c.ns, protocolPolicy), &v1alpha1.ProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtocolPolicy), err
}

// Update takes the representation of a protocolPolicy and updates it. Returns the server's representation of the protocolPolicy, and an error, if there is any.
func (c *FakeProtocolPolicies) Update(ctx context.Context, protocolPolicy *v1alpha1.ProtocolPolicy, opts v1.UpdateOptions) (result *v1alpha1.ProtocolPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(protocolpoliciesResource, c.ns, protocolPolicy), &v1alpha1.ProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtocolPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProtocolPolicies) UpdateStatus(ctx context.Context, protocolPolicy *v1alpha1.ProtocolPolicy, opts v1.UpdateOptions) (*v1alpha1.ProtocolPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(protocolpoliciesResource, "status", c.ns, protocolPolicy), &v1alpha1.ProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtocolPolicy), err
}

// Delete takes name of the protocolPolicy and deletes it. Returns an error if one occurs.
func (c *FakeProtocolPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(protocolpoliciesResource, c.ns, name), &v1alpha1.ProtocolPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProtocolPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(protocolpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ProtocolPolicyList{})
	return err
}

// Patch applies the patch and returns the patched protocolPolicy.
func (c *FakeProtocolPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ProtocolPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(protocolpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ProtocolPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ProtocolPolicy), err
}
