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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/msk/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScramSecretAssociations implements ScramSecretAssociationInterface
type FakeScramSecretAssociations struct {
	Fake *FakeMskV1alpha1
	ns   string
}

var scramsecretassociationsResource = schema.GroupVersionResource{Group: "msk.aws.kubeform.com", Version: "v1alpha1", Resource: "scramsecretassociations"}

var scramsecretassociationsKind = schema.GroupVersionKind{Group: "msk.aws.kubeform.com", Version: "v1alpha1", Kind: "ScramSecretAssociation"}

// Get takes name of the scramSecretAssociation, and returns the corresponding scramSecretAssociation object, and an error if there is any.
func (c *FakeScramSecretAssociations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScramSecretAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scramsecretassociationsResource, c.ns, name), &v1alpha1.ScramSecretAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScramSecretAssociation), err
}

// List takes label and field selectors, and returns the list of ScramSecretAssociations that match those selectors.
func (c *FakeScramSecretAssociations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScramSecretAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scramsecretassociationsResource, scramsecretassociationsKind, c.ns, opts), &v1alpha1.ScramSecretAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScramSecretAssociationList{ListMeta: obj.(*v1alpha1.ScramSecretAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScramSecretAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scramSecretAssociations.
func (c *FakeScramSecretAssociations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scramsecretassociationsResource, c.ns, opts))

}

// Create takes the representation of a scramSecretAssociation and creates it.  Returns the server's representation of the scramSecretAssociation, and an error, if there is any.
func (c *FakeScramSecretAssociations) Create(ctx context.Context, scramSecretAssociation *v1alpha1.ScramSecretAssociation, opts v1.CreateOptions) (result *v1alpha1.ScramSecretAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scramsecretassociationsResource, c.ns, scramSecretAssociation), &v1alpha1.ScramSecretAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScramSecretAssociation), err
}

// Update takes the representation of a scramSecretAssociation and updates it. Returns the server's representation of the scramSecretAssociation, and an error, if there is any.
func (c *FakeScramSecretAssociations) Update(ctx context.Context, scramSecretAssociation *v1alpha1.ScramSecretAssociation, opts v1.UpdateOptions) (result *v1alpha1.ScramSecretAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scramsecretassociationsResource, c.ns, scramSecretAssociation), &v1alpha1.ScramSecretAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScramSecretAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScramSecretAssociations) UpdateStatus(ctx context.Context, scramSecretAssociation *v1alpha1.ScramSecretAssociation, opts v1.UpdateOptions) (*v1alpha1.ScramSecretAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(scramsecretassociationsResource, "status", c.ns, scramSecretAssociation), &v1alpha1.ScramSecretAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScramSecretAssociation), err
}

// Delete takes name of the scramSecretAssociation and deletes it. Returns an error if one occurs.
func (c *FakeScramSecretAssociations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scramsecretassociationsResource, c.ns, name), &v1alpha1.ScramSecretAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScramSecretAssociations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scramsecretassociationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScramSecretAssociationList{})
	return err
}

// Patch applies the patch and returns the patched scramSecretAssociation.
func (c *FakeScramSecretAssociations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScramSecretAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scramsecretassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ScramSecretAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScramSecretAssociation), err
}
