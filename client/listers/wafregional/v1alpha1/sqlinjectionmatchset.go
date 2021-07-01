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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/wafregional/v1alpha1"
)

// SqlInjectionMatchSetLister helps list SqlInjectionMatchSets.
// All objects returned here must be treated as read-only.
type SqlInjectionMatchSetLister interface {
	// List lists all SqlInjectionMatchSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SqlInjectionMatchSet, err error)
	// SqlInjectionMatchSets returns an object that can list and get SqlInjectionMatchSets.
	SqlInjectionMatchSets(namespace string) SqlInjectionMatchSetNamespaceLister
	SqlInjectionMatchSetListerExpansion
}

// sqlInjectionMatchSetLister implements the SqlInjectionMatchSetLister interface.
type sqlInjectionMatchSetLister struct {
	indexer cache.Indexer
}

// NewSqlInjectionMatchSetLister returns a new SqlInjectionMatchSetLister.
func NewSqlInjectionMatchSetLister(indexer cache.Indexer) SqlInjectionMatchSetLister {
	return &sqlInjectionMatchSetLister{indexer: indexer}
}

// List lists all SqlInjectionMatchSets in the indexer.
func (s *sqlInjectionMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.SqlInjectionMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqlInjectionMatchSet))
	})
	return ret, err
}

// SqlInjectionMatchSets returns an object that can list and get SqlInjectionMatchSets.
func (s *sqlInjectionMatchSetLister) SqlInjectionMatchSets(namespace string) SqlInjectionMatchSetNamespaceLister {
	return sqlInjectionMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SqlInjectionMatchSetNamespaceLister helps list and get SqlInjectionMatchSets.
// All objects returned here must be treated as read-only.
type SqlInjectionMatchSetNamespaceLister interface {
	// List lists all SqlInjectionMatchSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SqlInjectionMatchSet, err error)
	// Get retrieves the SqlInjectionMatchSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SqlInjectionMatchSet, error)
	SqlInjectionMatchSetNamespaceListerExpansion
}

// sqlInjectionMatchSetNamespaceLister implements the SqlInjectionMatchSetNamespaceLister
// interface.
type sqlInjectionMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SqlInjectionMatchSets in the indexer for a given namespace.
func (s sqlInjectionMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SqlInjectionMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqlInjectionMatchSet))
	})
	return ret, err
}

// Get retrieves the SqlInjectionMatchSet from the indexer for a given namespace and name.
func (s sqlInjectionMatchSetNamespaceLister) Get(name string) (*v1alpha1.SqlInjectionMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sqlinjectionmatchset"), name)
	}
	return obj.(*v1alpha1.SqlInjectionMatchSet), nil
}
