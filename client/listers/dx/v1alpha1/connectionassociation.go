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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dx/v1alpha1"
)

// ConnectionAssociationLister helps list ConnectionAssociations.
// All objects returned here must be treated as read-only.
type ConnectionAssociationLister interface {
	// List lists all ConnectionAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConnectionAssociation, err error)
	// ConnectionAssociations returns an object that can list and get ConnectionAssociations.
	ConnectionAssociations(namespace string) ConnectionAssociationNamespaceLister
	ConnectionAssociationListerExpansion
}

// connectionAssociationLister implements the ConnectionAssociationLister interface.
type connectionAssociationLister struct {
	indexer cache.Indexer
}

// NewConnectionAssociationLister returns a new ConnectionAssociationLister.
func NewConnectionAssociationLister(indexer cache.Indexer) ConnectionAssociationLister {
	return &connectionAssociationLister{indexer: indexer}
}

// List lists all ConnectionAssociations in the indexer.
func (s *connectionAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.ConnectionAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConnectionAssociation))
	})
	return ret, err
}

// ConnectionAssociations returns an object that can list and get ConnectionAssociations.
func (s *connectionAssociationLister) ConnectionAssociations(namespace string) ConnectionAssociationNamespaceLister {
	return connectionAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConnectionAssociationNamespaceLister helps list and get ConnectionAssociations.
// All objects returned here must be treated as read-only.
type ConnectionAssociationNamespaceLister interface {
	// List lists all ConnectionAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConnectionAssociation, err error)
	// Get retrieves the ConnectionAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConnectionAssociation, error)
	ConnectionAssociationNamespaceListerExpansion
}

// connectionAssociationNamespaceLister implements the ConnectionAssociationNamespaceLister
// interface.
type connectionAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConnectionAssociations in the indexer for a given namespace.
func (s connectionAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConnectionAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConnectionAssociation))
	})
	return ret, err
}

// Get retrieves the ConnectionAssociation from the indexer for a given namespace and name.
func (s connectionAssociationNamespaceLister) Get(name string) (*v1alpha1.ConnectionAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("connectionassociation"), name)
	}
	return obj.(*v1alpha1.ConnectionAssociation), nil
}
