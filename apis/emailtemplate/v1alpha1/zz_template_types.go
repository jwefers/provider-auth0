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

type TemplateInitParameters struct {

	// (String) Body of the email template. You can include common variables.
	// Body of the email template. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (Boolean) Indicates whether the template is enabled.
	// Indicates whether the template is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Email address to use as the sender. You can include common variables.
	// Email address to use as the sender. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (Boolean) Whether the reset_email and verify_email templates should include the user's email address as the email parameter in the returnUrl (true) or whether no email address should be included in the redirect (false). Defaults to true.
	// Whether the `reset_email` and `verify_email` templates should include the user's email address as the email parameter in the `returnUrl` (true) or whether no email address should be included in the redirect (false). Defaults to `true`.
	IncludeEmailInRedirect *bool `json:"includeEmailInRedirect,omitempty" tf:"include_email_in_redirect,omitempty"`

	// (String) URL to redirect the user to after a successful action. Learn more.
	// URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/customize/email/email-templates#configure-template-fields).
	ResultURL *string `json:"resultUrl,omitempty" tf:"result_url,omitempty"`

	// (String) Subject line of the email. You can include common variables.
	// Subject line of the email. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String) Syntax of the template body. You can use either text or HTML with Liquid syntax.
	// Syntax of the template body. You can use either text or HTML with Liquid syntax.
	Syntax *string `json:"syntax,omitempty" tf:"syntax,omitempty"`

	// (String) Template name. Options include verify_email, verify_email_by_code, reset_email, reset_email_by_code, welcome_email, blocked_account, stolen_credentials, enrollment_email, mfa_oob_code, user_invitation, change_password (legacy), or password_reset (legacy).
	// Template name. Options include `verify_email`, `verify_email_by_code`, `reset_email`, `reset_email_by_code`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `user_invitation`, `change_password` (legacy), or `password_reset` (legacy).
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Number of seconds during which the link within the email will be valid.
	// Number of seconds during which the link within the email will be valid.
	URLLifetimeInSeconds *float64 `json:"urlLifetimeInSeconds,omitempty" tf:"url_lifetime_in_seconds,omitempty"`
}

type TemplateObservation struct {

	// (String) Body of the email template. You can include common variables.
	// Body of the email template. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (Boolean) Indicates whether the template is enabled.
	// Indicates whether the template is enabled.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Email address to use as the sender. You can include common variables.
	// Email address to use as the sender. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Boolean) Whether the reset_email and verify_email templates should include the user's email address as the email parameter in the returnUrl (true) or whether no email address should be included in the redirect (false). Defaults to true.
	// Whether the `reset_email` and `verify_email` templates should include the user's email address as the email parameter in the `returnUrl` (true) or whether no email address should be included in the redirect (false). Defaults to `true`.
	IncludeEmailInRedirect *bool `json:"includeEmailInRedirect,omitempty" tf:"include_email_in_redirect,omitempty"`

	// (String) URL to redirect the user to after a successful action. Learn more.
	// URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/customize/email/email-templates#configure-template-fields).
	ResultURL *string `json:"resultUrl,omitempty" tf:"result_url,omitempty"`

	// (String) Subject line of the email. You can include common variables.
	// Subject line of the email. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String) Syntax of the template body. You can use either text or HTML with Liquid syntax.
	// Syntax of the template body. You can use either text or HTML with Liquid syntax.
	Syntax *string `json:"syntax,omitempty" tf:"syntax,omitempty"`

	// (String) Template name. Options include verify_email, verify_email_by_code, reset_email, reset_email_by_code, welcome_email, blocked_account, stolen_credentials, enrollment_email, mfa_oob_code, user_invitation, change_password (legacy), or password_reset (legacy).
	// Template name. Options include `verify_email`, `verify_email_by_code`, `reset_email`, `reset_email_by_code`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `user_invitation`, `change_password` (legacy), or `password_reset` (legacy).
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Number of seconds during which the link within the email will be valid.
	// Number of seconds during which the link within the email will be valid.
	URLLifetimeInSeconds *float64 `json:"urlLifetimeInSeconds,omitempty" tf:"url_lifetime_in_seconds,omitempty"`
}

