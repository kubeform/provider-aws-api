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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ec2/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type Ec2V1alpha1Interface interface {
	RESTClient() rest.Interface
	AvailabilityZoneGroupsGetter
	CapacityReservationsGetter
	CarrierGatewaysGetter
	ClientVPNAuthorizationRulesGetter
	ClientVPNEndpointsGetter
	ClientVPNNetworkAssociationsGetter
	ClientVPNRoutesGetter
	FleetsGetter
	HostsGetter
	LocalGatewayRoutesGetter
	LocalGatewayRouteTableVpcAssociationsGetter
	ManagedPrefixListsGetter
	ManagedPrefixListEntriesGetter
	NetworkInsightsPathsGetter
	SerialConsoleAccessesGetter
	SubnetCIDRReservationsGetter
	TagsGetter
	TrafficMirrorFiltersGetter
	TrafficMirrorFilterRulesGetter
	TrafficMirrorSessionsGetter
	TrafficMirrorTargetsGetter
	TransitGatewaysGetter
	TransitGatewayConnectsGetter
	TransitGatewayConnectPeersGetter
	TransitGatewayMulticastDomainsGetter
	TransitGatewayMulticastDomainAssociationsGetter
	TransitGatewayMulticastGroupMembersGetter
	TransitGatewayMulticastGroupSourcesGetter
	TransitGatewayPeeringAttachmentsGetter
	TransitGatewayPeeringAttachmentAcceptersGetter
	TransitGatewayPrefixListReferencesGetter
	TransitGatewayRoutesGetter
	TransitGatewayRouteTablesGetter
	TransitGatewayRouteTableAssociationsGetter
	TransitGatewayRouteTablePropagationsGetter
	TransitGatewayVpcAttachmentsGetter
	TransitGatewayVpcAttachmentAcceptersGetter
}

// Ec2V1alpha1Client is used to interact with features provided by the ec2.aws.kubeform.com group.
type Ec2V1alpha1Client struct {
	restClient rest.Interface
}

func (c *Ec2V1alpha1Client) AvailabilityZoneGroups(namespace string) AvailabilityZoneGroupInterface {
	return newAvailabilityZoneGroups(c, namespace)
}

func (c *Ec2V1alpha1Client) CapacityReservations(namespace string) CapacityReservationInterface {
	return newCapacityReservations(c, namespace)
}

func (c *Ec2V1alpha1Client) CarrierGateways(namespace string) CarrierGatewayInterface {
	return newCarrierGateways(c, namespace)
}

func (c *Ec2V1alpha1Client) ClientVPNAuthorizationRules(namespace string) ClientVPNAuthorizationRuleInterface {
	return newClientVPNAuthorizationRules(c, namespace)
}

func (c *Ec2V1alpha1Client) ClientVPNEndpoints(namespace string) ClientVPNEndpointInterface {
	return newClientVPNEndpoints(c, namespace)
}

func (c *Ec2V1alpha1Client) ClientVPNNetworkAssociations(namespace string) ClientVPNNetworkAssociationInterface {
	return newClientVPNNetworkAssociations(c, namespace)
}

func (c *Ec2V1alpha1Client) ClientVPNRoutes(namespace string) ClientVPNRouteInterface {
	return newClientVPNRoutes(c, namespace)
}

func (c *Ec2V1alpha1Client) Fleets(namespace string) FleetInterface {
	return newFleets(c, namespace)
}

func (c *Ec2V1alpha1Client) Hosts(namespace string) HostInterface {
	return newHosts(c, namespace)
}

func (c *Ec2V1alpha1Client) LocalGatewayRoutes(namespace string) LocalGatewayRouteInterface {
	return newLocalGatewayRoutes(c, namespace)
}

func (c *Ec2V1alpha1Client) LocalGatewayRouteTableVpcAssociations(namespace string) LocalGatewayRouteTableVpcAssociationInterface {
	return newLocalGatewayRouteTableVpcAssociations(c, namespace)
}

func (c *Ec2V1alpha1Client) ManagedPrefixLists(namespace string) ManagedPrefixListInterface {
	return newManagedPrefixLists(c, namespace)
}

