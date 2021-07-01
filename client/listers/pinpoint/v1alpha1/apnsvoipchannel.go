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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/pinpoint/v1alpha1"
)

// ApnsVoipChannelLister helps list ApnsVoipChannels.
// All objects returned here must be treated as read-only.
type ApnsVoipChannelLister interface {
	// List lists all ApnsVoipChannels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApnsVoipChannel, err error)
	// ApnsVoipChannels returns an object that can list and get ApnsVoipChannels.
	ApnsVoipChannels(namespace string) ApnsVoipChannelNamespaceLister
	ApnsVoipChannelListerExpansion
}

// apnsVoipChannelLister implements the ApnsVoipChannelLister interface.
type apnsVoipChannelLister struct {
	indexer cache.Indexer
}

// NewApnsVoipChannelLister returns a new ApnsVoipChannelLister.
func NewApnsVoipChannelLister(indexer cache.Indexer) ApnsVoipChannelLister {
	return &apnsVoipChannelLister{indexer: indexer}
}

// List lists all ApnsVoipChannels in the indexer.
func (s *apnsVoipChannelLister) List(selector labels.Selector) (ret []*v1alpha1.ApnsVoipChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApnsVoipChannel))
	})
	return ret, err
}

// ApnsVoipChannels returns an object that can list and get ApnsVoipChannels.
func (s *apnsVoipChannelLister) ApnsVoipChannels(namespace string) ApnsVoipChannelNamespaceLister {
	return apnsVoipChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApnsVoipChannelNamespaceLister helps list and get ApnsVoipChannels.
// All objects returned here must be treated as read-only.
type ApnsVoipChannelNamespaceLister interface {
	// List lists all ApnsVoipChannels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApnsVoipChannel, err error)
	// Get retrieves the ApnsVoipChannel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ApnsVoipChannel, error)
	ApnsVoipChannelNamespaceListerExpansion
}

// apnsVoipChannelNamespaceLister implements the ApnsVoipChannelNamespaceLister
// interface.
type apnsVoipChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApnsVoipChannels in the indexer for a given namespace.
func (s apnsVoipChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApnsVoipChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApnsVoipChannel))
	})
	return ret, err
}

// Get retrieves the ApnsVoipChannel from the indexer for a given namespace and name.
func (s apnsVoipChannelNamespaceLister) Get(name string) (*v1alpha1.ApnsVoipChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apnsvoipchannel"), name)
	}
	return obj.(*v1alpha1.ApnsVoipChannel), nil
}
