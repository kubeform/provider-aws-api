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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/docdb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GlobalClusterLister helps list GlobalClusters.
// All objects returned here must be treated as read-only.
type GlobalClusterLister interface {
	// List lists all GlobalClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GlobalCluster, err error)
	// GlobalClusters returns an object that can list and get GlobalClusters.
	GlobalClusters(namespace string) GlobalClusterNamespaceLister
	GlobalClusterListerExpansion
}

// globalClusterLister implements the GlobalClusterLister interface.
type globalClusterLister struct {
	indexer cache.Indexer
}

// NewGlobalClusterLister returns a new GlobalClusterLister.
func NewGlobalClusterLister(indexer cache.Indexer) GlobalClusterLister {
	return &globalClusterLister{indexer: indexer}
}

// List lists all GlobalClusters in the indexer.
func (s *globalClusterLister) List(selector labels.Selector) (ret []*v1alpha1.GlobalCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlobalCluster))
	})
	return ret, err
}

// GlobalClusters returns an object that can list and get GlobalClusters.
func (s *globalClusterLister) GlobalClusters(namespace string) GlobalClusterNamespaceLister {
	return globalClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlobalClusterNamespaceLister helps list and get GlobalClusters.
// All objects returned here must be treated as read-only.
type GlobalClusterNamespaceLister interface {
	// List lists all GlobalClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GlobalCluster, err error)
	// Get retrieves the GlobalCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GlobalCluster, error)
	GlobalClusterNamespaceListerExpansion
}

// globalClusterNamespaceLister implements the GlobalClusterNamespaceLister
// interface.
type globalClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlobalClusters in the indexer for a given namespace.
func (s globalClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GlobalCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlobalCluster))
	})
	return ret, err
}

// Get retrieves the GlobalCluster from the indexer for a given namespace and name.
func (s globalClusterNamespaceLister) Get(name string) (*v1alpha1.GlobalCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("globalcluster"), name)
	}
	return obj.(*v1alpha1.GlobalCluster), nil
}
