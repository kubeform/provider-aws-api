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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReceiptRuleSetLister helps list ReceiptRuleSets.
// All objects returned here must be treated as read-only.
type ReceiptRuleSetLister interface {
	// List lists all ReceiptRuleSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReceiptRuleSet, err error)
	// ReceiptRuleSets returns an object that can list and get ReceiptRuleSets.
	ReceiptRuleSets(namespace string) ReceiptRuleSetNamespaceLister
	ReceiptRuleSetListerExpansion
}

// receiptRuleSetLister implements the ReceiptRuleSetLister interface.
type receiptRuleSetLister struct {
	indexer cache.Indexer
}

// NewReceiptRuleSetLister returns a new ReceiptRuleSetLister.
func NewReceiptRuleSetLister(indexer cache.Indexer) ReceiptRuleSetLister {
	return &receiptRuleSetLister{indexer: indexer}
}

// List lists all ReceiptRuleSets in the indexer.
func (s *receiptRuleSetLister) List(selector labels.Selector) (ret []*v1alpha1.ReceiptRuleSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReceiptRuleSet))
	})
	return ret, err
}

// ReceiptRuleSets returns an object that can list and get ReceiptRuleSets.
func (s *receiptRuleSetLister) ReceiptRuleSets(namespace string) ReceiptRuleSetNamespaceLister {
	return receiptRuleSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReceiptRuleSetNamespaceLister helps list and get ReceiptRuleSets.
// All objects returned here must be treated as read-only.
type ReceiptRuleSetNamespaceLister interface {
	// List lists all ReceiptRuleSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReceiptRuleSet, err error)
	// Get retrieves the ReceiptRuleSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ReceiptRuleSet, error)
	ReceiptRuleSetNamespaceListerExpansion
}

// receiptRuleSetNamespaceLister implements the ReceiptRuleSetNamespaceLister
// interface.
type receiptRuleSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReceiptRuleSets in the indexer for a given namespace.
func (s receiptRuleSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ReceiptRuleSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReceiptRuleSet))
	})
	return ret, err
}

// Get retrieves the ReceiptRuleSet from the indexer for a given namespace and name.
func (s receiptRuleSetNamespaceLister) Get(name string) (*v1alpha1.ReceiptRuleSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("receiptruleset"), name)
	}
	return obj.(*v1alpha1.ReceiptRuleSet), nil
}
