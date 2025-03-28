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

type BrandingInitParameters struct {

	// (Block List, Max: 1) Configuration settings for colors for branding. (see below for nested schema)
	// Configuration settings for colors for branding.
	Colors []ColorsInitParameters `json:"colors,omitempty" tf:"colors,omitempty"`

	// (String) URL of logo to display on login page.
	// URL of logo to display on login page.
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`
}

type BrandingObservation struct {

	// (Block List, Max: 1) Configuration settings for colors for branding. (see below for nested schema)
	// Configuration settings for colors for branding.
	Colors []ColorsObservation `json:"colors,omitempty" tf:"colors,omitempty"`

	// (String) URL of logo to display on login page.
	// URL of logo to display on login page.
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`
}

type BrandingParameters struct {

	// (Block List, Max: 1) Configuration settings for colors for branding. (see below for nested schema)
	// Configuration settings for colors for branding.
	// +kubebuilder:validation:Optional
	Colors []ColorsParameters `json:"colors,omitempty" tf:"colors,omitempty"`

	// (String) URL of logo to display on login page.
	// URL of logo to display on login page.
	// +kubebuilder:validation:Optional
	LogoURL *string `json:"logoUrl,omitempty" tf:"logo_url,omitempty"`
}

type ColorsInitParameters struct {

	// (String) Primary button background color in hexadecimal.
	// Primary button background color in hexadecimal.
	Primary *string `json:"primary,omitempty" tf:"primary,omitempty"`
}

type ColorsObservation struct {

	// (String) Primary button background color in hexadecimal.
	// Primary button background color in hexadecimal.
	Primary *string `json:"primary,omitempty" tf:"primary,omitempty"`
}

type ColorsParameters struct {

	// (String) Primary button background color in hexadecimal.
	// Primary button background color in hexadecimal.
	// +kubebuilder:validation:Optional
	Primary *string `json:"primary,omitempty" tf:"primary,omitempty"`
}

type ServiceProfileInitParameters struct {

	// Service SSO flow.
	// List of IdP strategies that will be shown to users during the Self-Service SSO flow.
	// +listType=set
	AllowedStrategies []*string `json:"allowedStrategies,omitempty" tf:"allowed_strategies,omitempty"`

	// (Block List, Max: 1) Field can be used to customize the look and feel of the wizard. (see below for nested schema)
	// Field can be used to customize the look and feel of the wizard.
	Branding []BrandingInitParameters `json:"branding,omitempty" tf:"branding,omitempty"`

	// service Profile
	// The description of the self-service Profile
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// service Profile
	// The name of the self-service Profile
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// SSO flow. The user will be prompted to map the attributes on their identity provider to ensure the specified attributes get passed to Auth0. (see below for nested schema)
	// This array stores the mapping information that will be shown to the user during the SS-SSO flow. The user will be prompted to map the attributes on their identity provider to ensure the specified attributes get passed to Auth0.
	UserAttributes []UserAttributesInitParameters `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`
}

type ServiceProfileObservation struct {

	// Service SSO flow.
	// List of IdP strategies that will be shown to users during the Self-Service SSO flow.
	// +listType=set
	AllowedStrategies []*string `json:"allowedStrategies,omitempty" tf:"allowed_strategies,omitempty"`

	// (Block List, Max: 1) Field can be used to customize the look and feel of the wizard. (see below for nested schema)
	// Field can be used to customize the look and feel of the wizard.
	Branding []BrandingObservation `json:"branding,omitempty" tf:"branding,omitempty"`

	// (String) The ISO 8601 formatted date the profile was created.
	// The ISO 8601 formatted date the profile was created.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// service Profile
	// The description of the self-service Profile
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// service Profile
	// The name of the self-service Profile
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The ISO 8601 formatted date the profile was updated.
	// The ISO 8601 formatted date the profile was updated.
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// SSO flow. The user will be prompted to map the attributes on their identity provider to ensure the specified attributes get passed to Auth0. (see below for nested schema)
	// This array stores the mapping information that will be shown to the user during the SS-SSO flow. The user will be prompted to map the attributes on their identity provider to ensure the specified attributes get passed to Auth0.
	UserAttributes []UserAttributesObservation `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`
}

