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
// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	reflect "reflect"
	sync "sync"
	time "time"

	versioned "awsoperator.io/pkg/generated/clientset/versioned"
	apigateway "awsoperator.io/pkg/generated/informers/externalversions/apigateway"
	apigatewayv2 "awsoperator.io/pkg/generated/informers/externalversions/apigatewayv2"
	applicationautoscaling "awsoperator.io/pkg/generated/informers/externalversions/applicationautoscaling"
	appmesh "awsoperator.io/pkg/generated/informers/externalversions/appmesh"
	appstream "awsoperator.io/pkg/generated/informers/externalversions/appstream"
	appsync "awsoperator.io/pkg/generated/informers/externalversions/appsync"
	ask "awsoperator.io/pkg/generated/informers/externalversions/ask"
	athena "awsoperator.io/pkg/generated/informers/externalversions/athena"
	autoscaling "awsoperator.io/pkg/generated/informers/externalversions/autoscaling"
	autoscalingplans "awsoperator.io/pkg/generated/informers/externalversions/autoscalingplans"
	batch "awsoperator.io/pkg/generated/informers/externalversions/batch"
	budgets "awsoperator.io/pkg/generated/informers/externalversions/budgets"
	cdk "awsoperator.io/pkg/generated/informers/externalversions/cdk"
	certificatemanager "awsoperator.io/pkg/generated/informers/externalversions/certificatemanager"
	cloud9 "awsoperator.io/pkg/generated/informers/externalversions/cloud9"
	cloudformation "awsoperator.io/pkg/generated/informers/externalversions/cloudformation"
	cloudfront "awsoperator.io/pkg/generated/informers/externalversions/cloudfront"
	cloudtrail "awsoperator.io/pkg/generated/informers/externalversions/cloudtrail"
	cloudwatch "awsoperator.io/pkg/generated/informers/externalversions/cloudwatch"
	codebuild "awsoperator.io/pkg/generated/informers/externalversions/codebuild"
	codecommit "awsoperator.io/pkg/generated/informers/externalversions/codecommit"
	codedeploy "awsoperator.io/pkg/generated/informers/externalversions/codedeploy"
	codepipeline "awsoperator.io/pkg/generated/informers/externalversions/codepipeline"
	cognito "awsoperator.io/pkg/generated/informers/externalversions/cognito"
	config "awsoperator.io/pkg/generated/informers/externalversions/config"
	datapipeline "awsoperator.io/pkg/generated/informers/externalversions/datapipeline"
	dax "awsoperator.io/pkg/generated/informers/externalversions/dax"
	directoryservice "awsoperator.io/pkg/generated/informers/externalversions/directoryservice"
	dlm "awsoperator.io/pkg/generated/informers/externalversions/dlm"
	dms "awsoperator.io/pkg/generated/informers/externalversions/dms"
	docdb "awsoperator.io/pkg/generated/informers/externalversions/docdb"
	dynamodb "awsoperator.io/pkg/generated/informers/externalversions/dynamodb"
	ec2 "awsoperator.io/pkg/generated/informers/externalversions/ec2"
	ecr "awsoperator.io/pkg/generated/informers/externalversions/ecr"
	ecs "awsoperator.io/pkg/generated/informers/externalversions/ecs"
	efs "awsoperator.io/pkg/generated/informers/externalversions/efs"
	eks "awsoperator.io/pkg/generated/informers/externalversions/eks"
	elasticache "awsoperator.io/pkg/generated/informers/externalversions/elasticache"
	elasticbeanstalk "awsoperator.io/pkg/generated/informers/externalversions/elasticbeanstalk"
	elasticloadbalancing "awsoperator.io/pkg/generated/informers/externalversions/elasticloadbalancing"
	elasticloadbalancingv2 "awsoperator.io/pkg/generated/informers/externalversions/elasticloadbalancingv2"
	elasticsearch "awsoperator.io/pkg/generated/informers/externalversions/elasticsearch"
	emr "awsoperator.io/pkg/generated/informers/externalversions/emr"
	events "awsoperator.io/pkg/generated/informers/externalversions/events"
	fsx "awsoperator.io/pkg/generated/informers/externalversions/fsx"
	gamelift "awsoperator.io/pkg/generated/informers/externalversions/gamelift"
	glue "awsoperator.io/pkg/generated/informers/externalversions/glue"
	greengrass "awsoperator.io/pkg/generated/informers/externalversions/greengrass"
	guardduty "awsoperator.io/pkg/generated/informers/externalversions/guardduty"
	iam "awsoperator.io/pkg/generated/informers/externalversions/iam"
	inspector "awsoperator.io/pkg/generated/informers/externalversions/inspector"
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
	iot "awsoperator.io/pkg/generated/informers/externalversions/iot"
	iot1click "awsoperator.io/pkg/generated/informers/externalversions/iot1click"
	iotanalytics "awsoperator.io/pkg/generated/informers/externalversions/iotanalytics"
	kinesis "awsoperator.io/pkg/generated/informers/externalversions/kinesis"
	kinesisanalytics "awsoperator.io/pkg/generated/informers/externalversions/kinesisanalytics"
	kinesisanalyticsv2 "awsoperator.io/pkg/generated/informers/externalversions/kinesisanalyticsv2"
	kinesisfirehose "awsoperator.io/pkg/generated/informers/externalversions/kinesisfirehose"
	kms "awsoperator.io/pkg/generated/informers/externalversions/kms"
	lambda "awsoperator.io/pkg/generated/informers/externalversions/lambda"
	logs "awsoperator.io/pkg/generated/informers/externalversions/logs"
	mediastore "awsoperator.io/pkg/generated/informers/externalversions/mediastore"
	mq "awsoperator.io/pkg/generated/informers/externalversions/mq"
	neptune "awsoperator.io/pkg/generated/informers/externalversions/neptune"
	opsworks "awsoperator.io/pkg/generated/informers/externalversions/opsworks"
	opsworkscm "awsoperator.io/pkg/generated/informers/externalversions/opsworkscm"
	pinpointemail "awsoperator.io/pkg/generated/informers/externalversions/pinpointemail"
	ram "awsoperator.io/pkg/generated/informers/externalversions/ram"
	rds "awsoperator.io/pkg/generated/informers/externalversions/rds"
	redshift "awsoperator.io/pkg/generated/informers/externalversions/redshift"
	robomaker "awsoperator.io/pkg/generated/informers/externalversions/robomaker"
	route53 "awsoperator.io/pkg/generated/informers/externalversions/route53"
	route53resolver "awsoperator.io/pkg/generated/informers/externalversions/route53resolver"
	s3 "awsoperator.io/pkg/generated/informers/externalversions/s3"
	sagemaker "awsoperator.io/pkg/generated/informers/externalversions/sagemaker"
	sdb "awsoperator.io/pkg/generated/informers/externalversions/sdb"
	secretsmanager "awsoperator.io/pkg/generated/informers/externalversions/secretsmanager"
	servicecatalog "awsoperator.io/pkg/generated/informers/externalversions/servicecatalog"
	servicediscovery "awsoperator.io/pkg/generated/informers/externalversions/servicediscovery"
	ses "awsoperator.io/pkg/generated/informers/externalversions/ses"
	sns "awsoperator.io/pkg/generated/informers/externalversions/sns"
	sqs "awsoperator.io/pkg/generated/informers/externalversions/sqs"
	ssm "awsoperator.io/pkg/generated/informers/externalversions/ssm"
	stepfunctions "awsoperator.io/pkg/generated/informers/externalversions/stepfunctions"
	transfer "awsoperator.io/pkg/generated/informers/externalversions/transfer"
	waf "awsoperator.io/pkg/generated/informers/externalversions/waf"
	wafregional "awsoperator.io/pkg/generated/informers/externalversions/wafregional"
	workspaces "awsoperator.io/pkg/generated/informers/externalversions/workspaces"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// SharedInformerOption defines the functional option type for SharedInformerFactory.
