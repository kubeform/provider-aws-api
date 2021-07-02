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
	v1alpha1 "kubeform.dev/provider-aws-api/client/clientset/versioned/typed/guardduty/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeGuarddutyV1alpha1 struct {
	*testing.Fake
}

func (c *FakeGuarddutyV1alpha1) Detectors(namespace string) v1alpha1.DetectorInterface {
	return &FakeDetectors{c, namespace}
}

func (c *FakeGuarddutyV1alpha1) Filters(namespace string) v1alpha1.FilterInterface {
	return &FakeFilters{c, namespace}
}

func (c *FakeGuarddutyV1alpha1) InviteAccepters(namespace string) v1alpha1.InviteAccepterInterface {
	return &FakeInviteAccepters{c, namespace}
}

func (c *FakeGuarddutyV1alpha1) Ipsets(namespace string) v1alpha1.IpsetInterface {
	return &FakeIpsets{c, namespace}
}

func (c *FakeGuarddutyV1alpha1) Members(namespace string) v1alpha1.MemberInterface {
	return &FakeMembers{c, namespace}
}

func (c *FakeGuarddutyV1alpha1) OrganizationAdminAccounts(namespace string) v1alpha1.OrganizationAdminAccountInterface {
	return &FakeOrganizationAdminAccounts{c, namespace}
}

func (c *FakeGuarddutyV1alpha1) OrganizationConfigurations(namespace string) v1alpha1.OrganizationConfigurationInterface {
	return &FakeOrganizationConfigurations{c, namespace}
}

func (c *FakeGuarddutyV1alpha1) PublishingDestinations(namespace string) v1alpha1.PublishingDestinationInterface {
	return &FakePublishingDestinations{c, namespace}
}

func (c *FakeGuarddutyV1alpha1) Threatintelsets(namespace string) v1alpha1.ThreatintelsetInterface {
	return &FakeThreatintelsets{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeGuarddutyV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
