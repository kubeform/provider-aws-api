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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/elb/v1alpha1"
)

// ElbLister helps list Elbs.
// All objects returned here must be treated as read-only.
type ElbLister interface {
	// List lists all Elbs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Elb, err error)
	// Elbs returns an object that can list and get Elbs.
	Elbs(namespace string) ElbNamespaceLister
	ElbListerExpansion
}

// elbLister implements the ElbLister interface.
type elbLister struct {
	indexer cache.Indexer
}

// NewElbLister returns a new ElbLister.
func NewElbLister(indexer cache.Indexer) ElbLister {
	return &elbLister{indexer: indexer}
}

// List lists all Elbs in the indexer.
func (s *elbLister) List(selector labels.Selector) (ret []*v1alpha1.Elb, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Elb))
	})
	return ret, err
}

// Elbs returns an object that can list and get Elbs.
func (s *elbLister) Elbs(namespace string) ElbNamespaceLister {
	return elbNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ElbNamespaceLister helps list and get Elbs.
// All objects returned here must be treated as read-only.
type ElbNamespaceLister interface {
	// List lists all Elbs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Elb, err error)
	// Get retrieves the Elb from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Elb, error)
	ElbNamespaceListerExpansion
}

// elbNamespaceLister implements the ElbNamespaceLister
// interface.
type elbNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Elbs in the indexer for a given namespace.
func (s elbNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Elb, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Elb))
	})
	return ret, err
}

// Get retrieves the Elb from the indexer for a given namespace and name.
func (s elbNamespaceLister) Get(name string) (*v1alpha1.Elb, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("elb"), name)
	}
	return obj.(*v1alpha1.Elb), nil
}
