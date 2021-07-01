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
	v1alpha1 "kubeform.dev/provider-aws-api/client/clientset/versioned/typed/glue/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeGlueV1alpha1 struct {
	*testing.Fake
}

func (c *FakeGlueV1alpha1) CatalogDatabases(namespace string) v1alpha1.CatalogDatabaseInterface {
	return &FakeCatalogDatabases{c, namespace}
}

func (c *FakeGlueV1alpha1) CatalogTables(namespace string) v1alpha1.CatalogTableInterface {
	return &FakeCatalogTables{c, namespace}
}

func (c *FakeGlueV1alpha1) Classifiers(namespace string) v1alpha1.ClassifierInterface {
	return &FakeClassifiers{c, namespace}
}

func (c *FakeGlueV1alpha1) Connections(namespace string) v1alpha1.ConnectionInterface {
	return &FakeConnections{c, namespace}
}

func (c *FakeGlueV1alpha1) Crawlers(namespace string) v1alpha1.CrawlerInterface {
	return &FakeCrawlers{c, namespace}
}

func (c *FakeGlueV1alpha1) DataCatalogEncryptionSettingses(namespace string) v1alpha1.DataCatalogEncryptionSettingsInterface {
	return &FakeDataCatalogEncryptionSettingses{c, namespace}
}

func (c *FakeGlueV1alpha1) DevEndpoints(namespace string) v1alpha1.DevEndpointInterface {
	return &FakeDevEndpoints{c, namespace}
}

func (c *FakeGlueV1alpha1) Jobs(namespace string) v1alpha1.JobInterface {
	return &FakeJobs{c, namespace}
}

func (c *FakeGlueV1alpha1) MlTransforms(namespace string) v1alpha1.MlTransformInterface {
	return &FakeMlTransforms{c, namespace}
}

func (c *FakeGlueV1alpha1) Partitions(namespace string) v1alpha1.PartitionInterface {
	return &FakePartitions{c, namespace}
}

func (c *FakeGlueV1alpha1) Registries(namespace string) v1alpha1.RegistryInterface {
	return &FakeRegistries{c, namespace}
}

func (c *FakeGlueV1alpha1) ResourcePolicies(namespace string) v1alpha1.ResourcePolicyInterface {
	return &FakeResourcePolicies{c, namespace}
}

func (c *FakeGlueV1alpha1) Schemas(namespace string) v1alpha1.SchemaInterface {
	return &FakeSchemas{c, namespace}
}

func (c *FakeGlueV1alpha1) SecurityConfigurations(namespace string) v1alpha1.SecurityConfigurationInterface {
	return &FakeSecurityConfigurations{c, namespace}
}

func (c *FakeGlueV1alpha1) Triggers(namespace string) v1alpha1.TriggerInterface {
	return &FakeTriggers{c, namespace}
}

func (c *FakeGlueV1alpha1) UserDefinedFunctions(namespace string) v1alpha1.UserDefinedFunctionInterface {
	return &FakeUserDefinedFunctions{c, namespace}
}

func (c *FakeGlueV1alpha1) Workflows(namespace string) v1alpha1.WorkflowInterface {
	return &FakeWorkflows{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeGlueV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
