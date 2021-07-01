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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/directoryservice/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConditionalForwarderLister helps list ConditionalForwarders.
// All objects returned here must be treated as read-only.
type ConditionalForwarderLister interface {
	// List lists all ConditionalForwarders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConditionalForwarder, err error)
	// ConditionalForwarders returns an object that can list and get ConditionalForwarders.
	ConditionalForwarders(namespace string) ConditionalForwarderNamespaceLister
	ConditionalForwarderListerExpansion
}

// conditionalForwarderLister implements the ConditionalForwarderLister interface.
type conditionalForwarderLister struct {
	indexer cache.Indexer
}

// NewConditionalForwarderLister returns a new ConditionalForwarderLister.
func NewConditionalForwarderLister(indexer cache.Indexer) ConditionalForwarderLister {
	return &conditionalForwarderLister{indexer: indexer}
}

// List lists all ConditionalForwarders in the indexer.
func (s *conditionalForwarderLister) List(selector labels.Selector) (ret []*v1alpha1.ConditionalForwarder, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConditionalForwarder))
	})
	return ret, err
}

// ConditionalForwarders returns an object that can list and get ConditionalForwarders.
func (s *conditionalForwarderLister) ConditionalForwarders(namespace string) ConditionalForwarderNamespaceLister {
	return conditionalForwarderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConditionalForwarderNamespaceLister helps list and get ConditionalForwarders.
// All objects returned here must be treated as read-only.
type ConditionalForwarderNamespaceLister interface {
	// List lists all ConditionalForwarders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConditionalForwarder, err error)
	// Get retrieves the ConditionalForwarder from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConditionalForwarder, error)
	ConditionalForwarderNamespaceListerExpansion
}

// conditionalForwarderNamespaceLister implements the ConditionalForwarderNamespaceLister
// interface.
type conditionalForwarderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConditionalForwarders in the indexer for a given namespace.
func (s conditionalForwarderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConditionalForwarder, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConditionalForwarder))
	})
	return ret, err
}

// Get retrieves the ConditionalForwarder from the indexer for a given namespace and name.
func (s conditionalForwarderNamespaceLister) Get(name string) (*v1alpha1.ConditionalForwarder, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("conditionalforwarder"), name)
	}
	return obj.(*v1alpha1.ConditionalForwarder), nil
}
