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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-aws-api/apis/vpc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IpamPreviewNextCIDRLister helps list IpamPreviewNextCIDRs.
// All objects returned here must be treated as read-only.
type IpamPreviewNextCIDRLister interface {
	// List lists all IpamPreviewNextCIDRs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IpamPreviewNextCIDR, err error)
	// IpamPreviewNextCIDRs returns an object that can list and get IpamPreviewNextCIDRs.
	IpamPreviewNextCIDRs(namespace string) IpamPreviewNextCIDRNamespaceLister
	IpamPreviewNextCIDRListerExpansion
}

// ipamPreviewNextCIDRLister implements the IpamPreviewNextCIDRLister interface.
type ipamPreviewNextCIDRLister struct {
	indexer cache.Indexer
}

// NewIpamPreviewNextCIDRLister returns a new IpamPreviewNextCIDRLister.
func NewIpamPreviewNextCIDRLister(indexer cache.Indexer) IpamPreviewNextCIDRLister {
	return &ipamPreviewNextCIDRLister{indexer: indexer}
}

// List lists all IpamPreviewNextCIDRs in the indexer.
func (s *ipamPreviewNextCIDRLister) List(selector labels.Selector) (ret []*v1alpha1.IpamPreviewNextCIDR, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IpamPreviewNextCIDR))
	})
	return ret, err
}

// IpamPreviewNextCIDRs returns an object that can list and get IpamPreviewNextCIDRs.
func (s *ipamPreviewNextCIDRLister) IpamPreviewNextCIDRs(namespace string) IpamPreviewNextCIDRNamespaceLister {
	return ipamPreviewNextCIDRNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IpamPreviewNextCIDRNamespaceLister helps list and get IpamPreviewNextCIDRs.
// All objects returned here must be treated as read-only.
type IpamPreviewNextCIDRNamespaceLister interface {
	// List lists all IpamPreviewNextCIDRs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IpamPreviewNextCIDR, err error)
	// Get retrieves the IpamPreviewNextCIDR from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IpamPreviewNextCIDR, error)
	IpamPreviewNextCIDRNamespaceListerExpansion
}

// ipamPreviewNextCIDRNamespaceLister implements the IpamPreviewNextCIDRNamespaceLister
// interface.
type ipamPreviewNextCIDRNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IpamPreviewNextCIDRs in the indexer for a given namespace.
func (s ipamPreviewNextCIDRNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IpamPreviewNextCIDR, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IpamPreviewNextCIDR))
	})
	return ret, err
}

// Get retrieves the IpamPreviewNextCIDR from the indexer for a given namespace and name.
func (s ipamPreviewNextCIDRNamespaceLister) Get(name string) (*v1alpha1.IpamPreviewNextCIDR, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ipampreviewnextcidr"), name)
	}
	return obj.(*v1alpha1.IpamPreviewNextCIDR), nil
}