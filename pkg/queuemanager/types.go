package queuemanager

import (
	"sync"

	"github.com/awslabs/aws-service-operator/pkg/message"
)

// Queue Manager allows you to register topics and a handler function
type QueueManager struct {
	lock     sync.RWMutex
	handlers map[string]message.Handler
}
