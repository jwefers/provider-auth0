package tenant

import "github.com/crossplane/upjet/pkg/config"

// Configure the auth0_tenant resource in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_tenant", func(r *config.Resource) {
		r.ShortGroup = "tenant"
	})
}
