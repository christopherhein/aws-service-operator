package queuemanager

import "github.com/awslabs/aws-service-operator/pkg/message"

// New will return the QueueManager
func New() *QueueManager {
	return &QueueManager{
		handlers: make(map[string]message.Handler),
	}
}

// Get will return the handler func
func (q *QueueManager) Get(name string) (handler Handler, ok bool) {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if handler, ok = q.handlers[name]; ok {
		return handler, ok
	}
	return handler, false
}

// Add will add a new handler func
func (q *QueueManager) Add(name string, handlerFunc Handler) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.handlers[name] = handlerFunc
}

// Keys will return the list of topic ARNs
func (q *QueueManager) Keys() []string {
	q.lock.RLock()
	defer q.lock.RUnlock()
	keys := []string{}
	for key, _ := range q.handlers {
		keys = append(keys, key)
	}
	return keys
}
