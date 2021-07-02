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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/eip/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AssociationLister helps list Associations.
// All objects returned here must be treated as read-only.
type AssociationLister interface {
	// List lists all Associations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Association, err error)
	// Associations returns an object that can list and get Associations.
	Associations(namespace string) AssociationNamespaceLister
	AssociationListerExpansion
}

// associationLister implements the AssociationLister interface.
type associationLister struct {
	indexer cache.Indexer
}

// NewAssociationLister returns a new AssociationLister.
func NewAssociationLister(indexer cache.Indexer) AssociationLister {
	return &associationLister{indexer: indexer}
}

// List lists all Associations in the indexer.
func (s *associationLister) List(selector labels.Selector) (ret []*v1alpha1.Association, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Association))
	})
	return ret, err
}

// Associations returns an object that can list and get Associations.
func (s *associationLister) Associations(namespace string) AssociationNamespaceLister {
	return associationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AssociationNamespaceLister helps list and get Associations.
// All objects returned here must be treated as read-only.
type AssociationNamespaceLister interface {
	// List lists all Associations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Association, err error)
	// Get retrieves the Association from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Association, error)
	AssociationNamespaceListerExpansion
}

// associationNamespaceLister implements the AssociationNamespaceLister
// interface.
type associationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Associations in the indexer for a given namespace.
func (s associationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Association, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Association))
	})
	return ret, err
}

// Get retrieves the Association from the indexer for a given namespace and name.
func (s associationNamespaceLister) Get(name string) (*v1alpha1.Association, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("association"), name)
	}
	return obj.(*v1alpha1.Association), nil
}
