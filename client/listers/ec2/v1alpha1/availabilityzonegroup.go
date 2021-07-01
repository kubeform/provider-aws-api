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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"
)

// AvailabilityZoneGroupLister helps list AvailabilityZoneGroups.
// All objects returned here must be treated as read-only.
type AvailabilityZoneGroupLister interface {
	// List lists all AvailabilityZoneGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AvailabilityZoneGroup, err error)
	// AvailabilityZoneGroups returns an object that can list and get AvailabilityZoneGroups.
	AvailabilityZoneGroups(namespace string) AvailabilityZoneGroupNamespaceLister
	AvailabilityZoneGroupListerExpansion
}

// availabilityZoneGroupLister implements the AvailabilityZoneGroupLister interface.
type availabilityZoneGroupLister struct {
	indexer cache.Indexer
}

// NewAvailabilityZoneGroupLister returns a new AvailabilityZoneGroupLister.
func NewAvailabilityZoneGroupLister(indexer cache.Indexer) AvailabilityZoneGroupLister {
	return &availabilityZoneGroupLister{indexer: indexer}
}

// List lists all AvailabilityZoneGroups in the indexer.
func (s *availabilityZoneGroupLister) List(selector labels.Selector) (ret []*v1alpha1.AvailabilityZoneGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AvailabilityZoneGroup))
	})
	return ret, err
}

// AvailabilityZoneGroups returns an object that can list and get AvailabilityZoneGroups.
func (s *availabilityZoneGroupLister) AvailabilityZoneGroups(namespace string) AvailabilityZoneGroupNamespaceLister {
	return availabilityZoneGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AvailabilityZoneGroupNamespaceLister helps list and get AvailabilityZoneGroups.
// All objects returned here must be treated as read-only.
type AvailabilityZoneGroupNamespaceLister interface {
	// List lists all AvailabilityZoneGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AvailabilityZoneGroup, err error)
	// Get retrieves the AvailabilityZoneGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AvailabilityZoneGroup, error)
	AvailabilityZoneGroupNamespaceListerExpansion
}

// availabilityZoneGroupNamespaceLister implements the AvailabilityZoneGroupNamespaceLister
// interface.
type availabilityZoneGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AvailabilityZoneGroups in the indexer for a given namespace.
func (s availabilityZoneGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AvailabilityZoneGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AvailabilityZoneGroup))
	})
	return ret, err
}

// Get retrieves the AvailabilityZoneGroup from the indexer for a given namespace and name.
func (s availabilityZoneGroupNamespaceLister) Get(name string) (*v1alpha1.AvailabilityZoneGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("availabilityzonegroup"), name)
	}
	return obj.(*v1alpha1.AvailabilityZoneGroup), nil
}
