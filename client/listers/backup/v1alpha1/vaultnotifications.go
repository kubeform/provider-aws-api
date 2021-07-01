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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/backup/v1alpha1"
)

// VaultNotificationsLister helps list VaultNotificationses.
// All objects returned here must be treated as read-only.
type VaultNotificationsLister interface {
	// List lists all VaultNotificationses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VaultNotifications, err error)
	// VaultNotificationses returns an object that can list and get VaultNotificationses.
	VaultNotificationses(namespace string) VaultNotificationsNamespaceLister
	VaultNotificationsListerExpansion
}

// vaultNotificationsLister implements the VaultNotificationsLister interface.
type vaultNotificationsLister struct {
	indexer cache.Indexer
}

// NewVaultNotificationsLister returns a new VaultNotificationsLister.
func NewVaultNotificationsLister(indexer cache.Indexer) VaultNotificationsLister {
	return &vaultNotificationsLister{indexer: indexer}
}

// List lists all VaultNotificationses in the indexer.
func (s *vaultNotificationsLister) List(selector labels.Selector) (ret []*v1alpha1.VaultNotifications, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultNotifications))
	})
	return ret, err
}

// VaultNotificationses returns an object that can list and get VaultNotificationses.
func (s *vaultNotificationsLister) VaultNotificationses(namespace string) VaultNotificationsNamespaceLister {
	return vaultNotificationsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VaultNotificationsNamespaceLister helps list and get VaultNotificationses.
// All objects returned here must be treated as read-only.
type VaultNotificationsNamespaceLister interface {
	// List lists all VaultNotificationses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VaultNotifications, err error)
	// Get retrieves the VaultNotifications from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VaultNotifications, error)
	VaultNotificationsNamespaceListerExpansion
}

// vaultNotificationsNamespaceLister implements the VaultNotificationsNamespaceLister
// interface.
type vaultNotificationsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VaultNotificationses in the indexer for a given namespace.
func (s vaultNotificationsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VaultNotifications, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VaultNotifications))
	})
	return ret, err
}

// Get retrieves the VaultNotifications from the indexer for a given namespace and name.
func (s vaultNotificationsNamespaceLister) Get(name string) (*v1alpha1.VaultNotifications, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vaultnotifications"), name)
	}
	return obj.(*v1alpha1.VaultNotifications), nil
}
