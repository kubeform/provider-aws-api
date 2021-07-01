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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/alb/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"
)

type AlbV1alpha1Interface interface {
	RESTClient() rest.Interface
	AlbsGetter
	ListenersGetter
	ListenerCertificatesGetter
	ListenerRulesGetter
	TargetGroupsGetter
	TargetGroupAttachmentsGetter
}

// AlbV1alpha1Client is used to interact with features provided by the alb.aws.kubeform.com group.
type AlbV1alpha1Client struct {
	restClient rest.Interface
}

func (c *AlbV1alpha1Client) Albs(namespace string) AlbInterface {
	return newAlbs(c, namespace)
}

func (c *AlbV1alpha1Client) Listeners(namespace string) ListenerInterface {
	return newListeners(c, namespace)
}

func (c *AlbV1alpha1Client) ListenerCertificates(namespace string) ListenerCertificateInterface {
	return newListenerCertificates(c, namespace)
}

func (c *AlbV1alpha1Client) ListenerRules(namespace string) ListenerRuleInterface {
	return newListenerRules(c, namespace)
}

func (c *AlbV1alpha1Client) TargetGroups(namespace string) TargetGroupInterface {
	return newTargetGroups(c, namespace)
}

func (c *AlbV1alpha1Client) TargetGroupAttachments(namespace string) TargetGroupAttachmentInterface {
	return newTargetGroupAttachments(c, namespace)
}

// NewForConfig creates a new AlbV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*AlbV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &AlbV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new AlbV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *AlbV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new AlbV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *AlbV1alpha1Client {
	return &AlbV1alpha1Client{c}
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
func (c *AlbV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
