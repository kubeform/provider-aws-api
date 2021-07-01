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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/codeartifact/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDomains implements DomainInterface
type FakeDomains struct {
	Fake *FakeCodeartifactV1alpha1
	ns   string
}

var domainsResource = schema.GroupVersionResource{Group: "codeartifact.aws.kubeform.com", Version: "v1alpha1", Resource: "domains"}

var domainsKind = schema.GroupVersionKind{Group: "codeartifact.aws.kubeform.com", Version: "v1alpha1", Kind: "Domain"}

// Get takes name of the domain, and returns the corresponding domain object, and an error if there is any.
func (c *FakeDomains) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Domain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(domainsResource, c.ns, name), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}

// List takes label and field selectors, and returns the list of Domains that match those selectors.
func (c *FakeDomains) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(domainsResource, domainsKind, c.ns, opts), &v1alpha1.DomainList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DomainList{ListMeta: obj.(*v1alpha1.DomainList).ListMeta}
	for _, item := range obj.(*v1alpha1.DomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested domains.
func (c *FakeDomains) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(domainsResource, c.ns, opts))

}

// Create takes the representation of a domain and creates it.  Returns the server's representation of the domain, and an error, if there is any.
func (c *FakeDomains) Create(ctx context.Context, domain *v1alpha1.Domain, opts v1.CreateOptions) (result *v1alpha1.Domain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(domainsResource, c.ns, domain), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}

// Update takes the representation of a domain and updates it. Returns the server's representation of the domain, and an error, if there is any.
func (c *FakeDomains) Update(ctx context.Context, domain *v1alpha1.Domain, opts v1.UpdateOptions) (result *v1alpha1.Domain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(domainsResource, c.ns, domain), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDomains) UpdateStatus(ctx context.Context, domain *v1alpha1.Domain, opts v1.UpdateOptions) (*v1alpha1.Domain, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(domainsResource, "status", c.ns, domain), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}

// Delete takes name of the domain and deletes it. Returns an error if one occurs.
func (c *FakeDomains) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(domainsResource, c.ns, name), &v1alpha1.Domain{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDomains) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(domainsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DomainList{})
	return err
}

// Patch applies the patch and returns the patched domain.
func (c *FakeDomains) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Domain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(domainsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}
