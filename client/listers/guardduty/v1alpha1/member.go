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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/guardduty/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MemberLister helps list Members.
// All objects returned here must be treated as read-only.
type MemberLister interface {
	// List lists all Members in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Member, err error)
	// Members returns an object that can list and get Members.
	Members(namespace string) MemberNamespaceLister
	MemberListerExpansion
}

// memberLister implements the MemberLister interface.
type memberLister struct {
	indexer cache.Indexer
}

// NewMemberLister returns a new MemberLister.
func NewMemberLister(indexer cache.Indexer) MemberLister {
	return &memberLister{indexer: indexer}
}

// List lists all Members in the indexer.
func (s *memberLister) List(selector labels.Selector) (ret []*v1alpha1.Member, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Member))
	})
	return ret, err
}

// Members returns an object that can list and get Members.
func (s *memberLister) Members(namespace string) MemberNamespaceLister {
	return memberNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MemberNamespaceLister helps list and get Members.
// All objects returned here must be treated as read-only.
type MemberNamespaceLister interface {
	// List lists all Members in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Member, err error)
	// Get retrieves the Member from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Member, error)
	MemberNamespaceListerExpansion
}

// memberNamespaceLister implements the MemberNamespaceLister
// interface.
type memberNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Members in the indexer for a given namespace.
func (s memberNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Member, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Member))
	})
	return ret, err
}

// Get retrieves the Member from the indexer for a given namespace and name.
func (s memberNamespaceLister) Get(name string) (*v1alpha1.Member, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("member"), name)
	}
	return obj.(*v1alpha1.Member), nil
}
