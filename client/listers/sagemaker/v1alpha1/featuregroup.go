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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/sagemaker/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FeatureGroupLister helps list FeatureGroups.
// All objects returned here must be treated as read-only.
type FeatureGroupLister interface {
	// List lists all FeatureGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FeatureGroup, err error)
	// FeatureGroups returns an object that can list and get FeatureGroups.
	FeatureGroups(namespace string) FeatureGroupNamespaceLister
	FeatureGroupListerExpansion
}

// featureGroupLister implements the FeatureGroupLister interface.
type featureGroupLister struct {
	indexer cache.Indexer
}

// NewFeatureGroupLister returns a new FeatureGroupLister.
func NewFeatureGroupLister(indexer cache.Indexer) FeatureGroupLister {
	return &featureGroupLister{indexer: indexer}
}

// List lists all FeatureGroups in the indexer.
func (s *featureGroupLister) List(selector labels.Selector) (ret []*v1alpha1.FeatureGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FeatureGroup))
	})
	return ret, err
}

// FeatureGroups returns an object that can list and get FeatureGroups.
func (s *featureGroupLister) FeatureGroups(namespace string) FeatureGroupNamespaceLister {
	return featureGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FeatureGroupNamespaceLister helps list and get FeatureGroups.
// All objects returned here must be treated as read-only.
type FeatureGroupNamespaceLister interface {
	// List lists all FeatureGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FeatureGroup, err error)
	// Get retrieves the FeatureGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FeatureGroup, error)
	FeatureGroupNamespaceListerExpansion
}

// featureGroupNamespaceLister implements the FeatureGroupNamespaceLister
// interface.
type featureGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FeatureGroups in the indexer for a given namespace.
func (s featureGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FeatureGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FeatureGroup))
	})
	return ret, err
}

// Get retrieves the FeatureGroup from the indexer for a given namespace and name.
func (s featureGroupNamespaceLister) Get(name string) (*v1alpha1.FeatureGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("featuregroup"), name)
	}
	return obj.(*v1alpha1.FeatureGroup), nil
}
