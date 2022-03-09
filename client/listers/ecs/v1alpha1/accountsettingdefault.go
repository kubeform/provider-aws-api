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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ecs/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AccountSettingDefaultLister helps list AccountSettingDefaults.
// All objects returned here must be treated as read-only.
type AccountSettingDefaultLister interface {
	// List lists all AccountSettingDefaults in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AccountSettingDefault, err error)
	// AccountSettingDefaults returns an object that can list and get AccountSettingDefaults.
	AccountSettingDefaults(namespace string) AccountSettingDefaultNamespaceLister
	AccountSettingDefaultListerExpansion
}

// accountSettingDefaultLister implements the AccountSettingDefaultLister interface.
type accountSettingDefaultLister struct {
	indexer cache.Indexer
}

// NewAccountSettingDefaultLister returns a new AccountSettingDefaultLister.
func NewAccountSettingDefaultLister(indexer cache.Indexer) AccountSettingDefaultLister {
	return &accountSettingDefaultLister{indexer: indexer}
}

// List lists all AccountSettingDefaults in the indexer.
func (s *accountSettingDefaultLister) List(selector labels.Selector) (ret []*v1alpha1.AccountSettingDefault, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AccountSettingDefault))
	})
	return ret, err
}

// AccountSettingDefaults returns an object that can list and get AccountSettingDefaults.
func (s *accountSettingDefaultLister) AccountSettingDefaults(namespace string) AccountSettingDefaultNamespaceLister {
	return accountSettingDefaultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AccountSettingDefaultNamespaceLister helps list and get AccountSettingDefaults.
// All objects returned here must be treated as read-only.
type AccountSettingDefaultNamespaceLister interface {
	// List lists all AccountSettingDefaults in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AccountSettingDefault, err error)
	// Get retrieves the AccountSettingDefault from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AccountSettingDefault, error)
	AccountSettingDefaultNamespaceListerExpansion
}

// accountSettingDefaultNamespaceLister implements the AccountSettingDefaultNamespaceLister
// interface.
type accountSettingDefaultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AccountSettingDefaults in the indexer for a given namespace.
func (s accountSettingDefaultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AccountSettingDefault, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AccountSettingDefault))
	})
	return ret, err
}

// Get retrieves the AccountSettingDefault from the indexer for a given namespace and name.
func (s accountSettingDefaultNamespaceLister) Get(name string) (*v1alpha1.AccountSettingDefault, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("accountsettingdefault"), name)
	}
	return obj.(*v1alpha1.AccountSettingDefault), nil
}
