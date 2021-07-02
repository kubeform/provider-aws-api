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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/pinpoint/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApnsSandboxChannelLister helps list ApnsSandboxChannels.
// All objects returned here must be treated as read-only.
type ApnsSandboxChannelLister interface {
	// List lists all ApnsSandboxChannels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApnsSandboxChannel, err error)
	// ApnsSandboxChannels returns an object that can list and get ApnsSandboxChannels.
	ApnsSandboxChannels(namespace string) ApnsSandboxChannelNamespaceLister
	ApnsSandboxChannelListerExpansion
}

// apnsSandboxChannelLister implements the ApnsSandboxChannelLister interface.
type apnsSandboxChannelLister struct {
	indexer cache.Indexer
}

// NewApnsSandboxChannelLister returns a new ApnsSandboxChannelLister.
func NewApnsSandboxChannelLister(indexer cache.Indexer) ApnsSandboxChannelLister {
	return &apnsSandboxChannelLister{indexer: indexer}
}

// List lists all ApnsSandboxChannels in the indexer.
func (s *apnsSandboxChannelLister) List(selector labels.Selector) (ret []*v1alpha1.ApnsSandboxChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApnsSandboxChannel))
	})
	return ret, err
}

// ApnsSandboxChannels returns an object that can list and get ApnsSandboxChannels.
func (s *apnsSandboxChannelLister) ApnsSandboxChannels(namespace string) ApnsSandboxChannelNamespaceLister {
	return apnsSandboxChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApnsSandboxChannelNamespaceLister helps list and get ApnsSandboxChannels.
// All objects returned here must be treated as read-only.
type ApnsSandboxChannelNamespaceLister interface {
	// List lists all ApnsSandboxChannels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApnsSandboxChannel, err error)
	// Get retrieves the ApnsSandboxChannel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ApnsSandboxChannel, error)
	ApnsSandboxChannelNamespaceListerExpansion
}

// apnsSandboxChannelNamespaceLister implements the ApnsSandboxChannelNamespaceLister
// interface.
type apnsSandboxChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApnsSandboxChannels in the indexer for a given namespace.
func (s apnsSandboxChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApnsSandboxChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApnsSandboxChannel))
	})
	return ret, err
}

// Get retrieves the ApnsSandboxChannel from the indexer for a given namespace and name.
func (s apnsSandboxChannelNamespaceLister) Get(name string) (*v1alpha1.ApnsSandboxChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apnssandboxchannel"), name)
	}
	return obj.(*v1alpha1.ApnsSandboxChannel), nil
}