type SharedInformerOption func(*sharedInformerFactory) *sharedInformerFactory

type sharedInformerFactory struct {
	client           versioned.Interface
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	lock             sync.Mutex
	defaultResync    time.Duration
	customResync     map[reflect.Type]time.Duration

	informers map[reflect.Type]cache.SharedIndexInformer
	// startedInformers is used for tracking which informers have been started.
	// This allows Start() to be called multiple times safely.
	startedInformers map[reflect.Type]bool
}

// WithCustomResyncConfig sets a custom resync period for the specified informer types.
func WithCustomResyncConfig(resyncConfig map[v1.Object]time.Duration) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		for k, v := range resyncConfig {
			factory.customResync[reflect.TypeOf(k)] = v
		}
		return factory
	}
}

// WithTweakListOptions sets a custom filter on all listers of the configured SharedInformerFactory.
func WithTweakListOptions(tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.tweakListOptions = tweakListOptions
		return factory
	}
}

// WithNamespace limits the SharedInformerFactory to the specified namespace.
func WithNamespace(namespace string) SharedInformerOption {
	return func(factory *sharedInformerFactory) *sharedInformerFactory {
		factory.namespace = namespace
		return factory
	}
}

// NewSharedInformerFactory constructs a new instance of sharedInformerFactory for all namespaces.
func NewSharedInformerFactory(client versioned.Interface, defaultResync time.Duration) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync)
}

