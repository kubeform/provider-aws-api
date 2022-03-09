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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cognito/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IdentityPoolProviderPrincipalTagLister helps list IdentityPoolProviderPrincipalTags.
// All objects returned here must be treated as read-only.
type IdentityPoolProviderPrincipalTagLister interface {
	// List lists all IdentityPoolProviderPrincipalTags in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IdentityPoolProviderPrincipalTag, err error)
	// IdentityPoolProviderPrincipalTags returns an object that can list and get IdentityPoolProviderPrincipalTags.
	IdentityPoolProviderPrincipalTags(namespace string) IdentityPoolProviderPrincipalTagNamespaceLister
	IdentityPoolProviderPrincipalTagListerExpansion
}

// identityPoolProviderPrincipalTagLister implements the IdentityPoolProviderPrincipalTagLister interface.
type identityPoolProviderPrincipalTagLister struct {
	indexer cache.Indexer
}

// NewIdentityPoolProviderPrincipalTagLister returns a new IdentityPoolProviderPrincipalTagLister.
func NewIdentityPoolProviderPrincipalTagLister(indexer cache.Indexer) IdentityPoolProviderPrincipalTagLister {
	return &identityPoolProviderPrincipalTagLister{indexer: indexer}
}

// List lists all IdentityPoolProviderPrincipalTags in the indexer.
func (s *identityPoolProviderPrincipalTagLister) List(selector labels.Selector) (ret []*v1alpha1.IdentityPoolProviderPrincipalTag, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IdentityPoolProviderPrincipalTag))
	})
	return ret, err
}

// IdentityPoolProviderPrincipalTags returns an object that can list and get IdentityPoolProviderPrincipalTags.
func (s *identityPoolProviderPrincipalTagLister) IdentityPoolProviderPrincipalTags(namespace string) IdentityPoolProviderPrincipalTagNamespaceLister {
	return identityPoolProviderPrincipalTagNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IdentityPoolProviderPrincipalTagNamespaceLister helps list and get IdentityPoolProviderPrincipalTags.
// All objects returned here must be treated as read-only.
type IdentityPoolProviderPrincipalTagNamespaceLister interface {
	// List lists all IdentityPoolProviderPrincipalTags in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IdentityPoolProviderPrincipalTag, err error)
	// Get retrieves the IdentityPoolProviderPrincipalTag from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IdentityPoolProviderPrincipalTag, error)
	IdentityPoolProviderPrincipalTagNamespaceListerExpansion
}

// identityPoolProviderPrincipalTagNamespaceLister implements the IdentityPoolProviderPrincipalTagNamespaceLister
// interface.
type identityPoolProviderPrincipalTagNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IdentityPoolProviderPrincipalTags in the indexer for a given namespace.
func (s identityPoolProviderPrincipalTagNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IdentityPoolProviderPrincipalTag, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IdentityPoolProviderPrincipalTag))
	})
	return ret, err
}

// Get retrieves the IdentityPoolProviderPrincipalTag from the indexer for a given namespace and name.
func (s identityPoolProviderPrincipalTagNamespaceLister) Get(name string) (*v1alpha1.IdentityPoolProviderPrincipalTag, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("identitypoolproviderprincipaltag"), name)
	}
	return obj.(*v1alpha1.IdentityPoolProviderPrincipalTag), nil
}
