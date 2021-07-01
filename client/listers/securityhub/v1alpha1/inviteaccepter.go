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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/securityhub/v1alpha1"
)

// InviteAccepterLister helps list InviteAccepters.
// All objects returned here must be treated as read-only.
type InviteAccepterLister interface {
	// List lists all InviteAccepters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InviteAccepter, err error)
	// InviteAccepters returns an object that can list and get InviteAccepters.
	InviteAccepters(namespace string) InviteAccepterNamespaceLister
	InviteAccepterListerExpansion
}

// inviteAccepterLister implements the InviteAccepterLister interface.
type inviteAccepterLister struct {
	indexer cache.Indexer
}

// NewInviteAccepterLister returns a new InviteAccepterLister.
func NewInviteAccepterLister(indexer cache.Indexer) InviteAccepterLister {
	return &inviteAccepterLister{indexer: indexer}
}

// List lists all InviteAccepters in the indexer.
func (s *inviteAccepterLister) List(selector labels.Selector) (ret []*v1alpha1.InviteAccepter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InviteAccepter))
	})
	return ret, err
}

// InviteAccepters returns an object that can list and get InviteAccepters.
func (s *inviteAccepterLister) InviteAccepters(namespace string) InviteAccepterNamespaceLister {
	return inviteAccepterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InviteAccepterNamespaceLister helps list and get InviteAccepters.
// All objects returned here must be treated as read-only.
type InviteAccepterNamespaceLister interface {
	// List lists all InviteAccepters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InviteAccepter, err error)
	// Get retrieves the InviteAccepter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InviteAccepter, error)
	InviteAccepterNamespaceListerExpansion
}

// inviteAccepterNamespaceLister implements the InviteAccepterNamespaceLister
// interface.
type inviteAccepterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InviteAccepters in the indexer for a given namespace.
func (s inviteAccepterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InviteAccepter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InviteAccepter))
	})
	return ret, err
}

// Get retrieves the InviteAccepter from the indexer for a given namespace and name.
func (s inviteAccepterNamespaceLister) Get(name string) (*v1alpha1.InviteAccepter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("inviteaccepter"), name)
	}
	return obj.(*v1alpha1.InviteAccepter), nil
}
