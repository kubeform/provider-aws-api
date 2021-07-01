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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ssoadmin/v1alpha1"
)

// PermissionSetInlinePolicyLister helps list PermissionSetInlinePolicies.
// All objects returned here must be treated as read-only.
type PermissionSetInlinePolicyLister interface {
	// List lists all PermissionSetInlinePolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PermissionSetInlinePolicy, err error)
	// PermissionSetInlinePolicies returns an object that can list and get PermissionSetInlinePolicies.
	PermissionSetInlinePolicies(namespace string) PermissionSetInlinePolicyNamespaceLister
	PermissionSetInlinePolicyListerExpansion
}

// permissionSetInlinePolicyLister implements the PermissionSetInlinePolicyLister interface.
type permissionSetInlinePolicyLister struct {
	indexer cache.Indexer
}

// NewPermissionSetInlinePolicyLister returns a new PermissionSetInlinePolicyLister.
func NewPermissionSetInlinePolicyLister(indexer cache.Indexer) PermissionSetInlinePolicyLister {
	return &permissionSetInlinePolicyLister{indexer: indexer}
}

// List lists all PermissionSetInlinePolicies in the indexer.
func (s *permissionSetInlinePolicyLister) List(selector labels.Selector) (ret []*v1alpha1.PermissionSetInlinePolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PermissionSetInlinePolicy))
	})
	return ret, err
}

// PermissionSetInlinePolicies returns an object that can list and get PermissionSetInlinePolicies.
func (s *permissionSetInlinePolicyLister) PermissionSetInlinePolicies(namespace string) PermissionSetInlinePolicyNamespaceLister {
	return permissionSetInlinePolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PermissionSetInlinePolicyNamespaceLister helps list and get PermissionSetInlinePolicies.
// All objects returned here must be treated as read-only.
type PermissionSetInlinePolicyNamespaceLister interface {
	// List lists all PermissionSetInlinePolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PermissionSetInlinePolicy, err error)
	// Get retrieves the PermissionSetInlinePolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PermissionSetInlinePolicy, error)
	PermissionSetInlinePolicyNamespaceListerExpansion
}

// permissionSetInlinePolicyNamespaceLister implements the PermissionSetInlinePolicyNamespaceLister
// interface.
type permissionSetInlinePolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PermissionSetInlinePolicies in the indexer for a given namespace.
func (s permissionSetInlinePolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PermissionSetInlinePolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PermissionSetInlinePolicy))
	})
	return ret, err
}

// Get retrieves the PermissionSetInlinePolicy from the indexer for a given namespace and name.
func (s permissionSetInlinePolicyNamespaceLister) Get(name string) (*v1alpha1.PermissionSetInlinePolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("permissionsetinlinepolicy"), name)
	}
	return obj.(*v1alpha1.PermissionSetInlinePolicy), nil
}
