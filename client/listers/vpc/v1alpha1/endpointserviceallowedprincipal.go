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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/vpc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EndpointServiceAllowedPrincipalLister helps list EndpointServiceAllowedPrincipals.
// All objects returned here must be treated as read-only.
type EndpointServiceAllowedPrincipalLister interface {
	// List lists all EndpointServiceAllowedPrincipals in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointServiceAllowedPrincipal, err error)
	// EndpointServiceAllowedPrincipals returns an object that can list and get EndpointServiceAllowedPrincipals.
	EndpointServiceAllowedPrincipals(namespace string) EndpointServiceAllowedPrincipalNamespaceLister
	EndpointServiceAllowedPrincipalListerExpansion
}

// endpointServiceAllowedPrincipalLister implements the EndpointServiceAllowedPrincipalLister interface.
type endpointServiceAllowedPrincipalLister struct {
	indexer cache.Indexer
}

// NewEndpointServiceAllowedPrincipalLister returns a new EndpointServiceAllowedPrincipalLister.
func NewEndpointServiceAllowedPrincipalLister(indexer cache.Indexer) EndpointServiceAllowedPrincipalLister {
	return &endpointServiceAllowedPrincipalLister{indexer: indexer}
}

// List lists all EndpointServiceAllowedPrincipals in the indexer.
func (s *endpointServiceAllowedPrincipalLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointServiceAllowedPrincipal, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointServiceAllowedPrincipal))
	})
	return ret, err
}

// EndpointServiceAllowedPrincipals returns an object that can list and get EndpointServiceAllowedPrincipals.
func (s *endpointServiceAllowedPrincipalLister) EndpointServiceAllowedPrincipals(namespace string) EndpointServiceAllowedPrincipalNamespaceLister {
	return endpointServiceAllowedPrincipalNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointServiceAllowedPrincipalNamespaceLister helps list and get EndpointServiceAllowedPrincipals.
// All objects returned here must be treated as read-only.
type EndpointServiceAllowedPrincipalNamespaceLister interface {
	// List lists all EndpointServiceAllowedPrincipals in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointServiceAllowedPrincipal, err error)
	// Get retrieves the EndpointServiceAllowedPrincipal from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EndpointServiceAllowedPrincipal, error)
	EndpointServiceAllowedPrincipalNamespaceListerExpansion
}

// endpointServiceAllowedPrincipalNamespaceLister implements the EndpointServiceAllowedPrincipalNamespaceLister
// interface.
type endpointServiceAllowedPrincipalNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EndpointServiceAllowedPrincipals in the indexer for a given namespace.
func (s endpointServiceAllowedPrincipalNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointServiceAllowedPrincipal, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointServiceAllowedPrincipal))
	})
	return ret, err
}

// Get retrieves the EndpointServiceAllowedPrincipal from the indexer for a given namespace and name.
func (s endpointServiceAllowedPrincipalNamespaceLister) Get(name string) (*v1alpha1.EndpointServiceAllowedPrincipal, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("endpointserviceallowedprincipal"), name)
	}
	return obj.(*v1alpha1.EndpointServiceAllowedPrincipal), nil
}
