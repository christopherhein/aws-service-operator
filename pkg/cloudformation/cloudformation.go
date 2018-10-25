package cloudformation

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/awslabs/aws-service-operator/pkg/config"
	"github.com/awslabs/aws-service-operator/pkg/helpers"
	"github.com/sirupsen/logrus"
)

// DELETE_COMPLETE represents that the stack has been deleted successfully
const DELETE_COMPLETE = "DELETE_COMPLETE"

// CREATE_COMPLETE respresents that the stack has been created successfully
const CREATE_COMPLETE = "CREATE_COMPLETE"

// UPDATE_COMPLETE represents that the stack has been updated successfully
const UPDATE_COMPLETE = "UPDATE_COMPLETE"

// ROLLBACK_COMPLETE represents that the stack has been rolledback successfully
const ROLLBACK_COMPLETE = "ROLLBACK_COMPLETE"

// Name will return the proper stack name for each component
func Name(clusterName string, resourceType string, name string, namespace string) string {
	nameParts := []string{clusterName, resourceType, name, namespace}
	return KubernetesResourceName(strings.Join(nameParts, "-"))
}

// GetStackOutputs will return a map of attributes from the stack
func GetOutputs(c *config.Config, stackName string) (map[string]string, error) {
	outputs := map[string]string{}
	sess := c.AWSSession
	svc := cloudformation.New(sess)

	stackInputs := cloudformation.DescribeStacksInput{
		StackName: aws.String(stackName),
	}

	output, err := svc.DescribeStacks(&stackInputs)
	if err != nil {
		return nil, err
	}
	// Not sure if this is even possible
	if len(output.Stacks) != 1 {
		return nil, errors.New("no stacks returned with that stack name")
	}

	for _, out := range output.Stacks[0].Outputs {
		outputs[*out.OutputKey] = *out.OutputValue
	}

	return outputs, err
}

// SetDefaults will set the basic attributes like Name, Template URL and
// the Notification ARN
func SetDefaults(input interface{}, name, template, topicARN string) interface{} {
	input.SetStackName(name)
	input.SetTemplateURL(template)
	input.SetNotificationARNs([]*string{aws.String(s.topicARN)})
	return input
}

// SetParams takes a list of attrs and adds them to the stackInput
func SetParams(input interface{}, attrs map[string]string, templateHelpers interface{}) (interface{}, error) {
	params := []*cloudformation.Parameter{}
	for key, value := range attrs {
		computedValue, err := helpers.Templatize(value, templateHelpers)
		if err != nil {
			return nil, err
		}
		params := append(params, helpers.CreateParam(key, computedValue))
	}
	input.SetParameters(params)
}

// CreateParam returns a new prefilled cloudformation param
func CreateParam(key string, value string) *cloudformation.Parameter {
	param := &cloudformation.Parameter{}
	param.SetParameterKey(key)
	param.SetParameterValue(value)
	return param
}

// SetTags will add tags to the stackInput
func SetTags(input interface{}, attrs map[string]string) interface{} {
	tags := []*cloudformation.Tag{}
	for key, value := range attrs {
		tags = append(tags, helpers.CreateTag(key, value))
	}
	input.SetTags(tags)
}

// CreateTag returns a new prefilled cloudformation tag
func CreateTag(key string, value string) *cloudformation.Tag {
	tag := &cloudformation.Tag{}
	tag.SetKey(key)
	tag.SetValue(value)
	return tag
}

// IsComplete will determine if it's in a state to process
func IsComplete(status string, defaultRet bool) bool {
	switch status {
	case CREATE_COMPLETE:
		return true
	case UPDATE_COMPLETE:
		return true
	case DELETE_COMPLETE:
		return false
	case ROLLBACK_COMPLETE:
		return false
	}
	return defaultRet
}

// GetTemplateURL will return the url to the CFT from the CFT resource
func GetTemplateURL(config *config.Config, rType string, name string, namespace string) string {
	logger := config.Logger
	clientSet, _ := awsclient.NewForConfig(config.RESTConfig)

	var cName string
	var cNamespace string
	if name == "" {
		cName = rType
	}
	if namespace == "" {
		cNamespace = config.DefaultNamespace
	}

	resource, err := clientSet.CloudFormationTemplates(cNamespace).Get(cName, metav1.GetOptions{})
	if err != nil {
		logger.WithError(err).Error("error getting cloudformation template returning fallback template")
		return "https://s3-us-west-2.amazonaws.com/cloudkit-templates/" + rType + ".yaml"
	}

	logger.WithFields(logrus.Fields{
		"namespace": cNamespace,
		"name":      cName,
		"url":       resource.Output.URL,
	}).Info("found cloudformation template")
	return resource.Output.URL
}
