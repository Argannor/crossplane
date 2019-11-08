// +build !ignore_autogenerated

/*
Copyright 2019 The Crossplane Authors.

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
	"encoding/json"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplication) DeepCopyInto(out *KubernetesApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplication.
func (in *KubernetesApplication) DeepCopy() *KubernetesApplication {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplicationList) DeepCopyInto(out *KubernetesApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubernetesApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplicationList.
func (in *KubernetesApplicationList) DeepCopy() *KubernetesApplicationList {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplicationResource) DeepCopyInto(out *KubernetesApplicationResource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplicationResource.
func (in *KubernetesApplicationResource) DeepCopy() *KubernetesApplicationResource {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplicationResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesApplicationResource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplicationResourceList) DeepCopyInto(out *KubernetesApplicationResourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubernetesApplicationResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplicationResourceList.
func (in *KubernetesApplicationResourceList) DeepCopy() *KubernetesApplicationResourceList {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplicationResourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubernetesApplicationResourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplicationResourceSpec) DeepCopyInto(out *KubernetesApplicationResourceSpec) {
	*out = *in
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = (*in).DeepCopy()
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplicationResourceSpec.
func (in *KubernetesApplicationResourceSpec) DeepCopy() *KubernetesApplicationResourceSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplicationResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplicationResourceStatus) DeepCopyInto(out *KubernetesApplicationResourceStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(KubernetesClusterReference)
		**out = **in
	}
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(RemoteStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplicationResourceStatus.
func (in *KubernetesApplicationResourceStatus) DeepCopy() *KubernetesApplicationResourceStatus {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplicationResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplicationResourceTemplate) DeepCopyInto(out *KubernetesApplicationResourceTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplicationResourceTemplate.
func (in *KubernetesApplicationResourceTemplate) DeepCopy() *KubernetesApplicationResourceTemplate {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplicationResourceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplicationSpec) DeepCopyInto(out *KubernetesApplicationSpec) {
	*out = *in
	if in.ResourceSelector != nil {
		in, out := &in.ResourceSelector, &out.ResourceSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterSelector != nil {
		in, out := &in.ClusterSelector, &out.ClusterSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourceTemplates != nil {
		in, out := &in.ResourceTemplates, &out.ResourceTemplates
		*out = make([]KubernetesApplicationResourceTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplicationSpec.
func (in *KubernetesApplicationSpec) DeepCopy() *KubernetesApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesApplicationStatus) DeepCopyInto(out *KubernetesApplicationStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(KubernetesClusterReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesApplicationStatus.
func (in *KubernetesApplicationStatus) DeepCopy() *KubernetesApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(KubernetesApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesClusterReference) DeepCopyInto(out *KubernetesClusterReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesClusterReference.
func (in *KubernetesClusterReference) DeepCopy() *KubernetesClusterReference {
	if in == nil {
		return nil
	}
	out := new(KubernetesClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteStatus) DeepCopyInto(out *RemoteStatus) {
	*out = *in
	if in.Raw != nil {
		in, out := &in.Raw, &out.Raw
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteStatus.
func (in *RemoteStatus) DeepCopy() *RemoteStatus {
	if in == nil {
		return nil
	}
	out := new(RemoteStatus)
	in.DeepCopyInto(out)
	return out
}
