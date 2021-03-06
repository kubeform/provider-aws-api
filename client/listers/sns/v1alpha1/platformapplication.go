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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/sns/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PlatformApplicationLister helps list PlatformApplications.
// All objects returned here must be treated as read-only.
type PlatformApplicationLister interface {
	// List lists all PlatformApplications in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformApplication, err error)
	// PlatformApplications returns an object that can list and get PlatformApplications.
	PlatformApplications(namespace string) PlatformApplicationNamespaceLister
	PlatformApplicationListerExpansion
}

// platformApplicationLister implements the PlatformApplicationLister interface.
type platformApplicationLister struct {
	indexer cache.Indexer
}

// NewPlatformApplicationLister returns a new PlatformApplicationLister.
func NewPlatformApplicationLister(indexer cache.Indexer) PlatformApplicationLister {
	return &platformApplicationLister{indexer: indexer}
}

// List lists all PlatformApplications in the indexer.
func (s *platformApplicationLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformApplication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformApplication))
	})
	return ret, err
}

// PlatformApplications returns an object that can list and get PlatformApplications.
func (s *platformApplicationLister) PlatformApplications(namespace string) PlatformApplicationNamespaceLister {
	return platformApplicationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PlatformApplicationNamespaceLister helps list and get PlatformApplications.
// All objects returned here must be treated as read-only.
type PlatformApplicationNamespaceLister interface {
	// List lists all PlatformApplications in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlatformApplication, err error)
	// Get retrieves the PlatformApplication from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PlatformApplication, error)
	PlatformApplicationNamespaceListerExpansion
}

// platformApplicationNamespaceLister implements the PlatformApplicationNamespaceLister
// interface.
type platformApplicationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PlatformApplications in the indexer for a given namespace.
func (s platformApplicationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PlatformApplication, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlatformApplication))
	})
	return ret, err
}

// Get retrieves the PlatformApplication from the indexer for a given namespace and name.
func (s platformApplicationNamespaceLister) Get(name string) (*v1alpha1.PlatformApplication, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("platformapplication"), name)
	}
	return obj.(*v1alpha1.PlatformApplication), nil
}
