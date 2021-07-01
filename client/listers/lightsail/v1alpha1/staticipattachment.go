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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubeform.dev/provider-aws-api/apis/lightsail/v1alpha1"
)

// StaticIPAttachmentLister helps list StaticIPAttachments.
// All objects returned here must be treated as read-only.
type StaticIPAttachmentLister interface {
	// List lists all StaticIPAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StaticIPAttachment, err error)
	// StaticIPAttachments returns an object that can list and get StaticIPAttachments.
	StaticIPAttachments(namespace string) StaticIPAttachmentNamespaceLister
	StaticIPAttachmentListerExpansion
}

// staticIPAttachmentLister implements the StaticIPAttachmentLister interface.
type staticIPAttachmentLister struct {
	indexer cache.Indexer
}

// NewStaticIPAttachmentLister returns a new StaticIPAttachmentLister.
func NewStaticIPAttachmentLister(indexer cache.Indexer) StaticIPAttachmentLister {
	return &staticIPAttachmentLister{indexer: indexer}
}

// List lists all StaticIPAttachments in the indexer.
func (s *staticIPAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.StaticIPAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StaticIPAttachment))
	})
	return ret, err
}

// StaticIPAttachments returns an object that can list and get StaticIPAttachments.
func (s *staticIPAttachmentLister) StaticIPAttachments(namespace string) StaticIPAttachmentNamespaceLister {
	return staticIPAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StaticIPAttachmentNamespaceLister helps list and get StaticIPAttachments.
// All objects returned here must be treated as read-only.
type StaticIPAttachmentNamespaceLister interface {
	// List lists all StaticIPAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StaticIPAttachment, err error)
	// Get retrieves the StaticIPAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StaticIPAttachment, error)
	StaticIPAttachmentNamespaceListerExpansion
}

// staticIPAttachmentNamespaceLister implements the StaticIPAttachmentNamespaceLister
// interface.
type staticIPAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StaticIPAttachments in the indexer for a given namespace.
func (s staticIPAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StaticIPAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StaticIPAttachment))
	})
	return ret, err
}

// Get retrieves the StaticIPAttachment from the indexer for a given namespace and name.
func (s staticIPAttachmentNamespaceLister) Get(name string) (*v1alpha1.StaticIPAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("staticipattachment"), name)
	}
	return obj.(*v1alpha1.StaticIPAttachment), nil
}
