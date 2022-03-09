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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BgpPeers returns a BgpPeerInformer.
	BgpPeers() BgpPeerInformer
	// Connections returns a ConnectionInformer.
	Connections() ConnectionInformer
	// ConnectionAssociations returns a ConnectionAssociationInformer.
	ConnectionAssociations() ConnectionAssociationInformer
	// ConnectionConfirmations returns a ConnectionConfirmationInformer.
	ConnectionConfirmations() ConnectionConfirmationInformer
	// Gateways returns a GatewayInformer.
	Gateways() GatewayInformer
	// GatewayAssociations returns a GatewayAssociationInformer.
	GatewayAssociations() GatewayAssociationInformer
	// GatewayAssociationProposals returns a GatewayAssociationProposalInformer.
	GatewayAssociationProposals() GatewayAssociationProposalInformer
	// HostedConnections returns a HostedConnectionInformer.
	HostedConnections() HostedConnectionInformer
	// HostedPrivateVirtualInterfaces returns a HostedPrivateVirtualInterfaceInformer.
	HostedPrivateVirtualInterfaces() HostedPrivateVirtualInterfaceInformer
	// HostedPrivateVirtualInterfaceAccepters returns a HostedPrivateVirtualInterfaceAccepterInformer.
	HostedPrivateVirtualInterfaceAccepters() HostedPrivateVirtualInterfaceAccepterInformer
	// HostedPublicVirtualInterfaces returns a HostedPublicVirtualInterfaceInformer.
	HostedPublicVirtualInterfaces() HostedPublicVirtualInterfaceInformer
	// HostedPublicVirtualInterfaceAccepters returns a HostedPublicVirtualInterfaceAccepterInformer.
	HostedPublicVirtualInterfaceAccepters() HostedPublicVirtualInterfaceAccepterInformer
	// HostedTransitVirtualInterfaces returns a HostedTransitVirtualInterfaceInformer.
	HostedTransitVirtualInterfaces() HostedTransitVirtualInterfaceInformer
	// HostedTransitVirtualInterfaceAccepters returns a HostedTransitVirtualInterfaceAccepterInformer.
	HostedTransitVirtualInterfaceAccepters() HostedTransitVirtualInterfaceAccepterInformer
	// Lags returns a LagInformer.
	Lags() LagInformer
	// PrivateVirtualInterfaces returns a PrivateVirtualInterfaceInformer.
	PrivateVirtualInterfaces() PrivateVirtualInterfaceInformer
	// PublicVirtualInterfaces returns a PublicVirtualInterfaceInformer.
	PublicVirtualInterfaces() PublicVirtualInterfaceInformer
	// TransitVirtualInterfaces returns a TransitVirtualInterfaceInformer.
	TransitVirtualInterfaces() TransitVirtualInterfaceInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BgpPeers returns a BgpPeerInformer.
func (v *version) BgpPeers() BgpPeerInformer {
	return &bgpPeerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Connections returns a ConnectionInformer.
func (v *version) Connections() ConnectionInformer {
	return &connectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConnectionAssociations returns a ConnectionAssociationInformer.
func (v *version) ConnectionAssociations() ConnectionAssociationInformer {
	return &connectionAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConnectionConfirmations returns a ConnectionConfirmationInformer.
func (v *version) ConnectionConfirmations() ConnectionConfirmationInformer {
	return &connectionConfirmationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Gateways returns a GatewayInformer.
func (v *version) Gateways() GatewayInformer {
	return &gatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GatewayAssociations returns a GatewayAssociationInformer.
func (v *version) GatewayAssociations() GatewayAssociationInformer {
	return &gatewayAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GatewayAssociationProposals returns a GatewayAssociationProposalInformer.
func (v *version) GatewayAssociationProposals() GatewayAssociationProposalInformer {
	return &gatewayAssociationProposalInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HostedConnections returns a HostedConnectionInformer.
func (v *version) HostedConnections() HostedConnectionInformer {
	return &hostedConnectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HostedPrivateVirtualInterfaces returns a HostedPrivateVirtualInterfaceInformer.
func (v *version) HostedPrivateVirtualInterfaces() HostedPrivateVirtualInterfaceInformer {
	return &hostedPrivateVirtualInterfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HostedPrivateVirtualInterfaceAccepters returns a HostedPrivateVirtualInterfaceAccepterInformer.
func (v *version) HostedPrivateVirtualInterfaceAccepters() HostedPrivateVirtualInterfaceAccepterInformer {
	return &hostedPrivateVirtualInterfaceAccepterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HostedPublicVirtualInterfaces returns a HostedPublicVirtualInterfaceInformer.
func (v *version) HostedPublicVirtualInterfaces() HostedPublicVirtualInterfaceInformer {
	return &hostedPublicVirtualInterfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HostedPublicVirtualInterfaceAccepters returns a HostedPublicVirtualInterfaceAccepterInformer.
func (v *version) HostedPublicVirtualInterfaceAccepters() HostedPublicVirtualInterfaceAccepterInformer {
	return &hostedPublicVirtualInterfaceAccepterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HostedTransitVirtualInterfaces returns a HostedTransitVirtualInterfaceInformer.
func (v *version) HostedTransitVirtualInterfaces() HostedTransitVirtualInterfaceInformer {
	return &hostedTransitVirtualInterfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HostedTransitVirtualInterfaceAccepters returns a HostedTransitVirtualInterfaceAccepterInformer.
func (v *version) HostedTransitVirtualInterfaceAccepters() HostedTransitVirtualInterfaceAccepterInformer {
	return &hostedTransitVirtualInterfaceAccepterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Lags returns a LagInformer.
func (v *version) Lags() LagInformer {
	return &lagInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PrivateVirtualInterfaces returns a PrivateVirtualInterfaceInformer.
func (v *version) PrivateVirtualInterfaces() PrivateVirtualInterfaceInformer {
	return &privateVirtualInterfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PublicVirtualInterfaces returns a PublicVirtualInterfaceInformer.
func (v *version) PublicVirtualInterfaces() PublicVirtualInterfaceInformer {
	return &publicVirtualInterfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TransitVirtualInterfaces returns a TransitVirtualInterfaceInformer.
func (v *version) TransitVirtualInterfaces() TransitVirtualInterfaceInformer {
	return &transitVirtualInterfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
