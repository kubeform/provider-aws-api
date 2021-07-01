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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/macie2/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClassificationJobLister helps list ClassificationJobs.
// All objects returned here must be treated as read-only.
type ClassificationJobLister interface {
	// List lists all ClassificationJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClassificationJob, err error)
	// ClassificationJobs returns an object that can list and get ClassificationJobs.
	ClassificationJobs(namespace string) ClassificationJobNamespaceLister
	ClassificationJobListerExpansion
}

// classificationJobLister implements the ClassificationJobLister interface.
type classificationJobLister struct {
	indexer cache.Indexer
}

// NewClassificationJobLister returns a new ClassificationJobLister.
func NewClassificationJobLister(indexer cache.Indexer) ClassificationJobLister {
	return &classificationJobLister{indexer: indexer}
}

// List lists all ClassificationJobs in the indexer.
func (s *classificationJobLister) List(selector labels.Selector) (ret []*v1alpha1.ClassificationJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClassificationJob))
	})
	return ret, err
}

// ClassificationJobs returns an object that can list and get ClassificationJobs.
func (s *classificationJobLister) ClassificationJobs(namespace string) ClassificationJobNamespaceLister {
	return classificationJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ClassificationJobNamespaceLister helps list and get ClassificationJobs.
// All objects returned here must be treated as read-only.
type ClassificationJobNamespaceLister interface {
	// List lists all ClassificationJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ClassificationJob, err error)
	// Get retrieves the ClassificationJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ClassificationJob, error)
	ClassificationJobNamespaceListerExpansion
}

// classificationJobNamespaceLister implements the ClassificationJobNamespaceLister
// interface.
type classificationJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ClassificationJobs in the indexer for a given namespace.
func (s classificationJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ClassificationJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ClassificationJob))
	})
	return ret, err
}

// Get retrieves the ClassificationJob from the indexer for a given namespace and name.
func (s classificationJobNamespaceLister) Get(name string) (*v1alpha1.ClassificationJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("classificationjob"), name)
	}
	return obj.(*v1alpha1.ClassificationJob), nil
}
