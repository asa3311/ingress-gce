/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	frontendconfigv1beta1 "k8s.io/ingress-gce/pkg/apis/frontendconfig/v1beta1"
	versioned "k8s.io/ingress-gce/pkg/frontendconfig/client/clientset/versioned"
	internalinterfaces "k8s.io/ingress-gce/pkg/frontendconfig/client/informers/externalversions/internalinterfaces"
	v1beta1 "k8s.io/ingress-gce/pkg/frontendconfig/client/listers/frontendconfig/v1beta1"
)

// FrontendConfigInformer provides access to a shared informer and lister for
// FrontendConfigs.
type FrontendConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FrontendConfigLister
}

type frontendConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFrontendConfigInformer constructs a new informer for FrontendConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFrontendConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFrontendConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFrontendConfigInformer constructs a new informer for FrontendConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFrontendConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1beta1().FrontendConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1beta1().FrontendConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&frontendconfigv1beta1.FrontendConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *frontendConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFrontendConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *frontendConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&frontendconfigv1beta1.FrontendConfig{}, f.defaultInformer)
}

func (f *frontendConfigInformer) Lister() v1beta1.FrontendConfigLister {
	return v1beta1.NewFrontendConfigLister(f.Informer().GetIndexer())
}
