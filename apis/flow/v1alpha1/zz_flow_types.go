// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FlowInitParameters struct {

	// (String) Actions of the flow.
	// Actions of the flow.
	Actions *string `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) Name of the flow.
	// Name of the flow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FlowObservation struct {

	// (String) Actions of the flow.
	// Actions of the flow.
	Actions *string `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Name of the flow.
	// Name of the flow.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type FlowParameters struct {

	// (String) Actions of the flow.
	// Actions of the flow.
	// +kubebuilder:validation:Optional
	Actions *string `json:"actions,omitempty" tf:"actions,omitempty"`

	// (String) Name of the flow.
	// Name of the flow.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// FlowSpec defines the desired state of Flow
type FlowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FlowParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FlowInitParameters `json:"initProvider,omitempty"`
}

// FlowStatus defines the observed state of Flow.
type FlowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FlowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Flow is the Schema for the Flows API. With this resource, you can create and manage Flows for a tenant.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,auth0}
type Flow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   FlowSpec   `json:"spec"`
	Status FlowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlowList contains a list of Flows
type FlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Flow `json:"items"`
}

// Repository type metadata.
var (
	Flow_Kind             = "Flow"
	Flow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Flow_Kind}.String()
	Flow_KindAPIVersion   = Flow_Kind + "." + CRDGroupVersion.String()
	Flow_GroupVersionKind = CRDGroupVersion.WithKind(Flow_Kind)
)

func init() {
	SchemeBuilder.Register(&Flow{}, &FlowList{})
}
