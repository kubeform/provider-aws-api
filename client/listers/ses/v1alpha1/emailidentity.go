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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"
)

// EmailIdentityLister helps list EmailIdentities.
// All objects returned here must be treated as read-only.
type EmailIdentityLister interface {
	// List lists all EmailIdentities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EmailIdentity, err error)
	// EmailIdentities returns an object that can list and get EmailIdentities.
	EmailIdentities(namespace string) EmailIdentityNamespaceLister
	EmailIdentityListerExpansion
}

// emailIdentityLister implements the EmailIdentityLister interface.
type emailIdentityLister struct {
	indexer cache.Indexer
}

// NewEmailIdentityLister returns a new EmailIdentityLister.
func NewEmailIdentityLister(indexer cache.Indexer) EmailIdentityLister {
	return &emailIdentityLister{indexer: indexer}
}

// List lists all EmailIdentities in the indexer.
func (s *emailIdentityLister) List(selector labels.Selector) (ret []*v1alpha1.EmailIdentity, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EmailIdentity))
	})
	return ret, err
}

// EmailIdentities returns an object that can list and get EmailIdentities.
func (s *emailIdentityLister) EmailIdentities(namespace string) EmailIdentityNamespaceLister {
	return emailIdentityNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EmailIdentityNamespaceLister helps list and get EmailIdentities.
// All objects returned here must be treated as read-only.
type EmailIdentityNamespaceLister interface {
	// List lists all EmailIdentities in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EmailIdentity, err error)
	// Get retrieves the EmailIdentity from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EmailIdentity, error)
	EmailIdentityNamespaceListerExpansion
}

// emailIdentityNamespaceLister implements the EmailIdentityNamespaceLister
// interface.
type emailIdentityNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EmailIdentities in the indexer for a given namespace.
func (s emailIdentityNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EmailIdentity, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EmailIdentity))
	})
	return ret, err
}

// Get retrieves the EmailIdentity from the indexer for a given namespace and name.
func (s emailIdentityNamespaceLister) Get(name string) (*v1alpha1.EmailIdentity, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("emailidentity"), name)
	}
	return obj.(*v1alpha1.EmailIdentity), nil
}
