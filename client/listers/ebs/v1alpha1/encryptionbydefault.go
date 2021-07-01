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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ebs/v1alpha1"
)

// EncryptionByDefaultLister helps list EncryptionByDefaults.
// All objects returned here must be treated as read-only.
type EncryptionByDefaultLister interface {
	// List lists all EncryptionByDefaults in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EncryptionByDefault, err error)
	// EncryptionByDefaults returns an object that can list and get EncryptionByDefaults.
	EncryptionByDefaults(namespace string) EncryptionByDefaultNamespaceLister
	EncryptionByDefaultListerExpansion
}

// encryptionByDefaultLister implements the EncryptionByDefaultLister interface.
type encryptionByDefaultLister struct {
	indexer cache.Indexer
}

// NewEncryptionByDefaultLister returns a new EncryptionByDefaultLister.
func NewEncryptionByDefaultLister(indexer cache.Indexer) EncryptionByDefaultLister {
	return &encryptionByDefaultLister{indexer: indexer}
}

// List lists all EncryptionByDefaults in the indexer.
func (s *encryptionByDefaultLister) List(selector labels.Selector) (ret []*v1alpha1.EncryptionByDefault, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EncryptionByDefault))
	})
	return ret, err
}

// EncryptionByDefaults returns an object that can list and get EncryptionByDefaults.
func (s *encryptionByDefaultLister) EncryptionByDefaults(namespace string) EncryptionByDefaultNamespaceLister {
	return encryptionByDefaultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EncryptionByDefaultNamespaceLister helps list and get EncryptionByDefaults.
// All objects returned here must be treated as read-only.
type EncryptionByDefaultNamespaceLister interface {
	// List lists all EncryptionByDefaults in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EncryptionByDefault, err error)
	// Get retrieves the EncryptionByDefault from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EncryptionByDefault, error)
	EncryptionByDefaultNamespaceListerExpansion
}

// encryptionByDefaultNamespaceLister implements the EncryptionByDefaultNamespaceLister
// interface.
type encryptionByDefaultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EncryptionByDefaults in the indexer for a given namespace.
func (s encryptionByDefaultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EncryptionByDefault, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EncryptionByDefault))
	})
	return ret, err
}

// Get retrieves the EncryptionByDefault from the indexer for a given namespace and name.
func (s encryptionByDefaultNamespaceLister) Get(name string) (*v1alpha1.EncryptionByDefault, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("encryptionbydefault"), name)
	}
	return obj.(*v1alpha1.EncryptionByDefault), nil
}
