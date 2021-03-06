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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/app/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CookieStickinessPolicyLister helps list CookieStickinessPolicies.
// All objects returned here must be treated as read-only.
type CookieStickinessPolicyLister interface {
	// List lists all CookieStickinessPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CookieStickinessPolicy, err error)
	// CookieStickinessPolicies returns an object that can list and get CookieStickinessPolicies.
	CookieStickinessPolicies(namespace string) CookieStickinessPolicyNamespaceLister
	CookieStickinessPolicyListerExpansion
}

// cookieStickinessPolicyLister implements the CookieStickinessPolicyLister interface.
type cookieStickinessPolicyLister struct {
	indexer cache.Indexer
}

// NewCookieStickinessPolicyLister returns a new CookieStickinessPolicyLister.
func NewCookieStickinessPolicyLister(indexer cache.Indexer) CookieStickinessPolicyLister {
	return &cookieStickinessPolicyLister{indexer: indexer}
}

// List lists all CookieStickinessPolicies in the indexer.
func (s *cookieStickinessPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.CookieStickinessPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CookieStickinessPolicy))
	})
	return ret, err
}

// CookieStickinessPolicies returns an object that can list and get CookieStickinessPolicies.
func (s *cookieStickinessPolicyLister) CookieStickinessPolicies(namespace string) CookieStickinessPolicyNamespaceLister {
	return cookieStickinessPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CookieStickinessPolicyNamespaceLister helps list and get CookieStickinessPolicies.
// All objects returned here must be treated as read-only.
type CookieStickinessPolicyNamespaceLister interface {
	// List lists all CookieStickinessPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CookieStickinessPolicy, err error)
	// Get retrieves the CookieStickinessPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CookieStickinessPolicy, error)
	CookieStickinessPolicyNamespaceListerExpansion
}

// cookieStickinessPolicyNamespaceLister implements the CookieStickinessPolicyNamespaceLister
// interface.
type cookieStickinessPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CookieStickinessPolicies in the indexer for a given namespace.
func (s cookieStickinessPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CookieStickinessPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CookieStickinessPolicy))
	})
	return ret, err
}

// Get retrieves the CookieStickinessPolicy from the indexer for a given namespace and name.
func (s cookieStickinessPolicyNamespaceLister) Get(name string) (*v1alpha1.CookieStickinessPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cookiestickinesspolicy"), name)
	}
	return obj.(*v1alpha1.CookieStickinessPolicy), nil
}
