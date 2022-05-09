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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterCapacityProviderses implements ClusterCapacityProvidersInterface
type FakeClusterCapacityProviderses struct {
	Fake *FakeEcsV1alpha1
	ns   string
}

var clustercapacityprovidersesResource = schema.GroupVersionResource{Group: "ecs.aws.kubeform.com", Version: "v1alpha1", Resource: "clustercapacityproviderses"}

var clustercapacityprovidersesKind = schema.GroupVersionKind{Group: "ecs.aws.kubeform.com", Version: "v1alpha1", Kind: "ClusterCapacityProviders"}

// Get takes name of the clusterCapacityProviders, and returns the corresponding clusterCapacityProviders object, and an error if there is any.
func (c *FakeClusterCapacityProviderses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterCapacityProviders, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustercapacityprovidersesResource, c.ns, name), &v1alpha1.ClusterCapacityProviders{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCapacityProviders), err
}

// List takes label and field selectors, and returns the list of ClusterCapacityProviderses that match those selectors.
func (c *FakeClusterCapacityProviderses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterCapacityProvidersList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustercapacityprovidersesResource, clustercapacityprovidersesKind, c.ns, opts), &v1alpha1.ClusterCapacityProvidersList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterCapacityProvidersList{ListMeta: obj.(*v1alpha1.ClusterCapacityProvidersList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterCapacityProvidersList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterCapacityProviderses.
func (c *FakeClusterCapacityProviderses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustercapacityprovidersesResource, c.ns, opts))

}

// Create takes the representation of a clusterCapacityProviders and creates it.  Returns the server's representation of the clusterCapacityProviders, and an error, if there is any.
func (c *FakeClusterCapacityProviderses) Create(ctx context.Context, clusterCapacityProviders *v1alpha1.ClusterCapacityProviders, opts v1.CreateOptions) (result *v1alpha1.ClusterCapacityProviders, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustercapacityprovidersesResource, c.ns, clusterCapacityProviders), &v1alpha1.ClusterCapacityProviders{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCapacityProviders), err
}

// Update takes the representation of a clusterCapacityProviders and updates it. Returns the server's representation of the clusterCapacityProviders, and an error, if there is any.
func (c *FakeClusterCapacityProviderses) Update(ctx context.Context, clusterCapacityProviders *v1alpha1.ClusterCapacityProviders, opts v1.UpdateOptions) (result *v1alpha1.ClusterCapacityProviders, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustercapacityprovidersesResource, c.ns, clusterCapacityProviders), &v1alpha1.ClusterCapacityProviders{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCapacityProviders), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterCapacityProviderses) UpdateStatus(ctx context.Context, clusterCapacityProviders *v1alpha1.ClusterCapacityProviders, opts v1.UpdateOptions) (*v1alpha1.ClusterCapacityProviders, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clustercapacityprovidersesResource, "status", c.ns, clusterCapacityProviders), &v1alpha1.ClusterCapacityProviders{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCapacityProviders), err
}

// Delete takes name of the clusterCapacityProviders and deletes it. Returns an error if one occurs.
func (c *FakeClusterCapacityProviderses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustercapacityprovidersesResource, c.ns, name), &v1alpha1.ClusterCapacityProviders{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterCapacityProviderses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustercapacityprovidersesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterCapacityProvidersList{})
	return err
}

// Patch applies the patch and returns the patched clusterCapacityProviders.
func (c *FakeClusterCapacityProviderses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterCapacityProviders, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustercapacityprovidersesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterCapacityProviders{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterCapacityProviders), err
}