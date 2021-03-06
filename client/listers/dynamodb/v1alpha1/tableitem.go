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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dynamodb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TableItemLister helps list TableItems.
// All objects returned here must be treated as read-only.
type TableItemLister interface {
	// List lists all TableItems in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TableItem, err error)
	// TableItems returns an object that can list and get TableItems.
	TableItems(namespace string) TableItemNamespaceLister
	TableItemListerExpansion
}

// tableItemLister implements the TableItemLister interface.
type tableItemLister struct {
	indexer cache.Indexer
}

// NewTableItemLister returns a new TableItemLister.
func NewTableItemLister(indexer cache.Indexer) TableItemLister {
	return &tableItemLister{indexer: indexer}
}

// List lists all TableItems in the indexer.
func (s *tableItemLister) List(selector labels.Selector) (ret []*v1alpha1.TableItem, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TableItem))
	})
	return ret, err
}

// TableItems returns an object that can list and get TableItems.
func (s *tableItemLister) TableItems(namespace string) TableItemNamespaceLister {
	return tableItemNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TableItemNamespaceLister helps list and get TableItems.
// All objects returned here must be treated as read-only.
type TableItemNamespaceLister interface {
	// List lists all TableItems in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TableItem, err error)
	// Get retrieves the TableItem from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TableItem, error)
	TableItemNamespaceListerExpansion
}

// tableItemNamespaceLister implements the TableItemNamespaceLister
// interface.
type tableItemNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TableItems in the indexer for a given namespace.
func (s tableItemNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TableItem, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TableItem))
	})
	return ret, err
}

// Get retrieves the TableItem from the indexer for a given namespace and name.
func (s tableItemNamespaceLister) Get(name string) (*v1alpha1.TableItem, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tableitem"), name)
	}
	return obj.(*v1alpha1.TableItem), nil
}
