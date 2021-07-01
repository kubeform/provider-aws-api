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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/appautoscaling/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScheduledActionLister helps list ScheduledActions.
// All objects returned here must be treated as read-only.
type ScheduledActionLister interface {
	// List lists all ScheduledActions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledAction, err error)
	// ScheduledActions returns an object that can list and get ScheduledActions.
	ScheduledActions(namespace string) ScheduledActionNamespaceLister
	ScheduledActionListerExpansion
}

// scheduledActionLister implements the ScheduledActionLister interface.
type scheduledActionLister struct {
	indexer cache.Indexer
}

// NewScheduledActionLister returns a new ScheduledActionLister.
func NewScheduledActionLister(indexer cache.Indexer) ScheduledActionLister {
	return &scheduledActionLister{indexer: indexer}
}

// List lists all ScheduledActions in the indexer.
func (s *scheduledActionLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledAction, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledAction))
	})
	return ret, err
}

// ScheduledActions returns an object that can list and get ScheduledActions.
func (s *scheduledActionLister) ScheduledActions(namespace string) ScheduledActionNamespaceLister {
	return scheduledActionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScheduledActionNamespaceLister helps list and get ScheduledActions.
// All objects returned here must be treated as read-only.
type ScheduledActionNamespaceLister interface {
	// List lists all ScheduledActions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScheduledAction, err error)
	// Get retrieves the ScheduledAction from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScheduledAction, error)
	ScheduledActionNamespaceListerExpansion
}

// scheduledActionNamespaceLister implements the ScheduledActionNamespaceLister
// interface.
type scheduledActionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScheduledActions in the indexer for a given namespace.
func (s scheduledActionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScheduledAction, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScheduledAction))
	})
	return ret, err
}

// Get retrieves the ScheduledAction from the indexer for a given namespace and name.
func (s scheduledActionNamespaceLister) Get(name string) (*v1alpha1.ScheduledAction, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scheduledaction"), name)
	}
	return obj.(*v1alpha1.ScheduledAction), nil
}
