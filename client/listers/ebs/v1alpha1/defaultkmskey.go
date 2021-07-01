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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ebs/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DefaultKmsKeyLister helps list DefaultKmsKeys.
// All objects returned here must be treated as read-only.
type DefaultKmsKeyLister interface {
	// List lists all DefaultKmsKeys in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultKmsKey, err error)
	// DefaultKmsKeys returns an object that can list and get DefaultKmsKeys.
	DefaultKmsKeys(namespace string) DefaultKmsKeyNamespaceLister
	DefaultKmsKeyListerExpansion
}

// defaultKmsKeyLister implements the DefaultKmsKeyLister interface.
type defaultKmsKeyLister struct {
	indexer cache.Indexer
}

// NewDefaultKmsKeyLister returns a new DefaultKmsKeyLister.
func NewDefaultKmsKeyLister(indexer cache.Indexer) DefaultKmsKeyLister {
	return &defaultKmsKeyLister{indexer: indexer}
}

// List lists all DefaultKmsKeys in the indexer.
func (s *defaultKmsKeyLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultKmsKey, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultKmsKey))
	})
	return ret, err
}

// DefaultKmsKeys returns an object that can list and get DefaultKmsKeys.
func (s *defaultKmsKeyLister) DefaultKmsKeys(namespace string) DefaultKmsKeyNamespaceLister {
	return defaultKmsKeyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DefaultKmsKeyNamespaceLister helps list and get DefaultKmsKeys.
// All objects returned here must be treated as read-only.
type DefaultKmsKeyNamespaceLister interface {
	// List lists all DefaultKmsKeys in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DefaultKmsKey, err error)
	// Get retrieves the DefaultKmsKey from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DefaultKmsKey, error)
	DefaultKmsKeyNamespaceListerExpansion
}

// defaultKmsKeyNamespaceLister implements the DefaultKmsKeyNamespaceLister
// interface.
type defaultKmsKeyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DefaultKmsKeys in the indexer for a given namespace.
func (s defaultKmsKeyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DefaultKmsKey, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DefaultKmsKey))
	})
	return ret, err
}

// Get retrieves the DefaultKmsKey from the indexer for a given namespace and name.
func (s defaultKmsKeyNamespaceLister) Get(name string) (*v1alpha1.DefaultKmsKey, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("defaultkmskey"), name)
	}
	return obj.(*v1alpha1.DefaultKmsKey), nil
}
