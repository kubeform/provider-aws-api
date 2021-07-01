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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudwatch/v1alpha1"
)

// FakeQueryDefinitions implements QueryDefinitionInterface
type FakeQueryDefinitions struct {
	Fake *FakeCloudwatchV1alpha1
	ns   string
}

var querydefinitionsResource = schema.GroupVersionResource{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Resource: "querydefinitions"}

var querydefinitionsKind = schema.GroupVersionKind{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Kind: "QueryDefinition"}

// Get takes name of the queryDefinition, and returns the corresponding queryDefinition object, and an error if there is any.
func (c *FakeQueryDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.QueryDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(querydefinitionsResource, c.ns, name), &v1alpha1.QueryDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QueryDefinition), err
}

// List takes label and field selectors, and returns the list of QueryDefinitions that match those selectors.
func (c *FakeQueryDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.QueryDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(querydefinitionsResource, querydefinitionsKind, c.ns, opts), &v1alpha1.QueryDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.QueryDefinitionList{ListMeta: obj.(*v1alpha1.QueryDefinitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.QueryDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested queryDefinitions.
func (c *FakeQueryDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(querydefinitionsResource, c.ns, opts))

}

// Create takes the representation of a queryDefinition and creates it.  Returns the server's representation of the queryDefinition, and an error, if there is any.
func (c *FakeQueryDefinitions) Create(ctx context.Context, queryDefinition *v1alpha1.QueryDefinition, opts v1.CreateOptions) (result *v1alpha1.QueryDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(querydefinitionsResource, c.ns, queryDefinition), &v1alpha1.QueryDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QueryDefinition), err
}

// Update takes the representation of a queryDefinition and updates it. Returns the server's representation of the queryDefinition, and an error, if there is any.
func (c *FakeQueryDefinitions) Update(ctx context.Context, queryDefinition *v1alpha1.QueryDefinition, opts v1.UpdateOptions) (result *v1alpha1.QueryDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(querydefinitionsResource, c.ns, queryDefinition), &v1alpha1.QueryDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QueryDefinition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeQueryDefinitions) UpdateStatus(ctx context.Context, queryDefinition *v1alpha1.QueryDefinition, opts v1.UpdateOptions) (*v1alpha1.QueryDefinition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(querydefinitionsResource, "status", c.ns, queryDefinition), &v1alpha1.QueryDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QueryDefinition), err
}

// Delete takes name of the queryDefinition and deletes it. Returns an error if one occurs.
func (c *FakeQueryDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(querydefinitionsResource, c.ns, name), &v1alpha1.QueryDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQueryDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(querydefinitionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.QueryDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched queryDefinition.
func (c *FakeQueryDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.QueryDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(querydefinitionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.QueryDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.QueryDefinition), err
}
