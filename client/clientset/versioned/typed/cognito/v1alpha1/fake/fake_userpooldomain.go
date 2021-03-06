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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cognito/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUserPoolDomains implements UserPoolDomainInterface
type FakeUserPoolDomains struct {
	Fake *FakeCognitoV1alpha1
	ns   string
}

var userpooldomainsResource = schema.GroupVersionResource{Group: "cognito.aws.kubeform.com", Version: "v1alpha1", Resource: "userpooldomains"}

var userpooldomainsKind = schema.GroupVersionKind{Group: "cognito.aws.kubeform.com", Version: "v1alpha1", Kind: "UserPoolDomain"}

// Get takes name of the userPoolDomain, and returns the corresponding userPoolDomain object, and an error if there is any.
func (c *FakeUserPoolDomains) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(userpooldomainsResource, c.ns, name), &v1alpha1.UserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolDomain), err
}

// List takes label and field selectors, and returns the list of UserPoolDomains that match those selectors.
func (c *FakeUserPoolDomains) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UserPoolDomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(userpooldomainsResource, userpooldomainsKind, c.ns, opts), &v1alpha1.UserPoolDomainList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UserPoolDomainList{ListMeta: obj.(*v1alpha1.UserPoolDomainList).ListMeta}
	for _, item := range obj.(*v1alpha1.UserPoolDomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userPoolDomains.
func (c *FakeUserPoolDomains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(userpooldomainsResource, c.ns, opts))

}

// Create takes the representation of a userPoolDomain and creates it.  Returns the server's representation of the userPoolDomain, and an error, if there is any.
func (c *FakeUserPoolDomains) Create(ctx context.Context, userPoolDomain *v1alpha1.UserPoolDomain, opts v1.CreateOptions) (result *v1alpha1.UserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(userpooldomainsResource, c.ns, userPoolDomain), &v1alpha1.UserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolDomain), err
}

// Update takes the representation of a userPoolDomain and updates it. Returns the server's representation of the userPoolDomain, and an error, if there is any.
func (c *FakeUserPoolDomains) Update(ctx context.Context, userPoolDomain *v1alpha1.UserPoolDomain, opts v1.UpdateOptions) (result *v1alpha1.UserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(userpooldomainsResource, c.ns, userPoolDomain), &v1alpha1.UserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolDomain), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserPoolDomains) UpdateStatus(ctx context.Context, userPoolDomain *v1alpha1.UserPoolDomain, opts v1.UpdateOptions) (*v1alpha1.UserPoolDomain, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(userpooldomainsResource, "status", c.ns, userPoolDomain), &v1alpha1.UserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolDomain), err
}

// Delete takes name of the userPoolDomain and deletes it. Returns an error if one occurs.
func (c *FakeUserPoolDomains) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(userpooldomainsResource, c.ns, name), &v1alpha1.UserPoolDomain{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserPoolDomains) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(userpooldomainsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UserPoolDomainList{})
	return err
}

// Patch applies the patch and returns the patched userPoolDomain.
func (c *FakeUserPoolDomains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UserPoolDomain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(userpooldomainsResource, c.ns, name, pt, data, subresources...), &v1alpha1.UserPoolDomain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolDomain), err
}
