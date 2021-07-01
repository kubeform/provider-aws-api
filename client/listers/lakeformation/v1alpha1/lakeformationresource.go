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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/lakeformation/v1alpha1"
)

// LakeformationResourceLister helps list LakeformationResources.
// All objects returned here must be treated as read-only.
type LakeformationResourceLister interface {
	// List lists all LakeformationResources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LakeformationResource, err error)
	// LakeformationResources returns an object that can list and get LakeformationResources.
	LakeformationResources(namespace string) LakeformationResourceNamespaceLister
	LakeformationResourceListerExpansion
}

// lakeformationResourceLister implements the LakeformationResourceLister interface.
type lakeformationResourceLister struct {
	indexer cache.Indexer
}

// NewLakeformationResourceLister returns a new LakeformationResourceLister.
func NewLakeformationResourceLister(indexer cache.Indexer) LakeformationResourceLister {
	return &lakeformationResourceLister{indexer: indexer}
}

// List lists all LakeformationResources in the indexer.
func (s *lakeformationResourceLister) List(selector labels.Selector) (ret []*v1alpha1.LakeformationResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LakeformationResource))
	})
	return ret, err
}

// LakeformationResources returns an object that can list and get LakeformationResources.
func (s *lakeformationResourceLister) LakeformationResources(namespace string) LakeformationResourceNamespaceLister {
	return lakeformationResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LakeformationResourceNamespaceLister helps list and get LakeformationResources.
// All objects returned here must be treated as read-only.
type LakeformationResourceNamespaceLister interface {
	// List lists all LakeformationResources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LakeformationResource, err error)
	// Get retrieves the LakeformationResource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LakeformationResource, error)
	LakeformationResourceNamespaceListerExpansion
}

// lakeformationResourceNamespaceLister implements the LakeformationResourceNamespaceLister
// interface.
type lakeformationResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LakeformationResources in the indexer for a given namespace.
func (s lakeformationResourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LakeformationResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LakeformationResource))
	})
	return ret, err
}

// Get retrieves the LakeformationResource from the indexer for a given namespace and name.
func (s lakeformationResourceNamespaceLister) Get(name string) (*v1alpha1.LakeformationResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("lakeformationresource"), name)
	}
	return obj.(*v1alpha1.LakeformationResource), nil
}
