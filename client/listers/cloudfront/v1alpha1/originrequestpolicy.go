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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudfront/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OriginRequestPolicyLister helps list OriginRequestPolicies.
// All objects returned here must be treated as read-only.
type OriginRequestPolicyLister interface {
	// List lists all OriginRequestPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OriginRequestPolicy, err error)
	// OriginRequestPolicies returns an object that can list and get OriginRequestPolicies.
	OriginRequestPolicies(namespace string) OriginRequestPolicyNamespaceLister
	OriginRequestPolicyListerExpansion
}

// originRequestPolicyLister implements the OriginRequestPolicyLister interface.
type originRequestPolicyLister struct {
	indexer cache.Indexer
}

// NewOriginRequestPolicyLister returns a new OriginRequestPolicyLister.
func NewOriginRequestPolicyLister(indexer cache.Indexer) OriginRequestPolicyLister {
	return &originRequestPolicyLister{indexer: indexer}
}

// List lists all OriginRequestPolicies in the indexer.
func (s *originRequestPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.OriginRequestPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OriginRequestPolicy))
	})
	return ret, err
}

// OriginRequestPolicies returns an object that can list and get OriginRequestPolicies.
func (s *originRequestPolicyLister) OriginRequestPolicies(namespace string) OriginRequestPolicyNamespaceLister {
	return originRequestPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OriginRequestPolicyNamespaceLister helps list and get OriginRequestPolicies.
// All objects returned here must be treated as read-only.
type OriginRequestPolicyNamespaceLister interface {
	// List lists all OriginRequestPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OriginRequestPolicy, err error)
	// Get retrieves the OriginRequestPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OriginRequestPolicy, error)
	OriginRequestPolicyNamespaceListerExpansion
}

// originRequestPolicyNamespaceLister implements the OriginRequestPolicyNamespaceLister
// interface.
type originRequestPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OriginRequestPolicies in the indexer for a given namespace.
func (s originRequestPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OriginRequestPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OriginRequestPolicy))
	})
	return ret, err
}

// Get retrieves the OriginRequestPolicy from the indexer for a given namespace and name.
func (s originRequestPolicyNamespaceLister) Get(name string) (*v1alpha1.OriginRequestPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("originrequestpolicy"), name)
	}
	return obj.(*v1alpha1.OriginRequestPolicy), nil
}
