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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecr/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RegistryScanningConfigurationLister helps list RegistryScanningConfigurations.
// All objects returned here must be treated as read-only.
type RegistryScanningConfigurationLister interface {
	// List lists all RegistryScanningConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegistryScanningConfiguration, err error)
	// RegistryScanningConfigurations returns an object that can list and get RegistryScanningConfigurations.
	RegistryScanningConfigurations(namespace string) RegistryScanningConfigurationNamespaceLister
	RegistryScanningConfigurationListerExpansion
}

// registryScanningConfigurationLister implements the RegistryScanningConfigurationLister interface.
type registryScanningConfigurationLister struct {
	indexer cache.Indexer
}

// NewRegistryScanningConfigurationLister returns a new RegistryScanningConfigurationLister.
func NewRegistryScanningConfigurationLister(indexer cache.Indexer) RegistryScanningConfigurationLister {
	return &registryScanningConfigurationLister{indexer: indexer}
}

// List lists all RegistryScanningConfigurations in the indexer.
func (s *registryScanningConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.RegistryScanningConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegistryScanningConfiguration))
	})
	return ret, err
}

// RegistryScanningConfigurations returns an object that can list and get RegistryScanningConfigurations.
func (s *registryScanningConfigurationLister) RegistryScanningConfigurations(namespace string) RegistryScanningConfigurationNamespaceLister {
	return registryScanningConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RegistryScanningConfigurationNamespaceLister helps list and get RegistryScanningConfigurations.
// All objects returned here must be treated as read-only.
type RegistryScanningConfigurationNamespaceLister interface {
	// List lists all RegistryScanningConfigurations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RegistryScanningConfiguration, err error)
	// Get retrieves the RegistryScanningConfiguration from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RegistryScanningConfiguration, error)
	RegistryScanningConfigurationNamespaceListerExpansion
}

// registryScanningConfigurationNamespaceLister implements the RegistryScanningConfigurationNamespaceLister
// interface.
type registryScanningConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RegistryScanningConfigurations in the indexer for a given namespace.
func (s registryScanningConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RegistryScanningConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RegistryScanningConfiguration))
	})
	return ret, err
}

// Get retrieves the RegistryScanningConfiguration from the indexer for a given namespace and name.
func (s registryScanningConfigurationNamespaceLister) Get(name string) (*v1alpha1.RegistryScanningConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("registryscanningconfiguration"), name)
	}
	return obj.(*v1alpha1.RegistryScanningConfiguration), nil
}
