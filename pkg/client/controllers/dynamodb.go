package controllers

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	awsV1alpha1 "github.com/awslabs/aws-service-operator/pkg/apis/service-operator.aws/v1alpha1"
	"github.com/awslabs/aws-service-operator/pkg/client/stack"
	"github.com/awslabs/aws-service-operator/pkg/cloudformation"
	"github.com/awslabs/aws-service-operator/pkg/config"
	"github.com/awslabs/aws-service-operator/pkg/controller"
	"github.com/awslabs/aws-service-operator/pkg/message"
	"github.com/awslabs/aws-service-operator/pkg/queue"
	"github.com/awslabs/aws-service-operator/pkg/queuemanager"
	"github.com/iancoleman/strcase"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/cache"
)

// DyanmoDBController returns the DyanmoDBInterface
type DyanmoDBController interface {
	DyanmoDB() DyanmoDBInterface
}

// DyanmoDBInterface implements the controller methods for the type
type DyanmoDBInterface interface {
	Create(interface{})
	Update(interface{}, interface{})
	Delete(interface{})
	Status(interface{})
	Register()
	Watch(context.Context, string)
	HandleQueue(interface{})
}

// dyanmoDB implements DyanmoDBInterface
type dyanmoDB struct {
	config   *config.Config
	topicARN string
}

// newDyanmoDB returns a dyanmoDB
func newDyanmoDB(c *config.Config) *dyanmoDB {
	return &dyanmoDB{
		config: c,
	}
}

// Create will handle the Create mechanism for the Informer
func (c *dyanmoDB) Create(obj interface{}) {
	s := obj.(*awsV1alpha1.DynamoDB).DeepCopy()

	client := stack.NewForConfig(c.config)
	stackID, err := client.DynamoDB(s, c.topicARN).Sync(&awsV1alpha1.DynamoDB{})
	if err != nil {
		c.config.Logger.WithError(err).Errorf("error creating dynamodb '%s'", s.Name)
		return
	}
	c.config.Logger.WithFields(logrus.Fields{
		"stack-id":  stackID,
		"kind":      "dynamodb",
		"name":      s.Name,
		"namespace": s.Namespace,
		"stack-url": fmt.Sprintf("https://console.aws.amazon.com/cloudformation/home?#/stack/detail?stackId=%s", stackID),
	}).Info("created stack")

	err = client

	err = client.DynamoDB(s, c.topicARN).Status(stackId, "CREATE_IN_PROGRESS", "")
	if err != nil {
		c.config.Logger.WithError(err).Error("error updating status")
		return
	}
}

// Update will handle the Update mechanism for the Informer
func (c *dyanmoDB) Update(oldObj, newObj interface{}) {
	oo := oldObj.(*awsV1alpha1.DynamoDB).DeepCopy()
	no := newObj.(*awsV1alpha1.DynamoDB).DeepCopy()

	if no.Status.ResourceStatus == cloudformation.DELETE_COMPLETE {
		c.Create(no)
		return
	}
	if helpers.IsStackComplete(oo.Status.ResourceStatus, false) && !reflect.DeepEqual(oo.Spec, no.Spec) {
		client := stack.NewForConfig(c.config)
		stackID, err := client.DynamoDB(oo, c.topicARN).Sync(no)
		if err != nil {
			c.config.Logger.WithError(err).Errorf("error updating dynamodb '%s'", no.Name)
			return
		}
		c.config.Logger.WithFields(logrus.Fields{
			"stack-id":  stackID,
			"kind":      "dynamodb",
			"name":      no.Name,
			"namespace": no.Namespace,
			"stack-url": fmt.Sprintf("https://console.aws.amazon.com/cloudformation/home?#/stack/detail?stackId=%s", stackID),
		}).Info("updated stack")

		err = client.DynamoDB(no, c.topicARN).Status(stackId, "UPDATE_IN_PROGRESS", "")
		if err != nil {
			c.config.Logger.WithError(err).Error("error updating status")
			return
		}
	}
}

