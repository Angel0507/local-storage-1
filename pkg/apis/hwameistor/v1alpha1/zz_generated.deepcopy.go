//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Contributors to The HwameiStor project.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessibilityTopology) DeepCopyInto(out *AccessibilityTopology) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessibilityTopology.
func (in *AccessibilityTopology) DeepCopy() *AccessibilityTopology {
	if in == nil {
		return nil
	}
	out := new(AccessibilityTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DRBDSystemConfig) DeepCopyInto(out *DRBDSystemConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DRBDSystemConfig.
func (in *DRBDSystemConfig) DeepCopy() *DRBDSystemConfig {
	if in == nil {
		return nil
	}
	out := new(DRBDSystemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HAState) DeepCopyInto(out *HAState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HAState.
func (in *HAState) DeepCopy() *HAState {
	if in == nil {
		return nil
	}
	out := new(HAState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalDisk) DeepCopyInto(out *LocalDisk) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalDisk.
func (in *LocalDisk) DeepCopy() *LocalDisk {
	if in == nil {
		return nil
	}
	out := new(LocalDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalPool) DeepCopyInto(out *LocalPool) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]LocalDisk, len(*in))
		copy(*out, *in)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalPool.
func (in *LocalPool) DeepCopy() *LocalPool {
	if in == nil {
		return nil
	}
	out := new(LocalPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageNode) DeepCopyInto(out *LocalStorageNode) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageNode.
func (in *LocalStorageNode) DeepCopy() *LocalStorageNode {
	if in == nil {
		return nil
	}
	out := new(LocalStorageNode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalStorageNode) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageNodeList) DeepCopyInto(out *LocalStorageNodeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalStorageNode, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageNodeList.
func (in *LocalStorageNodeList) DeepCopy() *LocalStorageNodeList {
	if in == nil {
		return nil
	}
	out := new(LocalStorageNodeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalStorageNodeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageNodeSpec) DeepCopyInto(out *LocalStorageNodeSpec) {
	*out = *in
	out.Topo = in.Topo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageNodeSpec.
func (in *LocalStorageNodeSpec) DeepCopy() *LocalStorageNodeSpec {
	if in == nil {
		return nil
	}
	out := new(LocalStorageNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalStorageNodeStatus) DeepCopyInto(out *LocalStorageNodeStatus) {
	*out = *in
	if in.Pools != nil {
		in, out := &in.Pools, &out.Pools
		*out = make(map[string]LocalPool, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalStorageNodeStatus.
func (in *LocalStorageNodeStatus) DeepCopy() *LocalStorageNodeStatus {
	if in == nil {
		return nil
	}
	out := new(LocalStorageNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolume) DeepCopyInto(out *LocalVolume) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolume.
func (in *LocalVolume) DeepCopy() *LocalVolume {
	if in == nil {
		return nil
	}
	out := new(LocalVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolume) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeConvert) DeepCopyInto(out *LocalVolumeConvert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeConvert.
func (in *LocalVolumeConvert) DeepCopy() *LocalVolumeConvert {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeConvert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeConvert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeConvertList) DeepCopyInto(out *LocalVolumeConvertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeConvert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeConvertList.
func (in *LocalVolumeConvertList) DeepCopy() *LocalVolumeConvertList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeConvertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeConvertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeConvertSpec) DeepCopyInto(out *LocalVolumeConvertSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeConvertSpec.
func (in *LocalVolumeConvertSpec) DeepCopy() *LocalVolumeConvertSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeConvertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeConvertStatus) DeepCopyInto(out *LocalVolumeConvertStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeConvertStatus.
func (in *LocalVolumeConvertStatus) DeepCopy() *LocalVolumeConvertStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeConvertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeExpand) DeepCopyInto(out *LocalVolumeExpand) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeExpand.
func (in *LocalVolumeExpand) DeepCopy() *LocalVolumeExpand {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeExpand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeExpand) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeExpandList) DeepCopyInto(out *LocalVolumeExpandList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeExpand, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeExpandList.
func (in *LocalVolumeExpandList) DeepCopy() *LocalVolumeExpandList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeExpandList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeExpandList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeExpandSpec) DeepCopyInto(out *LocalVolumeExpandSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeExpandSpec.
func (in *LocalVolumeExpandSpec) DeepCopy() *LocalVolumeExpandSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeExpandSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeExpandStatus) DeepCopyInto(out *LocalVolumeExpandStatus) {
	*out = *in
	if in.Subs != nil {
		in, out := &in.Subs, &out.Subs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeExpandStatus.
func (in *LocalVolumeExpandStatus) DeepCopy() *LocalVolumeExpandStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeExpandStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroup) DeepCopyInto(out *LocalVolumeGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroup.
func (in *LocalVolumeGroup) DeepCopy() *LocalVolumeGroup {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupConvert) DeepCopyInto(out *LocalVolumeGroupConvert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupConvert.
func (in *LocalVolumeGroupConvert) DeepCopy() *LocalVolumeGroupConvert {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupConvert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeGroupConvert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupConvertList) DeepCopyInto(out *LocalVolumeGroupConvertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeGroupConvert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupConvertList.
func (in *LocalVolumeGroupConvertList) DeepCopy() *LocalVolumeGroupConvertList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupConvertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeGroupConvertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupConvertSpec) DeepCopyInto(out *LocalVolumeGroupConvertSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupConvertSpec.
func (in *LocalVolumeGroupConvertSpec) DeepCopy() *LocalVolumeGroupConvertSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupConvertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupConvertStatus) DeepCopyInto(out *LocalVolumeGroupConvertStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupConvertStatus.
func (in *LocalVolumeGroupConvertStatus) DeepCopy() *LocalVolumeGroupConvertStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupConvertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupList) DeepCopyInto(out *LocalVolumeGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupList.
func (in *LocalVolumeGroupList) DeepCopy() *LocalVolumeGroupList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupMigrate) DeepCopyInto(out *LocalVolumeGroupMigrate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupMigrate.
func (in *LocalVolumeGroupMigrate) DeepCopy() *LocalVolumeGroupMigrate {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupMigrate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeGroupMigrate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupMigrateList) DeepCopyInto(out *LocalVolumeGroupMigrateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeGroupMigrate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupMigrateList.
func (in *LocalVolumeGroupMigrateList) DeepCopy() *LocalVolumeGroupMigrateList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupMigrateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeGroupMigrateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupMigrateSpec) DeepCopyInto(out *LocalVolumeGroupMigrateSpec) {
	*out = *in
	if in.TargetNodesNames != nil {
		in, out := &in.TargetNodesNames, &out.TargetNodesNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SourceNodesNames != nil {
		in, out := &in.SourceNodesNames, &out.SourceNodesNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupMigrateSpec.
func (in *LocalVolumeGroupMigrateSpec) DeepCopy() *LocalVolumeGroupMigrateSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupMigrateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupMigrateStatus) DeepCopyInto(out *LocalVolumeGroupMigrateStatus) {
	*out = *in
	if in.TargetNodesNames != nil {
		in, out := &in.TargetNodesNames, &out.TargetNodesNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupMigrateStatus.
func (in *LocalVolumeGroupMigrateStatus) DeepCopy() *LocalVolumeGroupMigrateStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupMigrateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupSpec) DeepCopyInto(out *LocalVolumeGroupSpec) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]VolumeInfo, len(*in))
		copy(*out, *in)
	}
	in.Accessibility.DeepCopyInto(&out.Accessibility)
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupSpec.
func (in *LocalVolumeGroupSpec) DeepCopy() *LocalVolumeGroupSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeGroupStatus) DeepCopyInto(out *LocalVolumeGroupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeGroupStatus.
func (in *LocalVolumeGroupStatus) DeepCopy() *LocalVolumeGroupStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeList) DeepCopyInto(out *LocalVolumeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeList.
func (in *LocalVolumeList) DeepCopy() *LocalVolumeList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeMigrate) DeepCopyInto(out *LocalVolumeMigrate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeMigrate.
func (in *LocalVolumeMigrate) DeepCopy() *LocalVolumeMigrate {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeMigrate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeMigrate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeMigrateList) DeepCopyInto(out *LocalVolumeMigrateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeMigrate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeMigrateList.
func (in *LocalVolumeMigrateList) DeepCopy() *LocalVolumeMigrateList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeMigrateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeMigrateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeMigrateSpec) DeepCopyInto(out *LocalVolumeMigrateSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeMigrateSpec.
func (in *LocalVolumeMigrateSpec) DeepCopy() *LocalVolumeMigrateSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeMigrateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeMigrateStatus) DeepCopyInto(out *LocalVolumeMigrateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeMigrateStatus.
func (in *LocalVolumeMigrateStatus) DeepCopy() *LocalVolumeMigrateStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeMigrateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeReplica) DeepCopyInto(out *LocalVolumeReplica) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeReplica.
func (in *LocalVolumeReplica) DeepCopy() *LocalVolumeReplica {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeReplica)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeReplica) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeReplicaList) DeepCopyInto(out *LocalVolumeReplicaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LocalVolumeReplica, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeReplicaList.
func (in *LocalVolumeReplicaList) DeepCopy() *LocalVolumeReplicaList {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeReplicaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LocalVolumeReplicaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeReplicaSpec) DeepCopyInto(out *LocalVolumeReplicaSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeReplicaSpec.
func (in *LocalVolumeReplicaSpec) DeepCopy() *LocalVolumeReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeReplicaStatus) DeepCopyInto(out *LocalVolumeReplicaStatus) {
	*out = *in
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HAState != nil {
		in, out := &in.HAState, &out.HAState
		*out = new(HAState)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeReplicaStatus.
func (in *LocalVolumeReplicaStatus) DeepCopy() *LocalVolumeReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeSpec) DeepCopyInto(out *LocalVolumeSpec) {
	*out = *in
	in.Accessibility.DeepCopyInto(&out.Accessibility)
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(VolumeConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeSpec.
func (in *LocalVolumeSpec) DeepCopy() *LocalVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalVolumeStatus) DeepCopyInto(out *LocalVolumeStatus) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalVolumeStatus.
func (in *LocalVolumeStatus) DeepCopy() *LocalVolumeStatus {
	if in == nil {
		return nil
	}
	out := new(LocalVolumeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeConfig) DeepCopyInto(out *NodeConfig) {
	*out = *in
	if in.Topology != nil {
		in, out := &in.Topology, &out.Topology
		*out = new(Topology)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeConfig.
func (in *NodeConfig) DeepCopy() *NodeConfig {
	if in == nil {
		return nil
	}
	out := new(NodeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemConfig) DeepCopyInto(out *SystemConfig) {
	*out = *in
	if in.DRBD != nil {
		in, out := &in.DRBD, &out.DRBD
		*out = new(DRBDSystemConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemConfig.
func (in *SystemConfig) DeepCopy() *SystemConfig {
	if in == nil {
		return nil
	}
	out := new(SystemConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topology) DeepCopyInto(out *Topology) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topology.
func (in *Topology) DeepCopy() *Topology {
	if in == nil {
		return nil
	}
	out := new(Topology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeConfig) DeepCopyInto(out *VolumeConfig) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = make([]VolumeReplica, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeConfig.
func (in *VolumeConfig) DeepCopy() *VolumeConfig {
	if in == nil {
		return nil
	}
	out := new(VolumeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeInfo) DeepCopyInto(out *VolumeInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeInfo.
func (in *VolumeInfo) DeepCopy() *VolumeInfo {
	if in == nil {
		return nil
	}
	out := new(VolumeInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeReplica) DeepCopyInto(out *VolumeReplica) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeReplica.
func (in *VolumeReplica) DeepCopy() *VolumeReplica {
	if in == nil {
		return nil
	}
	out := new(VolumeReplica)
	in.DeepCopyInto(out)
	return out
}
