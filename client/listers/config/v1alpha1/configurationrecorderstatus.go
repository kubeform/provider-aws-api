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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/config/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConfigurationRecorderStatusLister helps list ConfigurationRecorderStatuses.
// All objects returned here must be treated as read-only.
type ConfigurationRecorderStatusLister interface {
	// List lists all ConfigurationRecorderStatuses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigurationRecorderStatus, err error)
	// ConfigurationRecorderStatuses returns an object that can list and get ConfigurationRecorderStatuses.
	ConfigurationRecorderStatuses(namespace string) ConfigurationRecorderStatusNamespaceLister
	ConfigurationRecorderStatusListerExpansion
}

// configurationRecorderStatusLister implements the ConfigurationRecorderStatusLister interface.
type configurationRecorderStatusLister struct {
	indexer cache.Indexer
}

// NewConfigurationRecorderStatusLister returns a new ConfigurationRecorderStatusLister.
func NewConfigurationRecorderStatusLister(indexer cache.Indexer) ConfigurationRecorderStatusLister {
	return &configurationRecorderStatusLister{indexer: indexer}
}

// List lists all ConfigurationRecorderStatuses in the indexer.
func (s *configurationRecorderStatusLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigurationRecorderStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigurationRecorderStatus))
	})
	return ret, err
}

// ConfigurationRecorderStatuses returns an object that can list and get ConfigurationRecorderStatuses.
func (s *configurationRecorderStatusLister) ConfigurationRecorderStatuses(namespace string) ConfigurationRecorderStatusNamespaceLister {
	return configurationRecorderStatusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigurationRecorderStatusNamespaceLister helps list and get ConfigurationRecorderStatuses.
// All objects returned here must be treated as read-only.
type ConfigurationRecorderStatusNamespaceLister interface {
	// List lists all ConfigurationRecorderStatuses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigurationRecorderStatus, err error)
	// Get retrieves the ConfigurationRecorderStatus from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConfigurationRecorderStatus, error)
	ConfigurationRecorderStatusNamespaceListerExpansion
}

// configurationRecorderStatusNamespaceLister implements the ConfigurationRecorderStatusNamespaceLister
// interface.
type configurationRecorderStatusNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigurationRecorderStatuses in the indexer for a given namespace.
func (s configurationRecorderStatusNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigurationRecorderStatus, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigurationRecorderStatus))
	})
	return ret, err
}

// Get retrieves the ConfigurationRecorderStatus from the indexer for a given namespace and name.
func (s configurationRecorderStatusNamespaceLister) Get(name string) (*v1alpha1.ConfigurationRecorderStatus, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("configurationrecorderstatus"), name)
	}
	return obj.(*v1alpha1.ConfigurationRecorderStatus), nil
}