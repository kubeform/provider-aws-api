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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/batch/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeEnvironments implements ComputeEnvironmentInterface
type FakeComputeEnvironments struct {
	Fake *FakeBatchV1alpha1
	ns   string
}

var computeenvironmentsResource = schema.GroupVersionResource{Group: "batch.aws.kubeform.com", Version: "v1alpha1", Resource: "computeenvironments"}

var computeenvironmentsKind = schema.GroupVersionKind{Group: "batch.aws.kubeform.com", Version: "v1alpha1", Kind: "ComputeEnvironment"}

// Get takes name of the computeEnvironment, and returns the corresponding computeEnvironment object, and an error if there is any.
func (c *FakeComputeEnvironments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ComputeEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeenvironmentsResource, c.ns, name), &v1alpha1.ComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeEnvironment), err
}

// List takes label and field selectors, and returns the list of ComputeEnvironments that match those selectors.
func (c *FakeComputeEnvironments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ComputeEnvironmentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeenvironmentsResource, computeenvironmentsKind, c.ns, opts), &v1alpha1.ComputeEnvironmentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ComputeEnvironmentList{ListMeta: obj.(*v1alpha1.ComputeEnvironmentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ComputeEnvironmentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeEnvironments.
func (c *FakeComputeEnvironments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeenvironmentsResource, c.ns, opts))

}

// Create takes the representation of a computeEnvironment and creates it.  Returns the server's representation of the computeEnvironment, and an error, if there is any.
func (c *FakeComputeEnvironments) Create(ctx context.Context, computeEnvironment *v1alpha1.ComputeEnvironment, opts v1.CreateOptions) (result *v1alpha1.ComputeEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeenvironmentsResource, c.ns, computeEnvironment), &v1alpha1.ComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeEnvironment), err
}

// Update takes the representation of a computeEnvironment and updates it. Returns the server's representation of the computeEnvironment, and an error, if there is any.
func (c *FakeComputeEnvironments) Update(ctx context.Context, computeEnvironment *v1alpha1.ComputeEnvironment, opts v1.UpdateOptions) (result *v1alpha1.ComputeEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeenvironmentsResource, c.ns, computeEnvironment), &v1alpha1.ComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeEnvironment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeEnvironments) UpdateStatus(ctx context.Context, computeEnvironment *v1alpha1.ComputeEnvironment, opts v1.UpdateOptions) (*v1alpha1.ComputeEnvironment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeenvironmentsResource, "status", c.ns, computeEnvironment), &v1alpha1.ComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeEnvironment), err
}

// Delete takes name of the computeEnvironment and deletes it. Returns an error if one occurs.
func (c *FakeComputeEnvironments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(computeenvironmentsResource, c.ns, name), &v1alpha1.ComputeEnvironment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeEnvironments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeenvironmentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ComputeEnvironmentList{})
	return err
}

// Patch applies the patch and returns the patched computeEnvironment.
func (c *FakeComputeEnvironments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ComputeEnvironment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeenvironmentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ComputeEnvironment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ComputeEnvironment), err
}
