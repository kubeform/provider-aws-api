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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSerialConsoleAccesses implements SerialConsoleAccessInterface
type FakeSerialConsoleAccesses struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var serialconsoleaccessesResource = schema.GroupVersionResource{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Resource: "serialconsoleaccesses"}

var serialconsoleaccessesKind = schema.GroupVersionKind{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Kind: "SerialConsoleAccess"}

// Get takes name of the serialConsoleAccess, and returns the corresponding serialConsoleAccess object, and an error if there is any.
func (c *FakeSerialConsoleAccesses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SerialConsoleAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serialconsoleaccessesResource, c.ns, name), &v1alpha1.SerialConsoleAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SerialConsoleAccess), err
}

// List takes label and field selectors, and returns the list of SerialConsoleAccesses that match those selectors.
func (c *FakeSerialConsoleAccesses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SerialConsoleAccessList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serialconsoleaccessesResource, serialconsoleaccessesKind, c.ns, opts), &v1alpha1.SerialConsoleAccessList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SerialConsoleAccessList{ListMeta: obj.(*v1alpha1.SerialConsoleAccessList).ListMeta}
	for _, item := range obj.(*v1alpha1.SerialConsoleAccessList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serialConsoleAccesses.
func (c *FakeSerialConsoleAccesses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serialconsoleaccessesResource, c.ns, opts))

}

// Create takes the representation of a serialConsoleAccess and creates it.  Returns the server's representation of the serialConsoleAccess, and an error, if there is any.
func (c *FakeSerialConsoleAccesses) Create(ctx context.Context, serialConsoleAccess *v1alpha1.SerialConsoleAccess, opts v1.CreateOptions) (result *v1alpha1.SerialConsoleAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serialconsoleaccessesResource, c.ns, serialConsoleAccess), &v1alpha1.SerialConsoleAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SerialConsoleAccess), err
}

// Update takes the representation of a serialConsoleAccess and updates it. Returns the server's representation of the serialConsoleAccess, and an error, if there is any.
func (c *FakeSerialConsoleAccesses) Update(ctx context.Context, serialConsoleAccess *v1alpha1.SerialConsoleAccess, opts v1.UpdateOptions) (result *v1alpha1.SerialConsoleAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serialconsoleaccessesResource, c.ns, serialConsoleAccess), &v1alpha1.SerialConsoleAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SerialConsoleAccess), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSerialConsoleAccesses) UpdateStatus(ctx context.Context, serialConsoleAccess *v1alpha1.SerialConsoleAccess, opts v1.UpdateOptions) (*v1alpha1.SerialConsoleAccess, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serialconsoleaccessesResource, "status", c.ns, serialConsoleAccess), &v1alpha1.SerialConsoleAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SerialConsoleAccess), err
}

// Delete takes name of the serialConsoleAccess and deletes it. Returns an error if one occurs.
func (c *FakeSerialConsoleAccesses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serialconsoleaccessesResource, c.ns, name), &v1alpha1.SerialConsoleAccess{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSerialConsoleAccesses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serialconsoleaccessesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SerialConsoleAccessList{})
	return err
}

// Patch applies the patch and returns the patched serialConsoleAccess.
func (c *FakeSerialConsoleAccesses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SerialConsoleAccess, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serialconsoleaccessesResource, c.ns, name, pt, data, subresources...), &v1alpha1.SerialConsoleAccess{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SerialConsoleAccess), err
}
