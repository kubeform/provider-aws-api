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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecr/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PullThroughCacheRuleLister helps list PullThroughCacheRules.
// All objects returned here must be treated as read-only.
type PullThroughCacheRuleLister interface {
	// List lists all PullThroughCacheRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PullThroughCacheRule, err error)
	// PullThroughCacheRules returns an object that can list and get PullThroughCacheRules.
	PullThroughCacheRules(namespace string) PullThroughCacheRuleNamespaceLister
	PullThroughCacheRuleListerExpansion
}

// pullThroughCacheRuleLister implements the PullThroughCacheRuleLister interface.
type pullThroughCacheRuleLister struct {
	indexer cache.Indexer
}

// NewPullThroughCacheRuleLister returns a new PullThroughCacheRuleLister.
func NewPullThroughCacheRuleLister(indexer cache.Indexer) PullThroughCacheRuleLister {
	return &pullThroughCacheRuleLister{indexer: indexer}
}

// List lists all PullThroughCacheRules in the indexer.
func (s *pullThroughCacheRuleLister) List(selector labels.Selector) (ret []*v1alpha1.PullThroughCacheRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PullThroughCacheRule))
	})
	return ret, err
}

// PullThroughCacheRules returns an object that can list and get PullThroughCacheRules.
func (s *pullThroughCacheRuleLister) PullThroughCacheRules(namespace string) PullThroughCacheRuleNamespaceLister {
	return pullThroughCacheRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PullThroughCacheRuleNamespaceLister helps list and get PullThroughCacheRules.
// All objects returned here must be treated as read-only.
type PullThroughCacheRuleNamespaceLister interface {
	// List lists all PullThroughCacheRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PullThroughCacheRule, err error)
	// Get retrieves the PullThroughCacheRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PullThroughCacheRule, error)
	PullThroughCacheRuleNamespaceListerExpansion
}

// pullThroughCacheRuleNamespaceLister implements the PullThroughCacheRuleNamespaceLister
// interface.
type pullThroughCacheRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PullThroughCacheRules in the indexer for a given namespace.
func (s pullThroughCacheRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PullThroughCacheRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PullThroughCacheRule))
	})
	return ret, err
}

// Get retrieves the PullThroughCacheRule from the indexer for a given namespace and name.
func (s pullThroughCacheRuleNamespaceLister) Get(name string) (*v1alpha1.PullThroughCacheRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pullthroughcacherule"), name)
	}
	return obj.(*v1alpha1.PullThroughCacheRule), nil
}