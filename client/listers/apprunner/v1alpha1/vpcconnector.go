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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/apprunner/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VpcConnectorLister helps list VpcConnectors.
// All objects returned here must be treated as read-only.
type VpcConnectorLister interface {
	// List lists all VpcConnectors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpcConnector, err error)
	// VpcConnectors returns an object that can list and get VpcConnectors.
	VpcConnectors(namespace string) VpcConnectorNamespaceLister
	VpcConnectorListerExpansion
}

// vpcConnectorLister implements the VpcConnectorLister interface.
type vpcConnectorLister struct {
	indexer cache.Indexer
}

// NewVpcConnectorLister returns a new VpcConnectorLister.
func NewVpcConnectorLister(indexer cache.Indexer) VpcConnectorLister {
	return &vpcConnectorLister{indexer: indexer}
}

// List lists all VpcConnectors in the indexer.
func (s *vpcConnectorLister) List(selector labels.Selector) (ret []*v1alpha1.VpcConnector, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcConnector))
	})
	return ret, err
}

// VpcConnectors returns an object that can list and get VpcConnectors.
func (s *vpcConnectorLister) VpcConnectors(namespace string) VpcConnectorNamespaceLister {
	return vpcConnectorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcConnectorNamespaceLister helps list and get VpcConnectors.
// All objects returned here must be treated as read-only.
type VpcConnectorNamespaceLister interface {
	// List lists all VpcConnectors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpcConnector, err error)
	// Get retrieves the VpcConnector from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VpcConnector, error)
	VpcConnectorNamespaceListerExpansion
}

// vpcConnectorNamespaceLister implements the VpcConnectorNamespaceLister
// interface.
type vpcConnectorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcConnectors in the indexer for a given namespace.
func (s vpcConnectorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcConnector, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcConnector))
	})
	return ret, err
}

// Get retrieves the VpcConnector from the indexer for a given namespace and name.
func (s vpcConnectorNamespaceLister) Get(name string) (*v1alpha1.VpcConnector, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcconnector"), name)
	}
	return obj.(*v1alpha1.VpcConnector), nil
}
