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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/route53/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResolverQueryLogConfigLister helps list ResolverQueryLogConfigs.
// All objects returned here must be treated as read-only.
type ResolverQueryLogConfigLister interface {
	// List lists all ResolverQueryLogConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResolverQueryLogConfig, err error)
	// ResolverQueryLogConfigs returns an object that can list and get ResolverQueryLogConfigs.
	ResolverQueryLogConfigs(namespace string) ResolverQueryLogConfigNamespaceLister
	ResolverQueryLogConfigListerExpansion
}

// resolverQueryLogConfigLister implements the ResolverQueryLogConfigLister interface.
type resolverQueryLogConfigLister struct {
	indexer cache.Indexer
}

// NewResolverQueryLogConfigLister returns a new ResolverQueryLogConfigLister.
func NewResolverQueryLogConfigLister(indexer cache.Indexer) ResolverQueryLogConfigLister {
	return &resolverQueryLogConfigLister{indexer: indexer}
}

// List lists all ResolverQueryLogConfigs in the indexer.
func (s *resolverQueryLogConfigLister) List(selector labels.Selector) (ret []*v1alpha1.ResolverQueryLogConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResolverQueryLogConfig))
	})
	return ret, err
}

// ResolverQueryLogConfigs returns an object that can list and get ResolverQueryLogConfigs.
func (s *resolverQueryLogConfigLister) ResolverQueryLogConfigs(namespace string) ResolverQueryLogConfigNamespaceLister {
	return resolverQueryLogConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResolverQueryLogConfigNamespaceLister helps list and get ResolverQueryLogConfigs.
// All objects returned here must be treated as read-only.
type ResolverQueryLogConfigNamespaceLister interface {
	// List lists all ResolverQueryLogConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ResolverQueryLogConfig, err error)
	// Get retrieves the ResolverQueryLogConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ResolverQueryLogConfig, error)
	ResolverQueryLogConfigNamespaceListerExpansion
}

// resolverQueryLogConfigNamespaceLister implements the ResolverQueryLogConfigNamespaceLister
// interface.
type resolverQueryLogConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResolverQueryLogConfigs in the indexer for a given namespace.
func (s resolverQueryLogConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResolverQueryLogConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResolverQueryLogConfig))
	})
	return ret, err
}

// Get retrieves the ResolverQueryLogConfig from the indexer for a given namespace and name.
func (s resolverQueryLogConfigNamespaceLister) Get(name string) (*v1alpha1.ResolverQueryLogConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resolverquerylogconfig"), name)
	}
	return obj.(*v1alpha1.ResolverQueryLogConfig), nil
}
