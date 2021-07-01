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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/opsworks/v1alpha1"
)

// RdsDbInstanceLister helps list RdsDbInstances.
// All objects returned here must be treated as read-only.
type RdsDbInstanceLister interface {
	// List lists all RdsDbInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RdsDbInstance, err error)
	// RdsDbInstances returns an object that can list and get RdsDbInstances.
	RdsDbInstances(namespace string) RdsDbInstanceNamespaceLister
	RdsDbInstanceListerExpansion
}

// rdsDbInstanceLister implements the RdsDbInstanceLister interface.
type rdsDbInstanceLister struct {
	indexer cache.Indexer
}

// NewRdsDbInstanceLister returns a new RdsDbInstanceLister.
func NewRdsDbInstanceLister(indexer cache.Indexer) RdsDbInstanceLister {
	return &rdsDbInstanceLister{indexer: indexer}
}

// List lists all RdsDbInstances in the indexer.
func (s *rdsDbInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.RdsDbInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RdsDbInstance))
	})
	return ret, err
}

// RdsDbInstances returns an object that can list and get RdsDbInstances.
func (s *rdsDbInstanceLister) RdsDbInstances(namespace string) RdsDbInstanceNamespaceLister {
	return rdsDbInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RdsDbInstanceNamespaceLister helps list and get RdsDbInstances.
// All objects returned here must be treated as read-only.
type RdsDbInstanceNamespaceLister interface {
	// List lists all RdsDbInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RdsDbInstance, err error)
	// Get retrieves the RdsDbInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RdsDbInstance, error)
	RdsDbInstanceNamespaceListerExpansion
}

// rdsDbInstanceNamespaceLister implements the RdsDbInstanceNamespaceLister
// interface.
type rdsDbInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RdsDbInstances in the indexer for a given namespace.
func (s rdsDbInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RdsDbInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RdsDbInstance))
	})
	return ret, err
}

// Get retrieves the RdsDbInstance from the indexer for a given namespace and name.
func (s rdsDbInstanceNamespaceLister) Get(name string) (*v1alpha1.RdsDbInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rdsdbinstance"), name)
	}
	return obj.(*v1alpha1.RdsDbInstance), nil
}
