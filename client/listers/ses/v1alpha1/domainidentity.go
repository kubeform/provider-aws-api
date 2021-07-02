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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainIdentityLister helps list DomainIdentities.
// All objects returned here must be treated as read-only.
type DomainIdentityLister interface {
	// List lists all DomainIdentities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainIdentity, err error)
	// DomainIdentities returns an object that can list and get DomainIdentities.
	DomainIdentities(namespace string) DomainIdentityNamespaceLister
	DomainIdentityListerExpansion
}

// domainIdentityLister implements the DomainIdentityLister interface.
type domainIdentityLister struct {
	indexer cache.Indexer
}

// NewDomainIdentityLister returns a new DomainIdentityLister.
func NewDomainIdentityLister(indexer cache.Indexer) DomainIdentityLister {
	return &domainIdentityLister{indexer: indexer}
}

// List lists all DomainIdentities in the indexer.
func (s *domainIdentityLister) List(selector labels.Selector) (ret []*v1alpha1.DomainIdentity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainIdentity))
	})
	return ret, err
}

// DomainIdentities returns an object that can list and get DomainIdentities.
func (s *domainIdentityLister) DomainIdentities(namespace string) DomainIdentityNamespaceLister {
	return domainIdentityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainIdentityNamespaceLister helps list and get DomainIdentities.
// All objects returned here must be treated as read-only.
type DomainIdentityNamespaceLister interface {
	// List lists all DomainIdentities in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainIdentity, err error)
	// Get retrieves the DomainIdentity from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainIdentity, error)
	DomainIdentityNamespaceListerExpansion
}

// domainIdentityNamespaceLister implements the DomainIdentityNamespaceLister
// interface.
type domainIdentityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainIdentities in the indexer for a given namespace.
func (s domainIdentityNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainIdentity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainIdentity))
	})
	return ret, err
}

// Get retrieves the DomainIdentity from the indexer for a given namespace and name.
func (s domainIdentityNamespaceLister) Get(name string) (*v1alpha1.DomainIdentity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainidentity"), name)
	}
	return obj.(*v1alpha1.DomainIdentity), nil
}
