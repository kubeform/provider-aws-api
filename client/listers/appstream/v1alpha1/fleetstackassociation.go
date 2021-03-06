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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/appstream/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FleetStackAssociationLister helps list FleetStackAssociations.
// All objects returned here must be treated as read-only.
type FleetStackAssociationLister interface {
	// List lists all FleetStackAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FleetStackAssociation, err error)
	// FleetStackAssociations returns an object that can list and get FleetStackAssociations.
	FleetStackAssociations(namespace string) FleetStackAssociationNamespaceLister
	FleetStackAssociationListerExpansion
}

// fleetStackAssociationLister implements the FleetStackAssociationLister interface.
type fleetStackAssociationLister struct {
	indexer cache.Indexer
}

// NewFleetStackAssociationLister returns a new FleetStackAssociationLister.
func NewFleetStackAssociationLister(indexer cache.Indexer) FleetStackAssociationLister {
	return &fleetStackAssociationLister{indexer: indexer}
}

// List lists all FleetStackAssociations in the indexer.
func (s *fleetStackAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.FleetStackAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FleetStackAssociation))
	})
	return ret, err
}

// FleetStackAssociations returns an object that can list and get FleetStackAssociations.
func (s *fleetStackAssociationLister) FleetStackAssociations(namespace string) FleetStackAssociationNamespaceLister {
	return fleetStackAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FleetStackAssociationNamespaceLister helps list and get FleetStackAssociations.
// All objects returned here must be treated as read-only.
type FleetStackAssociationNamespaceLister interface {
	// List lists all FleetStackAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FleetStackAssociation, err error)
	// Get retrieves the FleetStackAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FleetStackAssociation, error)
	FleetStackAssociationNamespaceListerExpansion
}

// fleetStackAssociationNamespaceLister implements the FleetStackAssociationNamespaceLister
// interface.
type fleetStackAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FleetStackAssociations in the indexer for a given namespace.
func (s fleetStackAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FleetStackAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FleetStackAssociation))
	})
	return ret, err
}

// Get retrieves the FleetStackAssociation from the indexer for a given namespace and name.
func (s fleetStackAssociationNamespaceLister) Get(name string) (*v1alpha1.FleetStackAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fleetstackassociation"), name)
	}
	return obj.(*v1alpha1.FleetStackAssociation), nil
}
