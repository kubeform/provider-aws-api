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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubeform.dev/provider-aws-api/client/clientset/versioned/typed/secretsmanager/v1alpha1"
)

type FakeSecretsmanagerV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSecretsmanagerV1alpha1) Secrets(namespace string) v1alpha1.SecretInterface {
	return &FakeSecrets{c, namespace}
}

func (c *FakeSecretsmanagerV1alpha1) SecretPolicies(namespace string) v1alpha1.SecretPolicyInterface {
	return &FakeSecretPolicies{c, namespace}
}

func (c *FakeSecretsmanagerV1alpha1) SecretRotations(namespace string) v1alpha1.SecretRotationInterface {
	return &FakeSecretRotations{c, namespace}
}

func (c *FakeSecretsmanagerV1alpha1) SecretVersions(namespace string) v1alpha1.SecretVersionInterface {
	return &FakeSecretVersions{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSecretsmanagerV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
