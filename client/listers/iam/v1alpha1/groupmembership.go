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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/iam/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GroupMembershipLister helps list GroupMemberships.
// All objects returned here must be treated as read-only.
type GroupMembershipLister interface {
	// List lists all GroupMemberships in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GroupMembership, err error)
	// GroupMemberships returns an object that can list and get GroupMemberships.
	GroupMemberships(namespace string) GroupMembershipNamespaceLister
	GroupMembershipListerExpansion
}

// groupMembershipLister implements the GroupMembershipLister interface.
type groupMembershipLister struct {
	indexer cache.Indexer
}

// NewGroupMembershipLister returns a new GroupMembershipLister.
func NewGroupMembershipLister(indexer cache.Indexer) GroupMembershipLister {
	return &groupMembershipLister{indexer: indexer}
}

// List lists all GroupMemberships in the indexer.
func (s *groupMembershipLister) List(selector labels.Selector) (ret []*v1alpha1.GroupMembership, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GroupMembership))
	})
	return ret, err
}

// GroupMemberships returns an object that can list and get GroupMemberships.
func (s *groupMembershipLister) GroupMemberships(namespace string) GroupMembershipNamespaceLister {
	return groupMembershipNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GroupMembershipNamespaceLister helps list and get GroupMemberships.
// All objects returned here must be treated as read-only.
type GroupMembershipNamespaceLister interface {
	// List lists all GroupMemberships in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GroupMembership, err error)
	// Get retrieves the GroupMembership from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GroupMembership, error)
	GroupMembershipNamespaceListerExpansion
}

// groupMembershipNamespaceLister implements the GroupMembershipNamespaceLister
// interface.
type groupMembershipNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GroupMemberships in the indexer for a given namespace.
func (s groupMembershipNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GroupMembership, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GroupMembership))
	})
	return ret, err
}

// Get retrieves the GroupMembership from the indexer for a given namespace and name.
func (s groupMembershipNamespaceLister) Get(name string) (*v1alpha1.GroupMembership, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("groupmembership"), name)
	}
	return obj.(*v1alpha1.GroupMembership), nil
}
