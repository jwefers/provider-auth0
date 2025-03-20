package connectionscimconfiguration

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_connection_scim_configuration resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_connection_scim_configuration", func(r *config.Resource) {
		r.ShortGroup = "connectionscimconfiguration"
	})
}
