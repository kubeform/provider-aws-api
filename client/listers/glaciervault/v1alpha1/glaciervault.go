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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/glaciervault/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GlacierVaultLister helps list GlacierVaults.
// All objects returned here must be treated as read-only.
type GlacierVaultLister interface {
	// List lists all GlacierVaults in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GlacierVault, err error)
	// GlacierVaults returns an object that can list and get GlacierVaults.
	GlacierVaults(namespace string) GlacierVaultNamespaceLister
	GlacierVaultListerExpansion
}

// glacierVaultLister implements the GlacierVaultLister interface.
type glacierVaultLister struct {
	indexer cache.Indexer
}

// NewGlacierVaultLister returns a new GlacierVaultLister.
func NewGlacierVaultLister(indexer cache.Indexer) GlacierVaultLister {
	return &glacierVaultLister{indexer: indexer}
}

// List lists all GlacierVaults in the indexer.
func (s *glacierVaultLister) List(selector labels.Selector) (ret []*v1alpha1.GlacierVault, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlacierVault))
	})
	return ret, err
}

// GlacierVaults returns an object that can list and get GlacierVaults.
func (s *glacierVaultLister) GlacierVaults(namespace string) GlacierVaultNamespaceLister {
	return glacierVaultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlacierVaultNamespaceLister helps list and get GlacierVaults.
// All objects returned here must be treated as read-only.
type GlacierVaultNamespaceLister interface {
	// List lists all GlacierVaults in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GlacierVault, err error)
	// Get retrieves the GlacierVault from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GlacierVault, error)
	GlacierVaultNamespaceListerExpansion
}

// glacierVaultNamespaceLister implements the GlacierVaultNamespaceLister
// interface.
type glacierVaultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlacierVaults in the indexer for a given namespace.
func (s glacierVaultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GlacierVault, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlacierVault))
	})
	return ret, err
}

// Get retrieves the GlacierVault from the indexer for a given namespace and name.
func (s glacierVaultNamespaceLister) Get(name string) (*v1alpha1.GlacierVault, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("glaciervault"), name)
	}
	return obj.(*v1alpha1.GlacierVault), nil
}
