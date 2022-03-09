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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "kubeform.dev/provider-aws-api/client/clientset/versioned/typed/ecr/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeEcrV1alpha1 struct {
	*testing.Fake
}

func (c *FakeEcrV1alpha1) LifecyclePolicies(namespace string) v1alpha1.LifecyclePolicyInterface {
	return &FakeLifecyclePolicies{c, namespace}
}

func (c *FakeEcrV1alpha1) PullThroughCacheRules(namespace string) v1alpha1.PullThroughCacheRuleInterface {
	return &FakePullThroughCacheRules{c, namespace}
}

func (c *FakeEcrV1alpha1) RegistryPolicies(namespace string) v1alpha1.RegistryPolicyInterface {
	return &FakeRegistryPolicies{c, namespace}
}

func (c *FakeEcrV1alpha1) RegistryScanningConfigurations(namespace string) v1alpha1.RegistryScanningConfigurationInterface {
	return &FakeRegistryScanningConfigurations{c, namespace}
}

func (c *FakeEcrV1alpha1) ReplicationConfigurations(namespace string) v1alpha1.ReplicationConfigurationInterface {
	return &FakeReplicationConfigurations{c, namespace}
}

func (c *FakeEcrV1alpha1) Repositories(namespace string) v1alpha1.RepositoryInterface {
	return &FakeRepositories{c, namespace}
}

func (c *FakeEcrV1alpha1) RepositoryPolicies(namespace string) v1alpha1.RepositoryPolicyInterface {
	return &FakeRepositoryPolicies{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEcrV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
