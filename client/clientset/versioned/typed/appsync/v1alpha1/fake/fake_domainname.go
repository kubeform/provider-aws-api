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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/appsync/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDomainNames implements DomainNameInterface
type FakeDomainNames struct {
	Fake *FakeAppsyncV1alpha1
	ns   string
}

var domainnamesResource = schema.GroupVersionResource{Group: "appsync.aws.kubeform.com", Version: "v1alpha1", Resource: "domainnames"}

var domainnamesKind = schema.GroupVersionKind{Group: "appsync.aws.kubeform.com", Version: "v1alpha1", Kind: "DomainName"}

// Get takes name of the domainName, and returns the corresponding domainName object, and an error if there is any.
func (c *FakeDomainNames) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DomainName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(domainnamesResource, c.ns, name), &v1alpha1.DomainName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainName), err
}

// List takes label and field selectors, and returns the list of DomainNames that match those selectors.
func (c *FakeDomainNames) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DomainNameList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(domainnamesResource, domainnamesKind, c.ns, opts), &v1alpha1.DomainNameList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DomainNameList{ListMeta: obj.(*v1alpha1.DomainNameList).ListMeta}
	for _, item := range obj.(*v1alpha1.DomainNameList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested domainNames.
func (c *FakeDomainNames) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(domainnamesResource, c.ns, opts))

}

// Create takes the representation of a domainName and creates it.  Returns the server's representation of the domainName, and an error, if there is any.
func (c *FakeDomainNames) Create(ctx context.Context, domainName *v1alpha1.DomainName, opts v1.CreateOptions) (result *v1alpha1.DomainName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(domainnamesResource, c.ns, domainName), &v1alpha1.DomainName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainName), err
}

// Update takes the representation of a domainName and updates it. Returns the server's representation of the domainName, and an error, if there is any.
func (c *FakeDomainNames) Update(ctx context.Context, domainName *v1alpha1.DomainName, opts v1.UpdateOptions) (result *v1alpha1.DomainName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(domainnamesResource, c.ns, domainName), &v1alpha1.DomainName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainName), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDomainNames) UpdateStatus(ctx context.Context, domainName *v1alpha1.DomainName, opts v1.UpdateOptions) (*v1alpha1.DomainName, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(domainnamesResource, "status", c.ns, domainName), &v1alpha1.DomainName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainName), err
}

// Delete takes name of the domainName and deletes it. Returns an error if one occurs.
func (c *FakeDomainNames) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(domainnamesResource, c.ns, name), &v1alpha1.DomainName{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDomainNames) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(domainnamesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DomainNameList{})
	return err
}

// Patch applies the patch and returns the patched domainName.
func (c *FakeDomainNames) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DomainName, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(domainnamesResource, c.ns, name, pt, data, subresources...), &v1alpha1.DomainName{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainName), err
}
