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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/elasticache/v1alpha1"
)

// FakeClusters implements ClusterInterface
type FakeClusters struct {
	Fake *FakeElasticacheV1alpha1
	ns   string
}

var clustersResource = schema.GroupVersionResource{Group: "elasticache.aws.kubeform.com", Version: "v1alpha1", Resource: "clusters"}

var clustersKind = schema.GroupVersionKind{Group: "elasticache.aws.kubeform.com", Version: "v1alpha1", Kind: "Cluster"}

// Get takes name of the cluster, and returns the corresponding cluster object, and an error if there is any.
func (c *FakeClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustersResource, c.ns, name), &v1alpha1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cluster), err
}

// List takes label and field selectors, and returns the list of Clusters that match those selectors.
func (c *FakeClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustersResource, clustersKind, c.ns, opts), &v1alpha1.ClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterList{ListMeta: obj.(*v1alpha1.ClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusters.
func (c *FakeClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustersResource, c.ns, opts))

}

// Create takes the representation of a cluster and creates it.  Returns the server's representation of the cluster, and an error, if there is any.
func (c *FakeClusters) Create(ctx context.Context, cluster *v1alpha1.Cluster, opts v1.CreateOptions) (result *v1alpha1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustersResource, c.ns, cluster), &v1alpha1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cluster), err
}

// Update takes the representation of a cluster and updates it. Returns the server's representation of the cluster, and an error, if there is any.
func (c *FakeClusters) Update(ctx context.Context, cluster *v1alpha1.Cluster, opts v1.UpdateOptions) (result *v1alpha1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustersResource, c.ns, cluster), &v1alpha1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusters) UpdateStatus(ctx context.Context, cluster *v1alpha1.Cluster, opts v1.UpdateOptions) (*v1alpha1.Cluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clustersResource, "status", c.ns, cluster), &v1alpha1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cluster), err
}

// Delete takes name of the cluster and deletes it. Returns an error if one occurs.
func (c *FakeClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clustersResource, c.ns, name), &v1alpha1.Cluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterList{})
	return err
}

// Patch applies the patch and returns the patched cluster.
func (c *FakeClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Cluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.Cluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Cluster), err
}