type TemplateParameters struct {

	// (String) Body of the email template. You can include common variables.
	// Body of the email template. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (Boolean) Indicates whether the template is enabled.
	// Indicates whether the template is enabled.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Email address to use as the sender. You can include common variables.
	// Email address to use as the sender. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from,omitempty"`

	// (Boolean) Whether the reset_email and verify_email templates should include the user's email address as the email parameter in the returnUrl (true) or whether no email address should be included in the redirect (false). Defaults to true.
	// Whether the `reset_email` and `verify_email` templates should include the user's email address as the email parameter in the `returnUrl` (true) or whether no email address should be included in the redirect (false). Defaults to `true`.
	// +kubebuilder:validation:Optional
	IncludeEmailInRedirect *bool `json:"includeEmailInRedirect,omitempty" tf:"include_email_in_redirect,omitempty"`

	// (String) URL to redirect the user to after a successful action. Learn more.
	// URL to redirect the user to after a successful action. [Learn more](https://auth0.com/docs/customize/email/email-templates#configure-template-fields).
	// +kubebuilder:validation:Optional
	ResultURL *string `json:"resultUrl,omitempty" tf:"result_url,omitempty"`

	// (String) Subject line of the email. You can include common variables.
	// Subject line of the email. You can include [common variables](https://auth0.com/docs/customize/email/email-templates#common-variables).
	// +kubebuilder:validation:Optional
	Subject *string `json:"subject,omitempty" tf:"subject,omitempty"`

	// (String) Syntax of the template body. You can use either text or HTML with Liquid syntax.
	// Syntax of the template body. You can use either text or HTML with Liquid syntax.
	// +kubebuilder:validation:Optional
	Syntax *string `json:"syntax,omitempty" tf:"syntax,omitempty"`

	// (String) Template name. Options include verify_email, verify_email_by_code, reset_email, reset_email_by_code, welcome_email, blocked_account, stolen_credentials, enrollment_email, mfa_oob_code, user_invitation, change_password (legacy), or password_reset (legacy).
	// Template name. Options include `verify_email`, `verify_email_by_code`, `reset_email`, `reset_email_by_code`, `welcome_email`, `blocked_account`, `stolen_credentials`, `enrollment_email`, `mfa_oob_code`, `user_invitation`, `change_password` (legacy), or `password_reset` (legacy).
	// +kubebuilder:validation:Optional
	Template *string `json:"template,omitempty" tf:"template,omitempty"`

	// (Number) Number of seconds during which the link within the email will be valid.
	// Number of seconds during which the link within the email will be valid.
	// +kubebuilder:validation:Optional
	URLLifetimeInSeconds *float64 `json:"urlLifetimeInSeconds,omitempty" tf:"url_lifetime_in_seconds,omitempty"`
}

// TemplateSpec defines the desired state of Template
type TemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TemplateParameters `json:"forProvider"`
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
	InitProvider TemplateInitParameters `json:"initProvider,omitempty"`
}

// TemplateStatus defines the observed state of Template.
type TemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Template is the Schema for the Templates API. With Auth0, you can have standard welcome, password reset, and account verification email-based workflows built right into Auth0. This resource allows you to configure email templates to customize the look, feel, and sender identities of emails sent by Auth0. Used in conjunction with configured email providers.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,auth0}
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.body) || (has(self.initProvider) && has(self.initProvider.body))",message="spec.forProvider.body is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.enabled) || (has(self.initProvider) && has(self.initProvider.enabled))",message="spec.forProvider.enabled is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.from) || (has(self.initProvider) && has(self.initProvider.from))",message="spec.forProvider.from is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subject) || (has(self.initProvider) && has(self.initProvider.subject))",message="spec.forProvider.subject is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.syntax) || (has(self.initProvider) && has(self.initProvider.syntax))",message="spec.forProvider.syntax is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.template) || (has(self.initProvider) && has(self.initProvider.template))",message="spec.forProvider.template is a required parameter"
	Spec   TemplateSpec   `json:"spec"`
	Status TemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TemplateList contains a list of Templates
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Template `json:"items"`
}

// Repository type metadata.
var (
	Template_Kind             = "Template"
	Template_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Template_Kind}.String()
	Template_KindAPIVersion   = Template_Kind + "." + CRDGroupVersion.String()
	Template_GroupVersionKind = CRDGroupVersion.WithKind(Template_Kind)
)

func init() {
	SchemeBuilder.Register(&Template{}, &TemplateList{})
}
