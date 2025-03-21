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

type ScopesInitParameters struct {

	// friendly description of the scope (permission).
	// User-friendly description of the scope (permission).
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Name of the scope (permission). Examples include read:appointments or delete:appointments.
	// Name of the scope (permission). Examples include `read:appointments` or `delete:appointments`.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ScopesObservation struct {

	// friendly description of the scope (permission).
	// User-friendly description of the scope (permission).
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Name of the scope (permission). Examples include read:appointments or delete:appointments.
	// Name of the scope (permission). Examples include `read:appointments` or `delete:appointments`.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type ScopesParameters struct {

	// friendly description of the scope (permission).
	// User-friendly description of the scope (permission).
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) Name of the scope (permission). Examples include read:appointments or delete:appointments.
	// Name of the scope (permission). Examples include `read:appointments` or `delete:appointments`.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

type ServerScopesInitParameters struct {

	// (String) Identifier of the resource server that the scopes (permission) are associated with.
	// Identifier of the resource server that the scopes (permission) are associated with.
	ResourceServerIdentifier *string `json:"resourceServerIdentifier,omitempty" tf:"resource_server_identifier,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Scopes []ScopesInitParameters `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

type ServerScopesObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Identifier of the resource server that the scopes (permission) are associated with.
	// Identifier of the resource server that the scopes (permission) are associated with.
	ResourceServerIdentifier *string `json:"resourceServerIdentifier,omitempty" tf:"resource_server_identifier,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	Scopes []ScopesObservation `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

type ServerScopesParameters struct {

	// (String) Identifier of the resource server that the scopes (permission) are associated with.
	// Identifier of the resource server that the scopes (permission) are associated with.
	// +kubebuilder:validation:Optional
	ResourceServerIdentifier *string `json:"resourceServerIdentifier,omitempty" tf:"resource_server_identifier,omitempty"`

	// (Block Set, Min: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Scopes []ScopesParameters `json:"scopes,omitempty" tf:"scopes,omitempty"`
}

// ServerScopesSpec defines the desired state of ServerScopes
type ServerScopesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerScopesParameters `json:"forProvider"`
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
	InitProvider ServerScopesInitParameters `json:"initProvider,omitempty"`
}

// ServerScopesStatus defines the observed state of ServerScopes.
type ServerScopesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerScopesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServerScopes is the Schema for the ServerScopess API. With this resource, you can manage scopes (permissions) associated with a resource server (API).
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,auth0}
type ServerScopes struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.resourceServerIdentifier) || (has(self.initProvider) && has(self.initProvider.resourceServerIdentifier))",message="spec.forProvider.resourceServerIdentifier is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopes) || (has(self.initProvider) && has(self.initProvider.scopes))",message="spec.forProvider.scopes is a required parameter"
	Spec   ServerScopesSpec   `json:"spec"`
	Status ServerScopesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerScopesList contains a list of ServerScopess
type ServerScopesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerScopes `json:"items"`
}

// Repository type metadata.
var (
	ServerScopes_Kind             = "ServerScopes"
	ServerScopes_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerScopes_Kind}.String()
	ServerScopes_KindAPIVersion   = ServerScopes_Kind + "." + CRDGroupVersion.String()
	ServerScopes_GroupVersionKind = CRDGroupVersion.WithKind(ServerScopes_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerScopes{}, &ServerScopesList{})
}
