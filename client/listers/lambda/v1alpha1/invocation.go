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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/lambda/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InvocationLister helps list Invocations.
// All objects returned here must be treated as read-only.
type InvocationLister interface {
	// List lists all Invocations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Invocation, err error)
	// Invocations returns an object that can list and get Invocations.
	Invocations(namespace string) InvocationNamespaceLister
	InvocationListerExpansion
}

// invocationLister implements the InvocationLister interface.
type invocationLister struct {
	indexer cache.Indexer
}

// NewInvocationLister returns a new InvocationLister.
func NewInvocationLister(indexer cache.Indexer) InvocationLister {
	return &invocationLister{indexer: indexer}
}

// List lists all Invocations in the indexer.
func (s *invocationLister) List(selector labels.Selector) (ret []*v1alpha1.Invocation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Invocation))
	})
	return ret, err
}

// Invocations returns an object that can list and get Invocations.
func (s *invocationLister) Invocations(namespace string) InvocationNamespaceLister {
	return invocationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InvocationNamespaceLister helps list and get Invocations.
// All objects returned here must be treated as read-only.
type InvocationNamespaceLister interface {
	// List lists all Invocations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Invocation, err error)
	// Get retrieves the Invocation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Invocation, error)
	InvocationNamespaceListerExpansion
}

// invocationNamespaceLister implements the InvocationNamespaceLister
// interface.
type invocationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Invocations in the indexer for a given namespace.
func (s invocationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Invocation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Invocation))
	})
	return ret, err
}

// Get retrieves the Invocation from the indexer for a given namespace and name.
func (s invocationNamespaceLister) Get(name string) (*v1alpha1.Invocation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("invocation"), name)
	}
	return obj.(*v1alpha1.Invocation), nil
}
