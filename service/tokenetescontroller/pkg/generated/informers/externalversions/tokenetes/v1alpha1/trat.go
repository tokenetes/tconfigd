/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	time "time"

	tokenetesv1alpha1 "github.com/tokenetes/tconfigd/tokenetescontroller/pkg/apis/tokenetes/v1alpha1"
	versioned "github.com/tokenetes/tconfigd/tokenetescontroller/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/tokenetes/tconfigd/tokenetescontroller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/tokenetes/tconfigd/tokenetescontroller/pkg/generated/listers/tokenetes/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TraTInformer provides access to a shared informer and lister for
// TraTs.
type TraTInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TraTLister
}

type traTInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTraTInformer constructs a new informer for TraT type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory trattprint and number of connections to the server.
func NewTraTInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTraTInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTraTInformer constructs a new informer for TraT type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory trattprint and number of connections to the server.
func NewFilteredTraTInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TokenetesV1alpha1().TraTs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TokenetesV1alpha1().TraTs(namespace).Watch(context.TODO(), options)
			},
		},
		&tokenetesv1alpha1.TraT{},
		resyncPeriod,
		indexers,
	)
}

func (f *traTInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTraTInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *traTInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tokenetesv1alpha1.TraT{}, f.defaultInformer)
}

func (f *traTInformer) Lister() v1alpha1.TraTLister {
	return v1alpha1.NewTraTLister(f.Informer().GetIndexer())
}
