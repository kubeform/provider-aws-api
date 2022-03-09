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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HostLister helps list Hosts.
// All objects returned here must be treated as read-only.
type HostLister interface {
	// List lists all Hosts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Host, err error)
	// Hosts returns an object that can list and get Hosts.
	Hosts(namespace string) HostNamespaceLister
	HostListerExpansion
}

// hostLister implements the HostLister interface.
type hostLister struct {
	indexer cache.Indexer
}

// NewHostLister returns a new HostLister.
func NewHostLister(indexer cache.Indexer) HostLister {
	return &hostLister{indexer: indexer}
}

// List lists all Hosts in the indexer.
func (s *hostLister) List(selector labels.Selector) (ret []*v1alpha1.Host, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Host))
	})
	return ret, err
}

// Hosts returns an object that can list and get Hosts.
func (s *hostLister) Hosts(namespace string) HostNamespaceLister {
	return hostNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HostNamespaceLister helps list and get Hosts.
// All objects returned here must be treated as read-only.
type HostNamespaceLister interface {
	// List lists all Hosts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Host, err error)
	// Get retrieves the Host from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Host, error)
	HostNamespaceListerExpansion
}

// hostNamespaceLister implements the HostNamespaceLister
// interface.
type hostNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Hosts in the indexer for a given namespace.
func (s hostNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Host, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Host))
	})
	return ret, err
}

// Get retrieves the Host from the indexer for a given namespace and name.
func (s hostNamespaceLister) Get(name string) (*v1alpha1.Host, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("host"), name)
	}
	return obj.(*v1alpha1.Host), nil
}
