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

// FakeDomainNameAPIAssociations implements DomainNameAPIAssociationInterface
type FakeDomainNameAPIAssociations struct {
	Fake *FakeAppsyncV1alpha1
	ns   string
}

var domainnameapiassociationsResource = schema.GroupVersionResource{Group: "appsync.aws.kubeform.com", Version: "v1alpha1", Resource: "domainnameapiassociations"}

var domainnameapiassociationsKind = schema.GroupVersionKind{Group: "appsync.aws.kubeform.com", Version: "v1alpha1", Kind: "DomainNameAPIAssociation"}

// Get takes name of the domainNameAPIAssociation, and returns the corresponding domainNameAPIAssociation object, and an error if there is any.
func (c *FakeDomainNameAPIAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DomainNameAPIAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(domainnameapiassociationsResource, c.ns, name), &v1alpha1.DomainNameAPIAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainNameAPIAssociation), err
}

// List takes label and field selectors, and returns the list of DomainNameAPIAssociations that match those selectors.
func (c *FakeDomainNameAPIAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DomainNameAPIAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(domainnameapiassociationsResource, domainnameapiassociationsKind, c.ns, opts), &v1alpha1.DomainNameAPIAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DomainNameAPIAssociationList{ListMeta: obj.(*v1alpha1.DomainNameAPIAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.DomainNameAPIAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested domainNameAPIAssociations.
func (c *FakeDomainNameAPIAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(domainnameapiassociationsResource, c.ns, opts))

}

// Create takes the representation of a domainNameAPIAssociation and creates it.  Returns the server's representation of the domainNameAPIAssociation, and an error, if there is any.
func (c *FakeDomainNameAPIAssociations) Create(ctx context.Context, domainNameAPIAssociation *v1alpha1.DomainNameAPIAssociation, opts v1.CreateOptions) (result *v1alpha1.DomainNameAPIAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(domainnameapiassociationsResource, c.ns, domainNameAPIAssociation), &v1alpha1.DomainNameAPIAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainNameAPIAssociation), err
}

// Update takes the representation of a domainNameAPIAssociation and updates it. Returns the server's representation of the domainNameAPIAssociation, and an error, if there is any.
func (c *FakeDomainNameAPIAssociations) Update(ctx context.Context, domainNameAPIAssociation *v1alpha1.DomainNameAPIAssociation, opts v1.UpdateOptions) (result *v1alpha1.DomainNameAPIAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(domainnameapiassociationsResource, c.ns, domainNameAPIAssociation), &v1alpha1.DomainNameAPIAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainNameAPIAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDomainNameAPIAssociations) UpdateStatus(ctx context.Context, domainNameAPIAssociation *v1alpha1.DomainNameAPIAssociation, opts v1.UpdateOptions) (*v1alpha1.DomainNameAPIAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(domainnameapiassociationsResource, "status", c.ns, domainNameAPIAssociation), &v1alpha1.DomainNameAPIAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainNameAPIAssociation), err
}

// Delete takes name of the domainNameAPIAssociation and deletes it. Returns an error if one occurs.
func (c *FakeDomainNameAPIAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(domainnameapiassociationsResource, c.ns, name), &v1alpha1.DomainNameAPIAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDomainNameAPIAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(domainnameapiassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DomainNameAPIAssociationList{})
	return err
}

// Patch applies the patch and returns the patched domainNameAPIAssociation.
func (c *FakeDomainNameAPIAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DomainNameAPIAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(domainnameapiassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DomainNameAPIAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainNameAPIAssociation), err
}
