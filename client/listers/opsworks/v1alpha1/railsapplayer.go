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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/opsworks/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RailsAppLayerLister helps list RailsAppLayers.
// All objects returned here must be treated as read-only.
type RailsAppLayerLister interface {
	// List lists all RailsAppLayers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RailsAppLayer, err error)
	// RailsAppLayers returns an object that can list and get RailsAppLayers.
	RailsAppLayers(namespace string) RailsAppLayerNamespaceLister
	RailsAppLayerListerExpansion
}

// railsAppLayerLister implements the RailsAppLayerLister interface.
type railsAppLayerLister struct {
	indexer cache.Indexer
}

// NewRailsAppLayerLister returns a new RailsAppLayerLister.
func NewRailsAppLayerLister(indexer cache.Indexer) RailsAppLayerLister {
	return &railsAppLayerLister{indexer: indexer}
}

// List lists all RailsAppLayers in the indexer.
func (s *railsAppLayerLister) List(selector labels.Selector) (ret []*v1alpha1.RailsAppLayer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RailsAppLayer))
	})
	return ret, err
}

// RailsAppLayers returns an object that can list and get RailsAppLayers.
func (s *railsAppLayerLister) RailsAppLayers(namespace string) RailsAppLayerNamespaceLister {
	return railsAppLayerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RailsAppLayerNamespaceLister helps list and get RailsAppLayers.
// All objects returned here must be treated as read-only.
type RailsAppLayerNamespaceLister interface {
	// List lists all RailsAppLayers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RailsAppLayer, err error)
	// Get retrieves the RailsAppLayer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RailsAppLayer, error)
	RailsAppLayerNamespaceListerExpansion
}

// railsAppLayerNamespaceLister implements the RailsAppLayerNamespaceLister
// interface.
type railsAppLayerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RailsAppLayers in the indexer for a given namespace.
func (s railsAppLayerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RailsAppLayer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RailsAppLayer))
	})
	return ret, err
}

// Get retrieves the RailsAppLayer from the indexer for a given namespace and name.
func (s railsAppLayerNamespaceLister) Get(name string) (*v1alpha1.RailsAppLayer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("railsapplayer"), name)
	}
	return obj.(*v1alpha1.RailsAppLayer), nil
}
