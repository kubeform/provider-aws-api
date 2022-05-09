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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/gamelift/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScripts implements ScriptInterface
type FakeScripts struct {
	Fake *FakeGameliftV1alpha1
	ns   string
}

var scriptsResource = schema.GroupVersionResource{Group: "gamelift.aws.kubeform.com", Version: "v1alpha1", Resource: "scripts"}

var scriptsKind = schema.GroupVersionKind{Group: "gamelift.aws.kubeform.com", Version: "v1alpha1", Kind: "Script"}

// Get takes name of the script, and returns the corresponding script object, and an error if there is any.
func (c *FakeScripts) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Script, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scriptsResource, c.ns, name), &v1alpha1.Script{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Script), err
}

// List takes label and field selectors, and returns the list of Scripts that match those selectors.
func (c *FakeScripts) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScriptList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scriptsResource, scriptsKind, c.ns, opts), &v1alpha1.ScriptList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScriptList{ListMeta: obj.(*v1alpha1.ScriptList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScriptList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scripts.
func (c *FakeScripts) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scriptsResource, c.ns, opts))

}

// Create takes the representation of a script and creates it.  Returns the server's representation of the script, and an error, if there is any.
func (c *FakeScripts) Create(ctx context.Context, script *v1alpha1.Script, opts v1.CreateOptions) (result *v1alpha1.Script, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scriptsResource, c.ns, script), &v1alpha1.Script{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Script), err
}

// Update takes the representation of a script and updates it. Returns the server's representation of the script, and an error, if there is any.
func (c *FakeScripts) Update(ctx context.Context, script *v1alpha1.Script, opts v1.UpdateOptions) (result *v1alpha1.Script, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scriptsResource, c.ns, script), &v1alpha1.Script{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Script), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScripts) UpdateStatus(ctx context.Context, script *v1alpha1.Script, opts v1.UpdateOptions) (*v1alpha1.Script, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(scriptsResource, "status", c.ns, script), &v1alpha1.Script{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Script), err
}

// Delete takes name of the script and deletes it. Returns an error if one occurs.
func (c *FakeScripts) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scriptsResource, c.ns, name), &v1alpha1.Script{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScripts) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scriptsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScriptList{})
	return err
}

// Patch applies the patch and returns the patched script.
func (c *FakeScripts) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Script, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scriptsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Script{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Script), err
}