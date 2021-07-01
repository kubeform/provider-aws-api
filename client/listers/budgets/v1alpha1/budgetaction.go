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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/budgets/v1alpha1"
)

// BudgetActionLister helps list BudgetActions.
// All objects returned here must be treated as read-only.
type BudgetActionLister interface {
	// List lists all BudgetActions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BudgetAction, err error)
	// BudgetActions returns an object that can list and get BudgetActions.
	BudgetActions(namespace string) BudgetActionNamespaceLister
	BudgetActionListerExpansion
}

// budgetActionLister implements the BudgetActionLister interface.
type budgetActionLister struct {
	indexer cache.Indexer
}

// NewBudgetActionLister returns a new BudgetActionLister.
func NewBudgetActionLister(indexer cache.Indexer) BudgetActionLister {
	return &budgetActionLister{indexer: indexer}
}

// List lists all BudgetActions in the indexer.
func (s *budgetActionLister) List(selector labels.Selector) (ret []*v1alpha1.BudgetAction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BudgetAction))
	})
	return ret, err
}

// BudgetActions returns an object that can list and get BudgetActions.
func (s *budgetActionLister) BudgetActions(namespace string) BudgetActionNamespaceLister {
	return budgetActionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BudgetActionNamespaceLister helps list and get BudgetActions.
// All objects returned here must be treated as read-only.
type BudgetActionNamespaceLister interface {
	// List lists all BudgetActions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BudgetAction, err error)
	// Get retrieves the BudgetAction from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BudgetAction, error)
	BudgetActionNamespaceListerExpansion
}

// budgetActionNamespaceLister implements the BudgetActionNamespaceLister
// interface.
type budgetActionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BudgetActions in the indexer for a given namespace.
func (s budgetActionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BudgetAction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BudgetAction))
	})
	return ret, err
}

// Get retrieves the BudgetAction from the indexer for a given namespace and name.
func (s budgetActionNamespaceLister) Get(name string) (*v1alpha1.BudgetAction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("budgetaction"), name)
	}
	return obj.(*v1alpha1.BudgetAction), nil
}
