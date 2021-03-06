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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/wafv2/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IpSetLister helps list IpSets.
// All objects returned here must be treated as read-only.
type IpSetLister interface {
	// List lists all IpSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IpSet, err error)
	// IpSets returns an object that can list and get IpSets.
	IpSets(namespace string) IpSetNamespaceLister
	IpSetListerExpansion
}

// ipSetLister implements the IpSetLister interface.
type ipSetLister struct {
	indexer cache.Indexer
}

// NewIpSetLister returns a new IpSetLister.
func NewIpSetLister(indexer cache.Indexer) IpSetLister {
	return &ipSetLister{indexer: indexer}
}

// List lists all IpSets in the indexer.
func (s *ipSetLister) List(selector labels.Selector) (ret []*v1alpha1.IpSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IpSet))
	})
	return ret, err
}

// IpSets returns an object that can list and get IpSets.
func (s *ipSetLister) IpSets(namespace string) IpSetNamespaceLister {
	return ipSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IpSetNamespaceLister helps list and get IpSets.
// All objects returned here must be treated as read-only.
type IpSetNamespaceLister interface {
	// List lists all IpSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IpSet, err error)
	// Get retrieves the IpSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IpSet, error)
	IpSetNamespaceListerExpansion
}

// ipSetNamespaceLister implements the IpSetNamespaceLister
// interface.
type ipSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IpSets in the indexer for a given namespace.
func (s ipSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IpSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IpSet))
	})
	return ret, err
}

// Get retrieves the IpSet from the indexer for a given namespace and name.
func (s ipSetNamespaceLister) Get(name string) (*v1alpha1.IpSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ipset"), name)
	}
	return obj.(*v1alpha1.IpSet), nil
}
