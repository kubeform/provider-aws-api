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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/fsx/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OntapVolumeLister helps list OntapVolumes.
// All objects returned here must be treated as read-only.
type OntapVolumeLister interface {
	// List lists all OntapVolumes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OntapVolume, err error)
	// OntapVolumes returns an object that can list and get OntapVolumes.
	OntapVolumes(namespace string) OntapVolumeNamespaceLister
	OntapVolumeListerExpansion
}

// ontapVolumeLister implements the OntapVolumeLister interface.
type ontapVolumeLister struct {
	indexer cache.Indexer
}

// NewOntapVolumeLister returns a new OntapVolumeLister.
func NewOntapVolumeLister(indexer cache.Indexer) OntapVolumeLister {
	return &ontapVolumeLister{indexer: indexer}
}

// List lists all OntapVolumes in the indexer.
func (s *ontapVolumeLister) List(selector labels.Selector) (ret []*v1alpha1.OntapVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OntapVolume))
	})
	return ret, err
}

// OntapVolumes returns an object that can list and get OntapVolumes.
func (s *ontapVolumeLister) OntapVolumes(namespace string) OntapVolumeNamespaceLister {
	return ontapVolumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OntapVolumeNamespaceLister helps list and get OntapVolumes.
// All objects returned here must be treated as read-only.
type OntapVolumeNamespaceLister interface {
	// List lists all OntapVolumes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OntapVolume, err error)
	// Get retrieves the OntapVolume from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OntapVolume, error)
	OntapVolumeNamespaceListerExpansion
}

// ontapVolumeNamespaceLister implements the OntapVolumeNamespaceLister
// interface.
type ontapVolumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OntapVolumes in the indexer for a given namespace.
func (s ontapVolumeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OntapVolume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OntapVolume))
	})
	return ret, err
}

// Get retrieves the OntapVolume from the indexer for a given namespace and name.
func (s ontapVolumeNamespaceLister) Get(name string) (*v1alpha1.OntapVolume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ontapvolume"), name)
	}
	return obj.(*v1alpha1.OntapVolume), nil
}
