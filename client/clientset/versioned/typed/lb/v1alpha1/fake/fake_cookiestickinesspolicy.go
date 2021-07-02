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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/lb/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCookieStickinessPolicies implements CookieStickinessPolicyInterface
type FakeCookieStickinessPolicies struct {
	Fake *FakeLbV1alpha1
	ns   string
}

var cookiestickinesspoliciesResource = schema.GroupVersionResource{Group: "lb.aws.kubeform.com", Version: "v1alpha1", Resource: "cookiestickinesspolicies"}

var cookiestickinesspoliciesKind = schema.GroupVersionKind{Group: "lb.aws.kubeform.com", Version: "v1alpha1", Kind: "CookieStickinessPolicy"}

// Get takes name of the cookieStickinessPolicy, and returns the corresponding cookieStickinessPolicy object, and an error if there is any.
func (c *FakeCookieStickinessPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cookiestickinesspoliciesResource, c.ns, name), &v1alpha1.CookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CookieStickinessPolicy), err
}

// List takes label and field selectors, and returns the list of CookieStickinessPolicies that match those selectors.
func (c *FakeCookieStickinessPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CookieStickinessPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cookiestickinesspoliciesResource, cookiestickinesspoliciesKind, c.ns, opts), &v1alpha1.CookieStickinessPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CookieStickinessPolicyList{ListMeta: obj.(*v1alpha1.CookieStickinessPolicyList).ListMeta}
	for _, item := range obj.(*v1alpha1.CookieStickinessPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cookieStickinessPolicies.
func (c *FakeCookieStickinessPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cookiestickinesspoliciesResource, c.ns, opts))

}

// Create takes the representation of a cookieStickinessPolicy and creates it.  Returns the server's representation of the cookieStickinessPolicy, and an error, if there is any.
func (c *FakeCookieStickinessPolicies) Create(ctx context.Context, cookieStickinessPolicy *v1alpha1.CookieStickinessPolicy, opts v1.CreateOptions) (result *v1alpha1.CookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cookiestickinesspoliciesResource, c.ns, cookieStickinessPolicy), &v1alpha1.CookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CookieStickinessPolicy), err
}

// Update takes the representation of a cookieStickinessPolicy and updates it. Returns the server's representation of the cookieStickinessPolicy, and an error, if there is any.
func (c *FakeCookieStickinessPolicies) Update(ctx context.Context, cookieStickinessPolicy *v1alpha1.CookieStickinessPolicy, opts v1.UpdateOptions) (result *v1alpha1.CookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cookiestickinesspoliciesResource, c.ns, cookieStickinessPolicy), &v1alpha1.CookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CookieStickinessPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCookieStickinessPolicies) UpdateStatus(ctx context.Context, cookieStickinessPolicy *v1alpha1.CookieStickinessPolicy, opts v1.UpdateOptions) (*v1alpha1.CookieStickinessPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(cookiestickinesspoliciesResource, "status", c.ns, cookieStickinessPolicy), &v1alpha1.CookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CookieStickinessPolicy), err
}

// Delete takes name of the cookieStickinessPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCookieStickinessPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cookiestickinesspoliciesResource, c.ns, name), &v1alpha1.CookieStickinessPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCookieStickinessPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cookiestickinesspoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CookieStickinessPolicyList{})
	return err
}

// Patch applies the patch and returns the patched cookieStickinessPolicy.
func (c *FakeCookieStickinessPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CookieStickinessPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cookiestickinesspoliciesResource, c.ns, name, pt, data, subresources...), &v1alpha1.CookieStickinessPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CookieStickinessPolicy), err
}
