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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/dx/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type DxV1alpha1Interface interface {
	RESTClient() rest.Interface
	BgpPeersGetter
	ConnectionsGetter
	ConnectionAssociationsGetter
	ConnectionConfirmationsGetter
	GatewaysGetter
	GatewayAssociationsGetter
	GatewayAssociationProposalsGetter
	HostedConnectionsGetter
	HostedPrivateVirtualInterfacesGetter
	HostedPrivateVirtualInterfaceAcceptersGetter
	HostedPublicVirtualInterfacesGetter
	HostedPublicVirtualInterfaceAcceptersGetter
	HostedTransitVirtualInterfacesGetter
	HostedTransitVirtualInterfaceAcceptersGetter
	LagsGetter
	PrivateVirtualInterfacesGetter
	PublicVirtualInterfacesGetter
	TransitVirtualInterfacesGetter
}

// DxV1alpha1Client is used to interact with features provided by the dx.aws.kubeform.com group.
type DxV1alpha1Client struct {
	restClient rest.Interface
}

func (c *DxV1alpha1Client) BgpPeers(namespace string) BgpPeerInterface {
	return newBgpPeers(c, namespace)
}

func (c *DxV1alpha1Client) Connections(namespace string) ConnectionInterface {
	return newConnections(c, namespace)
}

func (c *DxV1alpha1Client) ConnectionAssociations(namespace string) ConnectionAssociationInterface {
	return newConnectionAssociations(c, namespace)
}

func (c *DxV1alpha1Client) ConnectionConfirmations(namespace string) ConnectionConfirmationInterface {
	return newConnectionConfirmations(c, namespace)
}

func (c *DxV1alpha1Client) Gateways(namespace string) GatewayInterface {
	return newGateways(c, namespace)
}

func (c *DxV1alpha1Client) GatewayAssociations(namespace string) GatewayAssociationInterface {
	return newGatewayAssociations(c, namespace)
}

func (c *DxV1alpha1Client) GatewayAssociationProposals(namespace string) GatewayAssociationProposalInterface {
	return newGatewayAssociationProposals(c, namespace)
}

func (c *DxV1alpha1Client) HostedConnections(namespace string) HostedConnectionInterface {
	return newHostedConnections(c, namespace)
}

func (c *DxV1alpha1Client) HostedPrivateVirtualInterfaces(namespace string) HostedPrivateVirtualInterfaceInterface {
	return newHostedPrivateVirtualInterfaces(c, namespace)
}

func (c *DxV1alpha1Client) HostedPrivateVirtualInterfaceAccepters(namespace string) HostedPrivateVirtualInterfaceAccepterInterface {
	return newHostedPrivateVirtualInterfaceAccepters(c, namespace)
}

func (c *DxV1alpha1Client) HostedPublicVirtualInterfaces(namespace string) HostedPublicVirtualInterfaceInterface {
	return newHostedPublicVirtualInterfaces(c, namespace)
}

func (c *DxV1alpha1Client) HostedPublicVirtualInterfaceAccepters(namespace string) HostedPublicVirtualInterfaceAccepterInterface {
	return newHostedPublicVirtualInterfaceAccepters(c, namespace)
}

func (c *DxV1alpha1Client) HostedTransitVirtualInterfaces(namespace string) HostedTransitVirtualInterfaceInterface {
	return newHostedTransitVirtualInterfaces(c, namespace)
}

func (c *DxV1alpha1Client) HostedTransitVirtualInterfaceAccepters(namespace string) HostedTransitVirtualInterfaceAccepterInterface {
	return newHostedTransitVirtualInterfaceAccepters(c, namespace)
}

func (c *DxV1alpha1Client) Lags(namespace string) LagInterface {
	return newLags(c, namespace)
}

func (c *DxV1alpha1Client) PrivateVirtualInterfaces(namespace string) PrivateVirtualInterfaceInterface {
	return newPrivateVirtualInterfaces(c, namespace)
}

func (c *DxV1alpha1Client) PublicVirtualInterfaces(namespace string) PublicVirtualInterfaceInterface {
	return newPublicVirtualInterfaces(c, namespace)
}

func (c *DxV1alpha1Client) TransitVirtualInterfaces(namespace string) TransitVirtualInterfaceInterface {
	return newTransitVirtualInterfaces(c, namespace)
}

// NewForConfig creates a new DxV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*DxV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &DxV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new DxV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *DxV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new DxV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *DxV1alpha1Client {
	return &DxV1alpha1Client{c}
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
func (c *DxV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
