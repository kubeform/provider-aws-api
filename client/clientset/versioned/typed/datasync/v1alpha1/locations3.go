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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/datasync/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"
)

// LocationS3sGetter has a method to return a LocationS3Interface.
// A group's client should implement this interface.
type LocationS3sGetter interface {
	LocationS3s(namespace string) LocationS3Interface
}

// LocationS3Interface has methods to work with LocationS3 resources.
type LocationS3Interface interface {
	Create(ctx context.Context, locationS3 *v1alpha1.LocationS3, opts v1.CreateOptions) (*v1alpha1.LocationS3, error)
	Update(ctx context.Context, locationS3 *v1alpha1.LocationS3, opts v1.UpdateOptions) (*v1alpha1.LocationS3, error)
	UpdateStatus(ctx context.Context, locationS3 *v1alpha1.LocationS3, opts v1.UpdateOptions) (*v1alpha1.LocationS3, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.LocationS3, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LocationS3List, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocationS3, err error)
	LocationS3Expansion
}

// locationS3s implements LocationS3Interface
type locationS3s struct {
	client rest.Interface
	ns     string
}

// newLocationS3s returns a LocationS3s
func newLocationS3s(c *DatasyncV1alpha1Client, namespace string) *locationS3s {
	return &locationS3s{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the locationS3, and returns the corresponding locationS3 object, and an error if there is any.
func (c *locationS3s) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LocationS3, err error) {
	result = &v1alpha1.LocationS3{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("locations3s").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocationS3s that match those selectors.
func (c *locationS3s) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LocationS3List, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LocationS3List{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("locations3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested locationS3s.
func (c *locationS3s) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("locations3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a locationS3 and creates it.  Returns the server's representation of the locationS3, and an error, if there is any.
func (c *locationS3s) Create(ctx context.Context, locationS3 *v1alpha1.LocationS3, opts v1.CreateOptions) (result *v1alpha1.LocationS3, err error) {
	result = &v1alpha1.LocationS3{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("locations3s").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(locationS3).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a locationS3 and updates it. Returns the server's representation of the locationS3, and an error, if there is any.
func (c *locationS3s) Update(ctx context.Context, locationS3 *v1alpha1.LocationS3, opts v1.UpdateOptions) (result *v1alpha1.LocationS3, err error) {
	result = &v1alpha1.LocationS3{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("locations3s").
		Name(locationS3.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(locationS3).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *locationS3s) UpdateStatus(ctx context.Context, locationS3 *v1alpha1.LocationS3, opts v1.UpdateOptions) (result *v1alpha1.LocationS3, err error) {
	result = &v1alpha1.LocationS3{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("locations3s").
		Name(locationS3.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(locationS3).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the locationS3 and deletes it. Returns an error if one occurs.
func (c *locationS3s) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("locations3s").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *locationS3s) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("locations3s").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched locationS3.
func (c *locationS3s) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LocationS3, err error) {
	result = &v1alpha1.LocationS3{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("locations3s").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
