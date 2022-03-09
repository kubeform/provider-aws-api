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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/connect/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UserHierarchyStructureLister helps list UserHierarchyStructures.
// All objects returned here must be treated as read-only.
type UserHierarchyStructureLister interface {
	// List lists all UserHierarchyStructures in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.UserHierarchyStructure, err error)
	// UserHierarchyStructures returns an object that can list and get UserHierarchyStructures.
	UserHierarchyStructures(namespace string) UserHierarchyStructureNamespaceLister
	UserHierarchyStructureListerExpansion
}

// userHierarchyStructureLister implements the UserHierarchyStructureLister interface.
type userHierarchyStructureLister struct {
	indexer cache.Indexer
}

// NewUserHierarchyStructureLister returns a new UserHierarchyStructureLister.
func NewUserHierarchyStructureLister(indexer cache.Indexer) UserHierarchyStructureLister {
	return &userHierarchyStructureLister{indexer: indexer}
}

// List lists all UserHierarchyStructures in the indexer.
func (s *userHierarchyStructureLister) List(selector labels.Selector) (ret []*v1alpha1.UserHierarchyStructure, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserHierarchyStructure))
	})
	return ret, err
}

// UserHierarchyStructures returns an object that can list and get UserHierarchyStructures.
func (s *userHierarchyStructureLister) UserHierarchyStructures(namespace string) UserHierarchyStructureNamespaceLister {
	return userHierarchyStructureNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserHierarchyStructureNamespaceLister helps list and get UserHierarchyStructures.
// All objects returned here must be treated as read-only.
type UserHierarchyStructureNamespaceLister interface {
	// List lists all UserHierarchyStructures in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.UserHierarchyStructure, err error)
	// Get retrieves the UserHierarchyStructure from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.UserHierarchyStructure, error)
	UserHierarchyStructureNamespaceListerExpansion
}

// userHierarchyStructureNamespaceLister implements the UserHierarchyStructureNamespaceLister
// interface.
type userHierarchyStructureNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all UserHierarchyStructures in the indexer for a given namespace.
func (s userHierarchyStructureNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.UserHierarchyStructure, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserHierarchyStructure))
	})
	return ret, err
}

// Get retrieves the UserHierarchyStructure from the indexer for a given namespace and name.
func (s userHierarchyStructureNamespaceLister) Get(name string) (*v1alpha1.UserHierarchyStructure, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("userhierarchystructure"), name)
	}
	return obj.(*v1alpha1.UserHierarchyStructure), nil
}
