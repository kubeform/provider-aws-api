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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cur/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReportDefinitionLister helps list ReportDefinitions.
// All objects returned here must be treated as read-only.
type ReportDefinitionLister interface {
	// List lists all ReportDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReportDefinition, err error)
	// ReportDefinitions returns an object that can list and get ReportDefinitions.
	ReportDefinitions(namespace string) ReportDefinitionNamespaceLister
	ReportDefinitionListerExpansion
}

// reportDefinitionLister implements the ReportDefinitionLister interface.
type reportDefinitionLister struct {
	indexer cache.Indexer
}

// NewReportDefinitionLister returns a new ReportDefinitionLister.
func NewReportDefinitionLister(indexer cache.Indexer) ReportDefinitionLister {
	return &reportDefinitionLister{indexer: indexer}
}

// List lists all ReportDefinitions in the indexer.
func (s *reportDefinitionLister) List(selector labels.Selector) (ret []*v1alpha1.ReportDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReportDefinition))
	})
	return ret, err
}

// ReportDefinitions returns an object that can list and get ReportDefinitions.
func (s *reportDefinitionLister) ReportDefinitions(namespace string) ReportDefinitionNamespaceLister {
	return reportDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReportDefinitionNamespaceLister helps list and get ReportDefinitions.
// All objects returned here must be treated as read-only.
type ReportDefinitionNamespaceLister interface {
	// List lists all ReportDefinitions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReportDefinition, err error)
	// Get retrieves the ReportDefinition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ReportDefinition, error)
	ReportDefinitionNamespaceListerExpansion
}

// reportDefinitionNamespaceLister implements the ReportDefinitionNamespaceLister
// interface.
type reportDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReportDefinitions in the indexer for a given namespace.
func (s reportDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ReportDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReportDefinition))
	})
	return ret, err
}

// Get retrieves the ReportDefinition from the indexer for a given namespace and name.
func (s reportDefinitionNamespaceLister) Get(name string) (*v1alpha1.ReportDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("reportdefinition"), name)
	}
	return obj.(*v1alpha1.ReportDefinition), nil
}
