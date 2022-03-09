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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecs/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TagLister helps list Tags.
// All objects returned here must be treated as read-only.
type TagLister interface {
	// List lists all Tags in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Tag, err error)
	// Tags returns an object that can list and get Tags.
	Tags(namespace string) TagNamespaceLister
	TagListerExpansion
}

// tagLister implements the TagLister interface.
type tagLister struct {
	indexer cache.Indexer
}

// NewTagLister returns a new TagLister.
func NewTagLister(indexer cache.Indexer) TagLister {
	return &tagLister{indexer: indexer}
}

// List lists all Tags in the indexer.
func (s *tagLister) List(selector labels.Selector) (ret []*v1alpha1.Tag, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Tag))
	})
	return ret, err
}

// Tags returns an object that can list and get Tags.
func (s *tagLister) Tags(namespace string) TagNamespaceLister {
	return tagNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TagNamespaceLister helps list and get Tags.
// All objects returned here must be treated as read-only.
type TagNamespaceLister interface {
	// List lists all Tags in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Tag, err error)
	// Get retrieves the Tag from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Tag, error)
	TagNamespaceListerExpansion
}

// tagNamespaceLister implements the TagNamespaceLister
// interface.
type tagNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Tags in the indexer for a given namespace.
func (s tagNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Tag, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Tag))
	})
	return ret, err
}

// Get retrieves the Tag from the indexer for a given namespace and name.
func (s tagNamespaceLister) Get(name string) (*v1alpha1.Tag, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tag"), name)
	}
	return obj.(*v1alpha1.Tag), nil
}
