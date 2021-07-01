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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/neptune/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"
)

// ClusterInstancesGetter has a method to return a ClusterInstanceInterface.
// A group's client should implement this interface.
type ClusterInstancesGetter interface {
	ClusterInstances(namespace string) ClusterInstanceInterface
}

// ClusterInstanceInterface has methods to work with ClusterInstance resources.
type ClusterInstanceInterface interface {
	Create(ctx context.Context, clusterInstance *v1alpha1.ClusterInstance, opts v1.CreateOptions) (*v1alpha1.ClusterInstance, error)
	Update(ctx context.Context, clusterInstance *v1alpha1.ClusterInstance, opts v1.UpdateOptions) (*v1alpha1.ClusterInstance, error)
	UpdateStatus(ctx context.Context, clusterInstance *v1alpha1.ClusterInstance, opts v1.UpdateOptions) (*v1alpha1.ClusterInstance, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ClusterInstance, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ClusterInstanceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterInstance, err error)
	ClusterInstanceExpansion
}

// clusterInstances implements ClusterInstanceInterface
type clusterInstances struct {
	client rest.Interface
	ns     string
}

// newClusterInstances returns a ClusterInstances
func newClusterInstances(c *NeptuneV1alpha1Client, namespace string) *clusterInstances {
	return &clusterInstances{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterInstance, and returns the corresponding clusterInstance object, and an error if there is any.
func (c *clusterInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterInstance, err error) {
	result = &v1alpha1.ClusterInstance{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterinstances").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterInstances that match those selectors.
func (c *clusterInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterInstanceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterInstanceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterInstances.
func (c *clusterInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterInstance and creates it.  Returns the server's representation of the clusterInstance, and an error, if there is any.
func (c *clusterInstances) Create(ctx context.Context, clusterInstance *v1alpha1.ClusterInstance, opts v1.CreateOptions) (result *v1alpha1.ClusterInstance, err error) {
	result = &v1alpha1.ClusterInstance{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterinstances").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterInstance).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterInstance and updates it. Returns the server's representation of the clusterInstance, and an error, if there is any.
func (c *clusterInstances) Update(ctx context.Context, clusterInstance *v1alpha1.ClusterInstance, opts v1.UpdateOptions) (result *v1alpha1.ClusterInstance, err error) {
	result = &v1alpha1.ClusterInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterinstances").
		Name(clusterInstance.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterInstance).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterInstances) UpdateStatus(ctx context.Context, clusterInstance *v1alpha1.ClusterInstance, opts v1.UpdateOptions) (result *v1alpha1.ClusterInstance, err error) {
	result = &v1alpha1.ClusterInstance{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterinstances").
		Name(clusterInstance.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterInstance).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterInstance and deletes it. Returns an error if one occurs.
func (c *clusterInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterinstances").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterinstances").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterInstance.
func (c *clusterInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterInstance, err error) {
	result = &v1alpha1.ClusterInstance{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterinstances").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
