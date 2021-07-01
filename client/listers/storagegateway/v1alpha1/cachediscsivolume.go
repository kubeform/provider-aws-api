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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/storagegateway/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CachedIscsiVolumeLister helps list CachedIscsiVolumes.
// All objects returned here must be treated as read-only.
type CachedIscsiVolumeLister interface {
	// List lists all CachedIscsiVolumes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CachedIscsiVolume, err error)
	// CachedIscsiVolumes returns an object that can list and get CachedIscsiVolumes.
	CachedIscsiVolumes(namespace string) CachedIscsiVolumeNamespaceLister
	CachedIscsiVolumeListerExpansion
}

// cachedIscsiVolumeLister implements the CachedIscsiVolumeLister interface.
type cachedIscsiVolumeLister struct {
	indexer cache.Indexer
}

// NewCachedIscsiVolumeLister returns a new CachedIscsiVolumeLister.
func NewCachedIscsiVolumeLister(indexer cache.Indexer) CachedIscsiVolumeLister {
	return &cachedIscsiVolumeLister{indexer: indexer}
}

// List lists all CachedIscsiVolumes in the indexer.
func (s *cachedIscsiVolumeLister) List(selector labels.Selector) (ret []*v1alpha1.CachedIscsiVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CachedIscsiVolume))
	})
	return ret, err
}

// CachedIscsiVolumes returns an object that can list and get CachedIscsiVolumes.
func (s *cachedIscsiVolumeLister) CachedIscsiVolumes(namespace string) CachedIscsiVolumeNamespaceLister {
	return cachedIscsiVolumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CachedIscsiVolumeNamespaceLister helps list and get CachedIscsiVolumes.
// All objects returned here must be treated as read-only.
type CachedIscsiVolumeNamespaceLister interface {
	// List lists all CachedIscsiVolumes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CachedIscsiVolume, err error)
	// Get retrieves the CachedIscsiVolume from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CachedIscsiVolume, error)
	CachedIscsiVolumeNamespaceListerExpansion
}

// cachedIscsiVolumeNamespaceLister implements the CachedIscsiVolumeNamespaceLister
// interface.
type cachedIscsiVolumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CachedIscsiVolumes in the indexer for a given namespace.
func (s cachedIscsiVolumeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CachedIscsiVolume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CachedIscsiVolume))
	})
	return ret, err
}

// Get retrieves the CachedIscsiVolume from the indexer for a given namespace and name.
func (s cachedIscsiVolumeNamespaceLister) Get(name string) (*v1alpha1.CachedIscsiVolume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("cachediscsivolume"), name)
	}
	return obj.(*v1alpha1.CachedIscsiVolume), nil
}
