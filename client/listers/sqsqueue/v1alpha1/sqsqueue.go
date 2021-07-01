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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/sqsqueue/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SqsQueueLister helps list SqsQueues.
// All objects returned here must be treated as read-only.
type SqsQueueLister interface {
	// List lists all SqsQueues in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SqsQueue, err error)
	// SqsQueues returns an object that can list and get SqsQueues.
	SqsQueues(namespace string) SqsQueueNamespaceLister
	SqsQueueListerExpansion
}

// sqsQueueLister implements the SqsQueueLister interface.
type sqsQueueLister struct {
	indexer cache.Indexer
}

// NewSqsQueueLister returns a new SqsQueueLister.
func NewSqsQueueLister(indexer cache.Indexer) SqsQueueLister {
	return &sqsQueueLister{indexer: indexer}
}

// List lists all SqsQueues in the indexer.
func (s *sqsQueueLister) List(selector labels.Selector) (ret []*v1alpha1.SqsQueue, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqsQueue))
	})
	return ret, err
}

// SqsQueues returns an object that can list and get SqsQueues.
func (s *sqsQueueLister) SqsQueues(namespace string) SqsQueueNamespaceLister {
	return sqsQueueNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SqsQueueNamespaceLister helps list and get SqsQueues.
// All objects returned here must be treated as read-only.
type SqsQueueNamespaceLister interface {
	// List lists all SqsQueues in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SqsQueue, err error)
	// Get retrieves the SqsQueue from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SqsQueue, error)
	SqsQueueNamespaceListerExpansion
}

// sqsQueueNamespaceLister implements the SqsQueueNamespaceLister
// interface.
type sqsQueueNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SqsQueues in the indexer for a given namespace.
func (s sqsQueueNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SqsQueue, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SqsQueue))
	})
	return ret, err
}

// Get retrieves the SqsQueue from the indexer for a given namespace and name.
func (s sqsQueueNamespaceLister) Get(name string) (*v1alpha1.SqsQueue, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sqsqueue"), name)
	}
	return obj.(*v1alpha1.SqsQueue), nil
}
