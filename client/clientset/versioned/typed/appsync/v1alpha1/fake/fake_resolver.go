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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/appsync/v1alpha1"
)

// FakeResolvers implements ResolverInterface
type FakeResolvers struct {
	Fake *FakeAppsyncV1alpha1
	ns   string
}

var resolversResource = schema.GroupVersionResource{Group: "appsync.aws.kubeform.com", Version: "v1alpha1", Resource: "resolvers"}

var resolversKind = schema.GroupVersionKind{Group: "appsync.aws.kubeform.com", Version: "v1alpha1", Kind: "Resolver"}

// Get takes name of the resolver, and returns the corresponding resolver object, and an error if there is any.
func (c *FakeResolvers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Resolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resolversResource, c.ns, name), &v1alpha1.Resolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resolver), err
}

// List takes label and field selectors, and returns the list of Resolvers that match those selectors.
func (c *FakeResolvers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ResolverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resolversResource, resolversKind, c.ns, opts), &v1alpha1.ResolverList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResolverList{ListMeta: obj.(*v1alpha1.ResolverList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResolverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resolvers.
func (c *FakeResolvers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resolversResource, c.ns, opts))

}

// Create takes the representation of a resolver and creates it.  Returns the server's representation of the resolver, and an error, if there is any.
func (c *FakeResolvers) Create(ctx context.Context, resolver *v1alpha1.Resolver, opts v1.CreateOptions) (result *v1alpha1.Resolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resolversResource, c.ns, resolver), &v1alpha1.Resolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resolver), err
}

// Update takes the representation of a resolver and updates it. Returns the server's representation of the resolver, and an error, if there is any.
func (c *FakeResolvers) Update(ctx context.Context, resolver *v1alpha1.Resolver, opts v1.UpdateOptions) (result *v1alpha1.Resolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resolversResource, c.ns, resolver), &v1alpha1.Resolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resolver), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResolvers) UpdateStatus(ctx context.Context, resolver *v1alpha1.Resolver, opts v1.UpdateOptions) (*v1alpha1.Resolver, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resolversResource, "status", c.ns, resolver), &v1alpha1.Resolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resolver), err
}

// Delete takes name of the resolver and deletes it. Returns an error if one occurs.
func (c *FakeResolvers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(resolversResource, c.ns, name), &v1alpha1.Resolver{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResolvers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resolversResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResolverList{})
	return err
}

// Patch applies the patch and returns the patched resolver.
func (c *FakeResolvers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Resolver, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resolversResource, c.ns, name, pt, data, subresources...), &v1alpha1.Resolver{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Resolver), err
}
