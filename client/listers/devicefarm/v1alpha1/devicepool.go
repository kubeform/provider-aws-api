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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/devicefarm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DevicePoolLister helps list DevicePools.
// All objects returned here must be treated as read-only.
type DevicePoolLister interface {
	// List lists all DevicePools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DevicePool, err error)
	// DevicePools returns an object that can list and get DevicePools.
	DevicePools(namespace string) DevicePoolNamespaceLister
	DevicePoolListerExpansion
}

// devicePoolLister implements the DevicePoolLister interface.
type devicePoolLister struct {
	indexer cache.Indexer
}

// NewDevicePoolLister returns a new DevicePoolLister.
func NewDevicePoolLister(indexer cache.Indexer) DevicePoolLister {
	return &devicePoolLister{indexer: indexer}
}

// List lists all DevicePools in the indexer.
func (s *devicePoolLister) List(selector labels.Selector) (ret []*v1alpha1.DevicePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DevicePool))
	})
	return ret, err
}

// DevicePools returns an object that can list and get DevicePools.
func (s *devicePoolLister) DevicePools(namespace string) DevicePoolNamespaceLister {
	return devicePoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DevicePoolNamespaceLister helps list and get DevicePools.
// All objects returned here must be treated as read-only.
type DevicePoolNamespaceLister interface {
	// List lists all DevicePools in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DevicePool, err error)
	// Get retrieves the DevicePool from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DevicePool, error)
	DevicePoolNamespaceListerExpansion
}

// devicePoolNamespaceLister implements the DevicePoolNamespaceLister
// interface.
type devicePoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DevicePools in the indexer for a given namespace.
func (s devicePoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DevicePool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DevicePool))
	})
	return ret, err
}

// Get retrieves the DevicePool from the indexer for a given namespace and name.
func (s devicePoolNamespaceLister) Get(name string) (*v1alpha1.DevicePool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("devicepool"), name)
	}
	return obj.(*v1alpha1.DevicePool), nil
}