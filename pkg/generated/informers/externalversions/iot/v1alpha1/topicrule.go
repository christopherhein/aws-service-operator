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

package v1alpha1

import (
	time "time"

	iotv1alpha1 "awsoperator.io/pkg/apis/iot/v1alpha1"
	versioned "awsoperator.io/pkg/generated/clientset/versioned"
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "awsoperator.io/pkg/generated/listers/iot/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TopicRuleInformer provides access to a shared informer and lister for
// TopicRules.
type TopicRuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TopicRuleLister
}

type topicRuleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTopicRuleInformer constructs a new informer for TopicRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTopicRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTopicRuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTopicRuleInformer constructs a new informer for TopicRule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTopicRuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IotV1alpha1().TopicRules(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IotV1alpha1().TopicRules(namespace).Watch(options)
			},
		},
		&iotv1alpha1.TopicRule{},
		resyncPeriod,
		indexers,
	)
}

func (f *topicRuleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTopicRuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *topicRuleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&iotv1alpha1.TopicRule{}, f.defaultInformer)
}

func (f *topicRuleInformer) Lister() v1alpha1.TopicRuleLister {
	return v1alpha1.NewTopicRuleLister(f.Informer().GetIndexer())
}