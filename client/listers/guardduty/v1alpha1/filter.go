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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/guardduty/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FilterLister helps list Filters.
// All objects returned here must be treated as read-only.
type FilterLister interface {
	// List lists all Filters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Filter, err error)
	// Filters returns an object that can list and get Filters.
	Filters(namespace string) FilterNamespaceLister
	FilterListerExpansion
}

// filterLister implements the FilterLister interface.
type filterLister struct {
	indexer cache.Indexer
}

// NewFilterLister returns a new FilterLister.
func NewFilterLister(indexer cache.Indexer) FilterLister {
	return &filterLister{indexer: indexer}
}

// List lists all Filters in the indexer.
func (s *filterLister) List(selector labels.Selector) (ret []*v1alpha1.Filter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Filter))
	})
	return ret, err
}

// Filters returns an object that can list and get Filters.
func (s *filterLister) Filters(namespace string) FilterNamespaceLister {
	return filterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FilterNamespaceLister helps list and get Filters.
// All objects returned here must be treated as read-only.
type FilterNamespaceLister interface {
	// List lists all Filters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Filter, err error)
	// Get retrieves the Filter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Filter, error)
	FilterNamespaceListerExpansion
}

// filterNamespaceLister implements the FilterNamespaceLister
// interface.
type filterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Filters in the indexer for a given namespace.
func (s filterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Filter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Filter))
	})
	return ret, err
}

// Get retrieves the Filter from the indexer for a given namespace and name.
func (s filterNamespaceLister) Get(name string) (*v1alpha1.Filter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("filter"), name)
	}
	return obj.(*v1alpha1.Filter), nil
}
