// +build !ignore_autogenerated

// Code generated by operator-sdk-v0.19.0-x86_64-linux-gnu. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDevice) DeepCopyInto(out *BlockDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDevice.
func (in *BlockDevice) DeepCopy() *BlockDevice {
	if in == nil {
		return nil
	}
	out := new(BlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceClaim) DeepCopyInto(out *BlockDeviceClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceClaim.
func (in *BlockDeviceClaim) DeepCopy() *BlockDeviceClaim {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockDeviceClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceClaimList) DeepCopyInto(out *BlockDeviceClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BlockDeviceClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceClaimList.
func (in *BlockDeviceClaimList) DeepCopy() *BlockDeviceClaimList {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockDeviceClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceList) DeepCopyInto(out *BlockDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BlockDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceList.
func (in *BlockDeviceList) DeepCopy() *BlockDeviceList {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BlockDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceNodeAttributes) DeepCopyInto(out *BlockDeviceNodeAttributes) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceNodeAttributes.
func (in *BlockDeviceNodeAttributes) DeepCopy() *BlockDeviceNodeAttributes {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceNodeAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceCapacity) DeepCopyInto(out *DeviceCapacity) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceCapacity.
func (in *DeviceCapacity) DeepCopy() *DeviceCapacity {
	if in == nil {
		return nil
	}
	out := new(DeviceCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClaimDetails) DeepCopyInto(out *DeviceClaimDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClaimDetails.
func (in *DeviceClaimDetails) DeepCopy() *DeviceClaimDetails {
	if in == nil {
		return nil
	}
	out := new(DeviceClaimDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClaimResources) DeepCopyInto(out *DeviceClaimResources) {
	*out = *in
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClaimResources.
func (in *DeviceClaimResources) DeepCopy() *DeviceClaimResources {
	if in == nil {
		return nil
	}
	out := new(DeviceClaimResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClaimSpec) DeepCopyInto(out *DeviceClaimSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	out.Details = in.Details
	out.BlockDeviceNodeAttributes = in.BlockDeviceNodeAttributes
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClaimSpec.
func (in *DeviceClaimSpec) DeepCopy() *DeviceClaimSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceClaimStatus) DeepCopyInto(out *DeviceClaimStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceClaimStatus.
func (in *DeviceClaimStatus) DeepCopy() *DeviceClaimStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceDetails) DeepCopyInto(out *DeviceDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceDetails.
func (in *DeviceDetails) DeepCopy() *DeviceDetails {
	if in == nil {
		return nil
	}
	out := new(DeviceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceDevLink) DeepCopyInto(out *DeviceDevLink) {
	*out = *in
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceDevLink.
func (in *DeviceDevLink) DeepCopy() *DeviceDevLink {
	if in == nil {
		return nil
	}
	out := new(DeviceDevLink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceSpec) DeepCopyInto(out *DeviceSpec) {
	*out = *in
	out.Capacity = in.Capacity
	if in.ClaimRef != nil {
		in, out := &in.ClaimRef, &out.ClaimRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
	out.Details = in.Details
	if in.DevLinks != nil {
		in, out := &in.DevLinks, &out.DevLinks
		*out = make([]DeviceDevLink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.FileSystem = in.FileSystem
	out.NodeAttributes = in.NodeAttributes
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceSpec.
func (in *DeviceSpec) DeepCopy() *DeviceSpec {
	if in == nil {
		return nil
	}
	out := new(DeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeviceStatus) DeepCopyInto(out *DeviceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeviceStatus.
func (in *DeviceStatus) DeepCopy() *DeviceStatus {
	if in == nil {
		return nil
	}
	out := new(DeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileSystemInfo) DeepCopyInto(out *FileSystemInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileSystemInfo.
func (in *FileSystemInfo) DeepCopy() *FileSystemInfo {
	if in == nil {
		return nil
	}
	out := new(FileSystemInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAttribute) DeepCopyInto(out *NodeAttribute) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAttribute.
func (in *NodeAttribute) DeepCopy() *NodeAttribute {
	if in == nil {
		return nil
	}
	out := new(NodeAttribute)
	in.DeepCopyInto(out)
	return out
}
