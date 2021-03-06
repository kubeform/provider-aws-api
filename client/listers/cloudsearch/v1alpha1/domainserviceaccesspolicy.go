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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudsearch/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainServiceAccessPolicyLister helps list DomainServiceAccessPolicies.
// All objects returned here must be treated as read-only.
type DomainServiceAccessPolicyLister interface {
	// List lists all DomainServiceAccessPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainServiceAccessPolicy, err error)
	// DomainServiceAccessPolicies returns an object that can list and get DomainServiceAccessPolicies.
	DomainServiceAccessPolicies(namespace string) DomainServiceAccessPolicyNamespaceLister
	DomainServiceAccessPolicyListerExpansion
}

// domainServiceAccessPolicyLister implements the DomainServiceAccessPolicyLister interface.
type domainServiceAccessPolicyLister struct {
	indexer cache.Indexer
}

// NewDomainServiceAccessPolicyLister returns a new DomainServiceAccessPolicyLister.
func NewDomainServiceAccessPolicyLister(indexer cache.Indexer) DomainServiceAccessPolicyLister {
	return &domainServiceAccessPolicyLister{indexer: indexer}
}

// List lists all DomainServiceAccessPolicies in the indexer.
func (s *domainServiceAccessPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.DomainServiceAccessPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainServiceAccessPolicy))
	})
	return ret, err
}

// DomainServiceAccessPolicies returns an object that can list and get DomainServiceAccessPolicies.
func (s *domainServiceAccessPolicyLister) DomainServiceAccessPolicies(namespace string) DomainServiceAccessPolicyNamespaceLister {
	return domainServiceAccessPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainServiceAccessPolicyNamespaceLister helps list and get DomainServiceAccessPolicies.
// All objects returned here must be treated as read-only.
type DomainServiceAccessPolicyNamespaceLister interface {
	// List lists all DomainServiceAccessPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainServiceAccessPolicy, err error)
	// Get retrieves the DomainServiceAccessPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainServiceAccessPolicy, error)
	DomainServiceAccessPolicyNamespaceListerExpansion
}

// domainServiceAccessPolicyNamespaceLister implements the DomainServiceAccessPolicyNamespaceLister
// interface.
type domainServiceAccessPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainServiceAccessPolicies in the indexer for a given namespace.
func (s domainServiceAccessPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainServiceAccessPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainServiceAccessPolicy))
	})
	return ret, err
}

// Get retrieves the DomainServiceAccessPolicy from the indexer for a given namespace and name.
func (s domainServiceAccessPolicyNamespaceLister) Get(name string) (*v1alpha1.DomainServiceAccessPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainserviceaccesspolicy"), name)
	}
	return obj.(*v1alpha1.DomainServiceAccessPolicy), nil
}
