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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/storagegateway/v1alpha1"
)

// CacheLister helps list Caches.
// All objects returned here must be treated as read-only.
type CacheLister interface {
	// List lists all Caches in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Cache, err error)
	// Caches returns an object that can list and get Caches.
	Caches(namespace string) CacheNamespaceLister
	CacheListerExpansion
}

// cacheLister implements the CacheLister interface.
type cacheLister struct {
	indexer cache.Indexer
}

// NewCacheLister returns a new CacheLister.
func NewCacheLister(indexer cache.Indexer) CacheLister {
	return &cacheLister{indexer: indexer}
}

// List lists all Caches in the indexer.
func (s *cacheLister) List(selector labels.Selector) (ret []*v1alpha1.Cache, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Cache))
	})
	return ret, err
}

// Caches returns an object that can list and get Caches.
func (s *cacheLister) Caches(namespace string) CacheNamespaceLister {
	return cacheNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CacheNamespaceLister helps list and get Caches.
// All objects returned here must be treated as read-only.
type CacheNamespaceLister interface {
	// List lists all Caches in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Cache, err error)
	// Get retrieves the Cache from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Cache, error)
	CacheNamespaceListerExpansion
}

// cacheNamespaceLister implements the CacheNamespaceLister
// interface.
type cacheNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Caches in the indexer for a given namespace.
func (s cacheNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Cache, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Cache))
	})
	return ret, err
}

// Get retrieves the Cache from the indexer for a given namespace and name.
func (s cacheNamespaceLister) Get(name string) (*v1alpha1.Cache, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cache"), name)
	}
	return obj.(*v1alpha1.Cache), nil
}
