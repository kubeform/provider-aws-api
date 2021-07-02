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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/servicequotas/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceQuotaLister helps list ServiceQuotas.
// All objects returned here must be treated as read-only.
type ServiceQuotaLister interface {
	// List lists all ServiceQuotas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceQuota, err error)
	// ServiceQuotas returns an object that can list and get ServiceQuotas.
	ServiceQuotas(namespace string) ServiceQuotaNamespaceLister
	ServiceQuotaListerExpansion
}

// serviceQuotaLister implements the ServiceQuotaLister interface.
type serviceQuotaLister struct {
	indexer cache.Indexer
}

// NewServiceQuotaLister returns a new ServiceQuotaLister.
func NewServiceQuotaLister(indexer cache.Indexer) ServiceQuotaLister {
	return &serviceQuotaLister{indexer: indexer}
}

// List lists all ServiceQuotas in the indexer.
func (s *serviceQuotaLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceQuota, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceQuota))
	})
	return ret, err
}

// ServiceQuotas returns an object that can list and get ServiceQuotas.
func (s *serviceQuotaLister) ServiceQuotas(namespace string) ServiceQuotaNamespaceLister {
	return serviceQuotaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceQuotaNamespaceLister helps list and get ServiceQuotas.
// All objects returned here must be treated as read-only.
type ServiceQuotaNamespaceLister interface {
	// List lists all ServiceQuotas in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServiceQuota, err error)
	// Get retrieves the ServiceQuota from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServiceQuota, error)
	ServiceQuotaNamespaceListerExpansion
}

// serviceQuotaNamespaceLister implements the ServiceQuotaNamespaceLister
// interface.
type serviceQuotaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceQuotas in the indexer for a given namespace.
func (s serviceQuotaNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServiceQuota, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServiceQuota))
	})
	return ret, err
}

// Get retrieves the ServiceQuota from the indexer for a given namespace and name.
func (s serviceQuotaNamespaceLister) Get(name string) (*v1alpha1.ServiceQuota, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("servicequota"), name)
	}
	return obj.(*v1alpha1.ServiceQuota), nil
}
