package config

import (
	"github.com/aws/aws-sdk-go/aws/session"
	awsclient "github.com/awslabs/aws-service-operator/pkg/client/clientset/versioned/typed/service-operator.aws/v1alpha1"
	"github.com/awslabs/aws-service-operator/pkg/queuemanager"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/record"
)

// Config defines the configuration for the operator
type Config struct {
	Region           string
	Kubeconfig       string
	MasterURL        string
	QueueURL         string
	QueueARN         string
	QueueManager     *queuemanager.QueueManager
	AWSSession       *session.Session
	AWSClientset     awsclient.ServiceoperatorV1alpha1Interface
	KubeClientset    kubernetes.Interface
	RESTConfig       *rest.Config
	LoggingConfig    *LoggingConfig
	Logger           *logrus.Entry
	Resources        map[string]bool
	ClusterName      string
	Bucket           string
	AccountID        string
	DefaultNamespace string
	Recorder         record.EventRecorder
}

// LoggingConfig defines the attributes for the logger
type LoggingConfig struct {
	File              string
	Level             string
	DisableTimestamps bool
	FullTimestamps    bool
}

func getKubeconfig(masterURL, kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	}
	return rest.InClusterConfig()
}

func CreateContext(masterURL, kubeconfig string) (
	config *rest.Config,
	clientset *kubernetes.Interface,
	awsclientset *awsclient.Interface,
	err error,
) {
	config, err = getKubeconfig(masterURL, kubeconfig)
	if err != nil {
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}

	awsclientset, err := awsclient.NewForConfig(config)
	if err != nil {
		return
	}

	return
}
