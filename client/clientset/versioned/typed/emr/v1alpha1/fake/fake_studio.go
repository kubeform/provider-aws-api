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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/emr/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStudios implements StudioInterface
type FakeStudios struct {
	Fake *FakeEmrV1alpha1
	ns   string
}

var studiosResource = schema.GroupVersionResource{Group: "emr.aws.kubeform.com", Version: "v1alpha1", Resource: "studios"}

var studiosKind = schema.GroupVersionKind{Group: "emr.aws.kubeform.com", Version: "v1alpha1", Kind: "Studio"}

// Get takes name of the studio, and returns the corresponding studio object, and an error if there is any.
func (c *FakeStudios) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Studio, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(studiosResource, c.ns, name), &v1alpha1.Studio{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Studio), err
}

// List takes label and field selectors, and returns the list of Studios that match those selectors.
func (c *FakeStudios) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StudioList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(studiosResource, studiosKind, c.ns, opts), &v1alpha1.StudioList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StudioList{ListMeta: obj.(*v1alpha1.StudioList).ListMeta}
	for _, item := range obj.(*v1alpha1.StudioList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested studios.
func (c *FakeStudios) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(studiosResource, c.ns, opts))

}

// Create takes the representation of a studio and creates it.  Returns the server's representation of the studio, and an error, if there is any.
func (c *FakeStudios) Create(ctx context.Context, studio *v1alpha1.Studio, opts v1.CreateOptions) (result *v1alpha1.Studio, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(studiosResource, c.ns, studio), &v1alpha1.Studio{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Studio), err
}

// Update takes the representation of a studio and updates it. Returns the server's representation of the studio, and an error, if there is any.
func (c *FakeStudios) Update(ctx context.Context, studio *v1alpha1.Studio, opts v1.UpdateOptions) (result *v1alpha1.Studio, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(studiosResource, c.ns, studio), &v1alpha1.Studio{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Studio), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStudios) UpdateStatus(ctx context.Context, studio *v1alpha1.Studio, opts v1.UpdateOptions) (*v1alpha1.Studio, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(studiosResource, "status", c.ns, studio), &v1alpha1.Studio{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Studio), err
}

// Delete takes name of the studio and deletes it. Returns an error if one occurs.
func (c *FakeStudios) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(studiosResource, c.ns, name), &v1alpha1.Studio{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStudios) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(studiosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StudioList{})
	return err
}

// Patch applies the patch and returns the patched studio.
func (c *FakeStudios) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Studio, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(studiosResource, c.ns, name, pt, data, subresources...), &v1alpha1.Studio{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Studio), err
}
