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

// VoiceConnectorTerminationCredentialsLister helps list VoiceConnectorTerminationCredentialses.
// All objects returned here must be treated as read-only.
type VoiceConnectorTerminationCredentialsLister interface {
	// List lists all VoiceConnectorTerminationCredentialses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VoiceConnectorTerminationCredentials, err error)
	// VoiceConnectorTerminationCredentialses returns an object that can list and get VoiceConnectorTerminationCredentialses.
	VoiceConnectorTerminationCredentialses(namespace string) VoiceConnectorTerminationCredentialsNamespaceLister
	VoiceConnectorTerminationCredentialsListerExpansion
}

// voiceConnectorTerminationCredentialsLister implements the VoiceConnectorTerminationCredentialsLister interface.
type voiceConnectorTerminationCredentialsLister struct {
	indexer cache.Indexer
}

// NewVoiceConnectorTerminationCredentialsLister returns a new VoiceConnectorTerminationCredentialsLister.
func NewVoiceConnectorTerminationCredentialsLister(indexer cache.Indexer) VoiceConnectorTerminationCredentialsLister {
	return &voiceConnectorTerminationCredentialsLister{indexer: indexer}
}

// List lists all VoiceConnectorTerminationCredentialses in the indexer.
func (s *voiceConnectorTerminationCredentialsLister) List(selector labels.Selector) (ret []*v1alpha1.VoiceConnectorTerminationCredentials, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VoiceConnectorTerminationCredentials))
	})
	return ret, err
}

// VoiceConnectorTerminationCredentialses returns an object that can list and get VoiceConnectorTerminationCredentialses.
func (s *voiceConnectorTerminationCredentialsLister) VoiceConnectorTerminationCredentialses(namespace string) VoiceConnectorTerminationCredentialsNamespaceLister {
	return voiceConnectorTerminationCredentialsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VoiceConnectorTerminationCredentialsNamespaceLister helps list and get VoiceConnectorTerminationCredentialses.
// All objects returned here must be treated as read-only.
type VoiceConnectorTerminationCredentialsNamespaceLister interface {
	// List lists all VoiceConnectorTerminationCredentialses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VoiceConnectorTerminationCredentials, err error)
	// Get retrieves the VoiceConnectorTerminationCredentials from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VoiceConnectorTerminationCredentials, error)
	VoiceConnectorTerminationCredentialsNamespaceListerExpansion
}

// voiceConnectorTerminationCredentialsNamespaceLister implements the VoiceConnectorTerminationCredentialsNamespaceLister
// interface.
type voiceConnectorTerminationCredentialsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VoiceConnectorTerminationCredentialses in the indexer for a given namespace.
func (s voiceConnectorTerminationCredentialsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VoiceConnectorTerminationCredentials, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VoiceConnectorTerminationCredentials))
	})
	return ret, err
}

// Get retrieves the VoiceConnectorTerminationCredentials from the indexer for a given namespace and name.
func (s voiceConnectorTerminationCredentialsNamespaceLister) Get(name string) (*v1alpha1.VoiceConnectorTerminationCredentials, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("voiceconnectorterminationcredentials"), name)
	}
	return obj.(*v1alpha1.VoiceConnectorTerminationCredentials), nil
}
