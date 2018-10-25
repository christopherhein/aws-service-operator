package stack

import (
	"github.com/aws/aws-sdk-go/service/cloudformation"
	cfnhelpers "github.com/aws/awslabs/aws-service-operator/pkg/cloudformation"
	"github.com/aws/awslabs/aws-service-operator/pkg/config"
	"github.com/awslabs/aws-service-operator/pkg/helpers"
	"k8s.io/client-go/util/retry"
)

// DyanmoDBStack turns the DynamoDBInterface
type DyanmoDBStack interface {
	DyanmoDB(*config.Config, interface{}, string) DyanmoDBInterface
}

// DyanmoDBInterface implements the controller methods for the type
type DyanmoDBInterface interface {
	Name() string
	Outputs() (map[string]string, error)
	Sync(*awsV1alpha1.DynamoDB) (*awsV1alpha1.DynamoDB, error)
	Describe() (*cloudformation.DescribeStackOutput, error)
	Create() (*cloudformation.CreateStackOutput, error)
	Update(*awsV1alpha1.DynamoDB) (*cloudformation.UpdateStackOutput, error)
	Delete() (*cloudformation.DeleteStackOutput, error)
	WaitUntilDeleted() error
	Status(string, string, string) error
	SyncAdditionalObjects() error
	GetTemplate() string
}

// dyanmoDB implements DyanmoDBInterface
type dyanmoDB struct {
	config   *config.Config
	resource *awsV1alpha1.DynamoDB
	topicARN string
}

// newDyanmoDB returns a dyanmoDB
func newDynamoDB(c *config.Config, obj interface{}, topicARN string) *dyanmoDB {
	return &dyanmoDB{
		config:   c,
		resource: obj.(*awsV1alpha1.DynamoDB),
		topicARN: topicARN,
	}
}

// Name returns the name of the stack based on the resource
func (s *dyanmoDB) Name() string {
	return cfnhelpers.Name(s.config.ClusterName, "dynamodb", s.resource.Name, s.resource.Namespace)
}

// Outputs returns a map of attirbutes related to the stack that was deployed
func (s *dynamodb) Outputs() (map[string]string, error) {
	return cfnhelpers.GetOutputs(s.config, s.Name())
}

func (s *dynamodb) Sync(updated *awsV1alpha1.DynamoDB) (stackID string, err error) {
	output, err := s.Get()
	if err != nil {
		return nil, err
	}

	if len(output.Stacks) == 0 {
		output, err := s.Create()
		if err != nil {
			return nil, err
		}
		return string(*output.StackId), err
	}

	output, err := s.Update(updated)
	if err != nil {
		return nil, err
	}
	return string(*output.StackId), err

	return
}

// Get will describe the stack if it exists
func (s *dynamodb) Get() (output *cloudformation.DescribeStacksOutput, err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	input := cloudformation.DescribeStackInput{}
	input.SetStackName(s.resource.Name())

	output, err = svc.DescribeStacks(&input)
	return
}

// Create will orchestrate creating the DynamoDB CloudFormation Stack
func (s *dynamodb) Create() (output *cloudformation.CreateStackOutput, err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	input := cfnhelpers.SetDefaults(&cloudformation.CreateStackInput{}, s.resource.Name(), s.GetTemplate(), s.topicARN)

	attributes := map[string]string{
		"ResourceName":       "{{.Obj.Name}}",
		"ResourceVersion":    "{{.Obj.ResourceVersion}}",
		"Namespace":          "{{.Obj.Namespace}}",
		"ClusterName":        "{{.Config.ClusterName}}",
		"TableName":          "{{.Obj.Name}}",
		"RangeAttributeName": "{{.Obj.Spec.RangeAttribute.Name}}",
		"RangeAttributeType": "{{.Obj.Spec.RangeAttribute.Type}}",
		"ReadCapacityUnits":  "{{.Obj.Spec.ReadCapacityUnits}}",
		"WriteCapacityUnits": "{{.Obj.Spec.WriteCapacityUnits}}",
		"HashAttributeName":  "{{.Obj.Spec.HashAttribute.Name}}",
		"HashAttributeType":  "{{.Obj.Spec.HashAttribute.Type}}",
	}

	templateHelpers := helpers.Data{
		Obj:     s.DynamoDB,
		Config:  s.config,
		Helpers: helpers.New(),
	}

	input, err = cfnhelpers.SetParams(input, attributes, templateHelpers)
	if err != nil {
		s.config.Logger.WithError(err).Error("error creating stack parameters")
	}

	tags := map[string]string{
		"ResourceName":    s.resource.Name,
		"ResourceVersion": s.resource.ResourceVersion,
		"Namespace":       s.resource.Namespace,
		"ClusterName":     s.config.ClusterName,
	}
	input = cfnhelpers.SetTags(input, tags)

	stackInputs := input.(*cloudformation.CreateStackInput)
	output, err = svc.CreateStack(stackInputs)
	return
}

