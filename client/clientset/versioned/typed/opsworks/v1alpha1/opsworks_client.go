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

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-aws-api/apis/opsworks/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type OpsworksV1alpha1Interface interface {
	RESTClient() rest.Interface
	ApplicationsGetter
	CustomLayersGetter
	GangliaLayersGetter
	HaproxyLayersGetter
	InstancesGetter
	JavaAppLayersGetter
	MemcachedLayersGetter
	MysqlLayersGetter
	NodejsAppLayersGetter
	PermissionsGetter
	PhpAppLayersGetter
	RailsAppLayersGetter
	RdsDbInstancesGetter
	StacksGetter
	StaticWebLayersGetter
	UserProfilesGetter
}

// OpsworksV1alpha1Client is used to interact with features provided by the opsworks.aws.kubeform.com group.
type OpsworksV1alpha1Client struct {
	restClient rest.Interface
}

func (c *OpsworksV1alpha1Client) Applications(namespace string) ApplicationInterface {
	return newApplications(c, namespace)
}

func (c *OpsworksV1alpha1Client) CustomLayers(namespace string) CustomLayerInterface {
	return newCustomLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) GangliaLayers(namespace string) GangliaLayerInterface {
	return newGangliaLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) HaproxyLayers(namespace string) HaproxyLayerInterface {
	return newHaproxyLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) Instances(namespace string) InstanceInterface {
	return newInstances(c, namespace)
}

func (c *OpsworksV1alpha1Client) JavaAppLayers(namespace string) JavaAppLayerInterface {
	return newJavaAppLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) MemcachedLayers(namespace string) MemcachedLayerInterface {
	return newMemcachedLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) MysqlLayers(namespace string) MysqlLayerInterface {
	return newMysqlLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) NodejsAppLayers(namespace string) NodejsAppLayerInterface {
	return newNodejsAppLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) Permissions(namespace string) PermissionInterface {
	return newPermissions(c, namespace)
}

func (c *OpsworksV1alpha1Client) PhpAppLayers(namespace string) PhpAppLayerInterface {
	return newPhpAppLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) RailsAppLayers(namespace string) RailsAppLayerInterface {
	return newRailsAppLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) RdsDbInstances(namespace string) RdsDbInstanceInterface {
	return newRdsDbInstances(c, namespace)
}

func (c *OpsworksV1alpha1Client) Stacks(namespace string) StackInterface {
	return newStacks(c, namespace)
}

func (c *OpsworksV1alpha1Client) StaticWebLayers(namespace string) StaticWebLayerInterface {
	return newStaticWebLayers(c, namespace)
}

func (c *OpsworksV1alpha1Client) UserProfiles(namespace string) UserProfileInterface {
	return newUserProfiles(c, namespace)
}

// NewForConfig creates a new OpsworksV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*OpsworksV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OpsworksV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new OpsworksV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *OpsworksV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OpsworksV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *OpsworksV1alpha1Client {
	return &OpsworksV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *OpsworksV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
