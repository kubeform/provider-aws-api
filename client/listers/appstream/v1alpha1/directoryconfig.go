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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/appstream/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DirectoryConfigLister helps list DirectoryConfigs.
// All objects returned here must be treated as read-only.
type DirectoryConfigLister interface {
	// List lists all DirectoryConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DirectoryConfig, err error)
	// DirectoryConfigs returns an object that can list and get DirectoryConfigs.
	DirectoryConfigs(namespace string) DirectoryConfigNamespaceLister
	DirectoryConfigListerExpansion
}

// directoryConfigLister implements the DirectoryConfigLister interface.
type directoryConfigLister struct {
	indexer cache.Indexer
}

// NewDirectoryConfigLister returns a new DirectoryConfigLister.
func NewDirectoryConfigLister(indexer cache.Indexer) DirectoryConfigLister {
	return &directoryConfigLister{indexer: indexer}
}

// List lists all DirectoryConfigs in the indexer.
func (s *directoryConfigLister) List(selector labels.Selector) (ret []*v1alpha1.DirectoryConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DirectoryConfig))
	})
	return ret, err
}

// DirectoryConfigs returns an object that can list and get DirectoryConfigs.
func (s *directoryConfigLister) DirectoryConfigs(namespace string) DirectoryConfigNamespaceLister {
	return directoryConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DirectoryConfigNamespaceLister helps list and get DirectoryConfigs.
// All objects returned here must be treated as read-only.
type DirectoryConfigNamespaceLister interface {
	// List lists all DirectoryConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DirectoryConfig, err error)
	// Get retrieves the DirectoryConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DirectoryConfig, error)
	DirectoryConfigNamespaceListerExpansion
}

// directoryConfigNamespaceLister implements the DirectoryConfigNamespaceLister
// interface.
type directoryConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DirectoryConfigs in the indexer for a given namespace.
func (s directoryConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DirectoryConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DirectoryConfig))
	})
	return ret, err
}

// Get retrieves the DirectoryConfig from the indexer for a given namespace and name.
func (s directoryConfigNamespaceLister) Get(name string) (*v1alpha1.DirectoryConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("directoryconfig"), name)
	}
	return obj.(*v1alpha1.DirectoryConfig), nil
}
