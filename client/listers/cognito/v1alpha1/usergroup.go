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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cognito/v1alpha1"
)

// UserGroupLister helps list UserGroups.
// All objects returned here must be treated as read-only.
type UserGroupLister interface {
	// List lists all UserGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.UserGroup, err error)
	// UserGroups returns an object that can list and get UserGroups.
	UserGroups(namespace string) UserGroupNamespaceLister
	UserGroupListerExpansion
}

// userGroupLister implements the UserGroupLister interface.
type userGroupLister struct {
	indexer cache.Indexer
}

// NewUserGroupLister returns a new UserGroupLister.
func NewUserGroupLister(indexer cache.Indexer) UserGroupLister {
	return &userGroupLister{indexer: indexer}
}

// List lists all UserGroups in the indexer.
func (s *userGroupLister) List(selector labels.Selector) (ret []*v1alpha1.UserGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserGroup))
	})
	return ret, err
}

// UserGroups returns an object that can list and get UserGroups.
func (s *userGroupLister) UserGroups(namespace string) UserGroupNamespaceLister {
	return userGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserGroupNamespaceLister helps list and get UserGroups.
// All objects returned here must be treated as read-only.
type UserGroupNamespaceLister interface {
	// List lists all UserGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.UserGroup, err error)
	// Get retrieves the UserGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.UserGroup, error)
	UserGroupNamespaceListerExpansion
}

// userGroupNamespaceLister implements the UserGroupNamespaceLister
// interface.
type userGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all UserGroups in the indexer for a given namespace.
func (s userGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.UserGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserGroup))
	})
	return ret, err
}

// Get retrieves the UserGroup from the indexer for a given namespace and name.
func (s userGroupNamespaceLister) Get(name string) (*v1alpha1.UserGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("usergroup"), name)
	}
	return obj.(*v1alpha1.UserGroup), nil
}
