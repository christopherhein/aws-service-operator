package base

import "github.com/aws/awslabs/aws-service-operator/pkg/config"

// Interface implements the main structure for each operator
type ControllerInterface interface {
	Config() *config.Config
	DynamoDBController
}

// ControllerClient implements the main controller client
type ControllerClient struct {
	config *config.Config
}

// DynamoDB returns the DynamoDB controller
func (c *ControllerClient) DyanmoDB() DyanmoDBInterface {
	return newDynamoDB(c)
}

// Config is a passthrough for the Config type
func (c *ControllerClient) Config() *config.Config {
	return c.config
}

// NewForConfig will return the ControllerClient initialized with the Config
func NewForConfig(c *config.Config) *ControllerClient {
	return &ControllerClient{config}
}
