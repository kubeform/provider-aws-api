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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/devicefarm/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// InstanceProfileLister helps list InstanceProfiles.
// All objects returned here must be treated as read-only.
type InstanceProfileLister interface {
	// List lists all InstanceProfiles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceProfile, err error)
	// InstanceProfiles returns an object that can list and get InstanceProfiles.
	InstanceProfiles(namespace string) InstanceProfileNamespaceLister
	InstanceProfileListerExpansion
}

// instanceProfileLister implements the InstanceProfileLister interface.
type instanceProfileLister struct {
	indexer cache.Indexer
}

// NewInstanceProfileLister returns a new InstanceProfileLister.
func NewInstanceProfileLister(indexer cache.Indexer) InstanceProfileLister {
	return &instanceProfileLister{indexer: indexer}
}

// List lists all InstanceProfiles in the indexer.
func (s *instanceProfileLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceProfile))
	})
	return ret, err
}

// InstanceProfiles returns an object that can list and get InstanceProfiles.
func (s *instanceProfileLister) InstanceProfiles(namespace string) InstanceProfileNamespaceLister {
	return instanceProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// InstanceProfileNamespaceLister helps list and get InstanceProfiles.
// All objects returned here must be treated as read-only.
type InstanceProfileNamespaceLister interface {
	// List lists all InstanceProfiles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.InstanceProfile, err error)
	// Get retrieves the InstanceProfile from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.InstanceProfile, error)
	InstanceProfileNamespaceListerExpansion
}

// instanceProfileNamespaceLister implements the InstanceProfileNamespaceLister
// interface.
type instanceProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all InstanceProfiles in the indexer for a given namespace.
func (s instanceProfileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.InstanceProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.InstanceProfile))
	})
	return ret, err
}

// Get retrieves the InstanceProfile from the indexer for a given namespace and name.
func (s instanceProfileNamespaceLister) Get(name string) (*v1alpha1.InstanceProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("instanceprofile"), name)
	}
	return obj.(*v1alpha1.InstanceProfile), nil
}
