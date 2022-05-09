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

// BucketVersioningLister helps list BucketVersionings.
// All objects returned here must be treated as read-only.
type BucketVersioningLister interface {
	// List lists all BucketVersionings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketVersioning, err error)
	// BucketVersionings returns an object that can list and get BucketVersionings.
	BucketVersionings(namespace string) BucketVersioningNamespaceLister
	BucketVersioningListerExpansion
}

// bucketVersioningLister implements the BucketVersioningLister interface.
type bucketVersioningLister struct {
	indexer cache.Indexer
}

// NewBucketVersioningLister returns a new BucketVersioningLister.
func NewBucketVersioningLister(indexer cache.Indexer) BucketVersioningLister {
	return &bucketVersioningLister{indexer: indexer}
}

// List lists all BucketVersionings in the indexer.
func (s *bucketVersioningLister) List(selector labels.Selector) (ret []*v1alpha1.BucketVersioning, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketVersioning))
	})
	return ret, err
}

// BucketVersionings returns an object that can list and get BucketVersionings.
func (s *bucketVersioningLister) BucketVersionings(namespace string) BucketVersioningNamespaceLister {
	return bucketVersioningNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BucketVersioningNamespaceLister helps list and get BucketVersionings.
// All objects returned here must be treated as read-only.
type BucketVersioningNamespaceLister interface {
	// List lists all BucketVersionings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketVersioning, err error)
	// Get retrieves the BucketVersioning from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BucketVersioning, error)
	BucketVersioningNamespaceListerExpansion
}

// bucketVersioningNamespaceLister implements the BucketVersioningNamespaceLister
// interface.
type bucketVersioningNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BucketVersionings in the indexer for a given namespace.
func (s bucketVersioningNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BucketVersioning, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketVersioning))
	})
	return ret, err
}

// Get retrieves the BucketVersioning from the indexer for a given namespace and name.
func (s bucketVersioningNamespaceLister) Get(name string) (*v1alpha1.BucketVersioning, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucketversioning"), name)
	}
	return obj.(*v1alpha1.BucketVersioning), nil
}