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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/glaciervault/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LockLister helps list Locks.
// All objects returned here must be treated as read-only.
type LockLister interface {
	// List lists all Locks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Lock, err error)
	// Locks returns an object that can list and get Locks.
	Locks(namespace string) LockNamespaceLister
	LockListerExpansion
}

// lockLister implements the LockLister interface.
type lockLister struct {
	indexer cache.Indexer
}

// NewLockLister returns a new LockLister.
func NewLockLister(indexer cache.Indexer) LockLister {
	return &lockLister{indexer: indexer}
}

// List lists all Locks in the indexer.
func (s *lockLister) List(selector labels.Selector) (ret []*v1alpha1.Lock, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Lock))
	})
	return ret, err
}

// Locks returns an object that can list and get Locks.
func (s *lockLister) Locks(namespace string) LockNamespaceLister {
	return lockNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LockNamespaceLister helps list and get Locks.
// All objects returned here must be treated as read-only.
type LockNamespaceLister interface {
	// List lists all Locks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Lock, err error)
	// Get retrieves the Lock from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Lock, error)
	LockNamespaceListerExpansion
}

// lockNamespaceLister implements the LockNamespaceLister
// interface.
type lockNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Locks in the indexer for a given namespace.
func (s lockNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Lock, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Lock))
	})
	return ret, err
}

// Get retrieves the Lock from the indexer for a given namespace and name.
func (s lockNamespaceLister) Get(name string) (*v1alpha1.Lock, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lock"), name)
	}
	return obj.(*v1alpha1.Lock), nil
}
