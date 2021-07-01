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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/signer/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSigningJobs implements SigningJobInterface
type FakeSigningJobs struct {
	Fake *FakeSignerV1alpha1
	ns   string
}

var signingjobsResource = schema.GroupVersionResource{Group: "signer.aws.kubeform.com", Version: "v1alpha1", Resource: "signingjobs"}

var signingjobsKind = schema.GroupVersionKind{Group: "signer.aws.kubeform.com", Version: "v1alpha1", Kind: "SigningJob"}

// Get takes name of the signingJob, and returns the corresponding signingJob object, and an error if there is any.
func (c *FakeSigningJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SigningJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(signingjobsResource, c.ns, name), &v1alpha1.SigningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningJob), err
}

// List takes label and field selectors, and returns the list of SigningJobs that match those selectors.
func (c *FakeSigningJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SigningJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(signingjobsResource, signingjobsKind, c.ns, opts), &v1alpha1.SigningJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SigningJobList{ListMeta: obj.(*v1alpha1.SigningJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.SigningJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested signingJobs.
func (c *FakeSigningJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(signingjobsResource, c.ns, opts))

}

// Create takes the representation of a signingJob and creates it.  Returns the server's representation of the signingJob, and an error, if there is any.
func (c *FakeSigningJobs) Create(ctx context.Context, signingJob *v1alpha1.SigningJob, opts v1.CreateOptions) (result *v1alpha1.SigningJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(signingjobsResource, c.ns, signingJob), &v1alpha1.SigningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningJob), err
}

// Update takes the representation of a signingJob and updates it. Returns the server's representation of the signingJob, and an error, if there is any.
func (c *FakeSigningJobs) Update(ctx context.Context, signingJob *v1alpha1.SigningJob, opts v1.UpdateOptions) (result *v1alpha1.SigningJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(signingjobsResource, c.ns, signingJob), &v1alpha1.SigningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSigningJobs) UpdateStatus(ctx context.Context, signingJob *v1alpha1.SigningJob, opts v1.UpdateOptions) (*v1alpha1.SigningJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(signingjobsResource, "status", c.ns, signingJob), &v1alpha1.SigningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningJob), err
}

// Delete takes name of the signingJob and deletes it. Returns an error if one occurs.
func (c *FakeSigningJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(signingjobsResource, c.ns, name), &v1alpha1.SigningJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSigningJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(signingjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SigningJobList{})
	return err
}

// Patch applies the patch and returns the patched signingJob.
func (c *FakeSigningJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SigningJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(signingjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SigningJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SigningJob), err
}
