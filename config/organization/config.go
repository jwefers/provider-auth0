package organization

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_organization resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_organization", func(r *config.Resource) {
		r.ShortGroup = "organization"
	})
}
