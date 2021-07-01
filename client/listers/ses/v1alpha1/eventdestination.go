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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"
)

// EventDestinationLister helps list EventDestinations.
// All objects returned here must be treated as read-only.
type EventDestinationLister interface {
	// List lists all EventDestinations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventDestination, err error)
	// EventDestinations returns an object that can list and get EventDestinations.
	EventDestinations(namespace string) EventDestinationNamespaceLister
	EventDestinationListerExpansion
}

// eventDestinationLister implements the EventDestinationLister interface.
type eventDestinationLister struct {
	indexer cache.Indexer
}

// NewEventDestinationLister returns a new EventDestinationLister.
func NewEventDestinationLister(indexer cache.Indexer) EventDestinationLister {
	return &eventDestinationLister{indexer: indexer}
}

// List lists all EventDestinations in the indexer.
func (s *eventDestinationLister) List(selector labels.Selector) (ret []*v1alpha1.EventDestination, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventDestination))
	})
	return ret, err
}

// EventDestinations returns an object that can list and get EventDestinations.
func (s *eventDestinationLister) EventDestinations(namespace string) EventDestinationNamespaceLister {
	return eventDestinationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventDestinationNamespaceLister helps list and get EventDestinations.
// All objects returned here must be treated as read-only.
type EventDestinationNamespaceLister interface {
	// List lists all EventDestinations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventDestination, err error)
	// Get retrieves the EventDestination from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EventDestination, error)
	EventDestinationNamespaceListerExpansion
}

// eventDestinationNamespaceLister implements the EventDestinationNamespaceLister
// interface.
type eventDestinationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventDestinations in the indexer for a given namespace.
func (s eventDestinationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventDestination, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventDestination))
	})
	return ret, err
}

// Get retrieves the EventDestination from the indexer for a given namespace and name.
func (s eventDestinationNamespaceLister) Get(name string) (*v1alpha1.EventDestination, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventdestination"), name)
	}
	return obj.(*v1alpha1.EventDestination), nil
}
