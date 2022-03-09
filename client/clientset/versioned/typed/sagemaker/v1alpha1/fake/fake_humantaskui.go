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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/sagemaker/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHumanTaskUis implements HumanTaskUiInterface
type FakeHumanTaskUis struct {
	Fake *FakeSagemakerV1alpha1
	ns   string
}

var humantaskuisResource = schema.GroupVersionResource{Group: "sagemaker.aws.kubeform.com", Version: "v1alpha1", Resource: "humantaskuis"}

var humantaskuisKind = schema.GroupVersionKind{Group: "sagemaker.aws.kubeform.com", Version: "v1alpha1", Kind: "HumanTaskUi"}

// Get takes name of the humanTaskUi, and returns the corresponding humanTaskUi object, and an error if there is any.
func (c *FakeHumanTaskUis) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HumanTaskUi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(humantaskuisResource, c.ns, name), &v1alpha1.HumanTaskUi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HumanTaskUi), err
}

// List takes label and field selectors, and returns the list of HumanTaskUis that match those selectors.
func (c *FakeHumanTaskUis) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HumanTaskUiList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(humantaskuisResource, humantaskuisKind, c.ns, opts), &v1alpha1.HumanTaskUiList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HumanTaskUiList{ListMeta: obj.(*v1alpha1.HumanTaskUiList).ListMeta}
	for _, item := range obj.(*v1alpha1.HumanTaskUiList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested humanTaskUis.
func (c *FakeHumanTaskUis) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(humantaskuisResource, c.ns, opts))

}

// Create takes the representation of a humanTaskUi and creates it.  Returns the server's representation of the humanTaskUi, and an error, if there is any.
func (c *FakeHumanTaskUis) Create(ctx context.Context, humanTaskUi *v1alpha1.HumanTaskUi, opts v1.CreateOptions) (result *v1alpha1.HumanTaskUi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(humantaskuisResource, c.ns, humanTaskUi), &v1alpha1.HumanTaskUi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HumanTaskUi), err
}

// Update takes the representation of a humanTaskUi and updates it. Returns the server's representation of the humanTaskUi, and an error, if there is any.
func (c *FakeHumanTaskUis) Update(ctx context.Context, humanTaskUi *v1alpha1.HumanTaskUi, opts v1.UpdateOptions) (result *v1alpha1.HumanTaskUi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(humantaskuisResource, c.ns, humanTaskUi), &v1alpha1.HumanTaskUi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HumanTaskUi), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHumanTaskUis) UpdateStatus(ctx context.Context, humanTaskUi *v1alpha1.HumanTaskUi, opts v1.UpdateOptions) (*v1alpha1.HumanTaskUi, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(humantaskuisResource, "status", c.ns, humanTaskUi), &v1alpha1.HumanTaskUi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HumanTaskUi), err
}

// Delete takes name of the humanTaskUi and deletes it. Returns an error if one occurs.
func (c *FakeHumanTaskUis) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(humantaskuisResource, c.ns, name), &v1alpha1.HumanTaskUi{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHumanTaskUis) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(humantaskuisResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HumanTaskUiList{})
	return err
}

// Patch applies the patch and returns the patched humanTaskUi.
func (c *FakeHumanTaskUis) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HumanTaskUi, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(humantaskuisResource, c.ns, name, pt, data, subresources...), &v1alpha1.HumanTaskUi{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HumanTaskUi), err
}