type ServiceProfileParameters struct {

	// Service SSO flow.
	// List of IdP strategies that will be shown to users during the Self-Service SSO flow.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedStrategies []*string `json:"allowedStrategies,omitempty" tf:"allowed_strategies,omitempty"`

	// (Block List, Max: 1) Field can be used to customize the look and feel of the wizard. (see below for nested schema)
	// Field can be used to customize the look and feel of the wizard.
	// +kubebuilder:validation:Optional
	Branding []BrandingParameters `json:"branding,omitempty" tf:"branding,omitempty"`

	// service Profile
	// The description of the self-service Profile
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// service Profile
	// The name of the self-service Profile
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// SSO flow. The user will be prompted to map the attributes on their identity provider to ensure the specified attributes get passed to Auth0. (see below for nested schema)
	// This array stores the mapping information that will be shown to the user during the SS-SSO flow. The user will be prompted to map the attributes on their identity provider to ensure the specified attributes get passed to Auth0.
	// +kubebuilder:validation:Optional
	UserAttributes []UserAttributesParameters `json:"userAttributes,omitempty" tf:"user_attributes,omitempty"`
}

type UserAttributesInitParameters struct {

	// service Profile
	// A human readable description of the attribute.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Indicates if this attribute is optional or if it has to be provided by the customer for the application to function.
	// Indicates if this attribute is optional or if it has to be provided by the customer for the application to function.
	IsOptional *bool `json:"isOptional,omitempty" tf:"is_optional,omitempty"`

	// service Profile
	// Attribute’s name on Auth0 side
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type UserAttributesObservation struct {

	// service Profile
	// A human readable description of the attribute.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Indicates if this attribute is optional or if it has to be provided by the customer for the application to function.
	// Indicates if this attribute is optional or if it has to be provided by the customer for the application to function.
	IsOptional *bool `json:"isOptional,omitempty" tf:"is_optional,omitempty"`

	// service Profile
	// Attribute’s name on Auth0 side
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type UserAttributesParameters struct {

	// service Profile
	// A human readable description of the attribute.
	// +kubebuilder:validation:Optional
	Description *string `json:"description" tf:"description,omitempty"`

	// (Boolean) Indicates if this attribute is optional or if it has to be provided by the customer for the application to function.
	// Indicates if this attribute is optional or if it has to be provided by the customer for the application to function.
	// +kubebuilder:validation:Optional
	IsOptional *bool `json:"isOptional" tf:"is_optional,omitempty"`

	// service Profile
	// Attribute’s name on Auth0 side
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`
}

// ServiceProfileSpec defines the desired state of ServiceProfile
type ServiceProfileSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceProfileParameters `json:"forProvider"`
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
	InitProvider ServiceProfileInitParameters `json:"initProvider,omitempty"`
}

// ServiceProfileStatus defines the observed state of ServiceProfile.
type ServiceProfileStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceProfileObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceProfile is the Schema for the ServiceProfiles API. With this resource, you can create and manage Self-Service Profile for a tenant.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,auth0}
type ServiceProfile struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ServiceProfileSpec   `json:"spec"`
	Status ServiceProfileStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceProfileList contains a list of ServiceProfiles
type ServiceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceProfile `json:"items"`
}

// Repository type metadata.
var (
	ServiceProfile_Kind             = "ServiceProfile"
	ServiceProfile_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceProfile_Kind}.String()
	ServiceProfile_KindAPIVersion   = ServiceProfile_Kind + "." + CRDGroupVersion.String()
	ServiceProfile_GroupVersionKind = CRDGroupVersion.WithKind(ServiceProfile_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceProfile{}, &ServiceProfileList{})
}
