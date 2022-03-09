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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/datapipeline/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PipelineDefinitionLister helps list PipelineDefinitions.
// All objects returned here must be treated as read-only.
type PipelineDefinitionLister interface {
	// List lists all PipelineDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PipelineDefinition, err error)
	// PipelineDefinitions returns an object that can list and get PipelineDefinitions.
	PipelineDefinitions(namespace string) PipelineDefinitionNamespaceLister
	PipelineDefinitionListerExpansion
}

// pipelineDefinitionLister implements the PipelineDefinitionLister interface.
type pipelineDefinitionLister struct {
	indexer cache.Indexer
}

// NewPipelineDefinitionLister returns a new PipelineDefinitionLister.
func NewPipelineDefinitionLister(indexer cache.Indexer) PipelineDefinitionLister {
	return &pipelineDefinitionLister{indexer: indexer}
}

// List lists all PipelineDefinitions in the indexer.
func (s *pipelineDefinitionLister) List(selector labels.Selector) (ret []*v1alpha1.PipelineDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PipelineDefinition))
	})
	return ret, err
}

// PipelineDefinitions returns an object that can list and get PipelineDefinitions.
func (s *pipelineDefinitionLister) PipelineDefinitions(namespace string) PipelineDefinitionNamespaceLister {
	return pipelineDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PipelineDefinitionNamespaceLister helps list and get PipelineDefinitions.
// All objects returned here must be treated as read-only.
type PipelineDefinitionNamespaceLister interface {
	// List lists all PipelineDefinitions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PipelineDefinition, err error)
	// Get retrieves the PipelineDefinition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PipelineDefinition, error)
	PipelineDefinitionNamespaceListerExpansion
}

// pipelineDefinitionNamespaceLister implements the PipelineDefinitionNamespaceLister
// interface.
type pipelineDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PipelineDefinitions in the indexer for a given namespace.
func (s pipelineDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PipelineDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PipelineDefinition))
	})
	return ret, err
}

// Get retrieves the PipelineDefinition from the indexer for a given namespace and name.
func (s pipelineDefinitionNamespaceLister) Get(name string) (*v1alpha1.PipelineDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pipelinedefinition"), name)
	}
	return obj.(*v1alpha1.PipelineDefinition), nil
}
