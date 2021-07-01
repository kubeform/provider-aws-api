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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/servicecatalog/v1alpha1"
)

// PrincipalPortfolioAssociationLister helps list PrincipalPortfolioAssociations.
// All objects returned here must be treated as read-only.
type PrincipalPortfolioAssociationLister interface {
	// List lists all PrincipalPortfolioAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PrincipalPortfolioAssociation, err error)
	// PrincipalPortfolioAssociations returns an object that can list and get PrincipalPortfolioAssociations.
	PrincipalPortfolioAssociations(namespace string) PrincipalPortfolioAssociationNamespaceLister
	PrincipalPortfolioAssociationListerExpansion
}

// principalPortfolioAssociationLister implements the PrincipalPortfolioAssociationLister interface.
type principalPortfolioAssociationLister struct {
	indexer cache.Indexer
}

// NewPrincipalPortfolioAssociationLister returns a new PrincipalPortfolioAssociationLister.
func NewPrincipalPortfolioAssociationLister(indexer cache.Indexer) PrincipalPortfolioAssociationLister {
	return &principalPortfolioAssociationLister{indexer: indexer}
}

// List lists all PrincipalPortfolioAssociations in the indexer.
func (s *principalPortfolioAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.PrincipalPortfolioAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrincipalPortfolioAssociation))
	})
	return ret, err
}

// PrincipalPortfolioAssociations returns an object that can list and get PrincipalPortfolioAssociations.
func (s *principalPortfolioAssociationLister) PrincipalPortfolioAssociations(namespace string) PrincipalPortfolioAssociationNamespaceLister {
	return principalPortfolioAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrincipalPortfolioAssociationNamespaceLister helps list and get PrincipalPortfolioAssociations.
// All objects returned here must be treated as read-only.
type PrincipalPortfolioAssociationNamespaceLister interface {
	// List lists all PrincipalPortfolioAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PrincipalPortfolioAssociation, err error)
	// Get retrieves the PrincipalPortfolioAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PrincipalPortfolioAssociation, error)
	PrincipalPortfolioAssociationNamespaceListerExpansion
}

// principalPortfolioAssociationNamespaceLister implements the PrincipalPortfolioAssociationNamespaceLister
// interface.
type principalPortfolioAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrincipalPortfolioAssociations in the indexer for a given namespace.
func (s principalPortfolioAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrincipalPortfolioAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrincipalPortfolioAssociation))
	})
	return ret, err
}

// Get retrieves the PrincipalPortfolioAssociation from the indexer for a given namespace and name.
func (s principalPortfolioAssociationNamespaceLister) Get(name string) (*v1alpha1.PrincipalPortfolioAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("principalportfolioassociation"), name)
	}
	return obj.(*v1alpha1.PrincipalPortfolioAssociation), nil
}
