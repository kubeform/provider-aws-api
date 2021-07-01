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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/wafregional/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WebACLAssociationLister helps list WebACLAssociations.
// All objects returned here must be treated as read-only.
type WebACLAssociationLister interface {
	// List lists all WebACLAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebACLAssociation, err error)
	// WebACLAssociations returns an object that can list and get WebACLAssociations.
	WebACLAssociations(namespace string) WebACLAssociationNamespaceLister
	WebACLAssociationListerExpansion
}

// webACLAssociationLister implements the WebACLAssociationLister interface.
type webACLAssociationLister struct {
	indexer cache.Indexer
}

// NewWebACLAssociationLister returns a new WebACLAssociationLister.
func NewWebACLAssociationLister(indexer cache.Indexer) WebACLAssociationLister {
	return &webACLAssociationLister{indexer: indexer}
}

// List lists all WebACLAssociations in the indexer.
func (s *webACLAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.WebACLAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebACLAssociation))
	})
	return ret, err
}

// WebACLAssociations returns an object that can list and get WebACLAssociations.
func (s *webACLAssociationLister) WebACLAssociations(namespace string) WebACLAssociationNamespaceLister {
	return webACLAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebACLAssociationNamespaceLister helps list and get WebACLAssociations.
// All objects returned here must be treated as read-only.
type WebACLAssociationNamespaceLister interface {
	// List lists all WebACLAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WebACLAssociation, err error)
	// Get retrieves the WebACLAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WebACLAssociation, error)
	WebACLAssociationNamespaceListerExpansion
}

// webACLAssociationNamespaceLister implements the WebACLAssociationNamespaceLister
// interface.
type webACLAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebACLAssociations in the indexer for a given namespace.
func (s webACLAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WebACLAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebACLAssociation))
	})
	return ret, err
}

// Get retrieves the WebACLAssociation from the indexer for a given namespace and name.
func (s webACLAssociationNamespaceLister) Get(name string) (*v1alpha1.WebACLAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("webaclassociation"), name)
	}
	return obj.(*v1alpha1.WebACLAssociation), nil
}
