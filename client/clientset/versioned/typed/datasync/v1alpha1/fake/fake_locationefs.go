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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/datasync/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLocationEfses implements LocationEfsInterface
type FakeLocationEfses struct {
	Fake *FakeDatasyncV1alpha1
	ns   string
}

var locationefsesResource = schema.GroupVersionResource{Group: "datasync.aws.kubeform.com", Version: "v1alpha1", Resource: "locationefses"}

var locationefsesKind = schema.GroupVersionKind{Group: "datasync.aws.kubeform.com", Version: "v1alpha1", Kind: "LocationEfs"}

// Get takes name of the locationEfs, and returns the corresponding locationEfs object, and an error if there is any.
func (c *FakeLocationEfses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocationEfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(locationefsesResource, c.ns, name), &v1alpha1.LocationEfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationEfs), err
}

// List takes label and field selectors, and returns the list of LocationEfses that match those selectors.
func (c *FakeLocationEfses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocationEfsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(locationefsesResource, locationefsesKind, c.ns, opts), &v1alpha1.LocationEfsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LocationEfsList{ListMeta: obj.(*v1alpha1.LocationEfsList).ListMeta}
	for _, item := range obj.(*v1alpha1.LocationEfsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested locationEfses.
func (c *FakeLocationEfses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(locationefsesResource, c.ns, opts))

}

// Create takes the representation of a locationEfs and creates it.  Returns the server's representation of the locationEfs, and an error, if there is any.
func (c *FakeLocationEfses) Create(ctx context.Context, locationEfs *v1alpha1.LocationEfs, opts v1.CreateOptions) (result *v1alpha1.LocationEfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(locationefsesResource, c.ns, locationEfs), &v1alpha1.LocationEfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationEfs), err
}

// Update takes the representation of a locationEfs and updates it. Returns the server's representation of the locationEfs, and an error, if there is any.
func (c *FakeLocationEfses) Update(ctx context.Context, locationEfs *v1alpha1.LocationEfs, opts v1.UpdateOptions) (result *v1alpha1.LocationEfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(locationefsesResource, c.ns, locationEfs), &v1alpha1.LocationEfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationEfs), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocationEfses) UpdateStatus(ctx context.Context, locationEfs *v1alpha1.LocationEfs, opts v1.UpdateOptions) (*v1alpha1.LocationEfs, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(locationefsesResource, "status", c.ns, locationEfs), &v1alpha1.LocationEfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationEfs), err
}

// Delete takes name of the locationEfs and deletes it. Returns an error if one occurs.
func (c *FakeLocationEfses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(locationefsesResource, c.ns, name), &v1alpha1.LocationEfs{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocationEfses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(locationefsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LocationEfsList{})
	return err
}

// Patch applies the patch and returns the patched locationEfs.
func (c *FakeLocationEfses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocationEfs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(locationefsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.LocationEfs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationEfs), err
}
