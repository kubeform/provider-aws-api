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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/fsx/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OntapFileSystemLister helps list OntapFileSystems.
// All objects returned here must be treated as read-only.
type OntapFileSystemLister interface {
	// List lists all OntapFileSystems in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OntapFileSystem, err error)
	// OntapFileSystems returns an object that can list and get OntapFileSystems.
	OntapFileSystems(namespace string) OntapFileSystemNamespaceLister
	OntapFileSystemListerExpansion
}

// ontapFileSystemLister implements the OntapFileSystemLister interface.
type ontapFileSystemLister struct {
	indexer cache.Indexer
}

// NewOntapFileSystemLister returns a new OntapFileSystemLister.
func NewOntapFileSystemLister(indexer cache.Indexer) OntapFileSystemLister {
	return &ontapFileSystemLister{indexer: indexer}
}

// List lists all OntapFileSystems in the indexer.
func (s *ontapFileSystemLister) List(selector labels.Selector) (ret []*v1alpha1.OntapFileSystem, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OntapFileSystem))
	})
	return ret, err
}

// OntapFileSystems returns an object that can list and get OntapFileSystems.
func (s *ontapFileSystemLister) OntapFileSystems(namespace string) OntapFileSystemNamespaceLister {
	return ontapFileSystemNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OntapFileSystemNamespaceLister helps list and get OntapFileSystems.
// All objects returned here must be treated as read-only.
type OntapFileSystemNamespaceLister interface {
	// List lists all OntapFileSystems in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OntapFileSystem, err error)
	// Get retrieves the OntapFileSystem from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OntapFileSystem, error)
	OntapFileSystemNamespaceListerExpansion
}

// ontapFileSystemNamespaceLister implements the OntapFileSystemNamespaceLister
// interface.
type ontapFileSystemNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OntapFileSystems in the indexer for a given namespace.
func (s ontapFileSystemNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OntapFileSystem, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OntapFileSystem))
	})
	return ret, err
}

// Get retrieves the OntapFileSystem from the indexer for a given namespace and name.
func (s ontapFileSystemNamespaceLister) Get(name string) (*v1alpha1.OntapFileSystem, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ontapfilesystem"), name)
	}
	return obj.(*v1alpha1.OntapFileSystem), nil
}