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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/apigatewayv2/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IntegrationResponseLister helps list IntegrationResponses.
// All objects returned here must be treated as read-only.
type IntegrationResponseLister interface {
	// List lists all IntegrationResponses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationResponse, err error)
	// IntegrationResponses returns an object that can list and get IntegrationResponses.
	IntegrationResponses(namespace string) IntegrationResponseNamespaceLister
	IntegrationResponseListerExpansion
}

// integrationResponseLister implements the IntegrationResponseLister interface.
type integrationResponseLister struct {
	indexer cache.Indexer
}

// NewIntegrationResponseLister returns a new IntegrationResponseLister.
func NewIntegrationResponseLister(indexer cache.Indexer) IntegrationResponseLister {
	return &integrationResponseLister{indexer: indexer}
}

// List lists all IntegrationResponses in the indexer.
func (s *integrationResponseLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationResponse, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationResponse))
	})
	return ret, err
}

// IntegrationResponses returns an object that can list and get IntegrationResponses.
func (s *integrationResponseLister) IntegrationResponses(namespace string) IntegrationResponseNamespaceLister {
	return integrationResponseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IntegrationResponseNamespaceLister helps list and get IntegrationResponses.
// All objects returned here must be treated as read-only.
type IntegrationResponseNamespaceLister interface {
	// List lists all IntegrationResponses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IntegrationResponse, err error)
	// Get retrieves the IntegrationResponse from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IntegrationResponse, error)
	IntegrationResponseNamespaceListerExpansion
}

// integrationResponseNamespaceLister implements the IntegrationResponseNamespaceLister
// interface.
type integrationResponseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IntegrationResponses in the indexer for a given namespace.
func (s integrationResponseNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IntegrationResponse, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IntegrationResponse))
	})
	return ret, err
}

// Get retrieves the IntegrationResponse from the indexer for a given namespace and name.
func (s integrationResponseNamespaceLister) Get(name string) (*v1alpha1.IntegrationResponse, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("integrationresponse"), name)
	}
	return obj.(*v1alpha1.IntegrationResponse), nil
}
