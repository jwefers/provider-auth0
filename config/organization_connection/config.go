package organizationconnection

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_organization_connection in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_connection", func(r *config.Resource) {
		r.ShortGroup = "organizationconnection"
	})
}
