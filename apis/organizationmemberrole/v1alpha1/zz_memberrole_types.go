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

type MemberRoleInitParameters struct {

	// (String) The ID of the organization.
	// The ID of the organization.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (String) The role ID to assign to the organization member.
	// The role ID to assign to the organization member.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (String) The user ID of the organization member.
	// The user ID of the organization member.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type MemberRoleObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The ID of the organization.
	// The ID of the organization.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (String) Description of the role.
	// Description of the role.
	RoleDescription *string `json:"roleDescription,omitempty" tf:"role_description,omitempty"`

	// (String) The role ID to assign to the organization member.
	// The role ID to assign to the organization member.
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (String) Name of the role.
	// Name of the role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// (String) The user ID of the organization member.
	// The user ID of the organization member.
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type MemberRoleParameters struct {

	// (String) The ID of the organization.
	// The ID of the organization.
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// (String) The role ID to assign to the organization member.
	// The role ID to assign to the organization member.
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (String) The user ID of the organization member.
	// The user ID of the organization member.
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// MemberRoleSpec defines the desired state of MemberRole
type MemberRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MemberRoleParameters `json:"forProvider"`
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
	InitProvider MemberRoleInitParameters `json:"initProvider,omitempty"`
}

// MemberRoleStatus defines the observed state of MemberRole.
type MemberRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MemberRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MemberRole is the Schema for the MemberRoles API. This resource is used to manage the roles assigned to an organization member.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,auth0}
type MemberRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.organizationId) || (has(self.initProvider) && has(self.initProvider.organizationId))",message="spec.forProvider.organizationId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleId) || (has(self.initProvider) && has(self.initProvider.roleId))",message="spec.forProvider.roleId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userId) || (has(self.initProvider) && has(self.initProvider.userId))",message="spec.forProvider.userId is a required parameter"
	Spec   MemberRoleSpec   `json:"spec"`
	Status MemberRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MemberRoleList contains a list of MemberRoles
type MemberRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MemberRole `json:"items"`
}

// Repository type metadata.
var (
	MemberRole_Kind             = "MemberRole"
	MemberRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MemberRole_Kind}.String()
	MemberRole_KindAPIVersion   = MemberRole_Kind + "." + CRDGroupVersion.String()
	MemberRole_GroupVersionKind = CRDGroupVersion.WithKind(MemberRole_Kind)
)

func init() {
	SchemeBuilder.Register(&MemberRole{}, &MemberRoleList{})
}
