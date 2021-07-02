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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cognito/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UserPoolLister helps list UserPools.
// All objects returned here must be treated as read-only.
type UserPoolLister interface {
	// List lists all UserPools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.UserPool, err error)
	// UserPools returns an object that can list and get UserPools.
	UserPools(namespace string) UserPoolNamespaceLister
	UserPoolListerExpansion
}

// userPoolLister implements the UserPoolLister interface.
type userPoolLister struct {
	indexer cache.Indexer
}

// NewUserPoolLister returns a new UserPoolLister.
func NewUserPoolLister(indexer cache.Indexer) UserPoolLister {
	return &userPoolLister{indexer: indexer}
}

// List lists all UserPools in the indexer.
func (s *userPoolLister) List(selector labels.Selector) (ret []*v1alpha1.UserPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserPool))
	})
	return ret, err
}

// UserPools returns an object that can list and get UserPools.
func (s *userPoolLister) UserPools(namespace string) UserPoolNamespaceLister {
	return userPoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserPoolNamespaceLister helps list and get UserPools.
// All objects returned here must be treated as read-only.
type UserPoolNamespaceLister interface {
	// List lists all UserPools in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.UserPool, err error)
	// Get retrieves the UserPool from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.UserPool, error)
	UserPoolNamespaceListerExpansion
}

// userPoolNamespaceLister implements the UserPoolNamespaceLister
// interface.
type userPoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all UserPools in the indexer for a given namespace.
func (s userPoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.UserPool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserPool))
	})
	return ret, err
}

// Get retrieves the UserPool from the indexer for a given namespace and name.
func (s userPoolNamespaceLister) Get(name string) (*v1alpha1.UserPool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("userpool"), name)
	}
	return obj.(*v1alpha1.UserPool), nil
}
