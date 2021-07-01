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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/synthetics/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CanaryLister helps list Canaries.
// All objects returned here must be treated as read-only.
type CanaryLister interface {
	// List lists all Canaries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Canary, err error)
	// Canaries returns an object that can list and get Canaries.
	Canaries(namespace string) CanaryNamespaceLister
	CanaryListerExpansion
}

// canaryLister implements the CanaryLister interface.
type canaryLister struct {
	indexer cache.Indexer
}

// NewCanaryLister returns a new CanaryLister.
func NewCanaryLister(indexer cache.Indexer) CanaryLister {
	return &canaryLister{indexer: indexer}
}

// List lists all Canaries in the indexer.
func (s *canaryLister) List(selector labels.Selector) (ret []*v1alpha1.Canary, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Canary))
	})
	return ret, err
}

// Canaries returns an object that can list and get Canaries.
func (s *canaryLister) Canaries(namespace string) CanaryNamespaceLister {
	return canaryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CanaryNamespaceLister helps list and get Canaries.
// All objects returned here must be treated as read-only.
type CanaryNamespaceLister interface {
	// List lists all Canaries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Canary, err error)
	// Get retrieves the Canary from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Canary, error)
	CanaryNamespaceListerExpansion
}

// canaryNamespaceLister implements the CanaryNamespaceLister
// interface.
type canaryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Canaries in the indexer for a given namespace.
func (s canaryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Canary, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Canary))
	})
	return ret, err
}

// Get retrieves the Canary from the indexer for a given namespace and name.
func (s canaryNamespaceLister) Get(name string) (*v1alpha1.Canary, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("canary"), name)
	}
	return obj.(*v1alpha1.Canary), nil
}
