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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/backup/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PlanLister helps list Plans.
// All objects returned here must be treated as read-only.
type PlanLister interface {
	// List lists all Plans in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Plan, err error)
	// Plans returns an object that can list and get Plans.
	Plans(namespace string) PlanNamespaceLister
	PlanListerExpansion
}

// planLister implements the PlanLister interface.
type planLister struct {
	indexer cache.Indexer
}

// NewPlanLister returns a new PlanLister.
func NewPlanLister(indexer cache.Indexer) PlanLister {
	return &planLister{indexer: indexer}
}

// List lists all Plans in the indexer.
func (s *planLister) List(selector labels.Selector) (ret []*v1alpha1.Plan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Plan))
	})
	return ret, err
}

// Plans returns an object that can list and get Plans.
func (s *planLister) Plans(namespace string) PlanNamespaceLister {
	return planNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PlanNamespaceLister helps list and get Plans.
// All objects returned here must be treated as read-only.
type PlanNamespaceLister interface {
	// List lists all Plans in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Plan, err error)
	// Get retrieves the Plan from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Plan, error)
	PlanNamespaceListerExpansion
}

// planNamespaceLister implements the PlanNamespaceLister
// interface.
type planNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Plans in the indexer for a given namespace.
func (s planNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Plan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Plan))
	})
	return ret, err
}

// Get retrieves the Plan from the indexer for a given namespace and name.
func (s planNamespaceLister) Get(name string) (*v1alpha1.Plan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("plan"), name)
	}
	return obj.(*v1alpha1.Plan), nil
}
