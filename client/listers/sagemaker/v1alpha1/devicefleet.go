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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/sagemaker/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DeviceFleetLister helps list DeviceFleets.
// All objects returned here must be treated as read-only.
type DeviceFleetLister interface {
	// List lists all DeviceFleets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DeviceFleet, err error)
	// DeviceFleets returns an object that can list and get DeviceFleets.
	DeviceFleets(namespace string) DeviceFleetNamespaceLister
	DeviceFleetListerExpansion
}

// deviceFleetLister implements the DeviceFleetLister interface.
type deviceFleetLister struct {
	indexer cache.Indexer
}

// NewDeviceFleetLister returns a new DeviceFleetLister.
func NewDeviceFleetLister(indexer cache.Indexer) DeviceFleetLister {
	return &deviceFleetLister{indexer: indexer}
}

// List lists all DeviceFleets in the indexer.
func (s *deviceFleetLister) List(selector labels.Selector) (ret []*v1alpha1.DeviceFleet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeviceFleet))
	})
	return ret, err
}

// DeviceFleets returns an object that can list and get DeviceFleets.
func (s *deviceFleetLister) DeviceFleets(namespace string) DeviceFleetNamespaceLister {
	return deviceFleetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DeviceFleetNamespaceLister helps list and get DeviceFleets.
// All objects returned here must be treated as read-only.
type DeviceFleetNamespaceLister interface {
	// List lists all DeviceFleets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DeviceFleet, err error)
	// Get retrieves the DeviceFleet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DeviceFleet, error)
	DeviceFleetNamespaceListerExpansion
}

// deviceFleetNamespaceLister implements the DeviceFleetNamespaceLister
// interface.
type deviceFleetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DeviceFleets in the indexer for a given namespace.
func (s deviceFleetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DeviceFleet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DeviceFleet))
	})
	return ret, err
}

// Get retrieves the DeviceFleet from the indexer for a given namespace and name.
func (s deviceFleetNamespaceLister) Get(name string) (*v1alpha1.DeviceFleet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("devicefleet"), name)
	}
	return obj.(*v1alpha1.DeviceFleet), nil
}
