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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/apigateway/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RequestValidatorLister helps list RequestValidators.
// All objects returned here must be treated as read-only.
type RequestValidatorLister interface {
	// List lists all RequestValidators in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RequestValidator, err error)
	// RequestValidators returns an object that can list and get RequestValidators.
	RequestValidators(namespace string) RequestValidatorNamespaceLister
	RequestValidatorListerExpansion
}

// requestValidatorLister implements the RequestValidatorLister interface.
type requestValidatorLister struct {
	indexer cache.Indexer
}

// NewRequestValidatorLister returns a new RequestValidatorLister.
func NewRequestValidatorLister(indexer cache.Indexer) RequestValidatorLister {
	return &requestValidatorLister{indexer: indexer}
}

// List lists all RequestValidators in the indexer.
func (s *requestValidatorLister) List(selector labels.Selector) (ret []*v1alpha1.RequestValidator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RequestValidator))
	})
	return ret, err
}

// RequestValidators returns an object that can list and get RequestValidators.
func (s *requestValidatorLister) RequestValidators(namespace string) RequestValidatorNamespaceLister {
	return requestValidatorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RequestValidatorNamespaceLister helps list and get RequestValidators.
// All objects returned here must be treated as read-only.
type RequestValidatorNamespaceLister interface {
	// List lists all RequestValidators in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RequestValidator, err error)
	// Get retrieves the RequestValidator from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RequestValidator, error)
	RequestValidatorNamespaceListerExpansion
}

// requestValidatorNamespaceLister implements the RequestValidatorNamespaceLister
// interface.
type requestValidatorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RequestValidators in the indexer for a given namespace.
func (s requestValidatorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RequestValidator, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RequestValidator))
	})
	return ret, err
}

// Get retrieves the RequestValidator from the indexer for a given namespace and name.
func (s requestValidatorNamespaceLister) Get(name string) (*v1alpha1.RequestValidator, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("requestvalidator"), name)
	}
	return obj.(*v1alpha1.RequestValidator), nil
}
