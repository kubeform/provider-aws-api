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

	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"
	scheme "kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TransitGatewayRouteTablesGetter has a method to return a TransitGatewayRouteTableInterface.
// A group's client should implement this interface.
type TransitGatewayRouteTablesGetter interface {
	TransitGatewayRouteTables(namespace string) TransitGatewayRouteTableInterface
}

// TransitGatewayRouteTableInterface has methods to work with TransitGatewayRouteTable resources.
type TransitGatewayRouteTableInterface interface {
	Create(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.CreateOptions) (*v1alpha1.TransitGatewayRouteTable, error)
	Update(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.UpdateOptions) (*v1alpha1.TransitGatewayRouteTable, error)
	UpdateStatus(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.UpdateOptions) (*v1alpha1.TransitGatewayRouteTable, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.TransitGatewayRouteTable, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.TransitGatewayRouteTableList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitGatewayRouteTable, err error)
	TransitGatewayRouteTableExpansion
}

// transitGatewayRouteTables implements TransitGatewayRouteTableInterface
type transitGatewayRouteTables struct {
	client rest.Interface
	ns     string
}

// newTransitGatewayRouteTables returns a TransitGatewayRouteTables
func newTransitGatewayRouteTables(c *Ec2V1alpha1Client, namespace string) *transitGatewayRouteTables {
	return &transitGatewayRouteTables{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the transitGatewayRouteTable, and returns the corresponding transitGatewayRouteTable object, and an error if there is any.
func (c *transitGatewayRouteTables) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TransitGatewayRouteTables that match those selectors.
func (c *transitGatewayRouteTables) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TransitGatewayRouteTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TransitGatewayRouteTableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested transitGatewayRouteTables.
func (c *transitGatewayRouteTables) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a transitGatewayRouteTable and creates it.  Returns the server's representation of the transitGatewayRouteTable, and an error, if there is any.
func (c *transitGatewayRouteTables) Create(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.CreateOptions) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(transitGatewayRouteTable).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a transitGatewayRouteTable and updates it. Returns the server's representation of the transitGatewayRouteTable, and an error, if there is any.
func (c *transitGatewayRouteTables) Update(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.UpdateOptions) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(transitGatewayRouteTable.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(transitGatewayRouteTable).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *transitGatewayRouteTables) UpdateStatus(ctx context.Context, transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable, opts v1.UpdateOptions) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(transitGatewayRouteTable.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(transitGatewayRouteTable).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the transitGatewayRouteTable and deletes it. Returns an error if one occurs.
func (c *transitGatewayRouteTables) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *transitGatewayRouteTables) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched transitGatewayRouteTable.
func (c *transitGatewayRouteTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
