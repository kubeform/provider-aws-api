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
	// Applications returns a ApplicationInformer.
	Applications() ApplicationInformer
	// CustomLayers returns a CustomLayerInformer.
	CustomLayers() CustomLayerInformer
	// EcsClusterLayers returns a EcsClusterLayerInformer.
	EcsClusterLayers() EcsClusterLayerInformer
	// GangliaLayers returns a GangliaLayerInformer.
	GangliaLayers() GangliaLayerInformer
	// HaproxyLayers returns a HaproxyLayerInformer.
	HaproxyLayers() HaproxyLayerInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// JavaAppLayers returns a JavaAppLayerInformer.
	JavaAppLayers() JavaAppLayerInformer
	// MemcachedLayers returns a MemcachedLayerInformer.
	MemcachedLayers() MemcachedLayerInformer
	// MysqlLayers returns a MysqlLayerInformer.
	MysqlLayers() MysqlLayerInformer
	// NodejsAppLayers returns a NodejsAppLayerInformer.
	NodejsAppLayers() NodejsAppLayerInformer
	// Permissions returns a PermissionInformer.
	Permissions() PermissionInformer
	// PhpAppLayers returns a PhpAppLayerInformer.
	PhpAppLayers() PhpAppLayerInformer
	// RailsAppLayers returns a RailsAppLayerInformer.
	RailsAppLayers() RailsAppLayerInformer
	// RdsDbInstances returns a RdsDbInstanceInformer.
	RdsDbInstances() RdsDbInstanceInformer
	// Stacks returns a StackInformer.
	Stacks() StackInformer
	// StaticWebLayers returns a StaticWebLayerInformer.
	StaticWebLayers() StaticWebLayerInformer
	// UserProfiles returns a UserProfileInformer.
	UserProfiles() UserProfileInformer
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

// Applications returns a ApplicationInformer.
func (v *version) Applications() ApplicationInformer {
	return &applicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CustomLayers returns a CustomLayerInformer.
func (v *version) CustomLayers() CustomLayerInformer {
	return &customLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EcsClusterLayers returns a EcsClusterLayerInformer.
func (v *version) EcsClusterLayers() EcsClusterLayerInformer {
	return &ecsClusterLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// GangliaLayers returns a GangliaLayerInformer.
func (v *version) GangliaLayers() GangliaLayerInformer {
	return &gangliaLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// HaproxyLayers returns a HaproxyLayerInformer.
func (v *version) HaproxyLayers() HaproxyLayerInformer {
	return &haproxyLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// JavaAppLayers returns a JavaAppLayerInformer.
func (v *version) JavaAppLayers() JavaAppLayerInformer {
	return &javaAppLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MemcachedLayers returns a MemcachedLayerInformer.
func (v *version) MemcachedLayers() MemcachedLayerInformer {
	return &memcachedLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MysqlLayers returns a MysqlLayerInformer.
func (v *version) MysqlLayers() MysqlLayerInformer {
	return &mysqlLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NodejsAppLayers returns a NodejsAppLayerInformer.
func (v *version) NodejsAppLayers() NodejsAppLayerInformer {
	return &nodejsAppLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Permissions returns a PermissionInformer.
func (v *version) Permissions() PermissionInformer {
	return &permissionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PhpAppLayers returns a PhpAppLayerInformer.
func (v *version) PhpAppLayers() PhpAppLayerInformer {
	return &phpAppLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RailsAppLayers returns a RailsAppLayerInformer.
func (v *version) RailsAppLayers() RailsAppLayerInformer {
	return &railsAppLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RdsDbInstances returns a RdsDbInstanceInformer.
func (v *version) RdsDbInstances() RdsDbInstanceInformer {
	return &rdsDbInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Stacks returns a StackInformer.
func (v *version) Stacks() StackInformer {
	return &stackInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StaticWebLayers returns a StaticWebLayerInformer.
func (v *version) StaticWebLayers() StaticWebLayerInformer {
	return &staticWebLayerInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// UserProfiles returns a UserProfileInformer.
func (v *version) UserProfiles() UserProfileInformer {
	return &userProfileInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
