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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/connect/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ContactFlowModuleLister helps list ContactFlowModules.
// All objects returned here must be treated as read-only.
type ContactFlowModuleLister interface {
	// List lists all ContactFlowModules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ContactFlowModule, err error)
	// ContactFlowModules returns an object that can list and get ContactFlowModules.
	ContactFlowModules(namespace string) ContactFlowModuleNamespaceLister
	ContactFlowModuleListerExpansion
}

// contactFlowModuleLister implements the ContactFlowModuleLister interface.
type contactFlowModuleLister struct {
	indexer cache.Indexer
}

// NewContactFlowModuleLister returns a new ContactFlowModuleLister.
func NewContactFlowModuleLister(indexer cache.Indexer) ContactFlowModuleLister {
	return &contactFlowModuleLister{indexer: indexer}
}

// List lists all ContactFlowModules in the indexer.
func (s *contactFlowModuleLister) List(selector labels.Selector) (ret []*v1alpha1.ContactFlowModule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContactFlowModule))
	})
	return ret, err
}

// ContactFlowModules returns an object that can list and get ContactFlowModules.
func (s *contactFlowModuleLister) ContactFlowModules(namespace string) ContactFlowModuleNamespaceLister {
	return contactFlowModuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ContactFlowModuleNamespaceLister helps list and get ContactFlowModules.
// All objects returned here must be treated as read-only.
type ContactFlowModuleNamespaceLister interface {
	// List lists all ContactFlowModules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ContactFlowModule, err error)
	// Get retrieves the ContactFlowModule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ContactFlowModule, error)
	ContactFlowModuleNamespaceListerExpansion
}

// contactFlowModuleNamespaceLister implements the ContactFlowModuleNamespaceLister
// interface.
type contactFlowModuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ContactFlowModules in the indexer for a given namespace.
func (s contactFlowModuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ContactFlowModule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ContactFlowModule))
	})
	return ret, err
}

// Get retrieves the ContactFlowModule from the indexer for a given namespace and name.
func (s contactFlowModuleNamespaceLister) Get(name string) (*v1alpha1.ContactFlowModule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("contactflowmodule"), name)
	}
	return obj.(*v1alpha1.ContactFlowModule), nil
}