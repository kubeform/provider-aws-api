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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/datasync/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LocationFsxLustreFileSystemLister helps list LocationFsxLustreFileSystems.
// All objects returned here must be treated as read-only.
type LocationFsxLustreFileSystemLister interface {
	// List lists all LocationFsxLustreFileSystems in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LocationFsxLustreFileSystem, err error)
	// LocationFsxLustreFileSystems returns an object that can list and get LocationFsxLustreFileSystems.
	LocationFsxLustreFileSystems(namespace string) LocationFsxLustreFileSystemNamespaceLister
	LocationFsxLustreFileSystemListerExpansion
}

// locationFsxLustreFileSystemLister implements the LocationFsxLustreFileSystemLister interface.
type locationFsxLustreFileSystemLister struct {
	indexer cache.Indexer
}

// NewLocationFsxLustreFileSystemLister returns a new LocationFsxLustreFileSystemLister.
func NewLocationFsxLustreFileSystemLister(indexer cache.Indexer) LocationFsxLustreFileSystemLister {
	return &locationFsxLustreFileSystemLister{indexer: indexer}
}

// List lists all LocationFsxLustreFileSystems in the indexer.
func (s *locationFsxLustreFileSystemLister) List(selector labels.Selector) (ret []*v1alpha1.LocationFsxLustreFileSystem, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LocationFsxLustreFileSystem))
	})
	return ret, err
}

// LocationFsxLustreFileSystems returns an object that can list and get LocationFsxLustreFileSystems.
func (s *locationFsxLustreFileSystemLister) LocationFsxLustreFileSystems(namespace string) LocationFsxLustreFileSystemNamespaceLister {
	return locationFsxLustreFileSystemNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LocationFsxLustreFileSystemNamespaceLister helps list and get LocationFsxLustreFileSystems.
// All objects returned here must be treated as read-only.
type LocationFsxLustreFileSystemNamespaceLister interface {
	// List lists all LocationFsxLustreFileSystems in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LocationFsxLustreFileSystem, err error)
	// Get retrieves the LocationFsxLustreFileSystem from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LocationFsxLustreFileSystem, error)
	LocationFsxLustreFileSystemNamespaceListerExpansion
}

// locationFsxLustreFileSystemNamespaceLister implements the LocationFsxLustreFileSystemNamespaceLister
// interface.
type locationFsxLustreFileSystemNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LocationFsxLustreFileSystems in the indexer for a given namespace.
func (s locationFsxLustreFileSystemNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LocationFsxLustreFileSystem, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LocationFsxLustreFileSystem))
	})
	return ret, err
}

// Get retrieves the LocationFsxLustreFileSystem from the indexer for a given namespace and name.
func (s locationFsxLustreFileSystemNamespaceLister) Get(name string) (*v1alpha1.LocationFsxLustreFileSystem, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("locationfsxlustrefilesystem"), name)
	}
	return obj.(*v1alpha1.LocationFsxLustreFileSystem), nil
}
