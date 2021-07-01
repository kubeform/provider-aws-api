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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/glue/v1alpha1"
)

// DevEndpointLister helps list DevEndpoints.
// All objects returned here must be treated as read-only.
type DevEndpointLister interface {
	// List lists all DevEndpoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DevEndpoint, err error)
	// DevEndpoints returns an object that can list and get DevEndpoints.
	DevEndpoints(namespace string) DevEndpointNamespaceLister
	DevEndpointListerExpansion
}

// devEndpointLister implements the DevEndpointLister interface.
type devEndpointLister struct {
	indexer cache.Indexer
}

// NewDevEndpointLister returns a new DevEndpointLister.
func NewDevEndpointLister(indexer cache.Indexer) DevEndpointLister {
	return &devEndpointLister{indexer: indexer}
}

// List lists all DevEndpoints in the indexer.
func (s *devEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.DevEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DevEndpoint))
	})
	return ret, err
}

// DevEndpoints returns an object that can list and get DevEndpoints.
func (s *devEndpointLister) DevEndpoints(namespace string) DevEndpointNamespaceLister {
	return devEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DevEndpointNamespaceLister helps list and get DevEndpoints.
// All objects returned here must be treated as read-only.
type DevEndpointNamespaceLister interface {
	// List lists all DevEndpoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DevEndpoint, err error)
	// Get retrieves the DevEndpoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DevEndpoint, error)
	DevEndpointNamespaceListerExpansion
}

// devEndpointNamespaceLister implements the DevEndpointNamespaceLister
// interface.
type devEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DevEndpoints in the indexer for a given namespace.
func (s devEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DevEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DevEndpoint))
	})
	return ret, err
}

// Get retrieves the DevEndpoint from the indexer for a given namespace and name.
func (s devEndpointNamespaceLister) Get(name string) (*v1alpha1.DevEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("devendpoint"), name)
	}
	return obj.(*v1alpha1.DevEndpoint), nil
}
