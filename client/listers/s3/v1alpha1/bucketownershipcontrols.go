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

// BucketOwnershipControlsLister helps list BucketOwnershipControlses.
// All objects returned here must be treated as read-only.
type BucketOwnershipControlsLister interface {
	// List lists all BucketOwnershipControlses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketOwnershipControls, err error)
	// BucketOwnershipControlses returns an object that can list and get BucketOwnershipControlses.
	BucketOwnershipControlses(namespace string) BucketOwnershipControlsNamespaceLister
	BucketOwnershipControlsListerExpansion
}

// bucketOwnershipControlsLister implements the BucketOwnershipControlsLister interface.
type bucketOwnershipControlsLister struct {
	indexer cache.Indexer
}

// NewBucketOwnershipControlsLister returns a new BucketOwnershipControlsLister.
func NewBucketOwnershipControlsLister(indexer cache.Indexer) BucketOwnershipControlsLister {
	return &bucketOwnershipControlsLister{indexer: indexer}
}

// List lists all BucketOwnershipControlses in the indexer.
func (s *bucketOwnershipControlsLister) List(selector labels.Selector) (ret []*v1alpha1.BucketOwnershipControls, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketOwnershipControls))
	})
	return ret, err
}

// BucketOwnershipControlses returns an object that can list and get BucketOwnershipControlses.
func (s *bucketOwnershipControlsLister) BucketOwnershipControlses(namespace string) BucketOwnershipControlsNamespaceLister {
	return bucketOwnershipControlsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BucketOwnershipControlsNamespaceLister helps list and get BucketOwnershipControlses.
// All objects returned here must be treated as read-only.
type BucketOwnershipControlsNamespaceLister interface {
	// List lists all BucketOwnershipControlses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BucketOwnershipControls, err error)
	// Get retrieves the BucketOwnershipControls from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BucketOwnershipControls, error)
	BucketOwnershipControlsNamespaceListerExpansion
}

// bucketOwnershipControlsNamespaceLister implements the BucketOwnershipControlsNamespaceLister
// interface.
type bucketOwnershipControlsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BucketOwnershipControlses in the indexer for a given namespace.
func (s bucketOwnershipControlsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BucketOwnershipControls, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketOwnershipControls))
	})
	return ret, err
}

// Get retrieves the BucketOwnershipControls from the indexer for a given namespace and name.
func (s bucketOwnershipControlsNamespaceLister) Get(name string) (*v1alpha1.BucketOwnershipControls, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucketownershipcontrols"), name)
	}
	return obj.(*v1alpha1.BucketOwnershipControls), nil
}
