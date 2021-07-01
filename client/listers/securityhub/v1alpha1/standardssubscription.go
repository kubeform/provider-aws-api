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

// StandardsSubscriptionLister helps list StandardsSubscriptions.
// All objects returned here must be treated as read-only.
type StandardsSubscriptionLister interface {
	// List lists all StandardsSubscriptions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StandardsSubscription, err error)
	// StandardsSubscriptions returns an object that can list and get StandardsSubscriptions.
	StandardsSubscriptions(namespace string) StandardsSubscriptionNamespaceLister
	StandardsSubscriptionListerExpansion
}

// standardsSubscriptionLister implements the StandardsSubscriptionLister interface.
type standardsSubscriptionLister struct {
	indexer cache.Indexer
}

// NewStandardsSubscriptionLister returns a new StandardsSubscriptionLister.
func NewStandardsSubscriptionLister(indexer cache.Indexer) StandardsSubscriptionLister {
	return &standardsSubscriptionLister{indexer: indexer}
}

// List lists all StandardsSubscriptions in the indexer.
func (s *standardsSubscriptionLister) List(selector labels.Selector) (ret []*v1alpha1.StandardsSubscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StandardsSubscription))
	})
	return ret, err
}

// StandardsSubscriptions returns an object that can list and get StandardsSubscriptions.
func (s *standardsSubscriptionLister) StandardsSubscriptions(namespace string) StandardsSubscriptionNamespaceLister {
	return standardsSubscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StandardsSubscriptionNamespaceLister helps list and get StandardsSubscriptions.
// All objects returned here must be treated as read-only.
type StandardsSubscriptionNamespaceLister interface {
	// List lists all StandardsSubscriptions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StandardsSubscription, err error)
	// Get retrieves the StandardsSubscription from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StandardsSubscription, error)
	StandardsSubscriptionNamespaceListerExpansion
}

// standardsSubscriptionNamespaceLister implements the StandardsSubscriptionNamespaceLister
// interface.
type standardsSubscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StandardsSubscriptions in the indexer for a given namespace.
func (s standardsSubscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StandardsSubscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StandardsSubscription))
	})
	return ret, err
}

// Get retrieves the StandardsSubscription from the indexer for a given namespace and name.
func (s standardsSubscriptionNamespaceLister) Get(name string) (*v1alpha1.StandardsSubscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("standardssubscription"), name)
	}
	return obj.(*v1alpha1.StandardsSubscription), nil
}
