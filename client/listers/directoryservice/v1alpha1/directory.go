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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/directoryservice/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DirectoryLister helps list Directories.
// All objects returned here must be treated as read-only.
type DirectoryLister interface {
	// List lists all Directories in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Directory, err error)
	// Directories returns an object that can list and get Directories.
	Directories(namespace string) DirectoryNamespaceLister
	DirectoryListerExpansion
}

// directoryLister implements the DirectoryLister interface.
type directoryLister struct {
	indexer cache.Indexer
}

// NewDirectoryLister returns a new DirectoryLister.
func NewDirectoryLister(indexer cache.Indexer) DirectoryLister {
	return &directoryLister{indexer: indexer}
}

// List lists all Directories in the indexer.
func (s *directoryLister) List(selector labels.Selector) (ret []*v1alpha1.Directory, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Directory))
	})
	return ret, err
}

// Directories returns an object that can list and get Directories.
func (s *directoryLister) Directories(namespace string) DirectoryNamespaceLister {
	return directoryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DirectoryNamespaceLister helps list and get Directories.
// All objects returned here must be treated as read-only.
type DirectoryNamespaceLister interface {
	// List lists all Directories in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Directory, err error)
	// Get retrieves the Directory from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Directory, error)
	DirectoryNamespaceListerExpansion
}

// directoryNamespaceLister implements the DirectoryNamespaceLister
// interface.
type directoryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Directories in the indexer for a given namespace.
func (s directoryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Directory, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Directory))
	})
	return ret, err
}

// Get retrieves the Directory from the indexer for a given namespace and name.
func (s directoryNamespaceLister) Get(name string) (*v1alpha1.Directory, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("directory"), name)
	}
	return obj.(*v1alpha1.Directory), nil
}
