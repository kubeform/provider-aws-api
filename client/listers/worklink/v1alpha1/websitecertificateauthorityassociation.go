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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/worklink/v1alpha1"
)

// WebsiteCertificateAuthorityAssociationLister helps list WebsiteCertificateAuthorityAssociations.
// All objects returned here must be treated as read-only.
type WebsiteCertificateAuthorityAssociationLister interface {
	// List lists all WebsiteCertificateAuthorityAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebsiteCertificateAuthorityAssociation, err error)
	// WebsiteCertificateAuthorityAssociations returns an object that can list and get WebsiteCertificateAuthorityAssociations.
	WebsiteCertificateAuthorityAssociations(namespace string) WebsiteCertificateAuthorityAssociationNamespaceLister
	WebsiteCertificateAuthorityAssociationListerExpansion
}

// websiteCertificateAuthorityAssociationLister implements the WebsiteCertificateAuthorityAssociationLister interface.
type websiteCertificateAuthorityAssociationLister struct {
	indexer cache.Indexer
}

// NewWebsiteCertificateAuthorityAssociationLister returns a new WebsiteCertificateAuthorityAssociationLister.
func NewWebsiteCertificateAuthorityAssociationLister(indexer cache.Indexer) WebsiteCertificateAuthorityAssociationLister {
	return &websiteCertificateAuthorityAssociationLister{indexer: indexer}
}

// List lists all WebsiteCertificateAuthorityAssociations in the indexer.
func (s *websiteCertificateAuthorityAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.WebsiteCertificateAuthorityAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebsiteCertificateAuthorityAssociation))
	})
	return ret, err
}

// WebsiteCertificateAuthorityAssociations returns an object that can list and get WebsiteCertificateAuthorityAssociations.
func (s *websiteCertificateAuthorityAssociationLister) WebsiteCertificateAuthorityAssociations(namespace string) WebsiteCertificateAuthorityAssociationNamespaceLister {
	return websiteCertificateAuthorityAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebsiteCertificateAuthorityAssociationNamespaceLister helps list and get WebsiteCertificateAuthorityAssociations.
// All objects returned here must be treated as read-only.
type WebsiteCertificateAuthorityAssociationNamespaceLister interface {
	// List lists all WebsiteCertificateAuthorityAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebsiteCertificateAuthorityAssociation, err error)
	// Get retrieves the WebsiteCertificateAuthorityAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WebsiteCertificateAuthorityAssociation, error)
	WebsiteCertificateAuthorityAssociationNamespaceListerExpansion
}

// websiteCertificateAuthorityAssociationNamespaceLister implements the WebsiteCertificateAuthorityAssociationNamespaceLister
// interface.
type websiteCertificateAuthorityAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebsiteCertificateAuthorityAssociations in the indexer for a given namespace.
func (s websiteCertificateAuthorityAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WebsiteCertificateAuthorityAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebsiteCertificateAuthorityAssociation))
	})
	return ret, err
}

// Get retrieves the WebsiteCertificateAuthorityAssociation from the indexer for a given namespace and name.
func (s websiteCertificateAuthorityAssociationNamespaceLister) Get(name string) (*v1alpha1.WebsiteCertificateAuthorityAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("websitecertificateauthorityassociation"), name)
	}
	return obj.(*v1alpha1.WebsiteCertificateAuthorityAssociation), nil
}
