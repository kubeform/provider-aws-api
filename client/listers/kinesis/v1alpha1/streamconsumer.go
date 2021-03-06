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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/kinesis/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StreamConsumerLister helps list StreamConsumers.
// All objects returned here must be treated as read-only.
type StreamConsumerLister interface {
	// List lists all StreamConsumers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StreamConsumer, err error)
	// StreamConsumers returns an object that can list and get StreamConsumers.
	StreamConsumers(namespace string) StreamConsumerNamespaceLister
	StreamConsumerListerExpansion
}

// streamConsumerLister implements the StreamConsumerLister interface.
type streamConsumerLister struct {
	indexer cache.Indexer
}

// NewStreamConsumerLister returns a new StreamConsumerLister.
func NewStreamConsumerLister(indexer cache.Indexer) StreamConsumerLister {
	return &streamConsumerLister{indexer: indexer}
}

// List lists all StreamConsumers in the indexer.
func (s *streamConsumerLister) List(selector labels.Selector) (ret []*v1alpha1.StreamConsumer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamConsumer))
	})
	return ret, err
}

// StreamConsumers returns an object that can list and get StreamConsumers.
func (s *streamConsumerLister) StreamConsumers(namespace string) StreamConsumerNamespaceLister {
	return streamConsumerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StreamConsumerNamespaceLister helps list and get StreamConsumers.
// All objects returned here must be treated as read-only.
type StreamConsumerNamespaceLister interface {
	// List lists all StreamConsumers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StreamConsumer, err error)
	// Get retrieves the StreamConsumer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StreamConsumer, error)
	StreamConsumerNamespaceListerExpansion
}

// streamConsumerNamespaceLister implements the StreamConsumerNamespaceLister
// interface.
type streamConsumerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StreamConsumers in the indexer for a given namespace.
func (s streamConsumerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StreamConsumer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StreamConsumer))
	})
	return ret, err
}

// Get retrieves the StreamConsumer from the indexer for a given namespace and name.
func (s streamConsumerNamespaceLister) Get(name string) (*v1alpha1.StreamConsumer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("streamconsumer"), name)
	}
	return obj.(*v1alpha1.StreamConsumer), nil
}
