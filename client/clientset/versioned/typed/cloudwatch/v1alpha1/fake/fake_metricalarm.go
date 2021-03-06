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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudwatch/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMetricAlarms implements MetricAlarmInterface
type FakeMetricAlarms struct {
	Fake *FakeCloudwatchV1alpha1
	ns   string
}

var metricalarmsResource = schema.GroupVersionResource{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Resource: "metricalarms"}

var metricalarmsKind = schema.GroupVersionKind{Group: "cloudwatch.aws.kubeform.com", Version: "v1alpha1", Kind: "MetricAlarm"}

// Get takes name of the metricAlarm, and returns the corresponding metricAlarm object, and an error if there is any.
func (c *FakeMetricAlarms) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MetricAlarm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(metricalarmsResource, c.ns, name), &v1alpha1.MetricAlarm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlarm), err
}

// List takes label and field selectors, and returns the list of MetricAlarms that match those selectors.
func (c *FakeMetricAlarms) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MetricAlarmList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(metricalarmsResource, metricalarmsKind, c.ns, opts), &v1alpha1.MetricAlarmList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MetricAlarmList{ListMeta: obj.(*v1alpha1.MetricAlarmList).ListMeta}
	for _, item := range obj.(*v1alpha1.MetricAlarmList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested metricAlarms.
func (c *FakeMetricAlarms) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(metricalarmsResource, c.ns, opts))

}

// Create takes the representation of a metricAlarm and creates it.  Returns the server's representation of the metricAlarm, and an error, if there is any.
func (c *FakeMetricAlarms) Create(ctx context.Context, metricAlarm *v1alpha1.MetricAlarm, opts v1.CreateOptions) (result *v1alpha1.MetricAlarm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(metricalarmsResource, c.ns, metricAlarm), &v1alpha1.MetricAlarm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlarm), err
}

// Update takes the representation of a metricAlarm and updates it. Returns the server's representation of the metricAlarm, and an error, if there is any.
func (c *FakeMetricAlarms) Update(ctx context.Context, metricAlarm *v1alpha1.MetricAlarm, opts v1.UpdateOptions) (result *v1alpha1.MetricAlarm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(metricalarmsResource, c.ns, metricAlarm), &v1alpha1.MetricAlarm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlarm), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMetricAlarms) UpdateStatus(ctx context.Context, metricAlarm *v1alpha1.MetricAlarm, opts v1.UpdateOptions) (*v1alpha1.MetricAlarm, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(metricalarmsResource, "status", c.ns, metricAlarm), &v1alpha1.MetricAlarm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlarm), err
}

// Delete takes name of the metricAlarm and deletes it. Returns an error if one occurs.
func (c *FakeMetricAlarms) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(metricalarmsResource, c.ns, name), &v1alpha1.MetricAlarm{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMetricAlarms) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(metricalarmsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MetricAlarmList{})
	return err
}

// Patch applies the patch and returns the patched metricAlarm.
func (c *FakeMetricAlarms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MetricAlarm, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(metricalarmsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MetricAlarm{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MetricAlarm), err
}
