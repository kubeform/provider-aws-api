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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/route53/v1alpha1"
)

// HostedZoneDnssecLister helps list HostedZoneDnssecs.
// All objects returned here must be treated as read-only.
type HostedZoneDnssecLister interface {
	// List lists all HostedZoneDnssecs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HostedZoneDnssec, err error)
	// HostedZoneDnssecs returns an object that can list and get HostedZoneDnssecs.
	HostedZoneDnssecs(namespace string) HostedZoneDnssecNamespaceLister
	HostedZoneDnssecListerExpansion
}

// hostedZoneDnssecLister implements the HostedZoneDnssecLister interface.
type hostedZoneDnssecLister struct {
	indexer cache.Indexer
}

// NewHostedZoneDnssecLister returns a new HostedZoneDnssecLister.
func NewHostedZoneDnssecLister(indexer cache.Indexer) HostedZoneDnssecLister {
	return &hostedZoneDnssecLister{indexer: indexer}
}

// List lists all HostedZoneDnssecs in the indexer.
func (s *hostedZoneDnssecLister) List(selector labels.Selector) (ret []*v1alpha1.HostedZoneDnssec, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HostedZoneDnssec))
	})
	return ret, err
}

// HostedZoneDnssecs returns an object that can list and get HostedZoneDnssecs.
func (s *hostedZoneDnssecLister) HostedZoneDnssecs(namespace string) HostedZoneDnssecNamespaceLister {
	return hostedZoneDnssecNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HostedZoneDnssecNamespaceLister helps list and get HostedZoneDnssecs.
// All objects returned here must be treated as read-only.
type HostedZoneDnssecNamespaceLister interface {
	// List lists all HostedZoneDnssecs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HostedZoneDnssec, err error)
	// Get retrieves the HostedZoneDnssec from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HostedZoneDnssec, error)
	HostedZoneDnssecNamespaceListerExpansion
}

// hostedZoneDnssecNamespaceLister implements the HostedZoneDnssecNamespaceLister
// interface.
type hostedZoneDnssecNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HostedZoneDnssecs in the indexer for a given namespace.
func (s hostedZoneDnssecNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HostedZoneDnssec, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HostedZoneDnssec))
	})
	return ret, err
}

// Get retrieves the HostedZoneDnssec from the indexer for a given namespace and name.
func (s hostedZoneDnssecNamespaceLister) Get(name string) (*v1alpha1.HostedZoneDnssec, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hostedzonednssec"), name)
	}
	return obj.(*v1alpha1.HostedZoneDnssec), nil
}
