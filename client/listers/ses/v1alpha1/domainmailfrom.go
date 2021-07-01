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

// DomainMailFromLister helps list DomainMailFroms.
// All objects returned here must be treated as read-only.
type DomainMailFromLister interface {
	// List lists all DomainMailFroms in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainMailFrom, err error)
	// DomainMailFroms returns an object that can list and get DomainMailFroms.
	DomainMailFroms(namespace string) DomainMailFromNamespaceLister
	DomainMailFromListerExpansion
}

// domainMailFromLister implements the DomainMailFromLister interface.
type domainMailFromLister struct {
	indexer cache.Indexer
}

// NewDomainMailFromLister returns a new DomainMailFromLister.
func NewDomainMailFromLister(indexer cache.Indexer) DomainMailFromLister {
	return &domainMailFromLister{indexer: indexer}
}

// List lists all DomainMailFroms in the indexer.
func (s *domainMailFromLister) List(selector labels.Selector) (ret []*v1alpha1.DomainMailFrom, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainMailFrom))
	})
	return ret, err
}

// DomainMailFroms returns an object that can list and get DomainMailFroms.
func (s *domainMailFromLister) DomainMailFroms(namespace string) DomainMailFromNamespaceLister {
	return domainMailFromNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainMailFromNamespaceLister helps list and get DomainMailFroms.
// All objects returned here must be treated as read-only.
type DomainMailFromNamespaceLister interface {
	// List lists all DomainMailFroms in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainMailFrom, err error)
	// Get retrieves the DomainMailFrom from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainMailFrom, error)
	DomainMailFromNamespaceListerExpansion
}

// domainMailFromNamespaceLister implements the DomainMailFromNamespaceLister
// interface.
type domainMailFromNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainMailFroms in the indexer for a given namespace.
func (s domainMailFromNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainMailFrom, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainMailFrom))
	})
	return ret, err
}

// Get retrieves the DomainMailFrom from the indexer for a given namespace and name.
func (s domainMailFromNamespaceLister) Get(name string) (*v1alpha1.DomainMailFrom, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainmailfrom"), name)
	}
	return obj.(*v1alpha1.DomainMailFrom), nil
}
