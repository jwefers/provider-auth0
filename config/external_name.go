/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/crossplane/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"auth0_action":                           config.IdentifierFromProvider,
	"auth0_attack_protection":                config.IdentifierFromProvider,
	"auth0_branding":                         config.IdentifierFromProvider,
	"auth0_branding_theme":                   config.IdentifierFromProvider,
	"auth0_client":                           config.IdentifierFromProvider,
	"auth0_client_credentials":               config.IdentifierFromProvider,
	"auth0_client_grant":                     config.IdentifierFromProvider,
	"auth0_connection":                       config.IdentifierFromProvider,
	"auth0_connection_client":                config.IdentifierFromProvider,
	"auth0_connection_clients":               config.IdentifierFromProvider,
	"auth0_connection_scim_configuration":    config.IdentifierFromProvider,
	"auth0_custom_domain":                    config.IdentifierFromProvider,
	"auth0_custom_domain_verification":       config.IdentifierFromProvider,
	"auth0_email_provider":                   config.IdentifierFromProvider,
	"auth0_email_template":                   config.IdentifierFromProvider,
	"auth0_encryption_key_manager":           config.IdentifierFromProvider,
	"auth0_flow":                             config.IdentifierFromProvider,
	"auth0_flow_vault_connection":            config.IdentifierFromProvider,
	"auth0_form":                             config.IdentifierFromProvider,
	"auth0_guardian":                         config.IdentifierFromProvider,
	"auth0_hook":                             config.IdentifierFromProvider,
	"auth0_log_stream":                       config.IdentifierFromProvider,
	"auth0_organization":                     config.IdentifierFromProvider,
	"auth0_organization_client_grant":        config.IdentifierFromProvider,
	"auth0_organization_connection":          config.IdentifierFromProvider,
	"auth0_organization_connections":         config.IdentifierFromProvider,
	"auth0_organization_member":              config.IdentifierFromProvider,
	"auth0_organization_member_role":         config.IdentifierFromProvider,
	"auth0_organization_member_roles":        config.IdentifierFromProvider,
	"auth0_organization_members":             config.IdentifierFromProvider,
	"auth0_pages":                            config.IdentifierFromProvider,
	"auth0_phone_provider":                   config.IdentifierFromProvider,
	"auth0_prompt":                           config.IdentifierFromProvider,
	"auth0_prompt_custom_text":               config.IdentifierFromProvider,
	"auth0_prompt_partials":                  config.IdentifierFromProvider,
	"auth0_prompt_screen_partial":            config.IdentifierFromProvider,
	"auth0_prompt_screen_partials":           config.IdentifierFromProvider,
	"auth0_prompt_screen_renderer":           config.IdentifierFromProvider,
	"auth0_resource_server":                  config.IdentifierFromProvider,
	"auth0_resource_server_scope":            config.IdentifierFromProvider,
	"auth0_resource_server_scopes":           config.IdentifierFromProvider,
	"auth0_role":                             config.IdentifierFromProvider,
	"auth0_role_permission":                  config.IdentifierFromProvider,
	"auth0_role_permissions":                 config.IdentifierFromProvider,
	"auth0_rule":                             config.IdentifierFromProvider,
	"auth0_rule_config":                      config.IdentifierFromProvider,
	"auth0_self_service_profile":             config.IdentifierFromProvider,
	"auth0_self_service_profile_custom_text": config.IdentifierFromProvider,
	"auth0_tenant":                           config.IdentifierFromProvider,
	"auth0_token_exchange_profile":           config.IdentifierFromProvider,
	"auth0_trigger_action":                   config.IdentifierFromProvider,
	"auth0_trigger_actions":                  config.IdentifierFromProvider,
	"auth0_user":                             config.IdentifierFromProvider,
	"auth0_user_permission":                  config.IdentifierFromProvider,
	"auth0_user_permissions":                 config.IdentifierFromProvider,
	"auth0_user_role":                        config.IdentifierFromProvider,
	"auth0_user_roles":                       config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
