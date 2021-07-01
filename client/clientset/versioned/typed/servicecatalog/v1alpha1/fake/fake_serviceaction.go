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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/servicecatalog/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceActions implements ServiceActionInterface
type FakeServiceActions struct {
	Fake *FakeServicecatalogV1alpha1
	ns   string
}

var serviceactionsResource = schema.GroupVersionResource{Group: "servicecatalog.aws.kubeform.com", Version: "v1alpha1", Resource: "serviceactions"}

var serviceactionsKind = schema.GroupVersionKind{Group: "servicecatalog.aws.kubeform.com", Version: "v1alpha1", Kind: "ServiceAction"}

// Get takes name of the serviceAction, and returns the corresponding serviceAction object, and an error if there is any.
func (c *FakeServiceActions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServiceAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(serviceactionsResource, c.ns, name), &v1alpha1.ServiceAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAction), err
}

// List takes label and field selectors, and returns the list of ServiceActions that match those selectors.
func (c *FakeServiceActions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServiceActionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(serviceactionsResource, serviceactionsKind, c.ns, opts), &v1alpha1.ServiceActionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServiceActionList{ListMeta: obj.(*v1alpha1.ServiceActionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServiceActionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceActions.
func (c *FakeServiceActions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(serviceactionsResource, c.ns, opts))

}

// Create takes the representation of a serviceAction and creates it.  Returns the server's representation of the serviceAction, and an error, if there is any.
func (c *FakeServiceActions) Create(ctx context.Context, serviceAction *v1alpha1.ServiceAction, opts v1.CreateOptions) (result *v1alpha1.ServiceAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(serviceactionsResource, c.ns, serviceAction), &v1alpha1.ServiceAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAction), err
}

// Update takes the representation of a serviceAction and updates it. Returns the server's representation of the serviceAction, and an error, if there is any.
func (c *FakeServiceActions) Update(ctx context.Context, serviceAction *v1alpha1.ServiceAction, opts v1.UpdateOptions) (result *v1alpha1.ServiceAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(serviceactionsResource, c.ns, serviceAction), &v1alpha1.ServiceAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAction), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceActions) UpdateStatus(ctx context.Context, serviceAction *v1alpha1.ServiceAction, opts v1.UpdateOptions) (*v1alpha1.ServiceAction, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(serviceactionsResource, "status", c.ns, serviceAction), &v1alpha1.ServiceAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAction), err
}

// Delete takes name of the serviceAction and deletes it. Returns an error if one occurs.
func (c *FakeServiceActions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(serviceactionsResource, c.ns, name), &v1alpha1.ServiceAction{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceActions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(serviceactionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServiceActionList{})
	return err
}

// Patch applies the patch and returns the patched serviceAction.
func (c *FakeServiceActions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServiceAction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(serviceactionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServiceAction{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServiceAction), err
}
