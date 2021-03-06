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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/config/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type ConfigV1alpha1Interface interface {
	RESTClient() rest.Interface
	AggregateAuthorizationsGetter
	ConfigRulesGetter
	ConfigurationAggregatorsGetter
	ConfigurationRecordersGetter
	ConfigurationRecorderStatusesGetter
	ConformancePacksGetter
	DeliveryChannelsGetter
	OrganizationConformancePacksGetter
	OrganizationCustomRulesGetter
	OrganizationManagedRulesGetter
	RemediationConfigurationsGetter
}

// ConfigV1alpha1Client is used to interact with features provided by the config.aws.kubeform.com group.
type ConfigV1alpha1Client struct {
	restClient rest.Interface
}

func (c *ConfigV1alpha1Client) AggregateAuthorizations(namespace string) AggregateAuthorizationInterface {
	return newAggregateAuthorizations(c, namespace)
}

func (c *ConfigV1alpha1Client) ConfigRules(namespace string) ConfigRuleInterface {
	return newConfigRules(c, namespace)
}

func (c *ConfigV1alpha1Client) ConfigurationAggregators(namespace string) ConfigurationAggregatorInterface {
	return newConfigurationAggregators(c, namespace)
}

func (c *ConfigV1alpha1Client) ConfigurationRecorders(namespace string) ConfigurationRecorderInterface {
	return newConfigurationRecorders(c, namespace)
}

func (c *ConfigV1alpha1Client) ConfigurationRecorderStatuses(namespace string) ConfigurationRecorderStatusInterface {
	return newConfigurationRecorderStatuses(c, namespace)
}

func (c *ConfigV1alpha1Client) ConformancePacks(namespace string) ConformancePackInterface {
	return newConformancePacks(c, namespace)
}

func (c *ConfigV1alpha1Client) DeliveryChannels(namespace string) DeliveryChannelInterface {
	return newDeliveryChannels(c, namespace)
}

func (c *ConfigV1alpha1Client) OrganizationConformancePacks(namespace string) OrganizationConformancePackInterface {
	return newOrganizationConformancePacks(c, namespace)
}

func (c *ConfigV1alpha1Client) OrganizationCustomRules(namespace string) OrganizationCustomRuleInterface {
	return newOrganizationCustomRules(c, namespace)
}

func (c *ConfigV1alpha1Client) OrganizationManagedRules(namespace string) OrganizationManagedRuleInterface {
	return newOrganizationManagedRules(c, namespace)
}

func (c *ConfigV1alpha1Client) RemediationConfigurations(namespace string) RemediationConfigurationInterface {
	return newRemediationConfigurations(c, namespace)
}

// NewForConfig creates a new ConfigV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*ConfigV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &ConfigV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new ConfigV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ConfigV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ConfigV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *ConfigV1alpha1Client {
	return &ConfigV1alpha1Client{c}
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
func (c *ConfigV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
