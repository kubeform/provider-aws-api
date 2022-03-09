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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/iam/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSigningCertificates implements SigningCertificateInterface
type FakeSigningCertificates struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var signingcertificatesResource = schema.GroupVersionResource{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Resource: "signingcertificates"}

var signingcertificatesKind = schema.GroupVersionKind{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Kind: "SigningCertificate"}

// Get takes name of the signingCertificate, and returns the corresponding signingCertificate object, and an error if there is any.
func (c *FakeSigningCertificates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SigningCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(signingcertificatesResource, c.ns, name), &v1alpha1.SigningCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningCertificate), err
}

// List takes label and field selectors, and returns the list of SigningCertificates that match those selectors.
func (c *FakeSigningCertificates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SigningCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(signingcertificatesResource, signingcertificatesKind, c.ns, opts), &v1alpha1.SigningCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SigningCertificateList{ListMeta: obj.(*v1alpha1.SigningCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.SigningCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested signingCertificates.
func (c *FakeSigningCertificates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(signingcertificatesResource, c.ns, opts))

}

// Create takes the representation of a signingCertificate and creates it.  Returns the server's representation of the signingCertificate, and an error, if there is any.
func (c *FakeSigningCertificates) Create(ctx context.Context, signingCertificate *v1alpha1.SigningCertificate, opts v1.CreateOptions) (result *v1alpha1.SigningCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(signingcertificatesResource, c.ns, signingCertificate), &v1alpha1.SigningCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningCertificate), err
}

// Update takes the representation of a signingCertificate and updates it. Returns the server's representation of the signingCertificate, and an error, if there is any.
func (c *FakeSigningCertificates) Update(ctx context.Context, signingCertificate *v1alpha1.SigningCertificate, opts v1.UpdateOptions) (result *v1alpha1.SigningCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(signingcertificatesResource, c.ns, signingCertificate), &v1alpha1.SigningCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSigningCertificates) UpdateStatus(ctx context.Context, signingCertificate *v1alpha1.SigningCertificate, opts v1.UpdateOptions) (*v1alpha1.SigningCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(signingcertificatesResource, "status", c.ns, signingCertificate), &v1alpha1.SigningCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningCertificate), err
}

// Delete takes name of the signingCertificate and deletes it. Returns an error if one occurs.
func (c *FakeSigningCertificates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(signingcertificatesResource, c.ns, name), &v1alpha1.SigningCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSigningCertificates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(signingcertificatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SigningCertificateList{})
	return err
}

// Patch applies the patch and returns the patched signingCertificate.
func (c *FakeSigningCertificates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SigningCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(signingcertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SigningCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningCertificate), err
}
