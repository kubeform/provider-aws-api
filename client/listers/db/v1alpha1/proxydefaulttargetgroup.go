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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/db/v1alpha1"
)

// ProxyDefaultTargetGroupLister helps list ProxyDefaultTargetGroups.
// All objects returned here must be treated as read-only.
type ProxyDefaultTargetGroupLister interface {
	// List lists all ProxyDefaultTargetGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProxyDefaultTargetGroup, err error)
	// ProxyDefaultTargetGroups returns an object that can list and get ProxyDefaultTargetGroups.
	ProxyDefaultTargetGroups(namespace string) ProxyDefaultTargetGroupNamespaceLister
	ProxyDefaultTargetGroupListerExpansion
}

// proxyDefaultTargetGroupLister implements the ProxyDefaultTargetGroupLister interface.
type proxyDefaultTargetGroupLister struct {
	indexer cache.Indexer
}

// NewProxyDefaultTargetGroupLister returns a new ProxyDefaultTargetGroupLister.
func NewProxyDefaultTargetGroupLister(indexer cache.Indexer) ProxyDefaultTargetGroupLister {
	return &proxyDefaultTargetGroupLister{indexer: indexer}
}

// List lists all ProxyDefaultTargetGroups in the indexer.
func (s *proxyDefaultTargetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.ProxyDefaultTargetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProxyDefaultTargetGroup))
	})
	return ret, err
}

// ProxyDefaultTargetGroups returns an object that can list and get ProxyDefaultTargetGroups.
func (s *proxyDefaultTargetGroupLister) ProxyDefaultTargetGroups(namespace string) ProxyDefaultTargetGroupNamespaceLister {
	return proxyDefaultTargetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProxyDefaultTargetGroupNamespaceLister helps list and get ProxyDefaultTargetGroups.
// All objects returned here must be treated as read-only.
type ProxyDefaultTargetGroupNamespaceLister interface {
	// List lists all ProxyDefaultTargetGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ProxyDefaultTargetGroup, err error)
	// Get retrieves the ProxyDefaultTargetGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ProxyDefaultTargetGroup, error)
	ProxyDefaultTargetGroupNamespaceListerExpansion
}

// proxyDefaultTargetGroupNamespaceLister implements the ProxyDefaultTargetGroupNamespaceLister
// interface.
type proxyDefaultTargetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProxyDefaultTargetGroups in the indexer for a given namespace.
func (s proxyDefaultTargetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ProxyDefaultTargetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ProxyDefaultTargetGroup))
	})
	return ret, err
}

// Get retrieves the ProxyDefaultTargetGroup from the indexer for a given namespace and name.
func (s proxyDefaultTargetGroupNamespaceLister) Get(name string) (*v1alpha1.ProxyDefaultTargetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("proxydefaulttargetgroup"), name)
	}
	return obj.(*v1alpha1.ProxyDefaultTargetGroup), nil
}
