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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/amplify/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainAssociationLister helps list DomainAssociations.
// All objects returned here must be treated as read-only.
type DomainAssociationLister interface {
	// List lists all DomainAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainAssociation, err error)
	// DomainAssociations returns an object that can list and get DomainAssociations.
	DomainAssociations(namespace string) DomainAssociationNamespaceLister
	DomainAssociationListerExpansion
}

// domainAssociationLister implements the DomainAssociationLister interface.
type domainAssociationLister struct {
	indexer cache.Indexer
}

// NewDomainAssociationLister returns a new DomainAssociationLister.
func NewDomainAssociationLister(indexer cache.Indexer) DomainAssociationLister {
	return &domainAssociationLister{indexer: indexer}
}

// List lists all DomainAssociations in the indexer.
func (s *domainAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.DomainAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainAssociation))
	})
	return ret, err
}

// DomainAssociations returns an object that can list and get DomainAssociations.
func (s *domainAssociationLister) DomainAssociations(namespace string) DomainAssociationNamespaceLister {
	return domainAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainAssociationNamespaceLister helps list and get DomainAssociations.
// All objects returned here must be treated as read-only.
type DomainAssociationNamespaceLister interface {
	// List lists all DomainAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainAssociation, err error)
	// Get retrieves the DomainAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainAssociation, error)
	DomainAssociationNamespaceListerExpansion
}

// domainAssociationNamespaceLister implements the DomainAssociationNamespaceLister
// interface.
type domainAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainAssociations in the indexer for a given namespace.
func (s domainAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainAssociation))
	})
	return ret, err
}

// Get retrieves the DomainAssociation from the indexer for a given namespace and name.
func (s domainAssociationNamespaceLister) Get(name string) (*v1alpha1.DomainAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainassociation"), name)
	}
	return obj.(*v1alpha1.DomainAssociation), nil
}
