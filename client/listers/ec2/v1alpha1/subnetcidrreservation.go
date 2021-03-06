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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SubnetCIDRReservationLister helps list SubnetCIDRReservations.
// All objects returned here must be treated as read-only.
type SubnetCIDRReservationLister interface {
	// List lists all SubnetCIDRReservations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SubnetCIDRReservation, err error)
	// SubnetCIDRReservations returns an object that can list and get SubnetCIDRReservations.
	SubnetCIDRReservations(namespace string) SubnetCIDRReservationNamespaceLister
	SubnetCIDRReservationListerExpansion
}

// subnetCIDRReservationLister implements the SubnetCIDRReservationLister interface.
type subnetCIDRReservationLister struct {
	indexer cache.Indexer
}

// NewSubnetCIDRReservationLister returns a new SubnetCIDRReservationLister.
func NewSubnetCIDRReservationLister(indexer cache.Indexer) SubnetCIDRReservationLister {
	return &subnetCIDRReservationLister{indexer: indexer}
}

// List lists all SubnetCIDRReservations in the indexer.
func (s *subnetCIDRReservationLister) List(selector labels.Selector) (ret []*v1alpha1.SubnetCIDRReservation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubnetCIDRReservation))
	})
	return ret, err
}

// SubnetCIDRReservations returns an object that can list and get SubnetCIDRReservations.
func (s *subnetCIDRReservationLister) SubnetCIDRReservations(namespace string) SubnetCIDRReservationNamespaceLister {
	return subnetCIDRReservationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubnetCIDRReservationNamespaceLister helps list and get SubnetCIDRReservations.
// All objects returned here must be treated as read-only.
type SubnetCIDRReservationNamespaceLister interface {
	// List lists all SubnetCIDRReservations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SubnetCIDRReservation, err error)
	// Get retrieves the SubnetCIDRReservation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SubnetCIDRReservation, error)
	SubnetCIDRReservationNamespaceListerExpansion
}

// subnetCIDRReservationNamespaceLister implements the SubnetCIDRReservationNamespaceLister
// interface.
type subnetCIDRReservationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SubnetCIDRReservations in the indexer for a given namespace.
func (s subnetCIDRReservationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SubnetCIDRReservation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubnetCIDRReservation))
	})
	return ret, err
}

// Get retrieves the SubnetCIDRReservation from the indexer for a given namespace and name.
func (s subnetCIDRReservationNamespaceLister) Get(name string) (*v1alpha1.SubnetCIDRReservation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("subnetcidrreservation"), name)
	}
	return obj.(*v1alpha1.SubnetCIDRReservation), nil
}
