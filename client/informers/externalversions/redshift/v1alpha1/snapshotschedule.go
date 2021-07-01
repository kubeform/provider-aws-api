/*
Copyright AppsCode Inc. and Contributors

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

	redshiftv1alpha1 "kubeform.dev/provider-aws-api/apis/redshift/v1alpha1"
	versioned "kubeform.dev/provider-aws-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-aws-api/client/listers/redshift/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SnapshotScheduleInformer provides access to a shared informer and lister for
// SnapshotSchedules.
type SnapshotScheduleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SnapshotScheduleLister
}

type snapshotScheduleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSnapshotScheduleInformer constructs a new informer for SnapshotSchedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSnapshotScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSnapshotScheduleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSnapshotScheduleInformer constructs a new informer for SnapshotSchedule type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSnapshotScheduleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RedshiftV1alpha1().SnapshotSchedules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RedshiftV1alpha1().SnapshotSchedules(namespace).Watch(context.TODO(), options)
			},
		},
		&redshiftv1alpha1.SnapshotSchedule{},
		resyncPeriod,
		indexers,
	)
}

func (f *snapshotScheduleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSnapshotScheduleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *snapshotScheduleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&redshiftv1alpha1.SnapshotSchedule{}, f.defaultInformer)
}

func (f *snapshotScheduleInformer) Lister() v1alpha1.SnapshotScheduleLister {
	return v1alpha1.NewSnapshotScheduleLister(f.Informer().GetIndexer())
}
