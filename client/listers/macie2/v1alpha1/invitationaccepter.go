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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/macie2/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InvitationAccepterLister helps list InvitationAccepters.
// All objects returned here must be treated as read-only.
type InvitationAccepterLister interface {
	// List lists all InvitationAccepters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InvitationAccepter, err error)
	// InvitationAccepters returns an object that can list and get InvitationAccepters.
	InvitationAccepters(namespace string) InvitationAccepterNamespaceLister
	InvitationAccepterListerExpansion
}

// invitationAccepterLister implements the InvitationAccepterLister interface.
type invitationAccepterLister struct {
	indexer cache.Indexer
}

// NewInvitationAccepterLister returns a new InvitationAccepterLister.
func NewInvitationAccepterLister(indexer cache.Indexer) InvitationAccepterLister {
	return &invitationAccepterLister{indexer: indexer}
}

// List lists all InvitationAccepters in the indexer.
func (s *invitationAccepterLister) List(selector labels.Selector) (ret []*v1alpha1.InvitationAccepter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InvitationAccepter))
	})
	return ret, err
}

// InvitationAccepters returns an object that can list and get InvitationAccepters.
func (s *invitationAccepterLister) InvitationAccepters(namespace string) InvitationAccepterNamespaceLister {
	return invitationAccepterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InvitationAccepterNamespaceLister helps list and get InvitationAccepters.
// All objects returned here must be treated as read-only.
type InvitationAccepterNamespaceLister interface {
	// List lists all InvitationAccepters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InvitationAccepter, err error)
	// Get retrieves the InvitationAccepter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InvitationAccepter, error)
	InvitationAccepterNamespaceListerExpansion
}

// invitationAccepterNamespaceLister implements the InvitationAccepterNamespaceLister
// interface.
type invitationAccepterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InvitationAccepters in the indexer for a given namespace.
func (s invitationAccepterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InvitationAccepter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InvitationAccepter))
	})
	return ret, err
}

// Get retrieves the InvitationAccepter from the indexer for a given namespace and name.
func (s invitationAccepterNamespaceLister) Get(name string) (*v1alpha1.InvitationAccepter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("invitationaccepter"), name)
	}
	return obj.(*v1alpha1.InvitationAccepter), nil
}
