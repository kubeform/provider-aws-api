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

// HumanTaskUiLister helps list HumanTaskUis.
// All objects returned here must be treated as read-only.
type HumanTaskUiLister interface {
	// List lists all HumanTaskUis in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HumanTaskUi, err error)
	// HumanTaskUis returns an object that can list and get HumanTaskUis.
	HumanTaskUis(namespace string) HumanTaskUiNamespaceLister
	HumanTaskUiListerExpansion
}

// humanTaskUiLister implements the HumanTaskUiLister interface.
type humanTaskUiLister struct {
	indexer cache.Indexer
}

// NewHumanTaskUiLister returns a new HumanTaskUiLister.
func NewHumanTaskUiLister(indexer cache.Indexer) HumanTaskUiLister {
	return &humanTaskUiLister{indexer: indexer}
}

// List lists all HumanTaskUis in the indexer.
func (s *humanTaskUiLister) List(selector labels.Selector) (ret []*v1alpha1.HumanTaskUi, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HumanTaskUi))
	})
	return ret, err
}

// HumanTaskUis returns an object that can list and get HumanTaskUis.
func (s *humanTaskUiLister) HumanTaskUis(namespace string) HumanTaskUiNamespaceLister {
	return humanTaskUiNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HumanTaskUiNamespaceLister helps list and get HumanTaskUis.
// All objects returned here must be treated as read-only.
type HumanTaskUiNamespaceLister interface {
	// List lists all HumanTaskUis in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HumanTaskUi, err error)
	// Get retrieves the HumanTaskUi from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HumanTaskUi, error)
	HumanTaskUiNamespaceListerExpansion
}

// humanTaskUiNamespaceLister implements the HumanTaskUiNamespaceLister
// interface.
type humanTaskUiNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HumanTaskUis in the indexer for a given namespace.
func (s humanTaskUiNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HumanTaskUi, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HumanTaskUi))
	})
	return ret, err
}

// Get retrieves the HumanTaskUi from the indexer for a given namespace and name.
func (s humanTaskUiNamespaceLister) Get(name string) (*v1alpha1.HumanTaskUi, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("humantaskui"), name)
	}
	return obj.(*v1alpha1.HumanTaskUi), nil
}
