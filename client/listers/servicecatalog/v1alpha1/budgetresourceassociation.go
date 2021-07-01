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

// BudgetResourceAssociationLister helps list BudgetResourceAssociations.
// All objects returned here must be treated as read-only.
type BudgetResourceAssociationLister interface {
	// List lists all BudgetResourceAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BudgetResourceAssociation, err error)
	// BudgetResourceAssociations returns an object that can list and get BudgetResourceAssociations.
	BudgetResourceAssociations(namespace string) BudgetResourceAssociationNamespaceLister
	BudgetResourceAssociationListerExpansion
}

// budgetResourceAssociationLister implements the BudgetResourceAssociationLister interface.
type budgetResourceAssociationLister struct {
	indexer cache.Indexer
}

// NewBudgetResourceAssociationLister returns a new BudgetResourceAssociationLister.
func NewBudgetResourceAssociationLister(indexer cache.Indexer) BudgetResourceAssociationLister {
	return &budgetResourceAssociationLister{indexer: indexer}
}

// List lists all BudgetResourceAssociations in the indexer.
func (s *budgetResourceAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.BudgetResourceAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BudgetResourceAssociation))
	})
	return ret, err
}

// BudgetResourceAssociations returns an object that can list and get BudgetResourceAssociations.
func (s *budgetResourceAssociationLister) BudgetResourceAssociations(namespace string) BudgetResourceAssociationNamespaceLister {
	return budgetResourceAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BudgetResourceAssociationNamespaceLister helps list and get BudgetResourceAssociations.
// All objects returned here must be treated as read-only.
type BudgetResourceAssociationNamespaceLister interface {
	// List lists all BudgetResourceAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BudgetResourceAssociation, err error)
	// Get retrieves the BudgetResourceAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BudgetResourceAssociation, error)
	BudgetResourceAssociationNamespaceListerExpansion
}

// budgetResourceAssociationNamespaceLister implements the BudgetResourceAssociationNamespaceLister
// interface.
type budgetResourceAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BudgetResourceAssociations in the indexer for a given namespace.
func (s budgetResourceAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BudgetResourceAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BudgetResourceAssociation))
	})
	return ret, err
}

// Get retrieves the BudgetResourceAssociation from the indexer for a given namespace and name.
func (s budgetResourceAssociationNamespaceLister) Get(name string) (*v1alpha1.BudgetResourceAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("budgetresourceassociation"), name)
	}
	return obj.(*v1alpha1.BudgetResourceAssociation), nil
}