// NewFilteredSharedInformerFactory constructs a new instance of sharedInformerFactory.
// Listers obtained via this SharedInformerFactory will be subject to the same filters
// as specified here.
// Deprecated: Please use NewSharedInformerFactoryWithOptions instead
func NewFilteredSharedInformerFactory(client versioned.Interface, defaultResync time.Duration, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) SharedInformerFactory {
	return NewSharedInformerFactoryWithOptions(client, defaultResync, WithNamespace(namespace), WithTweakListOptions(tweakListOptions))
}

// NewSharedInformerFactoryWithOptions constructs a new instance of a SharedInformerFactory with additional options.
func NewSharedInformerFactoryWithOptions(client versioned.Interface, defaultResync time.Duration, options ...SharedInformerOption) SharedInformerFactory {
	factory := &sharedInformerFactory{
		client:           client,
		namespace:        v1.NamespaceAll,
		defaultResync:    defaultResync,
		informers:        make(map[reflect.Type]cache.SharedIndexInformer),
		startedInformers: make(map[reflect.Type]bool),
		customResync:     make(map[reflect.Type]time.Duration),
	}

	// Apply all options
	for _, opt := range options {
		factory = opt(factory)
	}

	return factory
}

// Start initializes all requested informers.
func (f *sharedInformerFactory) Start(stopCh <-chan struct{}) {
	f.lock.Lock()
	defer f.lock.Unlock()

	for informerType, informer := range f.informers {
		if !f.startedInformers[informerType] {
			go informer.Run(stopCh)
			f.startedInformers[informerType] = true
		}
	}
}

// WaitForCacheSync waits for all started informers' cache were synced.
func (f *sharedInformerFactory) WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool {
	informers := func() map[reflect.Type]cache.SharedIndexInformer {
		f.lock.Lock()
		defer f.lock.Unlock()

		informers := map[reflect.Type]cache.SharedIndexInformer{}
		for informerType, informer := range f.informers {
			if f.startedInformers[informerType] {
				informers[informerType] = informer
			}
		}
		return informers
	}()

	res := map[reflect.Type]bool{}
	for informType, informer := range informers {
		res[informType] = cache.WaitForCacheSync(stopCh, informer.HasSynced)
	}
	return res
}

// InternalInformerFor returns the SharedIndexInformer for obj using an internal
// client.
func (f *sharedInformerFactory) InformerFor(obj runtime.Object, newFunc internalinterfaces.NewInformerFunc) cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(obj)
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}

	resyncPeriod, exists := f.customResync[informerType]
	if !exists {
		resyncPeriod = f.defaultResync
	}

	informer = newFunc(f.client, resyncPeriod)
	f.informers[informerType] = informer

	return informer
}