// Delete will handle the Delete mechanism for the Informer
func (c *dyanmoDB) Delete(obj interface{}) {
	s := obj.(*awsV1alpha1.DynamoDB).DeepCopy()

	client := stack.NewForConfig(c.config)
	_, err := client.DynamoDB(s, c.topicARN).Delete()
	if err != nil {
		c.config.Logger.WithError(err).Errorf("error deleting dynamodb '%s'", s.Name)
		return
	}
	c.config.Logger.WithFields(logrus.Fields{
		"stack-id":  s.Status.StackID,
		"kind":      "dynamodb",
		"name":      s.Name,
		"namespace": s.Namespace,
	}).Info("deleted stack")

}

// Register will register the queue status function
func (c *dyanmoDB) Register() {
	queuectrl := queue.New(c.config, c.config.AWSClientset, 10)
	c.topicARN, _ = queuectrl.Register("dynamodbs")
	c.Config.QueueManger.Add(c.topicARN, queuemanager.HandlerFunc(HandleQueue))
}

// Watch will configure the controller manager watcher
func (c *dyanmoDB) Watch(ctx context.Context, namespace string) {
	resourceHandlers := cache.ResourceEventHandlerFuncs{
		AddFunc:    c.Create,
		UpdateFunc: c.Update,
		DeleteFunc: c.Delete,
	}
	manager := controller.Manager("dynamodbs", namespace, resourceHandlers, c.config.AWSClientset.RESTClient())
	manager.Watch(&awsV1alpha1.DynamoDB{}, ctx.Done())
}

// HandleQueue will handle updating the status from the queue
func (c *dyanmoDB) HandleQueue(obj message.BodyInterface) {
	if !obj.Valid() {
		resources, err := c.config.AWSClientset.DynamoDBs("").List(metav1.ListOptions{})
		if err != nil {
			logger.WithError(err).Error("error getting dynamodbs")
			return err
		}
		for _, resource := range resources.Items {
			if resource.Status.StackID == obj.StackID() {
				obj.SetResourceName(resource.Name)
				obj.SetNamespace(resource.Namespace)
			}
		}
	}

	if name != "" && namespace != "" {
		annotations := map[string]string{
			"StackID":      obj.StackID(),
			"StackName":    obj.StackName(),
			"ResourceType": obj.ResourceType(),
		}

		if obj.ResourceStatus() == message.ROLLBACK_COMPLETE {

			obj, err := deleteStack(config, name, namespace, obj.StackID())
			if err != nil {
				return err
			}
			c.config.Recorder.AnnotatedEventf(obj, annotations, corev1.EventTypeWarning, strcase.ToCamel(strings.ToLower(obj.ResourceStatus())), obj.ResourceStatusReason())

		} else if obj.ResourceStatus() == message.DELETE_COMPLETE {

			obj, err := updateStatus(config, obj.ResourceName(), obj.Namespace(), obj.StackID(), obj.ResourceStatus(), obj.ResourceStatusReason())
			if err != nil {
				return err
			}
			config.Recorder.AnnotatedEventf(obj, annotations, corev1.EventTypeWarning, strcase.ToCamel(strings.ToLower(msg.ParsedMessage["ResourceStatus"])), msg.ParsedMessage["ResourceStatusReason"])
			err = incrementRollbackCount(config, name, namespace)
			if err != nil {
				return err
			}

		} else {

			obj, err := updateStatus(config, name, namespace, msg.ParsedMessage["StackId"], msg.ParsedMessage["ResourceStatus"], msg.ParsedMessage["ResourceStatusReason"])
			if err != nil {
				return err
			}
			config.Recorder.AnnotatedEventf(obj, annotations, corev1.EventTypeNormal, strcase.ToCamel(strings.ToLower(msg.ParsedMessage["ResourceStatus"])), msg.ParsedMessage["ResourceStatusReason"])

		}

	}

	return nil
}
