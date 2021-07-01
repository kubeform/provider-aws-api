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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/glue/v1alpha1"
)

// CatalogDatabaseLister helps list CatalogDatabases.
// All objects returned here must be treated as read-only.
type CatalogDatabaseLister interface {
	// List lists all CatalogDatabases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CatalogDatabase, err error)
	// CatalogDatabases returns an object that can list and get CatalogDatabases.
	CatalogDatabases(namespace string) CatalogDatabaseNamespaceLister
	CatalogDatabaseListerExpansion
}

// catalogDatabaseLister implements the CatalogDatabaseLister interface.
type catalogDatabaseLister struct {
	indexer cache.Indexer
}

// NewCatalogDatabaseLister returns a new CatalogDatabaseLister.
func NewCatalogDatabaseLister(indexer cache.Indexer) CatalogDatabaseLister {
	return &catalogDatabaseLister{indexer: indexer}
}

// List lists all CatalogDatabases in the indexer.
func (s *catalogDatabaseLister) List(selector labels.Selector) (ret []*v1alpha1.CatalogDatabase, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CatalogDatabase))
	})
	return ret, err
}

// CatalogDatabases returns an object that can list and get CatalogDatabases.
func (s *catalogDatabaseLister) CatalogDatabases(namespace string) CatalogDatabaseNamespaceLister {
	return catalogDatabaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CatalogDatabaseNamespaceLister helps list and get CatalogDatabases.
// All objects returned here must be treated as read-only.
type CatalogDatabaseNamespaceLister interface {
	// List lists all CatalogDatabases in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CatalogDatabase, err error)
	// Get retrieves the CatalogDatabase from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CatalogDatabase, error)
	CatalogDatabaseNamespaceListerExpansion
}

// catalogDatabaseNamespaceLister implements the CatalogDatabaseNamespaceLister
// interface.
type catalogDatabaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CatalogDatabases in the indexer for a given namespace.
func (s catalogDatabaseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CatalogDatabase, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CatalogDatabase))
	})
	return ret, err
}

// Get retrieves the CatalogDatabase from the indexer for a given namespace and name.
func (s catalogDatabaseNamespaceLister) Get(name string) (*v1alpha1.CatalogDatabase, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("catalogdatabase"), name)
	}
	return obj.(*v1alpha1.CatalogDatabase), nil
}
