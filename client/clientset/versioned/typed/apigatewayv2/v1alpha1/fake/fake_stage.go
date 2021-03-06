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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/apigatewayv2/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStages implements StageInterface
type FakeStages struct {
	Fake *FakeApigatewayv2V1alpha1
	ns   string
}

var stagesResource = schema.GroupVersionResource{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Resource: "stages"}

var stagesKind = schema.GroupVersionKind{Group: "apigatewayv2.aws.kubeform.com", Version: "v1alpha1", Kind: "Stage"}

// Get takes name of the stage, and returns the corresponding stage object, and an error if there is any.
func (c *FakeStages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Stage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(stagesResource, c.ns, name), &v1alpha1.Stage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stage), err
}

// List takes label and field selectors, and returns the list of Stages that match those selectors.
func (c *FakeStages) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.StageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(stagesResource, stagesKind, c.ns, opts), &v1alpha1.StageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StageList{ListMeta: obj.(*v1alpha1.StageList).ListMeta}
	for _, item := range obj.(*v1alpha1.StageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stages.
func (c *FakeStages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(stagesResource, c.ns, opts))

}

// Create takes the representation of a stage and creates it.  Returns the server's representation of the stage, and an error, if there is any.
func (c *FakeStages) Create(ctx context.Context, stage *v1alpha1.Stage, opts v1.CreateOptions) (result *v1alpha1.Stage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(stagesResource, c.ns, stage), &v1alpha1.Stage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stage), err
}

// Update takes the representation of a stage and updates it. Returns the server's representation of the stage, and an error, if there is any.
func (c *FakeStages) Update(ctx context.Context, stage *v1alpha1.Stage, opts v1.UpdateOptions) (result *v1alpha1.Stage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(stagesResource, c.ns, stage), &v1alpha1.Stage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStages) UpdateStatus(ctx context.Context, stage *v1alpha1.Stage, opts v1.UpdateOptions) (*v1alpha1.Stage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(stagesResource, "status", c.ns, stage), &v1alpha1.Stage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stage), err
}

// Delete takes name of the stage and deletes it. Returns an error if one occurs.
func (c *FakeStages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(stagesResource, c.ns, name), &v1alpha1.Stage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(stagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.StageList{})
	return err
}

// Patch applies the patch and returns the patched stage.
func (c *FakeStages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Stage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(stagesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Stage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Stage), err
}
