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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/lightsail/v1alpha1"
)

// FakeInstancePublicPortses implements InstancePublicPortsInterface
type FakeInstancePublicPortses struct {
	Fake *FakeLightsailV1alpha1
	ns   string
}

var instancepublicportsesResource = schema.GroupVersionResource{Group: "lightsail.aws.kubeform.com", Version: "v1alpha1", Resource: "instancepublicportses"}

var instancepublicportsesKind = schema.GroupVersionKind{Group: "lightsail.aws.kubeform.com", Version: "v1alpha1", Kind: "InstancePublicPorts"}

// Get takes name of the instancePublicPorts, and returns the corresponding instancePublicPorts object, and an error if there is any.
func (c *FakeInstancePublicPortses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.InstancePublicPorts, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(instancepublicportsesResource, c.ns, name), &v1alpha1.InstancePublicPorts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstancePublicPorts), err
}

// List takes label and field selectors, and returns the list of InstancePublicPortses that match those selectors.
func (c *FakeInstancePublicPortses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.InstancePublicPortsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(instancepublicportsesResource, instancepublicportsesKind, c.ns, opts), &v1alpha1.InstancePublicPortsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstancePublicPortsList{ListMeta: obj.(*v1alpha1.InstancePublicPortsList).ListMeta}
	for _, item := range obj.(*v1alpha1.InstancePublicPortsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instancePublicPortses.
func (c *FakeInstancePublicPortses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(instancepublicportsesResource, c.ns, opts))

}

// Create takes the representation of a instancePublicPorts and creates it.  Returns the server's representation of the instancePublicPorts, and an error, if there is any.
func (c *FakeInstancePublicPortses) Create(ctx context.Context, instancePublicPorts *v1alpha1.InstancePublicPorts, opts v1.CreateOptions) (result *v1alpha1.InstancePublicPorts, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(instancepublicportsesResource, c.ns, instancePublicPorts), &v1alpha1.InstancePublicPorts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstancePublicPorts), err
}

// Update takes the representation of a instancePublicPorts and updates it. Returns the server's representation of the instancePublicPorts, and an error, if there is any.
func (c *FakeInstancePublicPortses) Update(ctx context.Context, instancePublicPorts *v1alpha1.InstancePublicPorts, opts v1.UpdateOptions) (result *v1alpha1.InstancePublicPorts, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(instancepublicportsesResource, c.ns, instancePublicPorts), &v1alpha1.InstancePublicPorts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstancePublicPorts), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInstancePublicPortses) UpdateStatus(ctx context.Context, instancePublicPorts *v1alpha1.InstancePublicPorts, opts v1.UpdateOptions) (*v1alpha1.InstancePublicPorts, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(instancepublicportsesResource, "status", c.ns, instancePublicPorts), &v1alpha1.InstancePublicPorts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstancePublicPorts), err
}

// Delete takes name of the instancePublicPorts and deletes it. Returns an error if one occurs.
func (c *FakeInstancePublicPortses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(instancepublicportsesResource, c.ns, name), &v1alpha1.InstancePublicPorts{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstancePublicPortses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(instancepublicportsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstancePublicPortsList{})
	return err
}

// Patch applies the patch and returns the patched instancePublicPorts.
func (c *FakeInstancePublicPortses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.InstancePublicPorts, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(instancepublicportsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.InstancePublicPorts{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstancePublicPorts), err
}
