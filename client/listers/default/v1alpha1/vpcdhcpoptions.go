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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/default/v1alpha1"
)

// VpcDHCPOptionsLister helps list VpcDHCPOptionses.
// All objects returned here must be treated as read-only.
type VpcDHCPOptionsLister interface {
	// List lists all VpcDHCPOptionses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpcDHCPOptions, err error)
	// VpcDHCPOptionses returns an object that can list and get VpcDHCPOptionses.
	VpcDHCPOptionses(namespace string) VpcDHCPOptionsNamespaceLister
	VpcDHCPOptionsListerExpansion
}

// vpcDHCPOptionsLister implements the VpcDHCPOptionsLister interface.
type vpcDHCPOptionsLister struct {
	indexer cache.Indexer
}

// NewVpcDHCPOptionsLister returns a new VpcDHCPOptionsLister.
func NewVpcDHCPOptionsLister(indexer cache.Indexer) VpcDHCPOptionsLister {
	return &vpcDHCPOptionsLister{indexer: indexer}
}

// List lists all VpcDHCPOptionses in the indexer.
func (s *vpcDHCPOptionsLister) List(selector labels.Selector) (ret []*v1alpha1.VpcDHCPOptions, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcDHCPOptions))
	})
	return ret, err
}

// VpcDHCPOptionses returns an object that can list and get VpcDHCPOptionses.
func (s *vpcDHCPOptionsLister) VpcDHCPOptionses(namespace string) VpcDHCPOptionsNamespaceLister {
	return vpcDHCPOptionsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcDHCPOptionsNamespaceLister helps list and get VpcDHCPOptionses.
// All objects returned here must be treated as read-only.
type VpcDHCPOptionsNamespaceLister interface {
	// List lists all VpcDHCPOptionses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpcDHCPOptions, err error)
	// Get retrieves the VpcDHCPOptions from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VpcDHCPOptions, error)
	VpcDHCPOptionsNamespaceListerExpansion
}

// vpcDHCPOptionsNamespaceLister implements the VpcDHCPOptionsNamespaceLister
// interface.
type vpcDHCPOptionsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcDHCPOptionses in the indexer for a given namespace.
func (s vpcDHCPOptionsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcDHCPOptions, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcDHCPOptions))
	})
	return ret, err
}

// Get retrieves the VpcDHCPOptions from the indexer for a given namespace and name.
func (s vpcDHCPOptionsNamespaceLister) Get(name string) (*v1alpha1.VpcDHCPOptions, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcdhcpoptions"), name)
	}
	return obj.(*v1alpha1.VpcDHCPOptions), nil
}
