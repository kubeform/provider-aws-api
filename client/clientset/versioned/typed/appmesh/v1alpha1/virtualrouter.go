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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/appmesh/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"
)

// VirtualRoutersGetter has a method to return a VirtualRouterInterface.
// A group's client should implement this interface.
type VirtualRoutersGetter interface {
	VirtualRouters(namespace string) VirtualRouterInterface
}

// VirtualRouterInterface has methods to work with VirtualRouter resources.
type VirtualRouterInterface interface {
	Create(ctx context.Context, virtualRouter *v1alpha1.VirtualRouter, opts v1.CreateOptions) (*v1alpha1.VirtualRouter, error)
	Update(ctx context.Context, virtualRouter *v1alpha1.VirtualRouter, opts v1.UpdateOptions) (*v1alpha1.VirtualRouter, error)
	UpdateStatus(ctx context.Context, virtualRouter *v1alpha1.VirtualRouter, opts v1.UpdateOptions) (*v1alpha1.VirtualRouter, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VirtualRouter, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VirtualRouterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualRouter, err error)
	VirtualRouterExpansion
}

// virtualRouters implements VirtualRouterInterface
type virtualRouters struct {
	client rest.Interface
	ns     string
}

// newVirtualRouters returns a VirtualRouters
func newVirtualRouters(c *AppmeshV1alpha1Client, namespace string) *virtualRouters {
	return &virtualRouters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualRouter, and returns the corresponding virtualRouter object, and an error if there is any.
func (c *virtualRouters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualRouters that match those selectors.
func (c *virtualRouters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VirtualRouterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualRouterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualRouters.
func (c *virtualRouters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualRouter and creates it.  Returns the server's representation of the virtualRouter, and an error, if there is any.
func (c *virtualRouters) Create(ctx context.Context, virtualRouter *v1alpha1.VirtualRouter, opts v1.CreateOptions) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualRouter).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualRouter and updates it. Returns the server's representation of the virtualRouter, and an error, if there is any.
func (c *virtualRouters) Update(ctx context.Context, virtualRouter *v1alpha1.VirtualRouter, opts v1.UpdateOptions) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(virtualRouter.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualRouter).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *virtualRouters) UpdateStatus(ctx context.Context, virtualRouter *v1alpha1.VirtualRouter, opts v1.UpdateOptions) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(virtualRouter.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualRouter).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualRouter and deletes it. Returns an error if one occurs.
func (c *virtualRouters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualRouters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualrouters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualRouter.
func (c *virtualRouters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
