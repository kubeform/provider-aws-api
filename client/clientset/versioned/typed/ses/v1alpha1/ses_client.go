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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/ses/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type SesV1alpha1Interface interface {
	RESTClient() rest.Interface
	ActiveReceiptRuleSetsGetter
	ConfigurationSetsGetter
	DomainDkimsGetter
	DomainIdentitiesGetter
	DomainIdentityVerificationsGetter
	DomainMailFromsGetter
	EmailIdentitiesGetter
	EventDestinationsGetter
	IdentityNotificationTopicsGetter
	IdentityPoliciesGetter
	ReceiptFiltersGetter
	ReceiptRulesGetter
	ReceiptRuleSetsGetter
	TemplatesGetter
}

// SesV1alpha1Client is used to interact with features provided by the ses.aws.kubeform.com group.
type SesV1alpha1Client struct {
	restClient rest.Interface
}

func (c *SesV1alpha1Client) ActiveReceiptRuleSets(namespace string) ActiveReceiptRuleSetInterface {
	return newActiveReceiptRuleSets(c, namespace)
}

func (c *SesV1alpha1Client) ConfigurationSets(namespace string) ConfigurationSetInterface {
	return newConfigurationSets(c, namespace)
}

func (c *SesV1alpha1Client) DomainDkims(namespace string) DomainDkimInterface {
	return newDomainDkims(c, namespace)
}

func (c *SesV1alpha1Client) DomainIdentities(namespace string) DomainIdentityInterface {
	return newDomainIdentities(c, namespace)
}

func (c *SesV1alpha1Client) DomainIdentityVerifications(namespace string) DomainIdentityVerificationInterface {
	return newDomainIdentityVerifications(c, namespace)
}

func (c *SesV1alpha1Client) DomainMailFroms(namespace string) DomainMailFromInterface {
	return newDomainMailFroms(c, namespace)
}

func (c *SesV1alpha1Client) EmailIdentities(namespace string) EmailIdentityInterface {
	return newEmailIdentities(c, namespace)
}

func (c *SesV1alpha1Client) EventDestinations(namespace string) EventDestinationInterface {
	return newEventDestinations(c, namespace)
}

func (c *SesV1alpha1Client) IdentityNotificationTopics(namespace string) IdentityNotificationTopicInterface {
	return newIdentityNotificationTopics(c, namespace)
}

func (c *SesV1alpha1Client) IdentityPolicies(namespace string) IdentityPolicyInterface {
	return newIdentityPolicies(c, namespace)
}

func (c *SesV1alpha1Client) ReceiptFilters(namespace string) ReceiptFilterInterface {
	return newReceiptFilters(c, namespace)
}

func (c *SesV1alpha1Client) ReceiptRules(namespace string) ReceiptRuleInterface {
	return newReceiptRules(c, namespace)
}

func (c *SesV1alpha1Client) ReceiptRuleSets(namespace string) ReceiptRuleSetInterface {
	return newReceiptRuleSets(c, namespace)
}

func (c *SesV1alpha1Client) Templates(namespace string) TemplateInterface {
	return newTemplates(c, namespace)
}

// NewForConfig creates a new SesV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*SesV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &SesV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new SesV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SesV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SesV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *SesV1alpha1Client {
	return &SesV1alpha1Client{c}
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
func (c *SesV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
