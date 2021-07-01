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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/autoscalingplans/v1alpha1"
)

// ScalingPlanLister helps list ScalingPlans.
// All objects returned here must be treated as read-only.
type ScalingPlanLister interface {
	// List lists all ScalingPlans in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScalingPlan, err error)
	// ScalingPlans returns an object that can list and get ScalingPlans.
	ScalingPlans(namespace string) ScalingPlanNamespaceLister
	ScalingPlanListerExpansion
}

// scalingPlanLister implements the ScalingPlanLister interface.
type scalingPlanLister struct {
	indexer cache.Indexer
}

// NewScalingPlanLister returns a new ScalingPlanLister.
func NewScalingPlanLister(indexer cache.Indexer) ScalingPlanLister {
	return &scalingPlanLister{indexer: indexer}
}

// List lists all ScalingPlans in the indexer.
func (s *scalingPlanLister) List(selector labels.Selector) (ret []*v1alpha1.ScalingPlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScalingPlan))
	})
	return ret, err
}

// ScalingPlans returns an object that can list and get ScalingPlans.
func (s *scalingPlanLister) ScalingPlans(namespace string) ScalingPlanNamespaceLister {
	return scalingPlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScalingPlanNamespaceLister helps list and get ScalingPlans.
// All objects returned here must be treated as read-only.
type ScalingPlanNamespaceLister interface {
	// List lists all ScalingPlans in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScalingPlan, err error)
	// Get retrieves the ScalingPlan from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScalingPlan, error)
	ScalingPlanNamespaceListerExpansion
}

// scalingPlanNamespaceLister implements the ScalingPlanNamespaceLister
// interface.
type scalingPlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScalingPlans in the indexer for a given namespace.
func (s scalingPlanNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScalingPlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScalingPlan))
	})
	return ret, err
}

// Get retrieves the ScalingPlan from the indexer for a given namespace and name.
func (s scalingPlanNamespaceLister) Get(name string) (*v1alpha1.ScalingPlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scalingplan"), name)
	}
	return obj.(*v1alpha1.ScalingPlan), nil
}
