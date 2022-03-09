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
	// Backups returns a BackupInformer.
	Backups() BackupInformer
	// DataRepositoryAssociations returns a DataRepositoryAssociationInformer.
	DataRepositoryAssociations() DataRepositoryAssociationInformer
	// LustreFileSystems returns a LustreFileSystemInformer.
	LustreFileSystems() LustreFileSystemInformer
	// OntapFileSystems returns a OntapFileSystemInformer.
	OntapFileSystems() OntapFileSystemInformer
	// OntapStorageVirtualMachines returns a OntapStorageVirtualMachineInformer.
	OntapStorageVirtualMachines() OntapStorageVirtualMachineInformer
	// OntapVolumes returns a OntapVolumeInformer.
	OntapVolumes() OntapVolumeInformer
	// OpenzfsFileSystems returns a OpenzfsFileSystemInformer.
	OpenzfsFileSystems() OpenzfsFileSystemInformer
	// OpenzfsSnapshots returns a OpenzfsSnapshotInformer.
	OpenzfsSnapshots() OpenzfsSnapshotInformer
	// OpenzfsVolumes returns a OpenzfsVolumeInformer.
	OpenzfsVolumes() OpenzfsVolumeInformer
	// WindowsFileSystems returns a WindowsFileSystemInformer.
	WindowsFileSystems() WindowsFileSystemInformer
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

// Backups returns a BackupInformer.
func (v *version) Backups() BackupInformer {
	return &backupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DataRepositoryAssociations returns a DataRepositoryAssociationInformer.
func (v *version) DataRepositoryAssociations() DataRepositoryAssociationInformer {
	return &dataRepositoryAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LustreFileSystems returns a LustreFileSystemInformer.
func (v *version) LustreFileSystems() LustreFileSystemInformer {
	return &lustreFileSystemInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OntapFileSystems returns a OntapFileSystemInformer.
func (v *version) OntapFileSystems() OntapFileSystemInformer {
	return &ontapFileSystemInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OntapStorageVirtualMachines returns a OntapStorageVirtualMachineInformer.
func (v *version) OntapStorageVirtualMachines() OntapStorageVirtualMachineInformer {
	return &ontapStorageVirtualMachineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OntapVolumes returns a OntapVolumeInformer.
func (v *version) OntapVolumes() OntapVolumeInformer {
	return &ontapVolumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenzfsFileSystems returns a OpenzfsFileSystemInformer.
func (v *version) OpenzfsFileSystems() OpenzfsFileSystemInformer {
	return &openzfsFileSystemInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenzfsSnapshots returns a OpenzfsSnapshotInformer.
func (v *version) OpenzfsSnapshots() OpenzfsSnapshotInformer {
	return &openzfsSnapshotInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// OpenzfsVolumes returns a OpenzfsVolumeInformer.
func (v *version) OpenzfsVolumes() OpenzfsVolumeInformer {
	return &openzfsVolumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WindowsFileSystems returns a WindowsFileSystemInformer.
func (v *version) WindowsFileSystems() WindowsFileSystemInformer {
	return &windowsFileSystemInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
