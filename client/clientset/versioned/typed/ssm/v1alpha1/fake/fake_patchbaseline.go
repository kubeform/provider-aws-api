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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ssm/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePatchBaselines implements PatchBaselineInterface
type FakePatchBaselines struct {
	Fake *FakeSsmV1alpha1
	ns   string
}

var patchbaselinesResource = schema.GroupVersionResource{Group: "ssm.aws.kubeform.com", Version: "v1alpha1", Resource: "patchbaselines"}

var patchbaselinesKind = schema.GroupVersionKind{Group: "ssm.aws.kubeform.com", Version: "v1alpha1", Kind: "PatchBaseline"}

// Get takes name of the patchBaseline, and returns the corresponding patchBaseline object, and an error if there is any.
func (c *FakePatchBaselines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PatchBaseline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(patchbaselinesResource, c.ns, name), &v1alpha1.PatchBaseline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PatchBaseline), err
}

// List takes label and field selectors, and returns the list of PatchBaselines that match those selectors.
func (c *FakePatchBaselines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PatchBaselineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(patchbaselinesResource, patchbaselinesKind, c.ns, opts), &v1alpha1.PatchBaselineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PatchBaselineList{ListMeta: obj.(*v1alpha1.PatchBaselineList).ListMeta}
	for _, item := range obj.(*v1alpha1.PatchBaselineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested patchBaselines.
func (c *FakePatchBaselines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(patchbaselinesResource, c.ns, opts))

}

// Create takes the representation of a patchBaseline and creates it.  Returns the server's representation of the patchBaseline, and an error, if there is any.
func (c *FakePatchBaselines) Create(ctx context.Context, patchBaseline *v1alpha1.PatchBaseline, opts v1.CreateOptions) (result *v1alpha1.PatchBaseline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(patchbaselinesResource, c.ns, patchBaseline), &v1alpha1.PatchBaseline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PatchBaseline), err
}

// Update takes the representation of a patchBaseline and updates it. Returns the server's representation of the patchBaseline, and an error, if there is any.
func (c *FakePatchBaselines) Update(ctx context.Context, patchBaseline *v1alpha1.PatchBaseline, opts v1.UpdateOptions) (result *v1alpha1.PatchBaseline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(patchbaselinesResource, c.ns, patchBaseline), &v1alpha1.PatchBaseline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PatchBaseline), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePatchBaselines) UpdateStatus(ctx context.Context, patchBaseline *v1alpha1.PatchBaseline, opts v1.UpdateOptions) (*v1alpha1.PatchBaseline, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(patchbaselinesResource, "status", c.ns, patchBaseline), &v1alpha1.PatchBaseline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PatchBaseline), err
}

// Delete takes name of the patchBaseline and deletes it. Returns an error if one occurs.
func (c *FakePatchBaselines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(patchbaselinesResource, c.ns, name), &v1alpha1.PatchBaseline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePatchBaselines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(patchbaselinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PatchBaselineList{})
	return err
}

// Patch applies the patch and returns the patched patchBaseline.
func (c *FakePatchBaselines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PatchBaseline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(patchbaselinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PatchBaseline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PatchBaseline), err
}
