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

// TransitGatewayPeeringAttachmentAccepterLister helps list TransitGatewayPeeringAttachmentAccepters.
// All objects returned here must be treated as read-only.
type TransitGatewayPeeringAttachmentAccepterLister interface {
	// List lists all TransitGatewayPeeringAttachmentAccepters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TransitGatewayPeeringAttachmentAccepter, err error)
	// TransitGatewayPeeringAttachmentAccepters returns an object that can list and get TransitGatewayPeeringAttachmentAccepters.
	TransitGatewayPeeringAttachmentAccepters(namespace string) TransitGatewayPeeringAttachmentAccepterNamespaceLister
	TransitGatewayPeeringAttachmentAccepterListerExpansion
}

// transitGatewayPeeringAttachmentAccepterLister implements the TransitGatewayPeeringAttachmentAccepterLister interface.
type transitGatewayPeeringAttachmentAccepterLister struct {
	indexer cache.Indexer
}

// NewTransitGatewayPeeringAttachmentAccepterLister returns a new TransitGatewayPeeringAttachmentAccepterLister.
func NewTransitGatewayPeeringAttachmentAccepterLister(indexer cache.Indexer) TransitGatewayPeeringAttachmentAccepterLister {
	return &transitGatewayPeeringAttachmentAccepterLister{indexer: indexer}
}

// List lists all TransitGatewayPeeringAttachmentAccepters in the indexer.
func (s *transitGatewayPeeringAttachmentAccepterLister) List(selector labels.Selector) (ret []*v1alpha1.TransitGatewayPeeringAttachmentAccepter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TransitGatewayPeeringAttachmentAccepter))
	})
	return ret, err
}

// TransitGatewayPeeringAttachmentAccepters returns an object that can list and get TransitGatewayPeeringAttachmentAccepters.
func (s *transitGatewayPeeringAttachmentAccepterLister) TransitGatewayPeeringAttachmentAccepters(namespace string) TransitGatewayPeeringAttachmentAccepterNamespaceLister {
	return transitGatewayPeeringAttachmentAccepterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TransitGatewayPeeringAttachmentAccepterNamespaceLister helps list and get TransitGatewayPeeringAttachmentAccepters.
// All objects returned here must be treated as read-only.
type TransitGatewayPeeringAttachmentAccepterNamespaceLister interface {
	// List lists all TransitGatewayPeeringAttachmentAccepters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TransitGatewayPeeringAttachmentAccepter, err error)
	// Get retrieves the TransitGatewayPeeringAttachmentAccepter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TransitGatewayPeeringAttachmentAccepter, error)
	TransitGatewayPeeringAttachmentAccepterNamespaceListerExpansion
}

// transitGatewayPeeringAttachmentAccepterNamespaceLister implements the TransitGatewayPeeringAttachmentAccepterNamespaceLister
// interface.
type transitGatewayPeeringAttachmentAccepterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TransitGatewayPeeringAttachmentAccepters in the indexer for a given namespace.
func (s transitGatewayPeeringAttachmentAccepterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TransitGatewayPeeringAttachmentAccepter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TransitGatewayPeeringAttachmentAccepter))
	})
	return ret, err
}

// Get retrieves the TransitGatewayPeeringAttachmentAccepter from the indexer for a given namespace and name.
func (s transitGatewayPeeringAttachmentAccepterNamespaceLister) Get(name string) (*v1alpha1.TransitGatewayPeeringAttachmentAccepter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("transitgatewaypeeringattachmentaccepter"), name)
	}
	return obj.(*v1alpha1.TransitGatewayPeeringAttachmentAccepter), nil
}