// SharedInformerFactory provides shared informers for resources in all known
// API group versions.
type SharedInformerFactory interface {
	internalinterfaces.SharedInformerFactory
	ForResource(resource schema.GroupVersionResource) (GenericInformer, error)
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool

	Apigateway() apigateway.Interface
	Apigatewayv2() apigatewayv2.Interface
	Applicationautoscaling() applicationautoscaling.Interface
	Appmesh() appmesh.Interface
	Appstream() appstream.Interface
	Appsync() appsync.Interface
	Ask() ask.Interface
	Athena() athena.Interface
	Autoscaling() autoscaling.Interface
	Autoscalingplans() autoscalingplans.Interface
	Batch() batch.Interface
	Budgets() budgets.Interface
	Cdk() cdk.Interface
	Certificatemanager() certificatemanager.Interface
	Cloud9() cloud9.Interface
	Cloudformation() cloudformation.Interface
	Cloudfront() cloudfront.Interface
	Cloudtrail() cloudtrail.Interface
	Cloudwatch() cloudwatch.Interface
	Codebuild() codebuild.Interface
	Codecommit() codecommit.Interface
	Codedeploy() codedeploy.Interface
	Codepipeline() codepipeline.Interface
	Cognito() cognito.Interface
	Config() config.Interface
	Datapipeline() datapipeline.Interface
	Dax() dax.Interface
	Directoryservice() directoryservice.Interface
	Dlm() dlm.Interface
	Dms() dms.Interface
	Docdb() docdb.Interface
	Dynamodb() dynamodb.Interface
	Ec2() ec2.Interface
	Ecr() ecr.Interface
	Ecs() ecs.Interface
	Efs() efs.Interface
	Eks() eks.Interface
	Elasticache() elasticache.Interface
	Elasticbeanstalk() elasticbeanstalk.Interface
	Elasticloadbalancing() elasticloadbalancing.Interface
	Elasticloadbalancingv2() elasticloadbalancingv2.Interface
	Elasticsearch() elasticsearch.Interface
	Emr() emr.Interface
	Events() events.Interface
	Fsx() fsx.Interface
	Gamelift() gamelift.Interface
	Glue() glue.Interface
	Greengrass() greengrass.Interface
	Guardduty() guardduty.Interface
	Iam() iam.Interface
	Inspector() inspector.Interface
	Iot() iot.Interface
	Iot1click() iot1click.Interface
	Iotanalytics() iotanalytics.Interface
	Kinesis() kinesis.Interface
	Kinesisanalytics() kinesisanalytics.Interface
	Kinesisanalyticsv2() kinesisanalyticsv2.Interface
	Kinesisfirehose() kinesisfirehose.Interface
	Kms() kms.Interface
	Lambda() lambda.Interface
	Logs() logs.Interface
	Mediastore() mediastore.Interface
	Mq() mq.Interface
	Neptune() neptune.Interface
	Opsworks() opsworks.Interface
	Opsworkscm() opsworkscm.Interface
	Pinpointemail() pinpointemail.Interface
	Ram() ram.Interface
	Rds() rds.Interface
	Redshift() redshift.Interface
	Robomaker() robomaker.Interface
	Route53() route53.Interface
	Route53resolver() route53resolver.Interface
	S3() s3.Interface
	Sagemaker() sagemaker.Interface
	Sdb() sdb.Interface
	Secretsmanager() secretsmanager.Interface
	Servicecatalog() servicecatalog.Interface
	Servicediscovery() servicediscovery.Interface
	Ses() ses.Interface
	Sns() sns.Interface
	Sqs() sqs.Interface
	Ssm() ssm.Interface
	Stepfunctions() stepfunctions.Interface
	Transfer() transfer.Interface
	Waf() waf.Interface
	Wafregional() wafregional.Interface
	Workspaces() workspaces.Interface
}

