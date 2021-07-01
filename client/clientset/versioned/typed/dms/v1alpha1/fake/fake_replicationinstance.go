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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/dms/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicationInstances implements ReplicationInstanceInterface
type FakeReplicationInstances struct {
	Fake *FakeDmsV1alpha1
	ns   string
}

var replicationinstancesResource = schema.GroupVersionResource{Group: "dms.aws.kubeform.com", Version: "v1alpha1", Resource: "replicationinstances"}

var replicationinstancesKind = schema.GroupVersionKind{Group: "dms.aws.kubeform.com", Version: "v1alpha1", Kind: "ReplicationInstance"}

// Get takes name of the replicationInstance, and returns the corresponding replicationInstance object, and an error if there is any.
func (c *FakeReplicationInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ReplicationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicationinstancesResource, c.ns, name), &v1alpha1.ReplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationInstance), err
}

// List takes label and field selectors, and returns the list of ReplicationInstances that match those selectors.
func (c *FakeReplicationInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReplicationInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicationinstancesResource, replicationinstancesKind, c.ns, opts), &v1alpha1.ReplicationInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReplicationInstanceList{ListMeta: obj.(*v1alpha1.ReplicationInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReplicationInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicationInstances.
func (c *FakeReplicationInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicationinstancesResource, c.ns, opts))

}

// Create takes the representation of a replicationInstance and creates it.  Returns the server's representation of the replicationInstance, and an error, if there is any.
func (c *FakeReplicationInstances) Create(ctx context.Context, replicationInstance *v1alpha1.ReplicationInstance, opts v1.CreateOptions) (result *v1alpha1.ReplicationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicationinstancesResource, c.ns, replicationInstance), &v1alpha1.ReplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationInstance), err
}

// Update takes the representation of a replicationInstance and updates it. Returns the server's representation of the replicationInstance, and an error, if there is any.
func (c *FakeReplicationInstances) Update(ctx context.Context, replicationInstance *v1alpha1.ReplicationInstance, opts v1.UpdateOptions) (result *v1alpha1.ReplicationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicationinstancesResource, c.ns, replicationInstance), &v1alpha1.ReplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicationInstances) UpdateStatus(ctx context.Context, replicationInstance *v1alpha1.ReplicationInstance, opts v1.UpdateOptions) (*v1alpha1.ReplicationInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicationinstancesResource, "status", c.ns, replicationInstance), &v1alpha1.ReplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationInstance), err
}

// Delete takes name of the replicationInstance and deletes it. Returns an error if one occurs.
func (c *FakeReplicationInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicationinstancesResource, c.ns, name), &v1alpha1.ReplicationInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicationInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicationinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReplicationInstanceList{})
	return err
}

// Patch applies the patch and returns the patched replicationInstance.
func (c *FakeReplicationInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ReplicationInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicationinstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ReplicationInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationInstance), err
}
