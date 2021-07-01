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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/simpledb/v1alpha1"
)

// DomainLister helps list Domains.
// All objects returned here must be treated as read-only.
type DomainLister interface {
	// List lists all Domains in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Domain, err error)
	// Domains returns an object that can list and get Domains.
	Domains(namespace string) DomainNamespaceLister
	DomainListerExpansion
}

// domainLister implements the DomainLister interface.
type domainLister struct {
	indexer cache.Indexer
}

// NewDomainLister returns a new DomainLister.
func NewDomainLister(indexer cache.Indexer) DomainLister {
	return &domainLister{indexer: indexer}
}

// List lists all Domains in the indexer.
func (s *domainLister) List(selector labels.Selector) (ret []*v1alpha1.Domain, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Domain))
	})
	return ret, err
}

// Domains returns an object that can list and get Domains.
func (s *domainLister) Domains(namespace string) DomainNamespaceLister {
	return domainNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainNamespaceLister helps list and get Domains.
// All objects returned here must be treated as read-only.
type DomainNamespaceLister interface {
	// List lists all Domains in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Domain, err error)
	// Get retrieves the Domain from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Domain, error)
	DomainNamespaceListerExpansion
}

// domainNamespaceLister implements the DomainNamespaceLister
// interface.
type domainNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Domains in the indexer for a given namespace.
func (s domainNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Domain, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Domain))
	})
	return ret, err
}

// Get retrieves the Domain from the indexer for a given namespace and name.
func (s domainNamespaceLister) Get(name string) (*v1alpha1.Domain, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domain"), name)
	}
	return obj.(*v1alpha1.Domain), nil
}