// Update will orchestrate updating the DynamoDB CloudFormation Stack
func (s *dynamodb) Update(updated *awsV1alpha1.DynamoDB) (output *cloudformation.UpdateStackOutput, err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	input := cfnhelpers.SetDefaults(&cloudformation.UpdateStackInput{}, s.resource.Name(), s.GetTemplate(), s.topicARN)

	attributes := map[string]string{
		"ResourceName":       "{{.Obj.Name}}",
		"ResourceVersion":    "{{.Obj.ResourceVersion}}",
		"Namespace":          "{{.Obj.Namespace}}",
		"ClusterName":        "{{.Config.ClusterName}}",
		"TableName":          "{{.Obj.Name}}",
		"RangeAttributeName": "{{.Obj.Spec.RangeAttribute.Name}}",
		"RangeAttributeType": "{{.Obj.Spec.RangeAttribute.Type}}",
		"ReadCapacityUnits":  "{{.Obj.Spec.ReadCapacityUnits}}",
		"WriteCapacityUnits": "{{.Obj.Spec.WriteCapacityUnits}}",
		"HashAttributeName":  "{{.Obj.Spec.HashAttribute.Name}}",
		"HashAttributeType":  "{{.Obj.Spec.HashAttribute.Type}}",
	}

	templateHelpers := helpers.Data{
		Obj:     s.DynamoDB,
		Config:  s.config,
		Helpers: helpers.New(),
	}

	input, err = cfnhelpers.SetParams(input, attributes, templateHelpers)
	if err != nil {
		s.config.Logger.WithError(err).Error("error creating stack parameters")
	}

	tags := map[string]string{
		"ResourceName":    s.resource.Name,
		"ResourceVersion": s.resource.ResourceVersion,
		"Namespace":       s.resource.Namespace,
		"ClusterName":     s.config.ClusterName,
	}
	input = cfnhelpers.SetTags(input, tags)

	stackInputs := input.(*cloudformation.CreateStackInput)
	output, err = svc.UpdateStack(stackInputs)
	return
}

// Delete will orchestrate deleteing the DynamoDB CloudFormation stack
func (s *dynamodb) Delete() (output *cloudformation.DeleteStackOutput, err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	input := cloudformation.DeleteStackInput{}
	input.SetStackName(s.Name())

	output, err = svc.DeleteStack(&input)
	return
}

// WaitUntilDeleted will orchestrate deleting the DynamoDB CloudFormation stack
// and wait for the stack to delete before returning.
func (s *dynamodb) WaitUntilDeleted() (err error) {
	sess := s.config.AWSSession
	svc := cloudformation.New(sess)

	input := cloudformation.DescribeStacksInput{}
	input.SetStackName(s.Name())

	err = svc.WaitUntilStackDeleteComplete(&input)
	return
}

// Status will handle updating the status, which get's written to the resources
func (s *dynamodb) Status(stackID, status, reason string) (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		resource, err := s.config.AWSClientset.DynamoDBs(s.resource.Namespace).Get(s.resource.Name, metav1.GetOptions{})
		if err != nil {
			s.config.Logger.WithError(err).Error("error finding resource")
			return err
		}

		resourceCopy := resource.DeepCopy()
		resourceCopy.Status.ResourceStatus = status
		resourceCopy.Status.ResourceStatusReason = reason
		resourceCopy.Status.StackID = stackID

		if cfnhelpers.IsComplete(status, false) {
			outputs, err := s.Outputs()
			if err != nil {
				s.config.Logger.WithError(err).Error("error getting stack outputs")
				return err
			}
			resourceCopy.Output.TableName = outputs["TableName"]
			resourceCopy.Output.TableARN = outputs["TableArn"]
		}

		resourceCopy, err = s.config.AWSClientset.DynamoDBs(namespace).Update(resourceCopy)
		if err != nil {
			s.config.Logger.WithError(err).Error("error updating resource")
			return err
		}
		return
	})
	if err != nil {
		s.config.Logger.WithError(err).Error("error retrying status update, max backoff reached")
		return err
	}
	s.SyncAdditionalObjects()
	return
}

// SyncAdditionalObjects will take update any additional resources that the
// operator is suppose to manage
//
// TODO: Consider making SyncAdditionalObject 2-3 separate functions for CMs,
// Secrets, Services, etc. Or consider building an internal Event loop when
// status happen.
func (s *dynamodb) SyncAdditionalObjects() (err error) {
	err = retry.RetryOnConflict(retry.DefaultRetry, func() error {
		resource, err := s.config.AWSClientset.DynamoDBs(s.resource.Namespace).Get(s.resource.Name, metav1.GetOptions{})
		if err != nil {
			s.config.Logger.WithError(err).Error("error finding resource")
			return err
		}

		resourceCopy := resource.DeepCopy()
		configmaps := []string{}
		dynamoCMData := map[string]string{
			"tableName": "{{.Obj.Output.TableName}}",
			"tableARN":  "{{.Obj.Output.TableARN}}",
			"region":    "{{.Config.Region}}",
		}
		dynamoCM := helpers.SyncConfigMap(config, s, s.resource.Name, s.resource.Namespace, dynamoCMData)
		configmaps = append(configmaps, dynamoCM)
		resourceCopy.AdditionalResources.ConfigMaps = configmaps

		_, err = s.config.AWSClientset.DynamoDBs(s.resource.Namespace).Update(resourceCopy)
		if err != nil {
			s.config.Logger.WithError(err).Error("error updating resource")
			return err
		}
		return nil
	})
	if err != nil {
		s.config.Logger.WithError(err).Error("error retrying status update, max backoff reached")
		return err
	}
	return
}

// GetTemplate will return the URL to the DynamoDB CloudFormation Template
func (s *dynamodb) GetTemplate() string {
	return cloudformation.GetTemplateURL(s.config, "dynamodb", s.DynamoDB.Spec.CloudFormationTemplateName, s.DynamoDB.Spec.CloudFormationTemplateNamespace)
}
