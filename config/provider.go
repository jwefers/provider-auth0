/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	action "github.com/jwefers/provider-auth0/config/action"
	attackprotection "github.com/jwefers/provider-auth0/config/attack_protection"
	branding "github.com/jwefers/provider-auth0/config/branding"
	brandingtheme "github.com/jwefers/provider-auth0/config/branding_theme"
	client "github.com/jwefers/provider-auth0/config/client"
	clientcredentials "github.com/jwefers/provider-auth0/config/client_credentials"
	clientgrant "github.com/jwefers/provider-auth0/config/client_grant"
	connection "github.com/jwefers/provider-auth0/config/connection"
	connectionclient "github.com/jwefers/provider-auth0/config/connection_client"
	connectionclients "github.com/jwefers/provider-auth0/config/connection_clients"
	connectionscimconfiguration "github.com/jwefers/provider-auth0/config/connection_scim_configuration"
	customdomain "github.com/jwefers/provider-auth0/config/custom_domain"
	customdomainverification "github.com/jwefers/provider-auth0/config/custom_domain_verification"
	emailprovider "github.com/jwefers/provider-auth0/config/email_provider"
	emailtemplate "github.com/jwefers/provider-auth0/config/email_template"
	encryptionkeymanager "github.com/jwefers/provider-auth0/config/encryption_key_manager"
	flow "github.com/jwefers/provider-auth0/config/flow"
	flowvaultconnection "github.com/jwefers/provider-auth0/config/flow_vault_connection"
	form "github.com/jwefers/provider-auth0/config/form"
	guardian "github.com/jwefers/provider-auth0/config/guardian"
	hook "github.com/jwefers/provider-auth0/config/hook"
	logstream "github.com/jwefers/provider-auth0/config/log_stream"
	organization "github.com/jwefers/provider-auth0/config/organization"
	organizationclientgrant "github.com/jwefers/provider-auth0/config/organization_client_grant"
	organizationconnection "github.com/jwefers/provider-auth0/config/organization_connection"
	organizationconnections "github.com/jwefers/provider-auth0/config/organization_connections"
	organizationmember "github.com/jwefers/provider-auth0/config/organization_member"
	organizationmemberrole "github.com/jwefers/provider-auth0/config/organization_member_role"
	organizationmemberroles "github.com/jwefers/provider-auth0/config/organization_member_roles"
	organizationmembers "github.com/jwefers/provider-auth0/config/organization_members"
	pages "github.com/jwefers/provider-auth0/config/pages"
	phoneprovider "github.com/jwefers/provider-auth0/config/phone_provider"
	prompt "github.com/jwefers/provider-auth0/config/prompt"
	promptcustomtext "github.com/jwefers/provider-auth0/config/prompt_custom_text"
	promptpartials "github.com/jwefers/provider-auth0/config/prompt_partials"
	promptscreenpartial "github.com/jwefers/provider-auth0/config/prompt_screen_partial"
	promptscreenpartials "github.com/jwefers/provider-auth0/config/prompt_screen_partials"
	promptscreenrenderer "github.com/jwefers/provider-auth0/config/prompt_screen_renderer"
	resourceserver "github.com/jwefers/provider-auth0/config/resource_server"
	resourceserverscope "github.com/jwefers/provider-auth0/config/resource_server_scope"
	resourceserverscopes "github.com/jwefers/provider-auth0/config/resource_server_scopes"
	role "github.com/jwefers/provider-auth0/config/role"
	rolepermission "github.com/jwefers/provider-auth0/config/role_permission"
	rolepermissions "github.com/jwefers/provider-auth0/config/role_permissions"
	rule "github.com/jwefers/provider-auth0/config/rule"
	ruleconfig "github.com/jwefers/provider-auth0/config/rule_config"
	selfserviceprofile "github.com/jwefers/provider-auth0/config/self_service_profile"
	selfserviceprofilecustomtext "github.com/jwefers/provider-auth0/config/self_service_profile_custom_text"
	tenant "github.com/jwefers/provider-auth0/config/tenant"
	tokenexchangeprofile "github.com/jwefers/provider-auth0/config/token_exchange_profile"
	triggeraction "github.com/jwefers/provider-auth0/config/trigger_action"
	triggeractions "github.com/jwefers/provider-auth0/config/trigger_actions"
	user "github.com/jwefers/provider-auth0/config/user"
	userpermission "github.com/jwefers/provider-auth0/config/user_permission"
	userpermissions "github.com/jwefers/provider-auth0/config/user_permissions"
	userrole "github.com/jwefers/provider-auth0/config/user_role"
	userroles "github.com/jwefers/provider-auth0/config/user_roles"
)

const (
	resourcePrefix = "auth0"
	modulePath     = "github.com/jwefers/provider-auth0"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("auth0.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		action.Configure,
		attackprotection.Configure,
		branding.Configure,
		brandingtheme.Configure,
		client.Configure,
		clientcredentials.Configure,
		clientgrant.Configure,
		connection.Configure,
		connectionclient.Configure,
		connectionclients.Configure,
		connectionscimconfiguration.Configure,
		customdomain.Configure,
		customdomainverification.Configure,
		emailprovider.Configure,
		emailtemplate.Configure,
		encryptionkeymanager.Configure,
		flow.Configure,
		flowvaultconnection.Configure,
		form.Configure,
		guardian.Configure,
		hook.Configure,
		logstream.Configure,
		organization.Configure,
		organizationclientgrant.Configure,
		organizationconnection.Configure,
		organizationconnections.Configure,
		organizationmember.Configure,
		organizationmemberrole.Configure,
		organizationmemberroles.Configure,
		organizationmembers.Configure,
		pages.Configure,
		phoneprovider.Configure,
		prompt.Configure,
		promptcustomtext.Configure,
		promptpartials.Configure,
		promptscreenpartial.Configure,
		promptscreenpartials.Configure,
		promptscreenrenderer.Configure,
		resourceserver.Configure,
		resourceserverscope.Configure,
		resourceserverscopes.Configure,
		role.Configure,
		rolepermission.Configure,
		rolepermissions.Configure,
		rule.Configure,
		ruleconfig.Configure,
		selfserviceprofile.Configure,
		selfserviceprofilecustomtext.Configure,
		tenant.Configure,
		tokenexchangeprofile.Configure,
		triggeraction.Configure,
		triggeractions.Configure,
		user.Configure,
		userpermission.Configure,
		userpermissions.Configure,
		userrole.Configure,
		userroles.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
