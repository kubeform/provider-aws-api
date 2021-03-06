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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dax/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SubnetGroupLister helps list SubnetGroups.
// All objects returned here must be treated as read-only.
type SubnetGroupLister interface {
	// List lists all SubnetGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SubnetGroup, err error)
	// SubnetGroups returns an object that can list and get SubnetGroups.
	SubnetGroups(namespace string) SubnetGroupNamespaceLister
	SubnetGroupListerExpansion
}

// subnetGroupLister implements the SubnetGroupLister interface.
type subnetGroupLister struct {
	indexer cache.Indexer
}

// NewSubnetGroupLister returns a new SubnetGroupLister.
func NewSubnetGroupLister(indexer cache.Indexer) SubnetGroupLister {
	return &subnetGroupLister{indexer: indexer}
}

// List lists all SubnetGroups in the indexer.
func (s *subnetGroupLister) List(selector labels.Selector) (ret []*v1alpha1.SubnetGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubnetGroup))
	})
	return ret, err
}

// SubnetGroups returns an object that can list and get SubnetGroups.
func (s *subnetGroupLister) SubnetGroups(namespace string) SubnetGroupNamespaceLister {
	return subnetGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubnetGroupNamespaceLister helps list and get SubnetGroups.
// All objects returned here must be treated as read-only.
type SubnetGroupNamespaceLister interface {
	// List lists all SubnetGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SubnetGroup, err error)
	// Get retrieves the SubnetGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SubnetGroup, error)
	SubnetGroupNamespaceListerExpansion
}

// subnetGroupNamespaceLister implements the SubnetGroupNamespaceLister
// interface.
type subnetGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SubnetGroups in the indexer for a given namespace.
func (s subnetGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SubnetGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubnetGroup))
	})
	return ret, err
}

// Get retrieves the SubnetGroup from the indexer for a given namespace and name.
func (s subnetGroupNamespaceLister) Get(name string) (*v1alpha1.SubnetGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("subnetgroup"), name)
	}
	return obj.(*v1alpha1.SubnetGroup), nil
}
