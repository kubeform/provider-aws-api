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
	v1alpha1 "kubeform.dev/provider-aws-api/client/clientset/versioned/typed/gamelift/v1alpha1"
)

type FakeGameliftV1alpha1 struct {
	*testing.Fake
}

func (c *FakeGameliftV1alpha1) Aliases(namespace string) v1alpha1.AliasInterface {
	return &FakeAliases{c, namespace}
}

func (c *FakeGameliftV1alpha1) Builds(namespace string) v1alpha1.BuildInterface {
	return &FakeBuilds{c, namespace}
}

func (c *FakeGameliftV1alpha1) Fleets(namespace string) v1alpha1.FleetInterface {
	return &FakeFleets{c, namespace}
}

func (c *FakeGameliftV1alpha1) GameSessionQueues(namespace string) v1alpha1.GameSessionQueueInterface {
	return &FakeGameSessionQueues{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeGameliftV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
