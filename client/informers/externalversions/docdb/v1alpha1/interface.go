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
	internalinterfaces "kubeform.dev/provider-aws-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ClusterInstances returns a ClusterInstanceInformer.
	ClusterInstances() ClusterInstanceInformer
	// ClusterParameterGroups returns a ClusterParameterGroupInformer.
	ClusterParameterGroups() ClusterParameterGroupInformer
	// ClusterSnapshots returns a ClusterSnapshotInformer.
	ClusterSnapshots() ClusterSnapshotInformer
	// GlobalClusters returns a GlobalClusterInformer.
	GlobalClusters() GlobalClusterInformer
	// SubnetGroups returns a SubnetGroupInformer.
	SubnetGroups() SubnetGroupInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterInstances returns a ClusterInstanceInformer.
func (v *version) ClusterInstances() ClusterInstanceInformer {
	return &clusterInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterParameterGroups returns a ClusterParameterGroupInformer.
func (v *version) ClusterParameterGroups() ClusterParameterGroupInformer {
	return &clusterParameterGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterSnapshots returns a ClusterSnapshotInformer.
func (v *version) ClusterSnapshots() ClusterSnapshotInformer {
	return &clusterSnapshotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GlobalClusters returns a GlobalClusterInformer.
func (v *version) GlobalClusters() GlobalClusterInformer {
	return &globalClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubnetGroups returns a SubnetGroupInformer.
func (v *version) SubnetGroups() SubnetGroupInformer {
	return &subnetGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
