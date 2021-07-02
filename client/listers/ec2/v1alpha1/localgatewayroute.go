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

// LocalGatewayRouteLister helps list LocalGatewayRoutes.
// All objects returned here must be treated as read-only.
type LocalGatewayRouteLister interface {
	// List lists all LocalGatewayRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LocalGatewayRoute, err error)
	// LocalGatewayRoutes returns an object that can list and get LocalGatewayRoutes.
	LocalGatewayRoutes(namespace string) LocalGatewayRouteNamespaceLister
	LocalGatewayRouteListerExpansion
}

// localGatewayRouteLister implements the LocalGatewayRouteLister interface.
type localGatewayRouteLister struct {
	indexer cache.Indexer
}

// NewLocalGatewayRouteLister returns a new LocalGatewayRouteLister.
func NewLocalGatewayRouteLister(indexer cache.Indexer) LocalGatewayRouteLister {
	return &localGatewayRouteLister{indexer: indexer}
}

// List lists all LocalGatewayRoutes in the indexer.
func (s *localGatewayRouteLister) List(selector labels.Selector) (ret []*v1alpha1.LocalGatewayRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LocalGatewayRoute))
	})
	return ret, err
}

// LocalGatewayRoutes returns an object that can list and get LocalGatewayRoutes.
func (s *localGatewayRouteLister) LocalGatewayRoutes(namespace string) LocalGatewayRouteNamespaceLister {
	return localGatewayRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LocalGatewayRouteNamespaceLister helps list and get LocalGatewayRoutes.
// All objects returned here must be treated as read-only.
type LocalGatewayRouteNamespaceLister interface {
	// List lists all LocalGatewayRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LocalGatewayRoute, err error)
	// Get retrieves the LocalGatewayRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LocalGatewayRoute, error)
	LocalGatewayRouteNamespaceListerExpansion
}

// localGatewayRouteNamespaceLister implements the LocalGatewayRouteNamespaceLister
// interface.
type localGatewayRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LocalGatewayRoutes in the indexer for a given namespace.
func (s localGatewayRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LocalGatewayRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LocalGatewayRoute))
	})
	return ret, err
}

// Get retrieves the LocalGatewayRoute from the indexer for a given namespace and name.
func (s localGatewayRouteNamespaceLister) Get(name string) (*v1alpha1.LocalGatewayRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("localgatewayroute"), name)
	}
	return obj.(*v1alpha1.LocalGatewayRoute), nil
}
