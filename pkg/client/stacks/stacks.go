package stack

import "github.com/aws/awslabs/aws-service-operator/pkg/config"

// StackInterface implements the main structure for using CloudFormation
type StackInterface interface {
	Config() *config.Config
	DyanmoDBStack
}

// StackClient returns the client for managing stacks
type StackClient struct {
	config *config.Config
}

// DynamoDB returns the DynamoDB stack
func (s *StackClient) DynamoDB(obj interface{}, topicARN string) DynamoDBInterface {
	return newDyanmoDB(s.config, obj, topicARN)
}

// Config is a passthrough for the Config type
func (s *StackClient) Config() *config.Config {
	return c.config
}

// NewForConfig will return the StackClient initialized with the Config
func NewForConfig(c *config.Config) *StackClient {
	return &StackClient{c}
}
