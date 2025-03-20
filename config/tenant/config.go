package tenant

import "github.com/crossplane/upjet/pkg/config"

// Configures the resource auth0_tenant in the provider configuration.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("auth0_tenant", func(r *config.Resource) {
		r.ShortGroup = "tenant"
	})
}
