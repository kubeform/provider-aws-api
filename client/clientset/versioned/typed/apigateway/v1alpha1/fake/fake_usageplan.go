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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/apigateway/v1alpha1"
)

// FakeUsagePlans implements UsagePlanInterface
type FakeUsagePlans struct {
	Fake *FakeApigatewayV1alpha1
	ns   string
}

var usageplansResource = schema.GroupVersionResource{Group: "apigateway.aws.kubeform.com", Version: "v1alpha1", Resource: "usageplans"}

var usageplansKind = schema.GroupVersionKind{Group: "apigateway.aws.kubeform.com", Version: "v1alpha1", Kind: "UsagePlan"}

// Get takes name of the usagePlan, and returns the corresponding usagePlan object, and an error if there is any.
func (c *FakeUsagePlans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UsagePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(usageplansResource, c.ns, name), &v1alpha1.UsagePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsagePlan), err
}

// List takes label and field selectors, and returns the list of UsagePlans that match those selectors.
func (c *FakeUsagePlans) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UsagePlanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(usageplansResource, usageplansKind, c.ns, opts), &v1alpha1.UsagePlanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UsagePlanList{ListMeta: obj.(*v1alpha1.UsagePlanList).ListMeta}
	for _, item := range obj.(*v1alpha1.UsagePlanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested usagePlans.
func (c *FakeUsagePlans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(usageplansResource, c.ns, opts))

}

// Create takes the representation of a usagePlan and creates it.  Returns the server's representation of the usagePlan, and an error, if there is any.
func (c *FakeUsagePlans) Create(ctx context.Context, usagePlan *v1alpha1.UsagePlan, opts v1.CreateOptions) (result *v1alpha1.UsagePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(usageplansResource, c.ns, usagePlan), &v1alpha1.UsagePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsagePlan), err
}

// Update takes the representation of a usagePlan and updates it. Returns the server's representation of the usagePlan, and an error, if there is any.
func (c *FakeUsagePlans) Update(ctx context.Context, usagePlan *v1alpha1.UsagePlan, opts v1.UpdateOptions) (result *v1alpha1.UsagePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(usageplansResource, c.ns, usagePlan), &v1alpha1.UsagePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsagePlan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUsagePlans) UpdateStatus(ctx context.Context, usagePlan *v1alpha1.UsagePlan, opts v1.UpdateOptions) (*v1alpha1.UsagePlan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(usageplansResource, "status", c.ns, usagePlan), &v1alpha1.UsagePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsagePlan), err
}

// Delete takes name of the usagePlan and deletes it. Returns an error if one occurs.
func (c *FakeUsagePlans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(usageplansResource, c.ns, name), &v1alpha1.UsagePlan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUsagePlans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(usageplansResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UsagePlanList{})
	return err
}

// Patch applies the patch and returns the patched usagePlan.
func (c *FakeUsagePlans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UsagePlan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(usageplansResource, c.ns, name, pt, data, subresources...), &v1alpha1.UsagePlan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UsagePlan), err
}
