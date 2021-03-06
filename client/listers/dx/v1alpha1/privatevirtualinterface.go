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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dx/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PrivateVirtualInterfaceLister helps list PrivateVirtualInterfaces.
// All objects returned here must be treated as read-only.
type PrivateVirtualInterfaceLister interface {
	// List lists all PrivateVirtualInterfaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateVirtualInterface, err error)
	// PrivateVirtualInterfaces returns an object that can list and get PrivateVirtualInterfaces.
	PrivateVirtualInterfaces(namespace string) PrivateVirtualInterfaceNamespaceLister
	PrivateVirtualInterfaceListerExpansion
}

// privateVirtualInterfaceLister implements the PrivateVirtualInterfaceLister interface.
type privateVirtualInterfaceLister struct {
	indexer cache.Indexer
}

// NewPrivateVirtualInterfaceLister returns a new PrivateVirtualInterfaceLister.
func NewPrivateVirtualInterfaceLister(indexer cache.Indexer) PrivateVirtualInterfaceLister {
	return &privateVirtualInterfaceLister{indexer: indexer}
}

// List lists all PrivateVirtualInterfaces in the indexer.
func (s *privateVirtualInterfaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateVirtualInterface, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateVirtualInterface))
	})
	return ret, err
}

// PrivateVirtualInterfaces returns an object that can list and get PrivateVirtualInterfaces.
func (s *privateVirtualInterfaceLister) PrivateVirtualInterfaces(namespace string) PrivateVirtualInterfaceNamespaceLister {
	return privateVirtualInterfaceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrivateVirtualInterfaceNamespaceLister helps list and get PrivateVirtualInterfaces.
// All objects returned here must be treated as read-only.
type PrivateVirtualInterfaceNamespaceLister interface {
	// List lists all PrivateVirtualInterfaces in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PrivateVirtualInterface, err error)
	// Get retrieves the PrivateVirtualInterface from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PrivateVirtualInterface, error)
	PrivateVirtualInterfaceNamespaceListerExpansion
}

// privateVirtualInterfaceNamespaceLister implements the PrivateVirtualInterfaceNamespaceLister
// interface.
type privateVirtualInterfaceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PrivateVirtualInterfaces in the indexer for a given namespace.
func (s privateVirtualInterfaceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PrivateVirtualInterface, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PrivateVirtualInterface))
	})
	return ret, err
}

// Get retrieves the PrivateVirtualInterface from the indexer for a given namespace and name.
func (s privateVirtualInterfaceNamespaceLister) Get(name string) (*v1alpha1.PrivateVirtualInterface, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("privatevirtualinterface"), name)
	}
	return obj.(*v1alpha1.PrivateVirtualInterface), nil
}
