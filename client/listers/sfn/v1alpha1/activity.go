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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/sfn/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ActivityLister helps list Activities.
// All objects returned here must be treated as read-only.
type ActivityLister interface {
	// List lists all Activities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Activity, err error)
	// Activities returns an object that can list and get Activities.
	Activities(namespace string) ActivityNamespaceLister
	ActivityListerExpansion
}

// activityLister implements the ActivityLister interface.
type activityLister struct {
	indexer cache.Indexer
}

// NewActivityLister returns a new ActivityLister.
func NewActivityLister(indexer cache.Indexer) ActivityLister {
	return &activityLister{indexer: indexer}
}

// List lists all Activities in the indexer.
func (s *activityLister) List(selector labels.Selector) (ret []*v1alpha1.Activity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Activity))
	})
	return ret, err
}

// Activities returns an object that can list and get Activities.
func (s *activityLister) Activities(namespace string) ActivityNamespaceLister {
	return activityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ActivityNamespaceLister helps list and get Activities.
// All objects returned here must be treated as read-only.
type ActivityNamespaceLister interface {
	// List lists all Activities in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Activity, err error)
	// Get retrieves the Activity from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Activity, error)
	ActivityNamespaceListerExpansion
}

// activityNamespaceLister implements the ActivityNamespaceLister
// interface.
type activityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Activities in the indexer for a given namespace.
func (s activityNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Activity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Activity))
	})
	return ret, err
}

// Get retrieves the Activity from the indexer for a given namespace and name.
func (s activityNamespaceLister) Get(name string) (*v1alpha1.Activity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("activity"), name)
	}
	return obj.(*v1alpha1.Activity), nil
}
