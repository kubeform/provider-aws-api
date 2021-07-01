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

// EventTargetLister helps list EventTargets.
// All objects returned here must be treated as read-only.
type EventTargetLister interface {
	// List lists all EventTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventTarget, err error)
	// EventTargets returns an object that can list and get EventTargets.
	EventTargets(namespace string) EventTargetNamespaceLister
	EventTargetListerExpansion
}

// eventTargetLister implements the EventTargetLister interface.
type eventTargetLister struct {
	indexer cache.Indexer
}

// NewEventTargetLister returns a new EventTargetLister.
func NewEventTargetLister(indexer cache.Indexer) EventTargetLister {
	return &eventTargetLister{indexer: indexer}
}

// List lists all EventTargets in the indexer.
func (s *eventTargetLister) List(selector labels.Selector) (ret []*v1alpha1.EventTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventTarget))
	})
	return ret, err
}

// EventTargets returns an object that can list and get EventTargets.
func (s *eventTargetLister) EventTargets(namespace string) EventTargetNamespaceLister {
	return eventTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EventTargetNamespaceLister helps list and get EventTargets.
// All objects returned here must be treated as read-only.
type EventTargetNamespaceLister interface {
	// List lists all EventTargets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EventTarget, err error)
	// Get retrieves the EventTarget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EventTarget, error)
	EventTargetNamespaceListerExpansion
}

// eventTargetNamespaceLister implements the EventTargetNamespaceLister
// interface.
type eventTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EventTargets in the indexer for a given namespace.
func (s eventTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EventTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EventTarget))
	})
	return ret, err
}

// Get retrieves the EventTarget from the indexer for a given namespace and name.
func (s eventTargetNamespaceLister) Get(name string) (*v1alpha1.EventTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eventtarget"), name)
	}
	return obj.(*v1alpha1.EventTarget), nil
}
