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
	v1alpha1 "kubeform.dev/provider-aws-api/client/clientset/versioned/typed/cloudfront/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCloudfrontV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCloudfrontV1alpha1) CachePolicies(namespace string) v1alpha1.CachePolicyInterface {
	return &FakeCachePolicies{c, namespace}
}

func (c *FakeCloudfrontV1alpha1) Distributions(namespace string) v1alpha1.DistributionInterface {
	return &FakeDistributions{c, namespace}
}

func (c *FakeCloudfrontV1alpha1) Functions(namespace string) v1alpha1.FunctionInterface {
	return &FakeFunctions{c, namespace}
}

func (c *FakeCloudfrontV1alpha1) KeyGroups(namespace string) v1alpha1.KeyGroupInterface {
	return &FakeKeyGroups{c, namespace}
}

func (c *FakeCloudfrontV1alpha1) OriginAccessIdentities(namespace string) v1alpha1.OriginAccessIdentityInterface {
	return &FakeOriginAccessIdentities{c, namespace}
}

func (c *FakeCloudfrontV1alpha1) OriginRequestPolicies(namespace string) v1alpha1.OriginRequestPolicyInterface {
	return &FakeOriginRequestPolicies{c, namespace}
}

func (c *FakeCloudfrontV1alpha1) PublicKeys(namespace string) v1alpha1.PublicKeyInterface {
	return &FakePublicKeys{c, namespace}
}

func (c *FakeCloudfrontV1alpha1) RealtimeLogConfigs(namespace string) v1alpha1.RealtimeLogConfigInterface {
	return &FakeRealtimeLogConfigs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCloudfrontV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
