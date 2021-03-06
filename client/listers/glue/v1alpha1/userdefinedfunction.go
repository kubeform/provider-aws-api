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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/glue/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UserDefinedFunctionLister helps list UserDefinedFunctions.
// All objects returned here must be treated as read-only.
type UserDefinedFunctionLister interface {
	// List lists all UserDefinedFunctions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.UserDefinedFunction, err error)
	// UserDefinedFunctions returns an object that can list and get UserDefinedFunctions.
	UserDefinedFunctions(namespace string) UserDefinedFunctionNamespaceLister
	UserDefinedFunctionListerExpansion
}

// userDefinedFunctionLister implements the UserDefinedFunctionLister interface.
type userDefinedFunctionLister struct {
	indexer cache.Indexer
}

// NewUserDefinedFunctionLister returns a new UserDefinedFunctionLister.
func NewUserDefinedFunctionLister(indexer cache.Indexer) UserDefinedFunctionLister {
	return &userDefinedFunctionLister{indexer: indexer}
}

// List lists all UserDefinedFunctions in the indexer.
func (s *userDefinedFunctionLister) List(selector labels.Selector) (ret []*v1alpha1.UserDefinedFunction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserDefinedFunction))
	})
	return ret, err
}

// UserDefinedFunctions returns an object that can list and get UserDefinedFunctions.
func (s *userDefinedFunctionLister) UserDefinedFunctions(namespace string) UserDefinedFunctionNamespaceLister {
	return userDefinedFunctionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserDefinedFunctionNamespaceLister helps list and get UserDefinedFunctions.
// All objects returned here must be treated as read-only.
type UserDefinedFunctionNamespaceLister interface {
	// List lists all UserDefinedFunctions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.UserDefinedFunction, err error)
	// Get retrieves the UserDefinedFunction from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.UserDefinedFunction, error)
	UserDefinedFunctionNamespaceListerExpansion
}

// userDefinedFunctionNamespaceLister implements the UserDefinedFunctionNamespaceLister
// interface.
type userDefinedFunctionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all UserDefinedFunctions in the indexer for a given namespace.
func (s userDefinedFunctionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.UserDefinedFunction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserDefinedFunction))
	})
	return ret, err
}

// Get retrieves the UserDefinedFunction from the indexer for a given namespace and name.
func (s userDefinedFunctionNamespaceLister) Get(name string) (*v1alpha1.UserDefinedFunction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("userdefinedfunction"), name)
	}
	return obj.(*v1alpha1.UserDefinedFunction), nil
}
