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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/waf/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ByteMatchSetLister helps list ByteMatchSets.
// All objects returned here must be treated as read-only.
type ByteMatchSetLister interface {
	// List lists all ByteMatchSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ByteMatchSet, err error)
	// ByteMatchSets returns an object that can list and get ByteMatchSets.
	ByteMatchSets(namespace string) ByteMatchSetNamespaceLister
	ByteMatchSetListerExpansion
}

// byteMatchSetLister implements the ByteMatchSetLister interface.
type byteMatchSetLister struct {
	indexer cache.Indexer
}

// NewByteMatchSetLister returns a new ByteMatchSetLister.
func NewByteMatchSetLister(indexer cache.Indexer) ByteMatchSetLister {
	return &byteMatchSetLister{indexer: indexer}
}

// List lists all ByteMatchSets in the indexer.
func (s *byteMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.ByteMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ByteMatchSet))
	})
	return ret, err
}

// ByteMatchSets returns an object that can list and get ByteMatchSets.
func (s *byteMatchSetLister) ByteMatchSets(namespace string) ByteMatchSetNamespaceLister {
	return byteMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ByteMatchSetNamespaceLister helps list and get ByteMatchSets.
// All objects returned here must be treated as read-only.
type ByteMatchSetNamespaceLister interface {
	// List lists all ByteMatchSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ByteMatchSet, err error)
	// Get retrieves the ByteMatchSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ByteMatchSet, error)
	ByteMatchSetNamespaceListerExpansion
}

// byteMatchSetNamespaceLister implements the ByteMatchSetNamespaceLister
// interface.
type byteMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ByteMatchSets in the indexer for a given namespace.
func (s byteMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ByteMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ByteMatchSet))
	})
	return ret, err
}

// Get retrieves the ByteMatchSet from the indexer for a given namespace and name.
func (s byteMatchSetNamespaceLister) Get(name string) (*v1alpha1.ByteMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bytematchset"), name)
	}
	return obj.(*v1alpha1.ByteMatchSet), nil
}
