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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/lb/v1alpha1"
)

// ListenerCertificateLister helps list ListenerCertificates.
// All objects returned here must be treated as read-only.
type ListenerCertificateLister interface {
	// List lists all ListenerCertificates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ListenerCertificate, err error)
	// ListenerCertificates returns an object that can list and get ListenerCertificates.
	ListenerCertificates(namespace string) ListenerCertificateNamespaceLister
	ListenerCertificateListerExpansion
}

// listenerCertificateLister implements the ListenerCertificateLister interface.
type listenerCertificateLister struct {
	indexer cache.Indexer
}

// NewListenerCertificateLister returns a new ListenerCertificateLister.
func NewListenerCertificateLister(indexer cache.Indexer) ListenerCertificateLister {
	return &listenerCertificateLister{indexer: indexer}
}

// List lists all ListenerCertificates in the indexer.
func (s *listenerCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.ListenerCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ListenerCertificate))
	})
	return ret, err
}

// ListenerCertificates returns an object that can list and get ListenerCertificates.
func (s *listenerCertificateLister) ListenerCertificates(namespace string) ListenerCertificateNamespaceLister {
	return listenerCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ListenerCertificateNamespaceLister helps list and get ListenerCertificates.
// All objects returned here must be treated as read-only.
type ListenerCertificateNamespaceLister interface {
	// List lists all ListenerCertificates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ListenerCertificate, err error)
	// Get retrieves the ListenerCertificate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ListenerCertificate, error)
	ListenerCertificateNamespaceListerExpansion
}

// listenerCertificateNamespaceLister implements the ListenerCertificateNamespaceLister
// interface.
type listenerCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ListenerCertificates in the indexer for a given namespace.
func (s listenerCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ListenerCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ListenerCertificate))
	})
	return ret, err
}

// Get retrieves the ListenerCertificate from the indexer for a given namespace and name.
func (s listenerCertificateNamespaceLister) Get(name string) (*v1alpha1.ListenerCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("listenercertificate"), name)
	}
	return obj.(*v1alpha1.ListenerCertificate), nil
}
