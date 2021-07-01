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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/iam/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SamlProviderLister helps list SamlProviders.
// All objects returned here must be treated as read-only.
type SamlProviderLister interface {
	// List lists all SamlProviders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SamlProvider, err error)
	// SamlProviders returns an object that can list and get SamlProviders.
	SamlProviders(namespace string) SamlProviderNamespaceLister
	SamlProviderListerExpansion
}

// samlProviderLister implements the SamlProviderLister interface.
type samlProviderLister struct {
	indexer cache.Indexer
}

// NewSamlProviderLister returns a new SamlProviderLister.
func NewSamlProviderLister(indexer cache.Indexer) SamlProviderLister {
	return &samlProviderLister{indexer: indexer}
}

// List lists all SamlProviders in the indexer.
func (s *samlProviderLister) List(selector labels.Selector) (ret []*v1alpha1.SamlProvider, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SamlProvider))
	})
	return ret, err
}

// SamlProviders returns an object that can list and get SamlProviders.
func (s *samlProviderLister) SamlProviders(namespace string) SamlProviderNamespaceLister {
	return samlProviderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SamlProviderNamespaceLister helps list and get SamlProviders.
// All objects returned here must be treated as read-only.
type SamlProviderNamespaceLister interface {
	// List lists all SamlProviders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SamlProvider, err error)
	// Get retrieves the SamlProvider from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SamlProvider, error)
	SamlProviderNamespaceListerExpansion
}

// samlProviderNamespaceLister implements the SamlProviderNamespaceLister
// interface.
type samlProviderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SamlProviders in the indexer for a given namespace.
func (s samlProviderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SamlProvider, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SamlProvider))
	})
	return ret, err
}

// Get retrieves the SamlProvider from the indexer for a given namespace and name.
func (s samlProviderNamespaceLister) Get(name string) (*v1alpha1.SamlProvider, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("samlprovider"), name)
	}
	return obj.(*v1alpha1.SamlProvider), nil
}
