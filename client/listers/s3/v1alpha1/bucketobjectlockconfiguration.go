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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BucketObjectLockConfigurationLister helps list BucketObjectLockConfigurations.
// All objects returned here must be treated as read-only.
type BucketObjectLockConfigurationLister interface {
	// List lists all BucketObjectLockConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketObjectLockConfiguration, err error)
	// BucketObjectLockConfigurations returns an object that can list and get BucketObjectLockConfigurations.
	BucketObjectLockConfigurations(namespace string) BucketObjectLockConfigurationNamespaceLister
	BucketObjectLockConfigurationListerExpansion
}

// bucketObjectLockConfigurationLister implements the BucketObjectLockConfigurationLister interface.
type bucketObjectLockConfigurationLister struct {
	indexer cache.Indexer
}

// NewBucketObjectLockConfigurationLister returns a new BucketObjectLockConfigurationLister.
func NewBucketObjectLockConfigurationLister(indexer cache.Indexer) BucketObjectLockConfigurationLister {
	return &bucketObjectLockConfigurationLister{indexer: indexer}
}

// List lists all BucketObjectLockConfigurations in the indexer.
func (s *bucketObjectLockConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.BucketObjectLockConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketObjectLockConfiguration))
	})
	return ret, err
}

// BucketObjectLockConfigurations returns an object that can list and get BucketObjectLockConfigurations.
func (s *bucketObjectLockConfigurationLister) BucketObjectLockConfigurations(namespace string) BucketObjectLockConfigurationNamespaceLister {
	return bucketObjectLockConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BucketObjectLockConfigurationNamespaceLister helps list and get BucketObjectLockConfigurations.
// All objects returned here must be treated as read-only.
type BucketObjectLockConfigurationNamespaceLister interface {
	// List lists all BucketObjectLockConfigurations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketObjectLockConfiguration, err error)
	// Get retrieves the BucketObjectLockConfiguration from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BucketObjectLockConfiguration, error)
	BucketObjectLockConfigurationNamespaceListerExpansion
}

// bucketObjectLockConfigurationNamespaceLister implements the BucketObjectLockConfigurationNamespaceLister
// interface.
type bucketObjectLockConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BucketObjectLockConfigurations in the indexer for a given namespace.
func (s bucketObjectLockConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BucketObjectLockConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketObjectLockConfiguration))
	})
	return ret, err
}

// Get retrieves the BucketObjectLockConfiguration from the indexer for a given namespace and name.
func (s bucketObjectLockConfigurationNamespaceLister) Get(name string) (*v1alpha1.BucketObjectLockConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucketobjectlockconfiguration"), name)
	}
	return obj.(*v1alpha1.BucketObjectLockConfiguration), nil
}
