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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/signer/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SigningProfileLister helps list SigningProfiles.
// All objects returned here must be treated as read-only.
type SigningProfileLister interface {
	// List lists all SigningProfiles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SigningProfile, err error)
	// SigningProfiles returns an object that can list and get SigningProfiles.
	SigningProfiles(namespace string) SigningProfileNamespaceLister
	SigningProfileListerExpansion
}

// signingProfileLister implements the SigningProfileLister interface.
type signingProfileLister struct {
	indexer cache.Indexer
}

// NewSigningProfileLister returns a new SigningProfileLister.
func NewSigningProfileLister(indexer cache.Indexer) SigningProfileLister {
	return &signingProfileLister{indexer: indexer}
}

// List lists all SigningProfiles in the indexer.
func (s *signingProfileLister) List(selector labels.Selector) (ret []*v1alpha1.SigningProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SigningProfile))
	})
	return ret, err
}

// SigningProfiles returns an object that can list and get SigningProfiles.
func (s *signingProfileLister) SigningProfiles(namespace string) SigningProfileNamespaceLister {
	return signingProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SigningProfileNamespaceLister helps list and get SigningProfiles.
// All objects returned here must be treated as read-only.
type SigningProfileNamespaceLister interface {
	// List lists all SigningProfiles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SigningProfile, err error)
	// Get retrieves the SigningProfile from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SigningProfile, error)
	SigningProfileNamespaceListerExpansion
}

// signingProfileNamespaceLister implements the SigningProfileNamespaceLister
// interface.
type signingProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SigningProfiles in the indexer for a given namespace.
func (s signingProfileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SigningProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SigningProfile))
	})
	return ret, err
}

// Get retrieves the SigningProfile from the indexer for a given namespace and name.
func (s signingProfileNamespaceLister) Get(name string) (*v1alpha1.SigningProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("signingprofile"), name)
	}
	return obj.(*v1alpha1.SigningProfile), nil
}
