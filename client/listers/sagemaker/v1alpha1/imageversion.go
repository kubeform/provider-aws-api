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

// ImageVersionLister helps list ImageVersions.
// All objects returned here must be treated as read-only.
type ImageVersionLister interface {
	// List lists all ImageVersions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ImageVersion, err error)
	// ImageVersions returns an object that can list and get ImageVersions.
	ImageVersions(namespace string) ImageVersionNamespaceLister
	ImageVersionListerExpansion
}

// imageVersionLister implements the ImageVersionLister interface.
type imageVersionLister struct {
	indexer cache.Indexer
}

// NewImageVersionLister returns a new ImageVersionLister.
func NewImageVersionLister(indexer cache.Indexer) ImageVersionLister {
	return &imageVersionLister{indexer: indexer}
}

// List lists all ImageVersions in the indexer.
func (s *imageVersionLister) List(selector labels.Selector) (ret []*v1alpha1.ImageVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageVersion))
	})
	return ret, err
}

// ImageVersions returns an object that can list and get ImageVersions.
func (s *imageVersionLister) ImageVersions(namespace string) ImageVersionNamespaceLister {
	return imageVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImageVersionNamespaceLister helps list and get ImageVersions.
// All objects returned here must be treated as read-only.
type ImageVersionNamespaceLister interface {
	// List lists all ImageVersions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ImageVersion, err error)
	// Get retrieves the ImageVersion from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ImageVersion, error)
	ImageVersionNamespaceListerExpansion
}

// imageVersionNamespaceLister implements the ImageVersionNamespaceLister
// interface.
type imageVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ImageVersions in the indexer for a given namespace.
func (s imageVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ImageVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageVersion))
	})
	return ret, err
}

// Get retrieves the ImageVersion from the indexer for a given namespace and name.
func (s imageVersionNamespaceLister) Get(name string) (*v1alpha1.ImageVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("imageversion"), name)
	}
	return obj.(*v1alpha1.ImageVersion), nil
}
