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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClientVPNRouteLister helps list ClientVPNRoutes.
// All objects returned here must be treated as read-only.
type ClientVPNRouteLister interface {
	// List lists all ClientVPNRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClientVPNRoute, err error)
	// ClientVPNRoutes returns an object that can list and get ClientVPNRoutes.
	ClientVPNRoutes(namespace string) ClientVPNRouteNamespaceLister
	ClientVPNRouteListerExpansion
}

// clientVPNRouteLister implements the ClientVPNRouteLister interface.
type clientVPNRouteLister struct {
	indexer cache.Indexer
}

// NewClientVPNRouteLister returns a new ClientVPNRouteLister.
func NewClientVPNRouteLister(indexer cache.Indexer) ClientVPNRouteLister {
	return &clientVPNRouteLister{indexer: indexer}
}

// List lists all ClientVPNRoutes in the indexer.
func (s *clientVPNRouteLister) List(selector labels.Selector) (ret []*v1alpha1.ClientVPNRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClientVPNRoute))
	})
	return ret, err
}

// ClientVPNRoutes returns an object that can list and get ClientVPNRoutes.
func (s *clientVPNRouteLister) ClientVPNRoutes(namespace string) ClientVPNRouteNamespaceLister {
	return clientVPNRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClientVPNRouteNamespaceLister helps list and get ClientVPNRoutes.
// All objects returned here must be treated as read-only.
type ClientVPNRouteNamespaceLister interface {
	// List lists all ClientVPNRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClientVPNRoute, err error)
	// Get retrieves the ClientVPNRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClientVPNRoute, error)
	ClientVPNRouteNamespaceListerExpansion
}

// clientVPNRouteNamespaceLister implements the ClientVPNRouteNamespaceLister
// interface.
type clientVPNRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClientVPNRoutes in the indexer for a given namespace.
func (s clientVPNRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClientVPNRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClientVPNRoute))
	})
	return ret, err
}

// Get retrieves the ClientVPNRoute from the indexer for a given namespace and name.
func (s clientVPNRouteNamespaceLister) Get(name string) (*v1alpha1.ClientVPNRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("clientvpnroute"), name)
	}
	return obj.(*v1alpha1.ClientVPNRoute), nil
}
