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

// FakeMultiRegionAccessPointPolicies implements MultiRegionAccessPointPolicyInterface
type FakeMultiRegionAccessPointPolicies struct {
	Fake *FakeS3controlV1alpha1
	ns   string
}

var multiregionaccesspointpoliciesResource = schema.GroupVersionResource{Group: "s3control.aws.kubeform.com", Version: "v1alpha1", Resource: "multiregionaccesspointpolicies"}

var multiregionaccesspointpoliciesKind = schema.GroupVersionKind{Group: "s3control.aws.kubeform.com", Version: "v1alpha1", Kind: "MultiRegionAccessPointPolicy"}

// Get takes name of the multiRegionAccessPointPolicy, and returns the corresponding multiRegionAccessPointPolicy object, and an error if there is any.
func (c *FakeMultiRegionAccessPointPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiRegionAccessPointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiregionaccesspointpoliciesResource, c.ns, name), &v1alpha1.MultiRegionAccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiRegionAccessPointPolicy), err
}

// List takes label and field selectors, and returns the list of MultiRegionAccessPointPolicies that match those selectors.
func (c *FakeMultiRegionAccessPointPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiRegionAccessPointPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiregionaccesspointpoliciesResource, multiregionaccesspointpoliciesKind, c.ns, opts), &v1alpha1.MultiRegionAccessPointPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MultiRegionAccessPointPolicyList{ListMeta: obj.(*v1alpha1.MultiRegionAccessPointPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.MultiRegionAccessPointPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiRegionAccessPointPolicies.
func (c *FakeMultiRegionAccessPointPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiregionaccesspointpoliciesResource, c.ns, opts))

}

// Create takes the representation of a multiRegionAccessPointPolicy and creates it.  Returns the server's representation of the multiRegionAccessPointPolicy, and an error, if there is any.
func (c *FakeMultiRegionAccessPointPolicies) Create(ctx context.Context, multiRegionAccessPointPolicy *v1alpha1.MultiRegionAccessPointPolicy, opts v1.CreateOptions) (result *v1alpha1.MultiRegionAccessPointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiregionaccesspointpoliciesResource, c.ns, multiRegionAccessPointPolicy), &v1alpha1.MultiRegionAccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiRegionAccessPointPolicy), err
}

// Update takes the representation of a multiRegionAccessPointPolicy and updates it. Returns the server's representation of the multiRegionAccessPointPolicy, and an error, if there is any.
func (c *FakeMultiRegionAccessPointPolicies) Update(ctx context.Context, multiRegionAccessPointPolicy *v1alpha1.MultiRegionAccessPointPolicy, opts v1.UpdateOptions) (result *v1alpha1.MultiRegionAccessPointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiregionaccesspointpoliciesResource, c.ns, multiRegionAccessPointPolicy), &v1alpha1.MultiRegionAccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiRegionAccessPointPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiRegionAccessPointPolicies) UpdateStatus(ctx context.Context, multiRegionAccessPointPolicy *v1alpha1.MultiRegionAccessPointPolicy, opts v1.UpdateOptions) (*v1alpha1.MultiRegionAccessPointPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiregionaccesspointpoliciesResource, "status", c.ns, multiRegionAccessPointPolicy), &v1alpha1.MultiRegionAccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiRegionAccessPointPolicy), err
}

// Delete takes name of the multiRegionAccessPointPolicy and deletes it. Returns an error if one occurs.
func (c *FakeMultiRegionAccessPointPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(multiregionaccesspointpoliciesResource, c.ns, name), &v1alpha1.MultiRegionAccessPointPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiRegionAccessPointPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiregionaccesspointpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MultiRegionAccessPointPolicyList{})
	return err
}

// Patch applies the patch and returns the patched multiRegionAccessPointPolicy.
func (c *FakeMultiRegionAccessPointPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiRegionAccessPointPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiregionaccesspointpoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MultiRegionAccessPointPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiRegionAccessPointPolicy), err
}
