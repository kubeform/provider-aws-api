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

// FakeServiceSpecificCredentials implements ServiceSpecificCredentialInterface
type FakeServiceSpecificCredentials struct {
	Fake *FakeIamV1alpha1
	ns   string
}

var servicespecificcredentialsResource = schema.GroupVersionResource{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Resource: "servicespecificcredentials"}

var servicespecificcredentialsKind = schema.GroupVersionKind{Group: "iam.aws.kubeform.com", Version: "v1alpha1", Kind: "ServiceSpecificCredential"}

// Get takes name of the serviceSpecificCredential, and returns the corresponding serviceSpecificCredential object, and an error if there is any.
func (c *FakeServiceSpecificCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceSpecificCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicespecificcredentialsResource, c.ns, name), &v1alpha1.ServiceSpecificCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceSpecificCredential), err
}

// List takes label and field selectors, and returns the list of ServiceSpecificCredentials that match those selectors.
func (c *FakeServiceSpecificCredentials) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceSpecificCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicespecificcredentialsResource, servicespecificcredentialsKind, c.ns, opts), &v1alpha1.ServiceSpecificCredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceSpecificCredentialList{ListMeta: obj.(*v1alpha1.ServiceSpecificCredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceSpecificCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceSpecificCredentials.
func (c *FakeServiceSpecificCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicespecificcredentialsResource, c.ns, opts))

}

// Create takes the representation of a serviceSpecificCredential and creates it.  Returns the server's representation of the serviceSpecificCredential, and an error, if there is any.
func (c *FakeServiceSpecificCredentials) Create(ctx context.Context, serviceSpecificCredential *v1alpha1.ServiceSpecificCredential, opts v1.CreateOptions) (result *v1alpha1.ServiceSpecificCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicespecificcredentialsResource, c.ns, serviceSpecificCredential), &v1alpha1.ServiceSpecificCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceSpecificCredential), err
}

// Update takes the representation of a serviceSpecificCredential and updates it. Returns the server's representation of the serviceSpecificCredential, and an error, if there is any.
func (c *FakeServiceSpecificCredentials) Update(ctx context.Context, serviceSpecificCredential *v1alpha1.ServiceSpecificCredential, opts v1.UpdateOptions) (result *v1alpha1.ServiceSpecificCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicespecificcredentialsResource, c.ns, serviceSpecificCredential), &v1alpha1.ServiceSpecificCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceSpecificCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceSpecificCredentials) UpdateStatus(ctx context.Context, serviceSpecificCredential *v1alpha1.ServiceSpecificCredential, opts v1.UpdateOptions) (*v1alpha1.ServiceSpecificCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicespecificcredentialsResource, "status", c.ns, serviceSpecificCredential), &v1alpha1.ServiceSpecificCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceSpecificCredential), err
}

// Delete takes name of the serviceSpecificCredential and deletes it. Returns an error if one occurs.
func (c *FakeServiceSpecificCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicespecificcredentialsResource, c.ns, name), &v1alpha1.ServiceSpecificCredential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceSpecificCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicespecificcredentialsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceSpecificCredentialList{})
	return err
}

// Patch applies the patch and returns the patched serviceSpecificCredential.
func (c *FakeServiceSpecificCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceSpecificCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicespecificcredentialsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceSpecificCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceSpecificCredential), err
}
