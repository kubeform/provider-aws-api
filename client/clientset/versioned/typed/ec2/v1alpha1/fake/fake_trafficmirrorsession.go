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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"
)

// FakeTrafficMirrorSessions implements TrafficMirrorSessionInterface
type FakeTrafficMirrorSessions struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var trafficmirrorsessionsResource = schema.GroupVersionResource{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Resource: "trafficmirrorsessions"}

var trafficmirrorsessionsKind = schema.GroupVersionKind{Group: "ec2.aws.kubeform.com", Version: "v1alpha1", Kind: "TrafficMirrorSession"}

// Get takes name of the trafficMirrorSession, and returns the corresponding trafficMirrorSession object, and an error if there is any.
func (c *FakeTrafficMirrorSessions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TrafficMirrorSession, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(trafficmirrorsessionsResource, c.ns, name), &v1alpha1.TrafficMirrorSession{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorSession), err
}

// List takes label and field selectors, and returns the list of TrafficMirrorSessions that match those selectors.
func (c *FakeTrafficMirrorSessions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TrafficMirrorSessionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(trafficmirrorsessionsResource, trafficmirrorsessionsKind, c.ns, opts), &v1alpha1.TrafficMirrorSessionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TrafficMirrorSessionList{ListMeta: obj.(*v1alpha1.TrafficMirrorSessionList).ListMeta}
	for _, item := range obj.(*v1alpha1.TrafficMirrorSessionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested trafficMirrorSessions.
func (c *FakeTrafficMirrorSessions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(trafficmirrorsessionsResource, c.ns, opts))

}

// Create takes the representation of a trafficMirrorSession and creates it.  Returns the server's representation of the trafficMirrorSession, and an error, if there is any.
func (c *FakeTrafficMirrorSessions) Create(ctx context.Context, trafficMirrorSession *v1alpha1.TrafficMirrorSession, opts v1.CreateOptions) (result *v1alpha1.TrafficMirrorSession, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(trafficmirrorsessionsResource, c.ns, trafficMirrorSession), &v1alpha1.TrafficMirrorSession{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorSession), err
}

// Update takes the representation of a trafficMirrorSession and updates it. Returns the server's representation of the trafficMirrorSession, and an error, if there is any.
func (c *FakeTrafficMirrorSessions) Update(ctx context.Context, trafficMirrorSession *v1alpha1.TrafficMirrorSession, opts v1.UpdateOptions) (result *v1alpha1.TrafficMirrorSession, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(trafficmirrorsessionsResource, c.ns, trafficMirrorSession), &v1alpha1.TrafficMirrorSession{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorSession), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTrafficMirrorSessions) UpdateStatus(ctx context.Context, trafficMirrorSession *v1alpha1.TrafficMirrorSession, opts v1.UpdateOptions) (*v1alpha1.TrafficMirrorSession, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(trafficmirrorsessionsResource, "status", c.ns, trafficMirrorSession), &v1alpha1.TrafficMirrorSession{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorSession), err
}

// Delete takes name of the trafficMirrorSession and deletes it. Returns an error if one occurs.
func (c *FakeTrafficMirrorSessions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(trafficmirrorsessionsResource, c.ns, name), &v1alpha1.TrafficMirrorSession{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTrafficMirrorSessions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(trafficmirrorsessionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TrafficMirrorSessionList{})
	return err
}

// Patch applies the patch and returns the patched trafficMirrorSession.
func (c *FakeTrafficMirrorSessions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TrafficMirrorSession, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trafficmirrorsessionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TrafficMirrorSession{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorSession), err
}
