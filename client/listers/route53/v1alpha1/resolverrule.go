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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/route53/v1alpha1"
)

// ResolverRuleLister helps list ResolverRules.
// All objects returned here must be treated as read-only.
type ResolverRuleLister interface {
	// List lists all ResolverRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResolverRule, err error)
	// ResolverRules returns an object that can list and get ResolverRules.
	ResolverRules(namespace string) ResolverRuleNamespaceLister
	ResolverRuleListerExpansion
}

// resolverRuleLister implements the ResolverRuleLister interface.
type resolverRuleLister struct {
	indexer cache.Indexer
}

// NewResolverRuleLister returns a new ResolverRuleLister.
func NewResolverRuleLister(indexer cache.Indexer) ResolverRuleLister {
	return &resolverRuleLister{indexer: indexer}
}

// List lists all ResolverRules in the indexer.
func (s *resolverRuleLister) List(selector labels.Selector) (ret []*v1alpha1.ResolverRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResolverRule))
	})
	return ret, err
}

// ResolverRules returns an object that can list and get ResolverRules.
func (s *resolverRuleLister) ResolverRules(namespace string) ResolverRuleNamespaceLister {
	return resolverRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResolverRuleNamespaceLister helps list and get ResolverRules.
// All objects returned here must be treated as read-only.
type ResolverRuleNamespaceLister interface {
	// List lists all ResolverRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResolverRule, err error)
	// Get retrieves the ResolverRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResolverRule, error)
	ResolverRuleNamespaceListerExpansion
}

// resolverRuleNamespaceLister implements the ResolverRuleNamespaceLister
// interface.
type resolverRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResolverRules in the indexer for a given namespace.
func (s resolverRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResolverRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResolverRule))
	})
	return ret, err
}

// Get retrieves the ResolverRule from the indexer for a given namespace and name.
func (s resolverRuleNamespaceLister) Get(name string) (*v1alpha1.ResolverRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resolverrule"), name)
	}
	return obj.(*v1alpha1.ResolverRule), nil
}
