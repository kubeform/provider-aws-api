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
	// Agents returns a AgentInformer.
	Agents() AgentInformer
	// LocationEfses returns a LocationEfsInformer.
	LocationEfses() LocationEfsInformer
	// LocationFsxLustreFileSystems returns a LocationFsxLustreFileSystemInformer.
	LocationFsxLustreFileSystems() LocationFsxLustreFileSystemInformer
	// LocationFsxWindowsFileSystems returns a LocationFsxWindowsFileSystemInformer.
	LocationFsxWindowsFileSystems() LocationFsxWindowsFileSystemInformer
	// LocationHdfses returns a LocationHdfsInformer.
	LocationHdfses() LocationHdfsInformer
	// LocationNfses returns a LocationNfsInformer.
	LocationNfses() LocationNfsInformer
	// LocationS3s returns a LocationS3Informer.
	LocationS3s() LocationS3Informer
	// LocationSmbs returns a LocationSmbInformer.
	LocationSmbs() LocationSmbInformer
	// Tasks returns a TaskInformer.
	Tasks() TaskInformer
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

// Agents returns a AgentInformer.
func (v *version) Agents() AgentInformer {
	return &agentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LocationEfses returns a LocationEfsInformer.
func (v *version) LocationEfses() LocationEfsInformer {
	return &locationEfsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LocationFsxLustreFileSystems returns a LocationFsxLustreFileSystemInformer.
func (v *version) LocationFsxLustreFileSystems() LocationFsxLustreFileSystemInformer {
	return &locationFsxLustreFileSystemInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LocationFsxWindowsFileSystems returns a LocationFsxWindowsFileSystemInformer.
func (v *version) LocationFsxWindowsFileSystems() LocationFsxWindowsFileSystemInformer {
	return &locationFsxWindowsFileSystemInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LocationHdfses returns a LocationHdfsInformer.
func (v *version) LocationHdfses() LocationHdfsInformer {
	return &locationHdfsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LocationNfses returns a LocationNfsInformer.
func (v *version) LocationNfses() LocationNfsInformer {
	return &locationNfsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LocationS3s returns a LocationS3Informer.
func (v *version) LocationS3s() LocationS3Informer {
	return &locationS3Informer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LocationSmbs returns a LocationSmbInformer.
func (v *version) LocationSmbs() LocationSmbInformer {
	return &locationSmbInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Tasks returns a TaskInformer.
func (v *version) Tasks() TaskInformer {
	return &taskInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
