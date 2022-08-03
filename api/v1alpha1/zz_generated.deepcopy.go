//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CEosLabDevice) DeepCopyInto(out *CEosLabDevice) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CEosLabDevice.
func (in *CEosLabDevice) DeepCopy() *CEosLabDevice {
	if in == nil {
		return nil
	}
	out := new(CEosLabDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CEosLabDevice) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CEosLabDeviceList) DeepCopyInto(out *CEosLabDeviceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CEosLabDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CEosLabDeviceList.
func (in *CEosLabDeviceList) DeepCopy() *CEosLabDeviceList {
	if in == nil {
		return nil
	}
	out := new(CEosLabDeviceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CEosLabDeviceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CEosLabDeviceSpec) DeepCopyInto(out *CEosLabDeviceSpec) {
	*out = *in
	if in.EnvVar != nil {
		in, out := &in.EnvVar, &out.EnvVar
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make(map[string]ServiceConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	in.CertConfig.DeepCopyInto(&out.CertConfig)
	if in.IntfMapping != nil {
		in, out := &in.IntfMapping, &out.IntfMapping
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ToggleOverrides != nil {
		in, out := &in.ToggleOverrides, &out.ToggleOverrides
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CEosLabDeviceSpec.
func (in *CEosLabDeviceSpec) DeepCopy() *CEosLabDeviceSpec {
	if in == nil {
		return nil
	}
	out := new(CEosLabDeviceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CEosLabDeviceStatus) DeepCopyInto(out *CEosLabDeviceStatus) {
	*out = *in
	in.CertConfig.DeepCopyInto(&out.CertConfig)
	if in.IntfMapping != nil {
		in, out := &in.IntfMapping, &out.IntfMapping
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ToggleOverrides != nil {
		in, out := &in.ToggleOverrides, &out.ToggleOverrides
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CEosLabDeviceStatus.
func (in *CEosLabDeviceStatus) DeepCopy() *CEosLabDeviceStatus {
	if in == nil {
		return nil
	}
	out := new(CEosLabDeviceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertConfig) DeepCopyInto(out *CertConfig) {
	*out = *in
	if in.SelfSignedCerts != nil {
		in, out := &in.SelfSignedCerts, &out.SelfSignedCerts
		*out = make([]SelfSignedCertConfig, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertConfig.
func (in *CertConfig) DeepCopy() *CertConfig {
	if in == nil {
		return nil
	}
	out := new(CertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSignedCertConfig) DeepCopyInto(out *SelfSignedCertConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSignedCertConfig.
func (in *SelfSignedCertConfig) DeepCopy() *SelfSignedCertConfig {
	if in == nil {
		return nil
	}
	out := new(SelfSignedCertConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceConfig) DeepCopyInto(out *ServiceConfig) {
	*out = *in
	if in.TCPPorts != nil {
		in, out := &in.TCPPorts, &out.TCPPorts
		*out = make([]uint32, len(*in))
		copy(*out, *in)
	}
	if in.UDPPorts != nil {
		in, out := &in.UDPPorts, &out.UDPPorts
		*out = make([]uint32, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceConfig.
func (in *ServiceConfig) DeepCopy() *ServiceConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceConfig)
	in.DeepCopyInto(out)
	return out
}
