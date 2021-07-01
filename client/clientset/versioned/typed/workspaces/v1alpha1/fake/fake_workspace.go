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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/workspaces/v1alpha1"
)

// FakeWorkspaces implements WorkspaceInterface
type FakeWorkspaces struct {
	Fake *FakeWorkspacesV1alpha1
	ns   string
}

var workspacesResource = schema.GroupVersionResource{Group: "workspaces.aws.kubeform.com", Version: "v1alpha1", Resource: "workspaces"}

var workspacesKind = schema.GroupVersionKind{Group: "workspaces.aws.kubeform.com", Version: "v1alpha1", Kind: "Workspace"}

// Get takes name of the workspace, and returns the corresponding workspace object, and an error if there is any.
func (c *FakeWorkspaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Workspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(workspacesResource, c.ns, name), &v1alpha1.Workspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workspace), err
}

// List takes label and field selectors, and returns the list of Workspaces that match those selectors.
func (c *FakeWorkspaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.WorkspaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(workspacesResource, workspacesKind, c.ns, opts), &v1alpha1.WorkspaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WorkspaceList{ListMeta: obj.(*v1alpha1.WorkspaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.WorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested workspaces.
func (c *FakeWorkspaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(workspacesResource, c.ns, opts))

}

// Create takes the representation of a workspace and creates it.  Returns the server's representation of the workspace, and an error, if there is any.
func (c *FakeWorkspaces) Create(ctx context.Context, workspace *v1alpha1.Workspace, opts v1.CreateOptions) (result *v1alpha1.Workspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(workspacesResource, c.ns, workspace), &v1alpha1.Workspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workspace), err
}

// Update takes the representation of a workspace and updates it. Returns the server's representation of the workspace, and an error, if there is any.
func (c *FakeWorkspaces) Update(ctx context.Context, workspace *v1alpha1.Workspace, opts v1.UpdateOptions) (result *v1alpha1.Workspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(workspacesResource, c.ns, workspace), &v1alpha1.Workspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workspace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWorkspaces) UpdateStatus(ctx context.Context, workspace *v1alpha1.Workspace, opts v1.UpdateOptions) (*v1alpha1.Workspace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(workspacesResource, "status", c.ns, workspace), &v1alpha1.Workspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workspace), err
}

// Delete takes name of the workspace and deletes it. Returns an error if one occurs.
func (c *FakeWorkspaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(workspacesResource, c.ns, name), &v1alpha1.Workspace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWorkspaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(workspacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.WorkspaceList{})
	return err
}

// Patch applies the patch and returns the patched workspace.
func (c *FakeWorkspaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Workspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(workspacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Workspace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Workspace), err
}
