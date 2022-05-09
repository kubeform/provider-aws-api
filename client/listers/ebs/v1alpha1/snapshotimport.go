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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ebs/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SnapshotImportLister helps list SnapshotImports.
// All objects returned here must be treated as read-only.
type SnapshotImportLister interface {
	// List lists all SnapshotImports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SnapshotImport, err error)
	// SnapshotImports returns an object that can list and get SnapshotImports.
	SnapshotImports(namespace string) SnapshotImportNamespaceLister
	SnapshotImportListerExpansion
}

// snapshotImportLister implements the SnapshotImportLister interface.
type snapshotImportLister struct {
	indexer cache.Indexer
}

// NewSnapshotImportLister returns a new SnapshotImportLister.
func NewSnapshotImportLister(indexer cache.Indexer) SnapshotImportLister {
	return &snapshotImportLister{indexer: indexer}
}

// List lists all SnapshotImports in the indexer.
func (s *snapshotImportLister) List(selector labels.Selector) (ret []*v1alpha1.SnapshotImport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SnapshotImport))
	})
	return ret, err
}

// SnapshotImports returns an object that can list and get SnapshotImports.
func (s *snapshotImportLister) SnapshotImports(namespace string) SnapshotImportNamespaceLister {
	return snapshotImportNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SnapshotImportNamespaceLister helps list and get SnapshotImports.
// All objects returned here must be treated as read-only.
type SnapshotImportNamespaceLister interface {
	// List lists all SnapshotImports in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SnapshotImport, err error)
	// Get retrieves the SnapshotImport from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SnapshotImport, error)
	SnapshotImportNamespaceListerExpansion
}

// snapshotImportNamespaceLister implements the SnapshotImportNamespaceLister
// interface.
type snapshotImportNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SnapshotImports in the indexer for a given namespace.
func (s snapshotImportNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SnapshotImport, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SnapshotImport))
	})
	return ret, err
}

// Get retrieves the SnapshotImport from the indexer for a given namespace and name.
func (s snapshotImportNamespaceLister) Get(name string) (*v1alpha1.SnapshotImport, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("snapshotimport"), name)
	}
	return obj.(*v1alpha1.SnapshotImport), nil
}