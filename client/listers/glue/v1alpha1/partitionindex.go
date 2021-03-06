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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/glue/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PartitionIndexLister helps list PartitionIndexes.
// All objects returned here must be treated as read-only.
type PartitionIndexLister interface {
	// List lists all PartitionIndexes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PartitionIndex, err error)
	// PartitionIndexes returns an object that can list and get PartitionIndexes.
	PartitionIndexes(namespace string) PartitionIndexNamespaceLister
	PartitionIndexListerExpansion
}

// partitionIndexLister implements the PartitionIndexLister interface.
type partitionIndexLister struct {
	indexer cache.Indexer
}

// NewPartitionIndexLister returns a new PartitionIndexLister.
func NewPartitionIndexLister(indexer cache.Indexer) PartitionIndexLister {
	return &partitionIndexLister{indexer: indexer}
}

// List lists all PartitionIndexes in the indexer.
func (s *partitionIndexLister) List(selector labels.Selector) (ret []*v1alpha1.PartitionIndex, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PartitionIndex))
	})
	return ret, err
}

// PartitionIndexes returns an object that can list and get PartitionIndexes.
func (s *partitionIndexLister) PartitionIndexes(namespace string) PartitionIndexNamespaceLister {
	return partitionIndexNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PartitionIndexNamespaceLister helps list and get PartitionIndexes.
// All objects returned here must be treated as read-only.
type PartitionIndexNamespaceLister interface {
	// List lists all PartitionIndexes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PartitionIndex, err error)
	// Get retrieves the PartitionIndex from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PartitionIndex, error)
	PartitionIndexNamespaceListerExpansion
}

// partitionIndexNamespaceLister implements the PartitionIndexNamespaceLister
// interface.
type partitionIndexNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PartitionIndexes in the indexer for a given namespace.
func (s partitionIndexNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PartitionIndex, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PartitionIndex))
	})
	return ret, err
}

// Get retrieves the PartitionIndex from the indexer for a given namespace and name.
func (s partitionIndexNamespaceLister) Get(name string) (*v1alpha1.PartitionIndex, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("partitionindex"), name)
	}
	return obj.(*v1alpha1.PartitionIndex), nil
}
