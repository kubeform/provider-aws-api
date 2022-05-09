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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/shield/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProtectionHealthCheckAssociationLister helps list ProtectionHealthCheckAssociations.
// All objects returned here must be treated as read-only.
type ProtectionHealthCheckAssociationLister interface {
	// List lists all ProtectionHealthCheckAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionHealthCheckAssociation, err error)
	// ProtectionHealthCheckAssociations returns an object that can list and get ProtectionHealthCheckAssociations.
	ProtectionHealthCheckAssociations(namespace string) ProtectionHealthCheckAssociationNamespaceLister
	ProtectionHealthCheckAssociationListerExpansion
}

// protectionHealthCheckAssociationLister implements the ProtectionHealthCheckAssociationLister interface.
type protectionHealthCheckAssociationLister struct {
	indexer cache.Indexer
}

// NewProtectionHealthCheckAssociationLister returns a new ProtectionHealthCheckAssociationLister.
func NewProtectionHealthCheckAssociationLister(indexer cache.Indexer) ProtectionHealthCheckAssociationLister {
	return &protectionHealthCheckAssociationLister{indexer: indexer}
}

// List lists all ProtectionHealthCheckAssociations in the indexer.
func (s *protectionHealthCheckAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionHealthCheckAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionHealthCheckAssociation))
	})
	return ret, err
}

// ProtectionHealthCheckAssociations returns an object that can list and get ProtectionHealthCheckAssociations.
func (s *protectionHealthCheckAssociationLister) ProtectionHealthCheckAssociations(namespace string) ProtectionHealthCheckAssociationNamespaceLister {
	return protectionHealthCheckAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProtectionHealthCheckAssociationNamespaceLister helps list and get ProtectionHealthCheckAssociations.
// All objects returned here must be treated as read-only.
type ProtectionHealthCheckAssociationNamespaceLister interface {
	// List lists all ProtectionHealthCheckAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProtectionHealthCheckAssociation, err error)
	// Get retrieves the ProtectionHealthCheckAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProtectionHealthCheckAssociation, error)
	ProtectionHealthCheckAssociationNamespaceListerExpansion
}

// protectionHealthCheckAssociationNamespaceLister implements the ProtectionHealthCheckAssociationNamespaceLister
// interface.
type protectionHealthCheckAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProtectionHealthCheckAssociations in the indexer for a given namespace.
func (s protectionHealthCheckAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProtectionHealthCheckAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProtectionHealthCheckAssociation))
	})
	return ret, err
}

// Get retrieves the ProtectionHealthCheckAssociation from the indexer for a given namespace and name.
func (s protectionHealthCheckAssociationNamespaceLister) Get(name string) (*v1alpha1.ProtectionHealthCheckAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("protectionhealthcheckassociation"), name)
	}
	return obj.(*v1alpha1.ProtectionHealthCheckAssociation), nil
}