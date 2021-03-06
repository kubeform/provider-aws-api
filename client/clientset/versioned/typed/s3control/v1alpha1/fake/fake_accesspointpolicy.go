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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/s3control/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAccessPointPolicies implements AccessPointPolicyInterface
type FakeAccessPointPolicies struct {
	Fake *FakeS3controlV1alpha1
	ns   string
}

var accesspointpoliciesResource = schema.GroupVersionResource{Group: "s3control.aws.kubeform.com", Version: "v1alpha1", Resource: "accesspointpolicies"}

var accesspointpoliciesKind = schema.GroupVersionKind{Group: "s3control.aws.kubeform.com", Version: "v1alpha1", Kind: "AccessPointPolicy"}

// Get takes name of the accessPointPolicy, and returns the corresponding accessPointPolicy object, and an error if there is any.
func (c *FakeAccessPointPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccessPointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accesspointpoliciesResource, c.ns, name), &v1alpha1.AccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPointPolicy), err
}

// List takes label and field selectors, and returns the list of AccessPointPolicies that match those selectors.
func (c *FakeAccessPointPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccessPointPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accesspointpoliciesResource, accesspointpoliciesKind, c.ns, opts), &v1alpha1.AccessPointPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccessPointPolicyList{ListMeta: obj.(*v1alpha1.AccessPointPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccessPointPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accessPointPolicies.
func (c *FakeAccessPointPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accesspointpoliciesResource, c.ns, opts))

}

// Create takes the representation of a accessPointPolicy and creates it.  Returns the server's representation of the accessPointPolicy, and an error, if there is any.
func (c *FakeAccessPointPolicies) Create(ctx context.Context, accessPointPolicy *v1alpha1.AccessPointPolicy, opts v1.CreateOptions) (result *v1alpha1.AccessPointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accesspointpoliciesResource, c.ns, accessPointPolicy), &v1alpha1.AccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPointPolicy), err
}

// Update takes the representation of a accessPointPolicy and updates it. Returns the server's representation of the accessPointPolicy, and an error, if there is any.
func (c *FakeAccessPointPolicies) Update(ctx context.Context, accessPointPolicy *v1alpha1.AccessPointPolicy, opts v1.UpdateOptions) (result *v1alpha1.AccessPointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accesspointpoliciesResource, c.ns, accessPointPolicy), &v1alpha1.AccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPointPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccessPointPolicies) UpdateStatus(ctx context.Context, accessPointPolicy *v1alpha1.AccessPointPolicy, opts v1.UpdateOptions) (*v1alpha1.AccessPointPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accesspointpoliciesResource, "status", c.ns, accessPointPolicy), &v1alpha1.AccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPointPolicy), err
}

// Delete takes name of the accessPointPolicy and deletes it. Returns an error if one occurs.
func (c *FakeAccessPointPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accesspointpoliciesResource, c.ns, name), &v1alpha1.AccessPointPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccessPointPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accesspointpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccessPointPolicyList{})
	return err
}

// Patch applies the patch and returns the patched accessPointPolicy.
func (c *FakeAccessPointPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccessPointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accesspointpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccessPointPolicy), err
}
