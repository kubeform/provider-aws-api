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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/network/v1alpha1"
)

// AclRuleLister helps list AclRules.
// All objects returned here must be treated as read-only.
type AclRuleLister interface {
	// List lists all AclRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AclRule, err error)
	// AclRules returns an object that can list and get AclRules.
	AclRules(namespace string) AclRuleNamespaceLister
	AclRuleListerExpansion
}

// aclRuleLister implements the AclRuleLister interface.
type aclRuleLister struct {
	indexer cache.Indexer
}

// NewAclRuleLister returns a new AclRuleLister.
func NewAclRuleLister(indexer cache.Indexer) AclRuleLister {
	return &aclRuleLister{indexer: indexer}
}

// List lists all AclRules in the indexer.
func (s *aclRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AclRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AclRule))
	})
	return ret, err
}

// AclRules returns an object that can list and get AclRules.
func (s *aclRuleLister) AclRules(namespace string) AclRuleNamespaceLister {
	return aclRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AclRuleNamespaceLister helps list and get AclRules.
// All objects returned here must be treated as read-only.
type AclRuleNamespaceLister interface {
	// List lists all AclRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AclRule, err error)
	// Get retrieves the AclRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AclRule, error)
	AclRuleNamespaceListerExpansion
}

// aclRuleNamespaceLister implements the AclRuleNamespaceLister
// interface.
type aclRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AclRules in the indexer for a given namespace.
func (s aclRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AclRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AclRule))
	})
	return ret, err
}

// Get retrieves the AclRule from the indexer for a given namespace and name.
func (s aclRuleNamespaceLister) Get(name string) (*v1alpha1.AclRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("aclrule"), name)
	}
	return obj.(*v1alpha1.AclRule), nil
}
