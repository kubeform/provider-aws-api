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
	v1alpha1 "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"
	"kubeform.dev/provider-aws-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type S3V1alpha1Interface interface {
	RESTClient() rest.Interface
	AccessPointsGetter
	AccountPublicAccessBlocksGetter
	BucketsGetter
	BucketACLsGetter
	BucketAccelerateConfigurationsGetter
	BucketAnalyticsConfigurationsGetter
	BucketCorsConfigurationsGetter
	BucketIntelligentTieringConfigurationsGetter
	BucketInventoriesGetter
	BucketLifecycleConfigurationsGetter
	BucketLoggingsGetter
	BucketMetricsGetter
	BucketNotificationsGetter
	BucketObjectsGetter
	BucketObjectLockConfigurationsGetter
	BucketOwnershipControlsesGetter
	BucketPoliciesGetter
	BucketPublicAccessBlocksGetter
	BucketReplicationConfigurationsGetter
	BucketRequestPaymentConfigurationsGetter
	BucketServerSideEncryptionConfigurationsGetter
	BucketVersioningsGetter
	BucketWebsiteConfigurationsGetter
	ObjectsGetter
	ObjectCopiesGetter
}

// S3V1alpha1Client is used to interact with features provided by the s3.aws.kubeform.com group.
type S3V1alpha1Client struct {
	restClient rest.Interface
}

func (c *S3V1alpha1Client) AccessPoints(namespace string) AccessPointInterface {
	return newAccessPoints(c, namespace)
}

func (c *S3V1alpha1Client) AccountPublicAccessBlocks(namespace string) AccountPublicAccessBlockInterface {
	return newAccountPublicAccessBlocks(c, namespace)
}

func (c *S3V1alpha1Client) Buckets(namespace string) BucketInterface {
	return newBuckets(c, namespace)
}

func (c *S3V1alpha1Client) BucketACLs(namespace string) BucketACLInterface {
	return newBucketACLs(c, namespace)
}

func (c *S3V1alpha1Client) BucketAccelerateConfigurations(namespace string) BucketAccelerateConfigurationInterface {
	return newBucketAccelerateConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketAnalyticsConfigurations(namespace string) BucketAnalyticsConfigurationInterface {
	return newBucketAnalyticsConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketCorsConfigurations(namespace string) BucketCorsConfigurationInterface {
	return newBucketCorsConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketIntelligentTieringConfigurations(namespace string) BucketIntelligentTieringConfigurationInterface {
	return newBucketIntelligentTieringConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketInventories(namespace string) BucketInventoryInterface {
	return newBucketInventories(c, namespace)
}

func (c *S3V1alpha1Client) BucketLifecycleConfigurations(namespace string) BucketLifecycleConfigurationInterface {
	return newBucketLifecycleConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketLoggings(namespace string) BucketLoggingInterface {
	return newBucketLoggings(c, namespace)
}

func (c *S3V1alpha1Client) BucketMetrics(namespace string) BucketMetricInterface {
	return newBucketMetrics(c, namespace)
}

func (c *S3V1alpha1Client) BucketNotifications(namespace string) BucketNotificationInterface {
	return newBucketNotifications(c, namespace)
}

func (c *S3V1alpha1Client) BucketObjects(namespace string) BucketObjectInterface {
	return newBucketObjects(c, namespace)
}

func (c *S3V1alpha1Client) BucketObjectLockConfigurations(namespace string) BucketObjectLockConfigurationInterface {
	return newBucketObjectLockConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketOwnershipControlses(namespace string) BucketOwnershipControlsInterface {
	return newBucketOwnershipControlses(c, namespace)
}

func (c *S3V1alpha1Client) BucketPolicies(namespace string) BucketPolicyInterface {
	return newBucketPolicies(c, namespace)
}

func (c *S3V1alpha1Client) BucketPublicAccessBlocks(namespace string) BucketPublicAccessBlockInterface {
	return newBucketPublicAccessBlocks(c, namespace)
}

func (c *S3V1alpha1Client) BucketReplicationConfigurations(namespace string) BucketReplicationConfigurationInterface {
	return newBucketReplicationConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketRequestPaymentConfigurations(namespace string) BucketRequestPaymentConfigurationInterface {
	return newBucketRequestPaymentConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketServerSideEncryptionConfigurations(namespace string) BucketServerSideEncryptionConfigurationInterface {
	return newBucketServerSideEncryptionConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) BucketVersionings(namespace string) BucketVersioningInterface {
	return newBucketVersionings(c, namespace)
}

func (c *S3V1alpha1Client) BucketWebsiteConfigurations(namespace string) BucketWebsiteConfigurationInterface {
	return newBucketWebsiteConfigurations(c, namespace)
}

func (c *S3V1alpha1Client) Objects(namespace string) ObjectInterface {
	return newObjects(c, namespace)
}

func (c *S3V1alpha1Client) ObjectCopies(namespace string) ObjectCopyInterface {
	return newObjectCopies(c, namespace)
}

// NewForConfig creates a new S3V1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*S3V1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &S3V1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new S3V1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *S3V1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new S3V1alpha1Client for the given RESTClient.
func New(c rest.Interface) *S3V1alpha1Client {
	return &S3V1alpha1Client{c}
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
func (c *S3V1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
