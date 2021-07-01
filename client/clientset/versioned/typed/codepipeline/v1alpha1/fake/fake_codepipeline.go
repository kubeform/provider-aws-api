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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/codepipeline/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCodepipelines implements CodepipelineInterface
type FakeCodepipelines struct {
	Fake *FakeCodepipelineV1alpha1
	ns   string
}

var codepipelinesResource = schema.GroupVersionResource{Group: "codepipeline.aws.kubeform.com", Version: "v1alpha1", Resource: "codepipelines"}

var codepipelinesKind = schema.GroupVersionKind{Group: "codepipeline.aws.kubeform.com", Version: "v1alpha1", Kind: "Codepipeline"}

// Get takes name of the codepipeline, and returns the corresponding codepipeline object, and an error if there is any.
func (c *FakeCodepipelines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Codepipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(codepipelinesResource, c.ns, name), &v1alpha1.Codepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Codepipeline), err
}

// List takes label and field selectors, and returns the list of Codepipelines that match those selectors.
func (c *FakeCodepipelines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CodepipelineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(codepipelinesResource, codepipelinesKind, c.ns, opts), &v1alpha1.CodepipelineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CodepipelineList{ListMeta: obj.(*v1alpha1.CodepipelineList).ListMeta}
	for _, item := range obj.(*v1alpha1.CodepipelineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested codepipelines.
func (c *FakeCodepipelines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(codepipelinesResource, c.ns, opts))

}

// Create takes the representation of a codepipeline and creates it.  Returns the server's representation of the codepipeline, and an error, if there is any.
func (c *FakeCodepipelines) Create(ctx context.Context, codepipeline *v1alpha1.Codepipeline, opts v1.CreateOptions) (result *v1alpha1.Codepipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(codepipelinesResource, c.ns, codepipeline), &v1alpha1.Codepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Codepipeline), err
}

// Update takes the representation of a codepipeline and updates it. Returns the server's representation of the codepipeline, and an error, if there is any.
func (c *FakeCodepipelines) Update(ctx context.Context, codepipeline *v1alpha1.Codepipeline, opts v1.UpdateOptions) (result *v1alpha1.Codepipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(codepipelinesResource, c.ns, codepipeline), &v1alpha1.Codepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Codepipeline), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCodepipelines) UpdateStatus(ctx context.Context, codepipeline *v1alpha1.Codepipeline, opts v1.UpdateOptions) (*v1alpha1.Codepipeline, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(codepipelinesResource, "status", c.ns, codepipeline), &v1alpha1.Codepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Codepipeline), err
}

// Delete takes name of the codepipeline and deletes it. Returns an error if one occurs.
func (c *FakeCodepipelines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(codepipelinesResource, c.ns, name), &v1alpha1.Codepipeline{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCodepipelines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(codepipelinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CodepipelineList{})
	return err
}

// Patch applies the patch and returns the patched codepipeline.
func (c *FakeCodepipelines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Codepipeline, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(codepipelinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Codepipeline{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Codepipeline), err
}
