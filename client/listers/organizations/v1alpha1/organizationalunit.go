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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/organizations/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OrganizationalUnitLister helps list OrganizationalUnits.
// All objects returned here must be treated as read-only.
type OrganizationalUnitLister interface {
	// List lists all OrganizationalUnits in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationalUnit, err error)
	// OrganizationalUnits returns an object that can list and get OrganizationalUnits.
	OrganizationalUnits(namespace string) OrganizationalUnitNamespaceLister
	OrganizationalUnitListerExpansion
}

// organizationalUnitLister implements the OrganizationalUnitLister interface.
type organizationalUnitLister struct {
	indexer cache.Indexer
}

// NewOrganizationalUnitLister returns a new OrganizationalUnitLister.
func NewOrganizationalUnitLister(indexer cache.Indexer) OrganizationalUnitLister {
	return &organizationalUnitLister{indexer: indexer}
}

// List lists all OrganizationalUnits in the indexer.
func (s *organizationalUnitLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationalUnit, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationalUnit))
	})
	return ret, err
}

// OrganizationalUnits returns an object that can list and get OrganizationalUnits.
func (s *organizationalUnitLister) OrganizationalUnits(namespace string) OrganizationalUnitNamespaceLister {
	return organizationalUnitNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OrganizationalUnitNamespaceLister helps list and get OrganizationalUnits.
// All objects returned here must be treated as read-only.
type OrganizationalUnitNamespaceLister interface {
	// List lists all OrganizationalUnits in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OrganizationalUnit, err error)
	// Get retrieves the OrganizationalUnit from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OrganizationalUnit, error)
	OrganizationalUnitNamespaceListerExpansion
}

// organizationalUnitNamespaceLister implements the OrganizationalUnitNamespaceLister
// interface.
type organizationalUnitNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OrganizationalUnits in the indexer for a given namespace.
func (s organizationalUnitNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OrganizationalUnit, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OrganizationalUnit))
	})
	return ret, err
}

// Get retrieves the OrganizationalUnit from the indexer for a given namespace and name.
func (s organizationalUnitNamespaceLister) Get(name string) (*v1alpha1.OrganizationalUnit, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("organizationalunit"), name)
	}
	return obj.(*v1alpha1.OrganizationalUnit), nil
}
