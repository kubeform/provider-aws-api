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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeObjectCopies implements ObjectCopyInterface
type FakeObjectCopies struct {
	Fake *FakeS3V1alpha1
	ns   string
}

var objectcopiesResource = schema.GroupVersionResource{Group: "s3.aws.kubeform.com", Version: "v1alpha1", Resource: "objectcopies"}

var objectcopiesKind = schema.GroupVersionKind{Group: "s3.aws.kubeform.com", Version: "v1alpha1", Kind: "ObjectCopy"}

// Get takes name of the objectCopy, and returns the corresponding objectCopy object, and an error if there is any.
func (c *FakeObjectCopies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ObjectCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(objectcopiesResource, c.ns, name), &v1alpha1.ObjectCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectCopy), err
}

// List takes label and field selectors, and returns the list of ObjectCopies that match those selectors.
func (c *FakeObjectCopies) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ObjectCopyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(objectcopiesResource, objectcopiesKind, c.ns, opts), &v1alpha1.ObjectCopyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ObjectCopyList{ListMeta: obj.(*v1alpha1.ObjectCopyList).ListMeta}
	for _, item := range obj.(*v1alpha1.ObjectCopyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested objectCopies.
func (c *FakeObjectCopies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(objectcopiesResource, c.ns, opts))

}

// Create takes the representation of a objectCopy and creates it.  Returns the server's representation of the objectCopy, and an error, if there is any.
func (c *FakeObjectCopies) Create(ctx context.Context, objectCopy *v1alpha1.ObjectCopy, opts v1.CreateOptions) (result *v1alpha1.ObjectCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(objectcopiesResource, c.ns, objectCopy), &v1alpha1.ObjectCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectCopy), err
}

// Update takes the representation of a objectCopy and updates it. Returns the server's representation of the objectCopy, and an error, if there is any.
func (c *FakeObjectCopies) Update(ctx context.Context, objectCopy *v1alpha1.ObjectCopy, opts v1.UpdateOptions) (result *v1alpha1.ObjectCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(objectcopiesResource, c.ns, objectCopy), &v1alpha1.ObjectCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectCopy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeObjectCopies) UpdateStatus(ctx context.Context, objectCopy *v1alpha1.ObjectCopy, opts v1.UpdateOptions) (*v1alpha1.ObjectCopy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(objectcopiesResource, "status", c.ns, objectCopy), &v1alpha1.ObjectCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectCopy), err
}

// Delete takes name of the objectCopy and deletes it. Returns an error if one occurs.
func (c *FakeObjectCopies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(objectcopiesResource, c.ns, name), &v1alpha1.ObjectCopy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeObjectCopies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(objectcopiesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ObjectCopyList{})
	return err
}

// Patch applies the patch and returns the patched objectCopy.
func (c *FakeObjectCopies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ObjectCopy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(objectcopiesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ObjectCopy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ObjectCopy), err
}
