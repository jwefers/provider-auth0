package organizationclientgrant

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_organization_client_grant in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization_client_grant", func(r *config.Resource) {
		r.ShortGroup = "organizationclientgrant"
	})
}
