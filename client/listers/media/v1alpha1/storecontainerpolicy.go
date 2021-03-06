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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/media/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StoreContainerPolicyLister helps list StoreContainerPolicies.
// All objects returned here must be treated as read-only.
type StoreContainerPolicyLister interface {
	// List lists all StoreContainerPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StoreContainerPolicy, err error)
	// StoreContainerPolicies returns an object that can list and get StoreContainerPolicies.
	StoreContainerPolicies(namespace string) StoreContainerPolicyNamespaceLister
	StoreContainerPolicyListerExpansion
}

// storeContainerPolicyLister implements the StoreContainerPolicyLister interface.
type storeContainerPolicyLister struct {
	indexer cache.Indexer
}

// NewStoreContainerPolicyLister returns a new StoreContainerPolicyLister.
func NewStoreContainerPolicyLister(indexer cache.Indexer) StoreContainerPolicyLister {
	return &storeContainerPolicyLister{indexer: indexer}
}

// List lists all StoreContainerPolicies in the indexer.
func (s *storeContainerPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.StoreContainerPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoreContainerPolicy))
	})
	return ret, err
}

// StoreContainerPolicies returns an object that can list and get StoreContainerPolicies.
func (s *storeContainerPolicyLister) StoreContainerPolicies(namespace string) StoreContainerPolicyNamespaceLister {
	return storeContainerPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StoreContainerPolicyNamespaceLister helps list and get StoreContainerPolicies.
// All objects returned here must be treated as read-only.
type StoreContainerPolicyNamespaceLister interface {
	// List lists all StoreContainerPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StoreContainerPolicy, err error)
	// Get retrieves the StoreContainerPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StoreContainerPolicy, error)
	StoreContainerPolicyNamespaceListerExpansion
}

// storeContainerPolicyNamespaceLister implements the StoreContainerPolicyNamespaceLister
// interface.
type storeContainerPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StoreContainerPolicies in the indexer for a given namespace.
func (s storeContainerPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StoreContainerPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoreContainerPolicy))
	})
	return ret, err
}

// Get retrieves the StoreContainerPolicy from the indexer for a given namespace and name.
func (s storeContainerPolicyNamespaceLister) Get(name string) (*v1alpha1.StoreContainerPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storecontainerpolicy"), name)
	}
	return obj.(*v1alpha1.StoreContainerPolicy), nil
}
