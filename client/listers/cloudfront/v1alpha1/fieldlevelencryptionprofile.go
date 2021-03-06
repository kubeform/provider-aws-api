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

// FieldLevelEncryptionProfileLister helps list FieldLevelEncryptionProfiles.
// All objects returned here must be treated as read-only.
type FieldLevelEncryptionProfileLister interface {
	// List lists all FieldLevelEncryptionProfiles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FieldLevelEncryptionProfile, err error)
	// FieldLevelEncryptionProfiles returns an object that can list and get FieldLevelEncryptionProfiles.
	FieldLevelEncryptionProfiles(namespace string) FieldLevelEncryptionProfileNamespaceLister
	FieldLevelEncryptionProfileListerExpansion
}

// fieldLevelEncryptionProfileLister implements the FieldLevelEncryptionProfileLister interface.
type fieldLevelEncryptionProfileLister struct {
	indexer cache.Indexer
}

// NewFieldLevelEncryptionProfileLister returns a new FieldLevelEncryptionProfileLister.
func NewFieldLevelEncryptionProfileLister(indexer cache.Indexer) FieldLevelEncryptionProfileLister {
	return &fieldLevelEncryptionProfileLister{indexer: indexer}
}

// List lists all FieldLevelEncryptionProfiles in the indexer.
func (s *fieldLevelEncryptionProfileLister) List(selector labels.Selector) (ret []*v1alpha1.FieldLevelEncryptionProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FieldLevelEncryptionProfile))
	})
	return ret, err
}

// FieldLevelEncryptionProfiles returns an object that can list and get FieldLevelEncryptionProfiles.
func (s *fieldLevelEncryptionProfileLister) FieldLevelEncryptionProfiles(namespace string) FieldLevelEncryptionProfileNamespaceLister {
	return fieldLevelEncryptionProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FieldLevelEncryptionProfileNamespaceLister helps list and get FieldLevelEncryptionProfiles.
// All objects returned here must be treated as read-only.
type FieldLevelEncryptionProfileNamespaceLister interface {
	// List lists all FieldLevelEncryptionProfiles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FieldLevelEncryptionProfile, err error)
	// Get retrieves the FieldLevelEncryptionProfile from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FieldLevelEncryptionProfile, error)
	FieldLevelEncryptionProfileNamespaceListerExpansion
}

// fieldLevelEncryptionProfileNamespaceLister implements the FieldLevelEncryptionProfileNamespaceLister
// interface.
type fieldLevelEncryptionProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FieldLevelEncryptionProfiles in the indexer for a given namespace.
func (s fieldLevelEncryptionProfileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FieldLevelEncryptionProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FieldLevelEncryptionProfile))
	})
	return ret, err
}

// Get retrieves the FieldLevelEncryptionProfile from the indexer for a given namespace and name.
func (s fieldLevelEncryptionProfileNamespaceLister) Get(name string) (*v1alpha1.FieldLevelEncryptionProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("fieldlevelencryptionprofile"), name)
	}
	return obj.(*v1alpha1.FieldLevelEncryptionProfile), nil
}
