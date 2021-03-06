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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/chime/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VoiceConnectorLister helps list VoiceConnectors.
// All objects returned here must be treated as read-only.
type VoiceConnectorLister interface {
	// List lists all VoiceConnectors in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VoiceConnector, err error)
	// VoiceConnectors returns an object that can list and get VoiceConnectors.
	VoiceConnectors(namespace string) VoiceConnectorNamespaceLister
	VoiceConnectorListerExpansion
}

// voiceConnectorLister implements the VoiceConnectorLister interface.
type voiceConnectorLister struct {
	indexer cache.Indexer
}

// NewVoiceConnectorLister returns a new VoiceConnectorLister.
func NewVoiceConnectorLister(indexer cache.Indexer) VoiceConnectorLister {
	return &voiceConnectorLister{indexer: indexer}
}

// List lists all VoiceConnectors in the indexer.
func (s *voiceConnectorLister) List(selector labels.Selector) (ret []*v1alpha1.VoiceConnector, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VoiceConnector))
	})
	return ret, err
}

// VoiceConnectors returns an object that can list and get VoiceConnectors.
func (s *voiceConnectorLister) VoiceConnectors(namespace string) VoiceConnectorNamespaceLister {
	return voiceConnectorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VoiceConnectorNamespaceLister helps list and get VoiceConnectors.
// All objects returned here must be treated as read-only.
type VoiceConnectorNamespaceLister interface {
	// List lists all VoiceConnectors in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VoiceConnector, err error)
	// Get retrieves the VoiceConnector from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VoiceConnector, error)
	VoiceConnectorNamespaceListerExpansion
}

// voiceConnectorNamespaceLister implements the VoiceConnectorNamespaceLister
// interface.
type voiceConnectorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VoiceConnectors in the indexer for a given namespace.
func (s voiceConnectorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VoiceConnector, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VoiceConnector))
	})
	return ret, err
}

// Get retrieves the VoiceConnector from the indexer for a given namespace and name.
func (s voiceConnectorNamespaceLister) Get(name string) (*v1alpha1.VoiceConnector, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("voiceconnector"), name)
	}
	return obj.(*v1alpha1.VoiceConnector), nil
}
