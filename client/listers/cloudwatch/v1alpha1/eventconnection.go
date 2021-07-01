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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudwatch/v1alpha1"
)

// EventConnectionLister helps list EventConnections.
// All objects returned here must be treated as read-only.
type EventConnectionLister interface {
	// List lists all EventConnections in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventConnection, err error)
	// EventConnections returns an object that can list and get EventConnections.
	EventConnections(namespace string) EventConnectionNamespaceLister
	EventConnectionListerExpansion
}

// eventConnectionLister implements the EventConnectionLister interface.
type eventConnectionLister struct {
	indexer cache.Indexer
}

// NewEventConnectionLister returns a new EventConnectionLister.
func NewEventConnectionLister(indexer cache.Indexer) EventConnectionLister {
	return &eventConnectionLister{indexer: indexer}
}

// List lists all EventConnections in the indexer.
func (s *eventConnectionLister) List(selector labels.Selector) (ret []*v1alpha1.EventConnection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventConnection))
	})
	return ret, err
}

// EventConnections returns an object that can list and get EventConnections.
func (s *eventConnectionLister) EventConnections(namespace string) EventConnectionNamespaceLister {
	return eventConnectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventConnectionNamespaceLister helps list and get EventConnections.
// All objects returned here must be treated as read-only.
type EventConnectionNamespaceLister interface {
	// List lists all EventConnections in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventConnection, err error)
	// Get retrieves the EventConnection from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EventConnection, error)
	EventConnectionNamespaceListerExpansion
}

// eventConnectionNamespaceLister implements the EventConnectionNamespaceLister
// interface.
type eventConnectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventConnections in the indexer for a given namespace.
func (s eventConnectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventConnection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventConnection))
	})
	return ret, err
}

// Get retrieves the EventConnection from the indexer for a given namespace and name.
func (s eventConnectionNamespaceLister) Get(name string) (*v1alpha1.EventConnection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventconnection"), name)
	}
	return obj.(*v1alpha1.EventConnection), nil
}
