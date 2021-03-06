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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/config/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeConformancePacks implements ConformancePackInterface
type FakeConformancePacks struct {
	Fake *FakeConfigV1alpha1
	ns   string
}

var conformancepacksResource = schema.GroupVersionResource{Group: "config.aws.kubeform.com", Version: "v1alpha1", Resource: "conformancepacks"}

var conformancepacksKind = schema.GroupVersionKind{Group: "config.aws.kubeform.com", Version: "v1alpha1", Kind: "ConformancePack"}

// Get takes name of the conformancePack, and returns the corresponding conformancePack object, and an error if there is any.
func (c *FakeConformancePacks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ConformancePack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(conformancepacksResource, c.ns, name), &v1alpha1.ConformancePack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConformancePack), err
}

// List takes label and field selectors, and returns the list of ConformancePacks that match those selectors.
func (c *FakeConformancePacks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConformancePackList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(conformancepacksResource, conformancepacksKind, c.ns, opts), &v1alpha1.ConformancePackList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ConformancePackList{ListMeta: obj.(*v1alpha1.ConformancePackList).ListMeta}
	for _, item := range obj.(*v1alpha1.ConformancePackList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested conformancePacks.
func (c *FakeConformancePacks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(conformancepacksResource, c.ns, opts))

}

// Create takes the representation of a conformancePack and creates it.  Returns the server's representation of the conformancePack, and an error, if there is any.
func (c *FakeConformancePacks) Create(ctx context.Context, conformancePack *v1alpha1.ConformancePack, opts v1.CreateOptions) (result *v1alpha1.ConformancePack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(conformancepacksResource, c.ns, conformancePack), &v1alpha1.ConformancePack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConformancePack), err
}

// Update takes the representation of a conformancePack and updates it. Returns the server's representation of the conformancePack, and an error, if there is any.
func (c *FakeConformancePacks) Update(ctx context.Context, conformancePack *v1alpha1.ConformancePack, opts v1.UpdateOptions) (result *v1alpha1.ConformancePack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(conformancepacksResource, c.ns, conformancePack), &v1alpha1.ConformancePack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConformancePack), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConformancePacks) UpdateStatus(ctx context.Context, conformancePack *v1alpha1.ConformancePack, opts v1.UpdateOptions) (*v1alpha1.ConformancePack, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(conformancepacksResource, "status", c.ns, conformancePack), &v1alpha1.ConformancePack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConformancePack), err
}

// Delete takes name of the conformancePack and deletes it. Returns an error if one occurs.
func (c *FakeConformancePacks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(conformancepacksResource, c.ns, name), &v1alpha1.ConformancePack{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConformancePacks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(conformancepacksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ConformancePackList{})
	return err
}

// Patch applies the patch and returns the patched conformancePack.
func (c *FakeConformancePacks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ConformancePack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(conformancepacksResource, c.ns, name, pt, data, subresources...), &v1alpha1.ConformancePack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ConformancePack), err
}
