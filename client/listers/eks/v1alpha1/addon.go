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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/eks/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AddonLister helps list Addons.
// All objects returned here must be treated as read-only.
type AddonLister interface {
	// List lists all Addons in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Addon, err error)
	// Addons returns an object that can list and get Addons.
	Addons(namespace string) AddonNamespaceLister
	AddonListerExpansion
}

// addonLister implements the AddonLister interface.
type addonLister struct {
	indexer cache.Indexer
}

// NewAddonLister returns a new AddonLister.
func NewAddonLister(indexer cache.Indexer) AddonLister {
	return &addonLister{indexer: indexer}
}

// List lists all Addons in the indexer.
func (s *addonLister) List(selector labels.Selector) (ret []*v1alpha1.Addon, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Addon))
	})
	return ret, err
}

// Addons returns an object that can list and get Addons.
func (s *addonLister) Addons(namespace string) AddonNamespaceLister {
	return addonNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AddonNamespaceLister helps list and get Addons.
// All objects returned here must be treated as read-only.
type AddonNamespaceLister interface {
	// List lists all Addons in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Addon, err error)
	// Get retrieves the Addon from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Addon, error)
	AddonNamespaceListerExpansion
}

// addonNamespaceLister implements the AddonNamespaceLister
// interface.
type addonNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Addons in the indexer for a given namespace.
func (s addonNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Addon, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Addon))
	})
	return ret, err
}

// Get retrieves the Addon from the indexer for a given namespace and name.
func (s addonNamespaceLister) Get(name string) (*v1alpha1.Addon, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("addon"), name)
	}
	return obj.(*v1alpha1.Addon), nil
}
