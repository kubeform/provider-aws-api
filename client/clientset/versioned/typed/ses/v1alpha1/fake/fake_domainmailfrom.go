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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDomainMailFroms implements DomainMailFromInterface
type FakeDomainMailFroms struct {
	Fake *FakeSesV1alpha1
	ns   string
}

var domainmailfromsResource = schema.GroupVersionResource{Group: "ses.aws.kubeform.com", Version: "v1alpha1", Resource: "domainmailfroms"}

var domainmailfromsKind = schema.GroupVersionKind{Group: "ses.aws.kubeform.com", Version: "v1alpha1", Kind: "DomainMailFrom"}

// Get takes name of the domainMailFrom, and returns the corresponding domainMailFrom object, and an error if there is any.
func (c *FakeDomainMailFroms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DomainMailFrom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(domainmailfromsResource, c.ns, name), &v1alpha1.DomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainMailFrom), err
}

// List takes label and field selectors, and returns the list of DomainMailFroms that match those selectors.
func (c *FakeDomainMailFroms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DomainMailFromList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(domainmailfromsResource, domainmailfromsKind, c.ns, opts), &v1alpha1.DomainMailFromList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DomainMailFromList{ListMeta: obj.(*v1alpha1.DomainMailFromList).ListMeta}
	for _, item := range obj.(*v1alpha1.DomainMailFromList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested domainMailFroms.
func (c *FakeDomainMailFroms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(domainmailfromsResource, c.ns, opts))

}

// Create takes the representation of a domainMailFrom and creates it.  Returns the server's representation of the domainMailFrom, and an error, if there is any.
func (c *FakeDomainMailFroms) Create(ctx context.Context, domainMailFrom *v1alpha1.DomainMailFrom, opts v1.CreateOptions) (result *v1alpha1.DomainMailFrom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(domainmailfromsResource, c.ns, domainMailFrom), &v1alpha1.DomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainMailFrom), err
}

// Update takes the representation of a domainMailFrom and updates it. Returns the server's representation of the domainMailFrom, and an error, if there is any.
func (c *FakeDomainMailFroms) Update(ctx context.Context, domainMailFrom *v1alpha1.DomainMailFrom, opts v1.UpdateOptions) (result *v1alpha1.DomainMailFrom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(domainmailfromsResource, c.ns, domainMailFrom), &v1alpha1.DomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainMailFrom), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDomainMailFroms) UpdateStatus(ctx context.Context, domainMailFrom *v1alpha1.DomainMailFrom, opts v1.UpdateOptions) (*v1alpha1.DomainMailFrom, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(domainmailfromsResource, "status", c.ns, domainMailFrom), &v1alpha1.DomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainMailFrom), err
}

// Delete takes name of the domainMailFrom and deletes it. Returns an error if one occurs.
func (c *FakeDomainMailFroms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(domainmailfromsResource, c.ns, name), &v1alpha1.DomainMailFrom{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDomainMailFroms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(domainmailfromsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DomainMailFromList{})
	return err
}

// Patch applies the patch and returns the patched domainMailFrom.
func (c *FakeDomainMailFroms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DomainMailFrom, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(domainmailfromsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DomainMailFrom{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DomainMailFrom), err
}
