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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/datasync/v1alpha1"
)

// FakeLocationSmbs implements LocationSmbInterface
type FakeLocationSmbs struct {
	Fake *FakeDatasyncV1alpha1
	ns   string
}

var locationsmbsResource = schema.GroupVersionResource{Group: "datasync.aws.kubeform.com", Version: "v1alpha1", Resource: "locationsmbs"}

var locationsmbsKind = schema.GroupVersionKind{Group: "datasync.aws.kubeform.com", Version: "v1alpha1", Kind: "LocationSmb"}

// Get takes name of the locationSmb, and returns the corresponding locationSmb object, and an error if there is any.
func (c *FakeLocationSmbs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocationSmb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(locationsmbsResource, c.ns, name), &v1alpha1.LocationSmb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationSmb), err
}

// List takes label and field selectors, and returns the list of LocationSmbs that match those selectors.
func (c *FakeLocationSmbs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocationSmbList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(locationsmbsResource, locationsmbsKind, c.ns, opts), &v1alpha1.LocationSmbList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LocationSmbList{ListMeta: obj.(*v1alpha1.LocationSmbList).ListMeta}
	for _, item := range obj.(*v1alpha1.LocationSmbList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested locationSmbs.
func (c *FakeLocationSmbs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(locationsmbsResource, c.ns, opts))

}

// Create takes the representation of a locationSmb and creates it.  Returns the server's representation of the locationSmb, and an error, if there is any.
func (c *FakeLocationSmbs) Create(ctx context.Context, locationSmb *v1alpha1.LocationSmb, opts v1.CreateOptions) (result *v1alpha1.LocationSmb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(locationsmbsResource, c.ns, locationSmb), &v1alpha1.LocationSmb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationSmb), err
}

// Update takes the representation of a locationSmb and updates it. Returns the server's representation of the locationSmb, and an error, if there is any.
func (c *FakeLocationSmbs) Update(ctx context.Context, locationSmb *v1alpha1.LocationSmb, opts v1.UpdateOptions) (result *v1alpha1.LocationSmb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(locationsmbsResource, c.ns, locationSmb), &v1alpha1.LocationSmb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationSmb), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLocationSmbs) UpdateStatus(ctx context.Context, locationSmb *v1alpha1.LocationSmb, opts v1.UpdateOptions) (*v1alpha1.LocationSmb, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(locationsmbsResource, "status", c.ns, locationSmb), &v1alpha1.LocationSmb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationSmb), err
}

// Delete takes name of the locationSmb and deletes it. Returns an error if one occurs.
func (c *FakeLocationSmbs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(locationsmbsResource, c.ns, name), &v1alpha1.LocationSmb{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLocationSmbs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(locationsmbsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LocationSmbList{})
	return err
}

// Patch applies the patch and returns the patched locationSmb.
func (c *FakeLocationSmbs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocationSmb, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(locationsmbsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LocationSmb{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LocationSmb), err
}
