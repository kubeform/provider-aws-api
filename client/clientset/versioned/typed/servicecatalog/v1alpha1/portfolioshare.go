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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/servicecatalog/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PortfolioSharesGetter has a method to return a PortfolioShareInterface.
// A group's client should implement this interface.
type PortfolioSharesGetter interface {
	PortfolioShares(namespace string) PortfolioShareInterface
}

// PortfolioShareInterface has methods to work with PortfolioShare resources.
type PortfolioShareInterface interface {
	Create(ctx context.Context, portfolioShare *v1alpha1.PortfolioShare, opts v1.CreateOptions) (*v1alpha1.PortfolioShare, error)
	Update(ctx context.Context, portfolioShare *v1alpha1.PortfolioShare, opts v1.UpdateOptions) (*v1alpha1.PortfolioShare, error)
	UpdateStatus(ctx context.Context, portfolioShare *v1alpha1.PortfolioShare, opts v1.UpdateOptions) (*v1alpha1.PortfolioShare, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PortfolioShare, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PortfolioShareList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PortfolioShare, err error)
	PortfolioShareExpansion
}

// portfolioShares implements PortfolioShareInterface
type portfolioShares struct {
	client rest.Interface
	ns     string
}

// newPortfolioShares returns a PortfolioShares
func newPortfolioShares(c *ServicecatalogV1alpha1Client, namespace string) *portfolioShares {
	return &portfolioShares{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the portfolioShare, and returns the corresponding portfolioShare object, and an error if there is any.
func (c *portfolioShares) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PortfolioShare, err error) {
	result = &v1alpha1.PortfolioShare{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("portfolioshares").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PortfolioShares that match those selectors.
func (c *portfolioShares) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PortfolioShareList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PortfolioShareList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("portfolioshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested portfolioShares.
func (c *portfolioShares) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("portfolioshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a portfolioShare and creates it.  Returns the server's representation of the portfolioShare, and an error, if there is any.
func (c *portfolioShares) Create(ctx context.Context, portfolioShare *v1alpha1.PortfolioShare, opts v1.CreateOptions) (result *v1alpha1.PortfolioShare, err error) {
	result = &v1alpha1.PortfolioShare{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("portfolioshares").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(portfolioShare).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a portfolioShare and updates it. Returns the server's representation of the portfolioShare, and an error, if there is any.
func (c *portfolioShares) Update(ctx context.Context, portfolioShare *v1alpha1.PortfolioShare, opts v1.UpdateOptions) (result *v1alpha1.PortfolioShare, err error) {
	result = &v1alpha1.PortfolioShare{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("portfolioshares").
		Name(portfolioShare.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(portfolioShare).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *portfolioShares) UpdateStatus(ctx context.Context, portfolioShare *v1alpha1.PortfolioShare, opts v1.UpdateOptions) (result *v1alpha1.PortfolioShare, err error) {
	result = &v1alpha1.PortfolioShare{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("portfolioshares").
		Name(portfolioShare.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(portfolioShare).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the portfolioShare and deletes it. Returns an error if one occurs.
func (c *portfolioShares) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("portfolioshares").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *portfolioShares) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("portfolioshares").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched portfolioShare.
func (c *portfolioShares) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PortfolioShare, err error) {
	result = &v1alpha1.PortfolioShare{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("portfolioshares").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
