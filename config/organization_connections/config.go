package organizationconnections

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_organization_connections in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_connections", func(r *config.Resource) {
		r.ShortGroup = "organizationconnections"
	})
}
