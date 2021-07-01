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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ssm/v1alpha1"
)

// PatchGroupLister helps list PatchGroups.
// All objects returned here must be treated as read-only.
type PatchGroupLister interface {
	// List lists all PatchGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PatchGroup, err error)
	// PatchGroups returns an object that can list and get PatchGroups.
	PatchGroups(namespace string) PatchGroupNamespaceLister
	PatchGroupListerExpansion
}

// patchGroupLister implements the PatchGroupLister interface.
type patchGroupLister struct {
	indexer cache.Indexer
}

// NewPatchGroupLister returns a new PatchGroupLister.
func NewPatchGroupLister(indexer cache.Indexer) PatchGroupLister {
	return &patchGroupLister{indexer: indexer}
}

// List lists all PatchGroups in the indexer.
func (s *patchGroupLister) List(selector labels.Selector) (ret []*v1alpha1.PatchGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PatchGroup))
	})
	return ret, err
}

// PatchGroups returns an object that can list and get PatchGroups.
func (s *patchGroupLister) PatchGroups(namespace string) PatchGroupNamespaceLister {
	return patchGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PatchGroupNamespaceLister helps list and get PatchGroups.
// All objects returned here must be treated as read-only.
type PatchGroupNamespaceLister interface {
	// List lists all PatchGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PatchGroup, err error)
	// Get retrieves the PatchGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PatchGroup, error)
	PatchGroupNamespaceListerExpansion
}

// patchGroupNamespaceLister implements the PatchGroupNamespaceLister
// interface.
type patchGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PatchGroups in the indexer for a given namespace.
func (s patchGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PatchGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PatchGroup))
	})
	return ret, err
}

// Get retrieves the PatchGroup from the indexer for a given namespace and name.
func (s patchGroupNamespaceLister) Get(name string) (*v1alpha1.PatchGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("patchgroup"), name)
	}
	return obj.(*v1alpha1.PatchGroup), nil
}
