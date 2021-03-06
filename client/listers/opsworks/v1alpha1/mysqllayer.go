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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/opsworks/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MysqlLayerLister helps list MysqlLayers.
// All objects returned here must be treated as read-only.
type MysqlLayerLister interface {
	// List lists all MysqlLayers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MysqlLayer, err error)
	// MysqlLayers returns an object that can list and get MysqlLayers.
	MysqlLayers(namespace string) MysqlLayerNamespaceLister
	MysqlLayerListerExpansion
}

// mysqlLayerLister implements the MysqlLayerLister interface.
type mysqlLayerLister struct {
	indexer cache.Indexer
}

// NewMysqlLayerLister returns a new MysqlLayerLister.
func NewMysqlLayerLister(indexer cache.Indexer) MysqlLayerLister {
	return &mysqlLayerLister{indexer: indexer}
}

// List lists all MysqlLayers in the indexer.
func (s *mysqlLayerLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlLayer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlLayer))
	})
	return ret, err
}

// MysqlLayers returns an object that can list and get MysqlLayers.
func (s *mysqlLayerLister) MysqlLayers(namespace string) MysqlLayerNamespaceLister {
	return mysqlLayerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MysqlLayerNamespaceLister helps list and get MysqlLayers.
// All objects returned here must be treated as read-only.
type MysqlLayerNamespaceLister interface {
	// List lists all MysqlLayers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MysqlLayer, err error)
	// Get retrieves the MysqlLayer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MysqlLayer, error)
	MysqlLayerNamespaceListerExpansion
}

// mysqlLayerNamespaceLister implements the MysqlLayerNamespaceLister
// interface.
type mysqlLayerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MysqlLayers in the indexer for a given namespace.
func (s mysqlLayerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MysqlLayer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MysqlLayer))
	})
	return ret, err
}

// Get retrieves the MysqlLayer from the indexer for a given namespace and name.
func (s mysqlLayerNamespaceLister) Get(name string) (*v1alpha1.MysqlLayer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mysqllayer"), name)
	}
	return obj.(*v1alpha1.MysqlLayer), nil
}
