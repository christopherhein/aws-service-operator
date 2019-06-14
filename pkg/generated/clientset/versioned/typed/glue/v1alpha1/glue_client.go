/*
Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License"). You may
not use this file except in compliance with the License. A copy of the
License is located at

     http://aws.amazon.com/apache2.0/

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "awsoperator.io/pkg/apis/glue/v1alpha1"
	"awsoperator.io/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type GlueV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClassifiersGetter
	ConnectionsGetter
	CrawlersGetter
	DataCatalogEncryptionSettingsesGetter
	DatabasesGetter
	DevEndpointsGetter
	JobsGetter
	PartitionsGetter
	SecurityConfigurationsGetter
	TablesGetter
	TriggersGetter
}

// GlueV1alpha1Client is used to interact with features provided by the glue.awsoperator.io group.
type GlueV1alpha1Client struct {
	restClient rest.Interface
}

func (c *GlueV1alpha1Client) Classifiers(namespace string) ClassifierInterface {
	return newClassifiers(c, namespace)
}

func (c *GlueV1alpha1Client) Connections(namespace string) ConnectionInterface {
	return newConnections(c, namespace)
}

func (c *GlueV1alpha1Client) Crawlers(namespace string) CrawlerInterface {
	return newCrawlers(c, namespace)
}

func (c *GlueV1alpha1Client) DataCatalogEncryptionSettingses(namespace string) DataCatalogEncryptionSettingsInterface {
	return newDataCatalogEncryptionSettingses(c, namespace)
}

func (c *GlueV1alpha1Client) Databases(namespace string) DatabaseInterface {
	return newDatabases(c, namespace)
}

func (c *GlueV1alpha1Client) DevEndpoints(namespace string) DevEndpointInterface {
	return newDevEndpoints(c, namespace)
}

func (c *GlueV1alpha1Client) Jobs(namespace string) JobInterface {
	return newJobs(c, namespace)
}

func (c *GlueV1alpha1Client) Partitions(namespace string) PartitionInterface {
	return newPartitions(c, namespace)
}

func (c *GlueV1alpha1Client) SecurityConfigurations(namespace string) SecurityConfigurationInterface {
	return newSecurityConfigurations(c, namespace)
}

func (c *GlueV1alpha1Client) Tables(namespace string) TableInterface {
	return newTables(c, namespace)
}

func (c *GlueV1alpha1Client) Triggers(namespace string) TriggerInterface {
	return newTriggers(c, namespace)
}

// NewForConfig creates a new GlueV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*GlueV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &GlueV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new GlueV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *GlueV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new GlueV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *GlueV1alpha1Client {
	return &GlueV1alpha1Client{c}
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
func (c *GlueV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}