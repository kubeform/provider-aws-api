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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"
)

// CarrierGatewayLister helps list CarrierGateways.
// All objects returned here must be treated as read-only.
type CarrierGatewayLister interface {
	// List lists all CarrierGateways in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CarrierGateway, err error)
	// CarrierGateways returns an object that can list and get CarrierGateways.
	CarrierGateways(namespace string) CarrierGatewayNamespaceLister
	CarrierGatewayListerExpansion
}

// carrierGatewayLister implements the CarrierGatewayLister interface.
type carrierGatewayLister struct {
	indexer cache.Indexer
}

// NewCarrierGatewayLister returns a new CarrierGatewayLister.
func NewCarrierGatewayLister(indexer cache.Indexer) CarrierGatewayLister {
	return &carrierGatewayLister{indexer: indexer}
}

// List lists all CarrierGateways in the indexer.
func (s *carrierGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.CarrierGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CarrierGateway))
	})
	return ret, err
}

// CarrierGateways returns an object that can list and get CarrierGateways.
func (s *carrierGatewayLister) CarrierGateways(namespace string) CarrierGatewayNamespaceLister {
	return carrierGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CarrierGatewayNamespaceLister helps list and get CarrierGateways.
// All objects returned here must be treated as read-only.
type CarrierGatewayNamespaceLister interface {
	// List lists all CarrierGateways in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CarrierGateway, err error)
	// Get retrieves the CarrierGateway from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CarrierGateway, error)
	CarrierGatewayNamespaceListerExpansion
}

// carrierGatewayNamespaceLister implements the CarrierGatewayNamespaceLister
// interface.
type carrierGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CarrierGateways in the indexer for a given namespace.
func (s carrierGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CarrierGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CarrierGateway))
	})
	return ret, err
}

// Get retrieves the CarrierGateway from the indexer for a given namespace and name.
func (s carrierGatewayNamespaceLister) Get(name string) (*v1alpha1.CarrierGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("carriergateway"), name)
	}
	return obj.(*v1alpha1.CarrierGateway), nil
}
