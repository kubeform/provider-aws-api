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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/secretsmanager/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SecretVersionLister helps list SecretVersions.
// All objects returned here must be treated as read-only.
type SecretVersionLister interface {
	// List lists all SecretVersions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SecretVersion, err error)
	// SecretVersions returns an object that can list and get SecretVersions.
	SecretVersions(namespace string) SecretVersionNamespaceLister
	SecretVersionListerExpansion
}

// secretVersionLister implements the SecretVersionLister interface.
type secretVersionLister struct {
	indexer cache.Indexer
}

// NewSecretVersionLister returns a new SecretVersionLister.
func NewSecretVersionLister(indexer cache.Indexer) SecretVersionLister {
	return &secretVersionLister{indexer: indexer}
}

// List lists all SecretVersions in the indexer.
func (s *secretVersionLister) List(selector labels.Selector) (ret []*v1alpha1.SecretVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretVersion))
	})
	return ret, err
}

// SecretVersions returns an object that can list and get SecretVersions.
func (s *secretVersionLister) SecretVersions(namespace string) SecretVersionNamespaceLister {
	return secretVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SecretVersionNamespaceLister helps list and get SecretVersions.
// All objects returned here must be treated as read-only.
type SecretVersionNamespaceLister interface {
	// List lists all SecretVersions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SecretVersion, err error)
	// Get retrieves the SecretVersion from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SecretVersion, error)
	SecretVersionNamespaceListerExpansion
}

// secretVersionNamespaceLister implements the SecretVersionNamespaceLister
// interface.
type secretVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SecretVersions in the indexer for a given namespace.
func (s secretVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SecretVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SecretVersion))
	})
	return ret, err
}

// Get retrieves the SecretVersion from the indexer for a given namespace and name.
func (s secretVersionNamespaceLister) Get(name string) (*v1alpha1.SecretVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("secretversion"), name)
	}
	return obj.(*v1alpha1.SecretVersion), nil
}
