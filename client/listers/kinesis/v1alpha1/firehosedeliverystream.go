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

// FirehoseDeliveryStreamLister helps list FirehoseDeliveryStreams.
// All objects returned here must be treated as read-only.
type FirehoseDeliveryStreamLister interface {
	// List lists all FirehoseDeliveryStreams in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FirehoseDeliveryStream, err error)
	// FirehoseDeliveryStreams returns an object that can list and get FirehoseDeliveryStreams.
	FirehoseDeliveryStreams(namespace string) FirehoseDeliveryStreamNamespaceLister
	FirehoseDeliveryStreamListerExpansion
}

// firehoseDeliveryStreamLister implements the FirehoseDeliveryStreamLister interface.
type firehoseDeliveryStreamLister struct {
	indexer cache.Indexer
}

// NewFirehoseDeliveryStreamLister returns a new FirehoseDeliveryStreamLister.
func NewFirehoseDeliveryStreamLister(indexer cache.Indexer) FirehoseDeliveryStreamLister {
	return &firehoseDeliveryStreamLister{indexer: indexer}
}

// List lists all FirehoseDeliveryStreams in the indexer.
func (s *firehoseDeliveryStreamLister) List(selector labels.Selector) (ret []*v1alpha1.FirehoseDeliveryStream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirehoseDeliveryStream))
	})
	return ret, err
}

// FirehoseDeliveryStreams returns an object that can list and get FirehoseDeliveryStreams.
func (s *firehoseDeliveryStreamLister) FirehoseDeliveryStreams(namespace string) FirehoseDeliveryStreamNamespaceLister {
	return firehoseDeliveryStreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FirehoseDeliveryStreamNamespaceLister helps list and get FirehoseDeliveryStreams.
// All objects returned here must be treated as read-only.
type FirehoseDeliveryStreamNamespaceLister interface {
	// List lists all FirehoseDeliveryStreams in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FirehoseDeliveryStream, err error)
	// Get retrieves the FirehoseDeliveryStream from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FirehoseDeliveryStream, error)
	FirehoseDeliveryStreamNamespaceListerExpansion
}

// firehoseDeliveryStreamNamespaceLister implements the FirehoseDeliveryStreamNamespaceLister
// interface.
type firehoseDeliveryStreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FirehoseDeliveryStreams in the indexer for a given namespace.
func (s firehoseDeliveryStreamNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FirehoseDeliveryStream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FirehoseDeliveryStream))
	})
	return ret, err
}

// Get retrieves the FirehoseDeliveryStream from the indexer for a given namespace and name.
func (s firehoseDeliveryStreamNamespaceLister) Get(name string) (*v1alpha1.FirehoseDeliveryStream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("firehosedeliverystream"), name)
	}
	return obj.(*v1alpha1.FirehoseDeliveryStream), nil
}
