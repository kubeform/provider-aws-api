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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dms/v1alpha1"
)

// EventSubscriptionLister helps list EventSubscriptions.
// All objects returned here must be treated as read-only.
type EventSubscriptionLister interface {
	// List lists all EventSubscriptions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventSubscription, err error)
	// EventSubscriptions returns an object that can list and get EventSubscriptions.
	EventSubscriptions(namespace string) EventSubscriptionNamespaceLister
	EventSubscriptionListerExpansion
}

// eventSubscriptionLister implements the EventSubscriptionLister interface.
type eventSubscriptionLister struct {
	indexer cache.Indexer
}

// NewEventSubscriptionLister returns a new EventSubscriptionLister.
func NewEventSubscriptionLister(indexer cache.Indexer) EventSubscriptionLister {
	return &eventSubscriptionLister{indexer: indexer}
}

// List lists all EventSubscriptions in the indexer.
func (s *eventSubscriptionLister) List(selector labels.Selector) (ret []*v1alpha1.EventSubscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventSubscription))
	})
	return ret, err
}

// EventSubscriptions returns an object that can list and get EventSubscriptions.
func (s *eventSubscriptionLister) EventSubscriptions(namespace string) EventSubscriptionNamespaceLister {
	return eventSubscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventSubscriptionNamespaceLister helps list and get EventSubscriptions.
// All objects returned here must be treated as read-only.
type EventSubscriptionNamespaceLister interface {
	// List lists all EventSubscriptions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventSubscription, err error)
	// Get retrieves the EventSubscription from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EventSubscription, error)
	EventSubscriptionNamespaceListerExpansion
}

// eventSubscriptionNamespaceLister implements the EventSubscriptionNamespaceLister
// interface.
type eventSubscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventSubscriptions in the indexer for a given namespace.
func (s eventSubscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventSubscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventSubscription))
	})
	return ret, err
}

// Get retrieves the EventSubscription from the indexer for a given namespace and name.
func (s eventSubscriptionNamespaceLister) Get(name string) (*v1alpha1.EventSubscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventsubscription"), name)
	}
	return obj.(*v1alpha1.EventSubscription), nil
}
