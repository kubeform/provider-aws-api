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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/db/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstanceRoleAssociationLister helps list InstanceRoleAssociations.
// All objects returned here must be treated as read-only.
type InstanceRoleAssociationLister interface {
	// List lists all InstanceRoleAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceRoleAssociation, err error)
	// InstanceRoleAssociations returns an object that can list and get InstanceRoleAssociations.
	InstanceRoleAssociations(namespace string) InstanceRoleAssociationNamespaceLister
	InstanceRoleAssociationListerExpansion
}

// instanceRoleAssociationLister implements the InstanceRoleAssociationLister interface.
type instanceRoleAssociationLister struct {
	indexer cache.Indexer
}

// NewInstanceRoleAssociationLister returns a new InstanceRoleAssociationLister.
func NewInstanceRoleAssociationLister(indexer cache.Indexer) InstanceRoleAssociationLister {
	return &instanceRoleAssociationLister{indexer: indexer}
}

// List lists all InstanceRoleAssociations in the indexer.
func (s *instanceRoleAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceRoleAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceRoleAssociation))
	})
	return ret, err
}

// InstanceRoleAssociations returns an object that can list and get InstanceRoleAssociations.
func (s *instanceRoleAssociationLister) InstanceRoleAssociations(namespace string) InstanceRoleAssociationNamespaceLister {
	return instanceRoleAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceRoleAssociationNamespaceLister helps list and get InstanceRoleAssociations.
// All objects returned here must be treated as read-only.
type InstanceRoleAssociationNamespaceLister interface {
	// List lists all InstanceRoleAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceRoleAssociation, err error)
	// Get retrieves the InstanceRoleAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstanceRoleAssociation, error)
	InstanceRoleAssociationNamespaceListerExpansion
}

// instanceRoleAssociationNamespaceLister implements the InstanceRoleAssociationNamespaceLister
// interface.
type instanceRoleAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstanceRoleAssociations in the indexer for a given namespace.
func (s instanceRoleAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceRoleAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceRoleAssociation))
	})
	return ret, err
}

// Get retrieves the InstanceRoleAssociation from the indexer for a given namespace and name.
func (s instanceRoleAssociationNamespaceLister) Get(name string) (*v1alpha1.InstanceRoleAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instanceroleassociation"), name)
	}
	return obj.(*v1alpha1.InstanceRoleAssociation), nil
}
