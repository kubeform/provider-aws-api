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

// AccountAssignmentLister helps list AccountAssignments.
// All objects returned here must be treated as read-only.
type AccountAssignmentLister interface {
	// List lists all AccountAssignments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AccountAssignment, err error)
	// AccountAssignments returns an object that can list and get AccountAssignments.
	AccountAssignments(namespace string) AccountAssignmentNamespaceLister
	AccountAssignmentListerExpansion
}

// accountAssignmentLister implements the AccountAssignmentLister interface.
type accountAssignmentLister struct {
	indexer cache.Indexer
}

// NewAccountAssignmentLister returns a new AccountAssignmentLister.
func NewAccountAssignmentLister(indexer cache.Indexer) AccountAssignmentLister {
	return &accountAssignmentLister{indexer: indexer}
}

// List lists all AccountAssignments in the indexer.
func (s *accountAssignmentLister) List(selector labels.Selector) (ret []*v1alpha1.AccountAssignment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AccountAssignment))
	})
	return ret, err
}

// AccountAssignments returns an object that can list and get AccountAssignments.
func (s *accountAssignmentLister) AccountAssignments(namespace string) AccountAssignmentNamespaceLister {
	return accountAssignmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AccountAssignmentNamespaceLister helps list and get AccountAssignments.
// All objects returned here must be treated as read-only.
type AccountAssignmentNamespaceLister interface {
	// List lists all AccountAssignments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AccountAssignment, err error)
	// Get retrieves the AccountAssignment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AccountAssignment, error)
	AccountAssignmentNamespaceListerExpansion
}

// accountAssignmentNamespaceLister implements the AccountAssignmentNamespaceLister
// interface.
type accountAssignmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AccountAssignments in the indexer for a given namespace.
func (s accountAssignmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AccountAssignment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AccountAssignment))
	})
	return ret, err
}

// Get retrieves the AccountAssignment from the indexer for a given namespace and name.
func (s accountAssignmentNamespaceLister) Get(name string) (*v1alpha1.AccountAssignment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("accountassignment"), name)
	}
	return obj.(*v1alpha1.AccountAssignment), nil
}
