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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/provider-aws-api/apis/secretsmanager/v1alpha1"
)

// FakeSecretRotations implements SecretRotationInterface
type FakeSecretRotations struct {
	Fake *FakeSecretsmanagerV1alpha1
	ns   string
}

var secretrotationsResource = schema.GroupVersionResource{Group: "secretsmanager.aws.kubeform.com", Version: "v1alpha1", Resource: "secretrotations"}

var secretrotationsKind = schema.GroupVersionKind{Group: "secretsmanager.aws.kubeform.com", Version: "v1alpha1", Kind: "SecretRotation"}

// Get takes name of the secretRotation, and returns the corresponding secretRotation object, and an error if there is any.
func (c *FakeSecretRotations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SecretRotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(secretrotationsResource, c.ns, name), &v1alpha1.SecretRotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretRotation), err
}

// List takes label and field selectors, and returns the list of SecretRotations that match those selectors.
func (c *FakeSecretRotations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SecretRotationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(secretrotationsResource, secretrotationsKind, c.ns, opts), &v1alpha1.SecretRotationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SecretRotationList{ListMeta: obj.(*v1alpha1.SecretRotationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SecretRotationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested secretRotations.
func (c *FakeSecretRotations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(secretrotationsResource, c.ns, opts))

}

// Create takes the representation of a secretRotation and creates it.  Returns the server's representation of the secretRotation, and an error, if there is any.
func (c *FakeSecretRotations) Create(ctx context.Context, secretRotation *v1alpha1.SecretRotation, opts v1.CreateOptions) (result *v1alpha1.SecretRotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(secretrotationsResource, c.ns, secretRotation), &v1alpha1.SecretRotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretRotation), err
}

// Update takes the representation of a secretRotation and updates it. Returns the server's representation of the secretRotation, and an error, if there is any.
func (c *FakeSecretRotations) Update(ctx context.Context, secretRotation *v1alpha1.SecretRotation, opts v1.UpdateOptions) (result *v1alpha1.SecretRotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(secretrotationsResource, c.ns, secretRotation), &v1alpha1.SecretRotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretRotation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSecretRotations) UpdateStatus(ctx context.Context, secretRotation *v1alpha1.SecretRotation, opts v1.UpdateOptions) (*v1alpha1.SecretRotation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(secretrotationsResource, "status", c.ns, secretRotation), &v1alpha1.SecretRotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretRotation), err
}

// Delete takes name of the secretRotation and deletes it. Returns an error if one occurs.
func (c *FakeSecretRotations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(secretrotationsResource, c.ns, name), &v1alpha1.SecretRotation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSecretRotations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(secretrotationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SecretRotationList{})
	return err
}

// Patch applies the patch and returns the patched secretRotation.
func (c *FakeSecretRotations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SecretRotation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(secretrotationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SecretRotation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SecretRotation), err
}
