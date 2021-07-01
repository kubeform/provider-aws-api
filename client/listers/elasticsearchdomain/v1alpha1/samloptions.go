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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/elasticsearchdomain/v1alpha1"
)

// SamlOptionsLister helps list SamlOptionses.
// All objects returned here must be treated as read-only.
type SamlOptionsLister interface {
	// List lists all SamlOptionses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SamlOptions, err error)
	// SamlOptionses returns an object that can list and get SamlOptionses.
	SamlOptionses(namespace string) SamlOptionsNamespaceLister
	SamlOptionsListerExpansion
}

// samlOptionsLister implements the SamlOptionsLister interface.
type samlOptionsLister struct {
	indexer cache.Indexer
}

// NewSamlOptionsLister returns a new SamlOptionsLister.
func NewSamlOptionsLister(indexer cache.Indexer) SamlOptionsLister {
	return &samlOptionsLister{indexer: indexer}
}

// List lists all SamlOptionses in the indexer.
func (s *samlOptionsLister) List(selector labels.Selector) (ret []*v1alpha1.SamlOptions, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SamlOptions))
	})
	return ret, err
}

// SamlOptionses returns an object that can list and get SamlOptionses.
func (s *samlOptionsLister) SamlOptionses(namespace string) SamlOptionsNamespaceLister {
	return samlOptionsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SamlOptionsNamespaceLister helps list and get SamlOptionses.
// All objects returned here must be treated as read-only.
type SamlOptionsNamespaceLister interface {
	// List lists all SamlOptionses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SamlOptions, err error)
	// Get retrieves the SamlOptions from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SamlOptions, error)
	SamlOptionsNamespaceListerExpansion
}

// samlOptionsNamespaceLister implements the SamlOptionsNamespaceLister
// interface.
type samlOptionsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SamlOptionses in the indexer for a given namespace.
func (s samlOptionsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SamlOptions, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SamlOptions))
	})
	return ret, err
}

// Get retrieves the SamlOptions from the indexer for a given namespace and name.
func (s samlOptionsNamespaceLister) Get(name string) (*v1alpha1.SamlOptions, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("samloptions"), name)
	}
	return obj.(*v1alpha1.SamlOptions), nil
}
