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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/macie2/v1alpha1"
)

// FakeClassificationJobs implements ClassificationJobInterface
type FakeClassificationJobs struct {
	Fake *FakeMacie2V1alpha1
	ns   string
}

var classificationjobsResource = schema.GroupVersionResource{Group: "macie2.aws.kubeform.com", Version: "v1alpha1", Resource: "classificationjobs"}

var classificationjobsKind = schema.GroupVersionKind{Group: "macie2.aws.kubeform.com", Version: "v1alpha1", Kind: "ClassificationJob"}

// Get takes name of the classificationJob, and returns the corresponding classificationJob object, and an error if there is any.
func (c *FakeClassificationJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClassificationJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(classificationjobsResource, c.ns, name), &v1alpha1.ClassificationJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClassificationJob), err
}

// List takes label and field selectors, and returns the list of ClassificationJobs that match those selectors.
func (c *FakeClassificationJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClassificationJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(classificationjobsResource, classificationjobsKind, c.ns, opts), &v1alpha1.ClassificationJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClassificationJobList{ListMeta: obj.(*v1alpha1.ClassificationJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClassificationJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested classificationJobs.
func (c *FakeClassificationJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(classificationjobsResource, c.ns, opts))

}

// Create takes the representation of a classificationJob and creates it.  Returns the server's representation of the classificationJob, and an error, if there is any.
func (c *FakeClassificationJobs) Create(ctx context.Context, classificationJob *v1alpha1.ClassificationJob, opts v1.CreateOptions) (result *v1alpha1.ClassificationJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(classificationjobsResource, c.ns, classificationJob), &v1alpha1.ClassificationJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClassificationJob), err
}

// Update takes the representation of a classificationJob and updates it. Returns the server's representation of the classificationJob, and an error, if there is any.
func (c *FakeClassificationJobs) Update(ctx context.Context, classificationJob *v1alpha1.ClassificationJob, opts v1.UpdateOptions) (result *v1alpha1.ClassificationJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(classificationjobsResource, c.ns, classificationJob), &v1alpha1.ClassificationJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClassificationJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClassificationJobs) UpdateStatus(ctx context.Context, classificationJob *v1alpha1.ClassificationJob, opts v1.UpdateOptions) (*v1alpha1.ClassificationJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(classificationjobsResource, "status", c.ns, classificationJob), &v1alpha1.ClassificationJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClassificationJob), err
}

// Delete takes name of the classificationJob and deletes it. Returns an error if one occurs.
func (c *FakeClassificationJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(classificationjobsResource, c.ns, name), &v1alpha1.ClassificationJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClassificationJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(classificationjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClassificationJobList{})
	return err
}

// Patch applies the patch and returns the patched classificationJob.
func (c *FakeClassificationJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClassificationJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(classificationjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClassificationJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClassificationJob), err
}