func (c *Ec2V1alpha1Client) ManagedPrefixListEntries(namespace string) ManagedPrefixListEntryInterface {
	return newManagedPrefixListEntries(c, namespace)
}

func (c *Ec2V1alpha1Client) NetworkInsightsPaths(namespace string) NetworkInsightsPathInterface {
	return newNetworkInsightsPaths(c, namespace)
}

func (c *Ec2V1alpha1Client) SerialConsoleAccesses(namespace string) SerialConsoleAccessInterface {
	return newSerialConsoleAccesses(c, namespace)
}

func (c *Ec2V1alpha1Client) SubnetCIDRReservations(namespace string) SubnetCIDRReservationInterface {
	return newSubnetCIDRReservations(c, namespace)
}

func (c *Ec2V1alpha1Client) Tags(namespace string) TagInterface {
	return newTags(c, namespace)
}

func (c *Ec2V1alpha1Client) TrafficMirrorFilters(namespace string) TrafficMirrorFilterInterface {
	return newTrafficMirrorFilters(c, namespace)
}

func (c *Ec2V1alpha1Client) TrafficMirrorFilterRules(namespace string) TrafficMirrorFilterRuleInterface {
	return newTrafficMirrorFilterRules(c, namespace)
}

func (c *Ec2V1alpha1Client) TrafficMirrorSessions(namespace string) TrafficMirrorSessionInterface {
	return newTrafficMirrorSessions(c, namespace)
}

func (c *Ec2V1alpha1Client) TrafficMirrorTargets(namespace string) TrafficMirrorTargetInterface {
	return newTrafficMirrorTargets(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGateways(namespace string) TransitGatewayInterface {
	return newTransitGateways(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayConnects(namespace string) TransitGatewayConnectInterface {
	return newTransitGatewayConnects(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayConnectPeers(namespace string) TransitGatewayConnectPeerInterface {
	return newTransitGatewayConnectPeers(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayMulticastDomains(namespace string) TransitGatewayMulticastDomainInterface {
	return newTransitGatewayMulticastDomains(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayMulticastDomainAssociations(namespace string) TransitGatewayMulticastDomainAssociationInterface {
	return newTransitGatewayMulticastDomainAssociations(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayMulticastGroupMembers(namespace string) TransitGatewayMulticastGroupMemberInterface {
	return newTransitGatewayMulticastGroupMembers(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayMulticastGroupSources(namespace string) TransitGatewayMulticastGroupSourceInterface {
	return newTransitGatewayMulticastGroupSources(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayPeeringAttachments(namespace string) TransitGatewayPeeringAttachmentInterface {
	return newTransitGatewayPeeringAttachments(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayPeeringAttachmentAccepters(namespace string) TransitGatewayPeeringAttachmentAccepterInterface {
	return newTransitGatewayPeeringAttachmentAccepters(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayPrefixListReferences(namespace string) TransitGatewayPrefixListReferenceInterface {
	return newTransitGatewayPrefixListReferences(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayRoutes(namespace string) TransitGatewayRouteInterface {
	return newTransitGatewayRoutes(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayRouteTables(namespace string) TransitGatewayRouteTableInterface {
	return newTransitGatewayRouteTables(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayRouteTableAssociations(namespace string) TransitGatewayRouteTableAssociationInterface {
	return newTransitGatewayRouteTableAssociations(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayRouteTablePropagations(namespace string) TransitGatewayRouteTablePropagationInterface {
	return newTransitGatewayRouteTablePropagations(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayVpcAttachments(namespace string) TransitGatewayVpcAttachmentInterface {
	return newTransitGatewayVpcAttachments(c, namespace)
}

func (c *Ec2V1alpha1Client) TransitGatewayVpcAttachmentAccepters(namespace string) TransitGatewayVpcAttachmentAccepterInterface {
	return newTransitGatewayVpcAttachmentAccepters(c, namespace)
}

// NewForConfig creates a new Ec2V1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*Ec2V1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &Ec2V1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new Ec2V1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Ec2V1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new Ec2V1alpha1Client for the given RESTClient.
func New(c rest.Interface) *Ec2V1alpha1Client {
	return &Ec2V1alpha1Client{c}
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
func (c *Ec2V1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
