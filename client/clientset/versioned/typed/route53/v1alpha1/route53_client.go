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
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/provider-aws-api/apis/route53/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"
)

type Route53V1alpha1Interface interface {
	RESTClient() rest.Interface
	DelegationSetsGetter
	HealthChecksGetter
	HostedZoneDnssecsGetter
	KeySigningKeysGetter
	QueryLogsGetter
	RecordsGetter
	ResolverDnssecConfigsGetter
	ResolverEndpointsGetter
	ResolverFirewallConfigsGetter
	ResolverFirewallDomainListsGetter
	ResolverFirewallRulesGetter
	ResolverFirewallRuleGroupsGetter
	ResolverFirewallRuleGroupAssociationsGetter
	ResolverQueryLogConfigsGetter
	ResolverQueryLogConfigAssociationsGetter
	ResolverRulesGetter
	ResolverRuleAssociationsGetter
	VpcAssociationAuthorizationsGetter
	ZonesGetter
	ZoneAssociationsGetter
}

// Route53V1alpha1Client is used to interact with features provided by the route53.aws.kubeform.com group.
type Route53V1alpha1Client struct {
	restClient rest.Interface
}

func (c *Route53V1alpha1Client) DelegationSets(namespace string) DelegationSetInterface {
	return newDelegationSets(c, namespace)
}

func (c *Route53V1alpha1Client) HealthChecks(namespace string) HealthCheckInterface {
	return newHealthChecks(c, namespace)
}

func (c *Route53V1alpha1Client) HostedZoneDnssecs(namespace string) HostedZoneDnssecInterface {
	return newHostedZoneDnssecs(c, namespace)
}

func (c *Route53V1alpha1Client) KeySigningKeys(namespace string) KeySigningKeyInterface {
	return newKeySigningKeys(c, namespace)
}

func (c *Route53V1alpha1Client) QueryLogs(namespace string) QueryLogInterface {
	return newQueryLogs(c, namespace)
}

func (c *Route53V1alpha1Client) Records(namespace string) RecordInterface {
	return newRecords(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverDnssecConfigs(namespace string) ResolverDnssecConfigInterface {
	return newResolverDnssecConfigs(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverEndpoints(namespace string) ResolverEndpointInterface {
	return newResolverEndpoints(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverFirewallConfigs(namespace string) ResolverFirewallConfigInterface {
	return newResolverFirewallConfigs(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverFirewallDomainLists(namespace string) ResolverFirewallDomainListInterface {
	return newResolverFirewallDomainLists(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverFirewallRules(namespace string) ResolverFirewallRuleInterface {
	return newResolverFirewallRules(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverFirewallRuleGroups(namespace string) ResolverFirewallRuleGroupInterface {
	return newResolverFirewallRuleGroups(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverFirewallRuleGroupAssociations(namespace string) ResolverFirewallRuleGroupAssociationInterface {
	return newResolverFirewallRuleGroupAssociations(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverQueryLogConfigs(namespace string) ResolverQueryLogConfigInterface {
	return newResolverQueryLogConfigs(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverQueryLogConfigAssociations(namespace string) ResolverQueryLogConfigAssociationInterface {
	return newResolverQueryLogConfigAssociations(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverRules(namespace string) ResolverRuleInterface {
	return newResolverRules(c, namespace)
}

func (c *Route53V1alpha1Client) ResolverRuleAssociations(namespace string) ResolverRuleAssociationInterface {
	return newResolverRuleAssociations(c, namespace)
}

func (c *Route53V1alpha1Client) VpcAssociationAuthorizations(namespace string) VpcAssociationAuthorizationInterface {
	return newVpcAssociationAuthorizations(c, namespace)
}

func (c *Route53V1alpha1Client) Zones(namespace string) ZoneInterface {
	return newZones(c, namespace)
}

func (c *Route53V1alpha1Client) ZoneAssociations(namespace string) ZoneAssociationInterface {
	return newZoneAssociations(c, namespace)
}

// NewForConfig creates a new Route53V1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*Route53V1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &Route53V1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new Route53V1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Route53V1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new Route53V1alpha1Client for the given RESTClient.
func New(c rest.Interface) *Route53V1alpha1Client {
	return &Route53V1alpha1Client{c}
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
func (c *Route53V1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
