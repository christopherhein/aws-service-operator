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
	v1alpha1 "awsoperator.io/pkg/apis/iotanalytics/v1alpha1"
	"awsoperator.io/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type IotanalyticsV1alpha1Interface interface {
	RESTClient() rest.Interface
	ChannelsGetter
	DatasetsGetter
	DatastoresGetter
	PipelinesGetter
}

// IotanalyticsV1alpha1Client is used to interact with features provided by the iotanalytics.awsoperator.io group.
type IotanalyticsV1alpha1Client struct {
	restClient rest.Interface
}

func (c *IotanalyticsV1alpha1Client) Channels(namespace string) ChannelInterface {
	return newChannels(c, namespace)
}

func (c *IotanalyticsV1alpha1Client) Datasets(namespace string) DatasetInterface {
	return newDatasets(c, namespace)
}

func (c *IotanalyticsV1alpha1Client) Datastores(namespace string) DatastoreInterface {
	return newDatastores(c, namespace)
}

func (c *IotanalyticsV1alpha1Client) Pipelines(namespace string) PipelineInterface {
	return newPipelines(c, namespace)
}

// NewForConfig creates a new IotanalyticsV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*IotanalyticsV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &IotanalyticsV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new IotanalyticsV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *IotanalyticsV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new IotanalyticsV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *IotanalyticsV1alpha1Client {
	return &IotanalyticsV1alpha1Client{c}
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
func (c *IotanalyticsV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}