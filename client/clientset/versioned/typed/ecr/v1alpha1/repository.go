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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecr/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"
)

// RepositoriesGetter has a method to return a RepositoryInterface.
// A group's client should implement this interface.
type RepositoriesGetter interface {
	Repositories(namespace string) RepositoryInterface
}

// RepositoryInterface has methods to work with Repository resources.
type RepositoryInterface interface {
	Create(ctx context.Context, repository *v1alpha1.Repository, opts v1.CreateOptions) (*v1alpha1.Repository, error)
	Update(ctx context.Context, repository *v1alpha1.Repository, opts v1.UpdateOptions) (*v1alpha1.Repository, error)
	UpdateStatus(ctx context.Context, repository *v1alpha1.Repository, opts v1.UpdateOptions) (*v1alpha1.Repository, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Repository, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RepositoryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Repository, err error)
	RepositoryExpansion
}

// repositories implements RepositoryInterface
type repositories struct {
	client rest.Interface
	ns     string
}

// newRepositories returns a Repositories
func newRepositories(c *EcrV1alpha1Client, namespace string) *repositories {
	return &repositories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the repository, and returns the corresponding repository object, and an error if there is any.
func (c *repositories) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Repository, err error) {
	result = &v1alpha1.Repository{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("repositories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Repositories that match those selectors.
func (c *repositories) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RepositoryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RepositoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("repositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested repositories.
func (c *repositories) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("repositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a repository and creates it.  Returns the server's representation of the repository, and an error, if there is any.
func (c *repositories) Create(ctx context.Context, repository *v1alpha1.Repository, opts v1.CreateOptions) (result *v1alpha1.Repository, err error) {
	result = &v1alpha1.Repository{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("repositories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(repository).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a repository and updates it. Returns the server's representation of the repository, and an error, if there is any.
func (c *repositories) Update(ctx context.Context, repository *v1alpha1.Repository, opts v1.UpdateOptions) (result *v1alpha1.Repository, err error) {
	result = &v1alpha1.Repository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("repositories").
		Name(repository.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(repository).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *repositories) UpdateStatus(ctx context.Context, repository *v1alpha1.Repository, opts v1.UpdateOptions) (result *v1alpha1.Repository, err error) {
	result = &v1alpha1.Repository{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("repositories").
		Name(repository.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(repository).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the repository and deletes it. Returns an error if one occurs.
func (c *repositories) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("repositories").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *repositories) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("repositories").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched repository.
func (c *repositories) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Repository, err error) {
	result = &v1alpha1.Repository{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("repositories").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
