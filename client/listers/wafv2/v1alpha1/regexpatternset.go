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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/wafv2/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RegexPatternSetLister helps list RegexPatternSets.
// All objects returned here must be treated as read-only.
type RegexPatternSetLister interface {
	// List lists all RegexPatternSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegexPatternSet, err error)
	// RegexPatternSets returns an object that can list and get RegexPatternSets.
	RegexPatternSets(namespace string) RegexPatternSetNamespaceLister
	RegexPatternSetListerExpansion
}

// regexPatternSetLister implements the RegexPatternSetLister interface.
type regexPatternSetLister struct {
	indexer cache.Indexer
}

// NewRegexPatternSetLister returns a new RegexPatternSetLister.
func NewRegexPatternSetLister(indexer cache.Indexer) RegexPatternSetLister {
	return &regexPatternSetLister{indexer: indexer}
}

// List lists all RegexPatternSets in the indexer.
func (s *regexPatternSetLister) List(selector labels.Selector) (ret []*v1alpha1.RegexPatternSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegexPatternSet))
	})
	return ret, err
}

// RegexPatternSets returns an object that can list and get RegexPatternSets.
func (s *regexPatternSetLister) RegexPatternSets(namespace string) RegexPatternSetNamespaceLister {
	return regexPatternSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RegexPatternSetNamespaceLister helps list and get RegexPatternSets.
// All objects returned here must be treated as read-only.
type RegexPatternSetNamespaceLister interface {
	// List lists all RegexPatternSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegexPatternSet, err error)
	// Get retrieves the RegexPatternSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RegexPatternSet, error)
	RegexPatternSetNamespaceListerExpansion
}

// regexPatternSetNamespaceLister implements the RegexPatternSetNamespaceLister
// interface.
type regexPatternSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RegexPatternSets in the indexer for a given namespace.
func (s regexPatternSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RegexPatternSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegexPatternSet))
	})
	return ret, err
}

// Get retrieves the RegexPatternSet from the indexer for a given namespace and name.
func (s regexPatternSetNamespaceLister) Get(name string) (*v1alpha1.RegexPatternSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("regexpatternset"), name)
	}
	return obj.(*v1alpha1.RegexPatternSet), nil
}
