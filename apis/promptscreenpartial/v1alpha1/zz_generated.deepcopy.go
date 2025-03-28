//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsertionPointsInitParameters) DeepCopyInto(out *InsertionPointsInitParameters) {
	*out = *in
	if in.FormContent != nil {
		in, out := &in.FormContent, &out.FormContent
		*out = new(string)
		**out = **in
	}
	if in.FormContentEnd != nil {
		in, out := &in.FormContentEnd, &out.FormContentEnd
		*out = new(string)
		**out = **in
	}
	if in.FormContentStart != nil {
		in, out := &in.FormContentStart, &out.FormContentStart
		*out = new(string)
		**out = **in
	}
	if in.FormFooterEnd != nil {
		in, out := &in.FormFooterEnd, &out.FormFooterEnd
		*out = new(string)
		**out = **in
	}
	if in.FormFooterStart != nil {
		in, out := &in.FormFooterStart, &out.FormFooterStart
		*out = new(string)
		**out = **in
	}
	if in.SecondaryActionsEnd != nil {
		in, out := &in.SecondaryActionsEnd, &out.SecondaryActionsEnd
		*out = new(string)
		**out = **in
	}
	if in.SecondaryActionsStart != nil {
		in, out := &in.SecondaryActionsStart, &out.SecondaryActionsStart
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsertionPointsInitParameters.
func (in *InsertionPointsInitParameters) DeepCopy() *InsertionPointsInitParameters {
	if in == nil {
		return nil
	}
	out := new(InsertionPointsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsertionPointsObservation) DeepCopyInto(out *InsertionPointsObservation) {
	*out = *in
	if in.FormContent != nil {
		in, out := &in.FormContent, &out.FormContent
		*out = new(string)
		**out = **in
	}
	if in.FormContentEnd != nil {
		in, out := &in.FormContentEnd, &out.FormContentEnd
		*out = new(string)
		**out = **in
	}
	if in.FormContentStart != nil {
		in, out := &in.FormContentStart, &out.FormContentStart
		*out = new(string)
		**out = **in
	}
	if in.FormFooterEnd != nil {
		in, out := &in.FormFooterEnd, &out.FormFooterEnd
		*out = new(string)
		**out = **in
	}
	if in.FormFooterStart != nil {
		in, out := &in.FormFooterStart, &out.FormFooterStart
		*out = new(string)
		**out = **in
	}
	if in.SecondaryActionsEnd != nil {
		in, out := &in.SecondaryActionsEnd, &out.SecondaryActionsEnd
		*out = new(string)
		**out = **in
	}
	if in.SecondaryActionsStart != nil {
		in, out := &in.SecondaryActionsStart, &out.SecondaryActionsStart
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsertionPointsObservation.
func (in *InsertionPointsObservation) DeepCopy() *InsertionPointsObservation {
	if in == nil {
		return nil
	}
	out := new(InsertionPointsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InsertionPointsParameters) DeepCopyInto(out *InsertionPointsParameters) {
	*out = *in
	if in.FormContent != nil {
		in, out := &in.FormContent, &out.FormContent
		*out = new(string)
		**out = **in
	}
	if in.FormContentEnd != nil {
		in, out := &in.FormContentEnd, &out.FormContentEnd
		*out = new(string)
		**out = **in
	}
	if in.FormContentStart != nil {
		in, out := &in.FormContentStart, &out.FormContentStart
		*out = new(string)
		**out = **in
	}
	if in.FormFooterEnd != nil {
		in, out := &in.FormFooterEnd, &out.FormFooterEnd
		*out = new(string)
		**out = **in
	}
	if in.FormFooterStart != nil {
		in, out := &in.FormFooterStart, &out.FormFooterStart
		*out = new(string)
		**out = **in
	}
	if in.SecondaryActionsEnd != nil {
		in, out := &in.SecondaryActionsEnd, &out.SecondaryActionsEnd
		*out = new(string)
		**out = **in
	}
	if in.SecondaryActionsStart != nil {
		in, out := &in.SecondaryActionsStart, &out.SecondaryActionsStart
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InsertionPointsParameters.
func (in *InsertionPointsParameters) DeepCopy() *InsertionPointsParameters {
	if in == nil {
		return nil
	}
	out := new(InsertionPointsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScreenPartial) DeepCopyInto(out *ScreenPartial) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScreenPartial.
func (in *ScreenPartial) DeepCopy() *ScreenPartial {
	if in == nil {
		return nil
	}
	out := new(ScreenPartial)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScreenPartial) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScreenPartialInitParameters) DeepCopyInto(out *ScreenPartialInitParameters) {
	*out = *in
	if in.InsertionPoints != nil {
		in, out := &in.InsertionPoints, &out.InsertionPoints
		*out = make([]InsertionPointsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PromptType != nil {
		in, out := &in.PromptType, &out.PromptType
		*out = new(string)
		**out = **in
	}
	if in.ScreenName != nil {
		in, out := &in.ScreenName, &out.ScreenName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScreenPartialInitParameters.
func (in *ScreenPartialInitParameters) DeepCopy() *ScreenPartialInitParameters {
	if in == nil {
		return nil
	}
	out := new(ScreenPartialInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScreenPartialList) DeepCopyInto(out *ScreenPartialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScreenPartial, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScreenPartialList.
func (in *ScreenPartialList) DeepCopy() *ScreenPartialList {
	if in == nil {
		return nil
	}
	out := new(ScreenPartialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScreenPartialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScreenPartialObservation) DeepCopyInto(out *ScreenPartialObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.InsertionPoints != nil {
		in, out := &in.InsertionPoints, &out.InsertionPoints
		*out = make([]InsertionPointsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PromptType != nil {
		in, out := &in.PromptType, &out.PromptType
		*out = new(string)
		**out = **in
	}
	if in.ScreenName != nil {
		in, out := &in.ScreenName, &out.ScreenName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScreenPartialObservation.
func (in *ScreenPartialObservation) DeepCopy() *ScreenPartialObservation {
	if in == nil {
		return nil
	}
	out := new(ScreenPartialObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScreenPartialParameters) DeepCopyInto(out *ScreenPartialParameters) {
	*out = *in
	if in.InsertionPoints != nil {
		in, out := &in.InsertionPoints, &out.InsertionPoints
		*out = make([]InsertionPointsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PromptType != nil {
		in, out := &in.PromptType, &out.PromptType
		*out = new(string)
		**out = **in
	}
	if in.ScreenName != nil {
		in, out := &in.ScreenName, &out.ScreenName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScreenPartialParameters.
func (in *ScreenPartialParameters) DeepCopy() *ScreenPartialParameters {
	if in == nil {
		return nil
	}
	out := new(ScreenPartialParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScreenPartialSpec) DeepCopyInto(out *ScreenPartialSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScreenPartialSpec.
func (in *ScreenPartialSpec) DeepCopy() *ScreenPartialSpec {
	if in == nil {
		return nil
	}
	out := new(ScreenPartialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScreenPartialStatus) DeepCopyInto(out *ScreenPartialStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScreenPartialStatus.
func (in *ScreenPartialStatus) DeepCopy() *ScreenPartialStatus {
	if in == nil {
		return nil
	}
	out := new(ScreenPartialStatus)
	in.DeepCopyInto(out)
	return out
}
