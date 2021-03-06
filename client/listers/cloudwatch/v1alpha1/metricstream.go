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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/cloudwatch/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MetricStreamLister helps list MetricStreams.
// All objects returned here must be treated as read-only.
type MetricStreamLister interface {
	// List lists all MetricStreams in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MetricStream, err error)
	// MetricStreams returns an object that can list and get MetricStreams.
	MetricStreams(namespace string) MetricStreamNamespaceLister
	MetricStreamListerExpansion
}

// metricStreamLister implements the MetricStreamLister interface.
type metricStreamLister struct {
	indexer cache.Indexer
}

// NewMetricStreamLister returns a new MetricStreamLister.
func NewMetricStreamLister(indexer cache.Indexer) MetricStreamLister {
	return &metricStreamLister{indexer: indexer}
}

// List lists all MetricStreams in the indexer.
func (s *metricStreamLister) List(selector labels.Selector) (ret []*v1alpha1.MetricStream, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetricStream))
	})
	return ret, err
}

// MetricStreams returns an object that can list and get MetricStreams.
func (s *metricStreamLister) MetricStreams(namespace string) MetricStreamNamespaceLister {
	return metricStreamNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MetricStreamNamespaceLister helps list and get MetricStreams.
// All objects returned here must be treated as read-only.
type MetricStreamNamespaceLister interface {
	// List lists all MetricStreams in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MetricStream, err error)
	// Get retrieves the MetricStream from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MetricStream, error)
	MetricStreamNamespaceListerExpansion
}

// metricStreamNamespaceLister implements the MetricStreamNamespaceLister
// interface.
type metricStreamNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MetricStreams in the indexer for a given namespace.
func (s metricStreamNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MetricStream, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MetricStream))
	})
	return ret, err
}

// Get retrieves the MetricStream from the indexer for a given namespace and name.
func (s metricStreamNamespaceLister) Get(name string) (*v1alpha1.MetricStream, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("metricstream"), name)
	}
	return obj.(*v1alpha1.MetricStream), nil
}
