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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecs/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterCapacityProvidersLister helps list ClusterCapacityProviderses.
// All objects returned here must be treated as read-only.
type ClusterCapacityProvidersLister interface {
	// List lists all ClusterCapacityProviderses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterCapacityProviders, err error)
	// ClusterCapacityProviderses returns an object that can list and get ClusterCapacityProviderses.
	ClusterCapacityProviderses(namespace string) ClusterCapacityProvidersNamespaceLister
	ClusterCapacityProvidersListerExpansion
}

// clusterCapacityProvidersLister implements the ClusterCapacityProvidersLister interface.
type clusterCapacityProvidersLister struct {
	indexer cache.Indexer
}

// NewClusterCapacityProvidersLister returns a new ClusterCapacityProvidersLister.
func NewClusterCapacityProvidersLister(indexer cache.Indexer) ClusterCapacityProvidersLister {
	return &clusterCapacityProvidersLister{indexer: indexer}
}

// List lists all ClusterCapacityProviderses in the indexer.
func (s *clusterCapacityProvidersLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterCapacityProviders, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterCapacityProviders))
	})
	return ret, err
}

// ClusterCapacityProviderses returns an object that can list and get ClusterCapacityProviderses.
func (s *clusterCapacityProvidersLister) ClusterCapacityProviderses(namespace string) ClusterCapacityProvidersNamespaceLister {
	return clusterCapacityProvidersNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClusterCapacityProvidersNamespaceLister helps list and get ClusterCapacityProviderses.
// All objects returned here must be treated as read-only.
type ClusterCapacityProvidersNamespaceLister interface {
	// List lists all ClusterCapacityProviderses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClusterCapacityProviders, err error)
	// Get retrieves the ClusterCapacityProviders from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClusterCapacityProviders, error)
	ClusterCapacityProvidersNamespaceListerExpansion
}

// clusterCapacityProvidersNamespaceLister implements the ClusterCapacityProvidersNamespaceLister
// interface.
type clusterCapacityProvidersNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClusterCapacityProviderses in the indexer for a given namespace.
func (s clusterCapacityProvidersNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClusterCapacityProviders, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClusterCapacityProviders))
	})
	return ret, err
}

// Get retrieves the ClusterCapacityProviders from the indexer for a given namespace and name.
func (s clusterCapacityProvidersNamespaceLister) Get(name string) (*v1alpha1.ClusterCapacityProviders, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clustercapacityproviders"), name)
	}
	return obj.(*v1alpha1.ClusterCapacityProviders), nil
}
