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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/provider-aws-api/apis/guardduty/v1alpha1"
)

// IpsetLister helps list Ipsets.
// All objects returned here must be treated as read-only.
type IpsetLister interface {
	// List lists all Ipsets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Ipset, err error)
	// Ipsets returns an object that can list and get Ipsets.
	Ipsets(namespace string) IpsetNamespaceLister
	IpsetListerExpansion
}

// ipsetLister implements the IpsetLister interface.
type ipsetLister struct {
	indexer cache.Indexer
}

// NewIpsetLister returns a new IpsetLister.
func NewIpsetLister(indexer cache.Indexer) IpsetLister {
	return &ipsetLister{indexer: indexer}
}

// List lists all Ipsets in the indexer.
func (s *ipsetLister) List(selector labels.Selector) (ret []*v1alpha1.Ipset, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ipset))
	})
	return ret, err
}

// Ipsets returns an object that can list and get Ipsets.
func (s *ipsetLister) Ipsets(namespace string) IpsetNamespaceLister {
	return ipsetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IpsetNamespaceLister helps list and get Ipsets.
// All objects returned here must be treated as read-only.
type IpsetNamespaceLister interface {
	// List lists all Ipsets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Ipset, err error)
	// Get retrieves the Ipset from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Ipset, error)
	IpsetNamespaceListerExpansion
}

// ipsetNamespaceLister implements the IpsetNamespaceLister
// interface.
type ipsetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ipsets in the indexer for a given namespace.
func (s ipsetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Ipset, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ipset))
	})
	return ret, err
}

// Get retrieves the Ipset from the indexer for a given namespace and name.
func (s ipsetNamespaceLister) Get(name string) (*v1alpha1.Ipset, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ipset"), name)
	}
	return obj.(*v1alpha1.Ipset), nil
}
