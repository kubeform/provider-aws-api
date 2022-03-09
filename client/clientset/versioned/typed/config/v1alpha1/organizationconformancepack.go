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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-aws-api/apis/config/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OrganizationConformancePacksGetter has a method to return a OrganizationConformancePackInterface.
// A group's client should implement this interface.
type OrganizationConformancePacksGetter interface {
	OrganizationConformancePacks(namespace string) OrganizationConformancePackInterface
}

// OrganizationConformancePackInterface has methods to work with OrganizationConformancePack resources.
type OrganizationConformancePackInterface interface {
	Create(ctx context.Context, organizationConformancePack *v1alpha1.OrganizationConformancePack, opts v1.CreateOptions) (*v1alpha1.OrganizationConformancePack, error)
	Update(ctx context.Context, organizationConformancePack *v1alpha1.OrganizationConformancePack, opts v1.UpdateOptions) (*v1alpha1.OrganizationConformancePack, error)
	UpdateStatus(ctx context.Context, organizationConformancePack *v1alpha1.OrganizationConformancePack, opts v1.UpdateOptions) (*v1alpha1.OrganizationConformancePack, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OrganizationConformancePack, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OrganizationConformancePackList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrganizationConformancePack, err error)
	OrganizationConformancePackExpansion
}

// organizationConformancePacks implements OrganizationConformancePackInterface
type organizationConformancePacks struct {
	client rest.Interface
	ns     string
}

// newOrganizationConformancePacks returns a OrganizationConformancePacks
func newOrganizationConformancePacks(c *ConfigV1alpha1Client, namespace string) *organizationConformancePacks {
	return &organizationConformancePacks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the organizationConformancePack, and returns the corresponding organizationConformancePack object, and an error if there is any.
func (c *organizationConformancePacks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OrganizationConformancePack, err error) {
	result = &v1alpha1.OrganizationConformancePack{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OrganizationConformancePacks that match those selectors.
func (c *organizationConformancePacks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrganizationConformancePackList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OrganizationConformancePackList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested organizationConformancePacks.
func (c *organizationConformancePacks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a organizationConformancePack and creates it.  Returns the server's representation of the organizationConformancePack, and an error, if there is any.
func (c *organizationConformancePacks) Create(ctx context.Context, organizationConformancePack *v1alpha1.OrganizationConformancePack, opts v1.CreateOptions) (result *v1alpha1.OrganizationConformancePack, err error) {
	result = &v1alpha1.OrganizationConformancePack{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(organizationConformancePack).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a organizationConformancePack and updates it. Returns the server's representation of the organizationConformancePack, and an error, if there is any.
func (c *organizationConformancePacks) Update(ctx context.Context, organizationConformancePack *v1alpha1.OrganizationConformancePack, opts v1.UpdateOptions) (result *v1alpha1.OrganizationConformancePack, err error) {
	result = &v1alpha1.OrganizationConformancePack{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		Name(organizationConformancePack.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(organizationConformancePack).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *organizationConformancePacks) UpdateStatus(ctx context.Context, organizationConformancePack *v1alpha1.OrganizationConformancePack, opts v1.UpdateOptions) (result *v1alpha1.OrganizationConformancePack, err error) {
	result = &v1alpha1.OrganizationConformancePack{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		Name(organizationConformancePack.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(organizationConformancePack).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the organizationConformancePack and deletes it. Returns an error if one occurs.
func (c *organizationConformancePacks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *organizationConformancePacks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched organizationConformancePack.
func (c *organizationConformancePacks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrganizationConformancePack, err error) {
	result = &v1alpha1.OrganizationConformancePack{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("organizationconformancepacks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