func (f *sharedInformerFactory) Apigateway() apigateway.Interface {
	return apigateway.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Apigatewayv2() apigatewayv2.Interface {
	return apigatewayv2.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Applicationautoscaling() applicationautoscaling.Interface {
	return applicationautoscaling.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Appmesh() appmesh.Interface {
	return appmesh.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Appstream() appstream.Interface {
	return appstream.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Appsync() appsync.Interface {
	return appsync.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ask() ask.Interface {
	return ask.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Athena() athena.Interface {
	return athena.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Autoscaling() autoscaling.Interface {
	return autoscaling.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Autoscalingplans() autoscalingplans.Interface {
	return autoscalingplans.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Batch() batch.Interface {
	return batch.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Budgets() budgets.Interface {
	return budgets.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cdk() cdk.Interface {
	return cdk.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Certificatemanager() certificatemanager.Interface {
	return certificatemanager.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloud9() cloud9.Interface {
	return cloud9.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloudformation() cloudformation.Interface {
	return cloudformation.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloudfront() cloudfront.Interface {
	return cloudfront.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloudtrail() cloudtrail.Interface {
	return cloudtrail.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cloudwatch() cloudwatch.Interface {
	return cloudwatch.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Codebuild() codebuild.Interface {
	return codebuild.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Codecommit() codecommit.Interface {
	return codecommit.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Codedeploy() codedeploy.Interface {
	return codedeploy.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Codepipeline() codepipeline.Interface {
	return codepipeline.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Cognito() cognito.Interface {
	return cognito.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Config() config.Interface {
	return config.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Datapipeline() datapipeline.Interface {
	return datapipeline.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dax() dax.Interface {
	return dax.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Directoryservice() directoryservice.Interface {
	return directoryservice.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dlm() dlm.Interface {
	return dlm.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dms() dms.Interface {
	return dms.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Docdb() docdb.Interface {
	return docdb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Dynamodb() dynamodb.Interface {
	return dynamodb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ec2() ec2.Interface {
	return ec2.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ecr() ecr.Interface {
	return ecr.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ecs() ecs.Interface {
	return ecs.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Efs() efs.Interface {
	return efs.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Eks() eks.Interface {
	return eks.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Elasticache() elasticache.Interface {
	return elasticache.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Elasticbeanstalk() elasticbeanstalk.Interface {
	return elasticbeanstalk.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Elasticloadbalancing() elasticloadbalancing.Interface {
	return elasticloadbalancing.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Elasticloadbalancingv2() elasticloadbalancingv2.Interface {
	return elasticloadbalancingv2.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Elasticsearch() elasticsearch.Interface {
	return elasticsearch.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Emr() emr.Interface {
	return emr.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Events() events.Interface {
	return events.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Fsx() fsx.Interface {
	return fsx.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Gamelift() gamelift.Interface {
	return gamelift.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Glue() glue.Interface {
	return glue.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Greengrass() greengrass.Interface {
	return greengrass.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Guardduty() guardduty.Interface {
	return guardduty.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Iam() iam.Interface {
	return iam.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Inspector() inspector.Interface {
	return inspector.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Iot() iot.Interface {
	return iot.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Iot1click() iot1click.Interface {
	return iot1click.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Iotanalytics() iotanalytics.Interface {
	return iotanalytics.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Kinesis() kinesis.Interface {
	return kinesis.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Kinesisanalytics() kinesisanalytics.Interface {
	return kinesisanalytics.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Kinesisanalyticsv2() kinesisanalyticsv2.Interface {
	return kinesisanalyticsv2.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Kinesisfirehose() kinesisfirehose.Interface {
	return kinesisfirehose.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Kms() kms.Interface {
	return kms.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Lambda() lambda.Interface {
	return lambda.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Logs() logs.Interface {
	return logs.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Mediastore() mediastore.Interface {
	return mediastore.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Mq() mq.Interface {
	return mq.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Neptune() neptune.Interface {
	return neptune.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Opsworks() opsworks.Interface {
	return opsworks.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Opsworkscm() opsworkscm.Interface {
	return opsworkscm.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Pinpointemail() pinpointemail.Interface {
	return pinpointemail.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ram() ram.Interface {
	return ram.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Rds() rds.Interface {
	return rds.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Redshift() redshift.Interface {
	return redshift.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Robomaker() robomaker.Interface {
	return robomaker.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Route53() route53.Interface {
	return route53.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Route53resolver() route53resolver.Interface {
	return route53resolver.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) S3() s3.Interface {
	return s3.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sagemaker() sagemaker.Interface {
	return sagemaker.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sdb() sdb.Interface {
	return sdb.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Secretsmanager() secretsmanager.Interface {
	return secretsmanager.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Servicecatalog() servicecatalog.Interface {
	return servicecatalog.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Servicediscovery() servicediscovery.Interface {
	return servicediscovery.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ses() ses.Interface {
	return ses.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sns() sns.Interface {
	return sns.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Sqs() sqs.Interface {
	return sqs.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Ssm() ssm.Interface {
	return ssm.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Stepfunctions() stepfunctions.Interface {
	return stepfunctions.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Transfer() transfer.Interface {
	return transfer.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Waf() waf.Interface {
	return waf.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Wafregional() wafregional.Interface {
	return wafregional.New(f, f.namespace, f.tweakListOptions)
}

func (f *sharedInformerFactory) Workspaces() workspaces.Interface {
	return workspaces.New(f, f.namespace, f.tweakListOptions)
}