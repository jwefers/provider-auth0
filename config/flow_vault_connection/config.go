package flowvaultconnection

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_flow_vault_connection resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_flow_vault_connection", func(r *config.Resource) {
		r.ShortGroup = "flowvaultconnection"
	})
}
