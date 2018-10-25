package message

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/awslabs/aws-service-operator/pkg/config"
	"github.com/awslabs/aws-service-operator/pkg/helpers"
)

const (
	ROLLBACK_COMPLETE = "ROLLBACK_COMPLETE"
	DELETE_COMPLETE   = "DELETE_COMPLETE"
)

// HandlerFunc allows you to define a custom function for when a message is stored
type HandlerFunc func(config *config.Config, msg *MessageBody) error

// Handler allows a custom function to be passed
type Handler interface {
	HandleMessage(config *config.Config, msg *MessageBody) error
}

// BodyInterface will return an object which will expose methods for the queue
// processers to use.
type BodyInterface interface {
	ParseMessage() error
	IsComplete() bool
	ResourceName() string
	Namespace() string
	StackID() string
	ResourceStatus() string
	StackName() string
	ResourceType() string
	ResourceStatusReason() string
	Valid() bool
	TopicARN() string
	SetResourceName(string) string
	SetNamespace(string) string
}

// Body will parse the message from the Body of SQS
type Body struct {
	Type               string `json:"Type"`
	topicARN           string `json:"TopicArn"`
	Message            string `json:"Message"`
	ParsedMessage      map[string]string
	namespace          string
	resourceName       string
	ResourceProperties ResourceProperties
	valid              bool
}

// ResourceProperties will wrap the ResourceProperties object
type ResourceProperties struct {
	Tags []Tag `json:"Tags"`
}

// Tag represents a Tag
type Tag struct {
	Key   string `json:"Key"`
	Value string `json:"Value"`
}

// HandleMessage will stub the handler for processing messages
func (f HandlerFunc) HandleMessage(config *config.Config, msg *MessageBody) error {
	return f(config, msg)
}

// New will return a new pointer to a body
func New(item *sqs.Message) (b *Body, err error) {
	err = json.Unmarshal([]byte(*item.Body), b)
	if err != nil {
		return b, err
	}
	b.ParseMessage()

	return b, nil
}

// ParseMessage will take the message attribute and make it readable
func (m *Body) ParseMessage() error {
	m.valid = false
	resp := make(map[string]string)
	items := strings.Split(m.Message, "\n")
	for _, item := range items {
		x := strings.Split(item, "=")
		key := x[0]
		if key != "" {
			s := x[1]
			s = s[1 : len(s)-1]
			resp[key] = s
		}
	}
	m.ParsedMessage = resp

	var resourceProperties ResourceProperties
	if resp["ResourceProperties"] != "null" {
		err := json.Unmarshal([]byte(resp["ResourceProperties"]), &resourceProperties)
		if err != nil {
			return err
		}
		m.ResourceProperties = resourceProperties
		for _, tag := range resourceProperties.Tags {
			switch tag.Key {
			case "Namespace":
				m.Namespace = tag.Value
			case "ResourceName":
				m.ResourceName = tag.Value
			}
		}
		if m.Namespace != "" && m.ResourceName != "" {
			m.valid = true
		}
	}
	return nil
}

// IsComplete returns a simple status instead of the raw CFT resp
func (m *Body) IsComplete() bool {
	return helpers.IsStackComplete(m.ParsedMessage["ResourceStatus"], true)
}

// Namespace returns the namespace of the object
func (m *Body) Namespace() string {
	return m.namespace
}

// ResourceName returns the name of the object
func (m *Body) ResourceName() string {
	return m.resourceName
}

// StackID returns the stackID from the parsed message
func (m *Body) StackID() string {
	return m.parsedMessage["StackId"]
}

// ResourceStatus will return the parsed Resource Status
func (m *Body) ResourceStatus() string {
	return m.parsedMessage["ResourceStatus"]
}

// StackName will return the parsed stack name
func (m *Body) StackName() string {
	return m.parsedMessage["StackName"]
}

// ResourceType will return the parsed resource type
func (m *Body) ResourceType() string {
	return m.parsedMessage["ResourceType"]
}

// ResourceStatusReason will return the parse resource status reason
func (m *Body) ResourceStatusReason() string {
	return m.parsedMessage["ResourceStatusReason"]
}

// Valid will return true or false if the resource has all valid params
func (m *Body) Valid() bool {
	return m.valid
}

// TopicARN will return the proper Topic ARN
func (m *Body) TopicARN() string {
	return m.topicARN
}

// SetResourceName will set the name on the Body
func (m *Body) SetResourceName(name string) {
	m.resourceName = name
}

// SetNamespace will set the namespace on the Body
func (m *Body) SetNamespace(namespace string) {
	m.namespace = namespace
}
