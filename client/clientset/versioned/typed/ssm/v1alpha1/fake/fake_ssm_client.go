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
	v1alpha1 "kubeform.dev/provider-aws-api/client/clientset/versioned/typed/ssm/v1alpha1"
)

type FakeSsmV1alpha1 struct {
	*testing.Fake
}

func (c *FakeSsmV1alpha1) Activations(namespace string) v1alpha1.ActivationInterface {
	return &FakeActivations{c, namespace}
}

func (c *FakeSsmV1alpha1) Associations(namespace string) v1alpha1.AssociationInterface {
	return &FakeAssociations{c, namespace}
}

func (c *FakeSsmV1alpha1) Documents(namespace string) v1alpha1.DocumentInterface {
	return &FakeDocuments{c, namespace}
}

func (c *FakeSsmV1alpha1) MaintenanceWindows(namespace string) v1alpha1.MaintenanceWindowInterface {
	return &FakeMaintenanceWindows{c, namespace}
}

func (c *FakeSsmV1alpha1) MaintenanceWindowTargets(namespace string) v1alpha1.MaintenanceWindowTargetInterface {
	return &FakeMaintenanceWindowTargets{c, namespace}
}

func (c *FakeSsmV1alpha1) MaintenanceWindowTasks(namespace string) v1alpha1.MaintenanceWindowTaskInterface {
	return &FakeMaintenanceWindowTasks{c, namespace}
}

func (c *FakeSsmV1alpha1) Parameters(namespace string) v1alpha1.ParameterInterface {
	return &FakeParameters{c, namespace}
}

func (c *FakeSsmV1alpha1) PatchBaselines(namespace string) v1alpha1.PatchBaselineInterface {
	return &FakePatchBaselines{c, namespace}
}

func (c *FakeSsmV1alpha1) PatchGroups(namespace string) v1alpha1.PatchGroupInterface {
	return &FakePatchGroups{c, namespace}
}

func (c *FakeSsmV1alpha1) ResourceDataSyncs(namespace string) v1alpha1.ResourceDataSyncInterface {
	return &FakeResourceDataSyncs{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSsmV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
