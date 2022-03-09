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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/fsx/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOntapStorageVirtualMachines implements OntapStorageVirtualMachineInterface
type FakeOntapStorageVirtualMachines struct {
	Fake *FakeFsxV1alpha1
	ns   string
}

var ontapstoragevirtualmachinesResource = schema.GroupVersionResource{Group: "fsx.aws.kubeform.com", Version: "v1alpha1", Resource: "ontapstoragevirtualmachines"}

var ontapstoragevirtualmachinesKind = schema.GroupVersionKind{Group: "fsx.aws.kubeform.com", Version: "v1alpha1", Kind: "OntapStorageVirtualMachine"}

// Get takes name of the ontapStorageVirtualMachine, and returns the corresponding ontapStorageVirtualMachine object, and an error if there is any.
func (c *FakeOntapStorageVirtualMachines) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OntapStorageVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ontapstoragevirtualmachinesResource, c.ns, name), &v1alpha1.OntapStorageVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OntapStorageVirtualMachine), err
}

// List takes label and field selectors, and returns the list of OntapStorageVirtualMachines that match those selectors.
func (c *FakeOntapStorageVirtualMachines) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OntapStorageVirtualMachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ontapstoragevirtualmachinesResource, ontapstoragevirtualmachinesKind, c.ns, opts), &v1alpha1.OntapStorageVirtualMachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OntapStorageVirtualMachineList{ListMeta: obj.(*v1alpha1.OntapStorageVirtualMachineList).ListMeta}
	for _, item := range obj.(*v1alpha1.OntapStorageVirtualMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ontapStorageVirtualMachines.
func (c *FakeOntapStorageVirtualMachines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ontapstoragevirtualmachinesResource, c.ns, opts))

}

// Create takes the representation of a ontapStorageVirtualMachine and creates it.  Returns the server's representation of the ontapStorageVirtualMachine, and an error, if there is any.
func (c *FakeOntapStorageVirtualMachines) Create(ctx context.Context, ontapStorageVirtualMachine *v1alpha1.OntapStorageVirtualMachine, opts v1.CreateOptions) (result *v1alpha1.OntapStorageVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ontapstoragevirtualmachinesResource, c.ns, ontapStorageVirtualMachine), &v1alpha1.OntapStorageVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OntapStorageVirtualMachine), err
}

// Update takes the representation of a ontapStorageVirtualMachine and updates it. Returns the server's representation of the ontapStorageVirtualMachine, and an error, if there is any.
func (c *FakeOntapStorageVirtualMachines) Update(ctx context.Context, ontapStorageVirtualMachine *v1alpha1.OntapStorageVirtualMachine, opts v1.UpdateOptions) (result *v1alpha1.OntapStorageVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ontapstoragevirtualmachinesResource, c.ns, ontapStorageVirtualMachine), &v1alpha1.OntapStorageVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OntapStorageVirtualMachine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOntapStorageVirtualMachines) UpdateStatus(ctx context.Context, ontapStorageVirtualMachine *v1alpha1.OntapStorageVirtualMachine, opts v1.UpdateOptions) (*v1alpha1.OntapStorageVirtualMachine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ontapstoragevirtualmachinesResource, "status", c.ns, ontapStorageVirtualMachine), &v1alpha1.OntapStorageVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OntapStorageVirtualMachine), err
}

// Delete takes name of the ontapStorageVirtualMachine and deletes it. Returns an error if one occurs.
func (c *FakeOntapStorageVirtualMachines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ontapstoragevirtualmachinesResource, c.ns, name), &v1alpha1.OntapStorageVirtualMachine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOntapStorageVirtualMachines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ontapstoragevirtualmachinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OntapStorageVirtualMachineList{})
	return err
}

// Patch applies the patch and returns the patched ontapStorageVirtualMachine.
func (c *FakeOntapStorageVirtualMachines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OntapStorageVirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ontapstoragevirtualmachinesResource, c.ns, name, pt, data, subresources...), &v1alpha1.OntapStorageVirtualMachine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OntapStorageVirtualMachine), err
}
