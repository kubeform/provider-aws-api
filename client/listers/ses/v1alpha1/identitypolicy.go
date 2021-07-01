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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"
)

// IdentityPolicyLister helps list IdentityPolicies.
// All objects returned here must be treated as read-only.
type IdentityPolicyLister interface {
	// List lists all IdentityPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IdentityPolicy, err error)
	// IdentityPolicies returns an object that can list and get IdentityPolicies.
	IdentityPolicies(namespace string) IdentityPolicyNamespaceLister
	IdentityPolicyListerExpansion
}

// identityPolicyLister implements the IdentityPolicyLister interface.
type identityPolicyLister struct {
	indexer cache.Indexer
}

// NewIdentityPolicyLister returns a new IdentityPolicyLister.
func NewIdentityPolicyLister(indexer cache.Indexer) IdentityPolicyLister {
	return &identityPolicyLister{indexer: indexer}
}

// List lists all IdentityPolicies in the indexer.
func (s *identityPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.IdentityPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IdentityPolicy))
	})
	return ret, err
}

// IdentityPolicies returns an object that can list and get IdentityPolicies.
func (s *identityPolicyLister) IdentityPolicies(namespace string) IdentityPolicyNamespaceLister {
	return identityPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IdentityPolicyNamespaceLister helps list and get IdentityPolicies.
// All objects returned here must be treated as read-only.
type IdentityPolicyNamespaceLister interface {
	// List lists all IdentityPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IdentityPolicy, err error)
	// Get retrieves the IdentityPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IdentityPolicy, error)
	IdentityPolicyNamespaceListerExpansion
}

// identityPolicyNamespaceLister implements the IdentityPolicyNamespaceLister
// interface.
type identityPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IdentityPolicies in the indexer for a given namespace.
func (s identityPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IdentityPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IdentityPolicy))
	})
	return ret, err
}

// Get retrieves the IdentityPolicy from the indexer for a given namespace and name.
func (s identityPolicyNamespaceLister) Get(name string) (*v1alpha1.IdentityPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("identitypolicy"), name)
	}
	return obj.(*v1alpha1.IdentityPolicy), nil
}
