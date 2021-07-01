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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"
)

// FakeAccountPublicAccessBlocks implements AccountPublicAccessBlockInterface
type FakeAccountPublicAccessBlocks struct {
	Fake *FakeS3V1alpha1
	ns   string
}

var accountpublicaccessblocksResource = schema.GroupVersionResource{Group: "s3.aws.kubeform.com", Version: "v1alpha1", Resource: "accountpublicaccessblocks"}

var accountpublicaccessblocksKind = schema.GroupVersionKind{Group: "s3.aws.kubeform.com", Version: "v1alpha1", Kind: "AccountPublicAccessBlock"}

// Get takes name of the accountPublicAccessBlock, and returns the corresponding accountPublicAccessBlock object, and an error if there is any.
func (c *FakeAccountPublicAccessBlocks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AccountPublicAccessBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(accountpublicaccessblocksResource, c.ns, name), &v1alpha1.AccountPublicAccessBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPublicAccessBlock), err
}

// List takes label and field selectors, and returns the list of AccountPublicAccessBlocks that match those selectors.
func (c *FakeAccountPublicAccessBlocks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AccountPublicAccessBlockList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(accountpublicaccessblocksResource, accountpublicaccessblocksKind, c.ns, opts), &v1alpha1.AccountPublicAccessBlockList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AccountPublicAccessBlockList{ListMeta: obj.(*v1alpha1.AccountPublicAccessBlockList).ListMeta}
	for _, item := range obj.(*v1alpha1.AccountPublicAccessBlockList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested accountPublicAccessBlocks.
func (c *FakeAccountPublicAccessBlocks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(accountpublicaccessblocksResource, c.ns, opts))

}

// Create takes the representation of a accountPublicAccessBlock and creates it.  Returns the server's representation of the accountPublicAccessBlock, and an error, if there is any.
func (c *FakeAccountPublicAccessBlocks) Create(ctx context.Context, accountPublicAccessBlock *v1alpha1.AccountPublicAccessBlock, opts v1.CreateOptions) (result *v1alpha1.AccountPublicAccessBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(accountpublicaccessblocksResource, c.ns, accountPublicAccessBlock), &v1alpha1.AccountPublicAccessBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPublicAccessBlock), err
}

// Update takes the representation of a accountPublicAccessBlock and updates it. Returns the server's representation of the accountPublicAccessBlock, and an error, if there is any.
func (c *FakeAccountPublicAccessBlocks) Update(ctx context.Context, accountPublicAccessBlock *v1alpha1.AccountPublicAccessBlock, opts v1.UpdateOptions) (result *v1alpha1.AccountPublicAccessBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(accountpublicaccessblocksResource, c.ns, accountPublicAccessBlock), &v1alpha1.AccountPublicAccessBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPublicAccessBlock), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAccountPublicAccessBlocks) UpdateStatus(ctx context.Context, accountPublicAccessBlock *v1alpha1.AccountPublicAccessBlock, opts v1.UpdateOptions) (*v1alpha1.AccountPublicAccessBlock, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(accountpublicaccessblocksResource, "status", c.ns, accountPublicAccessBlock), &v1alpha1.AccountPublicAccessBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPublicAccessBlock), err
}

// Delete takes name of the accountPublicAccessBlock and deletes it. Returns an error if one occurs.
func (c *FakeAccountPublicAccessBlocks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(accountpublicaccessblocksResource, c.ns, name), &v1alpha1.AccountPublicAccessBlock{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAccountPublicAccessBlocks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(accountpublicaccessblocksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AccountPublicAccessBlockList{})
	return err
}

// Patch applies the patch and returns the patched accountPublicAccessBlock.
func (c *FakeAccountPublicAccessBlocks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AccountPublicAccessBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(accountpublicaccessblocksResource, c.ns, name, pt, data, subresources...), &v1alpha1.AccountPublicAccessBlock{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AccountPublicAccessBlock), err
}
