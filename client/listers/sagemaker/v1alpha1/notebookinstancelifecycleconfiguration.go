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

// NotebookInstanceLifecycleConfigurationLister helps list NotebookInstanceLifecycleConfigurations.
// All objects returned here must be treated as read-only.
type NotebookInstanceLifecycleConfigurationLister interface {
	// List lists all NotebookInstanceLifecycleConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NotebookInstanceLifecycleConfiguration, err error)
	// NotebookInstanceLifecycleConfigurations returns an object that can list and get NotebookInstanceLifecycleConfigurations.
	NotebookInstanceLifecycleConfigurations(namespace string) NotebookInstanceLifecycleConfigurationNamespaceLister
	NotebookInstanceLifecycleConfigurationListerExpansion
}

// notebookInstanceLifecycleConfigurationLister implements the NotebookInstanceLifecycleConfigurationLister interface.
type notebookInstanceLifecycleConfigurationLister struct {
	indexer cache.Indexer
}

// NewNotebookInstanceLifecycleConfigurationLister returns a new NotebookInstanceLifecycleConfigurationLister.
func NewNotebookInstanceLifecycleConfigurationLister(indexer cache.Indexer) NotebookInstanceLifecycleConfigurationLister {
	return &notebookInstanceLifecycleConfigurationLister{indexer: indexer}
}

// List lists all NotebookInstanceLifecycleConfigurations in the indexer.
func (s *notebookInstanceLifecycleConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.NotebookInstanceLifecycleConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NotebookInstanceLifecycleConfiguration))
	})
	return ret, err
}

// NotebookInstanceLifecycleConfigurations returns an object that can list and get NotebookInstanceLifecycleConfigurations.
func (s *notebookInstanceLifecycleConfigurationLister) NotebookInstanceLifecycleConfigurations(namespace string) NotebookInstanceLifecycleConfigurationNamespaceLister {
	return notebookInstanceLifecycleConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NotebookInstanceLifecycleConfigurationNamespaceLister helps list and get NotebookInstanceLifecycleConfigurations.
// All objects returned here must be treated as read-only.
type NotebookInstanceLifecycleConfigurationNamespaceLister interface {
	// List lists all NotebookInstanceLifecycleConfigurations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NotebookInstanceLifecycleConfiguration, err error)
	// Get retrieves the NotebookInstanceLifecycleConfiguration from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NotebookInstanceLifecycleConfiguration, error)
	NotebookInstanceLifecycleConfigurationNamespaceListerExpansion
}

// notebookInstanceLifecycleConfigurationNamespaceLister implements the NotebookInstanceLifecycleConfigurationNamespaceLister
// interface.
type notebookInstanceLifecycleConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NotebookInstanceLifecycleConfigurations in the indexer for a given namespace.
func (s notebookInstanceLifecycleConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NotebookInstanceLifecycleConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NotebookInstanceLifecycleConfiguration))
	})
	return ret, err
}

// Get retrieves the NotebookInstanceLifecycleConfiguration from the indexer for a given namespace and name.
func (s notebookInstanceLifecycleConfigurationNamespaceLister) Get(name string) (*v1alpha1.NotebookInstanceLifecycleConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("notebookinstancelifecycleconfiguration"), name)
	}
	return obj.(*v1alpha1.NotebookInstanceLifecycleConfiguration), nil
}
