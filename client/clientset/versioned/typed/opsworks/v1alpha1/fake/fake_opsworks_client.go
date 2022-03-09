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
	v1alpha1 "kubeform.dev/provider-aws-api/client/clientset/versioned/typed/opsworks/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeOpsworksV1alpha1 struct {
	*testing.Fake
}

func (c *FakeOpsworksV1alpha1) Applications(namespace string) v1alpha1.ApplicationInterface {
	return &FakeApplications{c, namespace}
}

func (c *FakeOpsworksV1alpha1) CustomLayers(namespace string) v1alpha1.CustomLayerInterface {
	return &FakeCustomLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) EcsClusterLayers(namespace string) v1alpha1.EcsClusterLayerInterface {
	return &FakeEcsClusterLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) GangliaLayers(namespace string) v1alpha1.GangliaLayerInterface {
	return &FakeGangliaLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) HaproxyLayers(namespace string) v1alpha1.HaproxyLayerInterface {
	return &FakeHaproxyLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) Instances(namespace string) v1alpha1.InstanceInterface {
	return &FakeInstances{c, namespace}
}

func (c *FakeOpsworksV1alpha1) JavaAppLayers(namespace string) v1alpha1.JavaAppLayerInterface {
	return &FakeJavaAppLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) MemcachedLayers(namespace string) v1alpha1.MemcachedLayerInterface {
	return &FakeMemcachedLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) MysqlLayers(namespace string) v1alpha1.MysqlLayerInterface {
	return &FakeMysqlLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) NodejsAppLayers(namespace string) v1alpha1.NodejsAppLayerInterface {
	return &FakeNodejsAppLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) Permissions(namespace string) v1alpha1.PermissionInterface {
	return &FakePermissions{c, namespace}
}

func (c *FakeOpsworksV1alpha1) PhpAppLayers(namespace string) v1alpha1.PhpAppLayerInterface {
	return &FakePhpAppLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) RailsAppLayers(namespace string) v1alpha1.RailsAppLayerInterface {
	return &FakeRailsAppLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) RdsDbInstances(namespace string) v1alpha1.RdsDbInstanceInterface {
	return &FakeRdsDbInstances{c, namespace}
}

func (c *FakeOpsworksV1alpha1) Stacks(namespace string) v1alpha1.StackInterface {
	return &FakeStacks{c, namespace}
}

func (c *FakeOpsworksV1alpha1) StaticWebLayers(namespace string) v1alpha1.StaticWebLayerInterface {
	return &FakeStaticWebLayers{c, namespace}
}

func (c *FakeOpsworksV1alpha1) UserProfiles(namespace string) v1alpha1.UserProfileInterface {
	return &FakeUserProfiles{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeOpsworksV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
