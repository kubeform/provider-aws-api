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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/launch/v1alpha1"
)

// FakeTemplates implements TemplateInterface
type FakeTemplates struct {
	Fake *FakeLaunchV1alpha1
	ns   string
}

var templatesResource = schema.GroupVersionResource{Group: "launch.aws.kubeform.com", Version: "v1alpha1", Resource: "templates"}

var templatesKind = schema.GroupVersionKind{Group: "launch.aws.kubeform.com", Version: "v1alpha1", Kind: "Template"}

// Get takes name of the template, and returns the corresponding template object, and an error if there is any.
func (c *FakeTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(templatesResource, c.ns, name), &v1alpha1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}

// List takes label and field selectors, and returns the list of Templates that match those selectors.
func (c *FakeTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(templatesResource, templatesKind, c.ns, opts), &v1alpha1.TemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TemplateList{ListMeta: obj.(*v1alpha1.TemplateList).ListMeta}
	for _, item := range obj.(*v1alpha1.TemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested templates.
func (c *FakeTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(templatesResource, c.ns, opts))

}

// Create takes the representation of a template and creates it.  Returns the server's representation of the template, and an error, if there is any.
func (c *FakeTemplates) Create(ctx context.Context, template *v1alpha1.Template, opts v1.CreateOptions) (result *v1alpha1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(templatesResource, c.ns, template), &v1alpha1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}

// Update takes the representation of a template and updates it. Returns the server's representation of the template, and an error, if there is any.
func (c *FakeTemplates) Update(ctx context.Context, template *v1alpha1.Template, opts v1.UpdateOptions) (result *v1alpha1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(templatesResource, c.ns, template), &v1alpha1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTemplates) UpdateStatus(ctx context.Context, template *v1alpha1.Template, opts v1.UpdateOptions) (*v1alpha1.Template, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(templatesResource, "status", c.ns, template), &v1alpha1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}

// Delete takes name of the template and deletes it. Returns an error if one occurs.
func (c *FakeTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(templatesResource, c.ns, name), &v1alpha1.Template{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(templatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TemplateList{})
	return err
}

// Patch applies the patch and returns the patched template.
func (c *FakeTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Template, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(templatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Template{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Template), err
}
