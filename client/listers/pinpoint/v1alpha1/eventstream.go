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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/pinpoint/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EventStreamLister helps list EventStreams.
// All objects returned here must be treated as read-only.
type EventStreamLister interface {
	// List lists all EventStreams in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventStream, err error)
	// EventStreams returns an object that can list and get EventStreams.
	EventStreams(namespace string) EventStreamNamespaceLister
	EventStreamListerExpansion
}

// eventStreamLister implements the EventStreamLister interface.
type eventStreamLister struct {
	indexer cache.Indexer
}

// NewEventStreamLister returns a new EventStreamLister.
func NewEventStreamLister(indexer cache.Indexer) EventStreamLister {
	return &eventStreamLister{indexer: indexer}
}

// List lists all EventStreams in the indexer.
func (s *eventStreamLister) List(selector labels.Selector) (ret []*v1alpha1.EventStream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventStream))
	})
	return ret, err
}

// EventStreams returns an object that can list and get EventStreams.
func (s *eventStreamLister) EventStreams(namespace string) EventStreamNamespaceLister {
	return eventStreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventStreamNamespaceLister helps list and get EventStreams.
// All objects returned here must be treated as read-only.
type EventStreamNamespaceLister interface {
	// List lists all EventStreams in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventStream, err error)
	// Get retrieves the EventStream from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EventStream, error)
	EventStreamNamespaceListerExpansion
}

// eventStreamNamespaceLister implements the EventStreamNamespaceLister
// interface.
type eventStreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventStreams in the indexer for a given namespace.
func (s eventStreamNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventStream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventStream))
	})
	return ret, err
}

// Get retrieves the EventStream from the indexer for a given namespace and name.
func (s eventStreamNamespaceLister) Get(name string) (*v1alpha1.EventStream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventstream"), name)
	}
	return obj.(*v1alpha1.EventStream), nil
}
