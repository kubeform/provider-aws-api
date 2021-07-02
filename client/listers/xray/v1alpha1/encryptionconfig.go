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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/xray/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EncryptionConfigLister helps list EncryptionConfigs.
// All objects returned here must be treated as read-only.
type EncryptionConfigLister interface {
	// List lists all EncryptionConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EncryptionConfig, err error)
	// EncryptionConfigs returns an object that can list and get EncryptionConfigs.
	EncryptionConfigs(namespace string) EncryptionConfigNamespaceLister
	EncryptionConfigListerExpansion
}

// encryptionConfigLister implements the EncryptionConfigLister interface.
type encryptionConfigLister struct {
	indexer cache.Indexer
}

// NewEncryptionConfigLister returns a new EncryptionConfigLister.
func NewEncryptionConfigLister(indexer cache.Indexer) EncryptionConfigLister {
	return &encryptionConfigLister{indexer: indexer}
}

// List lists all EncryptionConfigs in the indexer.
func (s *encryptionConfigLister) List(selector labels.Selector) (ret []*v1alpha1.EncryptionConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EncryptionConfig))
	})
	return ret, err
}

// EncryptionConfigs returns an object that can list and get EncryptionConfigs.
func (s *encryptionConfigLister) EncryptionConfigs(namespace string) EncryptionConfigNamespaceLister {
	return encryptionConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EncryptionConfigNamespaceLister helps list and get EncryptionConfigs.
// All objects returned here must be treated as read-only.
type EncryptionConfigNamespaceLister interface {
	// List lists all EncryptionConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EncryptionConfig, err error)
	// Get retrieves the EncryptionConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EncryptionConfig, error)
	EncryptionConfigNamespaceListerExpansion
}

// encryptionConfigNamespaceLister implements the EncryptionConfigNamespaceLister
// interface.
type encryptionConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EncryptionConfigs in the indexer for a given namespace.
func (s encryptionConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EncryptionConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EncryptionConfig))
	})
	return ret, err
}

// Get retrieves the EncryptionConfig from the indexer for a given namespace and name.
func (s encryptionConfigNamespaceLister) Get(name string) (*v1alpha1.EncryptionConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("encryptionconfig"), name)
	}
	return obj.(*v1alpha1.EncryptionConfig), nil
}
