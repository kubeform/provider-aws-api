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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/glaciervault/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"
)

// LocksGetter has a method to return a LockInterface.
// A group's client should implement this interface.
type LocksGetter interface {
	Locks(namespace string) LockInterface
}

// LockInterface has methods to work with Lock resources.
type LockInterface interface {
	Create(ctx context.Context, lock *v1alpha1.Lock, opts v1.CreateOptions) (*v1alpha1.Lock, error)
	Update(ctx context.Context, lock *v1alpha1.Lock, opts v1.UpdateOptions) (*v1alpha1.Lock, error)
	UpdateStatus(ctx context.Context, lock *v1alpha1.Lock, opts v1.UpdateOptions) (*v1alpha1.Lock, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Lock, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.LockList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lock, err error)
	LockExpansion
}

// locks implements LockInterface
type locks struct {
	client rest.Interface
	ns     string
}

// newLocks returns a Locks
func newLocks(c *GlaciervaultV1alpha1Client, namespace string) *locks {
	return &locks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the lock, and returns the corresponding lock object, and an error if there is any.
func (c *locks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Lock, err error) {
	result = &v1alpha1.Lock{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("locks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Locks that match those selectors.
func (c *locks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LockList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LockList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("locks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested locks.
func (c *locks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("locks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a lock and creates it.  Returns the server's representation of the lock, and an error, if there is any.
func (c *locks) Create(ctx context.Context, lock *v1alpha1.Lock, opts v1.CreateOptions) (result *v1alpha1.Lock, err error) {
	result = &v1alpha1.Lock{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("locks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lock).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a lock and updates it. Returns the server's representation of the lock, and an error, if there is any.
func (c *locks) Update(ctx context.Context, lock *v1alpha1.Lock, opts v1.UpdateOptions) (result *v1alpha1.Lock, err error) {
	result = &v1alpha1.Lock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("locks").
		Name(lock.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lock).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *locks) UpdateStatus(ctx context.Context, lock *v1alpha1.Lock, opts v1.UpdateOptions) (result *v1alpha1.Lock, err error) {
	result = &v1alpha1.Lock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("locks").
		Name(lock.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(lock).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the lock and deletes it. Returns an error if one occurs.
func (c *locks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("locks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *locks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("locks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched lock.
func (c *locks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Lock, err error) {
	result = &v1alpha1.Lock{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("locks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
