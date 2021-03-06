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

// AppImageConfigLister helps list AppImageConfigs.
// All objects returned here must be treated as read-only.
type AppImageConfigLister interface {
	// List lists all AppImageConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppImageConfig, err error)
	// AppImageConfigs returns an object that can list and get AppImageConfigs.
	AppImageConfigs(namespace string) AppImageConfigNamespaceLister
	AppImageConfigListerExpansion
}

// appImageConfigLister implements the AppImageConfigLister interface.
type appImageConfigLister struct {
	indexer cache.Indexer
}

// NewAppImageConfigLister returns a new AppImageConfigLister.
func NewAppImageConfigLister(indexer cache.Indexer) AppImageConfigLister {
	return &appImageConfigLister{indexer: indexer}
}

// List lists all AppImageConfigs in the indexer.
func (s *appImageConfigLister) List(selector labels.Selector) (ret []*v1alpha1.AppImageConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppImageConfig))
	})
	return ret, err
}

// AppImageConfigs returns an object that can list and get AppImageConfigs.
func (s *appImageConfigLister) AppImageConfigs(namespace string) AppImageConfigNamespaceLister {
	return appImageConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppImageConfigNamespaceLister helps list and get AppImageConfigs.
// All objects returned here must be treated as read-only.
type AppImageConfigNamespaceLister interface {
	// List lists all AppImageConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppImageConfig, err error)
	// Get retrieves the AppImageConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AppImageConfig, error)
	AppImageConfigNamespaceListerExpansion
}

// appImageConfigNamespaceLister implements the AppImageConfigNamespaceLister
// interface.
type appImageConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppImageConfigs in the indexer for a given namespace.
func (s appImageConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppImageConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppImageConfig))
	})
	return ret, err
}

// Get retrieves the AppImageConfig from the indexer for a given namespace and name.
func (s appImageConfigNamespaceLister) Get(name string) (*v1alpha1.AppImageConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appimageconfig"), name)
	}
	return obj.(*v1alpha1.AppImageConfig), nil
}
