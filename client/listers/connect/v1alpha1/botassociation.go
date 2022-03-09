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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/connect/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BotAssociationLister helps list BotAssociations.
// All objects returned here must be treated as read-only.
type BotAssociationLister interface {
	// List lists all BotAssociations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BotAssociation, err error)
	// BotAssociations returns an object that can list and get BotAssociations.
	BotAssociations(namespace string) BotAssociationNamespaceLister
	BotAssociationListerExpansion
}

// botAssociationLister implements the BotAssociationLister interface.
type botAssociationLister struct {
	indexer cache.Indexer
}

// NewBotAssociationLister returns a new BotAssociationLister.
func NewBotAssociationLister(indexer cache.Indexer) BotAssociationLister {
	return &botAssociationLister{indexer: indexer}
}

// List lists all BotAssociations in the indexer.
func (s *botAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.BotAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BotAssociation))
	})
	return ret, err
}

// BotAssociations returns an object that can list and get BotAssociations.
func (s *botAssociationLister) BotAssociations(namespace string) BotAssociationNamespaceLister {
	return botAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BotAssociationNamespaceLister helps list and get BotAssociations.
// All objects returned here must be treated as read-only.
type BotAssociationNamespaceLister interface {
	// List lists all BotAssociations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BotAssociation, err error)
	// Get retrieves the BotAssociation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BotAssociation, error)
	BotAssociationNamespaceListerExpansion
}

// botAssociationNamespaceLister implements the BotAssociationNamespaceLister
// interface.
type botAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BotAssociations in the indexer for a given namespace.
func (s botAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BotAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BotAssociation))
	})
	return ret, err
}

// Get retrieves the BotAssociation from the indexer for a given namespace and name.
func (s botAssociationNamespaceLister) Get(name string) (*v1alpha1.BotAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("botassociation"), name)
	}
	return obj.(*v1alpha1.BotAssociation), nil
}
