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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudfront/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FieldLevelEncryptionConfigLister helps list FieldLevelEncryptionConfigs.
// All objects returned here must be treated as read-only.
type FieldLevelEncryptionConfigLister interface {
	// List lists all FieldLevelEncryptionConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FieldLevelEncryptionConfig, err error)
	// FieldLevelEncryptionConfigs returns an object that can list and get FieldLevelEncryptionConfigs.
	FieldLevelEncryptionConfigs(namespace string) FieldLevelEncryptionConfigNamespaceLister
	FieldLevelEncryptionConfigListerExpansion
}

// fieldLevelEncryptionConfigLister implements the FieldLevelEncryptionConfigLister interface.
type fieldLevelEncryptionConfigLister struct {
	indexer cache.Indexer
}

// NewFieldLevelEncryptionConfigLister returns a new FieldLevelEncryptionConfigLister.
func NewFieldLevelEncryptionConfigLister(indexer cache.Indexer) FieldLevelEncryptionConfigLister {
	return &fieldLevelEncryptionConfigLister{indexer: indexer}
}

// List lists all FieldLevelEncryptionConfigs in the indexer.
func (s *fieldLevelEncryptionConfigLister) List(selector labels.Selector) (ret []*v1alpha1.FieldLevelEncryptionConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FieldLevelEncryptionConfig))
	})
	return ret, err
}

// FieldLevelEncryptionConfigs returns an object that can list and get FieldLevelEncryptionConfigs.
func (s *fieldLevelEncryptionConfigLister) FieldLevelEncryptionConfigs(namespace string) FieldLevelEncryptionConfigNamespaceLister {
	return fieldLevelEncryptionConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FieldLevelEncryptionConfigNamespaceLister helps list and get FieldLevelEncryptionConfigs.
// All objects returned here must be treated as read-only.
type FieldLevelEncryptionConfigNamespaceLister interface {
	// List lists all FieldLevelEncryptionConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FieldLevelEncryptionConfig, err error)
	// Get retrieves the FieldLevelEncryptionConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FieldLevelEncryptionConfig, error)
	FieldLevelEncryptionConfigNamespaceListerExpansion
}

// fieldLevelEncryptionConfigNamespaceLister implements the FieldLevelEncryptionConfigNamespaceLister
// interface.
type fieldLevelEncryptionConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FieldLevelEncryptionConfigs in the indexer for a given namespace.
func (s fieldLevelEncryptionConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FieldLevelEncryptionConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FieldLevelEncryptionConfig))
	})
	return ret, err
}

// Get retrieves the FieldLevelEncryptionConfig from the indexer for a given namespace and name.
func (s fieldLevelEncryptionConfigNamespaceLister) Get(name string) (*v1alpha1.FieldLevelEncryptionConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fieldlevelencryptionconfig"), name)
	}
	return obj.(*v1alpha1.FieldLevelEncryptionConfig), nil
}
