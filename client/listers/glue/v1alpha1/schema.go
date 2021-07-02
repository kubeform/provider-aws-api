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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/glue/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SchemaLister helps list Schemas.
// All objects returned here must be treated as read-only.
type SchemaLister interface {
	// List lists all Schemas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Schema, err error)
	// Schemas returns an object that can list and get Schemas.
	Schemas(namespace string) SchemaNamespaceLister
	SchemaListerExpansion
}

// schemaLister implements the SchemaLister interface.
type schemaLister struct {
	indexer cache.Indexer
}

// NewSchemaLister returns a new SchemaLister.
func NewSchemaLister(indexer cache.Indexer) SchemaLister {
	return &schemaLister{indexer: indexer}
}

// List lists all Schemas in the indexer.
func (s *schemaLister) List(selector labels.Selector) (ret []*v1alpha1.Schema, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Schema))
	})
	return ret, err
}

// Schemas returns an object that can list and get Schemas.
func (s *schemaLister) Schemas(namespace string) SchemaNamespaceLister {
	return schemaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SchemaNamespaceLister helps list and get Schemas.
// All objects returned here must be treated as read-only.
type SchemaNamespaceLister interface {
	// List lists all Schemas in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Schema, err error)
	// Get retrieves the Schema from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Schema, error)
	SchemaNamespaceListerExpansion
}

// schemaNamespaceLister implements the SchemaNamespaceLister
// interface.
type schemaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Schemas in the indexer for a given namespace.
func (s schemaNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Schema, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Schema))
	})
	return ret, err
}

// Get retrieves the Schema from the indexer for a given namespace and name.
func (s schemaNamespaceLister) Get(name string) (*v1alpha1.Schema, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("schema"), name)
	}
	return obj.(*v1alpha1.Schema), nil
}
