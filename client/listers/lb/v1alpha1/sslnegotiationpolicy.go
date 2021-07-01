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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/lb/v1alpha1"
)

// SslNegotiationPolicyLister helps list SslNegotiationPolicies.
// All objects returned here must be treated as read-only.
type SslNegotiationPolicyLister interface {
	// List lists all SslNegotiationPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SslNegotiationPolicy, err error)
	// SslNegotiationPolicies returns an object that can list and get SslNegotiationPolicies.
	SslNegotiationPolicies(namespace string) SslNegotiationPolicyNamespaceLister
	SslNegotiationPolicyListerExpansion
}

// sslNegotiationPolicyLister implements the SslNegotiationPolicyLister interface.
type sslNegotiationPolicyLister struct {
	indexer cache.Indexer
}

// NewSslNegotiationPolicyLister returns a new SslNegotiationPolicyLister.
func NewSslNegotiationPolicyLister(indexer cache.Indexer) SslNegotiationPolicyLister {
	return &sslNegotiationPolicyLister{indexer: indexer}
}

// List lists all SslNegotiationPolicies in the indexer.
func (s *sslNegotiationPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.SslNegotiationPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SslNegotiationPolicy))
	})
	return ret, err
}

// SslNegotiationPolicies returns an object that can list and get SslNegotiationPolicies.
func (s *sslNegotiationPolicyLister) SslNegotiationPolicies(namespace string) SslNegotiationPolicyNamespaceLister {
	return sslNegotiationPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SslNegotiationPolicyNamespaceLister helps list and get SslNegotiationPolicies.
// All objects returned here must be treated as read-only.
type SslNegotiationPolicyNamespaceLister interface {
	// List lists all SslNegotiationPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SslNegotiationPolicy, err error)
	// Get retrieves the SslNegotiationPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SslNegotiationPolicy, error)
	SslNegotiationPolicyNamespaceListerExpansion
}

// sslNegotiationPolicyNamespaceLister implements the SslNegotiationPolicyNamespaceLister
// interface.
type sslNegotiationPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SslNegotiationPolicies in the indexer for a given namespace.
func (s sslNegotiationPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SslNegotiationPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SslNegotiationPolicy))
	})
	return ret, err
}

// Get retrieves the SslNegotiationPolicy from the indexer for a given namespace and name.
func (s sslNegotiationPolicyNamespaceLister) Get(name string) (*v1alpha1.SslNegotiationPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sslnegotiationpolicy"), name)
	}
	return obj.(*v1alpha1.SslNegotiationPolicy), nil
}
