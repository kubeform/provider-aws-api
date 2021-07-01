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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ram/v1alpha1"
)

// ResourceShareAccepterLister helps list ResourceShareAccepters.
// All objects returned here must be treated as read-only.
type ResourceShareAccepterLister interface {
	// List lists all ResourceShareAccepters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceShareAccepter, err error)
	// ResourceShareAccepters returns an object that can list and get ResourceShareAccepters.
	ResourceShareAccepters(namespace string) ResourceShareAccepterNamespaceLister
	ResourceShareAccepterListerExpansion
}

// resourceShareAccepterLister implements the ResourceShareAccepterLister interface.
type resourceShareAccepterLister struct {
	indexer cache.Indexer
}

// NewResourceShareAccepterLister returns a new ResourceShareAccepterLister.
func NewResourceShareAccepterLister(indexer cache.Indexer) ResourceShareAccepterLister {
	return &resourceShareAccepterLister{indexer: indexer}
}

// List lists all ResourceShareAccepters in the indexer.
func (s *resourceShareAccepterLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceShareAccepter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceShareAccepter))
	})
	return ret, err
}

// ResourceShareAccepters returns an object that can list and get ResourceShareAccepters.
func (s *resourceShareAccepterLister) ResourceShareAccepters(namespace string) ResourceShareAccepterNamespaceLister {
	return resourceShareAccepterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceShareAccepterNamespaceLister helps list and get ResourceShareAccepters.
// All objects returned here must be treated as read-only.
type ResourceShareAccepterNamespaceLister interface {
	// List lists all ResourceShareAccepters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceShareAccepter, err error)
	// Get retrieves the ResourceShareAccepter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResourceShareAccepter, error)
	ResourceShareAccepterNamespaceListerExpansion
}

// resourceShareAccepterNamespaceLister implements the ResourceShareAccepterNamespaceLister
// interface.
type resourceShareAccepterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceShareAccepters in the indexer for a given namespace.
func (s resourceShareAccepterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceShareAccepter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceShareAccepter))
	})
	return ret, err
}

// Get retrieves the ResourceShareAccepter from the indexer for a given namespace and name.
func (s resourceShareAccepterNamespaceLister) Get(name string) (*v1alpha1.ResourceShareAccepter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourceshareaccepter"), name)
	}
	return obj.(*v1alpha1.ResourceShareAccepter), nil
}
